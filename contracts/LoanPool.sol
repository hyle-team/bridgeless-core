// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {AbstractCompoundRateKeeper} from "@solarity/solidity-lib/finance/compound-rate-keeper/AbstractCompoundRateKeeper.sol";
import {DecimalsConverter} from "@solarity/solidity-lib/libs/utils/DecimalsConverter.sol";
import {PRECISION, PERCENTAGE_100} from "@solarity/solidity-lib/utils/Globals.sol";

import {ILoanPool} from "./interfaces/ILoanPool.sol";
import {IOracle} from "./interfaces/IOracle.sol";

contract LoanPool is
ILoanPool,
UUPSUpgradeable,
OwnableUpgradeable,
PausableUpgradeable,
ReentrancyGuardUpgradeable,
AbstractCompoundRateKeeper
{
    using SafeERC20 for IERC20;
    using DecimalsConverter for uint256;
    using Address for address payable;
    using Math for uint256;

    address public constant ETHEREUM_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    address public supplyToken;
    address public borrowToken;

    address public oracle;
    address public treasury;
    address public liquidator;

    PoolData public poolData;

    mapping(address user => UserData) internal _users;

    constructor() {
        _disableInitializers();
    }

    function __LoanPool_init(LoanPoolInitData calldata initData_) external initializer {
        __UUPSUpgradeable_init();
        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();
        __CompoundRateKeeper_init(initData_.capitalizationRate, initData_.capitalizationPeriod);

        setPoolData(initData_.poolData);
        setOracle(initData_.oracle);
        setTreasury(initData_.treasury);
        setLiquidator(initData_.liquidator);

        supplyToken = initData_.supplyToken;
        borrowToken = initData_.borrowToken;
    }

    receive() external payable {}

    function setPoolData(PoolData calldata poolData_) public onlyOwner {
        require(
            poolData_.collateralRatio <= PERCENTAGE_100,
            "LoanPool: Invalid collateral ratio."
        );

        poolData = poolData_;

        emit PoolDataSet(poolData_);
    }

    function setOracle(address oracle_) public onlyOwner {
        oracle = oracle_;

        emit OracleSet(oracle_);
    }

    function setTreasury(address treasury_) public onlyOwner {
        treasury = treasury_;

        emit TreasurySet(treasury_);
    }

    function setLiquidator(address liquidator_) public onlyOwner {
        liquidator = liquidator_;

        emit LiquidatorSet(liquidator_);
    }

    function setCapitalizationRate(uint256 capitalizationRate_) external onlyOwner {
        _setCapitalizationRate(capitalizationRate_);
    }

    function setCapitalizationPeriod(uint64 capitalizationPeriod_) external onlyOwner {
        _setCapitalizationPeriod(capitalizationPeriod_);
    }

    function deposit(uint256 amount_) external payable nonReentrant {
        _validateUserInput(amount_);

        address user_ = _msgSender();

        _updateUserPosition(user_);

        _receiveFunds(supplyToken, user_, amount_);

        _users[user_].deposited += amount_;

        emit Deposited(user_, amount_);
    }

    function withdraw(uint256 amount_) external nonReentrant {
        address user_ = _msgSender();

        _updateUserPosition(user_);

        UserData storage userData = _users[user_];

        amount_ = amount_.min(userData.deposited);
        _validateUserInput(amount_);

        require(amount_ <= getMaxPossibleWithdrawAmount(user_), "LoanPool: Not enough balance.");

        userData.deposited -= amount_;

        _sendFunds(supplyToken, user_, amount_);

        emit Withdrawn(user_, amount_);
    }

    function borrow(uint256 amount_, address recipient_) external payable nonReentrant {
        _validateUserInput(amount_);

        address user_ = _msgSender();

        _updateUserPosition(user_);

        UserData storage userData = _users[user_];

        require(amount_ <= getMaxPossibleBorrowAmount(user_), "LoanPool: Unsafe borrow amount.");

        userData.borrowed += amount_;
        userData.debtNormalized += (amount_ * PRECISION) / getCompoundRate();

        _sendFunds(borrowToken, recipient_, amount_);

        emit Borrowed(user_, recipient_, amount_);
    }

    function repay(uint256 amount_) external payable nonReentrant {
        address user_ = _msgSender();

        _updateUserPosition(user_);

        UserData storage userData = _users[user_];

        uint256 borrowed_ = userData.borrowed;

        amount_ = amount_.min(userData.borrowed);

        _validateUserInput(amount_);

        _receiveFunds(borrowToken, user_, amount_);

        uint256 newBorrowed_ = borrowed_ - amount_;

        userData.borrowed = newBorrowed_;
        if (newBorrowed_ == 0) {
            userData.borrowInterestFeePaid = 0;
            userData.debtNormalized = 0;
        } else {
            userData.debtNormalized -= (newBorrowed_ * PRECISION) / getCompoundRate();
        }

        emit Repaid(user_, amount_);
    }

    function liquidate(address user_, address recipient_) external payable nonReentrant {
        require(liquidator == _msgSender(), "LoanPool: Not a liquidator.");

        _updateUserPosition(user_);

        require(canUserBeLiquidated(user_), "LoanPool: User can't be liquidated.");

        UserData storage userData = _users[user_];

        uint256 borrowed_ = userData.borrowed;

        _receiveFunds(borrowToken, liquidator, borrowed_);

        _sendFunds(supplyToken, recipient_, userData.deposited);

        delete _users[user_];

        emit Liquidated(user_, recipient_, borrowed_);
    }

    function updateUserPosition(address user_) public nonReentrant {
        _updateUserPosition(user_);
    }

    function getUserDebt(address user_) public view returns (uint256) {
        UserData storage userData = _users[user_];

        uint256 debt_ = (userData.debtNormalized * getCompoundRate()) /
                    PRECISION -
                        userData.borrowInterestFeePaid;

        if (debt_ < userData.borrowed) {
            debt_ = userData.borrowed;
        }

        return debt_;
    }

    function getMaxPossibleBorrowAmount(address user_) public view returns (uint256) {
        uint256 freeCollateral_ = (getMaxPossibleWithdrawAmount(user_) *
            poolData.collateralRatio) / PERCENTAGE_100;
        (uint256 freeCollateralInBorrowCurrency_, ) = IOracle(oracle).getPrice(
            supplyToken,
            freeCollateral_
        );

        return freeCollateralInBorrowCurrency_;
    }

    function getMaxPossibleWithdrawAmount(address user_) public view returns (uint256) {
        uint256 collateral_ = _users[user_].deposited;

        uint256 requiredCollateral_ = _getRequiredCollateral(user_);

        if (collateral_ <= requiredCollateral_) {
            return 0;
        }

        return collateral_ - requiredCollateral_;
    }

    function canUserBeLiquidated(address user_) public view returns (bool) {
        return _getRequiredCollateral(user_) > _users[user_].deposited;
    }

    function _updateUserPosition(address user_) internal {
        UserData storage userData = _users[user_];

        userData.lastUpdate = block.timestamp;

        (uint256 fee_, uint256 feeInSupplyCurrency_) = _feeToPay(user_);
        if (feeInSupplyCurrency_ == 0) {
            return;
        }

        userData.borrowInterestFeePaid += fee_;
        userData.deposited -= feeInSupplyCurrency_;

        _sendFunds(supplyToken, treasury, feeInSupplyCurrency_);

        emit FeePaid(user_, feeInSupplyCurrency_);
    }

    function _sendFunds(address token_, address recipient_, uint256 amount_) internal {
        _validateUserInput(amount_);

        if (token_ == ETHEREUM_ADDRESS) {
            payable(recipient_).sendValue(amount_);
        } else {
            uint256 amountWithDecimals_ = amount_.from18Safe(token_);

            IERC20(token_).safeTransfer(recipient_, amountWithDecimals_);
        }
    }

    function _receiveFunds(address token_, address from_, uint256 amount_) internal {
        if (token_ == ETHEREUM_ADDRESS) {
            require(msg.value >= amount_, "LoanPool: Invalid native amount.");

            uint256 excessAmount_ = msg.value - amount_;
            if (excessAmount_ > 0) {
                payable(from_).sendValue(excessAmount_);
            }
        } else {
            require(msg.value == 0, "LoanPool: Invalid native amount.");

            uint256 amountWithDecimals_ = amount_.from18Safe(token_);

            IERC20(token_).safeTransferFrom(from_, address(this), amountWithDecimals_);
        }
    }

    function users(address user_) external view returns (UserData memory) {
        return _users[user_];
    }

    function _feeToPay(
        address user_
    ) internal view returns (uint256 fee_, uint256 feeInSupplyCurrency_) {
        fee_ = getUserDebt(user_) - _users[user_].borrowed;

        if (fee_ == 0) {
            return (0, 0);
        }

        (feeInSupplyCurrency_, ) = IOracle(oracle).getPrice(borrowToken, fee_);
    }

    function _getRequiredCollateral(address user_) internal view returns (uint256) {
        uint256 borrowed_ = _users[user_].borrowed;
        if (borrowed_ == 0) {
            return 0;
        }

        (uint256 borrowedInSupplyCurrency_, ) = IOracle(oracle).getPrice(borrowToken, borrowed_);

        (, uint256 feeToPayInSupplyCurrency_) = _feeToPay(user_);

        return
            (borrowedInSupplyCurrency_ * PERCENTAGE_100) /
            poolData.collateralRatio +
            feeToPayInSupplyCurrency_;
    }

    function _validateUserInput(uint256 amount_) internal pure {
        require(amount_ != 0, "LoanPool: Invalid amount.");
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}

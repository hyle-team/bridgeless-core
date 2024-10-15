// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import "@solarity/solidity-lib/finance/compound-rate-keeper/presets/OwnableCompoundRateKeeper.sol";
import "@solarity/solidity-lib/utils/Globals.sol";

contract LoanPool is OwnableUpgradeable {
    struct LoanPositionInfo {
        uint256 lastAbsoluteLoanAmount;
        uint256 normalizedLoanAmount;
        uint256 lastUpdateTimestamp;
    }

    OwnableCompoundRateKeeper public rateKeeper;

    uint256 public nextPositionID;

    mapping (uint256 => LoanPositionInfo) internal _loanPositionsInfo;

    event PositionCreated(uint256 positionId);

    function initialize(
        uint64 capitalizationPeriod_,
        uint256 capitalizationRate_
    ) external initializer {
        __Ownable_init();

        rateKeeper = new OwnableCompoundRateKeeper();

        rateKeeper.__OwnableCompoundRateKeeper_init(capitalizationRate_, capitalizationPeriod_);
        rateKeeper.transferOwnership(msg.sender);
    }

    function createLoanPositions(uint256[] calldata loanTokensAmountArr_) external {
        for (uint256 i = 0; i < loanTokensAmountArr_.length; i++) {
            createLoanPosition(loanTokensAmountArr_[i]);
        }
    }

    function updatePositions(uint256[] calldata positionIds_) external {
        for (uint256 i = 0; i < positionIds_.length; i++) {
            updatePosition(positionIds_[i]);
        }
    }

    function getLoanPositionInfo(uint256 positionId_) external view returns (LoanPositionInfo memory) {
        return _loanPositionsInfo[positionId_];
    }

    function createLoanPosition(uint256 loanTokensAmount_) public {
        uint256 normalizedLoanAmount_ = loanTokensAmount_ * PERCENTAGE_100 / rateKeeper.getCompoundRate();

        uint256 newPositionId = nextPositionID++;

        _loanPositionsInfo[newPositionId] = LoanPositionInfo(
            loanTokensAmount_,
            normalizedLoanAmount_,
            block.timestamp
        );

        emit PositionCreated(newPositionId);
    }

    function updatePosition(uint256 positionId_) public {
        LoanPositionInfo storage _positionInfo = _loanPositionsInfo[positionId_];

        uint256 currentAbsoluteValue = _positionInfo.normalizedLoanAmount * rateKeeper.getCompoundRate() / PERCENTAGE_100;

        _positionInfo.lastAbsoluteLoanAmount = currentAbsoluteValue;
    }
}
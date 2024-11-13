// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package loanpool

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ILoanPoolLoanPoolInitData is an auto generated low-level Go binding around an user-defined struct.
type ILoanPoolLoanPoolInitData struct {
	PoolData             ILoanPoolPoolData
	SupplyToken          common.Address
	BorrowToken          common.Address
	Oracle               common.Address
	Treasury             common.Address
	Liquidator           common.Address
	CapitalizationRate   *big.Int
	CapitalizationPeriod uint64
}

// ILoanPoolPoolData is an auto generated low-level Go binding around an user-defined struct.
type ILoanPoolPoolData struct {
	CollateralRatio *big.Int
}

// ILoanPoolUserData is an auto generated low-level Go binding around an user-defined struct.
type ILoanPoolUserData struct {
	Deposited             *big.Int
	Borrowed              *big.Int
	DebtNormalized        *big.Int
	BorrowInterestFeePaid *big.Int
	LastUpdate            *big.Int
}

// LoanPoolMetaData contains all meta data concerning the LoanPool contract.
var LoanPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCapitalizationPeriod\",\"type\":\"uint256\"}],\"name\":\"CapitalizationPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCapitalizationRate\",\"type\":\"uint256\"}],\"name\":\"CapitalizationRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Liquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"LiquidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"OracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILoanPool.PoolData\",\"name\":\"poolData\",\"type\":\"tuple\"}],\"name\":\"PoolDataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ETHEREUM_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"}],\"internalType\":\"structILoanPool.PoolData\",\"name\":\"poolData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"supplyToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"capitalizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"capitalizationPeriod\",\"type\":\"uint64\"}],\"internalType\":\"structILoanPool.LoanPoolInitData\",\"name\":\"initData_\",\"type\":\"tuple\"}],\"name\":\"__LoanPool_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"canUserBeLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUpdateCompoundRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCapitalizationPeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"capitalizationPeriod_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCapitalizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"capitalizationRate_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentRate_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp_\",\"type\":\"uint64\"}],\"name\":\"getFutureCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsMaxRateReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isMaxRateReached_\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"lastUpdate_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getMaxPossibleBorrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getMaxPossibleWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getUserDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"capitalizationPeriod_\",\"type\":\"uint64\"}],\"name\":\"setCapitalizationPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capitalizationRate_\",\"type\":\"uint256\"}],\"name\":\"setCapitalizationRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator_\",\"type\":\"address\"}],\"name\":\"setLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"}],\"internalType\":\"structILoanPool.PoolData\",\"name\":\"poolData_\",\"type\":\"tuple\"}],\"name\":\"setPoolData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"updateUserPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtNormalized\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowInterestFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"internalType\":\"structILoanPool.UserData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LoanPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LoanPoolMetaData.ABI instead.
var LoanPoolABI = LoanPoolMetaData.ABI

// LoanPool is an auto generated Go binding around an Ethereum contract.
type LoanPool struct {
	LoanPoolCaller     // Read-only binding to the contract
	LoanPoolTransactor // Write-only binding to the contract
	LoanPoolFilterer   // Log filterer for contract events
}

// LoanPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoanPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoanPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoanPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoanPoolSession struct {
	Contract     *LoanPool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoanPoolCallerSession struct {
	Contract *LoanPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LoanPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoanPoolTransactorSession struct {
	Contract     *LoanPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LoanPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoanPoolRaw struct {
	Contract *LoanPool // Generic contract binding to access the raw methods on
}

// LoanPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoanPoolCallerRaw struct {
	Contract *LoanPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LoanPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoanPoolTransactorRaw struct {
	Contract *LoanPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoanPool creates a new instance of LoanPool, bound to a specific deployed contract.
func NewLoanPool(address common.Address, backend bind.ContractBackend) (*LoanPool, error) {
	contract, err := bindLoanPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoanPool{LoanPoolCaller: LoanPoolCaller{contract: contract}, LoanPoolTransactor: LoanPoolTransactor{contract: contract}, LoanPoolFilterer: LoanPoolFilterer{contract: contract}}, nil
}

// NewLoanPoolCaller creates a new read-only instance of LoanPool, bound to a specific deployed contract.
func NewLoanPoolCaller(address common.Address, caller bind.ContractCaller) (*LoanPoolCaller, error) {
	contract, err := bindLoanPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoanPoolCaller{contract: contract}, nil
}

// NewLoanPoolTransactor creates a new write-only instance of LoanPool, bound to a specific deployed contract.
func NewLoanPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LoanPoolTransactor, error) {
	contract, err := bindLoanPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoanPoolTransactor{contract: contract}, nil
}

// NewLoanPoolFilterer creates a new log filterer instance of LoanPool, bound to a specific deployed contract.
func NewLoanPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LoanPoolFilterer, error) {
	contract, err := bindLoanPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoanPoolFilterer{contract: contract}, nil
}

// bindLoanPool binds a generic wrapper to an already deployed contract.
func bindLoanPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoanPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanPool *LoanPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoanPool.Contract.LoanPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanPool *LoanPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanPool.Contract.LoanPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanPool *LoanPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanPool.Contract.LoanPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanPool *LoanPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoanPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanPool *LoanPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanPool *LoanPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanPool.Contract.contract.Transact(opts, method, params...)
}

// ETHEREUMADDRESS is a free data retrieval call binding the contract method 0x67779fe2.
//
// Solidity: function ETHEREUM_ADDRESS() view returns(address)
func (_LoanPool *LoanPoolCaller) ETHEREUMADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "ETHEREUM_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHEREUMADDRESS is a free data retrieval call binding the contract method 0x67779fe2.
//
// Solidity: function ETHEREUM_ADDRESS() view returns(address)
func (_LoanPool *LoanPoolSession) ETHEREUMADDRESS() (common.Address, error) {
	return _LoanPool.Contract.ETHEREUMADDRESS(&_LoanPool.CallOpts)
}

// ETHEREUMADDRESS is a free data retrieval call binding the contract method 0x67779fe2.
//
// Solidity: function ETHEREUM_ADDRESS() view returns(address)
func (_LoanPool *LoanPoolCallerSession) ETHEREUMADDRESS() (common.Address, error) {
	return _LoanPool.Contract.ETHEREUMADDRESS(&_LoanPool.CallOpts)
}

// BorrowToken is a free data retrieval call binding the contract method 0x456dc17a.
//
// Solidity: function borrowToken() view returns(address)
func (_LoanPool *LoanPoolCaller) BorrowToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "borrowToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowToken is a free data retrieval call binding the contract method 0x456dc17a.
//
// Solidity: function borrowToken() view returns(address)
func (_LoanPool *LoanPoolSession) BorrowToken() (common.Address, error) {
	return _LoanPool.Contract.BorrowToken(&_LoanPool.CallOpts)
}

// BorrowToken is a free data retrieval call binding the contract method 0x456dc17a.
//
// Solidity: function borrowToken() view returns(address)
func (_LoanPool *LoanPoolCallerSession) BorrowToken() (common.Address, error) {
	return _LoanPool.Contract.BorrowToken(&_LoanPool.CallOpts)
}

// CanUserBeLiquidated is a free data retrieval call binding the contract method 0xf5590562.
//
// Solidity: function canUserBeLiquidated(address user_) view returns(bool)
func (_LoanPool *LoanPoolCaller) CanUserBeLiquidated(opts *bind.CallOpts, user_ common.Address) (bool, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "canUserBeLiquidated", user_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanUserBeLiquidated is a free data retrieval call binding the contract method 0xf5590562.
//
// Solidity: function canUserBeLiquidated(address user_) view returns(bool)
func (_LoanPool *LoanPoolSession) CanUserBeLiquidated(user_ common.Address) (bool, error) {
	return _LoanPool.Contract.CanUserBeLiquidated(&_LoanPool.CallOpts, user_)
}

// CanUserBeLiquidated is a free data retrieval call binding the contract method 0xf5590562.
//
// Solidity: function canUserBeLiquidated(address user_) view returns(bool)
func (_LoanPool *LoanPoolCallerSession) CanUserBeLiquidated(user_ common.Address) (bool, error) {
	return _LoanPool.Contract.CanUserBeLiquidated(&_LoanPool.CallOpts, user_)
}

// GetCapitalizationPeriod is a free data retrieval call binding the contract method 0x179680fe.
//
// Solidity: function getCapitalizationPeriod() view returns(uint64 capitalizationPeriod_)
func (_LoanPool *LoanPoolCaller) GetCapitalizationPeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getCapitalizationPeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCapitalizationPeriod is a free data retrieval call binding the contract method 0x179680fe.
//
// Solidity: function getCapitalizationPeriod() view returns(uint64 capitalizationPeriod_)
func (_LoanPool *LoanPoolSession) GetCapitalizationPeriod() (uint64, error) {
	return _LoanPool.Contract.GetCapitalizationPeriod(&_LoanPool.CallOpts)
}

// GetCapitalizationPeriod is a free data retrieval call binding the contract method 0x179680fe.
//
// Solidity: function getCapitalizationPeriod() view returns(uint64 capitalizationPeriod_)
func (_LoanPool *LoanPoolCallerSession) GetCapitalizationPeriod() (uint64, error) {
	return _LoanPool.Contract.GetCapitalizationPeriod(&_LoanPool.CallOpts)
}

// GetCapitalizationRate is a free data retrieval call binding the contract method 0x912119af.
//
// Solidity: function getCapitalizationRate() view returns(uint256 capitalizationRate_)
func (_LoanPool *LoanPoolCaller) GetCapitalizationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getCapitalizationRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCapitalizationRate is a free data retrieval call binding the contract method 0x912119af.
//
// Solidity: function getCapitalizationRate() view returns(uint256 capitalizationRate_)
func (_LoanPool *LoanPoolSession) GetCapitalizationRate() (*big.Int, error) {
	return _LoanPool.Contract.GetCapitalizationRate(&_LoanPool.CallOpts)
}

// GetCapitalizationRate is a free data retrieval call binding the contract method 0x912119af.
//
// Solidity: function getCapitalizationRate() view returns(uint256 capitalizationRate_)
func (_LoanPool *LoanPoolCallerSession) GetCapitalizationRate() (*big.Int, error) {
	return _LoanPool.Contract.GetCapitalizationRate(&_LoanPool.CallOpts)
}

// GetCompoundRate is a free data retrieval call binding the contract method 0xd0f37f4c.
//
// Solidity: function getCompoundRate() view returns(uint256)
func (_LoanPool *LoanPoolCaller) GetCompoundRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getCompoundRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompoundRate is a free data retrieval call binding the contract method 0xd0f37f4c.
//
// Solidity: function getCompoundRate() view returns(uint256)
func (_LoanPool *LoanPoolSession) GetCompoundRate() (*big.Int, error) {
	return _LoanPool.Contract.GetCompoundRate(&_LoanPool.CallOpts)
}

// GetCompoundRate is a free data retrieval call binding the contract method 0xd0f37f4c.
//
// Solidity: function getCompoundRate() view returns(uint256)
func (_LoanPool *LoanPoolCallerSession) GetCompoundRate() (*big.Int, error) {
	return _LoanPool.Contract.GetCompoundRate(&_LoanPool.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256 currentRate_)
func (_LoanPool *LoanPoolCaller) GetCurrentRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getCurrentRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256 currentRate_)
func (_LoanPool *LoanPoolSession) GetCurrentRate() (*big.Int, error) {
	return _LoanPool.Contract.GetCurrentRate(&_LoanPool.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256 currentRate_)
func (_LoanPool *LoanPoolCallerSession) GetCurrentRate() (*big.Int, error) {
	return _LoanPool.Contract.GetCurrentRate(&_LoanPool.CallOpts)
}

// GetFutureCompoundRate is a free data retrieval call binding the contract method 0x7e7fc8eb.
//
// Solidity: function getFutureCompoundRate(uint64 timestamp_) view returns(uint256)
func (_LoanPool *LoanPoolCaller) GetFutureCompoundRate(opts *bind.CallOpts, timestamp_ uint64) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getFutureCompoundRate", timestamp_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFutureCompoundRate is a free data retrieval call binding the contract method 0x7e7fc8eb.
//
// Solidity: function getFutureCompoundRate(uint64 timestamp_) view returns(uint256)
func (_LoanPool *LoanPoolSession) GetFutureCompoundRate(timestamp_ uint64) (*big.Int, error) {
	return _LoanPool.Contract.GetFutureCompoundRate(&_LoanPool.CallOpts, timestamp_)
}

// GetFutureCompoundRate is a free data retrieval call binding the contract method 0x7e7fc8eb.
//
// Solidity: function getFutureCompoundRate(uint64 timestamp_) view returns(uint256)
func (_LoanPool *LoanPoolCallerSession) GetFutureCompoundRate(timestamp_ uint64) (*big.Int, error) {
	return _LoanPool.Contract.GetFutureCompoundRate(&_LoanPool.CallOpts, timestamp_)
}

// GetIsMaxRateReached is a free data retrieval call binding the contract method 0x6e8c78fd.
//
// Solidity: function getIsMaxRateReached() view returns(bool isMaxRateReached_)
func (_LoanPool *LoanPoolCaller) GetIsMaxRateReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getIsMaxRateReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsMaxRateReached is a free data retrieval call binding the contract method 0x6e8c78fd.
//
// Solidity: function getIsMaxRateReached() view returns(bool isMaxRateReached_)
func (_LoanPool *LoanPoolSession) GetIsMaxRateReached() (bool, error) {
	return _LoanPool.Contract.GetIsMaxRateReached(&_LoanPool.CallOpts)
}

// GetIsMaxRateReached is a free data retrieval call binding the contract method 0x6e8c78fd.
//
// Solidity: function getIsMaxRateReached() view returns(bool isMaxRateReached_)
func (_LoanPool *LoanPoolCallerSession) GetIsMaxRateReached() (bool, error) {
	return _LoanPool.Contract.GetIsMaxRateReached(&_LoanPool.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() view returns(uint64 lastUpdate_)
func (_LoanPool *LoanPoolCaller) GetLastUpdate(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getLastUpdate")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() view returns(uint64 lastUpdate_)
func (_LoanPool *LoanPoolSession) GetLastUpdate() (uint64, error) {
	return _LoanPool.Contract.GetLastUpdate(&_LoanPool.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() view returns(uint64 lastUpdate_)
func (_LoanPool *LoanPoolCallerSession) GetLastUpdate() (uint64, error) {
	return _LoanPool.Contract.GetLastUpdate(&_LoanPool.CallOpts)
}

// GetMaxPossibleBorrowAmount is a free data retrieval call binding the contract method 0x987b7605.
//
// Solidity: function getMaxPossibleBorrowAmount(address user_) view returns(uint256)
func (_LoanPool *LoanPoolCaller) GetMaxPossibleBorrowAmount(opts *bind.CallOpts, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getMaxPossibleBorrowAmount", user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPossibleBorrowAmount is a free data retrieval call binding the contract method 0x987b7605.
//
// Solidity: function getMaxPossibleBorrowAmount(address user_) view returns(uint256)
func (_LoanPool *LoanPoolSession) GetMaxPossibleBorrowAmount(user_ common.Address) (*big.Int, error) {
	return _LoanPool.Contract.GetMaxPossibleBorrowAmount(&_LoanPool.CallOpts, user_)
}

// GetMaxPossibleBorrowAmount is a free data retrieval call binding the contract method 0x987b7605.
//
// Solidity: function getMaxPossibleBorrowAmount(address user_) view returns(uint256)
func (_LoanPool *LoanPoolCallerSession) GetMaxPossibleBorrowAmount(user_ common.Address) (*big.Int, error) {
	return _LoanPool.Contract.GetMaxPossibleBorrowAmount(&_LoanPool.CallOpts, user_)
}

// GetMaxPossibleWithdrawAmount is a free data retrieval call binding the contract method 0x3d1b3d69.
//
// Solidity: function getMaxPossibleWithdrawAmount(address user_) view returns(uint256)
func (_LoanPool *LoanPoolCaller) GetMaxPossibleWithdrawAmount(opts *bind.CallOpts, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getMaxPossibleWithdrawAmount", user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPossibleWithdrawAmount is a free data retrieval call binding the contract method 0x3d1b3d69.
//
// Solidity: function getMaxPossibleWithdrawAmount(address user_) view returns(uint256)
func (_LoanPool *LoanPoolSession) GetMaxPossibleWithdrawAmount(user_ common.Address) (*big.Int, error) {
	return _LoanPool.Contract.GetMaxPossibleWithdrawAmount(&_LoanPool.CallOpts, user_)
}

// GetMaxPossibleWithdrawAmount is a free data retrieval call binding the contract method 0x3d1b3d69.
//
// Solidity: function getMaxPossibleWithdrawAmount(address user_) view returns(uint256)
func (_LoanPool *LoanPoolCallerSession) GetMaxPossibleWithdrawAmount(user_ common.Address) (*big.Int, error) {
	return _LoanPool.Contract.GetMaxPossibleWithdrawAmount(&_LoanPool.CallOpts, user_)
}

// GetUserDebt is a free data retrieval call binding the contract method 0xe65896c9.
//
// Solidity: function getUserDebt(address user_) view returns(uint256)
func (_LoanPool *LoanPoolCaller) GetUserDebt(opts *bind.CallOpts, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getUserDebt", user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDebt is a free data retrieval call binding the contract method 0xe65896c9.
//
// Solidity: function getUserDebt(address user_) view returns(uint256)
func (_LoanPool *LoanPoolSession) GetUserDebt(user_ common.Address) (*big.Int, error) {
	return _LoanPool.Contract.GetUserDebt(&_LoanPool.CallOpts, user_)
}

// GetUserDebt is a free data retrieval call binding the contract method 0xe65896c9.
//
// Solidity: function getUserDebt(address user_) view returns(uint256)
func (_LoanPool *LoanPoolCallerSession) GetUserDebt(user_ common.Address) (*big.Int, error) {
	return _LoanPool.Contract.GetUserDebt(&_LoanPool.CallOpts, user_)
}

// Liquidator is a free data retrieval call binding the contract method 0x4046ebae.
//
// Solidity: function liquidator() view returns(address)
func (_LoanPool *LoanPoolCaller) Liquidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "liquidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Liquidator is a free data retrieval call binding the contract method 0x4046ebae.
//
// Solidity: function liquidator() view returns(address)
func (_LoanPool *LoanPoolSession) Liquidator() (common.Address, error) {
	return _LoanPool.Contract.Liquidator(&_LoanPool.CallOpts)
}

// Liquidator is a free data retrieval call binding the contract method 0x4046ebae.
//
// Solidity: function liquidator() view returns(address)
func (_LoanPool *LoanPoolCallerSession) Liquidator() (common.Address, error) {
	return _LoanPool.Contract.Liquidator(&_LoanPool.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LoanPool *LoanPoolCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LoanPool *LoanPoolSession) Oracle() (common.Address, error) {
	return _LoanPool.Contract.Oracle(&_LoanPool.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_LoanPool *LoanPoolCallerSession) Oracle() (common.Address, error) {
	return _LoanPool.Contract.Oracle(&_LoanPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanPool *LoanPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanPool *LoanPoolSession) Owner() (common.Address, error) {
	return _LoanPool.Contract.Owner(&_LoanPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanPool *LoanPoolCallerSession) Owner() (common.Address, error) {
	return _LoanPool.Contract.Owner(&_LoanPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LoanPool *LoanPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LoanPool *LoanPoolSession) Paused() (bool, error) {
	return _LoanPool.Contract.Paused(&_LoanPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LoanPool *LoanPoolCallerSession) Paused() (bool, error) {
	return _LoanPool.Contract.Paused(&_LoanPool.CallOpts)
}

// PoolData is a free data retrieval call binding the contract method 0xfee151ae.
//
// Solidity: function poolData() view returns(uint256 collateralRatio)
func (_LoanPool *LoanPoolCaller) PoolData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "poolData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolData is a free data retrieval call binding the contract method 0xfee151ae.
//
// Solidity: function poolData() view returns(uint256 collateralRatio)
func (_LoanPool *LoanPoolSession) PoolData() (*big.Int, error) {
	return _LoanPool.Contract.PoolData(&_LoanPool.CallOpts)
}

// PoolData is a free data retrieval call binding the contract method 0xfee151ae.
//
// Solidity: function poolData() view returns(uint256 collateralRatio)
func (_LoanPool *LoanPoolCallerSession) PoolData() (*big.Int, error) {
	return _LoanPool.Contract.PoolData(&_LoanPool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LoanPool *LoanPoolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LoanPool *LoanPoolSession) ProxiableUUID() ([32]byte, error) {
	return _LoanPool.Contract.ProxiableUUID(&_LoanPool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LoanPool *LoanPoolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LoanPool.Contract.ProxiableUUID(&_LoanPool.CallOpts)
}

// SupplyToken is a free data retrieval call binding the contract method 0x52059756.
//
// Solidity: function supplyToken() view returns(address)
func (_LoanPool *LoanPoolCaller) SupplyToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "supplyToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupplyToken is a free data retrieval call binding the contract method 0x52059756.
//
// Solidity: function supplyToken() view returns(address)
func (_LoanPool *LoanPoolSession) SupplyToken() (common.Address, error) {
	return _LoanPool.Contract.SupplyToken(&_LoanPool.CallOpts)
}

// SupplyToken is a free data retrieval call binding the contract method 0x52059756.
//
// Solidity: function supplyToken() view returns(address)
func (_LoanPool *LoanPoolCallerSession) SupplyToken() (common.Address, error) {
	return _LoanPool.Contract.SupplyToken(&_LoanPool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_LoanPool *LoanPoolCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_LoanPool *LoanPoolSession) Treasury() (common.Address, error) {
	return _LoanPool.Contract.Treasury(&_LoanPool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_LoanPool *LoanPoolCallerSession) Treasury() (common.Address, error) {
	return _LoanPool.Contract.Treasury(&_LoanPool.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address user_) view returns((uint256,uint256,uint256,uint256,uint256))
func (_LoanPool *LoanPoolCaller) Users(opts *bind.CallOpts, user_ common.Address) (ILoanPoolUserData, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "users", user_)

	if err != nil {
		return *new(ILoanPoolUserData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILoanPoolUserData)).(*ILoanPoolUserData)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address user_) view returns((uint256,uint256,uint256,uint256,uint256))
func (_LoanPool *LoanPoolSession) Users(user_ common.Address) (ILoanPoolUserData, error) {
	return _LoanPool.Contract.Users(&_LoanPool.CallOpts, user_)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address user_) view returns((uint256,uint256,uint256,uint256,uint256))
func (_LoanPool *LoanPoolCallerSession) Users(user_ common.Address) (ILoanPoolUserData, error) {
	return _LoanPool.Contract.Users(&_LoanPool.CallOpts, user_)
}

// LoanPoolInit is a paid mutator transaction binding the contract method 0x85a37317.
//
// Solidity: function __LoanPool_init(((uint256),address,address,address,address,address,uint256,uint64) initData_) returns()
func (_LoanPool *LoanPoolTransactor) LoanPoolInit(opts *bind.TransactOpts, initData_ ILoanPoolLoanPoolInitData) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "__LoanPool_init", initData_)
}

// LoanPoolInit is a paid mutator transaction binding the contract method 0x85a37317.
//
// Solidity: function __LoanPool_init(((uint256),address,address,address,address,address,uint256,uint64) initData_) returns()
func (_LoanPool *LoanPoolSession) LoanPoolInit(initData_ ILoanPoolLoanPoolInitData) (*types.Transaction, error) {
	return _LoanPool.Contract.LoanPoolInit(&_LoanPool.TransactOpts, initData_)
}

// LoanPoolInit is a paid mutator transaction binding the contract method 0x85a37317.
//
// Solidity: function __LoanPool_init(((uint256),address,address,address,address,address,uint256,uint64) initData_) returns()
func (_LoanPool *LoanPoolTransactorSession) LoanPoolInit(initData_ ILoanPoolLoanPoolInitData) (*types.Transaction, error) {
	return _LoanPool.Contract.LoanPoolInit(&_LoanPool.TransactOpts, initData_)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount_, address recipient_) payable returns()
func (_LoanPool *LoanPoolTransactor) Borrow(opts *bind.TransactOpts, amount_ *big.Int, recipient_ common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "borrow", amount_, recipient_)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount_, address recipient_) payable returns()
func (_LoanPool *LoanPoolSession) Borrow(amount_ *big.Int, recipient_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.Borrow(&_LoanPool.TransactOpts, amount_, recipient_)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount_, address recipient_) payable returns()
func (_LoanPool *LoanPoolTransactorSession) Borrow(amount_ *big.Int, recipient_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.Borrow(&_LoanPool.TransactOpts, amount_, recipient_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) payable returns()
func (_LoanPool *LoanPoolTransactor) Deposit(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "deposit", amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) payable returns()
func (_LoanPool *LoanPoolSession) Deposit(amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Deposit(&_LoanPool.TransactOpts, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) payable returns()
func (_LoanPool *LoanPoolTransactorSession) Deposit(amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Deposit(&_LoanPool.TransactOpts, amount_)
}

// EmergencyUpdateCompoundRate is a paid mutator transaction binding the contract method 0x71faedef.
//
// Solidity: function emergencyUpdateCompoundRate() returns()
func (_LoanPool *LoanPoolTransactor) EmergencyUpdateCompoundRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "emergencyUpdateCompoundRate")
}

// EmergencyUpdateCompoundRate is a paid mutator transaction binding the contract method 0x71faedef.
//
// Solidity: function emergencyUpdateCompoundRate() returns()
func (_LoanPool *LoanPoolSession) EmergencyUpdateCompoundRate() (*types.Transaction, error) {
	return _LoanPool.Contract.EmergencyUpdateCompoundRate(&_LoanPool.TransactOpts)
}

// EmergencyUpdateCompoundRate is a paid mutator transaction binding the contract method 0x71faedef.
//
// Solidity: function emergencyUpdateCompoundRate() returns()
func (_LoanPool *LoanPoolTransactorSession) EmergencyUpdateCompoundRate() (*types.Transaction, error) {
	return _LoanPool.Contract.EmergencyUpdateCompoundRate(&_LoanPool.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address user_, address recipient_) payable returns()
func (_LoanPool *LoanPoolTransactor) Liquidate(opts *bind.TransactOpts, user_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "liquidate", user_, recipient_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address user_, address recipient_) payable returns()
func (_LoanPool *LoanPoolSession) Liquidate(user_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.Liquidate(&_LoanPool.TransactOpts, user_, recipient_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address user_, address recipient_) payable returns()
func (_LoanPool *LoanPoolTransactorSession) Liquidate(user_ common.Address, recipient_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.Liquidate(&_LoanPool.TransactOpts, user_, recipient_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoanPool *LoanPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoanPool *LoanPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _LoanPool.Contract.RenounceOwnership(&_LoanPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoanPool *LoanPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LoanPool.Contract.RenounceOwnership(&_LoanPool.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount_) payable returns()
func (_LoanPool *LoanPoolTransactor) Repay(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "repay", amount_)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount_) payable returns()
func (_LoanPool *LoanPoolSession) Repay(amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Repay(&_LoanPool.TransactOpts, amount_)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount_) payable returns()
func (_LoanPool *LoanPoolTransactorSession) Repay(amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Repay(&_LoanPool.TransactOpts, amount_)
}

// SetCapitalizationPeriod is a paid mutator transaction binding the contract method 0xf70f48a9.
//
// Solidity: function setCapitalizationPeriod(uint64 capitalizationPeriod_) returns()
func (_LoanPool *LoanPoolTransactor) SetCapitalizationPeriod(opts *bind.TransactOpts, capitalizationPeriod_ uint64) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "setCapitalizationPeriod", capitalizationPeriod_)
}

// SetCapitalizationPeriod is a paid mutator transaction binding the contract method 0xf70f48a9.
//
// Solidity: function setCapitalizationPeriod(uint64 capitalizationPeriod_) returns()
func (_LoanPool *LoanPoolSession) SetCapitalizationPeriod(capitalizationPeriod_ uint64) (*types.Transaction, error) {
	return _LoanPool.Contract.SetCapitalizationPeriod(&_LoanPool.TransactOpts, capitalizationPeriod_)
}

// SetCapitalizationPeriod is a paid mutator transaction binding the contract method 0xf70f48a9.
//
// Solidity: function setCapitalizationPeriod(uint64 capitalizationPeriod_) returns()
func (_LoanPool *LoanPoolTransactorSession) SetCapitalizationPeriod(capitalizationPeriod_ uint64) (*types.Transaction, error) {
	return _LoanPool.Contract.SetCapitalizationPeriod(&_LoanPool.TransactOpts, capitalizationPeriod_)
}

// SetCapitalizationRate is a paid mutator transaction binding the contract method 0x9aad5c9a.
//
// Solidity: function setCapitalizationRate(uint256 capitalizationRate_) returns()
func (_LoanPool *LoanPoolTransactor) SetCapitalizationRate(opts *bind.TransactOpts, capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "setCapitalizationRate", capitalizationRate_)
}

// SetCapitalizationRate is a paid mutator transaction binding the contract method 0x9aad5c9a.
//
// Solidity: function setCapitalizationRate(uint256 capitalizationRate_) returns()
func (_LoanPool *LoanPoolSession) SetCapitalizationRate(capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.SetCapitalizationRate(&_LoanPool.TransactOpts, capitalizationRate_)
}

// SetCapitalizationRate is a paid mutator transaction binding the contract method 0x9aad5c9a.
//
// Solidity: function setCapitalizationRate(uint256 capitalizationRate_) returns()
func (_LoanPool *LoanPoolTransactorSession) SetCapitalizationRate(capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.SetCapitalizationRate(&_LoanPool.TransactOpts, capitalizationRate_)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x01c76f81.
//
// Solidity: function setLiquidator(address liquidator_) returns()
func (_LoanPool *LoanPoolTransactor) SetLiquidator(opts *bind.TransactOpts, liquidator_ common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "setLiquidator", liquidator_)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x01c76f81.
//
// Solidity: function setLiquidator(address liquidator_) returns()
func (_LoanPool *LoanPoolSession) SetLiquidator(liquidator_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.SetLiquidator(&_LoanPool.TransactOpts, liquidator_)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x01c76f81.
//
// Solidity: function setLiquidator(address liquidator_) returns()
func (_LoanPool *LoanPoolTransactorSession) SetLiquidator(liquidator_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.SetLiquidator(&_LoanPool.TransactOpts, liquidator_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_LoanPool *LoanPoolTransactor) SetOracle(opts *bind.TransactOpts, oracle_ common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "setOracle", oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_LoanPool *LoanPoolSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.SetOracle(&_LoanPool.TransactOpts, oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_LoanPool *LoanPoolTransactorSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.SetOracle(&_LoanPool.TransactOpts, oracle_)
}

// SetPoolData is a paid mutator transaction binding the contract method 0xd46b6406.
//
// Solidity: function setPoolData((uint256) poolData_) returns()
func (_LoanPool *LoanPoolTransactor) SetPoolData(opts *bind.TransactOpts, poolData_ ILoanPoolPoolData) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "setPoolData", poolData_)
}

// SetPoolData is a paid mutator transaction binding the contract method 0xd46b6406.
//
// Solidity: function setPoolData((uint256) poolData_) returns()
func (_LoanPool *LoanPoolSession) SetPoolData(poolData_ ILoanPoolPoolData) (*types.Transaction, error) {
	return _LoanPool.Contract.SetPoolData(&_LoanPool.TransactOpts, poolData_)
}

// SetPoolData is a paid mutator transaction binding the contract method 0xd46b6406.
//
// Solidity: function setPoolData((uint256) poolData_) returns()
func (_LoanPool *LoanPoolTransactorSession) SetPoolData(poolData_ ILoanPoolPoolData) (*types.Transaction, error) {
	return _LoanPool.Contract.SetPoolData(&_LoanPool.TransactOpts, poolData_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_LoanPool *LoanPoolTransactor) SetTreasury(opts *bind.TransactOpts, treasury_ common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "setTreasury", treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_LoanPool *LoanPoolSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.SetTreasury(&_LoanPool.TransactOpts, treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_LoanPool *LoanPoolTransactorSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.SetTreasury(&_LoanPool.TransactOpts, treasury_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoanPool *LoanPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoanPool *LoanPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.TransferOwnership(&_LoanPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoanPool *LoanPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.TransferOwnership(&_LoanPool.TransactOpts, newOwner)
}

// UpdateUserPosition is a paid mutator transaction binding the contract method 0xea5cbe07.
//
// Solidity: function updateUserPosition(address user_) returns()
func (_LoanPool *LoanPoolTransactor) UpdateUserPosition(opts *bind.TransactOpts, user_ common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "updateUserPosition", user_)
}

// UpdateUserPosition is a paid mutator transaction binding the contract method 0xea5cbe07.
//
// Solidity: function updateUserPosition(address user_) returns()
func (_LoanPool *LoanPoolSession) UpdateUserPosition(user_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.UpdateUserPosition(&_LoanPool.TransactOpts, user_)
}

// UpdateUserPosition is a paid mutator transaction binding the contract method 0xea5cbe07.
//
// Solidity: function updateUserPosition(address user_) returns()
func (_LoanPool *LoanPoolTransactorSession) UpdateUserPosition(user_ common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.UpdateUserPosition(&_LoanPool.TransactOpts, user_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LoanPool *LoanPoolTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LoanPool *LoanPoolSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.UpgradeTo(&_LoanPool.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LoanPool *LoanPoolTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LoanPool.Contract.UpgradeTo(&_LoanPool.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LoanPool *LoanPoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LoanPool *LoanPoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LoanPool.Contract.UpgradeToAndCall(&_LoanPool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LoanPool *LoanPoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LoanPool.Contract.UpgradeToAndCall(&_LoanPool.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_LoanPool *LoanPoolTransactor) Withdraw(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "withdraw", amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_LoanPool *LoanPoolSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Withdraw(&_LoanPool.TransactOpts, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) returns()
func (_LoanPool *LoanPoolTransactorSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Withdraw(&_LoanPool.TransactOpts, amount_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LoanPool *LoanPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LoanPool *LoanPoolSession) Receive() (*types.Transaction, error) {
	return _LoanPool.Contract.Receive(&_LoanPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LoanPool *LoanPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _LoanPool.Contract.Receive(&_LoanPool.TransactOpts)
}

// LoanPoolAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LoanPool contract.
type LoanPoolAdminChangedIterator struct {
	Event *LoanPoolAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolAdminChanged represents a AdminChanged event raised by the LoanPool contract.
type LoanPoolAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LoanPool *LoanPoolFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LoanPoolAdminChangedIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LoanPoolAdminChangedIterator{contract: _LoanPool.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LoanPool *LoanPoolFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LoanPoolAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolAdminChanged)
				if err := _LoanPool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LoanPool *LoanPoolFilterer) ParseAdminChanged(log types.Log) (*LoanPoolAdminChanged, error) {
	event := new(LoanPoolAdminChanged)
	if err := _LoanPool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LoanPool contract.
type LoanPoolBeaconUpgradedIterator struct {
	Event *LoanPoolBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolBeaconUpgraded represents a BeaconUpgraded event raised by the LoanPool contract.
type LoanPoolBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LoanPool *LoanPoolFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LoanPoolBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolBeaconUpgradedIterator{contract: _LoanPool.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LoanPool *LoanPoolFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LoanPoolBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolBeaconUpgraded)
				if err := _LoanPool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LoanPool *LoanPoolFilterer) ParseBeaconUpgraded(log types.Log) (*LoanPoolBeaconUpgraded, error) {
	event := new(LoanPoolBeaconUpgraded)
	if err := _LoanPool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolBorrowedIterator is returned from FilterBorrowed and is used to iterate over the raw logs and unpacked data for Borrowed events raised by the LoanPool contract.
type LoanPoolBorrowedIterator struct {
	Event *LoanPoolBorrowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolBorrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolBorrowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolBorrowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolBorrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolBorrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolBorrowed represents a Borrowed event raised by the LoanPool contract.
type LoanPoolBorrowed struct {
	User      common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBorrowed is a free log retrieval operation binding the contract event 0x3fc499aeb0bb1cb58b6de8b02b3f86f4e7394e9690bef0110e32ced8a5631045.
//
// Solidity: event Borrowed(address indexed user, address recipient, uint256 amount)
func (_LoanPool *LoanPoolFilterer) FilterBorrowed(opts *bind.FilterOpts, user []common.Address) (*LoanPoolBorrowedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Borrowed", userRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolBorrowedIterator{contract: _LoanPool.contract, event: "Borrowed", logs: logs, sub: sub}, nil
}

// WatchBorrowed is a free log subscription operation binding the contract event 0x3fc499aeb0bb1cb58b6de8b02b3f86f4e7394e9690bef0110e32ced8a5631045.
//
// Solidity: event Borrowed(address indexed user, address recipient, uint256 amount)
func (_LoanPool *LoanPoolFilterer) WatchBorrowed(opts *bind.WatchOpts, sink chan<- *LoanPoolBorrowed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Borrowed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolBorrowed)
				if err := _LoanPool.contract.UnpackLog(event, "Borrowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrowed is a log parse operation binding the contract event 0x3fc499aeb0bb1cb58b6de8b02b3f86f4e7394e9690bef0110e32ced8a5631045.
//
// Solidity: event Borrowed(address indexed user, address recipient, uint256 amount)
func (_LoanPool *LoanPoolFilterer) ParseBorrowed(log types.Log) (*LoanPoolBorrowed, error) {
	event := new(LoanPoolBorrowed)
	if err := _LoanPool.contract.UnpackLog(event, "Borrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolCapitalizationPeriodChangedIterator is returned from FilterCapitalizationPeriodChanged and is used to iterate over the raw logs and unpacked data for CapitalizationPeriodChanged events raised by the LoanPool contract.
type LoanPoolCapitalizationPeriodChangedIterator struct {
	Event *LoanPoolCapitalizationPeriodChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolCapitalizationPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolCapitalizationPeriodChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolCapitalizationPeriodChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolCapitalizationPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolCapitalizationPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolCapitalizationPeriodChanged represents a CapitalizationPeriodChanged event raised by the LoanPool contract.
type LoanPoolCapitalizationPeriodChanged struct {
	NewCapitalizationPeriod *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterCapitalizationPeriodChanged is a free log retrieval operation binding the contract event 0x2bc43b269adc02ae2363dc39721192e387c6b28d764e0d64cbe3a772405fe3af.
//
// Solidity: event CapitalizationPeriodChanged(uint256 newCapitalizationPeriod)
func (_LoanPool *LoanPoolFilterer) FilterCapitalizationPeriodChanged(opts *bind.FilterOpts) (*LoanPoolCapitalizationPeriodChangedIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "CapitalizationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &LoanPoolCapitalizationPeriodChangedIterator{contract: _LoanPool.contract, event: "CapitalizationPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchCapitalizationPeriodChanged is a free log subscription operation binding the contract event 0x2bc43b269adc02ae2363dc39721192e387c6b28d764e0d64cbe3a772405fe3af.
//
// Solidity: event CapitalizationPeriodChanged(uint256 newCapitalizationPeriod)
func (_LoanPool *LoanPoolFilterer) WatchCapitalizationPeriodChanged(opts *bind.WatchOpts, sink chan<- *LoanPoolCapitalizationPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "CapitalizationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolCapitalizationPeriodChanged)
				if err := _LoanPool.contract.UnpackLog(event, "CapitalizationPeriodChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCapitalizationPeriodChanged is a log parse operation binding the contract event 0x2bc43b269adc02ae2363dc39721192e387c6b28d764e0d64cbe3a772405fe3af.
//
// Solidity: event CapitalizationPeriodChanged(uint256 newCapitalizationPeriod)
func (_LoanPool *LoanPoolFilterer) ParseCapitalizationPeriodChanged(log types.Log) (*LoanPoolCapitalizationPeriodChanged, error) {
	event := new(LoanPoolCapitalizationPeriodChanged)
	if err := _LoanPool.contract.UnpackLog(event, "CapitalizationPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolCapitalizationRateChangedIterator is returned from FilterCapitalizationRateChanged and is used to iterate over the raw logs and unpacked data for CapitalizationRateChanged events raised by the LoanPool contract.
type LoanPoolCapitalizationRateChangedIterator struct {
	Event *LoanPoolCapitalizationRateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolCapitalizationRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolCapitalizationRateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolCapitalizationRateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolCapitalizationRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolCapitalizationRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolCapitalizationRateChanged represents a CapitalizationRateChanged event raised by the LoanPool contract.
type LoanPoolCapitalizationRateChanged struct {
	NewCapitalizationRate *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCapitalizationRateChanged is a free log retrieval operation binding the contract event 0x2dc8d118d310506a292b6c0b094a1b967c084fa0a728952284baf5f91dd75c24.
//
// Solidity: event CapitalizationRateChanged(uint256 newCapitalizationRate)
func (_LoanPool *LoanPoolFilterer) FilterCapitalizationRateChanged(opts *bind.FilterOpts) (*LoanPoolCapitalizationRateChangedIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "CapitalizationRateChanged")
	if err != nil {
		return nil, err
	}
	return &LoanPoolCapitalizationRateChangedIterator{contract: _LoanPool.contract, event: "CapitalizationRateChanged", logs: logs, sub: sub}, nil
}

// WatchCapitalizationRateChanged is a free log subscription operation binding the contract event 0x2dc8d118d310506a292b6c0b094a1b967c084fa0a728952284baf5f91dd75c24.
//
// Solidity: event CapitalizationRateChanged(uint256 newCapitalizationRate)
func (_LoanPool *LoanPoolFilterer) WatchCapitalizationRateChanged(opts *bind.WatchOpts, sink chan<- *LoanPoolCapitalizationRateChanged) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "CapitalizationRateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolCapitalizationRateChanged)
				if err := _LoanPool.contract.UnpackLog(event, "CapitalizationRateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCapitalizationRateChanged is a log parse operation binding the contract event 0x2dc8d118d310506a292b6c0b094a1b967c084fa0a728952284baf5f91dd75c24.
//
// Solidity: event CapitalizationRateChanged(uint256 newCapitalizationRate)
func (_LoanPool *LoanPoolFilterer) ParseCapitalizationRateChanged(log types.Log) (*LoanPoolCapitalizationRateChanged, error) {
	event := new(LoanPoolCapitalizationRateChanged)
	if err := _LoanPool.contract.UnpackLog(event, "CapitalizationRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the LoanPool contract.
type LoanPoolDepositedIterator struct {
	Event *LoanPoolDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolDeposited represents a Deposited event raised by the LoanPool contract.
type LoanPoolDeposited struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*LoanPoolDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolDepositedIterator{contract: _LoanPool.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *LoanPoolDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolDeposited)
				if err := _LoanPool.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) ParseDeposited(log types.Log) (*LoanPoolDeposited, error) {
	event := new(LoanPoolDeposited)
	if err := _LoanPool.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolFeePaidIterator is returned from FilterFeePaid and is used to iterate over the raw logs and unpacked data for FeePaid events raised by the LoanPool contract.
type LoanPoolFeePaidIterator struct {
	Event *LoanPoolFeePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolFeePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolFeePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolFeePaid represents a FeePaid event raised by the LoanPool contract.
type LoanPoolFeePaid struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeePaid is a free log retrieval operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) FilterFeePaid(opts *bind.FilterOpts, user []common.Address) (*LoanPoolFeePaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "FeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolFeePaidIterator{contract: _LoanPool.contract, event: "FeePaid", logs: logs, sub: sub}, nil
}

// WatchFeePaid is a free log subscription operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) WatchFeePaid(opts *bind.WatchOpts, sink chan<- *LoanPoolFeePaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "FeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolFeePaid)
				if err := _LoanPool.contract.UnpackLog(event, "FeePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeePaid is a log parse operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) ParseFeePaid(log types.Log) (*LoanPoolFeePaid, error) {
	event := new(LoanPoolFeePaid)
	if err := _LoanPool.contract.UnpackLog(event, "FeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LoanPool contract.
type LoanPoolInitializedIterator struct {
	Event *LoanPoolInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolInitialized represents a Initialized event raised by the LoanPool contract.
type LoanPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LoanPool *LoanPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*LoanPoolInitializedIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LoanPoolInitializedIterator{contract: _LoanPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LoanPool *LoanPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LoanPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolInitialized)
				if err := _LoanPool.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LoanPool *LoanPoolFilterer) ParseInitialized(log types.Log) (*LoanPoolInitialized, error) {
	event := new(LoanPoolInitialized)
	if err := _LoanPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolLiquidatedIterator is returned from FilterLiquidated and is used to iterate over the raw logs and unpacked data for Liquidated events raised by the LoanPool contract.
type LoanPoolLiquidatedIterator struct {
	Event *LoanPoolLiquidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolLiquidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolLiquidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolLiquidated represents a Liquidated event raised by the LoanPool contract.
type LoanPoolLiquidated struct {
	User      common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLiquidated is a free log retrieval operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address recipient, uint256 amount)
func (_LoanPool *LoanPoolFilterer) FilterLiquidated(opts *bind.FilterOpts, user []common.Address) (*LoanPoolLiquidatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Liquidated", userRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolLiquidatedIterator{contract: _LoanPool.contract, event: "Liquidated", logs: logs, sub: sub}, nil
}

// WatchLiquidated is a free log subscription operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address recipient, uint256 amount)
func (_LoanPool *LoanPoolFilterer) WatchLiquidated(opts *bind.WatchOpts, sink chan<- *LoanPoolLiquidated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Liquidated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolLiquidated)
				if err := _LoanPool.contract.UnpackLog(event, "Liquidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidated is a log parse operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address recipient, uint256 amount)
func (_LoanPool *LoanPoolFilterer) ParseLiquidated(log types.Log) (*LoanPoolLiquidated, error) {
	event := new(LoanPoolLiquidated)
	if err := _LoanPool.contract.UnpackLog(event, "Liquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolLiquidatorSetIterator is returned from FilterLiquidatorSet and is used to iterate over the raw logs and unpacked data for LiquidatorSet events raised by the LoanPool contract.
type LoanPoolLiquidatorSetIterator struct {
	Event *LoanPoolLiquidatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolLiquidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolLiquidatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolLiquidatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolLiquidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolLiquidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolLiquidatorSet represents a LiquidatorSet event raised by the LoanPool contract.
type LoanPoolLiquidatorSet struct {
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidatorSet is a free log retrieval operation binding the contract event 0xe8c0056707ed008ae443ab5de430ae0701dacac38472cca13c9c1d02593ef88c.
//
// Solidity: event LiquidatorSet(address indexed liquidator)
func (_LoanPool *LoanPoolFilterer) FilterLiquidatorSet(opts *bind.FilterOpts, liquidator []common.Address) (*LoanPoolLiquidatorSetIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "LiquidatorSet", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolLiquidatorSetIterator{contract: _LoanPool.contract, event: "LiquidatorSet", logs: logs, sub: sub}, nil
}

// WatchLiquidatorSet is a free log subscription operation binding the contract event 0xe8c0056707ed008ae443ab5de430ae0701dacac38472cca13c9c1d02593ef88c.
//
// Solidity: event LiquidatorSet(address indexed liquidator)
func (_LoanPool *LoanPoolFilterer) WatchLiquidatorSet(opts *bind.WatchOpts, sink chan<- *LoanPoolLiquidatorSet, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "LiquidatorSet", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolLiquidatorSet)
				if err := _LoanPool.contract.UnpackLog(event, "LiquidatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidatorSet is a log parse operation binding the contract event 0xe8c0056707ed008ae443ab5de430ae0701dacac38472cca13c9c1d02593ef88c.
//
// Solidity: event LiquidatorSet(address indexed liquidator)
func (_LoanPool *LoanPoolFilterer) ParseLiquidatorSet(log types.Log) (*LoanPoolLiquidatorSet, error) {
	event := new(LoanPoolLiquidatorSet)
	if err := _LoanPool.contract.UnpackLog(event, "LiquidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolOracleSetIterator is returned from FilterOracleSet and is used to iterate over the raw logs and unpacked data for OracleSet events raised by the LoanPool contract.
type LoanPoolOracleSetIterator struct {
	Event *LoanPoolOracleSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolOracleSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolOracleSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolOracleSet represents a OracleSet event raised by the LoanPool contract.
type LoanPoolOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleSet is a free log retrieval operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_LoanPool *LoanPoolFilterer) FilterOracleSet(opts *bind.FilterOpts) (*LoanPoolOracleSetIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "OracleSet")
	if err != nil {
		return nil, err
	}
	return &LoanPoolOracleSetIterator{contract: _LoanPool.contract, event: "OracleSet", logs: logs, sub: sub}, nil
}

// WatchOracleSet is a free log subscription operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_LoanPool *LoanPoolFilterer) WatchOracleSet(opts *bind.WatchOpts, sink chan<- *LoanPoolOracleSet) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "OracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolOracleSet)
				if err := _LoanPool.contract.UnpackLog(event, "OracleSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleSet is a log parse operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_LoanPool *LoanPoolFilterer) ParseOracleSet(log types.Log) (*LoanPoolOracleSet, error) {
	event := new(LoanPoolOracleSet)
	if err := _LoanPool.contract.UnpackLog(event, "OracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LoanPool contract.
type LoanPoolOwnershipTransferredIterator struct {
	Event *LoanPoolOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolOwnershipTransferred represents a OwnershipTransferred event raised by the LoanPool contract.
type LoanPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanPool *LoanPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LoanPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolOwnershipTransferredIterator{contract: _LoanPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanPool *LoanPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LoanPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolOwnershipTransferred)
				if err := _LoanPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanPool *LoanPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LoanPoolOwnershipTransferred, error) {
	event := new(LoanPoolOwnershipTransferred)
	if err := _LoanPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LoanPool contract.
type LoanPoolPausedIterator struct {
	Event *LoanPoolPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolPaused represents a Paused event raised by the LoanPool contract.
type LoanPoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LoanPool *LoanPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LoanPoolPausedIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LoanPoolPausedIterator{contract: _LoanPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LoanPool *LoanPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LoanPoolPaused) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolPaused)
				if err := _LoanPool.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LoanPool *LoanPoolFilterer) ParsePaused(log types.Log) (*LoanPoolPaused, error) {
	event := new(LoanPoolPaused)
	if err := _LoanPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolPoolDataSetIterator is returned from FilterPoolDataSet and is used to iterate over the raw logs and unpacked data for PoolDataSet events raised by the LoanPool contract.
type LoanPoolPoolDataSetIterator struct {
	Event *LoanPoolPoolDataSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolPoolDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolPoolDataSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolPoolDataSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolPoolDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolPoolDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolPoolDataSet represents a PoolDataSet event raised by the LoanPool contract.
type LoanPoolPoolDataSet struct {
	PoolData ILoanPoolPoolData
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolDataSet is a free log retrieval operation binding the contract event 0x8916c297dad8c35c95ebb9bdc996c37d30aa3e5cb778e7b54d9c06160a318727.
//
// Solidity: event PoolDataSet((uint256) poolData)
func (_LoanPool *LoanPoolFilterer) FilterPoolDataSet(opts *bind.FilterOpts) (*LoanPoolPoolDataSetIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "PoolDataSet")
	if err != nil {
		return nil, err
	}
	return &LoanPoolPoolDataSetIterator{contract: _LoanPool.contract, event: "PoolDataSet", logs: logs, sub: sub}, nil
}

// WatchPoolDataSet is a free log subscription operation binding the contract event 0x8916c297dad8c35c95ebb9bdc996c37d30aa3e5cb778e7b54d9c06160a318727.
//
// Solidity: event PoolDataSet((uint256) poolData)
func (_LoanPool *LoanPoolFilterer) WatchPoolDataSet(opts *bind.WatchOpts, sink chan<- *LoanPoolPoolDataSet) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "PoolDataSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolPoolDataSet)
				if err := _LoanPool.contract.UnpackLog(event, "PoolDataSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolDataSet is a log parse operation binding the contract event 0x8916c297dad8c35c95ebb9bdc996c37d30aa3e5cb778e7b54d9c06160a318727.
//
// Solidity: event PoolDataSet((uint256) poolData)
func (_LoanPool *LoanPoolFilterer) ParsePoolDataSet(log types.Log) (*LoanPoolPoolDataSet, error) {
	event := new(LoanPoolPoolDataSet)
	if err := _LoanPool.contract.UnpackLog(event, "PoolDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolRepaidIterator is returned from FilterRepaid and is used to iterate over the raw logs and unpacked data for Repaid events raised by the LoanPool contract.
type LoanPoolRepaidIterator struct {
	Event *LoanPoolRepaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolRepaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolRepaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolRepaid represents a Repaid event raised by the LoanPool contract.
type LoanPoolRepaid struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepaid is a free log retrieval operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) FilterRepaid(opts *bind.FilterOpts, user []common.Address) (*LoanPoolRepaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Repaid", userRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolRepaidIterator{contract: _LoanPool.contract, event: "Repaid", logs: logs, sub: sub}, nil
}

// WatchRepaid is a free log subscription operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) WatchRepaid(opts *bind.WatchOpts, sink chan<- *LoanPoolRepaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Repaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolRepaid)
				if err := _LoanPool.contract.UnpackLog(event, "Repaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepaid is a log parse operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) ParseRepaid(log types.Log) (*LoanPoolRepaid, error) {
	event := new(LoanPoolRepaid)
	if err := _LoanPool.contract.UnpackLog(event, "Repaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolTreasurySetIterator is returned from FilterTreasurySet and is used to iterate over the raw logs and unpacked data for TreasurySet events raised by the LoanPool contract.
type LoanPoolTreasurySetIterator struct {
	Event *LoanPoolTreasurySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolTreasurySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolTreasurySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolTreasurySet represents a TreasurySet event raised by the LoanPool contract.
type LoanPoolTreasurySet struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTreasurySet is a free log retrieval operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address treasury)
func (_LoanPool *LoanPoolFilterer) FilterTreasurySet(opts *bind.FilterOpts) (*LoanPoolTreasurySetIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "TreasurySet")
	if err != nil {
		return nil, err
	}
	return &LoanPoolTreasurySetIterator{contract: _LoanPool.contract, event: "TreasurySet", logs: logs, sub: sub}, nil
}

// WatchTreasurySet is a free log subscription operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address treasury)
func (_LoanPool *LoanPoolFilterer) WatchTreasurySet(opts *bind.WatchOpts, sink chan<- *LoanPoolTreasurySet) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "TreasurySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolTreasurySet)
				if err := _LoanPool.contract.UnpackLog(event, "TreasurySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasurySet is a log parse operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address treasury)
func (_LoanPool *LoanPoolFilterer) ParseTreasurySet(log types.Log) (*LoanPoolTreasurySet, error) {
	event := new(LoanPoolTreasurySet)
	if err := _LoanPool.contract.UnpackLog(event, "TreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LoanPool contract.
type LoanPoolUnpausedIterator struct {
	Event *LoanPoolUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolUnpaused represents a Unpaused event raised by the LoanPool contract.
type LoanPoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LoanPool *LoanPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LoanPoolUnpausedIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LoanPoolUnpausedIterator{contract: _LoanPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LoanPool *LoanPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LoanPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolUnpaused)
				if err := _LoanPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LoanPool *LoanPoolFilterer) ParseUnpaused(log types.Log) (*LoanPoolUnpaused, error) {
	event := new(LoanPoolUnpaused)
	if err := _LoanPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LoanPool contract.
type LoanPoolUpgradedIterator struct {
	Event *LoanPoolUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolUpgraded represents a Upgraded event raised by the LoanPool contract.
type LoanPoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LoanPool *LoanPoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LoanPoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolUpgradedIterator{contract: _LoanPool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LoanPool *LoanPoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LoanPoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolUpgraded)
				if err := _LoanPool.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LoanPool *LoanPoolFilterer) ParseUpgraded(log types.Log) (*LoanPoolUpgraded, error) {
	event := new(LoanPoolUpgraded)
	if err := _LoanPool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoanPoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the LoanPool contract.
type LoanPoolWithdrawnIterator struct {
	Event *LoanPoolWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoanPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoanPoolWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoanPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolWithdrawn represents a Withdrawn event raised by the LoanPool contract.
type LoanPoolWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*LoanPoolWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &LoanPoolWithdrawnIterator{contract: _LoanPool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *LoanPoolWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolWithdrawn)
				if err := _LoanPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_LoanPool *LoanPoolFilterer) ParseWithdrawn(log types.Log) (*LoanPoolWithdrawn, error) {
	event := new(LoanPoolWithdrawn)
	if err := _LoanPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

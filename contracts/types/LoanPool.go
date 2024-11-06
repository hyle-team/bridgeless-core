// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iloanpool

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

// ILoanPoolPoolData is an auto generated low-level Go binding around an user-defined struct.
type ILoanPoolPoolData struct {
	CollateralRatio *big.Int
}

// ILoanPoolMetaData contains all meta data concerning the ILoanPool contract.
var ILoanPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recepient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recepient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Liquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLiquidator\",\"type\":\"bool\"}],\"name\":\"LiquidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"OracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILoanPool.PoolData\",\"name\":\"poolData\",\"type\":\"tuple\"}],\"name\":\"PoolDataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recepient_\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getPossibleBorrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getPossibleWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"getUserDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"isLiquidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"isUserCanBeLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recepient_\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"capitalizationPeriod_\",\"type\":\"uint64\"}],\"name\":\"setCapitalizationPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capitalizationRate_\",\"type\":\"uint256\"}],\"name\":\"setCapitalizationRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"liquidators_\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isLiquidators_\",\"type\":\"bool[]\"}],\"name\":\"setLiquidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"}],\"internalType\":\"structILoanPool.PoolData\",\"name\":\"poolData_\",\"type\":\"tuple\"}],\"name\":\"setPoolData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"updateUserPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ILoanPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ILoanPoolMetaData.ABI instead.
var ILoanPoolABI = ILoanPoolMetaData.ABI

// ILoanPool is an auto generated Go binding around an Ethereum contract.
type ILoanPool struct {
	ILoanPoolCaller     // Read-only binding to the contract
	ILoanPoolTransactor // Write-only binding to the contract
	ILoanPoolFilterer   // Log filterer for contract events
}

// ILoanPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILoanPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILoanPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILoanPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILoanPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILoanPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILoanPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILoanPoolSession struct {
	Contract     *ILoanPool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILoanPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILoanPoolCallerSession struct {
	Contract *ILoanPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ILoanPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILoanPoolTransactorSession struct {
	Contract     *ILoanPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ILoanPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILoanPoolRaw struct {
	Contract *ILoanPool // Generic contract binding to access the raw methods on
}

// ILoanPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILoanPoolCallerRaw struct {
	Contract *ILoanPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ILoanPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILoanPoolTransactorRaw struct {
	Contract *ILoanPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILoanPool creates a new instance of ILoanPool, bound to a specific deployed contract.
func NewILoanPool(address common.Address, backend bind.ContractBackend) (*ILoanPool, error) {
	contract, err := bindILoanPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILoanPool{ILoanPoolCaller: ILoanPoolCaller{contract: contract}, ILoanPoolTransactor: ILoanPoolTransactor{contract: contract}, ILoanPoolFilterer: ILoanPoolFilterer{contract: contract}}, nil
}

// NewILoanPoolCaller creates a new read-only instance of ILoanPool, bound to a specific deployed contract.
func NewILoanPoolCaller(address common.Address, caller bind.ContractCaller) (*ILoanPoolCaller, error) {
	contract, err := bindILoanPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolCaller{contract: contract}, nil
}

// NewILoanPoolTransactor creates a new write-only instance of ILoanPool, bound to a specific deployed contract.
func NewILoanPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ILoanPoolTransactor, error) {
	contract, err := bindILoanPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolTransactor{contract: contract}, nil
}

// NewILoanPoolFilterer creates a new log filterer instance of ILoanPool, bound to a specific deployed contract.
func NewILoanPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ILoanPoolFilterer, error) {
	contract, err := bindILoanPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolFilterer{contract: contract}, nil
}

// bindILoanPool binds a generic wrapper to an already deployed contract.
func bindILoanPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILoanPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILoanPool *ILoanPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILoanPool.Contract.ILoanPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILoanPool *ILoanPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILoanPool.Contract.ILoanPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILoanPool *ILoanPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILoanPool.Contract.ILoanPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILoanPool *ILoanPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILoanPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILoanPool *ILoanPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILoanPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILoanPool *ILoanPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILoanPool.Contract.contract.Transact(opts, method, params...)
}

// BorrowToken is a free data retrieval call binding the contract method 0x456dc17a.
//
// Solidity: function borrowToken() view returns(address)
func (_ILoanPool *ILoanPoolCaller) BorrowToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "borrowToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowToken is a free data retrieval call binding the contract method 0x456dc17a.
//
// Solidity: function borrowToken() view returns(address)
func (_ILoanPool *ILoanPoolSession) BorrowToken() (common.Address, error) {
	return _ILoanPool.Contract.BorrowToken(&_ILoanPool.CallOpts)
}

// BorrowToken is a free data retrieval call binding the contract method 0x456dc17a.
//
// Solidity: function borrowToken() view returns(address)
func (_ILoanPool *ILoanPoolCallerSession) BorrowToken() (common.Address, error) {
	return _ILoanPool.Contract.BorrowToken(&_ILoanPool.CallOpts)
}

// GetPossibleBorrowAmount is a free data retrieval call binding the contract method 0xa80e9252.
//
// Solidity: function getPossibleBorrowAmount(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolCaller) GetPossibleBorrowAmount(opts *bind.CallOpts, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "getPossibleBorrowAmount", user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPossibleBorrowAmount is a free data retrieval call binding the contract method 0xa80e9252.
//
// Solidity: function getPossibleBorrowAmount(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolSession) GetPossibleBorrowAmount(user_ common.Address) (*big.Int, error) {
	return _ILoanPool.Contract.GetPossibleBorrowAmount(&_ILoanPool.CallOpts, user_)
}

// GetPossibleBorrowAmount is a free data retrieval call binding the contract method 0xa80e9252.
//
// Solidity: function getPossibleBorrowAmount(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolCallerSession) GetPossibleBorrowAmount(user_ common.Address) (*big.Int, error) {
	return _ILoanPool.Contract.GetPossibleBorrowAmount(&_ILoanPool.CallOpts, user_)
}

// GetPossibleWithdrawAmount is a free data retrieval call binding the contract method 0xc9c67933.
//
// Solidity: function getPossibleWithdrawAmount(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolCaller) GetPossibleWithdrawAmount(opts *bind.CallOpts, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "getPossibleWithdrawAmount", user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPossibleWithdrawAmount is a free data retrieval call binding the contract method 0xc9c67933.
//
// Solidity: function getPossibleWithdrawAmount(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolSession) GetPossibleWithdrawAmount(user_ common.Address) (*big.Int, error) {
	return _ILoanPool.Contract.GetPossibleWithdrawAmount(&_ILoanPool.CallOpts, user_)
}

// GetPossibleWithdrawAmount is a free data retrieval call binding the contract method 0xc9c67933.
//
// Solidity: function getPossibleWithdrawAmount(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolCallerSession) GetPossibleWithdrawAmount(user_ common.Address) (*big.Int, error) {
	return _ILoanPool.Contract.GetPossibleWithdrawAmount(&_ILoanPool.CallOpts, user_)
}

// GetUserDebt is a free data retrieval call binding the contract method 0xe65896c9.
//
// Solidity: function getUserDebt(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolCaller) GetUserDebt(opts *bind.CallOpts, user_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "getUserDebt", user_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDebt is a free data retrieval call binding the contract method 0xe65896c9.
//
// Solidity: function getUserDebt(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolSession) GetUserDebt(user_ common.Address) (*big.Int, error) {
	return _ILoanPool.Contract.GetUserDebt(&_ILoanPool.CallOpts, user_)
}

// GetUserDebt is a free data retrieval call binding the contract method 0xe65896c9.
//
// Solidity: function getUserDebt(address user_) view returns(uint256)
func (_ILoanPool *ILoanPoolCallerSession) GetUserDebt(user_ common.Address) (*big.Int, error) {
	return _ILoanPool.Contract.GetUserDebt(&_ILoanPool.CallOpts, user_)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address user_) view returns(bool)
func (_ILoanPool *ILoanPoolCaller) IsLiquidator(opts *bind.CallOpts, user_ common.Address) (bool, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "isLiquidator", user_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address user_) view returns(bool)
func (_ILoanPool *ILoanPoolSession) IsLiquidator(user_ common.Address) (bool, error) {
	return _ILoanPool.Contract.IsLiquidator(&_ILoanPool.CallOpts, user_)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address user_) view returns(bool)
func (_ILoanPool *ILoanPoolCallerSession) IsLiquidator(user_ common.Address) (bool, error) {
	return _ILoanPool.Contract.IsLiquidator(&_ILoanPool.CallOpts, user_)
}

// IsUserCanBeLiquidated is a free data retrieval call binding the contract method 0x7acc7dbb.
//
// Solidity: function isUserCanBeLiquidated(address user_) view returns(bool)
func (_ILoanPool *ILoanPoolCaller) IsUserCanBeLiquidated(opts *bind.CallOpts, user_ common.Address) (bool, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "isUserCanBeLiquidated", user_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserCanBeLiquidated is a free data retrieval call binding the contract method 0x7acc7dbb.
//
// Solidity: function isUserCanBeLiquidated(address user_) view returns(bool)
func (_ILoanPool *ILoanPoolSession) IsUserCanBeLiquidated(user_ common.Address) (bool, error) {
	return _ILoanPool.Contract.IsUserCanBeLiquidated(&_ILoanPool.CallOpts, user_)
}

// IsUserCanBeLiquidated is a free data retrieval call binding the contract method 0x7acc7dbb.
//
// Solidity: function isUserCanBeLiquidated(address user_) view returns(bool)
func (_ILoanPool *ILoanPoolCallerSession) IsUserCanBeLiquidated(user_ common.Address) (bool, error) {
	return _ILoanPool.Contract.IsUserCanBeLiquidated(&_ILoanPool.CallOpts, user_)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ILoanPool *ILoanPoolCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ILoanPool *ILoanPoolSession) Oracle() (common.Address, error) {
	return _ILoanPool.Contract.Oracle(&_ILoanPool.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ILoanPool *ILoanPoolCallerSession) Oracle() (common.Address, error) {
	return _ILoanPool.Contract.Oracle(&_ILoanPool.CallOpts)
}

// SupplyToken is a free data retrieval call binding the contract method 0x52059756.
//
// Solidity: function supplyToken() view returns(address)
func (_ILoanPool *ILoanPoolCaller) SupplyToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "supplyToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupplyToken is a free data retrieval call binding the contract method 0x52059756.
//
// Solidity: function supplyToken() view returns(address)
func (_ILoanPool *ILoanPoolSession) SupplyToken() (common.Address, error) {
	return _ILoanPool.Contract.SupplyToken(&_ILoanPool.CallOpts)
}

// SupplyToken is a free data retrieval call binding the contract method 0x52059756.
//
// Solidity: function supplyToken() view returns(address)
func (_ILoanPool *ILoanPoolCallerSession) SupplyToken() (common.Address, error) {
	return _ILoanPool.Contract.SupplyToken(&_ILoanPool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ILoanPool *ILoanPoolCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILoanPool.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ILoanPool *ILoanPoolSession) Treasury() (common.Address, error) {
	return _ILoanPool.Contract.Treasury(&_ILoanPool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ILoanPool *ILoanPoolCallerSession) Treasury() (common.Address, error) {
	return _ILoanPool.Contract.Treasury(&_ILoanPool.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount_, address recepient_) payable returns()
func (_ILoanPool *ILoanPoolTransactor) Borrow(opts *bind.TransactOpts, amount_ *big.Int, recepient_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "borrow", amount_, recepient_)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount_, address recepient_) payable returns()
func (_ILoanPool *ILoanPoolSession) Borrow(amount_ *big.Int, recepient_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.Borrow(&_ILoanPool.TransactOpts, amount_, recepient_)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount_, address recepient_) payable returns()
func (_ILoanPool *ILoanPoolTransactorSession) Borrow(amount_ *big.Int, recepient_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.Borrow(&_ILoanPool.TransactOpts, amount_, recepient_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolTransactor) Deposit(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "deposit", amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolSession) Deposit(amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.Deposit(&_ILoanPool.TransactOpts, amount_)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolTransactorSession) Deposit(amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.Deposit(&_ILoanPool.TransactOpts, amount_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address user_, address recepient_) returns()
func (_ILoanPool *ILoanPoolTransactor) Liquidate(opts *bind.TransactOpts, user_ common.Address, recepient_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "liquidate", user_, recepient_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address user_, address recepient_) returns()
func (_ILoanPool *ILoanPoolSession) Liquidate(user_ common.Address, recepient_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.Liquidate(&_ILoanPool.TransactOpts, user_, recepient_)
}

// Liquidate is a paid mutator transaction binding the contract method 0x86b9d81f.
//
// Solidity: function liquidate(address user_, address recepient_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) Liquidate(user_ common.Address, recepient_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.Liquidate(&_ILoanPool.TransactOpts, user_, recepient_)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolTransactor) Repay(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "repay", amount_)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolSession) Repay(amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.Repay(&_ILoanPool.TransactOpts, amount_)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolTransactorSession) Repay(amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.Repay(&_ILoanPool.TransactOpts, amount_)
}

// SetCapitalizationPeriod is a paid mutator transaction binding the contract method 0xf70f48a9.
//
// Solidity: function setCapitalizationPeriod(uint64 capitalizationPeriod_) returns()
func (_ILoanPool *ILoanPoolTransactor) SetCapitalizationPeriod(opts *bind.TransactOpts, capitalizationPeriod_ uint64) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "setCapitalizationPeriod", capitalizationPeriod_)
}

// SetCapitalizationPeriod is a paid mutator transaction binding the contract method 0xf70f48a9.
//
// Solidity: function setCapitalizationPeriod(uint64 capitalizationPeriod_) returns()
func (_ILoanPool *ILoanPoolSession) SetCapitalizationPeriod(capitalizationPeriod_ uint64) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetCapitalizationPeriod(&_ILoanPool.TransactOpts, capitalizationPeriod_)
}

// SetCapitalizationPeriod is a paid mutator transaction binding the contract method 0xf70f48a9.
//
// Solidity: function setCapitalizationPeriod(uint64 capitalizationPeriod_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) SetCapitalizationPeriod(capitalizationPeriod_ uint64) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetCapitalizationPeriod(&_ILoanPool.TransactOpts, capitalizationPeriod_)
}

// SetCapitalizationRate is a paid mutator transaction binding the contract method 0x9aad5c9a.
//
// Solidity: function setCapitalizationRate(uint256 capitalizationRate_) returns()
func (_ILoanPool *ILoanPoolTransactor) SetCapitalizationRate(opts *bind.TransactOpts, capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "setCapitalizationRate", capitalizationRate_)
}

// SetCapitalizationRate is a paid mutator transaction binding the contract method 0x9aad5c9a.
//
// Solidity: function setCapitalizationRate(uint256 capitalizationRate_) returns()
func (_ILoanPool *ILoanPoolSession) SetCapitalizationRate(capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetCapitalizationRate(&_ILoanPool.TransactOpts, capitalizationRate_)
}

// SetCapitalizationRate is a paid mutator transaction binding the contract method 0x9aad5c9a.
//
// Solidity: function setCapitalizationRate(uint256 capitalizationRate_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) SetCapitalizationRate(capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetCapitalizationRate(&_ILoanPool.TransactOpts, capitalizationRate_)
}

// SetLiquidators is a paid mutator transaction binding the contract method 0x247260ba.
//
// Solidity: function setLiquidators(address[] liquidators_, bool[] isLiquidators_) returns()
func (_ILoanPool *ILoanPoolTransactor) SetLiquidators(opts *bind.TransactOpts, liquidators_ []common.Address, isLiquidators_ []bool) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "setLiquidators", liquidators_, isLiquidators_)
}

// SetLiquidators is a paid mutator transaction binding the contract method 0x247260ba.
//
// Solidity: function setLiquidators(address[] liquidators_, bool[] isLiquidators_) returns()
func (_ILoanPool *ILoanPoolSession) SetLiquidators(liquidators_ []common.Address, isLiquidators_ []bool) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetLiquidators(&_ILoanPool.TransactOpts, liquidators_, isLiquidators_)
}

// SetLiquidators is a paid mutator transaction binding the contract method 0x247260ba.
//
// Solidity: function setLiquidators(address[] liquidators_, bool[] isLiquidators_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) SetLiquidators(liquidators_ []common.Address, isLiquidators_ []bool) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetLiquidators(&_ILoanPool.TransactOpts, liquidators_, isLiquidators_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_ILoanPool *ILoanPoolTransactor) SetOracle(opts *bind.TransactOpts, oracle_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "setOracle", oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_ILoanPool *ILoanPoolSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetOracle(&_ILoanPool.TransactOpts, oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetOracle(&_ILoanPool.TransactOpts, oracle_)
}

// SetPoolData is a paid mutator transaction binding the contract method 0xd46b6406.
//
// Solidity: function setPoolData((uint256) poolData_) returns()
func (_ILoanPool *ILoanPoolTransactor) SetPoolData(opts *bind.TransactOpts, poolData_ ILoanPoolPoolData) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "setPoolData", poolData_)
}

// SetPoolData is a paid mutator transaction binding the contract method 0xd46b6406.
//
// Solidity: function setPoolData((uint256) poolData_) returns()
func (_ILoanPool *ILoanPoolSession) SetPoolData(poolData_ ILoanPoolPoolData) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetPoolData(&_ILoanPool.TransactOpts, poolData_)
}

// SetPoolData is a paid mutator transaction binding the contract method 0xd46b6406.
//
// Solidity: function setPoolData((uint256) poolData_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) SetPoolData(poolData_ ILoanPoolPoolData) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetPoolData(&_ILoanPool.TransactOpts, poolData_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ILoanPool *ILoanPoolTransactor) SetTreasury(opts *bind.TransactOpts, treasury_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "setTreasury", treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ILoanPool *ILoanPoolSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetTreasury(&_ILoanPool.TransactOpts, treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.SetTreasury(&_ILoanPool.TransactOpts, treasury_)
}

// UpdateUserPosition is a paid mutator transaction binding the contract method 0xea5cbe07.
//
// Solidity: function updateUserPosition(address user_) returns()
func (_ILoanPool *ILoanPoolTransactor) UpdateUserPosition(opts *bind.TransactOpts, user_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "updateUserPosition", user_)
}

// UpdateUserPosition is a paid mutator transaction binding the contract method 0xea5cbe07.
//
// Solidity: function updateUserPosition(address user_) returns()
func (_ILoanPool *ILoanPoolSession) UpdateUserPosition(user_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.UpdateUserPosition(&_ILoanPool.TransactOpts, user_)
}

// UpdateUserPosition is a paid mutator transaction binding the contract method 0xea5cbe07.
//
// Solidity: function updateUserPosition(address user_) returns()
func (_ILoanPool *ILoanPoolTransactorSession) UpdateUserPosition(user_ common.Address) (*types.Transaction, error) {
	return _ILoanPool.Contract.UpdateUserPosition(&_ILoanPool.TransactOpts, user_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolTransactor) Withdraw(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.contract.Transact(opts, "withdraw", amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.Withdraw(&_ILoanPool.TransactOpts, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount_) payable returns()
func (_ILoanPool *ILoanPoolTransactorSession) Withdraw(amount_ *big.Int) (*types.Transaction, error) {
	return _ILoanPool.Contract.Withdraw(&_ILoanPool.TransactOpts, amount_)
}

// ILoanPoolBorrowedIterator is returned from FilterBorrowed and is used to iterate over the raw logs and unpacked data for Borrowed events raised by the ILoanPool contract.
type ILoanPoolBorrowedIterator struct {
	Event *ILoanPoolBorrowed // Event containing the contract specifics and raw log

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
func (it *ILoanPoolBorrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolBorrowed)
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
		it.Event = new(ILoanPoolBorrowed)
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
func (it *ILoanPoolBorrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolBorrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolBorrowed represents a Borrowed event raised by the ILoanPool contract.
type ILoanPoolBorrowed struct {
	User      common.Address
	Recepient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBorrowed is a free log retrieval operation binding the contract event 0x3fc499aeb0bb1cb58b6de8b02b3f86f4e7394e9690bef0110e32ced8a5631045.
//
// Solidity: event Borrowed(address indexed user, address recepient, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) FilterBorrowed(opts *bind.FilterOpts, user []common.Address) (*ILoanPoolBorrowedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "Borrowed", userRule)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolBorrowedIterator{contract: _ILoanPool.contract, event: "Borrowed", logs: logs, sub: sub}, nil
}

// WatchBorrowed is a free log subscription operation binding the contract event 0x3fc499aeb0bb1cb58b6de8b02b3f86f4e7394e9690bef0110e32ced8a5631045.
//
// Solidity: event Borrowed(address indexed user, address recepient, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) WatchBorrowed(opts *bind.WatchOpts, sink chan<- *ILoanPoolBorrowed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "Borrowed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolBorrowed)
				if err := _ILoanPool.contract.UnpackLog(event, "Borrowed", log); err != nil {
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
// Solidity: event Borrowed(address indexed user, address recepient, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) ParseBorrowed(log types.Log) (*ILoanPoolBorrowed, error) {
	event := new(ILoanPoolBorrowed)
	if err := _ILoanPool.contract.UnpackLog(event, "Borrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ILoanPool contract.
type ILoanPoolDepositedIterator struct {
	Event *ILoanPoolDeposited // Event containing the contract specifics and raw log

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
func (it *ILoanPoolDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolDeposited)
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
		it.Event = new(ILoanPoolDeposited)
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
func (it *ILoanPoolDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolDeposited represents a Deposited event raised by the ILoanPool contract.
type ILoanPoolDeposited struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*ILoanPoolDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolDepositedIterator{contract: _ILoanPool.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ILoanPoolDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolDeposited)
				if err := _ILoanPool.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_ILoanPool *ILoanPoolFilterer) ParseDeposited(log types.Log) (*ILoanPoolDeposited, error) {
	event := new(ILoanPoolDeposited)
	if err := _ILoanPool.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolFeePaidIterator is returned from FilterFeePaid and is used to iterate over the raw logs and unpacked data for FeePaid events raised by the ILoanPool contract.
type ILoanPoolFeePaidIterator struct {
	Event *ILoanPoolFeePaid // Event containing the contract specifics and raw log

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
func (it *ILoanPoolFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolFeePaid)
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
		it.Event = new(ILoanPoolFeePaid)
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
func (it *ILoanPoolFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolFeePaid represents a FeePaid event raised by the ILoanPool contract.
type ILoanPoolFeePaid struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeePaid is a free log retrieval operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) FilterFeePaid(opts *bind.FilterOpts, user []common.Address) (*ILoanPoolFeePaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "FeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolFeePaidIterator{contract: _ILoanPool.contract, event: "FeePaid", logs: logs, sub: sub}, nil
}

// WatchFeePaid is a free log subscription operation binding the contract event 0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f.
//
// Solidity: event FeePaid(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) WatchFeePaid(opts *bind.WatchOpts, sink chan<- *ILoanPoolFeePaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "FeePaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolFeePaid)
				if err := _ILoanPool.contract.UnpackLog(event, "FeePaid", log); err != nil {
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
func (_ILoanPool *ILoanPoolFilterer) ParseFeePaid(log types.Log) (*ILoanPoolFeePaid, error) {
	event := new(ILoanPoolFeePaid)
	if err := _ILoanPool.contract.UnpackLog(event, "FeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolLiquidatedIterator is returned from FilterLiquidated and is used to iterate over the raw logs and unpacked data for Liquidated events raised by the ILoanPool contract.
type ILoanPoolLiquidatedIterator struct {
	Event *ILoanPoolLiquidated // Event containing the contract specifics and raw log

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
func (it *ILoanPoolLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolLiquidated)
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
		it.Event = new(ILoanPoolLiquidated)
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
func (it *ILoanPoolLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolLiquidated represents a Liquidated event raised by the ILoanPool contract.
type ILoanPoolLiquidated struct {
	User      common.Address
	Recepient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLiquidated is a free log retrieval operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address recepient, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) FilterLiquidated(opts *bind.FilterOpts, user []common.Address) (*ILoanPoolLiquidatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "Liquidated", userRule)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolLiquidatedIterator{contract: _ILoanPool.contract, event: "Liquidated", logs: logs, sub: sub}, nil
}

// WatchLiquidated is a free log subscription operation binding the contract event 0xde0aa27286f5cb3a4ed853dc4823ead62d63e92cdf13de09d6aece56970721a4.
//
// Solidity: event Liquidated(address indexed user, address recepient, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) WatchLiquidated(opts *bind.WatchOpts, sink chan<- *ILoanPoolLiquidated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "Liquidated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolLiquidated)
				if err := _ILoanPool.contract.UnpackLog(event, "Liquidated", log); err != nil {
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
// Solidity: event Liquidated(address indexed user, address recepient, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) ParseLiquidated(log types.Log) (*ILoanPoolLiquidated, error) {
	event := new(ILoanPoolLiquidated)
	if err := _ILoanPool.contract.UnpackLog(event, "Liquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolLiquidatorSetIterator is returned from FilterLiquidatorSet and is used to iterate over the raw logs and unpacked data for LiquidatorSet events raised by the ILoanPool contract.
type ILoanPoolLiquidatorSetIterator struct {
	Event *ILoanPoolLiquidatorSet // Event containing the contract specifics and raw log

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
func (it *ILoanPoolLiquidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolLiquidatorSet)
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
		it.Event = new(ILoanPoolLiquidatorSet)
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
func (it *ILoanPoolLiquidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolLiquidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolLiquidatorSet represents a LiquidatorSet event raised by the ILoanPool contract.
type ILoanPoolLiquidatorSet struct {
	Liquidator   common.Address
	IsLiquidator bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidatorSet is a free log retrieval operation binding the contract event 0x81e020344174972c59f6c11a8f6c90b141866214e3d9b544d030f0b532f5a10f.
//
// Solidity: event LiquidatorSet(address indexed liquidator, bool isLiquidator)
func (_ILoanPool *ILoanPoolFilterer) FilterLiquidatorSet(opts *bind.FilterOpts, liquidator []common.Address) (*ILoanPoolLiquidatorSetIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "LiquidatorSet", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolLiquidatorSetIterator{contract: _ILoanPool.contract, event: "LiquidatorSet", logs: logs, sub: sub}, nil
}

// WatchLiquidatorSet is a free log subscription operation binding the contract event 0x81e020344174972c59f6c11a8f6c90b141866214e3d9b544d030f0b532f5a10f.
//
// Solidity: event LiquidatorSet(address indexed liquidator, bool isLiquidator)
func (_ILoanPool *ILoanPoolFilterer) WatchLiquidatorSet(opts *bind.WatchOpts, sink chan<- *ILoanPoolLiquidatorSet, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "LiquidatorSet", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolLiquidatorSet)
				if err := _ILoanPool.contract.UnpackLog(event, "LiquidatorSet", log); err != nil {
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

// ParseLiquidatorSet is a log parse operation binding the contract event 0x81e020344174972c59f6c11a8f6c90b141866214e3d9b544d030f0b532f5a10f.
//
// Solidity: event LiquidatorSet(address indexed liquidator, bool isLiquidator)
func (_ILoanPool *ILoanPoolFilterer) ParseLiquidatorSet(log types.Log) (*ILoanPoolLiquidatorSet, error) {
	event := new(ILoanPoolLiquidatorSet)
	if err := _ILoanPool.contract.UnpackLog(event, "LiquidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolOracleSetIterator is returned from FilterOracleSet and is used to iterate over the raw logs and unpacked data for OracleSet events raised by the ILoanPool contract.
type ILoanPoolOracleSetIterator struct {
	Event *ILoanPoolOracleSet // Event containing the contract specifics and raw log

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
func (it *ILoanPoolOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolOracleSet)
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
		it.Event = new(ILoanPoolOracleSet)
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
func (it *ILoanPoolOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolOracleSet represents a OracleSet event raised by the ILoanPool contract.
type ILoanPoolOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleSet is a free log retrieval operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_ILoanPool *ILoanPoolFilterer) FilterOracleSet(opts *bind.FilterOpts) (*ILoanPoolOracleSetIterator, error) {

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "OracleSet")
	if err != nil {
		return nil, err
	}
	return &ILoanPoolOracleSetIterator{contract: _ILoanPool.contract, event: "OracleSet", logs: logs, sub: sub}, nil
}

// WatchOracleSet is a free log subscription operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_ILoanPool *ILoanPoolFilterer) WatchOracleSet(opts *bind.WatchOpts, sink chan<- *ILoanPoolOracleSet) (event.Subscription, error) {

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "OracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolOracleSet)
				if err := _ILoanPool.contract.UnpackLog(event, "OracleSet", log); err != nil {
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
func (_ILoanPool *ILoanPoolFilterer) ParseOracleSet(log types.Log) (*ILoanPoolOracleSet, error) {
	event := new(ILoanPoolOracleSet)
	if err := _ILoanPool.contract.UnpackLog(event, "OracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolPoolDataSetIterator is returned from FilterPoolDataSet and is used to iterate over the raw logs and unpacked data for PoolDataSet events raised by the ILoanPool contract.
type ILoanPoolPoolDataSetIterator struct {
	Event *ILoanPoolPoolDataSet // Event containing the contract specifics and raw log

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
func (it *ILoanPoolPoolDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolPoolDataSet)
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
		it.Event = new(ILoanPoolPoolDataSet)
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
func (it *ILoanPoolPoolDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolPoolDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolPoolDataSet represents a PoolDataSet event raised by the ILoanPool contract.
type ILoanPoolPoolDataSet struct {
	PoolData ILoanPoolPoolData
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolDataSet is a free log retrieval operation binding the contract event 0x8916c297dad8c35c95ebb9bdc996c37d30aa3e5cb778e7b54d9c06160a318727.
//
// Solidity: event PoolDataSet((uint256) poolData)
func (_ILoanPool *ILoanPoolFilterer) FilterPoolDataSet(opts *bind.FilterOpts) (*ILoanPoolPoolDataSetIterator, error) {

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "PoolDataSet")
	if err != nil {
		return nil, err
	}
	return &ILoanPoolPoolDataSetIterator{contract: _ILoanPool.contract, event: "PoolDataSet", logs: logs, sub: sub}, nil
}

// WatchPoolDataSet is a free log subscription operation binding the contract event 0x8916c297dad8c35c95ebb9bdc996c37d30aa3e5cb778e7b54d9c06160a318727.
//
// Solidity: event PoolDataSet((uint256) poolData)
func (_ILoanPool *ILoanPoolFilterer) WatchPoolDataSet(opts *bind.WatchOpts, sink chan<- *ILoanPoolPoolDataSet) (event.Subscription, error) {

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "PoolDataSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolPoolDataSet)
				if err := _ILoanPool.contract.UnpackLog(event, "PoolDataSet", log); err != nil {
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
func (_ILoanPool *ILoanPoolFilterer) ParsePoolDataSet(log types.Log) (*ILoanPoolPoolDataSet, error) {
	event := new(ILoanPoolPoolDataSet)
	if err := _ILoanPool.contract.UnpackLog(event, "PoolDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolRepaidIterator is returned from FilterRepaid and is used to iterate over the raw logs and unpacked data for Repaid events raised by the ILoanPool contract.
type ILoanPoolRepaidIterator struct {
	Event *ILoanPoolRepaid // Event containing the contract specifics and raw log

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
func (it *ILoanPoolRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolRepaid)
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
		it.Event = new(ILoanPoolRepaid)
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
func (it *ILoanPoolRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolRepaid represents a Repaid event raised by the ILoanPool contract.
type ILoanPoolRepaid struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepaid is a free log retrieval operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) FilterRepaid(opts *bind.FilterOpts, user []common.Address) (*ILoanPoolRepaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "Repaid", userRule)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolRepaidIterator{contract: _ILoanPool.contract, event: "Repaid", logs: logs, sub: sub}, nil
}

// WatchRepaid is a free log subscription operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) WatchRepaid(opts *bind.WatchOpts, sink chan<- *ILoanPoolRepaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "Repaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolRepaid)
				if err := _ILoanPool.contract.UnpackLog(event, "Repaid", log); err != nil {
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
func (_ILoanPool *ILoanPoolFilterer) ParseRepaid(log types.Log) (*ILoanPoolRepaid, error) {
	event := new(ILoanPoolRepaid)
	if err := _ILoanPool.contract.UnpackLog(event, "Repaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolTreasurySetIterator is returned from FilterTreasurySet and is used to iterate over the raw logs and unpacked data for TreasurySet events raised by the ILoanPool contract.
type ILoanPoolTreasurySetIterator struct {
	Event *ILoanPoolTreasurySet // Event containing the contract specifics and raw log

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
func (it *ILoanPoolTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolTreasurySet)
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
		it.Event = new(ILoanPoolTreasurySet)
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
func (it *ILoanPoolTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolTreasurySet represents a TreasurySet event raised by the ILoanPool contract.
type ILoanPoolTreasurySet struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTreasurySet is a free log retrieval operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address treasury)
func (_ILoanPool *ILoanPoolFilterer) FilterTreasurySet(opts *bind.FilterOpts) (*ILoanPoolTreasurySetIterator, error) {

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "TreasurySet")
	if err != nil {
		return nil, err
	}
	return &ILoanPoolTreasurySetIterator{contract: _ILoanPool.contract, event: "TreasurySet", logs: logs, sub: sub}, nil
}

// WatchTreasurySet is a free log subscription operation binding the contract event 0x3c864541ef71378c6229510ed90f376565ee42d9c5e0904a984a9e863e6db44f.
//
// Solidity: event TreasurySet(address treasury)
func (_ILoanPool *ILoanPoolFilterer) WatchTreasurySet(opts *bind.WatchOpts, sink chan<- *ILoanPoolTreasurySet) (event.Subscription, error) {

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "TreasurySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolTreasurySet)
				if err := _ILoanPool.contract.UnpackLog(event, "TreasurySet", log); err != nil {
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
func (_ILoanPool *ILoanPoolFilterer) ParseTreasurySet(log types.Log) (*ILoanPoolTreasurySet, error) {
	event := new(ILoanPoolTreasurySet)
	if err := _ILoanPool.contract.UnpackLog(event, "TreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILoanPoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ILoanPool contract.
type ILoanPoolWithdrawnIterator struct {
	Event *ILoanPoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *ILoanPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILoanPoolWithdrawn)
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
		it.Event = new(ILoanPoolWithdrawn)
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
func (it *ILoanPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILoanPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILoanPoolWithdrawn represents a Withdrawn event raised by the ILoanPool contract.
type ILoanPoolWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*ILoanPoolWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &ILoanPoolWithdrawnIterator{contract: _ILoanPool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_ILoanPool *ILoanPoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ILoanPoolWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILoanPool.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILoanPoolWithdrawn)
				if err := _ILoanPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_ILoanPool *ILoanPoolFilterer) ParseWithdrawn(log types.Log) (*ILoanPoolWithdrawn, error) {
	event := new(ILoanPoolWithdrawn)
	if err := _ILoanPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

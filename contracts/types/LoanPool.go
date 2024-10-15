// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package types

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

// LoanPoolLoanPositionInfo is an auto generated low-level Go binding around an user-defined struct.
type LoanPoolLoanPositionInfo struct {
	LastAbsoluteLoanAmount *big.Int
	NormalizedLoanAmount   *big.Int
	LastUpdateTimestamp    *big.Int
}

// LoanPoolMetaData contains all meta data concerning the LoanPool contract.
var LoanPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"}],\"name\":\"PositionCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"loanTokensAmount_\",\"type\":\"uint256\"}],\"name\":\"createLoanPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"loanTokensAmountArr_\",\"type\":\"uint256[]\"}],\"name\":\"createLoanPositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionId_\",\"type\":\"uint256\"}],\"name\":\"getLoanPositionInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lastAbsoluteLoanAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"normalizedLoanAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLoanPool.LoanPositionInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"capitalizationPeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"capitalizationRate_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPositionID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateKeeper\",\"outputs\":[{\"internalType\":\"contractOwnableCompoundRateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"positionId_\",\"type\":\"uint256\"}],\"name\":\"updatePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"positionIds_\",\"type\":\"uint256[]\"}],\"name\":\"updatePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetLoanPositionInfo is a free data retrieval call binding the contract method 0xd29bdc48.
//
// Solidity: function getLoanPositionInfo(uint256 positionId_) view returns((uint256,uint256,uint256))
func (_LoanPool *LoanPoolCaller) GetLoanPositionInfo(opts *bind.CallOpts, positionId_ *big.Int) (LoanPoolLoanPositionInfo, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "getLoanPositionInfo", positionId_)

	if err != nil {
		return *new(LoanPoolLoanPositionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(LoanPoolLoanPositionInfo)).(*LoanPoolLoanPositionInfo)

	return out0, err

}

// GetLoanPositionInfo is a free data retrieval call binding the contract method 0xd29bdc48.
//
// Solidity: function getLoanPositionInfo(uint256 positionId_) view returns((uint256,uint256,uint256))
func (_LoanPool *LoanPoolSession) GetLoanPositionInfo(positionId_ *big.Int) (LoanPoolLoanPositionInfo, error) {
	return _LoanPool.Contract.GetLoanPositionInfo(&_LoanPool.CallOpts, positionId_)
}

// GetLoanPositionInfo is a free data retrieval call binding the contract method 0xd29bdc48.
//
// Solidity: function getLoanPositionInfo(uint256 positionId_) view returns((uint256,uint256,uint256))
func (_LoanPool *LoanPoolCallerSession) GetLoanPositionInfo(positionId_ *big.Int) (LoanPoolLoanPositionInfo, error) {
	return _LoanPool.Contract.GetLoanPositionInfo(&_LoanPool.CallOpts, positionId_)
}

// NextPositionID is a free data retrieval call binding the contract method 0x1c824905.
//
// Solidity: function nextPositionID() view returns(uint256)
func (_LoanPool *LoanPoolCaller) NextPositionID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "nextPositionID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPositionID is a free data retrieval call binding the contract method 0x1c824905.
//
// Solidity: function nextPositionID() view returns(uint256)
func (_LoanPool *LoanPoolSession) NextPositionID() (*big.Int, error) {
	return _LoanPool.Contract.NextPositionID(&_LoanPool.CallOpts)
}

// NextPositionID is a free data retrieval call binding the contract method 0x1c824905.
//
// Solidity: function nextPositionID() view returns(uint256)
func (_LoanPool *LoanPoolCallerSession) NextPositionID() (*big.Int, error) {
	return _LoanPool.Contract.NextPositionID(&_LoanPool.CallOpts)
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

// RateKeeper is a free data retrieval call binding the contract method 0xb7782f39.
//
// Solidity: function rateKeeper() view returns(address)
func (_LoanPool *LoanPoolCaller) RateKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoanPool.contract.Call(opts, &out, "rateKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateKeeper is a free data retrieval call binding the contract method 0xb7782f39.
//
// Solidity: function rateKeeper() view returns(address)
func (_LoanPool *LoanPoolSession) RateKeeper() (common.Address, error) {
	return _LoanPool.Contract.RateKeeper(&_LoanPool.CallOpts)
}

// RateKeeper is a free data retrieval call binding the contract method 0xb7782f39.
//
// Solidity: function rateKeeper() view returns(address)
func (_LoanPool *LoanPoolCallerSession) RateKeeper() (common.Address, error) {
	return _LoanPool.Contract.RateKeeper(&_LoanPool.CallOpts)
}

// CreateLoanPosition is a paid mutator transaction binding the contract method 0x8f32cd9e.
//
// Solidity: function createLoanPosition(uint256 loanTokensAmount_) returns()
func (_LoanPool *LoanPoolTransactor) CreateLoanPosition(opts *bind.TransactOpts, loanTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "createLoanPosition", loanTokensAmount_)
}

// CreateLoanPosition is a paid mutator transaction binding the contract method 0x8f32cd9e.
//
// Solidity: function createLoanPosition(uint256 loanTokensAmount_) returns()
func (_LoanPool *LoanPoolSession) CreateLoanPosition(loanTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.CreateLoanPosition(&_LoanPool.TransactOpts, loanTokensAmount_)
}

// CreateLoanPosition is a paid mutator transaction binding the contract method 0x8f32cd9e.
//
// Solidity: function createLoanPosition(uint256 loanTokensAmount_) returns()
func (_LoanPool *LoanPoolTransactorSession) CreateLoanPosition(loanTokensAmount_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.CreateLoanPosition(&_LoanPool.TransactOpts, loanTokensAmount_)
}

// CreateLoanPositions is a paid mutator transaction binding the contract method 0x27427905.
//
// Solidity: function createLoanPositions(uint256[] loanTokensAmountArr_) returns()
func (_LoanPool *LoanPoolTransactor) CreateLoanPositions(opts *bind.TransactOpts, loanTokensAmountArr_ []*big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "createLoanPositions", loanTokensAmountArr_)
}

// CreateLoanPositions is a paid mutator transaction binding the contract method 0x27427905.
//
// Solidity: function createLoanPositions(uint256[] loanTokensAmountArr_) returns()
func (_LoanPool *LoanPoolSession) CreateLoanPositions(loanTokensAmountArr_ []*big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.CreateLoanPositions(&_LoanPool.TransactOpts, loanTokensAmountArr_)
}

// CreateLoanPositions is a paid mutator transaction binding the contract method 0x27427905.
//
// Solidity: function createLoanPositions(uint256[] loanTokensAmountArr_) returns()
func (_LoanPool *LoanPoolTransactorSession) CreateLoanPositions(loanTokensAmountArr_ []*big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.CreateLoanPositions(&_LoanPool.TransactOpts, loanTokensAmountArr_)
}

// Initialize is a paid mutator transaction binding the contract method 0x294c865a.
//
// Solidity: function initialize(uint64 capitalizationPeriod_, uint256 capitalizationRate_) returns()
func (_LoanPool *LoanPoolTransactor) Initialize(opts *bind.TransactOpts, capitalizationPeriod_ uint64, capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "initialize", capitalizationPeriod_, capitalizationRate_)
}

// Initialize is a paid mutator transaction binding the contract method 0x294c865a.
//
// Solidity: function initialize(uint64 capitalizationPeriod_, uint256 capitalizationRate_) returns()
func (_LoanPool *LoanPoolSession) Initialize(capitalizationPeriod_ uint64, capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Initialize(&_LoanPool.TransactOpts, capitalizationPeriod_, capitalizationRate_)
}

// Initialize is a paid mutator transaction binding the contract method 0x294c865a.
//
// Solidity: function initialize(uint64 capitalizationPeriod_, uint256 capitalizationRate_) returns()
func (_LoanPool *LoanPoolTransactorSession) Initialize(capitalizationPeriod_ uint64, capitalizationRate_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.Initialize(&_LoanPool.TransactOpts, capitalizationPeriod_, capitalizationRate_)
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

// UpdatePosition is a paid mutator transaction binding the contract method 0x09f1c80a.
//
// Solidity: function updatePosition(uint256 positionId_) returns()
func (_LoanPool *LoanPoolTransactor) UpdatePosition(opts *bind.TransactOpts, positionId_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "updatePosition", positionId_)
}

// UpdatePosition is a paid mutator transaction binding the contract method 0x09f1c80a.
//
// Solidity: function updatePosition(uint256 positionId_) returns()
func (_LoanPool *LoanPoolSession) UpdatePosition(positionId_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.UpdatePosition(&_LoanPool.TransactOpts, positionId_)
}

// UpdatePosition is a paid mutator transaction binding the contract method 0x09f1c80a.
//
// Solidity: function updatePosition(uint256 positionId_) returns()
func (_LoanPool *LoanPoolTransactorSession) UpdatePosition(positionId_ *big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.UpdatePosition(&_LoanPool.TransactOpts, positionId_)
}

// UpdatePositions is a paid mutator transaction binding the contract method 0x203485da.
//
// Solidity: function updatePositions(uint256[] positionIds_) returns()
func (_LoanPool *LoanPoolTransactor) UpdatePositions(opts *bind.TransactOpts, positionIds_ []*big.Int) (*types.Transaction, error) {
	return _LoanPool.contract.Transact(opts, "updatePositions", positionIds_)
}

// UpdatePositions is a paid mutator transaction binding the contract method 0x203485da.
//
// Solidity: function updatePositions(uint256[] positionIds_) returns()
func (_LoanPool *LoanPoolSession) UpdatePositions(positionIds_ []*big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.UpdatePositions(&_LoanPool.TransactOpts, positionIds_)
}

// UpdatePositions is a paid mutator transaction binding the contract method 0x203485da.
//
// Solidity: function updatePositions(uint256[] positionIds_) returns()
func (_LoanPool *LoanPoolTransactorSession) UpdatePositions(positionIds_ []*big.Int) (*types.Transaction, error) {
	return _LoanPool.Contract.UpdatePositions(&_LoanPool.TransactOpts, positionIds_)
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

// LoanPoolPositionCreatedIterator is returned from FilterPositionCreated and is used to iterate over the raw logs and unpacked data for PositionCreated events raised by the LoanPool contract.
type LoanPoolPositionCreatedIterator struct {
	Event *LoanPoolPositionCreated // Event containing the contract specifics and raw log

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
func (it *LoanPoolPositionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanPoolPositionCreated)
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
		it.Event = new(LoanPoolPositionCreated)
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
func (it *LoanPoolPositionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanPoolPositionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanPoolPositionCreated represents a PositionCreated event raised by the LoanPool contract.
type LoanPoolPositionCreated struct {
	PositionId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPositionCreated is a free log retrieval operation binding the contract event 0x97c3f5c9077358c7266488de6a3ebba41df38417797d90b665239fcb506c840a.
//
// Solidity: event PositionCreated(uint256 positionId)
func (_LoanPool *LoanPoolFilterer) FilterPositionCreated(opts *bind.FilterOpts) (*LoanPoolPositionCreatedIterator, error) {

	logs, sub, err := _LoanPool.contract.FilterLogs(opts, "PositionCreated")
	if err != nil {
		return nil, err
	}
	return &LoanPoolPositionCreatedIterator{contract: _LoanPool.contract, event: "PositionCreated", logs: logs, sub: sub}, nil
}

// WatchPositionCreated is a free log subscription operation binding the contract event 0x97c3f5c9077358c7266488de6a3ebba41df38417797d90b665239fcb506c840a.
//
// Solidity: event PositionCreated(uint256 positionId)
func (_LoanPool *LoanPoolFilterer) WatchPositionCreated(opts *bind.WatchOpts, sink chan<- *LoanPoolPositionCreated) (event.Subscription, error) {

	logs, sub, err := _LoanPool.contract.WatchLogs(opts, "PositionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanPoolPositionCreated)
				if err := _LoanPool.contract.UnpackLog(event, "PositionCreated", log); err != nil {
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

// ParsePositionCreated is a log parse operation binding the contract event 0x97c3f5c9077358c7266488de6a3ebba41df38417797d90b665239fcb506c840a.
//
// Solidity: event PositionCreated(uint256 positionId)
func (_LoanPool *LoanPoolFilterer) ParsePositionCreated(log types.Log) (*LoanPoolPositionCreated, error) {
	event := new(LoanPoolPositionCreated)
	if err := _LoanPool.contract.UnpackLog(event, "PositionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

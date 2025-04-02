// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plaza_sdk

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

// PlazaSdkMetaData contains all meta data concerning the PlazaSdk contract.
var PlazaSdkMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocate\",\"inputs\":[{\"name\":\"_amountToDistribute\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"couponAmountToDistribute\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolFactory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPoolFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ClaimedShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessDenied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPoolAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughCouponBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughSharesBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughSharesToDistribute\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NothingToClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedPool\",\"inputs\":[]}]",
}

// PlazaSdkABI is the input ABI used to generate the binding from.
// Deprecated: Use PlazaSdkMetaData.ABI instead.
var PlazaSdkABI = PlazaSdkMetaData.ABI

// PlazaSdk is an auto generated Go binding around an Ethereum contract.
type PlazaSdk struct {
	PlazaSdkCaller     // Read-only binding to the contract
	PlazaSdkTransactor // Write-only binding to the contract
	PlazaSdkFilterer   // Log filterer for contract events
}

// PlazaSdkCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlazaSdkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlazaSdkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlazaSdkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlazaSdkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlazaSdkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlazaSdkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlazaSdkSession struct {
	Contract     *PlazaSdk         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlazaSdkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlazaSdkCallerSession struct {
	Contract *PlazaSdkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PlazaSdkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlazaSdkTransactorSession struct {
	Contract     *PlazaSdkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PlazaSdkRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlazaSdkRaw struct {
	Contract *PlazaSdk // Generic contract binding to access the raw methods on
}

// PlazaSdkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlazaSdkCallerRaw struct {
	Contract *PlazaSdkCaller // Generic read-only contract binding to access the raw methods on
}

// PlazaSdkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlazaSdkTransactorRaw struct {
	Contract *PlazaSdkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlazaSdk creates a new instance of PlazaSdk, bound to a specific deployed contract.
func NewPlazaSdk(address common.Address, backend bind.ContractBackend) (*PlazaSdk, error) {
	contract, err := bindPlazaSdk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlazaSdk{PlazaSdkCaller: PlazaSdkCaller{contract: contract}, PlazaSdkTransactor: PlazaSdkTransactor{contract: contract}, PlazaSdkFilterer: PlazaSdkFilterer{contract: contract}}, nil
}

// NewPlazaSdkCaller creates a new read-only instance of PlazaSdk, bound to a specific deployed contract.
func NewPlazaSdkCaller(address common.Address, caller bind.ContractCaller) (*PlazaSdkCaller, error) {
	contract, err := bindPlazaSdk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkCaller{contract: contract}, nil
}

// NewPlazaSdkTransactor creates a new write-only instance of PlazaSdk, bound to a specific deployed contract.
func NewPlazaSdkTransactor(address common.Address, transactor bind.ContractTransactor) (*PlazaSdkTransactor, error) {
	contract, err := bindPlazaSdk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkTransactor{contract: contract}, nil
}

// NewPlazaSdkFilterer creates a new log filterer instance of PlazaSdk, bound to a specific deployed contract.
func NewPlazaSdkFilterer(address common.Address, filterer bind.ContractFilterer) (*PlazaSdkFilterer, error) {
	contract, err := bindPlazaSdk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkFilterer{contract: contract}, nil
}

// bindPlazaSdk binds a generic wrapper to an already deployed contract.
func bindPlazaSdk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlazaSdkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlazaSdk *PlazaSdkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlazaSdk.Contract.PlazaSdkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlazaSdk *PlazaSdkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.Contract.PlazaSdkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlazaSdk *PlazaSdkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlazaSdk.Contract.PlazaSdkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlazaSdk *PlazaSdkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlazaSdk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlazaSdk *PlazaSdkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlazaSdk *PlazaSdkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlazaSdk.Contract.contract.Transact(opts, method, params...)
}

// CouponAmountToDistribute is a free data retrieval call binding the contract method 0xa543915f.
//
// Solidity: function couponAmountToDistribute() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) CouponAmountToDistribute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "couponAmountToDistribute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CouponAmountToDistribute is a free data retrieval call binding the contract method 0xa543915f.
//
// Solidity: function couponAmountToDistribute() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) CouponAmountToDistribute() (*big.Int, error) {
	return _PlazaSdk.Contract.CouponAmountToDistribute(&_PlazaSdk.CallOpts)
}

// CouponAmountToDistribute is a free data retrieval call binding the contract method 0xa543915f.
//
// Solidity: function couponAmountToDistribute() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) CouponAmountToDistribute() (*big.Int, error) {
	return _PlazaSdk.Contract.CouponAmountToDistribute(&_PlazaSdk.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlazaSdk *PlazaSdkCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlazaSdk *PlazaSdkSession) Paused() (bool, error) {
	return _PlazaSdk.Contract.Paused(&_PlazaSdk.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PlazaSdk *PlazaSdkCallerSession) Paused() (bool, error) {
	return _PlazaSdk.Contract.Paused(&_PlazaSdk.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PlazaSdk *PlazaSdkSession) Pool() (common.Address, error) {
	return _PlazaSdk.Contract.Pool(&_PlazaSdk.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) Pool() (common.Address, error) {
	return _PlazaSdk.Contract.Pool(&_PlazaSdk.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_PlazaSdk *PlazaSdkSession) PoolFactory() (common.Address, error) {
	return _PlazaSdk.Contract.PoolFactory(&_PlazaSdk.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) PoolFactory() (common.Address, error) {
	return _PlazaSdk.Contract.PoolFactory(&_PlazaSdk.CallOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0x90ca796b.
//
// Solidity: function allocate(uint256 _amountToDistribute) returns()
func (_PlazaSdk *PlazaSdkTransactor) Allocate(opts *bind.TransactOpts, _amountToDistribute *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "allocate", _amountToDistribute)
}

// Allocate is a paid mutator transaction binding the contract method 0x90ca796b.
//
// Solidity: function allocate(uint256 _amountToDistribute) returns()
func (_PlazaSdk *PlazaSdkSession) Allocate(_amountToDistribute *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Allocate(&_PlazaSdk.TransactOpts, _amountToDistribute)
}

// Allocate is a paid mutator transaction binding the contract method 0x90ca796b.
//
// Solidity: function allocate(uint256 _amountToDistribute) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Allocate(_amountToDistribute *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Allocate(&_PlazaSdk.TransactOpts, _amountToDistribute)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PlazaSdk *PlazaSdkTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PlazaSdk *PlazaSdkSession) Claim() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Claim(&_PlazaSdk.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Claim() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Claim(&_PlazaSdk.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pool, address _poolFactory) returns()
func (_PlazaSdk *PlazaSdkTransactor) Initialize(opts *bind.TransactOpts, _pool common.Address, _poolFactory common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "initialize", _pool, _poolFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pool, address _poolFactory) returns()
func (_PlazaSdk *PlazaSdkSession) Initialize(_pool common.Address, _poolFactory common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _pool, _poolFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pool, address _poolFactory) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Initialize(_pool common.Address, _poolFactory common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _pool, _poolFactory)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlazaSdk *PlazaSdkTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlazaSdk *PlazaSdkSession) Pause() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Pause(&_PlazaSdk.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Pause() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Pause(&_PlazaSdk.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlazaSdk *PlazaSdkTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlazaSdk *PlazaSdkSession) Unpause() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Unpause(&_PlazaSdk.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Unpause() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Unpause(&_PlazaSdk.TransactOpts)
}

// PlazaSdkClaimedSharesIterator is returned from FilterClaimedShares and is used to iterate over the raw logs and unpacked data for ClaimedShares events raised by the PlazaSdk contract.
type PlazaSdkClaimedSharesIterator struct {
	Event *PlazaSdkClaimedShares // Event containing the contract specifics and raw log

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
func (it *PlazaSdkClaimedSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkClaimedShares)
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
		it.Event = new(PlazaSdkClaimedShares)
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
func (it *PlazaSdkClaimedSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkClaimedSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkClaimedShares represents a ClaimedShares event raised by the PlazaSdk contract.
type PlazaSdkClaimedShares struct {
	User   common.Address
	Period *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimedShares is a free log retrieval operation binding the contract event 0x2cda5c7d9ff70b5c9119b1ba7469a9c367b07c395944a87a72953d9598f5bc62.
//
// Solidity: event ClaimedShares(address user, uint256 period, uint256 shares)
func (_PlazaSdk *PlazaSdkFilterer) FilterClaimedShares(opts *bind.FilterOpts) (*PlazaSdkClaimedSharesIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "ClaimedShares")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkClaimedSharesIterator{contract: _PlazaSdk.contract, event: "ClaimedShares", logs: logs, sub: sub}, nil
}

// WatchClaimedShares is a free log subscription operation binding the contract event 0x2cda5c7d9ff70b5c9119b1ba7469a9c367b07c395944a87a72953d9598f5bc62.
//
// Solidity: event ClaimedShares(address user, uint256 period, uint256 shares)
func (_PlazaSdk *PlazaSdkFilterer) WatchClaimedShares(opts *bind.WatchOpts, sink chan<- *PlazaSdkClaimedShares) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "ClaimedShares")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkClaimedShares)
				if err := _PlazaSdk.contract.UnpackLog(event, "ClaimedShares", log); err != nil {
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

// ParseClaimedShares is a log parse operation binding the contract event 0x2cda5c7d9ff70b5c9119b1ba7469a9c367b07c395944a87a72953d9598f5bc62.
//
// Solidity: event ClaimedShares(address user, uint256 period, uint256 shares)
func (_PlazaSdk *PlazaSdkFilterer) ParseClaimedShares(log types.Log) (*PlazaSdkClaimedShares, error) {
	event := new(PlazaSdkClaimedShares)
	if err := _PlazaSdk.contract.UnpackLog(event, "ClaimedShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PlazaSdk contract.
type PlazaSdkInitializedIterator struct {
	Event *PlazaSdkInitialized // Event containing the contract specifics and raw log

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
func (it *PlazaSdkInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkInitialized)
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
		it.Event = new(PlazaSdkInitialized)
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
func (it *PlazaSdkInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkInitialized represents a Initialized event raised by the PlazaSdk contract.
type PlazaSdkInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PlazaSdk *PlazaSdkFilterer) FilterInitialized(opts *bind.FilterOpts) (*PlazaSdkInitializedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkInitializedIterator{contract: _PlazaSdk.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PlazaSdk *PlazaSdkFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PlazaSdkInitialized) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkInitialized)
				if err := _PlazaSdk.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PlazaSdk *PlazaSdkFilterer) ParseInitialized(log types.Log) (*PlazaSdkInitialized, error) {
	event := new(PlazaSdkInitialized)
	if err := _PlazaSdk.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PlazaSdk contract.
type PlazaSdkPausedIterator struct {
	Event *PlazaSdkPaused // Event containing the contract specifics and raw log

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
func (it *PlazaSdkPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkPaused)
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
		it.Event = new(PlazaSdkPaused)
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
func (it *PlazaSdkPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkPaused represents a Paused event raised by the PlazaSdk contract.
type PlazaSdkPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PlazaSdk *PlazaSdkFilterer) FilterPaused(opts *bind.FilterOpts) (*PlazaSdkPausedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkPausedIterator{contract: _PlazaSdk.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PlazaSdk *PlazaSdkFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PlazaSdkPaused) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkPaused)
				if err := _PlazaSdk.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PlazaSdk *PlazaSdkFilterer) ParsePaused(log types.Log) (*PlazaSdkPaused, error) {
	event := new(PlazaSdkPaused)
	if err := _PlazaSdk.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PlazaSdk contract.
type PlazaSdkUnpausedIterator struct {
	Event *PlazaSdkUnpaused // Event containing the contract specifics and raw log

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
func (it *PlazaSdkUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkUnpaused)
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
		it.Event = new(PlazaSdkUnpaused)
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
func (it *PlazaSdkUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkUnpaused represents a Unpaused event raised by the PlazaSdk contract.
type PlazaSdkUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PlazaSdk *PlazaSdkFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PlazaSdkUnpausedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkUnpausedIterator{contract: _PlazaSdk.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PlazaSdk *PlazaSdkFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PlazaSdkUnpaused) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkUnpaused)
				if err := _PlazaSdk.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PlazaSdk *PlazaSdkFilterer) ParseUnpaused(log types.Log) (*PlazaSdkUnpaused, error) {
	event := new(PlazaSdkUnpaused)
	if err := _PlazaSdk.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

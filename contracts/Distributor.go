// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// DistributorMetaData contains all meta data concerning the Distributor contract.
var DistributorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocate\",\"inputs\":[{\"name\":\"_amountToDistribute\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"couponAmountToDistribute\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolFactory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPoolFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ClaimedShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessDenied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPoolAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughCouponBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughSharesBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughSharesToDistribute\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NothingToClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedPool\",\"inputs\":[]}]",
}

// DistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributorMetaData.ABI instead.
var DistributorABI = DistributorMetaData.ABI

// Distributor is an auto generated Go binding around an Ethereum contract.
type Distributor struct {
	DistributorCaller     // Read-only binding to the contract
	DistributorTransactor // Write-only binding to the contract
	DistributorFilterer   // Log filterer for contract events
}

// DistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributorSession struct {
	Contract     *Distributor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributorCallerSession struct {
	Contract *DistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributorTransactorSession struct {
	Contract     *DistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributorRaw struct {
	Contract *Distributor // Generic contract binding to access the raw methods on
}

// DistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributorCallerRaw struct {
	Contract *DistributorCaller // Generic read-only contract binding to access the raw methods on
}

// DistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributorTransactorRaw struct {
	Contract *DistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributor creates a new instance of Distributor, bound to a specific deployed contract.
func NewDistributor(address common.Address, backend bind.ContractBackend) (*Distributor, error) {
	contract, err := bindDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distributor{DistributorCaller: DistributorCaller{contract: contract}, DistributorTransactor: DistributorTransactor{contract: contract}, DistributorFilterer: DistributorFilterer{contract: contract}}, nil
}

// NewDistributorCaller creates a new read-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorCaller(address common.Address, caller bind.ContractCaller) (*DistributorCaller, error) {
	contract, err := bindDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorCaller{contract: contract}, nil
}

// NewDistributorTransactor creates a new write-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributorTransactor, error) {
	contract, err := bindDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorTransactor{contract: contract}, nil
}

// NewDistributorFilterer creates a new log filterer instance of Distributor, bound to a specific deployed contract.
func NewDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributorFilterer, error) {
	contract, err := bindDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributorFilterer{contract: contract}, nil
}

// bindDistributor binds a generic wrapper to an already deployed contract.
func bindDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.DistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transact(opts, method, params...)
}

// CouponAmountToDistribute is a free data retrieval call binding the contract method 0xa543915f.
//
// Solidity: function couponAmountToDistribute() view returns(uint256)
func (_Distributor *DistributorCaller) CouponAmountToDistribute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "couponAmountToDistribute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CouponAmountToDistribute is a free data retrieval call binding the contract method 0xa543915f.
//
// Solidity: function couponAmountToDistribute() view returns(uint256)
func (_Distributor *DistributorSession) CouponAmountToDistribute() (*big.Int, error) {
	return _Distributor.Contract.CouponAmountToDistribute(&_Distributor.CallOpts)
}

// CouponAmountToDistribute is a free data retrieval call binding the contract method 0xa543915f.
//
// Solidity: function couponAmountToDistribute() view returns(uint256)
func (_Distributor *DistributorCallerSession) CouponAmountToDistribute() (*big.Int, error) {
	return _Distributor.Contract.CouponAmountToDistribute(&_Distributor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Distributor *DistributorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Distributor *DistributorSession) Paused() (bool, error) {
	return _Distributor.Contract.Paused(&_Distributor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Distributor *DistributorCallerSession) Paused() (bool, error) {
	return _Distributor.Contract.Paused(&_Distributor.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Distributor *DistributorCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Distributor *DistributorSession) Pool() (common.Address, error) {
	return _Distributor.Contract.Pool(&_Distributor.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Distributor *DistributorCallerSession) Pool() (common.Address, error) {
	return _Distributor.Contract.Pool(&_Distributor.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Distributor *DistributorCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Distributor *DistributorSession) PoolFactory() (common.Address, error) {
	return _Distributor.Contract.PoolFactory(&_Distributor.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Distributor *DistributorCallerSession) PoolFactory() (common.Address, error) {
	return _Distributor.Contract.PoolFactory(&_Distributor.CallOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0x90ca796b.
//
// Solidity: function allocate(uint256 _amountToDistribute) returns()
func (_Distributor *DistributorTransactor) Allocate(opts *bind.TransactOpts, _amountToDistribute *big.Int) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "allocate", _amountToDistribute)
}

// Allocate is a paid mutator transaction binding the contract method 0x90ca796b.
//
// Solidity: function allocate(uint256 _amountToDistribute) returns()
func (_Distributor *DistributorSession) Allocate(_amountToDistribute *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.Allocate(&_Distributor.TransactOpts, _amountToDistribute)
}

// Allocate is a paid mutator transaction binding the contract method 0x90ca796b.
//
// Solidity: function allocate(uint256 _amountToDistribute) returns()
func (_Distributor *DistributorTransactorSession) Allocate(_amountToDistribute *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.Allocate(&_Distributor.TransactOpts, _amountToDistribute)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Distributor *DistributorTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Distributor *DistributorSession) Claim() (*types.Transaction, error) {
	return _Distributor.Contract.Claim(&_Distributor.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Distributor *DistributorTransactorSession) Claim() (*types.Transaction, error) {
	return _Distributor.Contract.Claim(&_Distributor.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pool, address _poolFactory) returns()
func (_Distributor *DistributorTransactor) Initialize(opts *bind.TransactOpts, _pool common.Address, _poolFactory common.Address) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "initialize", _pool, _poolFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pool, address _poolFactory) returns()
func (_Distributor *DistributorSession) Initialize(_pool common.Address, _poolFactory common.Address) (*types.Transaction, error) {
	return _Distributor.Contract.Initialize(&_Distributor.TransactOpts, _pool, _poolFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _pool, address _poolFactory) returns()
func (_Distributor *DistributorTransactorSession) Initialize(_pool common.Address, _poolFactory common.Address) (*types.Transaction, error) {
	return _Distributor.Contract.Initialize(&_Distributor.TransactOpts, _pool, _poolFactory)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Distributor *DistributorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Distributor *DistributorSession) Pause() (*types.Transaction, error) {
	return _Distributor.Contract.Pause(&_Distributor.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Distributor *DistributorTransactorSession) Pause() (*types.Transaction, error) {
	return _Distributor.Contract.Pause(&_Distributor.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Distributor *DistributorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Distributor *DistributorSession) Unpause() (*types.Transaction, error) {
	return _Distributor.Contract.Unpause(&_Distributor.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Distributor *DistributorTransactorSession) Unpause() (*types.Transaction, error) {
	return _Distributor.Contract.Unpause(&_Distributor.TransactOpts)
}

// DistributorClaimedSharesIterator is returned from FilterClaimedShares and is used to iterate over the raw logs and unpacked data for ClaimedShares events raised by the Distributor contract.
type DistributorClaimedSharesIterator struct {
	Event *DistributorClaimedShares // Event containing the contract specifics and raw log

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
func (it *DistributorClaimedSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorClaimedShares)
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
		it.Event = new(DistributorClaimedShares)
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
func (it *DistributorClaimedSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorClaimedSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorClaimedShares represents a ClaimedShares event raised by the Distributor contract.
type DistributorClaimedShares struct {
	User   common.Address
	Period *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimedShares is a free log retrieval operation binding the contract event 0x2cda5c7d9ff70b5c9119b1ba7469a9c367b07c395944a87a72953d9598f5bc62.
//
// Solidity: event ClaimedShares(address user, uint256 period, uint256 shares)
func (_Distributor *DistributorFilterer) FilterClaimedShares(opts *bind.FilterOpts) (*DistributorClaimedSharesIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "ClaimedShares")
	if err != nil {
		return nil, err
	}
	return &DistributorClaimedSharesIterator{contract: _Distributor.contract, event: "ClaimedShares", logs: logs, sub: sub}, nil
}

// WatchClaimedShares is a free log subscription operation binding the contract event 0x2cda5c7d9ff70b5c9119b1ba7469a9c367b07c395944a87a72953d9598f5bc62.
//
// Solidity: event ClaimedShares(address user, uint256 period, uint256 shares)
func (_Distributor *DistributorFilterer) WatchClaimedShares(opts *bind.WatchOpts, sink chan<- *DistributorClaimedShares) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "ClaimedShares")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorClaimedShares)
				if err := _Distributor.contract.UnpackLog(event, "ClaimedShares", log); err != nil {
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
func (_Distributor *DistributorFilterer) ParseClaimedShares(log types.Log) (*DistributorClaimedShares, error) {
	event := new(DistributorClaimedShares)
	if err := _Distributor.contract.UnpackLog(event, "ClaimedShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Distributor contract.
type DistributorInitializedIterator struct {
	Event *DistributorInitialized // Event containing the contract specifics and raw log

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
func (it *DistributorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorInitialized)
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
		it.Event = new(DistributorInitialized)
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
func (it *DistributorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorInitialized represents a Initialized event raised by the Distributor contract.
type DistributorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Distributor *DistributorFilterer) FilterInitialized(opts *bind.FilterOpts) (*DistributorInitializedIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DistributorInitializedIterator{contract: _Distributor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Distributor *DistributorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DistributorInitialized) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorInitialized)
				if err := _Distributor.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Distributor *DistributorFilterer) ParseInitialized(log types.Log) (*DistributorInitialized, error) {
	event := new(DistributorInitialized)
	if err := _Distributor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Distributor contract.
type DistributorPausedIterator struct {
	Event *DistributorPaused // Event containing the contract specifics and raw log

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
func (it *DistributorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorPaused)
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
		it.Event = new(DistributorPaused)
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
func (it *DistributorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorPaused represents a Paused event raised by the Distributor contract.
type DistributorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Distributor *DistributorFilterer) FilterPaused(opts *bind.FilterOpts) (*DistributorPausedIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DistributorPausedIterator{contract: _Distributor.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Distributor *DistributorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DistributorPaused) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorPaused)
				if err := _Distributor.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Distributor *DistributorFilterer) ParsePaused(log types.Log) (*DistributorPaused, error) {
	event := new(DistributorPaused)
	if err := _Distributor.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Distributor contract.
type DistributorUnpausedIterator struct {
	Event *DistributorUnpaused // Event containing the contract specifics and raw log

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
func (it *DistributorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorUnpaused)
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
		it.Event = new(DistributorUnpaused)
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
func (it *DistributorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorUnpaused represents a Unpaused event raised by the Distributor contract.
type DistributorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Distributor *DistributorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DistributorUnpausedIterator, error) {

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DistributorUnpausedIterator{contract: _Distributor.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Distributor *DistributorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DistributorUnpaused) (event.Subscription, error) {

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorUnpaused)
				if err := _Distributor.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Distributor *DistributorFilterer) ParseUnpaused(log types.Log) (*DistributorUnpaused, error) {
	event := new(DistributorUnpaused)
	if err := _Distributor.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

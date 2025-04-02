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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_balancerVault\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balancerVault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVault\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exitPlazaAndBalancer\",\"inputs\":[{\"name\":\"balancerPoolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_plazaPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"contractIAsset[]\"},{\"name\":\"plazaTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountsOut\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"plazaTokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"minbalancerPoolTokenOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"joinBalancerAndPlaza\",\"inputs\":[{\"name\":\"balancerPoolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_plazaPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"contractIAsset[]\"},{\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"plazaTokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"minPlazaTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TokensRedeemed\",\"inputs\":[{\"name\":\"plazaPool\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"redeemedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
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

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) BalancerVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "balancerVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_PlazaSdk *PlazaSdkSession) BalancerVault() (common.Address, error) {
	return _PlazaSdk.Contract.BalancerVault(&_PlazaSdk.CallOpts)
}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) BalancerVault() (common.Address, error) {
	return _PlazaSdk.Contract.BalancerVault(&_PlazaSdk.CallOpts)
}

// ExitPlazaAndBalancer is a paid mutator transaction binding the contract method 0x9f4b851b.
//
// Solidity: function exitPlazaAndBalancer(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256 plazaTokenAmount, uint256[] minAmountsOut, bytes userData, uint8 plazaTokenType, uint256 minbalancerPoolTokenOut) returns()
func (_PlazaSdk *PlazaSdkTransactor) ExitPlazaAndBalancer(opts *bind.TransactOpts, balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, plazaTokenAmount *big.Int, minAmountsOut []*big.Int, userData []byte, plazaTokenType uint8, minbalancerPoolTokenOut *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "exitPlazaAndBalancer", balancerPoolId, _plazaPool, assets, plazaTokenAmount, minAmountsOut, userData, plazaTokenType, minbalancerPoolTokenOut)
}

// ExitPlazaAndBalancer is a paid mutator transaction binding the contract method 0x9f4b851b.
//
// Solidity: function exitPlazaAndBalancer(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256 plazaTokenAmount, uint256[] minAmountsOut, bytes userData, uint8 plazaTokenType, uint256 minbalancerPoolTokenOut) returns()
func (_PlazaSdk *PlazaSdkSession) ExitPlazaAndBalancer(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, plazaTokenAmount *big.Int, minAmountsOut []*big.Int, userData []byte, plazaTokenType uint8, minbalancerPoolTokenOut *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.ExitPlazaAndBalancer(&_PlazaSdk.TransactOpts, balancerPoolId, _plazaPool, assets, plazaTokenAmount, minAmountsOut, userData, plazaTokenType, minbalancerPoolTokenOut)
}

// ExitPlazaAndBalancer is a paid mutator transaction binding the contract method 0x9f4b851b.
//
// Solidity: function exitPlazaAndBalancer(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256 plazaTokenAmount, uint256[] minAmountsOut, bytes userData, uint8 plazaTokenType, uint256 minbalancerPoolTokenOut) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) ExitPlazaAndBalancer(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, plazaTokenAmount *big.Int, minAmountsOut []*big.Int, userData []byte, plazaTokenType uint8, minbalancerPoolTokenOut *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.ExitPlazaAndBalancer(&_PlazaSdk.TransactOpts, balancerPoolId, _plazaPool, assets, plazaTokenAmount, minAmountsOut, userData, plazaTokenType, minbalancerPoolTokenOut)
}

// JoinBalancerAndPlaza is a paid mutator transaction binding the contract method 0x61038bbf.
//
// Solidity: function joinBalancerAndPlaza(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256[] maxAmountsIn, bytes userData, uint8 plazaTokenType, uint256 minPlazaTokens, uint256 deadline) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactor) JoinBalancerAndPlaza(opts *bind.TransactOpts, balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, maxAmountsIn []*big.Int, userData []byte, plazaTokenType uint8, minPlazaTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "joinBalancerAndPlaza", balancerPoolId, _plazaPool, assets, maxAmountsIn, userData, plazaTokenType, minPlazaTokens, deadline)
}

// JoinBalancerAndPlaza is a paid mutator transaction binding the contract method 0x61038bbf.
//
// Solidity: function joinBalancerAndPlaza(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256[] maxAmountsIn, bytes userData, uint8 plazaTokenType, uint256 minPlazaTokens, uint256 deadline) returns(uint256)
func (_PlazaSdk *PlazaSdkSession) JoinBalancerAndPlaza(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, maxAmountsIn []*big.Int, userData []byte, plazaTokenType uint8, minPlazaTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.JoinBalancerAndPlaza(&_PlazaSdk.TransactOpts, balancerPoolId, _plazaPool, assets, maxAmountsIn, userData, plazaTokenType, minPlazaTokens, deadline)
}

// JoinBalancerAndPlaza is a paid mutator transaction binding the contract method 0x61038bbf.
//
// Solidity: function joinBalancerAndPlaza(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256[] maxAmountsIn, bytes userData, uint8 plazaTokenType, uint256 minPlazaTokens, uint256 deadline) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactorSession) JoinBalancerAndPlaza(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, maxAmountsIn []*big.Int, userData []byte, plazaTokenType uint8, minPlazaTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.JoinBalancerAndPlaza(&_PlazaSdk.TransactOpts, balancerPoolId, _plazaPool, assets, maxAmountsIn, userData, plazaTokenType, minPlazaTokens, deadline)
}

// PlazaSdkTokensRedeemedIterator is returned from FilterTokensRedeemed and is used to iterate over the raw logs and unpacked data for TokensRedeemed events raised by the PlazaSdk contract.
type PlazaSdkTokensRedeemedIterator struct {
	Event *PlazaSdkTokensRedeemed // Event containing the contract specifics and raw log

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
func (it *PlazaSdkTokensRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkTokensRedeemed)
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
		it.Event = new(PlazaSdkTokensRedeemed)
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
func (it *PlazaSdkTokensRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkTokensRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkTokensRedeemed represents a TokensRedeemed event raised by the PlazaSdk contract.
type PlazaSdkTokensRedeemed struct {
	PlazaPool       common.Address
	Caller          common.Address
	OnBehalfOf      common.Address
	TokenType       uint8
	DepositedAmount *big.Int
	RedeemedAmount  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokensRedeemed is a free log retrieval operation binding the contract event 0x9de4dbfa6bd5132363914fdf4e6027ec4d7ffc10d197d301d4ef0d5293f5a0e6.
//
// Solidity: event TokensRedeemed(address indexed plazaPool, address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterTokensRedeemed(opts *bind.FilterOpts, plazaPool []common.Address, onBehalfOf []common.Address) (*PlazaSdkTokensRedeemedIterator, error) {

	var plazaPoolRule []interface{}
	for _, plazaPoolItem := range plazaPool {
		plazaPoolRule = append(plazaPoolRule, plazaPoolItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "TokensRedeemed", plazaPoolRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkTokensRedeemedIterator{contract: _PlazaSdk.contract, event: "TokensRedeemed", logs: logs, sub: sub}, nil
}

// WatchTokensRedeemed is a free log subscription operation binding the contract event 0x9de4dbfa6bd5132363914fdf4e6027ec4d7ffc10d197d301d4ef0d5293f5a0e6.
//
// Solidity: event TokensRedeemed(address indexed plazaPool, address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchTokensRedeemed(opts *bind.WatchOpts, sink chan<- *PlazaSdkTokensRedeemed, plazaPool []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var plazaPoolRule []interface{}
	for _, plazaPoolItem := range plazaPool {
		plazaPoolRule = append(plazaPoolRule, plazaPoolItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "TokensRedeemed", plazaPoolRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkTokensRedeemed)
				if err := _PlazaSdk.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
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

// ParseTokensRedeemed is a log parse operation binding the contract event 0x9de4dbfa6bd5132363914fdf4e6027ec4d7ffc10d197d301d4ef0d5293f5a0e6.
//
// Solidity: event TokensRedeemed(address indexed plazaPool, address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseTokensRedeemed(log types.Log) (*PlazaSdkTokensRedeemed, error) {
	event := new(PlazaSdkTokensRedeemed)
	if err := _PlazaSdk.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

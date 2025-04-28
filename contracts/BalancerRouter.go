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

// BalancerRouterMetaData contains all meta data concerning the BalancerRouter contract.
var BalancerRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_balancerVault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balancerVault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVault\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"plazaPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"minPlazaTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"exitPlazaAndBalancer\",\"inputs\":[{\"name\":\"balancerPoolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_plazaPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"contractIAsset[]\"},{\"name\":\"plazaTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountsOut\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"plazaTokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"minbalancerPoolTokenOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"joinBalancerAndPlaza\",\"inputs\":[{\"name\":\"balancerPoolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_plazaPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"contractIAsset[]\"},{\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"plazaTokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"minPlazaTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_plazaPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWETH\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"plazaPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TokensRedeemed\",\"inputs\":[{\"name\":\"plazaPool\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"redeemedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedEth\",\"inputs\":[]}]",
}

// BalancerRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerRouterMetaData.ABI instead.
var BalancerRouterABI = BalancerRouterMetaData.ABI

// BalancerRouter is an auto generated Go binding around an Ethereum contract.
type BalancerRouter struct {
	BalancerRouterCaller     // Read-only binding to the contract
	BalancerRouterTransactor // Write-only binding to the contract
	BalancerRouterFilterer   // Log filterer for contract events
}

// BalancerRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerRouterSession struct {
	Contract     *BalancerRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerRouterCallerSession struct {
	Contract *BalancerRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalancerRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerRouterTransactorSession struct {
	Contract     *BalancerRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BalancerRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerRouterRaw struct {
	Contract *BalancerRouter // Generic contract binding to access the raw methods on
}

// BalancerRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerRouterCallerRaw struct {
	Contract *BalancerRouterCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerRouterTransactorRaw struct {
	Contract *BalancerRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerRouter creates a new instance of BalancerRouter, bound to a specific deployed contract.
func NewBalancerRouter(address common.Address, backend bind.ContractBackend) (*BalancerRouter, error) {
	contract, err := bindBalancerRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerRouter{BalancerRouterCaller: BalancerRouterCaller{contract: contract}, BalancerRouterTransactor: BalancerRouterTransactor{contract: contract}, BalancerRouterFilterer: BalancerRouterFilterer{contract: contract}}, nil
}

// NewBalancerRouterCaller creates a new read-only instance of BalancerRouter, bound to a specific deployed contract.
func NewBalancerRouterCaller(address common.Address, caller bind.ContractCaller) (*BalancerRouterCaller, error) {
	contract, err := bindBalancerRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerRouterCaller{contract: contract}, nil
}

// NewBalancerRouterTransactor creates a new write-only instance of BalancerRouter, bound to a specific deployed contract.
func NewBalancerRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerRouterTransactor, error) {
	contract, err := bindBalancerRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerRouterTransactor{contract: contract}, nil
}

// NewBalancerRouterFilterer creates a new log filterer instance of BalancerRouter, bound to a specific deployed contract.
func NewBalancerRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerRouterFilterer, error) {
	contract, err := bindBalancerRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerRouterFilterer{contract: contract}, nil
}

// bindBalancerRouter binds a generic wrapper to an already deployed contract.
func bindBalancerRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalancerRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerRouter *BalancerRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerRouter.Contract.BalancerRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerRouter *BalancerRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerRouter.Contract.BalancerRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerRouter *BalancerRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerRouter.Contract.BalancerRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerRouter *BalancerRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerRouter *BalancerRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerRouter *BalancerRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerRouter.Contract.contract.Transact(opts, method, params...)
}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_BalancerRouter *BalancerRouterCaller) BalancerVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerRouter.contract.Call(opts, &out, "balancerVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_BalancerRouter *BalancerRouterSession) BalancerVault() (common.Address, error) {
	return _BalancerRouter.Contract.BalancerVault(&_BalancerRouter.CallOpts)
}

// BalancerVault is a free data retrieval call binding the contract method 0x158274a5.
//
// Solidity: function balancerVault() view returns(address)
func (_BalancerRouter *BalancerRouterCallerSession) BalancerVault() (common.Address, error) {
	return _BalancerRouter.Contract.BalancerVault(&_BalancerRouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BalancerRouter *BalancerRouterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerRouter.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BalancerRouter *BalancerRouterSession) Weth() (common.Address, error) {
	return _BalancerRouter.Contract.Weth(&_BalancerRouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BalancerRouter *BalancerRouterCallerSession) Weth() (common.Address, error) {
	return _BalancerRouter.Contract.Weth(&_BalancerRouter.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe56281e8.
//
// Solidity: function deposit(address plazaPool, address depositToken, uint256 amount, uint8 tokenType, uint256 minPlazaTokens) payable returns(uint256)
func (_BalancerRouter *BalancerRouterTransactor) Deposit(opts *bind.TransactOpts, plazaPool common.Address, depositToken common.Address, amount *big.Int, tokenType uint8, minPlazaTokens *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.contract.Transact(opts, "deposit", plazaPool, depositToken, amount, tokenType, minPlazaTokens)
}

// Deposit is a paid mutator transaction binding the contract method 0xe56281e8.
//
// Solidity: function deposit(address plazaPool, address depositToken, uint256 amount, uint8 tokenType, uint256 minPlazaTokens) payable returns(uint256)
func (_BalancerRouter *BalancerRouterSession) Deposit(plazaPool common.Address, depositToken common.Address, amount *big.Int, tokenType uint8, minPlazaTokens *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.Deposit(&_BalancerRouter.TransactOpts, plazaPool, depositToken, amount, tokenType, minPlazaTokens)
}

// Deposit is a paid mutator transaction binding the contract method 0xe56281e8.
//
// Solidity: function deposit(address plazaPool, address depositToken, uint256 amount, uint8 tokenType, uint256 minPlazaTokens) payable returns(uint256)
func (_BalancerRouter *BalancerRouterTransactorSession) Deposit(plazaPool common.Address, depositToken common.Address, amount *big.Int, tokenType uint8, minPlazaTokens *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.Deposit(&_BalancerRouter.TransactOpts, plazaPool, depositToken, amount, tokenType, minPlazaTokens)
}

// ExitPlazaAndBalancer is a paid mutator transaction binding the contract method 0x9f4b851b.
//
// Solidity: function exitPlazaAndBalancer(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256 plazaTokenAmount, uint256[] minAmountsOut, bytes userData, uint8 plazaTokenType, uint256 minbalancerPoolTokenOut) returns()
func (_BalancerRouter *BalancerRouterTransactor) ExitPlazaAndBalancer(opts *bind.TransactOpts, balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, plazaTokenAmount *big.Int, minAmountsOut []*big.Int, userData []byte, plazaTokenType uint8, minbalancerPoolTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.contract.Transact(opts, "exitPlazaAndBalancer", balancerPoolId, _plazaPool, assets, plazaTokenAmount, minAmountsOut, userData, plazaTokenType, minbalancerPoolTokenOut)
}

// ExitPlazaAndBalancer is a paid mutator transaction binding the contract method 0x9f4b851b.
//
// Solidity: function exitPlazaAndBalancer(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256 plazaTokenAmount, uint256[] minAmountsOut, bytes userData, uint8 plazaTokenType, uint256 minbalancerPoolTokenOut) returns()
func (_BalancerRouter *BalancerRouterSession) ExitPlazaAndBalancer(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, plazaTokenAmount *big.Int, minAmountsOut []*big.Int, userData []byte, plazaTokenType uint8, minbalancerPoolTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.ExitPlazaAndBalancer(&_BalancerRouter.TransactOpts, balancerPoolId, _plazaPool, assets, plazaTokenAmount, minAmountsOut, userData, plazaTokenType, minbalancerPoolTokenOut)
}

// ExitPlazaAndBalancer is a paid mutator transaction binding the contract method 0x9f4b851b.
//
// Solidity: function exitPlazaAndBalancer(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256 plazaTokenAmount, uint256[] minAmountsOut, bytes userData, uint8 plazaTokenType, uint256 minbalancerPoolTokenOut) returns()
func (_BalancerRouter *BalancerRouterTransactorSession) ExitPlazaAndBalancer(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, plazaTokenAmount *big.Int, minAmountsOut []*big.Int, userData []byte, plazaTokenType uint8, minbalancerPoolTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.ExitPlazaAndBalancer(&_BalancerRouter.TransactOpts, balancerPoolId, _plazaPool, assets, plazaTokenAmount, minAmountsOut, userData, plazaTokenType, minbalancerPoolTokenOut)
}

// JoinBalancerAndPlaza is a paid mutator transaction binding the contract method 0x61038bbf.
//
// Solidity: function joinBalancerAndPlaza(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256[] maxAmountsIn, bytes userData, uint8 plazaTokenType, uint256 minPlazaTokens, uint256 deadline) returns(uint256)
func (_BalancerRouter *BalancerRouterTransactor) JoinBalancerAndPlaza(opts *bind.TransactOpts, balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, maxAmountsIn []*big.Int, userData []byte, plazaTokenType uint8, minPlazaTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.contract.Transact(opts, "joinBalancerAndPlaza", balancerPoolId, _plazaPool, assets, maxAmountsIn, userData, plazaTokenType, minPlazaTokens, deadline)
}

// JoinBalancerAndPlaza is a paid mutator transaction binding the contract method 0x61038bbf.
//
// Solidity: function joinBalancerAndPlaza(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256[] maxAmountsIn, bytes userData, uint8 plazaTokenType, uint256 minPlazaTokens, uint256 deadline) returns(uint256)
func (_BalancerRouter *BalancerRouterSession) JoinBalancerAndPlaza(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, maxAmountsIn []*big.Int, userData []byte, plazaTokenType uint8, minPlazaTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.JoinBalancerAndPlaza(&_BalancerRouter.TransactOpts, balancerPoolId, _plazaPool, assets, maxAmountsIn, userData, plazaTokenType, minPlazaTokens, deadline)
}

// JoinBalancerAndPlaza is a paid mutator transaction binding the contract method 0x61038bbf.
//
// Solidity: function joinBalancerAndPlaza(bytes32 balancerPoolId, address _plazaPool, address[] assets, uint256[] maxAmountsIn, bytes userData, uint8 plazaTokenType, uint256 minPlazaTokens, uint256 deadline) returns(uint256)
func (_BalancerRouter *BalancerRouterTransactorSession) JoinBalancerAndPlaza(balancerPoolId [32]byte, _plazaPool common.Address, assets []common.Address, maxAmountsIn []*big.Int, userData []byte, plazaTokenType uint8, minPlazaTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.JoinBalancerAndPlaza(&_BalancerRouter.TransactOpts, balancerPoolId, _plazaPool, assets, maxAmountsIn, userData, plazaTokenType, minPlazaTokens, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0xed325821.
//
// Solidity: function swap(address _plazaPool, uint8 tokenType, uint256 tokenAmount, uint256 minAmountOut, uint256 deadline) returns(uint256)
func (_BalancerRouter *BalancerRouterTransactor) Swap(opts *bind.TransactOpts, _plazaPool common.Address, tokenType uint8, tokenAmount *big.Int, minAmountOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.contract.Transact(opts, "swap", _plazaPool, tokenType, tokenAmount, minAmountOut, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0xed325821.
//
// Solidity: function swap(address _plazaPool, uint8 tokenType, uint256 tokenAmount, uint256 minAmountOut, uint256 deadline) returns(uint256)
func (_BalancerRouter *BalancerRouterSession) Swap(_plazaPool common.Address, tokenType uint8, tokenAmount *big.Int, minAmountOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.Swap(&_BalancerRouter.TransactOpts, _plazaPool, tokenType, tokenAmount, minAmountOut, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0xed325821.
//
// Solidity: function swap(address _plazaPool, uint8 tokenType, uint256 tokenAmount, uint256 minAmountOut, uint256 deadline) returns(uint256)
func (_BalancerRouter *BalancerRouterTransactorSession) Swap(_plazaPool common.Address, tokenType uint8, tokenAmount *big.Int, minAmountOut *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.Swap(&_BalancerRouter.TransactOpts, _plazaPool, tokenType, tokenAmount, minAmountOut, deadline)
}

// Withdraw is a paid mutator transaction binding the contract method 0xaa333ba6.
//
// Solidity: function withdraw(address plazaPool, uint8 tokenType, uint256 amount, address withdrawToken, uint256 minAmount) returns()
func (_BalancerRouter *BalancerRouterTransactor) Withdraw(opts *bind.TransactOpts, plazaPool common.Address, tokenType uint8, amount *big.Int, withdrawToken common.Address, minAmount *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.contract.Transact(opts, "withdraw", plazaPool, tokenType, amount, withdrawToken, minAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xaa333ba6.
//
// Solidity: function withdraw(address plazaPool, uint8 tokenType, uint256 amount, address withdrawToken, uint256 minAmount) returns()
func (_BalancerRouter *BalancerRouterSession) Withdraw(plazaPool common.Address, tokenType uint8, amount *big.Int, withdrawToken common.Address, minAmount *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.Withdraw(&_BalancerRouter.TransactOpts, plazaPool, tokenType, amount, withdrawToken, minAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xaa333ba6.
//
// Solidity: function withdraw(address plazaPool, uint8 tokenType, uint256 amount, address withdrawToken, uint256 minAmount) returns()
func (_BalancerRouter *BalancerRouterTransactorSession) Withdraw(plazaPool common.Address, tokenType uint8, amount *big.Int, withdrawToken common.Address, minAmount *big.Int) (*types.Transaction, error) {
	return _BalancerRouter.Contract.Withdraw(&_BalancerRouter.TransactOpts, plazaPool, tokenType, amount, withdrawToken, minAmount)
}

// BalancerRouterTokensRedeemedIterator is returned from FilterTokensRedeemed and is used to iterate over the raw logs and unpacked data for TokensRedeemed events raised by the BalancerRouter contract.
type BalancerRouterTokensRedeemedIterator struct {
	Event *BalancerRouterTokensRedeemed // Event containing the contract specifics and raw log

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
func (it *BalancerRouterTokensRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerRouterTokensRedeemed)
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
		it.Event = new(BalancerRouterTokensRedeemed)
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
func (it *BalancerRouterTokensRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerRouterTokensRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerRouterTokensRedeemed represents a TokensRedeemed event raised by the BalancerRouter contract.
type BalancerRouterTokensRedeemed struct {
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
func (_BalancerRouter *BalancerRouterFilterer) FilterTokensRedeemed(opts *bind.FilterOpts, plazaPool []common.Address, onBehalfOf []common.Address) (*BalancerRouterTokensRedeemedIterator, error) {

	var plazaPoolRule []interface{}
	for _, plazaPoolItem := range plazaPool {
		plazaPoolRule = append(plazaPoolRule, plazaPoolItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _BalancerRouter.contract.FilterLogs(opts, "TokensRedeemed", plazaPoolRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &BalancerRouterTokensRedeemedIterator{contract: _BalancerRouter.contract, event: "TokensRedeemed", logs: logs, sub: sub}, nil
}

// WatchTokensRedeemed is a free log subscription operation binding the contract event 0x9de4dbfa6bd5132363914fdf4e6027ec4d7ffc10d197d301d4ef0d5293f5a0e6.
//
// Solidity: event TokensRedeemed(address indexed plazaPool, address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_BalancerRouter *BalancerRouterFilterer) WatchTokensRedeemed(opts *bind.WatchOpts, sink chan<- *BalancerRouterTokensRedeemed, plazaPool []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var plazaPoolRule []interface{}
	for _, plazaPoolItem := range plazaPool {
		plazaPoolRule = append(plazaPoolRule, plazaPoolItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _BalancerRouter.contract.WatchLogs(opts, "TokensRedeemed", plazaPoolRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerRouterTokensRedeemed)
				if err := _BalancerRouter.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
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
func (_BalancerRouter *BalancerRouterFilterer) ParseTokensRedeemed(log types.Log) (*BalancerRouterTokensRedeemed, error) {
	event := new(BalancerRouterTokensRedeemed)
	if err := _BalancerRouter.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

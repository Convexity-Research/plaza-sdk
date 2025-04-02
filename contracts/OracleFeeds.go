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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feedHeartbeats\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeeds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"tokenA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenB\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"heartbeat\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.DEFAULTADMINROLE(&_PlazaSdk.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.DEFAULTADMINROLE(&_PlazaSdk.CallOpts)
}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCaller) GOVROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "GOV_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkSession) GOVROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.GOVROLE(&_PlazaSdk.CallOpts)
}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCallerSession) GOVROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.GOVROLE(&_PlazaSdk.CallOpts)
}

// FeedHeartbeats is a free data retrieval call binding the contract method 0xa83c3ab5.
//
// Solidity: function feedHeartbeats(address ) view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) FeedHeartbeats(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "feedHeartbeats", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeedHeartbeats is a free data retrieval call binding the contract method 0xa83c3ab5.
//
// Solidity: function feedHeartbeats(address ) view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) FeedHeartbeats(arg0 common.Address) (*big.Int, error) {
	return _PlazaSdk.Contract.FeedHeartbeats(&_PlazaSdk.CallOpts, arg0)
}

// FeedHeartbeats is a free data retrieval call binding the contract method 0xa83c3ab5.
//
// Solidity: function feedHeartbeats(address ) view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) FeedHeartbeats(arg0 common.Address) (*big.Int, error) {
	return _PlazaSdk.Contract.FeedHeartbeats(&_PlazaSdk.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PlazaSdk *PlazaSdkCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PlazaSdk *PlazaSdkSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PlazaSdk.Contract.GetRoleAdmin(&_PlazaSdk.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PlazaSdk *PlazaSdkCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PlazaSdk.Contract.GetRoleAdmin(&_PlazaSdk.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PlazaSdk *PlazaSdkCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PlazaSdk *PlazaSdkSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PlazaSdk.Contract.HasRole(&_PlazaSdk.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PlazaSdk *PlazaSdkCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PlazaSdk.Contract.HasRole(&_PlazaSdk.CallOpts, role, account)
}

// PriceFeeds is a free data retrieval call binding the contract method 0xf90c6906.
//
// Solidity: function priceFeeds(address , address ) view returns(address)
func (_PlazaSdk *PlazaSdkCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "priceFeeds", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0xf90c6906.
//
// Solidity: function priceFeeds(address , address ) view returns(address)
func (_PlazaSdk *PlazaSdkSession) PriceFeeds(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PlazaSdk.Contract.PriceFeeds(&_PlazaSdk.CallOpts, arg0, arg1)
}

// PriceFeeds is a free data retrieval call binding the contract method 0xf90c6906.
//
// Solidity: function priceFeeds(address , address ) view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) PriceFeeds(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _PlazaSdk.Contract.PriceFeeds(&_PlazaSdk.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PlazaSdk *PlazaSdkCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PlazaSdk *PlazaSdkSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PlazaSdk.Contract.SupportsInterface(&_PlazaSdk.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PlazaSdk *PlazaSdkCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PlazaSdk.Contract.SupportsInterface(&_PlazaSdk.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PlazaSdk *PlazaSdkTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PlazaSdk *PlazaSdkSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.GrantRole(&_PlazaSdk.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.GrantRole(&_PlazaSdk.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PlazaSdk *PlazaSdkTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PlazaSdk *PlazaSdkSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.RenounceRole(&_PlazaSdk.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.RenounceRole(&_PlazaSdk.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PlazaSdk *PlazaSdkTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PlazaSdk *PlazaSdkSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.RevokeRole(&_PlazaSdk.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.RevokeRole(&_PlazaSdk.TransactOpts, role, account)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x8fd525ec.
//
// Solidity: function setPriceFeed(address tokenA, address tokenB, address priceFeed, uint256 heartbeat) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetPriceFeed(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, priceFeed common.Address, heartbeat *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setPriceFeed", tokenA, tokenB, priceFeed, heartbeat)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x8fd525ec.
//
// Solidity: function setPriceFeed(address tokenA, address tokenB, address priceFeed, uint256 heartbeat) returns()
func (_PlazaSdk *PlazaSdkSession) SetPriceFeed(tokenA common.Address, tokenB common.Address, priceFeed common.Address, heartbeat *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetPriceFeed(&_PlazaSdk.TransactOpts, tokenA, tokenB, priceFeed, heartbeat)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x8fd525ec.
//
// Solidity: function setPriceFeed(address tokenA, address tokenB, address priceFeed, uint256 heartbeat) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetPriceFeed(tokenA common.Address, tokenB common.Address, priceFeed common.Address, heartbeat *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetPriceFeed(&_PlazaSdk.TransactOpts, tokenA, tokenB, priceFeed, heartbeat)
}

// PlazaSdkRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PlazaSdk contract.
type PlazaSdkRoleAdminChangedIterator struct {
	Event *PlazaSdkRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PlazaSdkRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkRoleAdminChanged)
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
		it.Event = new(PlazaSdkRoleAdminChanged)
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
func (it *PlazaSdkRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkRoleAdminChanged represents a RoleAdminChanged event raised by the PlazaSdk contract.
type PlazaSdkRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PlazaSdk *PlazaSdkFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PlazaSdkRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkRoleAdminChangedIterator{contract: _PlazaSdk.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PlazaSdk *PlazaSdkFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PlazaSdkRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkRoleAdminChanged)
				if err := _PlazaSdk.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PlazaSdk *PlazaSdkFilterer) ParseRoleAdminChanged(log types.Log) (*PlazaSdkRoleAdminChanged, error) {
	event := new(PlazaSdkRoleAdminChanged)
	if err := _PlazaSdk.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PlazaSdk contract.
type PlazaSdkRoleGrantedIterator struct {
	Event *PlazaSdkRoleGranted // Event containing the contract specifics and raw log

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
func (it *PlazaSdkRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkRoleGranted)
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
		it.Event = new(PlazaSdkRoleGranted)
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
func (it *PlazaSdkRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkRoleGranted represents a RoleGranted event raised by the PlazaSdk contract.
type PlazaSdkRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlazaSdk *PlazaSdkFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PlazaSdkRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkRoleGrantedIterator{contract: _PlazaSdk.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlazaSdk *PlazaSdkFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PlazaSdkRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkRoleGranted)
				if err := _PlazaSdk.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlazaSdk *PlazaSdkFilterer) ParseRoleGranted(log types.Log) (*PlazaSdkRoleGranted, error) {
	event := new(PlazaSdkRoleGranted)
	if err := _PlazaSdk.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PlazaSdk contract.
type PlazaSdkRoleRevokedIterator struct {
	Event *PlazaSdkRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PlazaSdkRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkRoleRevoked)
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
		it.Event = new(PlazaSdkRoleRevoked)
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
func (it *PlazaSdkRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkRoleRevoked represents a RoleRevoked event raised by the PlazaSdk contract.
type PlazaSdkRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlazaSdk *PlazaSdkFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PlazaSdkRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkRoleRevokedIterator{contract: _PlazaSdk.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlazaSdk *PlazaSdkFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PlazaSdkRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkRoleRevoked)
				if err := _PlazaSdk.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PlazaSdk *PlazaSdkFilterer) ParseRoleRevoked(log types.Log) (*PlazaSdkRoleRevoked, error) {
	event := new(PlazaSdkRoleRevoked)
	if err := _PlazaSdk.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

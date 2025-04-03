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

// OracleFeedsMetaData contains all meta data concerning the OracleFeeds contract.
var OracleFeedsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feedHeartbeats\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeeds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"tokenA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenB\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"heartbeat\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// OracleFeedsABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleFeedsMetaData.ABI instead.
var OracleFeedsABI = OracleFeedsMetaData.ABI

// OracleFeeds is an auto generated Go binding around an Ethereum contract.
type OracleFeeds struct {
	OracleFeedsCaller     // Read-only binding to the contract
	OracleFeedsTransactor // Write-only binding to the contract
	OracleFeedsFilterer   // Log filterer for contract events
}

// OracleFeedsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleFeedsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFeedsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleFeedsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFeedsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFeedsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFeedsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleFeedsSession struct {
	Contract     *OracleFeeds      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleFeedsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleFeedsCallerSession struct {
	Contract *OracleFeedsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OracleFeedsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleFeedsTransactorSession struct {
	Contract     *OracleFeedsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OracleFeedsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleFeedsRaw struct {
	Contract *OracleFeeds // Generic contract binding to access the raw methods on
}

// OracleFeedsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleFeedsCallerRaw struct {
	Contract *OracleFeedsCaller // Generic read-only contract binding to access the raw methods on
}

// OracleFeedsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleFeedsTransactorRaw struct {
	Contract *OracleFeedsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleFeeds creates a new instance of OracleFeeds, bound to a specific deployed contract.
func NewOracleFeeds(address common.Address, backend bind.ContractBackend) (*OracleFeeds, error) {
	contract, err := bindOracleFeeds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleFeeds{OracleFeedsCaller: OracleFeedsCaller{contract: contract}, OracleFeedsTransactor: OracleFeedsTransactor{contract: contract}, OracleFeedsFilterer: OracleFeedsFilterer{contract: contract}}, nil
}

// NewOracleFeedsCaller creates a new read-only instance of OracleFeeds, bound to a specific deployed contract.
func NewOracleFeedsCaller(address common.Address, caller bind.ContractCaller) (*OracleFeedsCaller, error) {
	contract, err := bindOracleFeeds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleFeedsCaller{contract: contract}, nil
}

// NewOracleFeedsTransactor creates a new write-only instance of OracleFeeds, bound to a specific deployed contract.
func NewOracleFeedsTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleFeedsTransactor, error) {
	contract, err := bindOracleFeeds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleFeedsTransactor{contract: contract}, nil
}

// NewOracleFeedsFilterer creates a new log filterer instance of OracleFeeds, bound to a specific deployed contract.
func NewOracleFeedsFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFeedsFilterer, error) {
	contract, err := bindOracleFeeds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFeedsFilterer{contract: contract}, nil
}

// bindOracleFeeds binds a generic wrapper to an already deployed contract.
func bindOracleFeeds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleFeedsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleFeeds *OracleFeedsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleFeeds.Contract.OracleFeedsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleFeeds *OracleFeedsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleFeeds.Contract.OracleFeedsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleFeeds *OracleFeedsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleFeeds.Contract.OracleFeedsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleFeeds *OracleFeedsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleFeeds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleFeeds *OracleFeedsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleFeeds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleFeeds *OracleFeedsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleFeeds.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OracleFeeds *OracleFeedsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OracleFeeds.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OracleFeeds *OracleFeedsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OracleFeeds.Contract.DEFAULTADMINROLE(&_OracleFeeds.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OracleFeeds *OracleFeedsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OracleFeeds.Contract.DEFAULTADMINROLE(&_OracleFeeds.CallOpts)
}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_OracleFeeds *OracleFeedsCaller) GOVROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OracleFeeds.contract.Call(opts, &out, "GOV_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_OracleFeeds *OracleFeedsSession) GOVROLE() ([32]byte, error) {
	return _OracleFeeds.Contract.GOVROLE(&_OracleFeeds.CallOpts)
}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_OracleFeeds *OracleFeedsCallerSession) GOVROLE() ([32]byte, error) {
	return _OracleFeeds.Contract.GOVROLE(&_OracleFeeds.CallOpts)
}

// FeedHeartbeats is a free data retrieval call binding the contract method 0xa83c3ab5.
//
// Solidity: function feedHeartbeats(address ) view returns(uint256)
func (_OracleFeeds *OracleFeedsCaller) FeedHeartbeats(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OracleFeeds.contract.Call(opts, &out, "feedHeartbeats", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeedHeartbeats is a free data retrieval call binding the contract method 0xa83c3ab5.
//
// Solidity: function feedHeartbeats(address ) view returns(uint256)
func (_OracleFeeds *OracleFeedsSession) FeedHeartbeats(arg0 common.Address) (*big.Int, error) {
	return _OracleFeeds.Contract.FeedHeartbeats(&_OracleFeeds.CallOpts, arg0)
}

// FeedHeartbeats is a free data retrieval call binding the contract method 0xa83c3ab5.
//
// Solidity: function feedHeartbeats(address ) view returns(uint256)
func (_OracleFeeds *OracleFeedsCallerSession) FeedHeartbeats(arg0 common.Address) (*big.Int, error) {
	return _OracleFeeds.Contract.FeedHeartbeats(&_OracleFeeds.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OracleFeeds *OracleFeedsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OracleFeeds.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OracleFeeds *OracleFeedsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OracleFeeds.Contract.GetRoleAdmin(&_OracleFeeds.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OracleFeeds *OracleFeedsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OracleFeeds.Contract.GetRoleAdmin(&_OracleFeeds.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OracleFeeds *OracleFeedsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OracleFeeds.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OracleFeeds *OracleFeedsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OracleFeeds.Contract.HasRole(&_OracleFeeds.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OracleFeeds *OracleFeedsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OracleFeeds.Contract.HasRole(&_OracleFeeds.CallOpts, role, account)
}

// PriceFeeds is a free data retrieval call binding the contract method 0xf90c6906.
//
// Solidity: function priceFeeds(address , address ) view returns(address)
func (_OracleFeeds *OracleFeedsCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _OracleFeeds.contract.Call(opts, &out, "priceFeeds", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0xf90c6906.
//
// Solidity: function priceFeeds(address , address ) view returns(address)
func (_OracleFeeds *OracleFeedsSession) PriceFeeds(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _OracleFeeds.Contract.PriceFeeds(&_OracleFeeds.CallOpts, arg0, arg1)
}

// PriceFeeds is a free data retrieval call binding the contract method 0xf90c6906.
//
// Solidity: function priceFeeds(address , address ) view returns(address)
func (_OracleFeeds *OracleFeedsCallerSession) PriceFeeds(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _OracleFeeds.Contract.PriceFeeds(&_OracleFeeds.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OracleFeeds *OracleFeedsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OracleFeeds.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OracleFeeds *OracleFeedsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OracleFeeds.Contract.SupportsInterface(&_OracleFeeds.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OracleFeeds *OracleFeedsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OracleFeeds.Contract.SupportsInterface(&_OracleFeeds.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OracleFeeds *OracleFeedsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleFeeds.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OracleFeeds *OracleFeedsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleFeeds.Contract.GrantRole(&_OracleFeeds.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OracleFeeds *OracleFeedsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleFeeds.Contract.GrantRole(&_OracleFeeds.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OracleFeeds *OracleFeedsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OracleFeeds.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OracleFeeds *OracleFeedsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OracleFeeds.Contract.RenounceRole(&_OracleFeeds.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OracleFeeds *OracleFeedsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OracleFeeds.Contract.RenounceRole(&_OracleFeeds.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OracleFeeds *OracleFeedsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleFeeds.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OracleFeeds *OracleFeedsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleFeeds.Contract.RevokeRole(&_OracleFeeds.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OracleFeeds *OracleFeedsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OracleFeeds.Contract.RevokeRole(&_OracleFeeds.TransactOpts, role, account)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x8fd525ec.
//
// Solidity: function setPriceFeed(address tokenA, address tokenB, address priceFeed, uint256 heartbeat) returns()
func (_OracleFeeds *OracleFeedsTransactor) SetPriceFeed(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, priceFeed common.Address, heartbeat *big.Int) (*types.Transaction, error) {
	return _OracleFeeds.contract.Transact(opts, "setPriceFeed", tokenA, tokenB, priceFeed, heartbeat)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x8fd525ec.
//
// Solidity: function setPriceFeed(address tokenA, address tokenB, address priceFeed, uint256 heartbeat) returns()
func (_OracleFeeds *OracleFeedsSession) SetPriceFeed(tokenA common.Address, tokenB common.Address, priceFeed common.Address, heartbeat *big.Int) (*types.Transaction, error) {
	return _OracleFeeds.Contract.SetPriceFeed(&_OracleFeeds.TransactOpts, tokenA, tokenB, priceFeed, heartbeat)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x8fd525ec.
//
// Solidity: function setPriceFeed(address tokenA, address tokenB, address priceFeed, uint256 heartbeat) returns()
func (_OracleFeeds *OracleFeedsTransactorSession) SetPriceFeed(tokenA common.Address, tokenB common.Address, priceFeed common.Address, heartbeat *big.Int) (*types.Transaction, error) {
	return _OracleFeeds.Contract.SetPriceFeed(&_OracleFeeds.TransactOpts, tokenA, tokenB, priceFeed, heartbeat)
}

// OracleFeedsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OracleFeeds contract.
type OracleFeedsRoleAdminChangedIterator struct {
	Event *OracleFeedsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OracleFeedsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFeedsRoleAdminChanged)
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
		it.Event = new(OracleFeedsRoleAdminChanged)
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
func (it *OracleFeedsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFeedsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFeedsRoleAdminChanged represents a RoleAdminChanged event raised by the OracleFeeds contract.
type OracleFeedsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OracleFeeds *OracleFeedsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OracleFeedsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _OracleFeeds.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OracleFeedsRoleAdminChangedIterator{contract: _OracleFeeds.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OracleFeeds *OracleFeedsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OracleFeedsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OracleFeeds.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFeedsRoleAdminChanged)
				if err := _OracleFeeds.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_OracleFeeds *OracleFeedsFilterer) ParseRoleAdminChanged(log types.Log) (*OracleFeedsRoleAdminChanged, error) {
	event := new(OracleFeedsRoleAdminChanged)
	if err := _OracleFeeds.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFeedsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OracleFeeds contract.
type OracleFeedsRoleGrantedIterator struct {
	Event *OracleFeedsRoleGranted // Event containing the contract specifics and raw log

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
func (it *OracleFeedsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFeedsRoleGranted)
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
		it.Event = new(OracleFeedsRoleGranted)
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
func (it *OracleFeedsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFeedsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFeedsRoleGranted represents a RoleGranted event raised by the OracleFeeds contract.
type OracleFeedsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleFeeds *OracleFeedsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OracleFeedsRoleGrantedIterator, error) {

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

	logs, sub, err := _OracleFeeds.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleFeedsRoleGrantedIterator{contract: _OracleFeeds.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleFeeds *OracleFeedsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OracleFeedsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OracleFeeds.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFeedsRoleGranted)
				if err := _OracleFeeds.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_OracleFeeds *OracleFeedsFilterer) ParseRoleGranted(log types.Log) (*OracleFeedsRoleGranted, error) {
	event := new(OracleFeedsRoleGranted)
	if err := _OracleFeeds.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleFeedsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OracleFeeds contract.
type OracleFeedsRoleRevokedIterator struct {
	Event *OracleFeedsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OracleFeedsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleFeedsRoleRevoked)
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
		it.Event = new(OracleFeedsRoleRevoked)
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
func (it *OracleFeedsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleFeedsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleFeedsRoleRevoked represents a RoleRevoked event raised by the OracleFeeds contract.
type OracleFeedsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleFeeds *OracleFeedsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OracleFeedsRoleRevokedIterator, error) {

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

	logs, sub, err := _OracleFeeds.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleFeedsRoleRevokedIterator{contract: _OracleFeeds.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OracleFeeds *OracleFeedsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OracleFeedsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OracleFeeds.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleFeedsRoleRevoked)
				if err := _OracleFeeds.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_OracleFeeds *OracleFeedsFilterer) ParseRoleRevoked(log types.Log) (*OracleFeedsRoleRevoked, error) {
	event := new(OracleFeedsRoleRevoked)
	if err := _OracleFeeds.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// ChainlinkFeedMetaData contains all meta data concerning the ChainlinkFeed contract.
var ChainlinkFeedMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_removeFirstElement\",\"inputs\":[{\"name\":\"arr\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"description\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getOracleDecimals\",\"inputs\":[{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOraclePrice\",\"inputs\":[{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoundData\",\"inputs\":[{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getSingleAssetPrice\",\"inputs\":[{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_oracleFeeds\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestRoundData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracleFeeds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBalancerPoolAddress\",\"inputs\":[{\"name\":\"_balancerPoolAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BaseOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExponentOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExponent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoFeedFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPriceFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PriceTooLargeForIntConversion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProductOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroInvariant\",\"inputs\":[]}]",
}

// ChainlinkFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainlinkFeedMetaData.ABI instead.
var ChainlinkFeedABI = ChainlinkFeedMetaData.ABI

// ChainlinkFeed is an auto generated Go binding around an Ethereum contract.
type ChainlinkFeed struct {
	ChainlinkFeedCaller     // Read-only binding to the contract
	ChainlinkFeedTransactor // Write-only binding to the contract
	ChainlinkFeedFilterer   // Log filterer for contract events
}

// ChainlinkFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkFeedSession struct {
	Contract     *ChainlinkFeed    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainlinkFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkFeedCallerSession struct {
	Contract *ChainlinkFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ChainlinkFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkFeedTransactorSession struct {
	Contract     *ChainlinkFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ChainlinkFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkFeedRaw struct {
	Contract *ChainlinkFeed // Generic contract binding to access the raw methods on
}

// ChainlinkFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkFeedCallerRaw struct {
	Contract *ChainlinkFeedCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkFeedTransactorRaw struct {
	Contract *ChainlinkFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkFeed creates a new instance of ChainlinkFeed, bound to a specific deployed contract.
func NewChainlinkFeed(address common.Address, backend bind.ContractBackend) (*ChainlinkFeed, error) {
	contract, err := bindChainlinkFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeed{ChainlinkFeedCaller: ChainlinkFeedCaller{contract: contract}, ChainlinkFeedTransactor: ChainlinkFeedTransactor{contract: contract}, ChainlinkFeedFilterer: ChainlinkFeedFilterer{contract: contract}}, nil
}

// NewChainlinkFeedCaller creates a new read-only instance of ChainlinkFeed, bound to a specific deployed contract.
func NewChainlinkFeedCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkFeedCaller, error) {
	contract, err := bindChainlinkFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedCaller{contract: contract}, nil
}

// NewChainlinkFeedTransactor creates a new write-only instance of ChainlinkFeed, bound to a specific deployed contract.
func NewChainlinkFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkFeedTransactor, error) {
	contract, err := bindChainlinkFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedTransactor{contract: contract}, nil
}

// NewChainlinkFeedFilterer creates a new log filterer instance of ChainlinkFeed, bound to a specific deployed contract.
func NewChainlinkFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkFeedFilterer, error) {
	contract, err := bindChainlinkFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedFilterer{contract: contract}, nil
}

// bindChainlinkFeed binds a generic wrapper to an already deployed contract.
func bindChainlinkFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChainlinkFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkFeed *ChainlinkFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkFeed.Contract.ChainlinkFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkFeed *ChainlinkFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.ChainlinkFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkFeed *ChainlinkFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.ChainlinkFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkFeed *ChainlinkFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkFeed *ChainlinkFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkFeed *ChainlinkFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.contract.Transact(opts, method, params...)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCaller) ETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "ETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedSession) ETH() (common.Address, error) {
	return _ChainlinkFeed.Contract.ETH(&_ChainlinkFeed.CallOpts)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) ETH() (common.Address, error) {
	return _ChainlinkFeed.Contract.ETH(&_ChainlinkFeed.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChainlinkFeed *ChainlinkFeedCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChainlinkFeed *ChainlinkFeedSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChainlinkFeed.Contract.UPGRADEINTERFACEVERSION(&_ChainlinkFeed.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ChainlinkFeed.Contract.UPGRADEINTERFACEVERSION(&_ChainlinkFeed.CallOpts)
}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCaller) USD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "USD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedSession) USD() (common.Address, error) {
	return _ChainlinkFeed.Contract.USD(&_ChainlinkFeed.CallOpts)
}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) USD() (common.Address, error) {
	return _ChainlinkFeed.Contract.USD(&_ChainlinkFeed.CallOpts)
}

// RemoveFirstElement is a free data retrieval call binding the contract method 0x8ea9e265.
//
// Solidity: function _removeFirstElement(uint256[] arr) pure returns(uint256[])
func (_ChainlinkFeed *ChainlinkFeedCaller) RemoveFirstElement(opts *bind.CallOpts, arr []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "_removeFirstElement", arr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RemoveFirstElement is a free data retrieval call binding the contract method 0x8ea9e265.
//
// Solidity: function _removeFirstElement(uint256[] arr) pure returns(uint256[])
func (_ChainlinkFeed *ChainlinkFeedSession) RemoveFirstElement(arr []*big.Int) ([]*big.Int, error) {
	return _ChainlinkFeed.Contract.RemoveFirstElement(&_ChainlinkFeed.CallOpts, arr)
}

// RemoveFirstElement is a free data retrieval call binding the contract method 0x8ea9e265.
//
// Solidity: function _removeFirstElement(uint256[] arr) pure returns(uint256[])
func (_ChainlinkFeed *ChainlinkFeedCallerSession) RemoveFirstElement(arr []*big.Int) ([]*big.Int, error) {
	return _ChainlinkFeed.Contract.RemoveFirstElement(&_ChainlinkFeed.CallOpts, arr)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkFeed *ChainlinkFeedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkFeed *ChainlinkFeedSession) Decimals() (uint8, error) {
	return _ChainlinkFeed.Contract.Decimals(&_ChainlinkFeed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) Decimals() (uint8, error) {
	return _ChainlinkFeed.Contract.Decimals(&_ChainlinkFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkFeed *ChainlinkFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkFeed *ChainlinkFeedSession) Description() (string, error) {
	return _ChainlinkFeed.Contract.Description(&_ChainlinkFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) Description() (string, error) {
	return _ChainlinkFeed.Contract.Description(&_ChainlinkFeed.CallOpts)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_ChainlinkFeed *ChainlinkFeedCaller) GetOracleDecimals(opts *bind.CallOpts, quote common.Address, base common.Address) (uint8, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "getOracleDecimals", quote, base)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_ChainlinkFeed *ChainlinkFeedSession) GetOracleDecimals(quote common.Address, base common.Address) (uint8, error) {
	return _ChainlinkFeed.Contract.GetOracleDecimals(&_ChainlinkFeed.CallOpts, quote, base)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) GetOracleDecimals(quote common.Address, base common.Address) (uint8, error) {
	return _ChainlinkFeed.Contract.GetOracleDecimals(&_ChainlinkFeed.CallOpts, quote, base)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedCaller) GetOraclePrice(opts *bind.CallOpts, quote common.Address, base common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "getOraclePrice", quote, base)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedSession) GetOraclePrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _ChainlinkFeed.Contract.GetOraclePrice(&_ChainlinkFeed.CallOpts, quote, base)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) GetOraclePrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _ChainlinkFeed.Contract.GetOraclePrice(&_ChainlinkFeed.CallOpts, quote, base)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) pure returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkFeed *ChainlinkFeedCaller) GetRoundData(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "getRoundData", arg0)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) pure returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkFeed *ChainlinkFeedSession) GetRoundData(arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkFeed.Contract.GetRoundData(&_ChainlinkFeed.CallOpts, arg0)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) pure returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) GetRoundData(arg0 *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkFeed.Contract.GetRoundData(&_ChainlinkFeed.CallOpts, arg0)
}

// GetSingleAssetPrice is a free data retrieval call binding the contract method 0x9907cba8.
//
// Solidity: function getSingleAssetPrice(address quote, address base) view returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedCaller) GetSingleAssetPrice(opts *bind.CallOpts, quote common.Address, base common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "getSingleAssetPrice", quote, base)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSingleAssetPrice is a free data retrieval call binding the contract method 0x9907cba8.
//
// Solidity: function getSingleAssetPrice(address quote, address base) view returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedSession) GetSingleAssetPrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _ChainlinkFeed.Contract.GetSingleAssetPrice(&_ChainlinkFeed.CallOpts, quote, base)
}

// GetSingleAssetPrice is a free data retrieval call binding the contract method 0x9907cba8.
//
// Solidity: function getSingleAssetPrice(address quote, address base) view returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) GetSingleAssetPrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _ChainlinkFeed.Contract.GetSingleAssetPrice(&_ChainlinkFeed.CallOpts, quote, base)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkFeed *ChainlinkFeedCaller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "latestRoundData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkFeed *ChainlinkFeedSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkFeed.Contract.LatestRoundData(&_ChainlinkFeed.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ChainlinkFeed.Contract.LatestRoundData(&_ChainlinkFeed.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCaller) OracleFeeds(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "oracleFeeds")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedSession) OracleFeeds() (common.Address, error) {
	return _ChainlinkFeed.Contract.OracleFeeds(&_ChainlinkFeed.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) OracleFeeds() (common.Address, error) {
	return _ChainlinkFeed.Contract.OracleFeeds(&_ChainlinkFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedSession) Owner() (common.Address, error) {
	return _ChainlinkFeed.Contract.Owner(&_ChainlinkFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) Owner() (common.Address, error) {
	return _ChainlinkFeed.Contract.Owner(&_ChainlinkFeed.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCaller) PoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "poolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedSession) PoolAddress() (common.Address, error) {
	return _ChainlinkFeed.Contract.PoolAddress(&_ChainlinkFeed.CallOpts)
}

// PoolAddress is a free data retrieval call binding the contract method 0x1755ff21.
//
// Solidity: function poolAddress() view returns(address)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) PoolAddress() (common.Address, error) {
	return _ChainlinkFeed.Contract.PoolAddress(&_ChainlinkFeed.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChainlinkFeed *ChainlinkFeedCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChainlinkFeed *ChainlinkFeedSession) ProxiableUUID() ([32]byte, error) {
	return _ChainlinkFeed.Contract.ProxiableUUID(&_ChainlinkFeed.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ChainlinkFeed.Contract.ProxiableUUID(&_ChainlinkFeed.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkFeed.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedSession) Version() (*big.Int, error) {
	return _ChainlinkFeed.Contract.Version(&_ChainlinkFeed.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkFeed *ChainlinkFeedCallerSession) Version() (*big.Int, error) {
	return _ChainlinkFeed.Contract.Version(&_ChainlinkFeed.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address _poolAddress, uint8 _decimals, address _oracleFeeds, address _owner) returns()
func (_ChainlinkFeed *ChainlinkFeedTransactor) Initialize(opts *bind.TransactOpts, _poolAddress common.Address, _decimals uint8, _oracleFeeds common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.contract.Transact(opts, "initialize", _poolAddress, _decimals, _oracleFeeds, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address _poolAddress, uint8 _decimals, address _oracleFeeds, address _owner) returns()
func (_ChainlinkFeed *ChainlinkFeedSession) Initialize(_poolAddress common.Address, _decimals uint8, _oracleFeeds common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.Initialize(&_ChainlinkFeed.TransactOpts, _poolAddress, _decimals, _oracleFeeds, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x9c2c58f8.
//
// Solidity: function initialize(address _poolAddress, uint8 _decimals, address _oracleFeeds, address _owner) returns()
func (_ChainlinkFeed *ChainlinkFeedTransactorSession) Initialize(_poolAddress common.Address, _decimals uint8, _oracleFeeds common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.Initialize(&_ChainlinkFeed.TransactOpts, _poolAddress, _decimals, _oracleFeeds, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkFeed *ChainlinkFeedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkFeed.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkFeed *ChainlinkFeedSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.RenounceOwnership(&_ChainlinkFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkFeed *ChainlinkFeedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.RenounceOwnership(&_ChainlinkFeed.TransactOpts)
}

// SetBalancerPoolAddress is a paid mutator transaction binding the contract method 0xd8e9b2e0.
//
// Solidity: function setBalancerPoolAddress(address _balancerPoolAddress) returns()
func (_ChainlinkFeed *ChainlinkFeedTransactor) SetBalancerPoolAddress(opts *bind.TransactOpts, _balancerPoolAddress common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.contract.Transact(opts, "setBalancerPoolAddress", _balancerPoolAddress)
}

// SetBalancerPoolAddress is a paid mutator transaction binding the contract method 0xd8e9b2e0.
//
// Solidity: function setBalancerPoolAddress(address _balancerPoolAddress) returns()
func (_ChainlinkFeed *ChainlinkFeedSession) SetBalancerPoolAddress(_balancerPoolAddress common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.SetBalancerPoolAddress(&_ChainlinkFeed.TransactOpts, _balancerPoolAddress)
}

// SetBalancerPoolAddress is a paid mutator transaction binding the contract method 0xd8e9b2e0.
//
// Solidity: function setBalancerPoolAddress(address _balancerPoolAddress) returns()
func (_ChainlinkFeed *ChainlinkFeedTransactorSession) SetBalancerPoolAddress(_balancerPoolAddress common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.SetBalancerPoolAddress(&_ChainlinkFeed.TransactOpts, _balancerPoolAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainlinkFeed *ChainlinkFeedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainlinkFeed *ChainlinkFeedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.TransferOwnership(&_ChainlinkFeed.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainlinkFeed *ChainlinkFeedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.TransferOwnership(&_ChainlinkFeed.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChainlinkFeed *ChainlinkFeedTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChainlinkFeed.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChainlinkFeed *ChainlinkFeedSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.UpgradeToAndCall(&_ChainlinkFeed.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ChainlinkFeed *ChainlinkFeedTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ChainlinkFeed.Contract.UpgradeToAndCall(&_ChainlinkFeed.TransactOpts, newImplementation, data)
}

// ChainlinkFeedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ChainlinkFeed contract.
type ChainlinkFeedInitializedIterator struct {
	Event *ChainlinkFeedInitialized // Event containing the contract specifics and raw log

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
func (it *ChainlinkFeedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkFeedInitialized)
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
		it.Event = new(ChainlinkFeedInitialized)
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
func (it *ChainlinkFeedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkFeedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkFeedInitialized represents a Initialized event raised by the ChainlinkFeed contract.
type ChainlinkFeedInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChainlinkFeed *ChainlinkFeedFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChainlinkFeedInitializedIterator, error) {

	logs, sub, err := _ChainlinkFeed.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedInitializedIterator{contract: _ChainlinkFeed.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ChainlinkFeed *ChainlinkFeedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChainlinkFeedInitialized) (event.Subscription, error) {

	logs, sub, err := _ChainlinkFeed.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkFeedInitialized)
				if err := _ChainlinkFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ChainlinkFeed *ChainlinkFeedFilterer) ParseInitialized(log types.Log) (*ChainlinkFeedInitialized, error) {
	event := new(ChainlinkFeedInitialized)
	if err := _ChainlinkFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainlinkFeedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainlinkFeed contract.
type ChainlinkFeedOwnershipTransferredIterator struct {
	Event *ChainlinkFeedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChainlinkFeedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkFeedOwnershipTransferred)
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
		it.Event = new(ChainlinkFeedOwnershipTransferred)
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
func (it *ChainlinkFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkFeedOwnershipTransferred represents a OwnershipTransferred event raised by the ChainlinkFeed contract.
type ChainlinkFeedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainlinkFeed *ChainlinkFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChainlinkFeedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainlinkFeed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedOwnershipTransferredIterator{contract: _ChainlinkFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainlinkFeed *ChainlinkFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainlinkFeedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainlinkFeed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkFeedOwnershipTransferred)
				if err := _ChainlinkFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChainlinkFeed *ChainlinkFeedFilterer) ParseOwnershipTransferred(log types.Log) (*ChainlinkFeedOwnershipTransferred, error) {
	event := new(ChainlinkFeedOwnershipTransferred)
	if err := _ChainlinkFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainlinkFeedUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ChainlinkFeed contract.
type ChainlinkFeedUpgradedIterator struct {
	Event *ChainlinkFeedUpgraded // Event containing the contract specifics and raw log

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
func (it *ChainlinkFeedUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkFeedUpgraded)
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
		it.Event = new(ChainlinkFeedUpgraded)
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
func (it *ChainlinkFeedUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkFeedUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkFeedUpgraded represents a Upgraded event raised by the ChainlinkFeed contract.
type ChainlinkFeedUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChainlinkFeed *ChainlinkFeedFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChainlinkFeedUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChainlinkFeed.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedUpgradedIterator{contract: _ChainlinkFeed.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ChainlinkFeed *ChainlinkFeedFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChainlinkFeedUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ChainlinkFeed.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkFeedUpgraded)
				if err := _ChainlinkFeed.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ChainlinkFeed *ChainlinkFeedFilterer) ParseUpgraded(log types.Log) (*ChainlinkFeedUpgraded, error) {
	event := new(ChainlinkFeedUpgraded)
	if err := _ChainlinkFeed.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

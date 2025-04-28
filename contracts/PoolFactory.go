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

// PoolFactoryPoolParams is an auto generated low-level Go binding around an user-defined struct.
type PoolFactoryPoolParams struct {
	Fee                *big.Int
	ReserveToken       common.Address
	CouponToken        common.Address
	DistributionPeriod *big.Int
	SharesPerToken     *big.Int
	FeeBeneficiary     common.Address
}

// PoolFactoryMetaData contains all meta data concerning the PoolFactory contract.
var PoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOL_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SECURITY_COUNCIL_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPool\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structPoolFactory.PoolParams\",\"components\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"couponToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"leverageAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"bondSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"leverageName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"leverageSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"pauseOnCreation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossChainController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractDeployer\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributorBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributorIntegrationAdapterBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributorIntegrationAdapters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributors\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_governance\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracleFeeds\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bondImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_leverageImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_distributorImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"leverageBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracleFeeds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pools\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCrossChainController\",\"inputs\":[{\"name\":\"_crossChainController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDeployer\",\"inputs\":[{\"name\":\"_deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDistributorIntegrationAdapterBeacon\",\"inputs\":[{\"name\":\"_distributorIntegrationAdapterBeacon\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGovernance\",\"inputs\":[{\"name\":\"_governance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolCreated\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"leverageAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorCreatingContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorCreatingProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TargetAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroDebtAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLeverageAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroReserveAmount\",\"inputs\":[]}]",
}

// PoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolFactoryMetaData.ABI instead.
var PoolFactoryABI = PoolFactoryMetaData.ABI

// PoolFactory is an auto generated Go binding around an Ethereum contract.
type PoolFactory struct {
	PoolFactoryCaller     // Read-only binding to the contract
	PoolFactoryTransactor // Write-only binding to the contract
	PoolFactoryFilterer   // Log filterer for contract events
}

// PoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolFactorySession struct {
	Contract     *PoolFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolFactoryCallerSession struct {
	Contract *PoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolFactoryTransactorSession struct {
	Contract     *PoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolFactoryRaw struct {
	Contract *PoolFactory // Generic contract binding to access the raw methods on
}

// PoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolFactoryCallerRaw struct {
	Contract *PoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolFactoryTransactorRaw struct {
	Contract *PoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolFactory creates a new instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactory(address common.Address, backend bind.ContractBackend) (*PoolFactory, error) {
	contract, err := bindPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolFactory{PoolFactoryCaller: PoolFactoryCaller{contract: contract}, PoolFactoryTransactor: PoolFactoryTransactor{contract: contract}, PoolFactoryFilterer: PoolFactoryFilterer{contract: contract}}, nil
}

// NewPoolFactoryCaller creates a new read-only instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*PoolFactoryCaller, error) {
	contract, err := bindPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryCaller{contract: contract}, nil
}

// NewPoolFactoryTransactor creates a new write-only instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolFactoryTransactor, error) {
	contract, err := bindPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryTransactor{contract: contract}, nil
}

// NewPoolFactoryFilterer creates a new log filterer instance of PoolFactory, bound to a specific deployed contract.
func NewPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFactoryFilterer, error) {
	contract, err := bindPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryFilterer{contract: contract}, nil
}

// bindPoolFactory binds a generic wrapper to an already deployed contract.
func bindPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolFactory *PoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolFactory.Contract.PoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolFactory *PoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.Contract.PoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolFactory *PoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolFactory.Contract.PoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolFactory *PoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolFactory *PoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolFactory *PoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolFactory.Contract.DEFAULTADMINROLE(&_PoolFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PoolFactory.Contract.DEFAULTADMINROLE(&_PoolFactory.CallOpts)
}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCaller) GOVROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "GOV_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactorySession) GOVROLE() ([32]byte, error) {
	return _PoolFactory.Contract.GOVROLE(&_PoolFactory.CallOpts)
}

// GOVROLE is a free data retrieval call binding the contract method 0xb536818a.
//
// Solidity: function GOV_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCallerSession) GOVROLE() ([32]byte, error) {
	return _PoolFactory.Contract.GOVROLE(&_PoolFactory.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactorySession) MINTERROLE() ([32]byte, error) {
	return _PoolFactory.Contract.MINTERROLE(&_PoolFactory.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCallerSession) MINTERROLE() ([32]byte, error) {
	return _PoolFactory.Contract.MINTERROLE(&_PoolFactory.CallOpts)
}

// POOLROLE is a free data retrieval call binding the contract method 0x404ccd07.
//
// Solidity: function POOL_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCaller) POOLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "POOL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POOLROLE is a free data retrieval call binding the contract method 0x404ccd07.
//
// Solidity: function POOL_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactorySession) POOLROLE() ([32]byte, error) {
	return _PoolFactory.Contract.POOLROLE(&_PoolFactory.CallOpts)
}

// POOLROLE is a free data retrieval call binding the contract method 0x404ccd07.
//
// Solidity: function POOL_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCallerSession) POOLROLE() ([32]byte, error) {
	return _PoolFactory.Contract.POOLROLE(&_PoolFactory.CallOpts)
}

// SECURITYCOUNCILROLE is a free data retrieval call binding the contract method 0xb9fe5cf7.
//
// Solidity: function SECURITY_COUNCIL_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCaller) SECURITYCOUNCILROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "SECURITY_COUNCIL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SECURITYCOUNCILROLE is a free data retrieval call binding the contract method 0xb9fe5cf7.
//
// Solidity: function SECURITY_COUNCIL_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactorySession) SECURITYCOUNCILROLE() ([32]byte, error) {
	return _PoolFactory.Contract.SECURITYCOUNCILROLE(&_PoolFactory.CallOpts)
}

// SECURITYCOUNCILROLE is a free data retrieval call binding the contract method 0xb9fe5cf7.
//
// Solidity: function SECURITY_COUNCIL_ROLE() view returns(bytes32)
func (_PoolFactory *PoolFactoryCallerSession) SECURITYCOUNCILROLE() ([32]byte, error) {
	return _PoolFactory.Contract.SECURITYCOUNCILROLE(&_PoolFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PoolFactory *PoolFactoryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PoolFactory *PoolFactorySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PoolFactory.Contract.UPGRADEINTERFACEVERSION(&_PoolFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PoolFactory *PoolFactoryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PoolFactory.Contract.UPGRADEINTERFACEVERSION(&_PoolFactory.CallOpts)
}

// BondBeacon is a free data retrieval call binding the contract method 0xc8f12bde.
//
// Solidity: function bondBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCaller) BondBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "bondBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondBeacon is a free data retrieval call binding the contract method 0xc8f12bde.
//
// Solidity: function bondBeacon() view returns(address)
func (_PoolFactory *PoolFactorySession) BondBeacon() (common.Address, error) {
	return _PoolFactory.Contract.BondBeacon(&_PoolFactory.CallOpts)
}

// BondBeacon is a free data retrieval call binding the contract method 0xc8f12bde.
//
// Solidity: function bondBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) BondBeacon() (common.Address, error) {
	return _PoolFactory.Contract.BondBeacon(&_PoolFactory.CallOpts)
}

// CrossChainController is a free data retrieval call binding the contract method 0xcbd5920f.
//
// Solidity: function crossChainController() view returns(address)
func (_PoolFactory *PoolFactoryCaller) CrossChainController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "crossChainController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrossChainController is a free data retrieval call binding the contract method 0xcbd5920f.
//
// Solidity: function crossChainController() view returns(address)
func (_PoolFactory *PoolFactorySession) CrossChainController() (common.Address, error) {
	return _PoolFactory.Contract.CrossChainController(&_PoolFactory.CallOpts)
}

// CrossChainController is a free data retrieval call binding the contract method 0xcbd5920f.
//
// Solidity: function crossChainController() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) CrossChainController() (common.Address, error) {
	return _PoolFactory.Contract.CrossChainController(&_PoolFactory.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PoolFactory *PoolFactoryCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PoolFactory *PoolFactorySession) Deployer() (common.Address, error) {
	return _PoolFactory.Contract.Deployer(&_PoolFactory.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Deployer() (common.Address, error) {
	return _PoolFactory.Contract.Deployer(&_PoolFactory.CallOpts)
}

// DistributorBeacon is a free data retrieval call binding the contract method 0x22a06f46.
//
// Solidity: function distributorBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCaller) DistributorBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "distributorBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DistributorBeacon is a free data retrieval call binding the contract method 0x22a06f46.
//
// Solidity: function distributorBeacon() view returns(address)
func (_PoolFactory *PoolFactorySession) DistributorBeacon() (common.Address, error) {
	return _PoolFactory.Contract.DistributorBeacon(&_PoolFactory.CallOpts)
}

// DistributorBeacon is a free data retrieval call binding the contract method 0x22a06f46.
//
// Solidity: function distributorBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) DistributorBeacon() (common.Address, error) {
	return _PoolFactory.Contract.DistributorBeacon(&_PoolFactory.CallOpts)
}

// DistributorIntegrationAdapterBeacon is a free data retrieval call binding the contract method 0x05d79327.
//
// Solidity: function distributorIntegrationAdapterBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCaller) DistributorIntegrationAdapterBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "distributorIntegrationAdapterBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DistributorIntegrationAdapterBeacon is a free data retrieval call binding the contract method 0x05d79327.
//
// Solidity: function distributorIntegrationAdapterBeacon() view returns(address)
func (_PoolFactory *PoolFactorySession) DistributorIntegrationAdapterBeacon() (common.Address, error) {
	return _PoolFactory.Contract.DistributorIntegrationAdapterBeacon(&_PoolFactory.CallOpts)
}

// DistributorIntegrationAdapterBeacon is a free data retrieval call binding the contract method 0x05d79327.
//
// Solidity: function distributorIntegrationAdapterBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) DistributorIntegrationAdapterBeacon() (common.Address, error) {
	return _PoolFactory.Contract.DistributorIntegrationAdapterBeacon(&_PoolFactory.CallOpts)
}

// DistributorIntegrationAdapters is a free data retrieval call binding the contract method 0x74db3b82.
//
// Solidity: function distributorIntegrationAdapters(address ) view returns(address)
func (_PoolFactory *PoolFactoryCaller) DistributorIntegrationAdapters(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "distributorIntegrationAdapters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DistributorIntegrationAdapters is a free data retrieval call binding the contract method 0x74db3b82.
//
// Solidity: function distributorIntegrationAdapters(address ) view returns(address)
func (_PoolFactory *PoolFactorySession) DistributorIntegrationAdapters(arg0 common.Address) (common.Address, error) {
	return _PoolFactory.Contract.DistributorIntegrationAdapters(&_PoolFactory.CallOpts, arg0)
}

// DistributorIntegrationAdapters is a free data retrieval call binding the contract method 0x74db3b82.
//
// Solidity: function distributorIntegrationAdapters(address ) view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) DistributorIntegrationAdapters(arg0 common.Address) (common.Address, error) {
	return _PoolFactory.Contract.DistributorIntegrationAdapters(&_PoolFactory.CallOpts, arg0)
}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors(address ) view returns(address)
func (_PoolFactory *PoolFactoryCaller) Distributors(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "distributors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors(address ) view returns(address)
func (_PoolFactory *PoolFactorySession) Distributors(arg0 common.Address) (common.Address, error) {
	return _PoolFactory.Contract.Distributors(&_PoolFactory.CallOpts, arg0)
}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors(address ) view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Distributors(arg0 common.Address) (common.Address, error) {
	return _PoolFactory.Contract.Distributors(&_PoolFactory.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolFactory *PoolFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolFactory *PoolFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolFactory.Contract.GetRoleAdmin(&_PoolFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PoolFactory *PoolFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PoolFactory.Contract.GetRoleAdmin(&_PoolFactory.CallOpts, role)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PoolFactory *PoolFactoryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PoolFactory *PoolFactorySession) Governance() (common.Address, error) {
	return _PoolFactory.Contract.Governance(&_PoolFactory.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Governance() (common.Address, error) {
	return _PoolFactory.Contract.Governance(&_PoolFactory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolFactory *PoolFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolFactory *PoolFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolFactory.Contract.HasRole(&_PoolFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PoolFactory *PoolFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PoolFactory.Contract.HasRole(&_PoolFactory.CallOpts, role, account)
}

// LeverageBeacon is a free data retrieval call binding the contract method 0x2808df1f.
//
// Solidity: function leverageBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCaller) LeverageBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "leverageBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LeverageBeacon is a free data retrieval call binding the contract method 0x2808df1f.
//
// Solidity: function leverageBeacon() view returns(address)
func (_PoolFactory *PoolFactorySession) LeverageBeacon() (common.Address, error) {
	return _PoolFactory.Contract.LeverageBeacon(&_PoolFactory.CallOpts)
}

// LeverageBeacon is a free data retrieval call binding the contract method 0x2808df1f.
//
// Solidity: function leverageBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) LeverageBeacon() (common.Address, error) {
	return _PoolFactory.Contract.LeverageBeacon(&_PoolFactory.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_PoolFactory *PoolFactoryCaller) OracleFeeds(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "oracleFeeds")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_PoolFactory *PoolFactorySession) OracleFeeds() (common.Address, error) {
	return _PoolFactory.Contract.OracleFeeds(&_PoolFactory.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) OracleFeeds() (common.Address, error) {
	return _PoolFactory.Contract.OracleFeeds(&_PoolFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolFactory *PoolFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolFactory *PoolFactorySession) Paused() (bool, error) {
	return _PoolFactory.Contract.Paused(&_PoolFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolFactory *PoolFactoryCallerSession) Paused() (bool, error) {
	return _PoolFactory.Contract.Paused(&_PoolFactory.CallOpts)
}

// PoolBeacon is a free data retrieval call binding the contract method 0xc5a7b2ea.
//
// Solidity: function poolBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCaller) PoolBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "poolBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolBeacon is a free data retrieval call binding the contract method 0xc5a7b2ea.
//
// Solidity: function poolBeacon() view returns(address)
func (_PoolFactory *PoolFactorySession) PoolBeacon() (common.Address, error) {
	return _PoolFactory.Contract.PoolBeacon(&_PoolFactory.CallOpts)
}

// PoolBeacon is a free data retrieval call binding the contract method 0xc5a7b2ea.
//
// Solidity: function poolBeacon() view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) PoolBeacon() (common.Address, error) {
	return _PoolFactory.Contract.PoolBeacon(&_PoolFactory.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_PoolFactory *PoolFactoryCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_PoolFactory *PoolFactorySession) Pools(arg0 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.Pools(&_PoolFactory.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_PoolFactory *PoolFactoryCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _PoolFactory.Contract.Pools(&_PoolFactory.CallOpts, arg0)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_PoolFactory *PoolFactoryCaller) PoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "poolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_PoolFactory *PoolFactorySession) PoolsLength() (*big.Int, error) {
	return _PoolFactory.Contract.PoolsLength(&_PoolFactory.CallOpts)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_PoolFactory *PoolFactoryCallerSession) PoolsLength() (*big.Int, error) {
	return _PoolFactory.Contract.PoolsLength(&_PoolFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PoolFactory *PoolFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PoolFactory *PoolFactorySession) ProxiableUUID() ([32]byte, error) {
	return _PoolFactory.Contract.ProxiableUUID(&_PoolFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PoolFactory *PoolFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PoolFactory.Contract.ProxiableUUID(&_PoolFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolFactory *PoolFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PoolFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolFactory *PoolFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolFactory.Contract.SupportsInterface(&_PoolFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PoolFactory *PoolFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PoolFactory.Contract.SupportsInterface(&_PoolFactory.CallOpts, interfaceId)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa3c31dd5.
//
// Solidity: function createPool((uint256,address,address,uint256,uint256,address) params, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount, string bondName, string bondSymbol, string leverageName, string leverageSymbol, bool pauseOnCreation) returns(address)
func (_PoolFactory *PoolFactoryTransactor) CreatePool(opts *bind.TransactOpts, params PoolFactoryPoolParams, reserveAmount *big.Int, bondAmount *big.Int, leverageAmount *big.Int, bondName string, bondSymbol string, leverageName string, leverageSymbol string, pauseOnCreation bool) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "createPool", params, reserveAmount, bondAmount, leverageAmount, bondName, bondSymbol, leverageName, leverageSymbol, pauseOnCreation)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa3c31dd5.
//
// Solidity: function createPool((uint256,address,address,uint256,uint256,address) params, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount, string bondName, string bondSymbol, string leverageName, string leverageSymbol, bool pauseOnCreation) returns(address)
func (_PoolFactory *PoolFactorySession) CreatePool(params PoolFactoryPoolParams, reserveAmount *big.Int, bondAmount *big.Int, leverageAmount *big.Int, bondName string, bondSymbol string, leverageName string, leverageSymbol string, pauseOnCreation bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool(&_PoolFactory.TransactOpts, params, reserveAmount, bondAmount, leverageAmount, bondName, bondSymbol, leverageName, leverageSymbol, pauseOnCreation)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa3c31dd5.
//
// Solidity: function createPool((uint256,address,address,uint256,uint256,address) params, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount, string bondName, string bondSymbol, string leverageName, string leverageSymbol, bool pauseOnCreation) returns(address)
func (_PoolFactory *PoolFactoryTransactorSession) CreatePool(params PoolFactoryPoolParams, reserveAmount *big.Int, bondAmount *big.Int, leverageAmount *big.Int, bondName string, bondSymbol string, leverageName string, leverageSymbol string, pauseOnCreation bool) (*types.Transaction, error) {
	return _PoolFactory.Contract.CreatePool(&_PoolFactory.TransactOpts, params, reserveAmount, bondAmount, leverageAmount, bondName, bondSymbol, leverageName, leverageSymbol, pauseOnCreation)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolFactory *PoolFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolFactory *PoolFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.GrantRole(&_PoolFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PoolFactory *PoolFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.GrantRole(&_PoolFactory.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _governance, address _deployer, address _oracleFeeds, address _poolImplementation, address _bondImplementation, address _leverageImplementation, address _distributorImplementation) returns()
func (_PoolFactory *PoolFactoryTransactor) Initialize(opts *bind.TransactOpts, _governance common.Address, _deployer common.Address, _oracleFeeds common.Address, _poolImplementation common.Address, _bondImplementation common.Address, _leverageImplementation common.Address, _distributorImplementation common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "initialize", _governance, _deployer, _oracleFeeds, _poolImplementation, _bondImplementation, _leverageImplementation, _distributorImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _governance, address _deployer, address _oracleFeeds, address _poolImplementation, address _bondImplementation, address _leverageImplementation, address _distributorImplementation) returns()
func (_PoolFactory *PoolFactorySession) Initialize(_governance common.Address, _deployer common.Address, _oracleFeeds common.Address, _poolImplementation common.Address, _bondImplementation common.Address, _leverageImplementation common.Address, _distributorImplementation common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.Initialize(&_PoolFactory.TransactOpts, _governance, _deployer, _oracleFeeds, _poolImplementation, _bondImplementation, _leverageImplementation, _distributorImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _governance, address _deployer, address _oracleFeeds, address _poolImplementation, address _bondImplementation, address _leverageImplementation, address _distributorImplementation) returns()
func (_PoolFactory *PoolFactoryTransactorSession) Initialize(_governance common.Address, _deployer common.Address, _oracleFeeds common.Address, _poolImplementation common.Address, _bondImplementation common.Address, _leverageImplementation common.Address, _distributorImplementation common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.Initialize(&_PoolFactory.TransactOpts, _governance, _deployer, _oracleFeeds, _poolImplementation, _bondImplementation, _leverageImplementation, _distributorImplementation)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolFactory *PoolFactoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolFactory *PoolFactorySession) Pause() (*types.Transaction, error) {
	return _PoolFactory.Contract.Pause(&_PoolFactory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolFactory *PoolFactoryTransactorSession) Pause() (*types.Transaction, error) {
	return _PoolFactory.Contract.Pause(&_PoolFactory.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PoolFactory *PoolFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PoolFactory *PoolFactorySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.RenounceRole(&_PoolFactory.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_PoolFactory *PoolFactoryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.RenounceRole(&_PoolFactory.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolFactory *PoolFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolFactory *PoolFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.RevokeRole(&_PoolFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PoolFactory *PoolFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.RevokeRole(&_PoolFactory.TransactOpts, role, account)
}

// SetCrossChainController is a paid mutator transaction binding the contract method 0x70880b32.
//
// Solidity: function setCrossChainController(address _crossChainController) returns()
func (_PoolFactory *PoolFactoryTransactor) SetCrossChainController(opts *bind.TransactOpts, _crossChainController common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setCrossChainController", _crossChainController)
}

// SetCrossChainController is a paid mutator transaction binding the contract method 0x70880b32.
//
// Solidity: function setCrossChainController(address _crossChainController) returns()
func (_PoolFactory *PoolFactorySession) SetCrossChainController(_crossChainController common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetCrossChainController(&_PoolFactory.TransactOpts, _crossChainController)
}

// SetCrossChainController is a paid mutator transaction binding the contract method 0x70880b32.
//
// Solidity: function setCrossChainController(address _crossChainController) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetCrossChainController(_crossChainController common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetCrossChainController(&_PoolFactory.TransactOpts, _crossChainController)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _deployer) returns()
func (_PoolFactory *PoolFactoryTransactor) SetDeployer(opts *bind.TransactOpts, _deployer common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setDeployer", _deployer)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _deployer) returns()
func (_PoolFactory *PoolFactorySession) SetDeployer(_deployer common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetDeployer(&_PoolFactory.TransactOpts, _deployer)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _deployer) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetDeployer(_deployer common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetDeployer(&_PoolFactory.TransactOpts, _deployer)
}

// SetDistributorIntegrationAdapterBeacon is a paid mutator transaction binding the contract method 0xc3f7eab7.
//
// Solidity: function setDistributorIntegrationAdapterBeacon(address _distributorIntegrationAdapterBeacon) returns()
func (_PoolFactory *PoolFactoryTransactor) SetDistributorIntegrationAdapterBeacon(opts *bind.TransactOpts, _distributorIntegrationAdapterBeacon common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setDistributorIntegrationAdapterBeacon", _distributorIntegrationAdapterBeacon)
}

// SetDistributorIntegrationAdapterBeacon is a paid mutator transaction binding the contract method 0xc3f7eab7.
//
// Solidity: function setDistributorIntegrationAdapterBeacon(address _distributorIntegrationAdapterBeacon) returns()
func (_PoolFactory *PoolFactorySession) SetDistributorIntegrationAdapterBeacon(_distributorIntegrationAdapterBeacon common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetDistributorIntegrationAdapterBeacon(&_PoolFactory.TransactOpts, _distributorIntegrationAdapterBeacon)
}

// SetDistributorIntegrationAdapterBeacon is a paid mutator transaction binding the contract method 0xc3f7eab7.
//
// Solidity: function setDistributorIntegrationAdapterBeacon(address _distributorIntegrationAdapterBeacon) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetDistributorIntegrationAdapterBeacon(_distributorIntegrationAdapterBeacon common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetDistributorIntegrationAdapterBeacon(&_PoolFactory.TransactOpts, _distributorIntegrationAdapterBeacon)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PoolFactory *PoolFactoryTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PoolFactory *PoolFactorySession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetGovernance(&_PoolFactory.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PoolFactory *PoolFactoryTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PoolFactory.Contract.SetGovernance(&_PoolFactory.TransactOpts, _governance)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolFactory *PoolFactoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolFactory *PoolFactorySession) Unpause() (*types.Transaction, error) {
	return _PoolFactory.Contract.Unpause(&_PoolFactory.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolFactory *PoolFactoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _PoolFactory.Contract.Unpause(&_PoolFactory.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PoolFactory *PoolFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PoolFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PoolFactory *PoolFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PoolFactory.Contract.UpgradeToAndCall(&_PoolFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PoolFactory *PoolFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PoolFactory.Contract.UpgradeToAndCall(&_PoolFactory.TransactOpts, newImplementation, data)
}

// PoolFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoolFactory contract.
type PoolFactoryInitializedIterator struct {
	Event *PoolFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *PoolFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryInitialized)
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
		it.Event = new(PoolFactoryInitialized)
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
func (it *PoolFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryInitialized represents a Initialized event raised by the PoolFactory contract.
type PoolFactoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoolFactory *PoolFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolFactoryInitializedIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolFactoryInitializedIterator{contract: _PoolFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PoolFactory *PoolFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryInitialized)
				if err := _PoolFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PoolFactory *PoolFactoryFilterer) ParseInitialized(log types.Log) (*PoolFactoryInitialized, error) {
	event := new(PoolFactoryInitialized)
	if err := _PoolFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PoolFactory contract.
type PoolFactoryPausedIterator struct {
	Event *PoolFactoryPaused // Event containing the contract specifics and raw log

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
func (it *PoolFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryPaused)
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
		it.Event = new(PoolFactoryPaused)
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
func (it *PoolFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryPaused represents a Paused event raised by the PoolFactory contract.
type PoolFactoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolFactory *PoolFactoryFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolFactoryPausedIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolFactoryPausedIterator{contract: _PoolFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolFactory *PoolFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolFactoryPaused) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryPaused)
				if err := _PoolFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PoolFactory *PoolFactoryFilterer) ParsePaused(log types.Log) (*PoolFactoryPaused, error) {
	event := new(PoolFactoryPaused)
	if err := _PoolFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the PoolFactory contract.
type PoolFactoryPoolCreatedIterator struct {
	Event *PoolFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *PoolFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryPoolCreated)
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
		it.Event = new(PoolFactoryPoolCreated)
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
func (it *PoolFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryPoolCreated represents a PoolCreated event raised by the PoolFactory contract.
type PoolFactoryPoolCreated struct {
	Pool           common.Address
	ReserveAmount  *big.Int
	BondAmount     *big.Int
	LeverageAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xc8e774967d7f45e5814dbb800335ed47014999568397f08e79b0f322047f743f.
//
// Solidity: event PoolCreated(address pool, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount)
func (_PoolFactory *PoolFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts) (*PoolFactoryPoolCreatedIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return &PoolFactoryPoolCreatedIterator{contract: _PoolFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xc8e774967d7f45e5814dbb800335ed47014999568397f08e79b0f322047f743f.
//
// Solidity: event PoolCreated(address pool, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount)
func (_PoolFactory *PoolFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PoolFactoryPoolCreated) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryPoolCreated)
				if err := _PoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0xc8e774967d7f45e5814dbb800335ed47014999568397f08e79b0f322047f743f.
//
// Solidity: event PoolCreated(address pool, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount)
func (_PoolFactory *PoolFactoryFilterer) ParsePoolCreated(log types.Log) (*PoolFactoryPoolCreated, error) {
	event := new(PoolFactoryPoolCreated)
	if err := _PoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PoolFactory contract.
type PoolFactoryRoleAdminChangedIterator struct {
	Event *PoolFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PoolFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryRoleAdminChanged)
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
		it.Event = new(PoolFactoryRoleAdminChanged)
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
func (it *PoolFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the PoolFactory contract.
type PoolFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PoolFactory *PoolFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PoolFactoryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryRoleAdminChangedIterator{contract: _PoolFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PoolFactory *PoolFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PoolFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryRoleAdminChanged)
				if err := _PoolFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PoolFactory *PoolFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*PoolFactoryRoleAdminChanged, error) {
	event := new(PoolFactoryRoleAdminChanged)
	if err := _PoolFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PoolFactory contract.
type PoolFactoryRoleGrantedIterator struct {
	Event *PoolFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *PoolFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryRoleGranted)
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
		it.Event = new(PoolFactoryRoleGranted)
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
func (it *PoolFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryRoleGranted represents a RoleGranted event raised by the PoolFactory contract.
type PoolFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolFactory *PoolFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolFactoryRoleGrantedIterator, error) {

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

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryRoleGrantedIterator{contract: _PoolFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolFactory *PoolFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PoolFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryRoleGranted)
				if err := _PoolFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PoolFactory *PoolFactoryFilterer) ParseRoleGranted(log types.Log) (*PoolFactoryRoleGranted, error) {
	event := new(PoolFactoryRoleGranted)
	if err := _PoolFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PoolFactory contract.
type PoolFactoryRoleRevokedIterator struct {
	Event *PoolFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PoolFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryRoleRevoked)
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
		it.Event = new(PoolFactoryRoleRevoked)
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
func (it *PoolFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryRoleRevoked represents a RoleRevoked event raised by the PoolFactory contract.
type PoolFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolFactory *PoolFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PoolFactoryRoleRevokedIterator, error) {

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

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryRoleRevokedIterator{contract: _PoolFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PoolFactory *PoolFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PoolFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryRoleRevoked)
				if err := _PoolFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PoolFactory *PoolFactoryFilterer) ParseRoleRevoked(log types.Log) (*PoolFactoryRoleRevoked, error) {
	event := new(PoolFactoryRoleRevoked)
	if err := _PoolFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PoolFactory contract.
type PoolFactoryUnpausedIterator struct {
	Event *PoolFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryUnpaused)
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
		it.Event = new(PoolFactoryUnpaused)
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
func (it *PoolFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryUnpaused represents a Unpaused event raised by the PoolFactory contract.
type PoolFactoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolFactory *PoolFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolFactoryUnpausedIterator, error) {

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolFactoryUnpausedIterator{contract: _PoolFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolFactory *PoolFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolFactoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryUnpaused)
				if err := _PoolFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PoolFactory *PoolFactoryFilterer) ParseUnpaused(log types.Log) (*PoolFactoryUnpaused, error) {
	event := new(PoolFactoryUnpaused)
	if err := _PoolFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PoolFactory contract.
type PoolFactoryUpgradedIterator struct {
	Event *PoolFactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *PoolFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFactoryUpgraded)
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
		it.Event = new(PoolFactoryUpgraded)
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
func (it *PoolFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFactoryUpgraded represents a Upgraded event raised by the PoolFactory contract.
type PoolFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PoolFactory *PoolFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PoolFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PoolFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PoolFactoryUpgradedIterator{contract: _PoolFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PoolFactory *PoolFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PoolFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PoolFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFactoryUpgraded)
				if err := _PoolFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PoolFactory *PoolFactoryFilterer) ParseUpgraded(log types.Log) (*PoolFactoryUpgraded, error) {
	event := new(PoolFactoryUpgraded)
	if err := _PoolFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

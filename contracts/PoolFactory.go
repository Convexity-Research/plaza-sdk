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

// PoolFactoryPoolParams is an auto generated low-level Go binding around an user-defined struct.
type PoolFactoryPoolParams struct {
	Fee                *big.Int
	ReserveToken       common.Address
	CouponToken        common.Address
	DistributionPeriod *big.Int
	SharesPerToken     *big.Int
	FeeBeneficiary     common.Address
}

// PlazaSdkMetaData contains all meta data concerning the PlazaSdk contract.
var PlazaSdkMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOV_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POOL_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SECURITY_COUNCIL_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPool\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structPoolFactory.PoolParams\",\"components\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"couponToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"leverageAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"bondSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"leverageName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"leverageSymbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"pauseOnCreation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractDeployer\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributorBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributors\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"governance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_governance\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracleFeeds\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bondImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_leverageImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_distributorImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"leverageBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracleFeeds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pools\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDeployer\",\"inputs\":[{\"name\":\"_deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGovernance\",\"inputs\":[{\"name\":\"_governance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolCreated\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"leverageAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorCreatingContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrorCreatingProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TargetAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroDebtAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLeverageAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroReserveAmount\",\"inputs\":[]}]",
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

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkSession) MINTERROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.MINTERROLE(&_PlazaSdk.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCallerSession) MINTERROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.MINTERROLE(&_PlazaSdk.CallOpts)
}

// POOLROLE is a free data retrieval call binding the contract method 0x404ccd07.
//
// Solidity: function POOL_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCaller) POOLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "POOL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POOLROLE is a free data retrieval call binding the contract method 0x404ccd07.
//
// Solidity: function POOL_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkSession) POOLROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.POOLROLE(&_PlazaSdk.CallOpts)
}

// POOLROLE is a free data retrieval call binding the contract method 0x404ccd07.
//
// Solidity: function POOL_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCallerSession) POOLROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.POOLROLE(&_PlazaSdk.CallOpts)
}

// SECURITYCOUNCILROLE is a free data retrieval call binding the contract method 0xb9fe5cf7.
//
// Solidity: function SECURITY_COUNCIL_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCaller) SECURITYCOUNCILROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "SECURITY_COUNCIL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SECURITYCOUNCILROLE is a free data retrieval call binding the contract method 0xb9fe5cf7.
//
// Solidity: function SECURITY_COUNCIL_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkSession) SECURITYCOUNCILROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.SECURITYCOUNCILROLE(&_PlazaSdk.CallOpts)
}

// SECURITYCOUNCILROLE is a free data retrieval call binding the contract method 0xb9fe5cf7.
//
// Solidity: function SECURITY_COUNCIL_ROLE() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCallerSession) SECURITYCOUNCILROLE() ([32]byte, error) {
	return _PlazaSdk.Contract.SECURITYCOUNCILROLE(&_PlazaSdk.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PlazaSdk *PlazaSdkCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PlazaSdk *PlazaSdkSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PlazaSdk.Contract.UPGRADEINTERFACEVERSION(&_PlazaSdk.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_PlazaSdk *PlazaSdkCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _PlazaSdk.Contract.UPGRADEINTERFACEVERSION(&_PlazaSdk.CallOpts)
}

// BondBeacon is a free data retrieval call binding the contract method 0xc8f12bde.
//
// Solidity: function bondBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) BondBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "bondBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondBeacon is a free data retrieval call binding the contract method 0xc8f12bde.
//
// Solidity: function bondBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkSession) BondBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.BondBeacon(&_PlazaSdk.CallOpts)
}

// BondBeacon is a free data retrieval call binding the contract method 0xc8f12bde.
//
// Solidity: function bondBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) BondBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.BondBeacon(&_PlazaSdk.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PlazaSdk *PlazaSdkSession) Deployer() (common.Address, error) {
	return _PlazaSdk.Contract.Deployer(&_PlazaSdk.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) Deployer() (common.Address, error) {
	return _PlazaSdk.Contract.Deployer(&_PlazaSdk.CallOpts)
}

// DistributorBeacon is a free data retrieval call binding the contract method 0x22a06f46.
//
// Solidity: function distributorBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) DistributorBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "distributorBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DistributorBeacon is a free data retrieval call binding the contract method 0x22a06f46.
//
// Solidity: function distributorBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkSession) DistributorBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.DistributorBeacon(&_PlazaSdk.CallOpts)
}

// DistributorBeacon is a free data retrieval call binding the contract method 0x22a06f46.
//
// Solidity: function distributorBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) DistributorBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.DistributorBeacon(&_PlazaSdk.CallOpts)
}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors(address ) view returns(address)
func (_PlazaSdk *PlazaSdkCaller) Distributors(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "distributors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors(address ) view returns(address)
func (_PlazaSdk *PlazaSdkSession) Distributors(arg0 common.Address) (common.Address, error) {
	return _PlazaSdk.Contract.Distributors(&_PlazaSdk.CallOpts, arg0)
}

// Distributors is a free data retrieval call binding the contract method 0xcc642784.
//
// Solidity: function distributors(address ) view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) Distributors(arg0 common.Address) (common.Address, error) {
	return _PlazaSdk.Contract.Distributors(&_PlazaSdk.CallOpts, arg0)
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

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PlazaSdk *PlazaSdkSession) Governance() (common.Address, error) {
	return _PlazaSdk.Contract.Governance(&_PlazaSdk.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) Governance() (common.Address, error) {
	return _PlazaSdk.Contract.Governance(&_PlazaSdk.CallOpts)
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

// LeverageBeacon is a free data retrieval call binding the contract method 0x2808df1f.
//
// Solidity: function leverageBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) LeverageBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "leverageBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LeverageBeacon is a free data retrieval call binding the contract method 0x2808df1f.
//
// Solidity: function leverageBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkSession) LeverageBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.LeverageBeacon(&_PlazaSdk.CallOpts)
}

// LeverageBeacon is a free data retrieval call binding the contract method 0x2808df1f.
//
// Solidity: function leverageBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) LeverageBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.LeverageBeacon(&_PlazaSdk.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) OracleFeeds(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "oracleFeeds")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_PlazaSdk *PlazaSdkSession) OracleFeeds() (common.Address, error) {
	return _PlazaSdk.Contract.OracleFeeds(&_PlazaSdk.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) OracleFeeds() (common.Address, error) {
	return _PlazaSdk.Contract.OracleFeeds(&_PlazaSdk.CallOpts)
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

// PoolBeacon is a free data retrieval call binding the contract method 0xc5a7b2ea.
//
// Solidity: function poolBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) PoolBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "poolBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolBeacon is a free data retrieval call binding the contract method 0xc5a7b2ea.
//
// Solidity: function poolBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkSession) PoolBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.PoolBeacon(&_PlazaSdk.CallOpts)
}

// PoolBeacon is a free data retrieval call binding the contract method 0xc5a7b2ea.
//
// Solidity: function poolBeacon() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) PoolBeacon() (common.Address, error) {
	return _PlazaSdk.Contract.PoolBeacon(&_PlazaSdk.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_PlazaSdk *PlazaSdkCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "pools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_PlazaSdk *PlazaSdkSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _PlazaSdk.Contract.Pools(&_PlazaSdk.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) Pools(arg0 *big.Int) (common.Address, error) {
	return _PlazaSdk.Contract.Pools(&_PlazaSdk.CallOpts, arg0)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) PoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "poolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) PoolsLength() (*big.Int, error) {
	return _PlazaSdk.Contract.PoolsLength(&_PlazaSdk.CallOpts)
}

// PoolsLength is a free data retrieval call binding the contract method 0x2716ae66.
//
// Solidity: function poolsLength() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) PoolsLength() (*big.Int, error) {
	return _PlazaSdk.Contract.PoolsLength(&_PlazaSdk.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PlazaSdk *PlazaSdkSession) ProxiableUUID() ([32]byte, error) {
	return _PlazaSdk.Contract.ProxiableUUID(&_PlazaSdk.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PlazaSdk *PlazaSdkCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PlazaSdk.Contract.ProxiableUUID(&_PlazaSdk.CallOpts)
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

// CreatePool is a paid mutator transaction binding the contract method 0xa3c31dd5.
//
// Solidity: function createPool((uint256,address,address,uint256,uint256,address) params, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount, string bondName, string bondSymbol, string leverageName, string leverageSymbol, bool pauseOnCreation) returns(address)
func (_PlazaSdk *PlazaSdkTransactor) CreatePool(opts *bind.TransactOpts, params PoolFactoryPoolParams, reserveAmount *big.Int, bondAmount *big.Int, leverageAmount *big.Int, bondName string, bondSymbol string, leverageName string, leverageSymbol string, pauseOnCreation bool) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "createPool", params, reserveAmount, bondAmount, leverageAmount, bondName, bondSymbol, leverageName, leverageSymbol, pauseOnCreation)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa3c31dd5.
//
// Solidity: function createPool((uint256,address,address,uint256,uint256,address) params, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount, string bondName, string bondSymbol, string leverageName, string leverageSymbol, bool pauseOnCreation) returns(address)
func (_PlazaSdk *PlazaSdkSession) CreatePool(params PoolFactoryPoolParams, reserveAmount *big.Int, bondAmount *big.Int, leverageAmount *big.Int, bondName string, bondSymbol string, leverageName string, leverageSymbol string, pauseOnCreation bool) (*types.Transaction, error) {
	return _PlazaSdk.Contract.CreatePool(&_PlazaSdk.TransactOpts, params, reserveAmount, bondAmount, leverageAmount, bondName, bondSymbol, leverageName, leverageSymbol, pauseOnCreation)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa3c31dd5.
//
// Solidity: function createPool((uint256,address,address,uint256,uint256,address) params, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount, string bondName, string bondSymbol, string leverageName, string leverageSymbol, bool pauseOnCreation) returns(address)
func (_PlazaSdk *PlazaSdkTransactorSession) CreatePool(params PoolFactoryPoolParams, reserveAmount *big.Int, bondAmount *big.Int, leverageAmount *big.Int, bondName string, bondSymbol string, leverageName string, leverageSymbol string, pauseOnCreation bool) (*types.Transaction, error) {
	return _PlazaSdk.Contract.CreatePool(&_PlazaSdk.TransactOpts, params, reserveAmount, bondAmount, leverageAmount, bondName, bondSymbol, leverageName, leverageSymbol, pauseOnCreation)
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

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _governance, address _deployer, address _oracleFeeds, address _poolImplementation, address _bondImplementation, address _leverageImplementation, address _distributorImplementation) returns()
func (_PlazaSdk *PlazaSdkTransactor) Initialize(opts *bind.TransactOpts, _governance common.Address, _deployer common.Address, _oracleFeeds common.Address, _poolImplementation common.Address, _bondImplementation common.Address, _leverageImplementation common.Address, _distributorImplementation common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "initialize", _governance, _deployer, _oracleFeeds, _poolImplementation, _bondImplementation, _leverageImplementation, _distributorImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _governance, address _deployer, address _oracleFeeds, address _poolImplementation, address _bondImplementation, address _leverageImplementation, address _distributorImplementation) returns()
func (_PlazaSdk *PlazaSdkSession) Initialize(_governance common.Address, _deployer common.Address, _oracleFeeds common.Address, _poolImplementation common.Address, _bondImplementation common.Address, _leverageImplementation common.Address, _distributorImplementation common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _governance, _deployer, _oracleFeeds, _poolImplementation, _bondImplementation, _leverageImplementation, _distributorImplementation)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _governance, address _deployer, address _oracleFeeds, address _poolImplementation, address _bondImplementation, address _leverageImplementation, address _distributorImplementation) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Initialize(_governance common.Address, _deployer common.Address, _oracleFeeds common.Address, _poolImplementation common.Address, _bondImplementation common.Address, _leverageImplementation common.Address, _distributorImplementation common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _governance, _deployer, _oracleFeeds, _poolImplementation, _bondImplementation, _leverageImplementation, _distributorImplementation)
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

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _deployer) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetDeployer(opts *bind.TransactOpts, _deployer common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setDeployer", _deployer)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _deployer) returns()
func (_PlazaSdk *PlazaSdkSession) SetDeployer(_deployer common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetDeployer(&_PlazaSdk.TransactOpts, _deployer)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _deployer) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetDeployer(_deployer common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetDeployer(&_PlazaSdk.TransactOpts, _deployer)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PlazaSdk *PlazaSdkSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetGovernance(&_PlazaSdk.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetGovernance(&_PlazaSdk.TransactOpts, _governance)
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

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PlazaSdk *PlazaSdkTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PlazaSdk *PlazaSdkSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PlazaSdk.Contract.UpgradeToAndCall(&_PlazaSdk.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PlazaSdk *PlazaSdkTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PlazaSdk.Contract.UpgradeToAndCall(&_PlazaSdk.TransactOpts, newImplementation, data)
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

// PlazaSdkPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the PlazaSdk contract.
type PlazaSdkPoolCreatedIterator struct {
	Event *PlazaSdkPoolCreated // Event containing the contract specifics and raw log

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
func (it *PlazaSdkPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkPoolCreated)
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
		it.Event = new(PlazaSdkPoolCreated)
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
func (it *PlazaSdkPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkPoolCreated represents a PoolCreated event raised by the PlazaSdk contract.
type PlazaSdkPoolCreated struct {
	Pool           common.Address
	ReserveAmount  *big.Int
	BondAmount     *big.Int
	LeverageAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xc8e774967d7f45e5814dbb800335ed47014999568397f08e79b0f322047f743f.
//
// Solidity: event PoolCreated(address pool, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterPoolCreated(opts *bind.FilterOpts) (*PlazaSdkPoolCreatedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkPoolCreatedIterator{contract: _PlazaSdk.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xc8e774967d7f45e5814dbb800335ed47014999568397f08e79b0f322047f743f.
//
// Solidity: event PoolCreated(address pool, uint256 reserveAmount, uint256 bondAmount, uint256 leverageAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PlazaSdkPoolCreated) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkPoolCreated)
				if err := _PlazaSdk.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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
func (_PlazaSdk *PlazaSdkFilterer) ParsePoolCreated(log types.Log) (*PlazaSdkPoolCreated, error) {
	event := new(PlazaSdkPoolCreated)
	if err := _PlazaSdk.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// PlazaSdkUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PlazaSdk contract.
type PlazaSdkUpgradedIterator struct {
	Event *PlazaSdkUpgraded // Event containing the contract specifics and raw log

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
func (it *PlazaSdkUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkUpgraded)
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
		it.Event = new(PlazaSdkUpgraded)
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
func (it *PlazaSdkUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkUpgraded represents a Upgraded event raised by the PlazaSdk contract.
type PlazaSdkUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PlazaSdk *PlazaSdkFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PlazaSdkUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkUpgradedIterator{contract: _PlazaSdk.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PlazaSdk *PlazaSdkFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PlazaSdkUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkUpgraded)
				if err := _PlazaSdk.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_PlazaSdk *PlazaSdkFilterer) ParseUpgraded(log types.Log) (*PlazaSdkUpgraded, error) {
	event := new(PlazaSdkUpgraded)
	if err := _PlazaSdk.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

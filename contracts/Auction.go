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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AccessDenied\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"AddressEmptyCode\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"AddressInsufficientBalance\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AlreadyClaimed\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionAlreadyEnded\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionFailed\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionHasEnded\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionStillOngoing\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionSucceededOrOngoing\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"BidAmountTooHigh\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"BidAmountTooLow\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"ERC1967NonPayable\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"EnforcedPause\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"ExpectedPause\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"FailedInnerCall\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"InvalidInitialization\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"InvalidSellAmount\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"NotInitializing\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"NothingToClaim\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\"},{\"inputs\":[{\"internalType\":\"enumAuction.State\",\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"totalSellReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"totalBuyCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"AuctionEnded\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidClaimed\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidPlaced\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidReduced\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidRefundAllocated\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidRemoved\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FailedAuctionBidRefundClaimed\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"LosingBidRefundClaimed\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Paused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Unpaused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"Upgraded\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MAX_BID_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"bid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bidCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextBidIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevBidIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"buyCouponToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claimBid\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claimRefund\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claimRefund\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"currentCouponAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"endAuction\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"highestBidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyCouponToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sellReserveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalBuyCouponAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBids\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_poolSaleLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lastBidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lowestBidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"pause\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pendingRefunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"poolSaleLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"sellReserveToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"slotSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumAuction.State\",\"name\":\"\",\"type\":\"uint8\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalBuyCouponAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSellReserveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"unpause\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"upgradeToAndCall\"}]",
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

// MAXBIDAMOUNT is a free data retrieval call binding the contract method 0xed0f8888.
//
// Solidity: function MAX_BID_AMOUNT() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) MAXBIDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "MAX_BID_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBIDAMOUNT is a free data retrieval call binding the contract method 0xed0f8888.
//
// Solidity: function MAX_BID_AMOUNT() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) MAXBIDAMOUNT() (*big.Int, error) {
	return _PlazaSdk.Contract.MAXBIDAMOUNT(&_PlazaSdk.CallOpts)
}

// MAXBIDAMOUNT is a free data retrieval call binding the contract method 0xed0f8888.
//
// Solidity: function MAX_BID_AMOUNT() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) MAXBIDAMOUNT() (*big.Int, error) {
	return _PlazaSdk.Contract.MAXBIDAMOUNT(&_PlazaSdk.CallOpts)
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

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_PlazaSdk *PlazaSdkSession) Beneficiary() (common.Address, error) {
	return _PlazaSdk.Contract.Beneficiary(&_PlazaSdk.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) Beneficiary() (common.Address, error) {
	return _PlazaSdk.Contract.Beneficiary(&_PlazaSdk.CallOpts)
}

// BidCount is a free data retrieval call binding the contract method 0xb40a5627.
//
// Solidity: function bidCount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) BidCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "bidCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidCount is a free data retrieval call binding the contract method 0xb40a5627.
//
// Solidity: function bidCount() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) BidCount() (*big.Int, error) {
	return _PlazaSdk.Contract.BidCount(&_PlazaSdk.CallOpts)
}

// BidCount is a free data retrieval call binding the contract method 0xb40a5627.
//
// Solidity: function bidCount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) BidCount() (*big.Int, error) {
	return _PlazaSdk.Contract.BidCount(&_PlazaSdk.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(address bidder, uint256 buyReserveAmount, uint256 sellCouponAmount, uint256 nextBidIndex, uint256 prevBidIndex, bool claimed)
func (_PlazaSdk *PlazaSdkCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	NextBidIndex     *big.Int
	PrevBidIndex     *big.Int
	Claimed          bool
}, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "bids", arg0)

	outstruct := new(struct {
		Bidder           common.Address
		BuyReserveAmount *big.Int
		SellCouponAmount *big.Int
		NextBidIndex     *big.Int
		PrevBidIndex     *big.Int
		Claimed          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BuyReserveAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellCouponAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NextBidIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PrevBidIndex = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(address bidder, uint256 buyReserveAmount, uint256 sellCouponAmount, uint256 nextBidIndex, uint256 prevBidIndex, bool claimed)
func (_PlazaSdk *PlazaSdkSession) Bids(arg0 *big.Int) (struct {
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	NextBidIndex     *big.Int
	PrevBidIndex     *big.Int
	Claimed          bool
}, error) {
	return _PlazaSdk.Contract.Bids(&_PlazaSdk.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(address bidder, uint256 buyReserveAmount, uint256 sellCouponAmount, uint256 nextBidIndex, uint256 prevBidIndex, bool claimed)
func (_PlazaSdk *PlazaSdkCallerSession) Bids(arg0 *big.Int) (struct {
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	NextBidIndex     *big.Int
	PrevBidIndex     *big.Int
	Claimed          bool
}, error) {
	return _PlazaSdk.Contract.Bids(&_PlazaSdk.CallOpts, arg0)
}

// BuyCouponToken is a free data retrieval call binding the contract method 0x119fdb0f.
//
// Solidity: function buyCouponToken() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) BuyCouponToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "buyCouponToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BuyCouponToken is a free data retrieval call binding the contract method 0x119fdb0f.
//
// Solidity: function buyCouponToken() view returns(address)
func (_PlazaSdk *PlazaSdkSession) BuyCouponToken() (common.Address, error) {
	return _PlazaSdk.Contract.BuyCouponToken(&_PlazaSdk.CallOpts)
}

// BuyCouponToken is a free data retrieval call binding the contract method 0x119fdb0f.
//
// Solidity: function buyCouponToken() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) BuyCouponToken() (common.Address, error) {
	return _PlazaSdk.Contract.BuyCouponToken(&_PlazaSdk.CallOpts)
}

// CurrentCouponAmount is a free data retrieval call binding the contract method 0xa9b001d9.
//
// Solidity: function currentCouponAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) CurrentCouponAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "currentCouponAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentCouponAmount is a free data retrieval call binding the contract method 0xa9b001d9.
//
// Solidity: function currentCouponAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) CurrentCouponAmount() (*big.Int, error) {
	return _PlazaSdk.Contract.CurrentCouponAmount(&_PlazaSdk.CallOpts)
}

// CurrentCouponAmount is a free data retrieval call binding the contract method 0xa9b001d9.
//
// Solidity: function currentCouponAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) CurrentCouponAmount() (*big.Int, error) {
	return _PlazaSdk.Contract.CurrentCouponAmount(&_PlazaSdk.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) EndTime() (*big.Int, error) {
	return _PlazaSdk.Contract.EndTime(&_PlazaSdk.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) EndTime() (*big.Int, error) {
	return _PlazaSdk.Contract.EndTime(&_PlazaSdk.CallOpts)
}

// HighestBidIndex is a free data retrieval call binding the contract method 0xead6f9ab.
//
// Solidity: function highestBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) HighestBidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "highestBidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBidIndex is a free data retrieval call binding the contract method 0xead6f9ab.
//
// Solidity: function highestBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) HighestBidIndex() (*big.Int, error) {
	return _PlazaSdk.Contract.HighestBidIndex(&_PlazaSdk.CallOpts)
}

// HighestBidIndex is a free data retrieval call binding the contract method 0xead6f9ab.
//
// Solidity: function highestBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) HighestBidIndex() (*big.Int, error) {
	return _PlazaSdk.Contract.HighestBidIndex(&_PlazaSdk.CallOpts)
}

// LastBidIndex is a free data retrieval call binding the contract method 0x26885d09.
//
// Solidity: function lastBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) LastBidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "lastBidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBidIndex is a free data retrieval call binding the contract method 0x26885d09.
//
// Solidity: function lastBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) LastBidIndex() (*big.Int, error) {
	return _PlazaSdk.Contract.LastBidIndex(&_PlazaSdk.CallOpts)
}

// LastBidIndex is a free data retrieval call binding the contract method 0x26885d09.
//
// Solidity: function lastBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) LastBidIndex() (*big.Int, error) {
	return _PlazaSdk.Contract.LastBidIndex(&_PlazaSdk.CallOpts)
}

// LowestBidIndex is a free data retrieval call binding the contract method 0xec1e4873.
//
// Solidity: function lowestBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) LowestBidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "lowestBidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowestBidIndex is a free data retrieval call binding the contract method 0xec1e4873.
//
// Solidity: function lowestBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) LowestBidIndex() (*big.Int, error) {
	return _PlazaSdk.Contract.LowestBidIndex(&_PlazaSdk.CallOpts)
}

// LowestBidIndex is a free data retrieval call binding the contract method 0xec1e4873.
//
// Solidity: function lowestBidIndex() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) LowestBidIndex() (*big.Int, error) {
	return _PlazaSdk.Contract.LowestBidIndex(&_PlazaSdk.CallOpts)
}

// MaxBids is a free data retrieval call binding the contract method 0x9f235ccd.
//
// Solidity: function maxBids() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) MaxBids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "maxBids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBids is a free data retrieval call binding the contract method 0x9f235ccd.
//
// Solidity: function maxBids() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) MaxBids() (*big.Int, error) {
	return _PlazaSdk.Contract.MaxBids(&_PlazaSdk.CallOpts)
}

// MaxBids is a free data retrieval call binding the contract method 0x9f235ccd.
//
// Solidity: function maxBids() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) MaxBids() (*big.Int, error) {
	return _PlazaSdk.Contract.MaxBids(&_PlazaSdk.CallOpts)
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

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) PendingRefunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "pendingRefunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _PlazaSdk.Contract.PendingRefunds(&_PlazaSdk.CallOpts, arg0)
}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _PlazaSdk.Contract.PendingRefunds(&_PlazaSdk.CallOpts, arg0)
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

// PoolSaleLimit is a free data retrieval call binding the contract method 0xc8c4e0fd.
//
// Solidity: function poolSaleLimit() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) PoolSaleLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "poolSaleLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSaleLimit is a free data retrieval call binding the contract method 0xc8c4e0fd.
//
// Solidity: function poolSaleLimit() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) PoolSaleLimit() (*big.Int, error) {
	return _PlazaSdk.Contract.PoolSaleLimit(&_PlazaSdk.CallOpts)
}

// PoolSaleLimit is a free data retrieval call binding the contract method 0xc8c4e0fd.
//
// Solidity: function poolSaleLimit() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) PoolSaleLimit() (*big.Int, error) {
	return _PlazaSdk.Contract.PoolSaleLimit(&_PlazaSdk.CallOpts)
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

// SellReserveToken is a free data retrieval call binding the contract method 0xde365a1b.
//
// Solidity: function sellReserveToken() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) SellReserveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "sellReserveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SellReserveToken is a free data retrieval call binding the contract method 0xde365a1b.
//
// Solidity: function sellReserveToken() view returns(address)
func (_PlazaSdk *PlazaSdkSession) SellReserveToken() (common.Address, error) {
	return _PlazaSdk.Contract.SellReserveToken(&_PlazaSdk.CallOpts)
}

// SellReserveToken is a free data retrieval call binding the contract method 0xde365a1b.
//
// Solidity: function sellReserveToken() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) SellReserveToken() (common.Address, error) {
	return _PlazaSdk.Contract.SellReserveToken(&_PlazaSdk.CallOpts)
}

// SlotSize is a free data retrieval call binding the contract method 0x628a0cf8.
//
// Solidity: function slotSize() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) SlotSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "slotSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotSize is a free data retrieval call binding the contract method 0x628a0cf8.
//
// Solidity: function slotSize() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) SlotSize() (*big.Int, error) {
	return _PlazaSdk.Contract.SlotSize(&_PlazaSdk.CallOpts)
}

// SlotSize is a free data retrieval call binding the contract method 0x628a0cf8.
//
// Solidity: function slotSize() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) SlotSize() (*big.Int, error) {
	return _PlazaSdk.Contract.SlotSize(&_PlazaSdk.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_PlazaSdk *PlazaSdkCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_PlazaSdk *PlazaSdkSession) State() (uint8, error) {
	return _PlazaSdk.Contract.State(&_PlazaSdk.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_PlazaSdk *PlazaSdkCallerSession) State() (uint8, error) {
	return _PlazaSdk.Contract.State(&_PlazaSdk.CallOpts)
}

// TotalBuyCouponAmount is a free data retrieval call binding the contract method 0xdcc8eaf6.
//
// Solidity: function totalBuyCouponAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) TotalBuyCouponAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "totalBuyCouponAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuyCouponAmount is a free data retrieval call binding the contract method 0xdcc8eaf6.
//
// Solidity: function totalBuyCouponAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) TotalBuyCouponAmount() (*big.Int, error) {
	return _PlazaSdk.Contract.TotalBuyCouponAmount(&_PlazaSdk.CallOpts)
}

// TotalBuyCouponAmount is a free data retrieval call binding the contract method 0xdcc8eaf6.
//
// Solidity: function totalBuyCouponAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) TotalBuyCouponAmount() (*big.Int, error) {
	return _PlazaSdk.Contract.TotalBuyCouponAmount(&_PlazaSdk.CallOpts)
}

// TotalSellReserveAmount is a free data retrieval call binding the contract method 0x26191ffa.
//
// Solidity: function totalSellReserveAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) TotalSellReserveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "totalSellReserveAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSellReserveAmount is a free data retrieval call binding the contract method 0x26191ffa.
//
// Solidity: function totalSellReserveAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) TotalSellReserveAmount() (*big.Int, error) {
	return _PlazaSdk.Contract.TotalSellReserveAmount(&_PlazaSdk.CallOpts)
}

// TotalSellReserveAmount is a free data retrieval call binding the contract method 0x26191ffa.
//
// Solidity: function totalSellReserveAmount() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) TotalSellReserveAmount() (*big.Int, error) {
	return _PlazaSdk.Contract.TotalSellReserveAmount(&_PlazaSdk.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 buyReserveAmount, uint256 sellCouponAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactor) Bid(opts *bind.TransactOpts, buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "bid", buyReserveAmount, sellCouponAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 buyReserveAmount, uint256 sellCouponAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkSession) Bid(buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Bid(&_PlazaSdk.TransactOpts, buyReserveAmount, sellCouponAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 buyReserveAmount, uint256 sellCouponAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactorSession) Bid(buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Bid(&_PlazaSdk.TransactOpts, buyReserveAmount, sellCouponAmount)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 bidIndex) returns()
func (_PlazaSdk *PlazaSdkTransactor) ClaimBid(opts *bind.TransactOpts, bidIndex *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "claimBid", bidIndex)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 bidIndex) returns()
func (_PlazaSdk *PlazaSdkSession) ClaimBid(bidIndex *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimBid(&_PlazaSdk.TransactOpts, bidIndex)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 bidIndex) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) ClaimBid(bidIndex *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimBid(&_PlazaSdk.TransactOpts, bidIndex)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 bidIndex) returns()
func (_PlazaSdk *PlazaSdkTransactor) ClaimRefund(opts *bind.TransactOpts, bidIndex *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "claimRefund", bidIndex)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 bidIndex) returns()
func (_PlazaSdk *PlazaSdkSession) ClaimRefund(bidIndex *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimRefund(&_PlazaSdk.TransactOpts, bidIndex)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 bidIndex) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) ClaimRefund(bidIndex *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimRefund(&_PlazaSdk.TransactOpts, bidIndex)
}

// ClaimRefund0 is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_PlazaSdk *PlazaSdkTransactor) ClaimRefund0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "claimRefund0")
}

// ClaimRefund0 is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_PlazaSdk *PlazaSdkSession) ClaimRefund0() (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimRefund0(&_PlazaSdk.TransactOpts)
}

// ClaimRefund0 is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) ClaimRefund0() (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimRefund0(&_PlazaSdk.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_PlazaSdk *PlazaSdkTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_PlazaSdk *PlazaSdkSession) EndAuction() (*types.Transaction, error) {
	return _PlazaSdk.Contract.EndAuction(&_PlazaSdk.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) EndAuction() (*types.Transaction, error) {
	return _PlazaSdk.Contract.EndAuction(&_PlazaSdk.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8aabdb8b.
//
// Solidity: function initialize(address _pool, address _buyCouponToken, address _sellReserveToken, uint256 _totalBuyCouponAmount, uint256 _endTime, uint256 _maxBids, address _beneficiary, uint256 _poolSaleLimit) returns()
func (_PlazaSdk *PlazaSdkTransactor) Initialize(opts *bind.TransactOpts, _pool common.Address, _buyCouponToken common.Address, _sellReserveToken common.Address, _totalBuyCouponAmount *big.Int, _endTime *big.Int, _maxBids *big.Int, _beneficiary common.Address, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "initialize", _pool, _buyCouponToken, _sellReserveToken, _totalBuyCouponAmount, _endTime, _maxBids, _beneficiary, _poolSaleLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x8aabdb8b.
//
// Solidity: function initialize(address _pool, address _buyCouponToken, address _sellReserveToken, uint256 _totalBuyCouponAmount, uint256 _endTime, uint256 _maxBids, address _beneficiary, uint256 _poolSaleLimit) returns()
func (_PlazaSdk *PlazaSdkSession) Initialize(_pool common.Address, _buyCouponToken common.Address, _sellReserveToken common.Address, _totalBuyCouponAmount *big.Int, _endTime *big.Int, _maxBids *big.Int, _beneficiary common.Address, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _pool, _buyCouponToken, _sellReserveToken, _totalBuyCouponAmount, _endTime, _maxBids, _beneficiary, _poolSaleLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x8aabdb8b.
//
// Solidity: function initialize(address _pool, address _buyCouponToken, address _sellReserveToken, uint256 _totalBuyCouponAmount, uint256 _endTime, uint256 _maxBids, address _beneficiary, uint256 _poolSaleLimit) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Initialize(_pool common.Address, _buyCouponToken common.Address, _sellReserveToken common.Address, _totalBuyCouponAmount *big.Int, _endTime *big.Int, _maxBids *big.Int, _beneficiary common.Address, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _pool, _buyCouponToken, _sellReserveToken, _totalBuyCouponAmount, _endTime, _maxBids, _beneficiary, _poolSaleLimit)
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

// PlazaSdkAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the PlazaSdk contract.
type PlazaSdkAuctionEndedIterator struct {
	Event *PlazaSdkAuctionEnded // Event containing the contract specifics and raw log

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
func (it *PlazaSdkAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkAuctionEnded)
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
		it.Event = new(PlazaSdkAuctionEnded)
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
func (it *PlazaSdkAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkAuctionEnded represents a AuctionEnded event raised by the PlazaSdk contract.
type PlazaSdkAuctionEnded struct {
	State                  uint8
	TotalSellReserveAmount *big.Int
	TotalBuyCouponAmount   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xbd6e0990969f60d91021509c9d5849868677e6886c512d34298688a9cd5e458c.
//
// Solidity: event AuctionEnded(uint8 state, uint256 totalSellReserveAmount, uint256 totalBuyCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*PlazaSdkAuctionEndedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkAuctionEndedIterator{contract: _PlazaSdk.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xbd6e0990969f60d91021509c9d5849868677e6886c512d34298688a9cd5e458c.
//
// Solidity: event AuctionEnded(uint8 state, uint256 totalSellReserveAmount, uint256 totalBuyCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *PlazaSdkAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkAuctionEnded)
				if err := _PlazaSdk.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xbd6e0990969f60d91021509c9d5849868677e6886c512d34298688a9cd5e458c.
//
// Solidity: event AuctionEnded(uint8 state, uint256 totalSellReserveAmount, uint256 totalBuyCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseAuctionEnded(log types.Log) (*PlazaSdkAuctionEnded, error) {
	event := new(PlazaSdkAuctionEnded)
	if err := _PlazaSdk.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkBidClaimedIterator is returned from FilterBidClaimed and is used to iterate over the raw logs and unpacked data for BidClaimed events raised by the PlazaSdk contract.
type PlazaSdkBidClaimedIterator struct {
	Event *PlazaSdkBidClaimed // Event containing the contract specifics and raw log

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
func (it *PlazaSdkBidClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkBidClaimed)
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
		it.Event = new(PlazaSdkBidClaimed)
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
func (it *PlazaSdkBidClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkBidClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkBidClaimed represents a BidClaimed event raised by the PlazaSdk contract.
type PlazaSdkBidClaimed struct {
	BidIndex         *big.Int
	Bidder           common.Address
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidClaimed is a free log retrieval operation binding the contract event 0xe68daa9f0be9a647eb4b9d00ba8932ad8171b86ec8cd767d14fd2ae946af3d21.
//
// Solidity: event BidClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterBidClaimed(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*PlazaSdkBidClaimedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "BidClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkBidClaimedIterator{contract: _PlazaSdk.contract, event: "BidClaimed", logs: logs, sub: sub}, nil
}

// WatchBidClaimed is a free log subscription operation binding the contract event 0xe68daa9f0be9a647eb4b9d00ba8932ad8171b86ec8cd767d14fd2ae946af3d21.
//
// Solidity: event BidClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchBidClaimed(opts *bind.WatchOpts, sink chan<- *PlazaSdkBidClaimed, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "BidClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkBidClaimed)
				if err := _PlazaSdk.contract.UnpackLog(event, "BidClaimed", log); err != nil {
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

// ParseBidClaimed is a log parse operation binding the contract event 0xe68daa9f0be9a647eb4b9d00ba8932ad8171b86ec8cd767d14fd2ae946af3d21.
//
// Solidity: event BidClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseBidClaimed(log types.Log) (*PlazaSdkBidClaimed, error) {
	event := new(PlazaSdkBidClaimed)
	if err := _PlazaSdk.contract.UnpackLog(event, "BidClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the PlazaSdk contract.
type PlazaSdkBidPlacedIterator struct {
	Event *PlazaSdkBidPlaced // Event containing the contract specifics and raw log

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
func (it *PlazaSdkBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkBidPlaced)
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
		it.Event = new(PlazaSdkBidPlaced)
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
func (it *PlazaSdkBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkBidPlaced represents a BidPlaced event raised by the PlazaSdk contract.
type PlazaSdkBidPlaced struct {
	BidIndex         *big.Int
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterBidPlaced(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*PlazaSdkBidPlacedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "BidPlaced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkBidPlacedIterator{contract: _PlazaSdk.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *PlazaSdkBidPlaced, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "BidPlaced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkBidPlaced)
				if err := _PlazaSdk.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseBidPlaced(log types.Log) (*PlazaSdkBidPlaced, error) {
	event := new(PlazaSdkBidPlaced)
	if err := _PlazaSdk.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkBidReducedIterator is returned from FilterBidReduced and is used to iterate over the raw logs and unpacked data for BidReduced events raised by the PlazaSdk contract.
type PlazaSdkBidReducedIterator struct {
	Event *PlazaSdkBidReduced // Event containing the contract specifics and raw log

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
func (it *PlazaSdkBidReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkBidReduced)
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
		it.Event = new(PlazaSdkBidReduced)
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
func (it *PlazaSdkBidReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkBidReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkBidReduced represents a BidReduced event raised by the PlazaSdk contract.
type PlazaSdkBidReduced struct {
	BidIndex         *big.Int
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidReduced is a free log retrieval operation binding the contract event 0x938f570a90ec62a7633da171c778b26404bb32a25e03756d49d4b79ecc4b785b.
//
// Solidity: event BidReduced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterBidReduced(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*PlazaSdkBidReducedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "BidReduced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkBidReducedIterator{contract: _PlazaSdk.contract, event: "BidReduced", logs: logs, sub: sub}, nil
}

// WatchBidReduced is a free log subscription operation binding the contract event 0x938f570a90ec62a7633da171c778b26404bb32a25e03756d49d4b79ecc4b785b.
//
// Solidity: event BidReduced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchBidReduced(opts *bind.WatchOpts, sink chan<- *PlazaSdkBidReduced, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "BidReduced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkBidReduced)
				if err := _PlazaSdk.contract.UnpackLog(event, "BidReduced", log); err != nil {
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

// ParseBidReduced is a log parse operation binding the contract event 0x938f570a90ec62a7633da171c778b26404bb32a25e03756d49d4b79ecc4b785b.
//
// Solidity: event BidReduced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseBidReduced(log types.Log) (*PlazaSdkBidReduced, error) {
	event := new(PlazaSdkBidReduced)
	if err := _PlazaSdk.contract.UnpackLog(event, "BidReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkBidRefundAllocatedIterator is returned from FilterBidRefundAllocated and is used to iterate over the raw logs and unpacked data for BidRefundAllocated events raised by the PlazaSdk contract.
type PlazaSdkBidRefundAllocatedIterator struct {
	Event *PlazaSdkBidRefundAllocated // Event containing the contract specifics and raw log

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
func (it *PlazaSdkBidRefundAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkBidRefundAllocated)
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
		it.Event = new(PlazaSdkBidRefundAllocated)
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
func (it *PlazaSdkBidRefundAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkBidRefundAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkBidRefundAllocated represents a BidRefundAllocated event raised by the PlazaSdk contract.
type PlazaSdkBidRefundAllocated struct {
	Bidder       common.Address
	CouponAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidRefundAllocated is a free log retrieval operation binding the contract event 0x1ab15874f0b1e2a964800d35f4bdd4a8e9c9f102b0262ad7f5aa21f195c5ece6.
//
// Solidity: event BidRefundAllocated(address indexed bidder, uint256 couponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterBidRefundAllocated(opts *bind.FilterOpts, bidder []common.Address) (*PlazaSdkBidRefundAllocatedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "BidRefundAllocated", bidderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkBidRefundAllocatedIterator{contract: _PlazaSdk.contract, event: "BidRefundAllocated", logs: logs, sub: sub}, nil
}

// WatchBidRefundAllocated is a free log subscription operation binding the contract event 0x1ab15874f0b1e2a964800d35f4bdd4a8e9c9f102b0262ad7f5aa21f195c5ece6.
//
// Solidity: event BidRefundAllocated(address indexed bidder, uint256 couponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchBidRefundAllocated(opts *bind.WatchOpts, sink chan<- *PlazaSdkBidRefundAllocated, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "BidRefundAllocated", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkBidRefundAllocated)
				if err := _PlazaSdk.contract.UnpackLog(event, "BidRefundAllocated", log); err != nil {
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

// ParseBidRefundAllocated is a log parse operation binding the contract event 0x1ab15874f0b1e2a964800d35f4bdd4a8e9c9f102b0262ad7f5aa21f195c5ece6.
//
// Solidity: event BidRefundAllocated(address indexed bidder, uint256 couponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseBidRefundAllocated(log types.Log) (*PlazaSdkBidRefundAllocated, error) {
	event := new(PlazaSdkBidRefundAllocated)
	if err := _PlazaSdk.contract.UnpackLog(event, "BidRefundAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkBidRemovedIterator is returned from FilterBidRemoved and is used to iterate over the raw logs and unpacked data for BidRemoved events raised by the PlazaSdk contract.
type PlazaSdkBidRemovedIterator struct {
	Event *PlazaSdkBidRemoved // Event containing the contract specifics and raw log

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
func (it *PlazaSdkBidRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkBidRemoved)
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
		it.Event = new(PlazaSdkBidRemoved)
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
func (it *PlazaSdkBidRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkBidRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkBidRemoved represents a BidRemoved event raised by the PlazaSdk contract.
type PlazaSdkBidRemoved struct {
	BidIndex         *big.Int
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidRemoved is a free log retrieval operation binding the contract event 0xc93523f51b5bdae47e5ae3ad4c018c4b8314a458265d3840bbe3ce8ca194e5d1.
//
// Solidity: event BidRemoved(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterBidRemoved(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*PlazaSdkBidRemovedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "BidRemoved", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkBidRemovedIterator{contract: _PlazaSdk.contract, event: "BidRemoved", logs: logs, sub: sub}, nil
}

// WatchBidRemoved is a free log subscription operation binding the contract event 0xc93523f51b5bdae47e5ae3ad4c018c4b8314a458265d3840bbe3ce8ca194e5d1.
//
// Solidity: event BidRemoved(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchBidRemoved(opts *bind.WatchOpts, sink chan<- *PlazaSdkBidRemoved, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "BidRemoved", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkBidRemoved)
				if err := _PlazaSdk.contract.UnpackLog(event, "BidRemoved", log); err != nil {
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

// ParseBidRemoved is a log parse operation binding the contract event 0xc93523f51b5bdae47e5ae3ad4c018c4b8314a458265d3840bbe3ce8ca194e5d1.
//
// Solidity: event BidRemoved(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseBidRemoved(log types.Log) (*PlazaSdkBidRemoved, error) {
	event := new(PlazaSdkBidRemoved)
	if err := _PlazaSdk.contract.UnpackLog(event, "BidRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkFailedAuctionBidRefundClaimedIterator is returned from FilterFailedAuctionBidRefundClaimed and is used to iterate over the raw logs and unpacked data for FailedAuctionBidRefundClaimed events raised by the PlazaSdk contract.
type PlazaSdkFailedAuctionBidRefundClaimedIterator struct {
	Event *PlazaSdkFailedAuctionBidRefundClaimed // Event containing the contract specifics and raw log

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
func (it *PlazaSdkFailedAuctionBidRefundClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkFailedAuctionBidRefundClaimed)
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
		it.Event = new(PlazaSdkFailedAuctionBidRefundClaimed)
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
func (it *PlazaSdkFailedAuctionBidRefundClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkFailedAuctionBidRefundClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkFailedAuctionBidRefundClaimed represents a FailedAuctionBidRefundClaimed event raised by the PlazaSdk contract.
type PlazaSdkFailedAuctionBidRefundClaimed struct {
	BidIndex         *big.Int
	Bidder           common.Address
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFailedAuctionBidRefundClaimed is a free log retrieval operation binding the contract event 0xe53799b6cbe78f2d3592976e014800e6a394c793001be11eb8faf86f68a5de4f.
//
// Solidity: event FailedAuctionBidRefundClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterFailedAuctionBidRefundClaimed(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*PlazaSdkFailedAuctionBidRefundClaimedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "FailedAuctionBidRefundClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkFailedAuctionBidRefundClaimedIterator{contract: _PlazaSdk.contract, event: "FailedAuctionBidRefundClaimed", logs: logs, sub: sub}, nil
}

// WatchFailedAuctionBidRefundClaimed is a free log subscription operation binding the contract event 0xe53799b6cbe78f2d3592976e014800e6a394c793001be11eb8faf86f68a5de4f.
//
// Solidity: event FailedAuctionBidRefundClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchFailedAuctionBidRefundClaimed(opts *bind.WatchOpts, sink chan<- *PlazaSdkFailedAuctionBidRefundClaimed, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "FailedAuctionBidRefundClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkFailedAuctionBidRefundClaimed)
				if err := _PlazaSdk.contract.UnpackLog(event, "FailedAuctionBidRefundClaimed", log); err != nil {
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

// ParseFailedAuctionBidRefundClaimed is a log parse operation binding the contract event 0xe53799b6cbe78f2d3592976e014800e6a394c793001be11eb8faf86f68a5de4f.
//
// Solidity: event FailedAuctionBidRefundClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseFailedAuctionBidRefundClaimed(log types.Log) (*PlazaSdkFailedAuctionBidRefundClaimed, error) {
	event := new(PlazaSdkFailedAuctionBidRefundClaimed)
	if err := _PlazaSdk.contract.UnpackLog(event, "FailedAuctionBidRefundClaimed", log); err != nil {
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

// PlazaSdkLosingBidRefundClaimedIterator is returned from FilterLosingBidRefundClaimed and is used to iterate over the raw logs and unpacked data for LosingBidRefundClaimed events raised by the PlazaSdk contract.
type PlazaSdkLosingBidRefundClaimedIterator struct {
	Event *PlazaSdkLosingBidRefundClaimed // Event containing the contract specifics and raw log

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
func (it *PlazaSdkLosingBidRefundClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkLosingBidRefundClaimed)
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
		it.Event = new(PlazaSdkLosingBidRefundClaimed)
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
func (it *PlazaSdkLosingBidRefundClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkLosingBidRefundClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkLosingBidRefundClaimed represents a LosingBidRefundClaimed event raised by the PlazaSdk contract.
type PlazaSdkLosingBidRefundClaimed struct {
	Bidder           common.Address
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLosingBidRefundClaimed is a free log retrieval operation binding the contract event 0xb11b7956ad36f0bf4cee602c022aec98ee0cec801ba8b54def5b23729eb7bfc9.
//
// Solidity: event LosingBidRefundClaimed(address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterLosingBidRefundClaimed(opts *bind.FilterOpts, bidder []common.Address) (*PlazaSdkLosingBidRefundClaimedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "LosingBidRefundClaimed", bidderRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkLosingBidRefundClaimedIterator{contract: _PlazaSdk.contract, event: "LosingBidRefundClaimed", logs: logs, sub: sub}, nil
}

// WatchLosingBidRefundClaimed is a free log subscription operation binding the contract event 0xb11b7956ad36f0bf4cee602c022aec98ee0cec801ba8b54def5b23729eb7bfc9.
//
// Solidity: event LosingBidRefundClaimed(address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchLosingBidRefundClaimed(opts *bind.WatchOpts, sink chan<- *PlazaSdkLosingBidRefundClaimed, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "LosingBidRefundClaimed", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkLosingBidRefundClaimed)
				if err := _PlazaSdk.contract.UnpackLog(event, "LosingBidRefundClaimed", log); err != nil {
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

// ParseLosingBidRefundClaimed is a log parse operation binding the contract event 0xb11b7956ad36f0bf4cee602c022aec98ee0cec801ba8b54def5b23729eb7bfc9.
//
// Solidity: event LosingBidRefundClaimed(address indexed bidder, uint256 sellCouponAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseLosingBidRefundClaimed(log types.Log) (*PlazaSdkLosingBidRefundClaimed, error) {
	event := new(PlazaSdkLosingBidRefundClaimed)
	if err := _PlazaSdk.contract.UnpackLog(event, "LosingBidRefundClaimed", log); err != nil {
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

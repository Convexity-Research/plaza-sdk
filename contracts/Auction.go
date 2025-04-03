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

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AccessDenied\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"AddressEmptyCode\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"AddressInsufficientBalance\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AlreadyClaimed\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionAlreadyEnded\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionFailed\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionHasEnded\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionStillOngoing\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"AuctionSucceededOrOngoing\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"BidAmountTooHigh\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"BidAmountTooLow\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"ERC1967NonPayable\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"EnforcedPause\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"ExpectedPause\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"FailedInnerCall\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"InvalidInitialization\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"InvalidSellAmount\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"NotInitializing\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"NothingToClaim\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\"},{\"inputs\":[],\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\"},{\"inputs\":[{\"internalType\":\"enumAuction.State\",\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"totalSellReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"totalBuyCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"AuctionEnded\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidClaimed\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidPlaced\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidReduced\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"couponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidRefundAllocated\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BidRemoved\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FailedAuctionBidRefundClaimed\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"LosingBidRefundClaimed\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Paused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Unpaused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"Upgraded\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MAX_BID_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"bid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bidCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"bids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"buyReserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellCouponAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextBidIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevBidIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"buyCouponToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claimBid\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claimRefund\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claimRefund\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"currentCouponAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"endAuction\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"highestBidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyCouponToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sellReserveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalBuyCouponAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBids\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_poolSaleLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lastBidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"lowestBidIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"pause\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pendingRefunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"poolSaleLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"sellReserveToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"slotSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumAuction.State\",\"name\":\"\",\"type\":\"uint8\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalBuyCouponAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSellReserveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"unpause\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"upgradeToAndCall\"}]",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// MAXBIDAMOUNT is a free data retrieval call binding the contract method 0xed0f8888.
//
// Solidity: function MAX_BID_AMOUNT() view returns(uint256)
func (_Auction *AuctionCaller) MAXBIDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "MAX_BID_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBIDAMOUNT is a free data retrieval call binding the contract method 0xed0f8888.
//
// Solidity: function MAX_BID_AMOUNT() view returns(uint256)
func (_Auction *AuctionSession) MAXBIDAMOUNT() (*big.Int, error) {
	return _Auction.Contract.MAXBIDAMOUNT(&_Auction.CallOpts)
}

// MAXBIDAMOUNT is a free data retrieval call binding the contract method 0xed0f8888.
//
// Solidity: function MAX_BID_AMOUNT() view returns(uint256)
func (_Auction *AuctionCallerSession) MAXBIDAMOUNT() (*big.Int, error) {
	return _Auction.Contract.MAXBIDAMOUNT(&_Auction.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Auction *AuctionCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Auction *AuctionSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Auction.Contract.UPGRADEINTERFACEVERSION(&_Auction.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Auction *AuctionCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Auction.Contract.UPGRADEINTERFACEVERSION(&_Auction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Auction *AuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Auction *AuctionSession) Beneficiary() (common.Address, error) {
	return _Auction.Contract.Beneficiary(&_Auction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_Auction *AuctionCallerSession) Beneficiary() (common.Address, error) {
	return _Auction.Contract.Beneficiary(&_Auction.CallOpts)
}

// BidCount is a free data retrieval call binding the contract method 0xb40a5627.
//
// Solidity: function bidCount() view returns(uint256)
func (_Auction *AuctionCaller) BidCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "bidCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidCount is a free data retrieval call binding the contract method 0xb40a5627.
//
// Solidity: function bidCount() view returns(uint256)
func (_Auction *AuctionSession) BidCount() (*big.Int, error) {
	return _Auction.Contract.BidCount(&_Auction.CallOpts)
}

// BidCount is a free data retrieval call binding the contract method 0xb40a5627.
//
// Solidity: function bidCount() view returns(uint256)
func (_Auction *AuctionCallerSession) BidCount() (*big.Int, error) {
	return _Auction.Contract.BidCount(&_Auction.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(address bidder, uint256 buyReserveAmount, uint256 sellCouponAmount, uint256 nextBidIndex, uint256 prevBidIndex, bool claimed)
func (_Auction *AuctionCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	NextBidIndex     *big.Int
	PrevBidIndex     *big.Int
	Claimed          bool
}, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "bids", arg0)

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
func (_Auction *AuctionSession) Bids(arg0 *big.Int) (struct {
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	NextBidIndex     *big.Int
	PrevBidIndex     *big.Int
	Claimed          bool
}, error) {
	return _Auction.Contract.Bids(&_Auction.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(address bidder, uint256 buyReserveAmount, uint256 sellCouponAmount, uint256 nextBidIndex, uint256 prevBidIndex, bool claimed)
func (_Auction *AuctionCallerSession) Bids(arg0 *big.Int) (struct {
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	NextBidIndex     *big.Int
	PrevBidIndex     *big.Int
	Claimed          bool
}, error) {
	return _Auction.Contract.Bids(&_Auction.CallOpts, arg0)
}

// BuyCouponToken is a free data retrieval call binding the contract method 0x119fdb0f.
//
// Solidity: function buyCouponToken() view returns(address)
func (_Auction *AuctionCaller) BuyCouponToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "buyCouponToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BuyCouponToken is a free data retrieval call binding the contract method 0x119fdb0f.
//
// Solidity: function buyCouponToken() view returns(address)
func (_Auction *AuctionSession) BuyCouponToken() (common.Address, error) {
	return _Auction.Contract.BuyCouponToken(&_Auction.CallOpts)
}

// BuyCouponToken is a free data retrieval call binding the contract method 0x119fdb0f.
//
// Solidity: function buyCouponToken() view returns(address)
func (_Auction *AuctionCallerSession) BuyCouponToken() (common.Address, error) {
	return _Auction.Contract.BuyCouponToken(&_Auction.CallOpts)
}

// CurrentCouponAmount is a free data retrieval call binding the contract method 0xa9b001d9.
//
// Solidity: function currentCouponAmount() view returns(uint256)
func (_Auction *AuctionCaller) CurrentCouponAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "currentCouponAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentCouponAmount is a free data retrieval call binding the contract method 0xa9b001d9.
//
// Solidity: function currentCouponAmount() view returns(uint256)
func (_Auction *AuctionSession) CurrentCouponAmount() (*big.Int, error) {
	return _Auction.Contract.CurrentCouponAmount(&_Auction.CallOpts)
}

// CurrentCouponAmount is a free data retrieval call binding the contract method 0xa9b001d9.
//
// Solidity: function currentCouponAmount() view returns(uint256)
func (_Auction *AuctionCallerSession) CurrentCouponAmount() (*big.Int, error) {
	return _Auction.Contract.CurrentCouponAmount(&_Auction.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCallerSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// HighestBidIndex is a free data retrieval call binding the contract method 0xead6f9ab.
//
// Solidity: function highestBidIndex() view returns(uint256)
func (_Auction *AuctionCaller) HighestBidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBidIndex is a free data retrieval call binding the contract method 0xead6f9ab.
//
// Solidity: function highestBidIndex() view returns(uint256)
func (_Auction *AuctionSession) HighestBidIndex() (*big.Int, error) {
	return _Auction.Contract.HighestBidIndex(&_Auction.CallOpts)
}

// HighestBidIndex is a free data retrieval call binding the contract method 0xead6f9ab.
//
// Solidity: function highestBidIndex() view returns(uint256)
func (_Auction *AuctionCallerSession) HighestBidIndex() (*big.Int, error) {
	return _Auction.Contract.HighestBidIndex(&_Auction.CallOpts)
}

// LastBidIndex is a free data retrieval call binding the contract method 0x26885d09.
//
// Solidity: function lastBidIndex() view returns(uint256)
func (_Auction *AuctionCaller) LastBidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "lastBidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBidIndex is a free data retrieval call binding the contract method 0x26885d09.
//
// Solidity: function lastBidIndex() view returns(uint256)
func (_Auction *AuctionSession) LastBidIndex() (*big.Int, error) {
	return _Auction.Contract.LastBidIndex(&_Auction.CallOpts)
}

// LastBidIndex is a free data retrieval call binding the contract method 0x26885d09.
//
// Solidity: function lastBidIndex() view returns(uint256)
func (_Auction *AuctionCallerSession) LastBidIndex() (*big.Int, error) {
	return _Auction.Contract.LastBidIndex(&_Auction.CallOpts)
}

// LowestBidIndex is a free data retrieval call binding the contract method 0xec1e4873.
//
// Solidity: function lowestBidIndex() view returns(uint256)
func (_Auction *AuctionCaller) LowestBidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "lowestBidIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowestBidIndex is a free data retrieval call binding the contract method 0xec1e4873.
//
// Solidity: function lowestBidIndex() view returns(uint256)
func (_Auction *AuctionSession) LowestBidIndex() (*big.Int, error) {
	return _Auction.Contract.LowestBidIndex(&_Auction.CallOpts)
}

// LowestBidIndex is a free data retrieval call binding the contract method 0xec1e4873.
//
// Solidity: function lowestBidIndex() view returns(uint256)
func (_Auction *AuctionCallerSession) LowestBidIndex() (*big.Int, error) {
	return _Auction.Contract.LowestBidIndex(&_Auction.CallOpts)
}

// MaxBids is a free data retrieval call binding the contract method 0x9f235ccd.
//
// Solidity: function maxBids() view returns(uint256)
func (_Auction *AuctionCaller) MaxBids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "maxBids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBids is a free data retrieval call binding the contract method 0x9f235ccd.
//
// Solidity: function maxBids() view returns(uint256)
func (_Auction *AuctionSession) MaxBids() (*big.Int, error) {
	return _Auction.Contract.MaxBids(&_Auction.CallOpts)
}

// MaxBids is a free data retrieval call binding the contract method 0x9f235ccd.
//
// Solidity: function maxBids() view returns(uint256)
func (_Auction *AuctionCallerSession) MaxBids() (*big.Int, error) {
	return _Auction.Contract.MaxBids(&_Auction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Auction *AuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Auction *AuctionSession) Paused() (bool, error) {
	return _Auction.Contract.Paused(&_Auction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Auction *AuctionCallerSession) Paused() (bool, error) {
	return _Auction.Contract.Paused(&_Auction.CallOpts)
}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_Auction *AuctionCaller) PendingRefunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "pendingRefunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_Auction *AuctionSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.PendingRefunds(&_Auction.CallOpts, arg0)
}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_Auction *AuctionCallerSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.PendingRefunds(&_Auction.CallOpts, arg0)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Auction *AuctionCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Auction *AuctionSession) Pool() (common.Address, error) {
	return _Auction.Contract.Pool(&_Auction.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Auction *AuctionCallerSession) Pool() (common.Address, error) {
	return _Auction.Contract.Pool(&_Auction.CallOpts)
}

// PoolSaleLimit is a free data retrieval call binding the contract method 0xc8c4e0fd.
//
// Solidity: function poolSaleLimit() view returns(uint256)
func (_Auction *AuctionCaller) PoolSaleLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "poolSaleLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolSaleLimit is a free data retrieval call binding the contract method 0xc8c4e0fd.
//
// Solidity: function poolSaleLimit() view returns(uint256)
func (_Auction *AuctionSession) PoolSaleLimit() (*big.Int, error) {
	return _Auction.Contract.PoolSaleLimit(&_Auction.CallOpts)
}

// PoolSaleLimit is a free data retrieval call binding the contract method 0xc8c4e0fd.
//
// Solidity: function poolSaleLimit() view returns(uint256)
func (_Auction *AuctionCallerSession) PoolSaleLimit() (*big.Int, error) {
	return _Auction.Contract.PoolSaleLimit(&_Auction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Auction *AuctionCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Auction *AuctionSession) ProxiableUUID() ([32]byte, error) {
	return _Auction.Contract.ProxiableUUID(&_Auction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Auction *AuctionCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Auction.Contract.ProxiableUUID(&_Auction.CallOpts)
}

// SellReserveToken is a free data retrieval call binding the contract method 0xde365a1b.
//
// Solidity: function sellReserveToken() view returns(address)
func (_Auction *AuctionCaller) SellReserveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "sellReserveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SellReserveToken is a free data retrieval call binding the contract method 0xde365a1b.
//
// Solidity: function sellReserveToken() view returns(address)
func (_Auction *AuctionSession) SellReserveToken() (common.Address, error) {
	return _Auction.Contract.SellReserveToken(&_Auction.CallOpts)
}

// SellReserveToken is a free data retrieval call binding the contract method 0xde365a1b.
//
// Solidity: function sellReserveToken() view returns(address)
func (_Auction *AuctionCallerSession) SellReserveToken() (common.Address, error) {
	return _Auction.Contract.SellReserveToken(&_Auction.CallOpts)
}

// SlotSize is a free data retrieval call binding the contract method 0x628a0cf8.
//
// Solidity: function slotSize() view returns(uint256)
func (_Auction *AuctionCaller) SlotSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "slotSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotSize is a free data retrieval call binding the contract method 0x628a0cf8.
//
// Solidity: function slotSize() view returns(uint256)
func (_Auction *AuctionSession) SlotSize() (*big.Int, error) {
	return _Auction.Contract.SlotSize(&_Auction.CallOpts)
}

// SlotSize is a free data retrieval call binding the contract method 0x628a0cf8.
//
// Solidity: function slotSize() view returns(uint256)
func (_Auction *AuctionCallerSession) SlotSize() (*big.Int, error) {
	return _Auction.Contract.SlotSize(&_Auction.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Auction *AuctionCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Auction *AuctionSession) State() (uint8, error) {
	return _Auction.Contract.State(&_Auction.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Auction *AuctionCallerSession) State() (uint8, error) {
	return _Auction.Contract.State(&_Auction.CallOpts)
}

// TotalBuyCouponAmount is a free data retrieval call binding the contract method 0xdcc8eaf6.
//
// Solidity: function totalBuyCouponAmount() view returns(uint256)
func (_Auction *AuctionCaller) TotalBuyCouponAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "totalBuyCouponAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuyCouponAmount is a free data retrieval call binding the contract method 0xdcc8eaf6.
//
// Solidity: function totalBuyCouponAmount() view returns(uint256)
func (_Auction *AuctionSession) TotalBuyCouponAmount() (*big.Int, error) {
	return _Auction.Contract.TotalBuyCouponAmount(&_Auction.CallOpts)
}

// TotalBuyCouponAmount is a free data retrieval call binding the contract method 0xdcc8eaf6.
//
// Solidity: function totalBuyCouponAmount() view returns(uint256)
func (_Auction *AuctionCallerSession) TotalBuyCouponAmount() (*big.Int, error) {
	return _Auction.Contract.TotalBuyCouponAmount(&_Auction.CallOpts)
}

// TotalSellReserveAmount is a free data retrieval call binding the contract method 0x26191ffa.
//
// Solidity: function totalSellReserveAmount() view returns(uint256)
func (_Auction *AuctionCaller) TotalSellReserveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "totalSellReserveAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSellReserveAmount is a free data retrieval call binding the contract method 0x26191ffa.
//
// Solidity: function totalSellReserveAmount() view returns(uint256)
func (_Auction *AuctionSession) TotalSellReserveAmount() (*big.Int, error) {
	return _Auction.Contract.TotalSellReserveAmount(&_Auction.CallOpts)
}

// TotalSellReserveAmount is a free data retrieval call binding the contract method 0x26191ffa.
//
// Solidity: function totalSellReserveAmount() view returns(uint256)
func (_Auction *AuctionCallerSession) TotalSellReserveAmount() (*big.Int, error) {
	return _Auction.Contract.TotalSellReserveAmount(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 buyReserveAmount, uint256 sellCouponAmount) returns(uint256)
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts, buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bid", buyReserveAmount, sellCouponAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 buyReserveAmount, uint256 sellCouponAmount) returns(uint256)
func (_Auction *AuctionSession) Bid(buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, buyReserveAmount, sellCouponAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 buyReserveAmount, uint256 sellCouponAmount) returns(uint256)
func (_Auction *AuctionTransactorSession) Bid(buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts, buyReserveAmount, sellCouponAmount)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 bidIndex) returns()
func (_Auction *AuctionTransactor) ClaimBid(opts *bind.TransactOpts, bidIndex *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "claimBid", bidIndex)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 bidIndex) returns()
func (_Auction *AuctionSession) ClaimBid(bidIndex *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.ClaimBid(&_Auction.TransactOpts, bidIndex)
}

// ClaimBid is a paid mutator transaction binding the contract method 0x21113057.
//
// Solidity: function claimBid(uint256 bidIndex) returns()
func (_Auction *AuctionTransactorSession) ClaimBid(bidIndex *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.ClaimBid(&_Auction.TransactOpts, bidIndex)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 bidIndex) returns()
func (_Auction *AuctionTransactor) ClaimRefund(opts *bind.TransactOpts, bidIndex *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "claimRefund", bidIndex)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 bidIndex) returns()
func (_Auction *AuctionSession) ClaimRefund(bidIndex *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.ClaimRefund(&_Auction.TransactOpts, bidIndex)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 bidIndex) returns()
func (_Auction *AuctionTransactorSession) ClaimRefund(bidIndex *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.ClaimRefund(&_Auction.TransactOpts, bidIndex)
}

// ClaimRefund0 is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_Auction *AuctionTransactor) ClaimRefund0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "claimRefund0")
}

// ClaimRefund0 is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_Auction *AuctionSession) ClaimRefund0() (*types.Transaction, error) {
	return _Auction.Contract.ClaimRefund0(&_Auction.TransactOpts)
}

// ClaimRefund0 is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_Auction *AuctionTransactorSession) ClaimRefund0() (*types.Transaction, error) {
	return _Auction.Contract.ClaimRefund0(&_Auction.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Auction *AuctionTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Auction *AuctionSession) EndAuction() (*types.Transaction, error) {
	return _Auction.Contract.EndAuction(&_Auction.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Auction *AuctionTransactorSession) EndAuction() (*types.Transaction, error) {
	return _Auction.Contract.EndAuction(&_Auction.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8aabdb8b.
//
// Solidity: function initialize(address _pool, address _buyCouponToken, address _sellReserveToken, uint256 _totalBuyCouponAmount, uint256 _endTime, uint256 _maxBids, address _beneficiary, uint256 _poolSaleLimit) returns()
func (_Auction *AuctionTransactor) Initialize(opts *bind.TransactOpts, _pool common.Address, _buyCouponToken common.Address, _sellReserveToken common.Address, _totalBuyCouponAmount *big.Int, _endTime *big.Int, _maxBids *big.Int, _beneficiary common.Address, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "initialize", _pool, _buyCouponToken, _sellReserveToken, _totalBuyCouponAmount, _endTime, _maxBids, _beneficiary, _poolSaleLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x8aabdb8b.
//
// Solidity: function initialize(address _pool, address _buyCouponToken, address _sellReserveToken, uint256 _totalBuyCouponAmount, uint256 _endTime, uint256 _maxBids, address _beneficiary, uint256 _poolSaleLimit) returns()
func (_Auction *AuctionSession) Initialize(_pool common.Address, _buyCouponToken common.Address, _sellReserveToken common.Address, _totalBuyCouponAmount *big.Int, _endTime *big.Int, _maxBids *big.Int, _beneficiary common.Address, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Initialize(&_Auction.TransactOpts, _pool, _buyCouponToken, _sellReserveToken, _totalBuyCouponAmount, _endTime, _maxBids, _beneficiary, _poolSaleLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x8aabdb8b.
//
// Solidity: function initialize(address _pool, address _buyCouponToken, address _sellReserveToken, uint256 _totalBuyCouponAmount, uint256 _endTime, uint256 _maxBids, address _beneficiary, uint256 _poolSaleLimit) returns()
func (_Auction *AuctionTransactorSession) Initialize(_pool common.Address, _buyCouponToken common.Address, _sellReserveToken common.Address, _totalBuyCouponAmount *big.Int, _endTime *big.Int, _maxBids *big.Int, _beneficiary common.Address, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Initialize(&_Auction.TransactOpts, _pool, _buyCouponToken, _sellReserveToken, _totalBuyCouponAmount, _endTime, _maxBids, _beneficiary, _poolSaleLimit)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Auction *AuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Auction *AuctionSession) Pause() (*types.Transaction, error) {
	return _Auction.Contract.Pause(&_Auction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Auction *AuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _Auction.Contract.Pause(&_Auction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Auction *AuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Auction *AuctionSession) Unpause() (*types.Transaction, error) {
	return _Auction.Contract.Unpause(&_Auction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Auction *AuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _Auction.Contract.Unpause(&_Auction.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Auction *AuctionTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Auction *AuctionSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Auction.Contract.UpgradeToAndCall(&_Auction.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Auction *AuctionTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Auction.Contract.UpgradeToAndCall(&_Auction.TransactOpts, newImplementation, data)
}

// AuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the Auction contract.
type AuctionAuctionEndedIterator struct {
	Event *AuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *AuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionAuctionEnded)
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
		it.Event = new(AuctionAuctionEnded)
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
func (it *AuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionAuctionEnded represents a AuctionEnded event raised by the Auction contract.
type AuctionAuctionEnded struct {
	State                  uint8
	TotalSellReserveAmount *big.Int
	TotalBuyCouponAmount   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xbd6e0990969f60d91021509c9d5849868677e6886c512d34298688a9cd5e458c.
//
// Solidity: event AuctionEnded(uint8 state, uint256 totalSellReserveAmount, uint256 totalBuyCouponAmount)
func (_Auction *AuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*AuctionAuctionEndedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &AuctionAuctionEndedIterator{contract: _Auction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xbd6e0990969f60d91021509c9d5849868677e6886c512d34298688a9cd5e458c.
//
// Solidity: event AuctionEnded(uint8 state, uint256 totalSellReserveAmount, uint256 totalBuyCouponAmount)
func (_Auction *AuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *AuctionAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionAuctionEnded)
				if err := _Auction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseAuctionEnded(log types.Log) (*AuctionAuctionEnded, error) {
	event := new(AuctionAuctionEnded)
	if err := _Auction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionBidClaimedIterator is returned from FilterBidClaimed and is used to iterate over the raw logs and unpacked data for BidClaimed events raised by the Auction contract.
type AuctionBidClaimedIterator struct {
	Event *AuctionBidClaimed // Event containing the contract specifics and raw log

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
func (it *AuctionBidClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidClaimed)
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
		it.Event = new(AuctionBidClaimed)
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
func (it *AuctionBidClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidClaimed represents a BidClaimed event raised by the Auction contract.
type AuctionBidClaimed struct {
	BidIndex         *big.Int
	Bidder           common.Address
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidClaimed is a free log retrieval operation binding the contract event 0xe68daa9f0be9a647eb4b9d00ba8932ad8171b86ec8cd767d14fd2ae946af3d21.
//
// Solidity: event BidClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) FilterBidClaimed(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*AuctionBidClaimedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidClaimedIterator{contract: _Auction.contract, event: "BidClaimed", logs: logs, sub: sub}, nil
}

// WatchBidClaimed is a free log subscription operation binding the contract event 0xe68daa9f0be9a647eb4b9d00ba8932ad8171b86ec8cd767d14fd2ae946af3d21.
//
// Solidity: event BidClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) WatchBidClaimed(opts *bind.WatchOpts, sink chan<- *AuctionBidClaimed, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidClaimed)
				if err := _Auction.contract.UnpackLog(event, "BidClaimed", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseBidClaimed(log types.Log) (*AuctionBidClaimed, error) {
	event := new(AuctionBidClaimed)
	if err := _Auction.contract.UnpackLog(event, "BidClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the Auction contract.
type AuctionBidPlacedIterator struct {
	Event *AuctionBidPlaced // Event containing the contract specifics and raw log

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
func (it *AuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidPlaced)
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
		it.Event = new(AuctionBidPlaced)
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
func (it *AuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidPlaced represents a BidPlaced event raised by the Auction contract.
type AuctionBidPlaced struct {
	BidIndex         *big.Int
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*AuctionBidPlacedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidPlaced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidPlacedIterator{contract: _Auction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x51db8e23b3f4479b162fd48823b8402895442b8f6cfd94f66239391881ec7b6f.
//
// Solidity: event BidPlaced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *AuctionBidPlaced, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidPlaced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidPlaced)
				if err := _Auction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseBidPlaced(log types.Log) (*AuctionBidPlaced, error) {
	event := new(AuctionBidPlaced)
	if err := _Auction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionBidReducedIterator is returned from FilterBidReduced and is used to iterate over the raw logs and unpacked data for BidReduced events raised by the Auction contract.
type AuctionBidReducedIterator struct {
	Event *AuctionBidReduced // Event containing the contract specifics and raw log

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
func (it *AuctionBidReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidReduced)
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
		it.Event = new(AuctionBidReduced)
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
func (it *AuctionBidReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidReduced represents a BidReduced event raised by the Auction contract.
type AuctionBidReduced struct {
	BidIndex         *big.Int
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidReduced is a free log retrieval operation binding the contract event 0x938f570a90ec62a7633da171c778b26404bb32a25e03756d49d4b79ecc4b785b.
//
// Solidity: event BidReduced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) FilterBidReduced(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*AuctionBidReducedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidReduced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidReducedIterator{contract: _Auction.contract, event: "BidReduced", logs: logs, sub: sub}, nil
}

// WatchBidReduced is a free log subscription operation binding the contract event 0x938f570a90ec62a7633da171c778b26404bb32a25e03756d49d4b79ecc4b785b.
//
// Solidity: event BidReduced(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) WatchBidReduced(opts *bind.WatchOpts, sink chan<- *AuctionBidReduced, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidReduced", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidReduced)
				if err := _Auction.contract.UnpackLog(event, "BidReduced", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseBidReduced(log types.Log) (*AuctionBidReduced, error) {
	event := new(AuctionBidReduced)
	if err := _Auction.contract.UnpackLog(event, "BidReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionBidRefundAllocatedIterator is returned from FilterBidRefundAllocated and is used to iterate over the raw logs and unpacked data for BidRefundAllocated events raised by the Auction contract.
type AuctionBidRefundAllocatedIterator struct {
	Event *AuctionBidRefundAllocated // Event containing the contract specifics and raw log

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
func (it *AuctionBidRefundAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidRefundAllocated)
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
		it.Event = new(AuctionBidRefundAllocated)
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
func (it *AuctionBidRefundAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidRefundAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidRefundAllocated represents a BidRefundAllocated event raised by the Auction contract.
type AuctionBidRefundAllocated struct {
	Bidder       common.Address
	CouponAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidRefundAllocated is a free log retrieval operation binding the contract event 0x1ab15874f0b1e2a964800d35f4bdd4a8e9c9f102b0262ad7f5aa21f195c5ece6.
//
// Solidity: event BidRefundAllocated(address indexed bidder, uint256 couponAmount)
func (_Auction *AuctionFilterer) FilterBidRefundAllocated(opts *bind.FilterOpts, bidder []common.Address) (*AuctionBidRefundAllocatedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidRefundAllocated", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidRefundAllocatedIterator{contract: _Auction.contract, event: "BidRefundAllocated", logs: logs, sub: sub}, nil
}

// WatchBidRefundAllocated is a free log subscription operation binding the contract event 0x1ab15874f0b1e2a964800d35f4bdd4a8e9c9f102b0262ad7f5aa21f195c5ece6.
//
// Solidity: event BidRefundAllocated(address indexed bidder, uint256 couponAmount)
func (_Auction *AuctionFilterer) WatchBidRefundAllocated(opts *bind.WatchOpts, sink chan<- *AuctionBidRefundAllocated, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidRefundAllocated", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidRefundAllocated)
				if err := _Auction.contract.UnpackLog(event, "BidRefundAllocated", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseBidRefundAllocated(log types.Log) (*AuctionBidRefundAllocated, error) {
	event := new(AuctionBidRefundAllocated)
	if err := _Auction.contract.UnpackLog(event, "BidRefundAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionBidRemovedIterator is returned from FilterBidRemoved and is used to iterate over the raw logs and unpacked data for BidRemoved events raised by the Auction contract.
type AuctionBidRemovedIterator struct {
	Event *AuctionBidRemoved // Event containing the contract specifics and raw log

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
func (it *AuctionBidRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidRemoved)
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
		it.Event = new(AuctionBidRemoved)
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
func (it *AuctionBidRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidRemoved represents a BidRemoved event raised by the Auction contract.
type AuctionBidRemoved struct {
	BidIndex         *big.Int
	Bidder           common.Address
	BuyReserveAmount *big.Int
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBidRemoved is a free log retrieval operation binding the contract event 0xc93523f51b5bdae47e5ae3ad4c018c4b8314a458265d3840bbe3ce8ca194e5d1.
//
// Solidity: event BidRemoved(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) FilterBidRemoved(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*AuctionBidRemovedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidRemoved", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidRemovedIterator{contract: _Auction.contract, event: "BidRemoved", logs: logs, sub: sub}, nil
}

// WatchBidRemoved is a free log subscription operation binding the contract event 0xc93523f51b5bdae47e5ae3ad4c018c4b8314a458265d3840bbe3ce8ca194e5d1.
//
// Solidity: event BidRemoved(uint256 indexed bidIndex, address indexed bidder, uint256 buyReserveAmount, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) WatchBidRemoved(opts *bind.WatchOpts, sink chan<- *AuctionBidRemoved, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidRemoved", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidRemoved)
				if err := _Auction.contract.UnpackLog(event, "BidRemoved", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseBidRemoved(log types.Log) (*AuctionBidRemoved, error) {
	event := new(AuctionBidRemoved)
	if err := _Auction.contract.UnpackLog(event, "BidRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionFailedAuctionBidRefundClaimedIterator is returned from FilterFailedAuctionBidRefundClaimed and is used to iterate over the raw logs and unpacked data for FailedAuctionBidRefundClaimed events raised by the Auction contract.
type AuctionFailedAuctionBidRefundClaimedIterator struct {
	Event *AuctionFailedAuctionBidRefundClaimed // Event containing the contract specifics and raw log

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
func (it *AuctionFailedAuctionBidRefundClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionFailedAuctionBidRefundClaimed)
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
		it.Event = new(AuctionFailedAuctionBidRefundClaimed)
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
func (it *AuctionFailedAuctionBidRefundClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionFailedAuctionBidRefundClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionFailedAuctionBidRefundClaimed represents a FailedAuctionBidRefundClaimed event raised by the Auction contract.
type AuctionFailedAuctionBidRefundClaimed struct {
	BidIndex         *big.Int
	Bidder           common.Address
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFailedAuctionBidRefundClaimed is a free log retrieval operation binding the contract event 0xe53799b6cbe78f2d3592976e014800e6a394c793001be11eb8faf86f68a5de4f.
//
// Solidity: event FailedAuctionBidRefundClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) FilterFailedAuctionBidRefundClaimed(opts *bind.FilterOpts, bidIndex []*big.Int, bidder []common.Address) (*AuctionFailedAuctionBidRefundClaimedIterator, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "FailedAuctionBidRefundClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionFailedAuctionBidRefundClaimedIterator{contract: _Auction.contract, event: "FailedAuctionBidRefundClaimed", logs: logs, sub: sub}, nil
}

// WatchFailedAuctionBidRefundClaimed is a free log subscription operation binding the contract event 0xe53799b6cbe78f2d3592976e014800e6a394c793001be11eb8faf86f68a5de4f.
//
// Solidity: event FailedAuctionBidRefundClaimed(uint256 indexed bidIndex, address indexed bidder, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) WatchFailedAuctionBidRefundClaimed(opts *bind.WatchOpts, sink chan<- *AuctionFailedAuctionBidRefundClaimed, bidIndex []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var bidIndexRule []interface{}
	for _, bidIndexItem := range bidIndex {
		bidIndexRule = append(bidIndexRule, bidIndexItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "FailedAuctionBidRefundClaimed", bidIndexRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionFailedAuctionBidRefundClaimed)
				if err := _Auction.contract.UnpackLog(event, "FailedAuctionBidRefundClaimed", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseFailedAuctionBidRefundClaimed(log types.Log) (*AuctionFailedAuctionBidRefundClaimed, error) {
	event := new(AuctionFailedAuctionBidRefundClaimed)
	if err := _Auction.contract.UnpackLog(event, "FailedAuctionBidRefundClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Auction contract.
type AuctionInitializedIterator struct {
	Event *AuctionInitialized // Event containing the contract specifics and raw log

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
func (it *AuctionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionInitialized)
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
		it.Event = new(AuctionInitialized)
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
func (it *AuctionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionInitialized represents a Initialized event raised by the Auction contract.
type AuctionInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Auction *AuctionFilterer) FilterInitialized(opts *bind.FilterOpts) (*AuctionInitializedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuctionInitializedIterator{contract: _Auction.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Auction *AuctionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuctionInitialized) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionInitialized)
				if err := _Auction.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseInitialized(log types.Log) (*AuctionInitialized, error) {
	event := new(AuctionInitialized)
	if err := _Auction.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionLosingBidRefundClaimedIterator is returned from FilterLosingBidRefundClaimed and is used to iterate over the raw logs and unpacked data for LosingBidRefundClaimed events raised by the Auction contract.
type AuctionLosingBidRefundClaimedIterator struct {
	Event *AuctionLosingBidRefundClaimed // Event containing the contract specifics and raw log

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
func (it *AuctionLosingBidRefundClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionLosingBidRefundClaimed)
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
		it.Event = new(AuctionLosingBidRefundClaimed)
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
func (it *AuctionLosingBidRefundClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionLosingBidRefundClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionLosingBidRefundClaimed represents a LosingBidRefundClaimed event raised by the Auction contract.
type AuctionLosingBidRefundClaimed struct {
	Bidder           common.Address
	SellCouponAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLosingBidRefundClaimed is a free log retrieval operation binding the contract event 0xb11b7956ad36f0bf4cee602c022aec98ee0cec801ba8b54def5b23729eb7bfc9.
//
// Solidity: event LosingBidRefundClaimed(address indexed bidder, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) FilterLosingBidRefundClaimed(opts *bind.FilterOpts, bidder []common.Address) (*AuctionLosingBidRefundClaimedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "LosingBidRefundClaimed", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionLosingBidRefundClaimedIterator{contract: _Auction.contract, event: "LosingBidRefundClaimed", logs: logs, sub: sub}, nil
}

// WatchLosingBidRefundClaimed is a free log subscription operation binding the contract event 0xb11b7956ad36f0bf4cee602c022aec98ee0cec801ba8b54def5b23729eb7bfc9.
//
// Solidity: event LosingBidRefundClaimed(address indexed bidder, uint256 sellCouponAmount)
func (_Auction *AuctionFilterer) WatchLosingBidRefundClaimed(opts *bind.WatchOpts, sink chan<- *AuctionLosingBidRefundClaimed, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "LosingBidRefundClaimed", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionLosingBidRefundClaimed)
				if err := _Auction.contract.UnpackLog(event, "LosingBidRefundClaimed", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseLosingBidRefundClaimed(log types.Log) (*AuctionLosingBidRefundClaimed, error) {
	event := new(AuctionLosingBidRefundClaimed)
	if err := _Auction.contract.UnpackLog(event, "LosingBidRefundClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Auction contract.
type AuctionPausedIterator struct {
	Event *AuctionPaused // Event containing the contract specifics and raw log

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
func (it *AuctionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionPaused)
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
		it.Event = new(AuctionPaused)
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
func (it *AuctionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionPaused represents a Paused event raised by the Auction contract.
type AuctionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Auction *AuctionFilterer) FilterPaused(opts *bind.FilterOpts) (*AuctionPausedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AuctionPausedIterator{contract: _Auction.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Auction *AuctionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AuctionPaused) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionPaused)
				if err := _Auction.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Auction *AuctionFilterer) ParsePaused(log types.Log) (*AuctionPaused, error) {
	event := new(AuctionPaused)
	if err := _Auction.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Auction contract.
type AuctionUnpausedIterator struct {
	Event *AuctionUnpaused // Event containing the contract specifics and raw log

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
func (it *AuctionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionUnpaused)
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
		it.Event = new(AuctionUnpaused)
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
func (it *AuctionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionUnpaused represents a Unpaused event raised by the Auction contract.
type AuctionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Auction *AuctionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AuctionUnpausedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AuctionUnpausedIterator{contract: _Auction.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Auction *AuctionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AuctionUnpaused) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionUnpaused)
				if err := _Auction.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseUnpaused(log types.Log) (*AuctionUnpaused, error) {
	event := new(AuctionUnpaused)
	if err := _Auction.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Auction contract.
type AuctionUpgradedIterator struct {
	Event *AuctionUpgraded // Event containing the contract specifics and raw log

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
func (it *AuctionUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionUpgraded)
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
		it.Event = new(AuctionUpgraded)
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
func (it *AuctionUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionUpgraded represents a Upgraded event raised by the Auction contract.
type AuctionUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Auction *AuctionFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AuctionUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AuctionUpgradedIterator{contract: _Auction.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Auction *AuctionFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AuctionUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionUpgraded)
				if err := _Auction.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Auction *AuctionFilterer) ParseUpgraded(log types.Log) (*AuctionUpgraded, error) {
	event := new(AuctionUpgraded)
	if err := _Auction.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

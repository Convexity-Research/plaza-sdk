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

// PoolPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type PoolPoolInfo struct {
	Fee                *big.Int
	Reserve            *big.Int
	BondSupply         *big.Int
	LevSupply          *big.Int
	SharesPerToken     *big.Int
	CurrentPeriod      *big.Int
	LastDistribution   *big.Int
	DistributionPeriod *big.Int
	AuctionPeriod      *big.Int
	FeeBeneficiary     common.Address
}

// PlazaSdkMetaData contains all meta data concerning the PlazaSdk contract.
var PlazaSdkMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBondToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimFees\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"couponToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distribute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeBeneficiary\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreateAmount\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"levSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolReserves\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oracleDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getOracleDecimals\",\"inputs\":[{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOraclePrice\",\"inputs\":[{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structPool.PoolInfo\",\"components\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserve\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"levSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastDistribution\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"auctionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedeemAmount\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"levSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolReserves\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oracleDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"marketRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_poolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_reserveToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_couponToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracleFeeds\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauseOnCreation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractLeverageToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastAuctionStart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracleFeeds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPoolFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reserveToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAuctionPeriod\",\"inputs\":[{\"name\":\"_auctionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDistributionPeriod\",\"inputs\":[{\"name\":\"_distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFee\",\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeBeneficiary\",\"inputs\":[{\"name\":\"_feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolSaleLimit\",\"inputs\":[{\"name\":\"_poolSaleLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSharesPerToken\",\"inputs\":[{\"name\":\"_sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"simulateCreate\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"simulateRedeem\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startAuction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferReserveToAuction\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zeroLastSharesPerToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuctionPeriodChanged\",\"inputs\":[{\"name\":\"oldPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionStarted\",\"inputs\":[{\"name\":\"auction\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"couponAmountToDistribute\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Distributed\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"distributor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionPeriodChanged\",\"inputs\":[{\"name\":\"oldPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRollOver\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeChanged\",\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeClaimed\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NoFeesToClaim\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolSaleLimitChanged\",\"inputs\":[{\"name\":\"oldThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SharesPerTokenChanged\",\"inputs\":[{\"name\":\"oldSharesPerToken\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sharesPerToken\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensRedeemed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"redeemedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessDenied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionAlreadyStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionIsOngoing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionPeriodPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionRecentlyStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotAuction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DistributionPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DistributionPeriodNotPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoFeedFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPriceFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotBeneficiary\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolSaleLimitTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransactionTooOld\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroDebtSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLeverageSupply\",\"inputs\":[]}]",
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

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) ETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "ETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_PlazaSdk *PlazaSdkSession) ETH() (common.Address, error) {
	return _PlazaSdk.Contract.ETH(&_PlazaSdk.CallOpts)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) ETH() (common.Address, error) {
	return _PlazaSdk.Contract.ETH(&_PlazaSdk.CallOpts)
}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) USD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "USD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_PlazaSdk *PlazaSdkSession) USD() (common.Address, error) {
	return _PlazaSdk.Contract.USD(&_PlazaSdk.CallOpts)
}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) USD() (common.Address, error) {
	return _PlazaSdk.Contract.USD(&_PlazaSdk.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address)
func (_PlazaSdk *PlazaSdkCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "auctions", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address)
func (_PlazaSdk *PlazaSdkSession) Auctions(arg0 *big.Int) (common.Address, error) {
	return _PlazaSdk.Contract.Auctions(&_PlazaSdk.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) Auctions(arg0 *big.Int) (common.Address, error) {
	return _PlazaSdk.Contract.Auctions(&_PlazaSdk.CallOpts, arg0)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) BondToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "bondToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_PlazaSdk *PlazaSdkSession) BondToken() (common.Address, error) {
	return _PlazaSdk.Contract.BondToken(&_PlazaSdk.CallOpts)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) BondToken() (common.Address, error) {
	return _PlazaSdk.Contract.BondToken(&_PlazaSdk.CallOpts)
}

// CouponToken is a free data retrieval call binding the contract method 0x457cf77a.
//
// Solidity: function couponToken() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) CouponToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "couponToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CouponToken is a free data retrieval call binding the contract method 0x457cf77a.
//
// Solidity: function couponToken() view returns(address)
func (_PlazaSdk *PlazaSdkSession) CouponToken() (common.Address, error) {
	return _PlazaSdk.Contract.CouponToken(&_PlazaSdk.CallOpts)
}

// CouponToken is a free data retrieval call binding the contract method 0x457cf77a.
//
// Solidity: function couponToken() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) CouponToken() (common.Address, error) {
	return _PlazaSdk.Contract.CouponToken(&_PlazaSdk.CallOpts)
}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) FeeBeneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "feeBeneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_PlazaSdk *PlazaSdkSession) FeeBeneficiary() (common.Address, error) {
	return _PlazaSdk.Contract.FeeBeneficiary(&_PlazaSdk.CallOpts)
}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) FeeBeneficiary() (common.Address, error) {
	return _PlazaSdk.Contract.FeeBeneficiary(&_PlazaSdk.CallOpts)
}

// GetCreateAmount is a free data retrieval call binding the contract method 0x9fef3e62.
//
// Solidity: function getCreateAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals) pure returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) GetCreateAmount(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "getCreateAmount", tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreateAmount is a free data retrieval call binding the contract method 0x9fef3e62.
//
// Solidity: function getCreateAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals) pure returns(uint256)
func (_PlazaSdk *PlazaSdkSession) GetCreateAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8) (*big.Int, error) {
	return _PlazaSdk.Contract.GetCreateAmount(&_PlazaSdk.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals)
}

// GetCreateAmount is a free data retrieval call binding the contract method 0x9fef3e62.
//
// Solidity: function getCreateAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals) pure returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) GetCreateAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8) (*big.Int, error) {
	return _PlazaSdk.Contract.GetCreateAmount(&_PlazaSdk.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_PlazaSdk *PlazaSdkCaller) GetOracleDecimals(opts *bind.CallOpts, quote common.Address, base common.Address) (uint8, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "getOracleDecimals", quote, base)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_PlazaSdk *PlazaSdkSession) GetOracleDecimals(quote common.Address, base common.Address) (uint8, error) {
	return _PlazaSdk.Contract.GetOracleDecimals(&_PlazaSdk.CallOpts, quote, base)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_PlazaSdk *PlazaSdkCallerSession) GetOracleDecimals(quote common.Address, base common.Address) (uint8, error) {
	return _PlazaSdk.Contract.GetOracleDecimals(&_PlazaSdk.CallOpts, quote, base)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) GetOraclePrice(opts *bind.CallOpts, quote common.Address, base common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "getOraclePrice", quote, base)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) GetOraclePrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _PlazaSdk.Contract.GetOraclePrice(&_PlazaSdk.CallOpts, quote, base)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) GetOraclePrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _PlazaSdk.Contract.GetOraclePrice(&_PlazaSdk.CallOpts, quote, base)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address) info)
func (_PlazaSdk *PlazaSdkCaller) GetPoolInfo(opts *bind.CallOpts) (PoolPoolInfo, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "getPoolInfo")

	if err != nil {
		return *new(PoolPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolPoolInfo)).(*PoolPoolInfo)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address) info)
func (_PlazaSdk *PlazaSdkSession) GetPoolInfo() (PoolPoolInfo, error) {
	return _PlazaSdk.Contract.GetPoolInfo(&_PlazaSdk.CallOpts)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address) info)
func (_PlazaSdk *PlazaSdkCallerSession) GetPoolInfo() (PoolPoolInfo, error) {
	return _PlazaSdk.Contract.GetPoolInfo(&_PlazaSdk.CallOpts)
}

// GetRedeemAmount is a free data retrieval call binding the contract method 0x98c76b23.
//
// Solidity: function getRedeemAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals, uint256 marketRate) pure returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) GetRedeemAmount(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8, marketRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "getRedeemAmount", tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals, marketRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedeemAmount is a free data retrieval call binding the contract method 0x98c76b23.
//
// Solidity: function getRedeemAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals, uint256 marketRate) pure returns(uint256)
func (_PlazaSdk *PlazaSdkSession) GetRedeemAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8, marketRate *big.Int) (*big.Int, error) {
	return _PlazaSdk.Contract.GetRedeemAmount(&_PlazaSdk.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals, marketRate)
}

// GetRedeemAmount is a free data retrieval call binding the contract method 0x98c76b23.
//
// Solidity: function getRedeemAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals, uint256 marketRate) pure returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) GetRedeemAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8, marketRate *big.Int) (*big.Int, error) {
	return _PlazaSdk.Contract.GetRedeemAmount(&_PlazaSdk.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals, marketRate)
}

// LToken is a free data retrieval call binding the contract method 0x010ee184.
//
// Solidity: function lToken() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) LToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "lToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LToken is a free data retrieval call binding the contract method 0x010ee184.
//
// Solidity: function lToken() view returns(address)
func (_PlazaSdk *PlazaSdkSession) LToken() (common.Address, error) {
	return _PlazaSdk.Contract.LToken(&_PlazaSdk.CallOpts)
}

// LToken is a free data retrieval call binding the contract method 0x010ee184.
//
// Solidity: function lToken() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) LToken() (common.Address, error) {
	return _PlazaSdk.Contract.LToken(&_PlazaSdk.CallOpts)
}

// LastAuctionStart is a free data retrieval call binding the contract method 0x578277e3.
//
// Solidity: function lastAuctionStart() view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) LastAuctionStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "lastAuctionStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAuctionStart is a free data retrieval call binding the contract method 0x578277e3.
//
// Solidity: function lastAuctionStart() view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) LastAuctionStart() (*big.Int, error) {
	return _PlazaSdk.Contract.LastAuctionStart(&_PlazaSdk.CallOpts)
}

// LastAuctionStart is a free data retrieval call binding the contract method 0x578277e3.
//
// Solidity: function lastAuctionStart() view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) LastAuctionStart() (*big.Int, error) {
	return _PlazaSdk.Contract.LastAuctionStart(&_PlazaSdk.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PlazaSdk *PlazaSdkCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PlazaSdk *PlazaSdkSession) Name() (string, error) {
	return _PlazaSdk.Contract.Name(&_PlazaSdk.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PlazaSdk *PlazaSdkCallerSession) Name() (string, error) {
	return _PlazaSdk.Contract.Name(&_PlazaSdk.CallOpts)
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

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_PlazaSdk *PlazaSdkSession) PoolFactory() (common.Address, error) {
	return _PlazaSdk.Contract.PoolFactory(&_PlazaSdk.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) PoolFactory() (common.Address, error) {
	return _PlazaSdk.Contract.PoolFactory(&_PlazaSdk.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_PlazaSdk *PlazaSdkCaller) ReserveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "reserveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_PlazaSdk *PlazaSdkSession) ReserveToken() (common.Address, error) {
	return _PlazaSdk.Contract.ReserveToken(&_PlazaSdk.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_PlazaSdk *PlazaSdkCallerSession) ReserveToken() (common.Address, error) {
	return _PlazaSdk.Contract.ReserveToken(&_PlazaSdk.CallOpts)
}

// SimulateCreate is a free data retrieval call binding the contract method 0x8afb27b0.
//
// Solidity: function simulateCreate(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) SimulateCreate(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "simulateCreate", tokenType, depositAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateCreate is a free data retrieval call binding the contract method 0x8afb27b0.
//
// Solidity: function simulateCreate(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) SimulateCreate(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _PlazaSdk.Contract.SimulateCreate(&_PlazaSdk.CallOpts, tokenType, depositAmount)
}

// SimulateCreate is a free data retrieval call binding the contract method 0x8afb27b0.
//
// Solidity: function simulateCreate(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) SimulateCreate(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _PlazaSdk.Contract.SimulateCreate(&_PlazaSdk.CallOpts, tokenType, depositAmount)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0x9d64bca4.
//
// Solidity: function simulateRedeem(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_PlazaSdk *PlazaSdkCaller) SimulateRedeem(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlazaSdk.contract.Call(opts, &out, "simulateRedeem", tokenType, depositAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateRedeem is a free data retrieval call binding the contract method 0x9d64bca4.
//
// Solidity: function simulateRedeem(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_PlazaSdk *PlazaSdkSession) SimulateRedeem(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _PlazaSdk.Contract.SimulateRedeem(&_PlazaSdk.CallOpts, tokenType, depositAmount)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0x9d64bca4.
//
// Solidity: function simulateRedeem(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_PlazaSdk *PlazaSdkCallerSession) SimulateRedeem(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _PlazaSdk.Contract.SimulateRedeem(&_PlazaSdk.CallOpts, tokenType, depositAmount)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_PlazaSdk *PlazaSdkTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_PlazaSdk *PlazaSdkSession) ClaimFees() (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimFees(&_PlazaSdk.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _PlazaSdk.Contract.ClaimFees(&_PlazaSdk.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0x6e530e97.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactor) Create(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "create", tokenType, depositAmount, minAmount)
}

// Create is a paid mutator transaction binding the contract method 0x6e530e97.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkSession) Create(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Create(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount)
}

// Create is a paid mutator transaction binding the contract method 0x6e530e97.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactorSession) Create(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Create(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount)
}

// Create0 is a paid mutator transaction binding the contract method 0xf9d22289.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactor) Create0(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "create0", tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Create0 is a paid mutator transaction binding the contract method 0xf9d22289.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_PlazaSdk *PlazaSdkSession) Create0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Create0(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Create0 is a paid mutator transaction binding the contract method 0xf9d22289.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactorSession) Create0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Create0(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_PlazaSdk *PlazaSdkTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_PlazaSdk *PlazaSdkSession) Distribute() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Distribute(&_PlazaSdk.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Distribute() (*types.Transaction, error) {
	return _PlazaSdk.Contract.Distribute(&_PlazaSdk.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e50408e.
//
// Solidity: function initialize(address _poolFactory, uint256 _fee, address _reserveToken, address _dToken, address _lToken, address _couponToken, uint256 _sharesPerToken, uint256 _distributionPeriod, address _feeBeneficiary, address _oracleFeeds, bool _pauseOnCreation) returns()
func (_PlazaSdk *PlazaSdkTransactor) Initialize(opts *bind.TransactOpts, _poolFactory common.Address, _fee *big.Int, _reserveToken common.Address, _dToken common.Address, _lToken common.Address, _couponToken common.Address, _sharesPerToken *big.Int, _distributionPeriod *big.Int, _feeBeneficiary common.Address, _oracleFeeds common.Address, _pauseOnCreation bool) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "initialize", _poolFactory, _fee, _reserveToken, _dToken, _lToken, _couponToken, _sharesPerToken, _distributionPeriod, _feeBeneficiary, _oracleFeeds, _pauseOnCreation)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e50408e.
//
// Solidity: function initialize(address _poolFactory, uint256 _fee, address _reserveToken, address _dToken, address _lToken, address _couponToken, uint256 _sharesPerToken, uint256 _distributionPeriod, address _feeBeneficiary, address _oracleFeeds, bool _pauseOnCreation) returns()
func (_PlazaSdk *PlazaSdkSession) Initialize(_poolFactory common.Address, _fee *big.Int, _reserveToken common.Address, _dToken common.Address, _lToken common.Address, _couponToken common.Address, _sharesPerToken *big.Int, _distributionPeriod *big.Int, _feeBeneficiary common.Address, _oracleFeeds common.Address, _pauseOnCreation bool) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _poolFactory, _fee, _reserveToken, _dToken, _lToken, _couponToken, _sharesPerToken, _distributionPeriod, _feeBeneficiary, _oracleFeeds, _pauseOnCreation)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e50408e.
//
// Solidity: function initialize(address _poolFactory, uint256 _fee, address _reserveToken, address _dToken, address _lToken, address _couponToken, uint256 _sharesPerToken, uint256 _distributionPeriod, address _feeBeneficiary, address _oracleFeeds, bool _pauseOnCreation) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) Initialize(_poolFactory common.Address, _fee *big.Int, _reserveToken common.Address, _dToken common.Address, _lToken common.Address, _couponToken common.Address, _sharesPerToken *big.Int, _distributionPeriod *big.Int, _feeBeneficiary common.Address, _oracleFeeds common.Address, _pauseOnCreation bool) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Initialize(&_PlazaSdk.TransactOpts, _poolFactory, _fee, _reserveToken, _dToken, _lToken, _couponToken, _sharesPerToken, _distributionPeriod, _feeBeneficiary, _oracleFeeds, _pauseOnCreation)
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

// Redeem is a paid mutator transaction binding the contract method 0x724786cf.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactor) Redeem(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "redeem", tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Redeem is a paid mutator transaction binding the contract method 0x724786cf.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_PlazaSdk *PlazaSdkSession) Redeem(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Redeem(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Redeem is a paid mutator transaction binding the contract method 0x724786cf.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactorSession) Redeem(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Redeem(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xf0fae20f.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactor) Redeem0(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "redeem0", tokenType, depositAmount, minAmount)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xf0fae20f.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkSession) Redeem0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Redeem0(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xf0fae20f.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_PlazaSdk *PlazaSdkTransactorSession) Redeem0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.Redeem0(&_PlazaSdk.TransactOpts, tokenType, depositAmount, minAmount)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xc8c73e7c.
//
// Solidity: function setAuctionPeriod(uint256 _auctionPeriod) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetAuctionPeriod(opts *bind.TransactOpts, _auctionPeriod *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setAuctionPeriod", _auctionPeriod)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xc8c73e7c.
//
// Solidity: function setAuctionPeriod(uint256 _auctionPeriod) returns()
func (_PlazaSdk *PlazaSdkSession) SetAuctionPeriod(_auctionPeriod *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetAuctionPeriod(&_PlazaSdk.TransactOpts, _auctionPeriod)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xc8c73e7c.
//
// Solidity: function setAuctionPeriod(uint256 _auctionPeriod) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetAuctionPeriod(_auctionPeriod *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetAuctionPeriod(&_PlazaSdk.TransactOpts, _auctionPeriod)
}

// SetDistributionPeriod is a paid mutator transaction binding the contract method 0x7710c6d8.
//
// Solidity: function setDistributionPeriod(uint256 _distributionPeriod) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetDistributionPeriod(opts *bind.TransactOpts, _distributionPeriod *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setDistributionPeriod", _distributionPeriod)
}

// SetDistributionPeriod is a paid mutator transaction binding the contract method 0x7710c6d8.
//
// Solidity: function setDistributionPeriod(uint256 _distributionPeriod) returns()
func (_PlazaSdk *PlazaSdkSession) SetDistributionPeriod(_distributionPeriod *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetDistributionPeriod(&_PlazaSdk.TransactOpts, _distributionPeriod)
}

// SetDistributionPeriod is a paid mutator transaction binding the contract method 0x7710c6d8.
//
// Solidity: function setDistributionPeriod(uint256 _distributionPeriod) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetDistributionPeriod(_distributionPeriod *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetDistributionPeriod(&_PlazaSdk.TransactOpts, _distributionPeriod)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_PlazaSdk *PlazaSdkSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetFee(&_PlazaSdk.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetFee(&_PlazaSdk.TransactOpts, _fee)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address _feeBeneficiary) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetFeeBeneficiary(opts *bind.TransactOpts, _feeBeneficiary common.Address) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setFeeBeneficiary", _feeBeneficiary)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address _feeBeneficiary) returns()
func (_PlazaSdk *PlazaSdkSession) SetFeeBeneficiary(_feeBeneficiary common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetFeeBeneficiary(&_PlazaSdk.TransactOpts, _feeBeneficiary)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address _feeBeneficiary) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetFeeBeneficiary(_feeBeneficiary common.Address) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetFeeBeneficiary(&_PlazaSdk.TransactOpts, _feeBeneficiary)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_PlazaSdk *PlazaSdkSession) SetName(_name string) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetName(&_PlazaSdk.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetName(&_PlazaSdk.TransactOpts, _name)
}

// SetPoolSaleLimit is a paid mutator transaction binding the contract method 0xe14a81a9.
//
// Solidity: function setPoolSaleLimit(uint256 _poolSaleLimit) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetPoolSaleLimit(opts *bind.TransactOpts, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setPoolSaleLimit", _poolSaleLimit)
}

// SetPoolSaleLimit is a paid mutator transaction binding the contract method 0xe14a81a9.
//
// Solidity: function setPoolSaleLimit(uint256 _poolSaleLimit) returns()
func (_PlazaSdk *PlazaSdkSession) SetPoolSaleLimit(_poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetPoolSaleLimit(&_PlazaSdk.TransactOpts, _poolSaleLimit)
}

// SetPoolSaleLimit is a paid mutator transaction binding the contract method 0xe14a81a9.
//
// Solidity: function setPoolSaleLimit(uint256 _poolSaleLimit) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetPoolSaleLimit(_poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetPoolSaleLimit(&_PlazaSdk.TransactOpts, _poolSaleLimit)
}

// SetSharesPerToken is a paid mutator transaction binding the contract method 0x7362d149.
//
// Solidity: function setSharesPerToken(uint256 _sharesPerToken) returns()
func (_PlazaSdk *PlazaSdkTransactor) SetSharesPerToken(opts *bind.TransactOpts, _sharesPerToken *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "setSharesPerToken", _sharesPerToken)
}

// SetSharesPerToken is a paid mutator transaction binding the contract method 0x7362d149.
//
// Solidity: function setSharesPerToken(uint256 _sharesPerToken) returns()
func (_PlazaSdk *PlazaSdkSession) SetSharesPerToken(_sharesPerToken *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetSharesPerToken(&_PlazaSdk.TransactOpts, _sharesPerToken)
}

// SetSharesPerToken is a paid mutator transaction binding the contract method 0x7362d149.
//
// Solidity: function setSharesPerToken(uint256 _sharesPerToken) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) SetSharesPerToken(_sharesPerToken *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.SetSharesPerToken(&_PlazaSdk.TransactOpts, _sharesPerToken)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_PlazaSdk *PlazaSdkTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_PlazaSdk *PlazaSdkSession) StartAuction() (*types.Transaction, error) {
	return _PlazaSdk.Contract.StartAuction(&_PlazaSdk.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) StartAuction() (*types.Transaction, error) {
	return _PlazaSdk.Contract.StartAuction(&_PlazaSdk.TransactOpts)
}

// TransferReserveToAuction is a paid mutator transaction binding the contract method 0xec27344c.
//
// Solidity: function transferReserveToAuction(uint256 amount) returns()
func (_PlazaSdk *PlazaSdkTransactor) TransferReserveToAuction(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "transferReserveToAuction", amount)
}

// TransferReserveToAuction is a paid mutator transaction binding the contract method 0xec27344c.
//
// Solidity: function transferReserveToAuction(uint256 amount) returns()
func (_PlazaSdk *PlazaSdkSession) TransferReserveToAuction(amount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.TransferReserveToAuction(&_PlazaSdk.TransactOpts, amount)
}

// TransferReserveToAuction is a paid mutator transaction binding the contract method 0xec27344c.
//
// Solidity: function transferReserveToAuction(uint256 amount) returns()
func (_PlazaSdk *PlazaSdkTransactorSession) TransferReserveToAuction(amount *big.Int) (*types.Transaction, error) {
	return _PlazaSdk.Contract.TransferReserveToAuction(&_PlazaSdk.TransactOpts, amount)
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

// ZeroLastSharesPerToken is a paid mutator transaction binding the contract method 0x10c04932.
//
// Solidity: function zeroLastSharesPerToken() returns()
func (_PlazaSdk *PlazaSdkTransactor) ZeroLastSharesPerToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlazaSdk.contract.Transact(opts, "zeroLastSharesPerToken")
}

// ZeroLastSharesPerToken is a paid mutator transaction binding the contract method 0x10c04932.
//
// Solidity: function zeroLastSharesPerToken() returns()
func (_PlazaSdk *PlazaSdkSession) ZeroLastSharesPerToken() (*types.Transaction, error) {
	return _PlazaSdk.Contract.ZeroLastSharesPerToken(&_PlazaSdk.TransactOpts)
}

// ZeroLastSharesPerToken is a paid mutator transaction binding the contract method 0x10c04932.
//
// Solidity: function zeroLastSharesPerToken() returns()
func (_PlazaSdk *PlazaSdkTransactorSession) ZeroLastSharesPerToken() (*types.Transaction, error) {
	return _PlazaSdk.Contract.ZeroLastSharesPerToken(&_PlazaSdk.TransactOpts)
}

// PlazaSdkAuctionPeriodChangedIterator is returned from FilterAuctionPeriodChanged and is used to iterate over the raw logs and unpacked data for AuctionPeriodChanged events raised by the PlazaSdk contract.
type PlazaSdkAuctionPeriodChangedIterator struct {
	Event *PlazaSdkAuctionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *PlazaSdkAuctionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkAuctionPeriodChanged)
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
		it.Event = new(PlazaSdkAuctionPeriodChanged)
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
func (it *PlazaSdkAuctionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkAuctionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkAuctionPeriodChanged represents a AuctionPeriodChanged event raised by the PlazaSdk contract.
type PlazaSdkAuctionPeriodChanged struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionPeriodChanged is a free log retrieval operation binding the contract event 0x8b63bdaf0f5981cb255f97be2272e1ea65195664462f04777a64848795236aed.
//
// Solidity: event AuctionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_PlazaSdk *PlazaSdkFilterer) FilterAuctionPeriodChanged(opts *bind.FilterOpts) (*PlazaSdkAuctionPeriodChangedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "AuctionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkAuctionPeriodChangedIterator{contract: _PlazaSdk.contract, event: "AuctionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchAuctionPeriodChanged is a free log subscription operation binding the contract event 0x8b63bdaf0f5981cb255f97be2272e1ea65195664462f04777a64848795236aed.
//
// Solidity: event AuctionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_PlazaSdk *PlazaSdkFilterer) WatchAuctionPeriodChanged(opts *bind.WatchOpts, sink chan<- *PlazaSdkAuctionPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "AuctionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkAuctionPeriodChanged)
				if err := _PlazaSdk.contract.UnpackLog(event, "AuctionPeriodChanged", log); err != nil {
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

// ParseAuctionPeriodChanged is a log parse operation binding the contract event 0x8b63bdaf0f5981cb255f97be2272e1ea65195664462f04777a64848795236aed.
//
// Solidity: event AuctionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_PlazaSdk *PlazaSdkFilterer) ParseAuctionPeriodChanged(log types.Log) (*PlazaSdkAuctionPeriodChanged, error) {
	event := new(PlazaSdkAuctionPeriodChanged)
	if err := _PlazaSdk.contract.UnpackLog(event, "AuctionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the PlazaSdk contract.
type PlazaSdkAuctionStartedIterator struct {
	Event *PlazaSdkAuctionStarted // Event containing the contract specifics and raw log

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
func (it *PlazaSdkAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkAuctionStarted)
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
		it.Event = new(PlazaSdkAuctionStarted)
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
func (it *PlazaSdkAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkAuctionStarted represents a AuctionStarted event raised by the PlazaSdk contract.
type PlazaSdkAuctionStarted struct {
	Auction                  common.Address
	Period                   *big.Int
	CouponAmountToDistribute *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x92db7303de5c78e58e4570f0b4acf392e3dac09394e96c4fa62b323b6337ee65.
//
// Solidity: event AuctionStarted(address auction, uint256 period, uint256 couponAmountToDistribute)
func (_PlazaSdk *PlazaSdkFilterer) FilterAuctionStarted(opts *bind.FilterOpts) (*PlazaSdkAuctionStartedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkAuctionStartedIterator{contract: _PlazaSdk.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x92db7303de5c78e58e4570f0b4acf392e3dac09394e96c4fa62b323b6337ee65.
//
// Solidity: event AuctionStarted(address auction, uint256 period, uint256 couponAmountToDistribute)
func (_PlazaSdk *PlazaSdkFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *PlazaSdkAuctionStarted) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkAuctionStarted)
				if err := _PlazaSdk.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0x92db7303de5c78e58e4570f0b4acf392e3dac09394e96c4fa62b323b6337ee65.
//
// Solidity: event AuctionStarted(address auction, uint256 period, uint256 couponAmountToDistribute)
func (_PlazaSdk *PlazaSdkFilterer) ParseAuctionStarted(log types.Log) (*PlazaSdkAuctionStarted, error) {
	event := new(PlazaSdkAuctionStarted)
	if err := _PlazaSdk.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkDistributedIterator is returned from FilterDistributed and is used to iterate over the raw logs and unpacked data for Distributed events raised by the PlazaSdk contract.
type PlazaSdkDistributedIterator struct {
	Event *PlazaSdkDistributed // Event containing the contract specifics and raw log

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
func (it *PlazaSdkDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkDistributed)
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
		it.Event = new(PlazaSdkDistributed)
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
func (it *PlazaSdkDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkDistributed represents a Distributed event raised by the PlazaSdk contract.
type PlazaSdkDistributed struct {
	Period      *big.Int
	Amount      *big.Int
	Distributor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDistributed is a free log retrieval operation binding the contract event 0x8dda631b265286fb6a61a6658e6e75969e952f2787646e616dfc15a84f298c18.
//
// Solidity: event Distributed(uint256 period, uint256 amount, address distributor)
func (_PlazaSdk *PlazaSdkFilterer) FilterDistributed(opts *bind.FilterOpts) (*PlazaSdkDistributedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "Distributed")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkDistributedIterator{contract: _PlazaSdk.contract, event: "Distributed", logs: logs, sub: sub}, nil
}

// WatchDistributed is a free log subscription operation binding the contract event 0x8dda631b265286fb6a61a6658e6e75969e952f2787646e616dfc15a84f298c18.
//
// Solidity: event Distributed(uint256 period, uint256 amount, address distributor)
func (_PlazaSdk *PlazaSdkFilterer) WatchDistributed(opts *bind.WatchOpts, sink chan<- *PlazaSdkDistributed) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "Distributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkDistributed)
				if err := _PlazaSdk.contract.UnpackLog(event, "Distributed", log); err != nil {
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

// ParseDistributed is a log parse operation binding the contract event 0x8dda631b265286fb6a61a6658e6e75969e952f2787646e616dfc15a84f298c18.
//
// Solidity: event Distributed(uint256 period, uint256 amount, address distributor)
func (_PlazaSdk *PlazaSdkFilterer) ParseDistributed(log types.Log) (*PlazaSdkDistributed, error) {
	event := new(PlazaSdkDistributed)
	if err := _PlazaSdk.contract.UnpackLog(event, "Distributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkDistributionPeriodChangedIterator is returned from FilterDistributionPeriodChanged and is used to iterate over the raw logs and unpacked data for DistributionPeriodChanged events raised by the PlazaSdk contract.
type PlazaSdkDistributionPeriodChangedIterator struct {
	Event *PlazaSdkDistributionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *PlazaSdkDistributionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkDistributionPeriodChanged)
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
		it.Event = new(PlazaSdkDistributionPeriodChanged)
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
func (it *PlazaSdkDistributionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkDistributionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkDistributionPeriodChanged represents a DistributionPeriodChanged event raised by the PlazaSdk contract.
type PlazaSdkDistributionPeriodChanged struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributionPeriodChanged is a free log retrieval operation binding the contract event 0x614ada6efad9c4a86e8541da67aa075e9e21e37f888c79f23807667abb742ce3.
//
// Solidity: event DistributionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_PlazaSdk *PlazaSdkFilterer) FilterDistributionPeriodChanged(opts *bind.FilterOpts) (*PlazaSdkDistributionPeriodChangedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "DistributionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkDistributionPeriodChangedIterator{contract: _PlazaSdk.contract, event: "DistributionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchDistributionPeriodChanged is a free log subscription operation binding the contract event 0x614ada6efad9c4a86e8541da67aa075e9e21e37f888c79f23807667abb742ce3.
//
// Solidity: event DistributionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_PlazaSdk *PlazaSdkFilterer) WatchDistributionPeriodChanged(opts *bind.WatchOpts, sink chan<- *PlazaSdkDistributionPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "DistributionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkDistributionPeriodChanged)
				if err := _PlazaSdk.contract.UnpackLog(event, "DistributionPeriodChanged", log); err != nil {
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

// ParseDistributionPeriodChanged is a log parse operation binding the contract event 0x614ada6efad9c4a86e8541da67aa075e9e21e37f888c79f23807667abb742ce3.
//
// Solidity: event DistributionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_PlazaSdk *PlazaSdkFilterer) ParseDistributionPeriodChanged(log types.Log) (*PlazaSdkDistributionPeriodChanged, error) {
	event := new(PlazaSdkDistributionPeriodChanged)
	if err := _PlazaSdk.contract.UnpackLog(event, "DistributionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkDistributionRollOverIterator is returned from FilterDistributionRollOver and is used to iterate over the raw logs and unpacked data for DistributionRollOver events raised by the PlazaSdk contract.
type PlazaSdkDistributionRollOverIterator struct {
	Event *PlazaSdkDistributionRollOver // Event containing the contract specifics and raw log

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
func (it *PlazaSdkDistributionRollOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkDistributionRollOver)
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
		it.Event = new(PlazaSdkDistributionRollOver)
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
func (it *PlazaSdkDistributionRollOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkDistributionRollOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkDistributionRollOver represents a DistributionRollOver event raised by the PlazaSdk contract.
type PlazaSdkDistributionRollOver struct {
	Period *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributionRollOver is a free log retrieval operation binding the contract event 0xb07ac059d2a0dd966a73cbd0cb9b6c1b7af9468cab03dc1fd5d34c6fb4a6b30c.
//
// Solidity: event DistributionRollOver(uint256 period, uint256 shares)
func (_PlazaSdk *PlazaSdkFilterer) FilterDistributionRollOver(opts *bind.FilterOpts) (*PlazaSdkDistributionRollOverIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "DistributionRollOver")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkDistributionRollOverIterator{contract: _PlazaSdk.contract, event: "DistributionRollOver", logs: logs, sub: sub}, nil
}

// WatchDistributionRollOver is a free log subscription operation binding the contract event 0xb07ac059d2a0dd966a73cbd0cb9b6c1b7af9468cab03dc1fd5d34c6fb4a6b30c.
//
// Solidity: event DistributionRollOver(uint256 period, uint256 shares)
func (_PlazaSdk *PlazaSdkFilterer) WatchDistributionRollOver(opts *bind.WatchOpts, sink chan<- *PlazaSdkDistributionRollOver) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "DistributionRollOver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkDistributionRollOver)
				if err := _PlazaSdk.contract.UnpackLog(event, "DistributionRollOver", log); err != nil {
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

// ParseDistributionRollOver is a log parse operation binding the contract event 0xb07ac059d2a0dd966a73cbd0cb9b6c1b7af9468cab03dc1fd5d34c6fb4a6b30c.
//
// Solidity: event DistributionRollOver(uint256 period, uint256 shares)
func (_PlazaSdk *PlazaSdkFilterer) ParseDistributionRollOver(log types.Log) (*PlazaSdkDistributionRollOver, error) {
	event := new(PlazaSdkDistributionRollOver)
	if err := _PlazaSdk.contract.UnpackLog(event, "DistributionRollOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkFeeChangedIterator is returned from FilterFeeChanged and is used to iterate over the raw logs and unpacked data for FeeChanged events raised by the PlazaSdk contract.
type PlazaSdkFeeChangedIterator struct {
	Event *PlazaSdkFeeChanged // Event containing the contract specifics and raw log

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
func (it *PlazaSdkFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkFeeChanged)
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
		it.Event = new(PlazaSdkFeeChanged)
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
func (it *PlazaSdkFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkFeeChanged represents a FeeChanged event raised by the PlazaSdk contract.
type PlazaSdkFeeChanged struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeChanged is a free log retrieval operation binding the contract event 0x5fc463da23c1b063e66f9e352006a7fbe8db7223c455dc429e881a2dfe2f94f1.
//
// Solidity: event FeeChanged(uint256 oldFee, uint256 newFee)
func (_PlazaSdk *PlazaSdkFilterer) FilterFeeChanged(opts *bind.FilterOpts) (*PlazaSdkFeeChangedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkFeeChangedIterator{contract: _PlazaSdk.contract, event: "FeeChanged", logs: logs, sub: sub}, nil
}

// WatchFeeChanged is a free log subscription operation binding the contract event 0x5fc463da23c1b063e66f9e352006a7fbe8db7223c455dc429e881a2dfe2f94f1.
//
// Solidity: event FeeChanged(uint256 oldFee, uint256 newFee)
func (_PlazaSdk *PlazaSdkFilterer) WatchFeeChanged(opts *bind.WatchOpts, sink chan<- *PlazaSdkFeeChanged) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkFeeChanged)
				if err := _PlazaSdk.contract.UnpackLog(event, "FeeChanged", log); err != nil {
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

// ParseFeeChanged is a log parse operation binding the contract event 0x5fc463da23c1b063e66f9e352006a7fbe8db7223c455dc429e881a2dfe2f94f1.
//
// Solidity: event FeeChanged(uint256 oldFee, uint256 newFee)
func (_PlazaSdk *PlazaSdkFilterer) ParseFeeChanged(log types.Log) (*PlazaSdkFeeChanged, error) {
	event := new(PlazaSdkFeeChanged)
	if err := _PlazaSdk.contract.UnpackLog(event, "FeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkFeeClaimedIterator is returned from FilterFeeClaimed and is used to iterate over the raw logs and unpacked data for FeeClaimed events raised by the PlazaSdk contract.
type PlazaSdkFeeClaimedIterator struct {
	Event *PlazaSdkFeeClaimed // Event containing the contract specifics and raw log

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
func (it *PlazaSdkFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkFeeClaimed)
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
		it.Event = new(PlazaSdkFeeClaimed)
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
func (it *PlazaSdkFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkFeeClaimed represents a FeeClaimed event raised by the PlazaSdk contract.
type PlazaSdkFeeClaimed struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeClaimed is a free log retrieval operation binding the contract event 0x20ca5094f3a20c321cbe4123d0f01b276b81df0fa24cd4d83d9253956035d863.
//
// Solidity: event FeeClaimed(address beneficiary, uint256 amount)
func (_PlazaSdk *PlazaSdkFilterer) FilterFeeClaimed(opts *bind.FilterOpts) (*PlazaSdkFeeClaimedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "FeeClaimed")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkFeeClaimedIterator{contract: _PlazaSdk.contract, event: "FeeClaimed", logs: logs, sub: sub}, nil
}

// WatchFeeClaimed is a free log subscription operation binding the contract event 0x20ca5094f3a20c321cbe4123d0f01b276b81df0fa24cd4d83d9253956035d863.
//
// Solidity: event FeeClaimed(address beneficiary, uint256 amount)
func (_PlazaSdk *PlazaSdkFilterer) WatchFeeClaimed(opts *bind.WatchOpts, sink chan<- *PlazaSdkFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "FeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkFeeClaimed)
				if err := _PlazaSdk.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
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

// ParseFeeClaimed is a log parse operation binding the contract event 0x20ca5094f3a20c321cbe4123d0f01b276b81df0fa24cd4d83d9253956035d863.
//
// Solidity: event FeeClaimed(address beneficiary, uint256 amount)
func (_PlazaSdk *PlazaSdkFilterer) ParseFeeClaimed(log types.Log) (*PlazaSdkFeeClaimed, error) {
	event := new(PlazaSdkFeeClaimed)
	if err := _PlazaSdk.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
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

// PlazaSdkNoFeesToClaimIterator is returned from FilterNoFeesToClaim and is used to iterate over the raw logs and unpacked data for NoFeesToClaim events raised by the PlazaSdk contract.
type PlazaSdkNoFeesToClaimIterator struct {
	Event *PlazaSdkNoFeesToClaim // Event containing the contract specifics and raw log

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
func (it *PlazaSdkNoFeesToClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkNoFeesToClaim)
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
		it.Event = new(PlazaSdkNoFeesToClaim)
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
func (it *PlazaSdkNoFeesToClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkNoFeesToClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkNoFeesToClaim represents a NoFeesToClaim event raised by the PlazaSdk contract.
type PlazaSdkNoFeesToClaim struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNoFeesToClaim is a free log retrieval operation binding the contract event 0x846d8c5c020eac85233c31fc79d004ef12b5e20abfbe63b23c49f764393b4482.
//
// Solidity: event NoFeesToClaim()
func (_PlazaSdk *PlazaSdkFilterer) FilterNoFeesToClaim(opts *bind.FilterOpts) (*PlazaSdkNoFeesToClaimIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "NoFeesToClaim")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkNoFeesToClaimIterator{contract: _PlazaSdk.contract, event: "NoFeesToClaim", logs: logs, sub: sub}, nil
}

// WatchNoFeesToClaim is a free log subscription operation binding the contract event 0x846d8c5c020eac85233c31fc79d004ef12b5e20abfbe63b23c49f764393b4482.
//
// Solidity: event NoFeesToClaim()
func (_PlazaSdk *PlazaSdkFilterer) WatchNoFeesToClaim(opts *bind.WatchOpts, sink chan<- *PlazaSdkNoFeesToClaim) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "NoFeesToClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkNoFeesToClaim)
				if err := _PlazaSdk.contract.UnpackLog(event, "NoFeesToClaim", log); err != nil {
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

// ParseNoFeesToClaim is a log parse operation binding the contract event 0x846d8c5c020eac85233c31fc79d004ef12b5e20abfbe63b23c49f764393b4482.
//
// Solidity: event NoFeesToClaim()
func (_PlazaSdk *PlazaSdkFilterer) ParseNoFeesToClaim(log types.Log) (*PlazaSdkNoFeesToClaim, error) {
	event := new(PlazaSdkNoFeesToClaim)
	if err := _PlazaSdk.contract.UnpackLog(event, "NoFeesToClaim", log); err != nil {
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

// PlazaSdkPoolSaleLimitChangedIterator is returned from FilterPoolSaleLimitChanged and is used to iterate over the raw logs and unpacked data for PoolSaleLimitChanged events raised by the PlazaSdk contract.
type PlazaSdkPoolSaleLimitChangedIterator struct {
	Event *PlazaSdkPoolSaleLimitChanged // Event containing the contract specifics and raw log

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
func (it *PlazaSdkPoolSaleLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkPoolSaleLimitChanged)
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
		it.Event = new(PlazaSdkPoolSaleLimitChanged)
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
func (it *PlazaSdkPoolSaleLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkPoolSaleLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkPoolSaleLimitChanged represents a PoolSaleLimitChanged event raised by the PlazaSdk contract.
type PlazaSdkPoolSaleLimitChanged struct {
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolSaleLimitChanged is a free log retrieval operation binding the contract event 0xbd1900a27c42d19889a9199e391bdf6738af2b388f81b7d73db6013548a27e7d.
//
// Solidity: event PoolSaleLimitChanged(uint256 oldThreshold, uint256 newThreshold)
func (_PlazaSdk *PlazaSdkFilterer) FilterPoolSaleLimitChanged(opts *bind.FilterOpts) (*PlazaSdkPoolSaleLimitChangedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "PoolSaleLimitChanged")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkPoolSaleLimitChangedIterator{contract: _PlazaSdk.contract, event: "PoolSaleLimitChanged", logs: logs, sub: sub}, nil
}

// WatchPoolSaleLimitChanged is a free log subscription operation binding the contract event 0xbd1900a27c42d19889a9199e391bdf6738af2b388f81b7d73db6013548a27e7d.
//
// Solidity: event PoolSaleLimitChanged(uint256 oldThreshold, uint256 newThreshold)
func (_PlazaSdk *PlazaSdkFilterer) WatchPoolSaleLimitChanged(opts *bind.WatchOpts, sink chan<- *PlazaSdkPoolSaleLimitChanged) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "PoolSaleLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkPoolSaleLimitChanged)
				if err := _PlazaSdk.contract.UnpackLog(event, "PoolSaleLimitChanged", log); err != nil {
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

// ParsePoolSaleLimitChanged is a log parse operation binding the contract event 0xbd1900a27c42d19889a9199e391bdf6738af2b388f81b7d73db6013548a27e7d.
//
// Solidity: event PoolSaleLimitChanged(uint256 oldThreshold, uint256 newThreshold)
func (_PlazaSdk *PlazaSdkFilterer) ParsePoolSaleLimitChanged(log types.Log) (*PlazaSdkPoolSaleLimitChanged, error) {
	event := new(PlazaSdkPoolSaleLimitChanged)
	if err := _PlazaSdk.contract.UnpackLog(event, "PoolSaleLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkSharesPerTokenChangedIterator is returned from FilterSharesPerTokenChanged and is used to iterate over the raw logs and unpacked data for SharesPerTokenChanged events raised by the PlazaSdk contract.
type PlazaSdkSharesPerTokenChangedIterator struct {
	Event *PlazaSdkSharesPerTokenChanged // Event containing the contract specifics and raw log

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
func (it *PlazaSdkSharesPerTokenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkSharesPerTokenChanged)
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
		it.Event = new(PlazaSdkSharesPerTokenChanged)
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
func (it *PlazaSdkSharesPerTokenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkSharesPerTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkSharesPerTokenChanged represents a SharesPerTokenChanged event raised by the PlazaSdk contract.
type PlazaSdkSharesPerTokenChanged struct {
	OldSharesPerToken *big.Int
	SharesPerToken    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSharesPerTokenChanged is a free log retrieval operation binding the contract event 0x67b8059073c48bc22a124c60cf344bd97cf9b265f273fa09afe215b64eb5d2c2.
//
// Solidity: event SharesPerTokenChanged(uint256 oldSharesPerToken, uint256 sharesPerToken)
func (_PlazaSdk *PlazaSdkFilterer) FilterSharesPerTokenChanged(opts *bind.FilterOpts) (*PlazaSdkSharesPerTokenChangedIterator, error) {

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "SharesPerTokenChanged")
	if err != nil {
		return nil, err
	}
	return &PlazaSdkSharesPerTokenChangedIterator{contract: _PlazaSdk.contract, event: "SharesPerTokenChanged", logs: logs, sub: sub}, nil
}

// WatchSharesPerTokenChanged is a free log subscription operation binding the contract event 0x67b8059073c48bc22a124c60cf344bd97cf9b265f273fa09afe215b64eb5d2c2.
//
// Solidity: event SharesPerTokenChanged(uint256 oldSharesPerToken, uint256 sharesPerToken)
func (_PlazaSdk *PlazaSdkFilterer) WatchSharesPerTokenChanged(opts *bind.WatchOpts, sink chan<- *PlazaSdkSharesPerTokenChanged) (event.Subscription, error) {

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "SharesPerTokenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkSharesPerTokenChanged)
				if err := _PlazaSdk.contract.UnpackLog(event, "SharesPerTokenChanged", log); err != nil {
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

// ParseSharesPerTokenChanged is a log parse operation binding the contract event 0x67b8059073c48bc22a124c60cf344bd97cf9b265f273fa09afe215b64eb5d2c2.
//
// Solidity: event SharesPerTokenChanged(uint256 oldSharesPerToken, uint256 sharesPerToken)
func (_PlazaSdk *PlazaSdkFilterer) ParseSharesPerTokenChanged(log types.Log) (*PlazaSdkSharesPerTokenChanged, error) {
	event := new(PlazaSdkSharesPerTokenChanged)
	if err := _PlazaSdk.contract.UnpackLog(event, "SharesPerTokenChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlazaSdkTokensCreatedIterator is returned from FilterTokensCreated and is used to iterate over the raw logs and unpacked data for TokensCreated events raised by the PlazaSdk contract.
type PlazaSdkTokensCreatedIterator struct {
	Event *PlazaSdkTokensCreated // Event containing the contract specifics and raw log

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
func (it *PlazaSdkTokensCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlazaSdkTokensCreated)
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
		it.Event = new(PlazaSdkTokensCreated)
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
func (it *PlazaSdkTokensCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlazaSdkTokensCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlazaSdkTokensCreated represents a TokensCreated event raised by the PlazaSdk contract.
type PlazaSdkTokensCreated struct {
	Caller          common.Address
	OnBehalfOf      common.Address
	TokenType       uint8
	DepositedAmount *big.Int
	MintedAmount    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokensCreated is a free log retrieval operation binding the contract event 0xbd52f81ce4beba21c7c6a29213c7c718167799f216ed3ebe0f7224c3947fa525.
//
// Solidity: event TokensCreated(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 mintedAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterTokensCreated(opts *bind.FilterOpts, onBehalfOf []common.Address) (*PlazaSdkTokensCreatedIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "TokensCreated", onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkTokensCreatedIterator{contract: _PlazaSdk.contract, event: "TokensCreated", logs: logs, sub: sub}, nil
}

// WatchTokensCreated is a free log subscription operation binding the contract event 0xbd52f81ce4beba21c7c6a29213c7c718167799f216ed3ebe0f7224c3947fa525.
//
// Solidity: event TokensCreated(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 mintedAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchTokensCreated(opts *bind.WatchOpts, sink chan<- *PlazaSdkTokensCreated, onBehalfOf []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "TokensCreated", onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlazaSdkTokensCreated)
				if err := _PlazaSdk.contract.UnpackLog(event, "TokensCreated", log); err != nil {
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

// ParseTokensCreated is a log parse operation binding the contract event 0xbd52f81ce4beba21c7c6a29213c7c718167799f216ed3ebe0f7224c3947fa525.
//
// Solidity: event TokensCreated(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 mintedAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseTokensCreated(log types.Log) (*PlazaSdkTokensCreated, error) {
	event := new(PlazaSdkTokensCreated)
	if err := _PlazaSdk.contract.UnpackLog(event, "TokensCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Caller          common.Address
	OnBehalfOf      common.Address
	TokenType       uint8
	DepositedAmount *big.Int
	RedeemedAmount  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokensRedeemed is a free log retrieval operation binding the contract event 0x2441fb9f349bdfdf06755d1f5239efc1851c81c59e35378f3df9ac880bca3094.
//
// Solidity: event TokensRedeemed(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_PlazaSdk *PlazaSdkFilterer) FilterTokensRedeemed(opts *bind.FilterOpts, onBehalfOf []common.Address) (*PlazaSdkTokensRedeemedIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PlazaSdk.contract.FilterLogs(opts, "TokensRedeemed", onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &PlazaSdkTokensRedeemedIterator{contract: _PlazaSdk.contract, event: "TokensRedeemed", logs: logs, sub: sub}, nil
}

// WatchTokensRedeemed is a free log subscription operation binding the contract event 0x2441fb9f349bdfdf06755d1f5239efc1851c81c59e35378f3df9ac880bca3094.
//
// Solidity: event TokensRedeemed(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_PlazaSdk *PlazaSdkFilterer) WatchTokensRedeemed(opts *bind.WatchOpts, sink chan<- *PlazaSdkTokensRedeemed, onBehalfOf []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _PlazaSdk.contract.WatchLogs(opts, "TokensRedeemed", onBehalfOfRule)
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

// ParseTokensRedeemed is a log parse operation binding the contract event 0x2441fb9f349bdfdf06755d1f5239efc1851c81c59e35378f3df9ac880bca3094.
//
// Solidity: event TokensRedeemed(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_PlazaSdk *PlazaSdkFilterer) ParseTokensRedeemed(log types.Log) (*PlazaSdkTokensRedeemed, error) {
	event := new(PlazaSdkTokensRedeemed)
	if err := _PlazaSdk.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
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

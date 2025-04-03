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

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBondToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimFees\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"couponToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"create\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distribute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeBeneficiary\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreateAmount\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"levSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolReserves\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oracleDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getOracleDecimals\",\"inputs\":[{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOraclePrice\",\"inputs\":[{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structPool.PoolInfo\",\"components\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserve\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"levSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastDistribution\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"auctionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedeemAmount\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"levSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolReserves\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oracleDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"marketRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_poolFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_reserveToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_couponToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracleFeeds\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauseOnCreation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractLeverageToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastAuctionStart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracleFeeds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPoolFactory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reserveToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAuctionPeriod\",\"inputs\":[{\"name\":\"_auctionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDistributionPeriod\",\"inputs\":[{\"name\":\"_distributionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFee\",\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeBeneficiary\",\"inputs\":[{\"name\":\"_feeBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolSaleLimit\",\"inputs\":[{\"name\":\"_poolSaleLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSharesPerToken\",\"inputs\":[{\"name\":\"_sharesPerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"simulateCreate\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"simulateRedeem\",\"inputs\":[{\"name\":\"tokenType\",\"type\":\"uint8\",\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startAuction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferReserveToAuction\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zeroLastSharesPerToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuctionPeriodChanged\",\"inputs\":[{\"name\":\"oldPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionStarted\",\"inputs\":[{\"name\":\"auction\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"couponAmountToDistribute\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Distributed\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"distributor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionPeriodChanged\",\"inputs\":[{\"name\":\"oldPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRollOver\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeChanged\",\"inputs\":[{\"name\":\"oldFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeClaimed\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NoFeesToClaim\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolSaleLimitChanged\",\"inputs\":[{\"name\":\"oldThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SharesPerTokenChanged\",\"inputs\":[{\"name\":\"oldSharesPerToken\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sharesPerToken\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensRedeemed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumPool.TokenType\"},{\"name\":\"depositedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"redeemedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessDenied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionAlreadyStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionIsOngoing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionPeriodPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionRecentlyStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotAuction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DistributionPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DistributionPeriodNotPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MinAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoFeedFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPriceFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotBeneficiary\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolSaleLimitTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransactionTooOld\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroDebtSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLeverageSupply\",\"inputs\":[]}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_Pool *PoolCaller) ETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "ETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_Pool *PoolSession) ETH() (common.Address, error) {
	return _Pool.Contract.ETH(&_Pool.CallOpts)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_Pool *PoolCallerSession) ETH() (common.Address, error) {
	return _Pool.Contract.ETH(&_Pool.CallOpts)
}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_Pool *PoolCaller) USD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "USD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_Pool *PoolSession) USD() (common.Address, error) {
	return _Pool.Contract.USD(&_Pool.CallOpts)
}

// USD is a free data retrieval call binding the contract method 0x1bf6c21b.
//
// Solidity: function USD() view returns(address)
func (_Pool *PoolCallerSession) USD() (common.Address, error) {
	return _Pool.Contract.USD(&_Pool.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address)
func (_Pool *PoolCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "auctions", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address)
func (_Pool *PoolSession) Auctions(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.Auctions(&_Pool.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address)
func (_Pool *PoolCallerSession) Auctions(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.Auctions(&_Pool.CallOpts, arg0)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_Pool *PoolCaller) BondToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "bondToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_Pool *PoolSession) BondToken() (common.Address, error) {
	return _Pool.Contract.BondToken(&_Pool.CallOpts)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_Pool *PoolCallerSession) BondToken() (common.Address, error) {
	return _Pool.Contract.BondToken(&_Pool.CallOpts)
}

// CouponToken is a free data retrieval call binding the contract method 0x457cf77a.
//
// Solidity: function couponToken() view returns(address)
func (_Pool *PoolCaller) CouponToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "couponToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CouponToken is a free data retrieval call binding the contract method 0x457cf77a.
//
// Solidity: function couponToken() view returns(address)
func (_Pool *PoolSession) CouponToken() (common.Address, error) {
	return _Pool.Contract.CouponToken(&_Pool.CallOpts)
}

// CouponToken is a free data retrieval call binding the contract method 0x457cf77a.
//
// Solidity: function couponToken() view returns(address)
func (_Pool *PoolCallerSession) CouponToken() (common.Address, error) {
	return _Pool.Contract.CouponToken(&_Pool.CallOpts)
}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_Pool *PoolCaller) FeeBeneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "feeBeneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_Pool *PoolSession) FeeBeneficiary() (common.Address, error) {
	return _Pool.Contract.FeeBeneficiary(&_Pool.CallOpts)
}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_Pool *PoolCallerSession) FeeBeneficiary() (common.Address, error) {
	return _Pool.Contract.FeeBeneficiary(&_Pool.CallOpts)
}

// GetCreateAmount is a free data retrieval call binding the contract method 0x9fef3e62.
//
// Solidity: function getCreateAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals) pure returns(uint256)
func (_Pool *PoolCaller) GetCreateAmount(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getCreateAmount", tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreateAmount is a free data retrieval call binding the contract method 0x9fef3e62.
//
// Solidity: function getCreateAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals) pure returns(uint256)
func (_Pool *PoolSession) GetCreateAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8) (*big.Int, error) {
	return _Pool.Contract.GetCreateAmount(&_Pool.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals)
}

// GetCreateAmount is a free data retrieval call binding the contract method 0x9fef3e62.
//
// Solidity: function getCreateAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals) pure returns(uint256)
func (_Pool *PoolCallerSession) GetCreateAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8) (*big.Int, error) {
	return _Pool.Contract.GetCreateAmount(&_Pool.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_Pool *PoolCaller) GetOracleDecimals(opts *bind.CallOpts, quote common.Address, base common.Address) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getOracleDecimals", quote, base)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_Pool *PoolSession) GetOracleDecimals(quote common.Address, base common.Address) (uint8, error) {
	return _Pool.Contract.GetOracleDecimals(&_Pool.CallOpts, quote, base)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0xd7032f2d.
//
// Solidity: function getOracleDecimals(address quote, address base) view returns(uint8 decimals)
func (_Pool *PoolCallerSession) GetOracleDecimals(quote common.Address, base common.Address) (uint8, error) {
	return _Pool.Contract.GetOracleDecimals(&_Pool.CallOpts, quote, base)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_Pool *PoolCaller) GetOraclePrice(opts *bind.CallOpts, quote common.Address, base common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getOraclePrice", quote, base)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_Pool *PoolSession) GetOraclePrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _Pool.Contract.GetOraclePrice(&_Pool.CallOpts, quote, base)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x4c2d8eff.
//
// Solidity: function getOraclePrice(address quote, address base) view returns(uint256)
func (_Pool *PoolCallerSession) GetOraclePrice(quote common.Address, base common.Address) (*big.Int, error) {
	return _Pool.Contract.GetOraclePrice(&_Pool.CallOpts, quote, base)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address) info)
func (_Pool *PoolCaller) GetPoolInfo(opts *bind.CallOpts) (PoolPoolInfo, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getPoolInfo")

	if err != nil {
		return *new(PoolPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolPoolInfo)).(*PoolPoolInfo)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address) info)
func (_Pool *PoolSession) GetPoolInfo() (PoolPoolInfo, error) {
	return _Pool.Contract.GetPoolInfo(&_Pool.CallOpts)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address) info)
func (_Pool *PoolCallerSession) GetPoolInfo() (PoolPoolInfo, error) {
	return _Pool.Contract.GetPoolInfo(&_Pool.CallOpts)
}

// GetRedeemAmount is a free data retrieval call binding the contract method 0x98c76b23.
//
// Solidity: function getRedeemAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals, uint256 marketRate) pure returns(uint256)
func (_Pool *PoolCaller) GetRedeemAmount(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8, marketRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRedeemAmount", tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals, marketRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedeemAmount is a free data retrieval call binding the contract method 0x98c76b23.
//
// Solidity: function getRedeemAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals, uint256 marketRate) pure returns(uint256)
func (_Pool *PoolSession) GetRedeemAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8, marketRate *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetRedeemAmount(&_Pool.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals, marketRate)
}

// GetRedeemAmount is a free data retrieval call binding the contract method 0x98c76b23.
//
// Solidity: function getRedeemAmount(uint8 tokenType, uint256 depositAmount, uint256 bondSupply, uint256 levSupply, uint256 poolReserves, uint256 ethPrice, uint8 oracleDecimals, uint256 marketRate) pure returns(uint256)
func (_Pool *PoolCallerSession) GetRedeemAmount(tokenType uint8, depositAmount *big.Int, bondSupply *big.Int, levSupply *big.Int, poolReserves *big.Int, ethPrice *big.Int, oracleDecimals uint8, marketRate *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetRedeemAmount(&_Pool.CallOpts, tokenType, depositAmount, bondSupply, levSupply, poolReserves, ethPrice, oracleDecimals, marketRate)
}

// LToken is a free data retrieval call binding the contract method 0x010ee184.
//
// Solidity: function lToken() view returns(address)
func (_Pool *PoolCaller) LToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "lToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LToken is a free data retrieval call binding the contract method 0x010ee184.
//
// Solidity: function lToken() view returns(address)
func (_Pool *PoolSession) LToken() (common.Address, error) {
	return _Pool.Contract.LToken(&_Pool.CallOpts)
}

// LToken is a free data retrieval call binding the contract method 0x010ee184.
//
// Solidity: function lToken() view returns(address)
func (_Pool *PoolCallerSession) LToken() (common.Address, error) {
	return _Pool.Contract.LToken(&_Pool.CallOpts)
}

// LastAuctionStart is a free data retrieval call binding the contract method 0x578277e3.
//
// Solidity: function lastAuctionStart() view returns(uint256)
func (_Pool *PoolCaller) LastAuctionStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "lastAuctionStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAuctionStart is a free data retrieval call binding the contract method 0x578277e3.
//
// Solidity: function lastAuctionStart() view returns(uint256)
func (_Pool *PoolSession) LastAuctionStart() (*big.Int, error) {
	return _Pool.Contract.LastAuctionStart(&_Pool.CallOpts)
}

// LastAuctionStart is a free data retrieval call binding the contract method 0x578277e3.
//
// Solidity: function lastAuctionStart() view returns(uint256)
func (_Pool *PoolCallerSession) LastAuctionStart() (*big.Int, error) {
	return _Pool.Contract.LastAuctionStart(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCallerSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_Pool *PoolCaller) OracleFeeds(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "oracleFeeds")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_Pool *PoolSession) OracleFeeds() (common.Address, error) {
	return _Pool.Contract.OracleFeeds(&_Pool.CallOpts)
}

// OracleFeeds is a free data retrieval call binding the contract method 0x2bbc2643.
//
// Solidity: function oracleFeeds() view returns(address)
func (_Pool *PoolCallerSession) OracleFeeds() (common.Address, error) {
	return _Pool.Contract.OracleFeeds(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCallerSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Pool *PoolCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Pool *PoolSession) PoolFactory() (common.Address, error) {
	return _Pool.Contract.PoolFactory(&_Pool.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Pool *PoolCallerSession) PoolFactory() (common.Address, error) {
	return _Pool.Contract.PoolFactory(&_Pool.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_Pool *PoolCaller) ReserveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "reserveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_Pool *PoolSession) ReserveToken() (common.Address, error) {
	return _Pool.Contract.ReserveToken(&_Pool.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_Pool *PoolCallerSession) ReserveToken() (common.Address, error) {
	return _Pool.Contract.ReserveToken(&_Pool.CallOpts)
}

// SimulateCreate is a free data retrieval call binding the contract method 0x8afb27b0.
//
// Solidity: function simulateCreate(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_Pool *PoolCaller) SimulateCreate(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "simulateCreate", tokenType, depositAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateCreate is a free data retrieval call binding the contract method 0x8afb27b0.
//
// Solidity: function simulateCreate(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_Pool *PoolSession) SimulateCreate(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.SimulateCreate(&_Pool.CallOpts, tokenType, depositAmount)
}

// SimulateCreate is a free data retrieval call binding the contract method 0x8afb27b0.
//
// Solidity: function simulateCreate(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_Pool *PoolCallerSession) SimulateCreate(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.SimulateCreate(&_Pool.CallOpts, tokenType, depositAmount)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0x9d64bca4.
//
// Solidity: function simulateRedeem(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_Pool *PoolCaller) SimulateRedeem(opts *bind.CallOpts, tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "simulateRedeem", tokenType, depositAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SimulateRedeem is a free data retrieval call binding the contract method 0x9d64bca4.
//
// Solidity: function simulateRedeem(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_Pool *PoolSession) SimulateRedeem(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.SimulateRedeem(&_Pool.CallOpts, tokenType, depositAmount)
}

// SimulateRedeem is a free data retrieval call binding the contract method 0x9d64bca4.
//
// Solidity: function simulateRedeem(uint8 tokenType, uint256 depositAmount) view returns(uint256)
func (_Pool *PoolCallerSession) SimulateRedeem(tokenType uint8, depositAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.SimulateRedeem(&_Pool.CallOpts, tokenType, depositAmount)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Pool *PoolTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Pool *PoolSession) ClaimFees() (*types.Transaction, error) {
	return _Pool.Contract.ClaimFees(&_Pool.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_Pool *PoolTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _Pool.Contract.ClaimFees(&_Pool.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0x6e530e97.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_Pool *PoolTransactor) Create(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "create", tokenType, depositAmount, minAmount)
}

// Create is a paid mutator transaction binding the contract method 0x6e530e97.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_Pool *PoolSession) Create(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Create(&_Pool.TransactOpts, tokenType, depositAmount, minAmount)
}

// Create is a paid mutator transaction binding the contract method 0x6e530e97.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_Pool *PoolTransactorSession) Create(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Create(&_Pool.TransactOpts, tokenType, depositAmount, minAmount)
}

// Create0 is a paid mutator transaction binding the contract method 0xf9d22289.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_Pool *PoolTransactor) Create0(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "create0", tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Create0 is a paid mutator transaction binding the contract method 0xf9d22289.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_Pool *PoolSession) Create0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Create0(&_Pool.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Create0 is a paid mutator transaction binding the contract method 0xf9d22289.
//
// Solidity: function create(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_Pool *PoolTransactorSession) Create0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Create0(&_Pool.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Pool *PoolTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Pool *PoolSession) Distribute() (*types.Transaction, error) {
	return _Pool.Contract.Distribute(&_Pool.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_Pool *PoolTransactorSession) Distribute() (*types.Transaction, error) {
	return _Pool.Contract.Distribute(&_Pool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e50408e.
//
// Solidity: function initialize(address _poolFactory, uint256 _fee, address _reserveToken, address _dToken, address _lToken, address _couponToken, uint256 _sharesPerToken, uint256 _distributionPeriod, address _feeBeneficiary, address _oracleFeeds, bool _pauseOnCreation) returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts, _poolFactory common.Address, _fee *big.Int, _reserveToken common.Address, _dToken common.Address, _lToken common.Address, _couponToken common.Address, _sharesPerToken *big.Int, _distributionPeriod *big.Int, _feeBeneficiary common.Address, _oracleFeeds common.Address, _pauseOnCreation bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize", _poolFactory, _fee, _reserveToken, _dToken, _lToken, _couponToken, _sharesPerToken, _distributionPeriod, _feeBeneficiary, _oracleFeeds, _pauseOnCreation)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e50408e.
//
// Solidity: function initialize(address _poolFactory, uint256 _fee, address _reserveToken, address _dToken, address _lToken, address _couponToken, uint256 _sharesPerToken, uint256 _distributionPeriod, address _feeBeneficiary, address _oracleFeeds, bool _pauseOnCreation) returns()
func (_Pool *PoolSession) Initialize(_poolFactory common.Address, _fee *big.Int, _reserveToken common.Address, _dToken common.Address, _lToken common.Address, _couponToken common.Address, _sharesPerToken *big.Int, _distributionPeriod *big.Int, _feeBeneficiary common.Address, _oracleFeeds common.Address, _pauseOnCreation bool) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, _poolFactory, _fee, _reserveToken, _dToken, _lToken, _couponToken, _sharesPerToken, _distributionPeriod, _feeBeneficiary, _oracleFeeds, _pauseOnCreation)
}

// Initialize is a paid mutator transaction binding the contract method 0x2e50408e.
//
// Solidity: function initialize(address _poolFactory, uint256 _fee, address _reserveToken, address _dToken, address _lToken, address _couponToken, uint256 _sharesPerToken, uint256 _distributionPeriod, address _feeBeneficiary, address _oracleFeeds, bool _pauseOnCreation) returns()
func (_Pool *PoolTransactorSession) Initialize(_poolFactory common.Address, _fee *big.Int, _reserveToken common.Address, _dToken common.Address, _lToken common.Address, _couponToken common.Address, _sharesPerToken *big.Int, _distributionPeriod *big.Int, _feeBeneficiary common.Address, _oracleFeeds common.Address, _pauseOnCreation bool) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, _poolFactory, _fee, _reserveToken, _dToken, _lToken, _couponToken, _sharesPerToken, _distributionPeriod, _feeBeneficiary, _oracleFeeds, _pauseOnCreation)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactorSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x724786cf.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_Pool *PoolTransactor) Redeem(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "redeem", tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Redeem is a paid mutator transaction binding the contract method 0x724786cf.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_Pool *PoolSession) Redeem(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Redeem(&_Pool.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Redeem is a paid mutator transaction binding the contract method 0x724786cf.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount, uint256 deadline, address onBehalfOf) returns(uint256)
func (_Pool *PoolTransactorSession) Redeem(tokenType uint8, depositAmount *big.Int, minAmount *big.Int, deadline *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Redeem(&_Pool.TransactOpts, tokenType, depositAmount, minAmount, deadline, onBehalfOf)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xf0fae20f.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_Pool *PoolTransactor) Redeem0(opts *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "redeem0", tokenType, depositAmount, minAmount)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xf0fae20f.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_Pool *PoolSession) Redeem0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Redeem0(&_Pool.TransactOpts, tokenType, depositAmount, minAmount)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xf0fae20f.
//
// Solidity: function redeem(uint8 tokenType, uint256 depositAmount, uint256 minAmount) returns(uint256)
func (_Pool *PoolTransactorSession) Redeem0(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Redeem0(&_Pool.TransactOpts, tokenType, depositAmount, minAmount)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xc8c73e7c.
//
// Solidity: function setAuctionPeriod(uint256 _auctionPeriod) returns()
func (_Pool *PoolTransactor) SetAuctionPeriod(opts *bind.TransactOpts, _auctionPeriod *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setAuctionPeriod", _auctionPeriod)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xc8c73e7c.
//
// Solidity: function setAuctionPeriod(uint256 _auctionPeriod) returns()
func (_Pool *PoolSession) SetAuctionPeriod(_auctionPeriod *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetAuctionPeriod(&_Pool.TransactOpts, _auctionPeriod)
}

// SetAuctionPeriod is a paid mutator transaction binding the contract method 0xc8c73e7c.
//
// Solidity: function setAuctionPeriod(uint256 _auctionPeriod) returns()
func (_Pool *PoolTransactorSession) SetAuctionPeriod(_auctionPeriod *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetAuctionPeriod(&_Pool.TransactOpts, _auctionPeriod)
}

// SetDistributionPeriod is a paid mutator transaction binding the contract method 0x7710c6d8.
//
// Solidity: function setDistributionPeriod(uint256 _distributionPeriod) returns()
func (_Pool *PoolTransactor) SetDistributionPeriod(opts *bind.TransactOpts, _distributionPeriod *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setDistributionPeriod", _distributionPeriod)
}

// SetDistributionPeriod is a paid mutator transaction binding the contract method 0x7710c6d8.
//
// Solidity: function setDistributionPeriod(uint256 _distributionPeriod) returns()
func (_Pool *PoolSession) SetDistributionPeriod(_distributionPeriod *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetDistributionPeriod(&_Pool.TransactOpts, _distributionPeriod)
}

// SetDistributionPeriod is a paid mutator transaction binding the contract method 0x7710c6d8.
//
// Solidity: function setDistributionPeriod(uint256 _distributionPeriod) returns()
func (_Pool *PoolTransactorSession) SetDistributionPeriod(_distributionPeriod *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetDistributionPeriod(&_Pool.TransactOpts, _distributionPeriod)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Pool *PoolTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Pool *PoolSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetFee(&_Pool.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Pool *PoolTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetFee(&_Pool.TransactOpts, _fee)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address _feeBeneficiary) returns()
func (_Pool *PoolTransactor) SetFeeBeneficiary(opts *bind.TransactOpts, _feeBeneficiary common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setFeeBeneficiary", _feeBeneficiary)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address _feeBeneficiary) returns()
func (_Pool *PoolSession) SetFeeBeneficiary(_feeBeneficiary common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetFeeBeneficiary(&_Pool.TransactOpts, _feeBeneficiary)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address _feeBeneficiary) returns()
func (_Pool *PoolTransactorSession) SetFeeBeneficiary(_feeBeneficiary common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetFeeBeneficiary(&_Pool.TransactOpts, _feeBeneficiary)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Pool *PoolTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Pool *PoolSession) SetName(_name string) (*types.Transaction, error) {
	return _Pool.Contract.SetName(&_Pool.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Pool *PoolTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _Pool.Contract.SetName(&_Pool.TransactOpts, _name)
}

// SetPoolSaleLimit is a paid mutator transaction binding the contract method 0xe14a81a9.
//
// Solidity: function setPoolSaleLimit(uint256 _poolSaleLimit) returns()
func (_Pool *PoolTransactor) SetPoolSaleLimit(opts *bind.TransactOpts, _poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setPoolSaleLimit", _poolSaleLimit)
}

// SetPoolSaleLimit is a paid mutator transaction binding the contract method 0xe14a81a9.
//
// Solidity: function setPoolSaleLimit(uint256 _poolSaleLimit) returns()
func (_Pool *PoolSession) SetPoolSaleLimit(_poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetPoolSaleLimit(&_Pool.TransactOpts, _poolSaleLimit)
}

// SetPoolSaleLimit is a paid mutator transaction binding the contract method 0xe14a81a9.
//
// Solidity: function setPoolSaleLimit(uint256 _poolSaleLimit) returns()
func (_Pool *PoolTransactorSession) SetPoolSaleLimit(_poolSaleLimit *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetPoolSaleLimit(&_Pool.TransactOpts, _poolSaleLimit)
}

// SetSharesPerToken is a paid mutator transaction binding the contract method 0x7362d149.
//
// Solidity: function setSharesPerToken(uint256 _sharesPerToken) returns()
func (_Pool *PoolTransactor) SetSharesPerToken(opts *bind.TransactOpts, _sharesPerToken *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setSharesPerToken", _sharesPerToken)
}

// SetSharesPerToken is a paid mutator transaction binding the contract method 0x7362d149.
//
// Solidity: function setSharesPerToken(uint256 _sharesPerToken) returns()
func (_Pool *PoolSession) SetSharesPerToken(_sharesPerToken *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetSharesPerToken(&_Pool.TransactOpts, _sharesPerToken)
}

// SetSharesPerToken is a paid mutator transaction binding the contract method 0x7362d149.
//
// Solidity: function setSharesPerToken(uint256 _sharesPerToken) returns()
func (_Pool *PoolTransactorSession) SetSharesPerToken(_sharesPerToken *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetSharesPerToken(&_Pool.TransactOpts, _sharesPerToken)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Pool *PoolTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Pool *PoolSession) StartAuction() (*types.Transaction, error) {
	return _Pool.Contract.StartAuction(&_Pool.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns()
func (_Pool *PoolTransactorSession) StartAuction() (*types.Transaction, error) {
	return _Pool.Contract.StartAuction(&_Pool.TransactOpts)
}

// TransferReserveToAuction is a paid mutator transaction binding the contract method 0xec27344c.
//
// Solidity: function transferReserveToAuction(uint256 amount) returns()
func (_Pool *PoolTransactor) TransferReserveToAuction(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferReserveToAuction", amount)
}

// TransferReserveToAuction is a paid mutator transaction binding the contract method 0xec27344c.
//
// Solidity: function transferReserveToAuction(uint256 amount) returns()
func (_Pool *PoolSession) TransferReserveToAuction(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferReserveToAuction(&_Pool.TransactOpts, amount)
}

// TransferReserveToAuction is a paid mutator transaction binding the contract method 0xec27344c.
//
// Solidity: function transferReserveToAuction(uint256 amount) returns()
func (_Pool *PoolTransactorSession) TransferReserveToAuction(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferReserveToAuction(&_Pool.TransactOpts, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// ZeroLastSharesPerToken is a paid mutator transaction binding the contract method 0x10c04932.
//
// Solidity: function zeroLastSharesPerToken() returns()
func (_Pool *PoolTransactor) ZeroLastSharesPerToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "zeroLastSharesPerToken")
}

// ZeroLastSharesPerToken is a paid mutator transaction binding the contract method 0x10c04932.
//
// Solidity: function zeroLastSharesPerToken() returns()
func (_Pool *PoolSession) ZeroLastSharesPerToken() (*types.Transaction, error) {
	return _Pool.Contract.ZeroLastSharesPerToken(&_Pool.TransactOpts)
}

// ZeroLastSharesPerToken is a paid mutator transaction binding the contract method 0x10c04932.
//
// Solidity: function zeroLastSharesPerToken() returns()
func (_Pool *PoolTransactorSession) ZeroLastSharesPerToken() (*types.Transaction, error) {
	return _Pool.Contract.ZeroLastSharesPerToken(&_Pool.TransactOpts)
}

// PoolAuctionPeriodChangedIterator is returned from FilterAuctionPeriodChanged and is used to iterate over the raw logs and unpacked data for AuctionPeriodChanged events raised by the Pool contract.
type PoolAuctionPeriodChangedIterator struct {
	Event *PoolAuctionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *PoolAuctionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAuctionPeriodChanged)
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
		it.Event = new(PoolAuctionPeriodChanged)
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
func (it *PoolAuctionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAuctionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAuctionPeriodChanged represents a AuctionPeriodChanged event raised by the Pool contract.
type PoolAuctionPeriodChanged struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionPeriodChanged is a free log retrieval operation binding the contract event 0x8b63bdaf0f5981cb255f97be2272e1ea65195664462f04777a64848795236aed.
//
// Solidity: event AuctionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_Pool *PoolFilterer) FilterAuctionPeriodChanged(opts *bind.FilterOpts) (*PoolAuctionPeriodChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AuctionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &PoolAuctionPeriodChangedIterator{contract: _Pool.contract, event: "AuctionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchAuctionPeriodChanged is a free log subscription operation binding the contract event 0x8b63bdaf0f5981cb255f97be2272e1ea65195664462f04777a64848795236aed.
//
// Solidity: event AuctionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_Pool *PoolFilterer) WatchAuctionPeriodChanged(opts *bind.WatchOpts, sink chan<- *PoolAuctionPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AuctionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAuctionPeriodChanged)
				if err := _Pool.contract.UnpackLog(event, "AuctionPeriodChanged", log); err != nil {
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
func (_Pool *PoolFilterer) ParseAuctionPeriodChanged(log types.Log) (*PoolAuctionPeriodChanged, error) {
	event := new(PoolAuctionPeriodChanged)
	if err := _Pool.contract.UnpackLog(event, "AuctionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the Pool contract.
type PoolAuctionStartedIterator struct {
	Event *PoolAuctionStarted // Event containing the contract specifics and raw log

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
func (it *PoolAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAuctionStarted)
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
		it.Event = new(PoolAuctionStarted)
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
func (it *PoolAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAuctionStarted represents a AuctionStarted event raised by the Pool contract.
type PoolAuctionStarted struct {
	Auction                  common.Address
	Period                   *big.Int
	CouponAmountToDistribute *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x92db7303de5c78e58e4570f0b4acf392e3dac09394e96c4fa62b323b6337ee65.
//
// Solidity: event AuctionStarted(address auction, uint256 period, uint256 couponAmountToDistribute)
func (_Pool *PoolFilterer) FilterAuctionStarted(opts *bind.FilterOpts) (*PoolAuctionStartedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return &PoolAuctionStartedIterator{contract: _Pool.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x92db7303de5c78e58e4570f0b4acf392e3dac09394e96c4fa62b323b6337ee65.
//
// Solidity: event AuctionStarted(address auction, uint256 period, uint256 couponAmountToDistribute)
func (_Pool *PoolFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *PoolAuctionStarted) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAuctionStarted)
				if err := _Pool.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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
func (_Pool *PoolFilterer) ParseAuctionStarted(log types.Log) (*PoolAuctionStarted, error) {
	event := new(PoolAuctionStarted)
	if err := _Pool.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDistributedIterator is returned from FilterDistributed and is used to iterate over the raw logs and unpacked data for Distributed events raised by the Pool contract.
type PoolDistributedIterator struct {
	Event *PoolDistributed // Event containing the contract specifics and raw log

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
func (it *PoolDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDistributed)
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
		it.Event = new(PoolDistributed)
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
func (it *PoolDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDistributed represents a Distributed event raised by the Pool contract.
type PoolDistributed struct {
	Period      *big.Int
	Amount      *big.Int
	Distributor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDistributed is a free log retrieval operation binding the contract event 0x8dda631b265286fb6a61a6658e6e75969e952f2787646e616dfc15a84f298c18.
//
// Solidity: event Distributed(uint256 period, uint256 amount, address distributor)
func (_Pool *PoolFilterer) FilterDistributed(opts *bind.FilterOpts) (*PoolDistributedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Distributed")
	if err != nil {
		return nil, err
	}
	return &PoolDistributedIterator{contract: _Pool.contract, event: "Distributed", logs: logs, sub: sub}, nil
}

// WatchDistributed is a free log subscription operation binding the contract event 0x8dda631b265286fb6a61a6658e6e75969e952f2787646e616dfc15a84f298c18.
//
// Solidity: event Distributed(uint256 period, uint256 amount, address distributor)
func (_Pool *PoolFilterer) WatchDistributed(opts *bind.WatchOpts, sink chan<- *PoolDistributed) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Distributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDistributed)
				if err := _Pool.contract.UnpackLog(event, "Distributed", log); err != nil {
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
func (_Pool *PoolFilterer) ParseDistributed(log types.Log) (*PoolDistributed, error) {
	event := new(PoolDistributed)
	if err := _Pool.contract.UnpackLog(event, "Distributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDistributionPeriodChangedIterator is returned from FilterDistributionPeriodChanged and is used to iterate over the raw logs and unpacked data for DistributionPeriodChanged events raised by the Pool contract.
type PoolDistributionPeriodChangedIterator struct {
	Event *PoolDistributionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *PoolDistributionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDistributionPeriodChanged)
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
		it.Event = new(PoolDistributionPeriodChanged)
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
func (it *PoolDistributionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDistributionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDistributionPeriodChanged represents a DistributionPeriodChanged event raised by the Pool contract.
type PoolDistributionPeriodChanged struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributionPeriodChanged is a free log retrieval operation binding the contract event 0x614ada6efad9c4a86e8541da67aa075e9e21e37f888c79f23807667abb742ce3.
//
// Solidity: event DistributionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_Pool *PoolFilterer) FilterDistributionPeriodChanged(opts *bind.FilterOpts) (*PoolDistributionPeriodChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DistributionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &PoolDistributionPeriodChangedIterator{contract: _Pool.contract, event: "DistributionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchDistributionPeriodChanged is a free log subscription operation binding the contract event 0x614ada6efad9c4a86e8541da67aa075e9e21e37f888c79f23807667abb742ce3.
//
// Solidity: event DistributionPeriodChanged(uint256 oldPeriod, uint256 newPeriod)
func (_Pool *PoolFilterer) WatchDistributionPeriodChanged(opts *bind.WatchOpts, sink chan<- *PoolDistributionPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DistributionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDistributionPeriodChanged)
				if err := _Pool.contract.UnpackLog(event, "DistributionPeriodChanged", log); err != nil {
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
func (_Pool *PoolFilterer) ParseDistributionPeriodChanged(log types.Log) (*PoolDistributionPeriodChanged, error) {
	event := new(PoolDistributionPeriodChanged)
	if err := _Pool.contract.UnpackLog(event, "DistributionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDistributionRollOverIterator is returned from FilterDistributionRollOver and is used to iterate over the raw logs and unpacked data for DistributionRollOver events raised by the Pool contract.
type PoolDistributionRollOverIterator struct {
	Event *PoolDistributionRollOver // Event containing the contract specifics and raw log

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
func (it *PoolDistributionRollOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDistributionRollOver)
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
		it.Event = new(PoolDistributionRollOver)
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
func (it *PoolDistributionRollOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDistributionRollOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDistributionRollOver represents a DistributionRollOver event raised by the Pool contract.
type PoolDistributionRollOver struct {
	Period *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributionRollOver is a free log retrieval operation binding the contract event 0xb07ac059d2a0dd966a73cbd0cb9b6c1b7af9468cab03dc1fd5d34c6fb4a6b30c.
//
// Solidity: event DistributionRollOver(uint256 period, uint256 shares)
func (_Pool *PoolFilterer) FilterDistributionRollOver(opts *bind.FilterOpts) (*PoolDistributionRollOverIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DistributionRollOver")
	if err != nil {
		return nil, err
	}
	return &PoolDistributionRollOverIterator{contract: _Pool.contract, event: "DistributionRollOver", logs: logs, sub: sub}, nil
}

// WatchDistributionRollOver is a free log subscription operation binding the contract event 0xb07ac059d2a0dd966a73cbd0cb9b6c1b7af9468cab03dc1fd5d34c6fb4a6b30c.
//
// Solidity: event DistributionRollOver(uint256 period, uint256 shares)
func (_Pool *PoolFilterer) WatchDistributionRollOver(opts *bind.WatchOpts, sink chan<- *PoolDistributionRollOver) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DistributionRollOver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDistributionRollOver)
				if err := _Pool.contract.UnpackLog(event, "DistributionRollOver", log); err != nil {
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
func (_Pool *PoolFilterer) ParseDistributionRollOver(log types.Log) (*PoolDistributionRollOver, error) {
	event := new(PoolDistributionRollOver)
	if err := _Pool.contract.UnpackLog(event, "DistributionRollOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFeeChangedIterator is returned from FilterFeeChanged and is used to iterate over the raw logs and unpacked data for FeeChanged events raised by the Pool contract.
type PoolFeeChangedIterator struct {
	Event *PoolFeeChanged // Event containing the contract specifics and raw log

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
func (it *PoolFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFeeChanged)
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
		it.Event = new(PoolFeeChanged)
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
func (it *PoolFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFeeChanged represents a FeeChanged event raised by the Pool contract.
type PoolFeeChanged struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeChanged is a free log retrieval operation binding the contract event 0x5fc463da23c1b063e66f9e352006a7fbe8db7223c455dc429e881a2dfe2f94f1.
//
// Solidity: event FeeChanged(uint256 oldFee, uint256 newFee)
func (_Pool *PoolFilterer) FilterFeeChanged(opts *bind.FilterOpts) (*PoolFeeChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return &PoolFeeChangedIterator{contract: _Pool.contract, event: "FeeChanged", logs: logs, sub: sub}, nil
}

// WatchFeeChanged is a free log subscription operation binding the contract event 0x5fc463da23c1b063e66f9e352006a7fbe8db7223c455dc429e881a2dfe2f94f1.
//
// Solidity: event FeeChanged(uint256 oldFee, uint256 newFee)
func (_Pool *PoolFilterer) WatchFeeChanged(opts *bind.WatchOpts, sink chan<- *PoolFeeChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "FeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFeeChanged)
				if err := _Pool.contract.UnpackLog(event, "FeeChanged", log); err != nil {
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
func (_Pool *PoolFilterer) ParseFeeChanged(log types.Log) (*PoolFeeChanged, error) {
	event := new(PoolFeeChanged)
	if err := _Pool.contract.UnpackLog(event, "FeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFeeClaimedIterator is returned from FilterFeeClaimed and is used to iterate over the raw logs and unpacked data for FeeClaimed events raised by the Pool contract.
type PoolFeeClaimedIterator struct {
	Event *PoolFeeClaimed // Event containing the contract specifics and raw log

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
func (it *PoolFeeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFeeClaimed)
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
		it.Event = new(PoolFeeClaimed)
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
func (it *PoolFeeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFeeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFeeClaimed represents a FeeClaimed event raised by the Pool contract.
type PoolFeeClaimed struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeClaimed is a free log retrieval operation binding the contract event 0x20ca5094f3a20c321cbe4123d0f01b276b81df0fa24cd4d83d9253956035d863.
//
// Solidity: event FeeClaimed(address beneficiary, uint256 amount)
func (_Pool *PoolFilterer) FilterFeeClaimed(opts *bind.FilterOpts) (*PoolFeeClaimedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "FeeClaimed")
	if err != nil {
		return nil, err
	}
	return &PoolFeeClaimedIterator{contract: _Pool.contract, event: "FeeClaimed", logs: logs, sub: sub}, nil
}

// WatchFeeClaimed is a free log subscription operation binding the contract event 0x20ca5094f3a20c321cbe4123d0f01b276b81df0fa24cd4d83d9253956035d863.
//
// Solidity: event FeeClaimed(address beneficiary, uint256 amount)
func (_Pool *PoolFilterer) WatchFeeClaimed(opts *bind.WatchOpts, sink chan<- *PoolFeeClaimed) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "FeeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFeeClaimed)
				if err := _Pool.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
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
func (_Pool *PoolFilterer) ParseFeeClaimed(log types.Log) (*PoolFeeClaimed, error) {
	event := new(PoolFeeClaimed)
	if err := _Pool.contract.UnpackLog(event, "FeeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pool contract.
type PoolInitializedIterator struct {
	Event *PoolInitialized // Event containing the contract specifics and raw log

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
func (it *PoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolInitialized)
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
		it.Event = new(PoolInitialized)
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
func (it *PoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolInitialized represents a Initialized event raised by the Pool contract.
type PoolInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Pool *PoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolInitializedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolInitializedIterator{contract: _Pool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Pool *PoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolInitialized) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolInitialized)
				if err := _Pool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Pool *PoolFilterer) ParseInitialized(log types.Log) (*PoolInitialized, error) {
	event := new(PoolInitialized)
	if err := _Pool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolNoFeesToClaimIterator is returned from FilterNoFeesToClaim and is used to iterate over the raw logs and unpacked data for NoFeesToClaim events raised by the Pool contract.
type PoolNoFeesToClaimIterator struct {
	Event *PoolNoFeesToClaim // Event containing the contract specifics and raw log

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
func (it *PoolNoFeesToClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolNoFeesToClaim)
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
		it.Event = new(PoolNoFeesToClaim)
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
func (it *PoolNoFeesToClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolNoFeesToClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolNoFeesToClaim represents a NoFeesToClaim event raised by the Pool contract.
type PoolNoFeesToClaim struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNoFeesToClaim is a free log retrieval operation binding the contract event 0x846d8c5c020eac85233c31fc79d004ef12b5e20abfbe63b23c49f764393b4482.
//
// Solidity: event NoFeesToClaim()
func (_Pool *PoolFilterer) FilterNoFeesToClaim(opts *bind.FilterOpts) (*PoolNoFeesToClaimIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "NoFeesToClaim")
	if err != nil {
		return nil, err
	}
	return &PoolNoFeesToClaimIterator{contract: _Pool.contract, event: "NoFeesToClaim", logs: logs, sub: sub}, nil
}

// WatchNoFeesToClaim is a free log subscription operation binding the contract event 0x846d8c5c020eac85233c31fc79d004ef12b5e20abfbe63b23c49f764393b4482.
//
// Solidity: event NoFeesToClaim()
func (_Pool *PoolFilterer) WatchNoFeesToClaim(opts *bind.WatchOpts, sink chan<- *PoolNoFeesToClaim) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "NoFeesToClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolNoFeesToClaim)
				if err := _Pool.contract.UnpackLog(event, "NoFeesToClaim", log); err != nil {
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
func (_Pool *PoolFilterer) ParseNoFeesToClaim(log types.Log) (*PoolNoFeesToClaim, error) {
	event := new(PoolNoFeesToClaim)
	if err := _Pool.contract.UnpackLog(event, "NoFeesToClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pool contract.
type PoolPausedIterator struct {
	Event *PoolPaused // Event containing the contract specifics and raw log

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
func (it *PoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPaused)
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
		it.Event = new(PoolPaused)
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
func (it *PoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPaused represents a Paused event raised by the Pool contract.
type PoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolPausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolPausedIterator{contract: _Pool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolPaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPaused)
				if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pool *PoolFilterer) ParsePaused(log types.Log) (*PoolPaused, error) {
	event := new(PoolPaused)
	if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPoolSaleLimitChangedIterator is returned from FilterPoolSaleLimitChanged and is used to iterate over the raw logs and unpacked data for PoolSaleLimitChanged events raised by the Pool contract.
type PoolPoolSaleLimitChangedIterator struct {
	Event *PoolPoolSaleLimitChanged // Event containing the contract specifics and raw log

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
func (it *PoolPoolSaleLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPoolSaleLimitChanged)
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
		it.Event = new(PoolPoolSaleLimitChanged)
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
func (it *PoolPoolSaleLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPoolSaleLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPoolSaleLimitChanged represents a PoolSaleLimitChanged event raised by the Pool contract.
type PoolPoolSaleLimitChanged struct {
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolSaleLimitChanged is a free log retrieval operation binding the contract event 0xbd1900a27c42d19889a9199e391bdf6738af2b388f81b7d73db6013548a27e7d.
//
// Solidity: event PoolSaleLimitChanged(uint256 oldThreshold, uint256 newThreshold)
func (_Pool *PoolFilterer) FilterPoolSaleLimitChanged(opts *bind.FilterOpts) (*PoolPoolSaleLimitChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "PoolSaleLimitChanged")
	if err != nil {
		return nil, err
	}
	return &PoolPoolSaleLimitChangedIterator{contract: _Pool.contract, event: "PoolSaleLimitChanged", logs: logs, sub: sub}, nil
}

// WatchPoolSaleLimitChanged is a free log subscription operation binding the contract event 0xbd1900a27c42d19889a9199e391bdf6738af2b388f81b7d73db6013548a27e7d.
//
// Solidity: event PoolSaleLimitChanged(uint256 oldThreshold, uint256 newThreshold)
func (_Pool *PoolFilterer) WatchPoolSaleLimitChanged(opts *bind.WatchOpts, sink chan<- *PoolPoolSaleLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "PoolSaleLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPoolSaleLimitChanged)
				if err := _Pool.contract.UnpackLog(event, "PoolSaleLimitChanged", log); err != nil {
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
func (_Pool *PoolFilterer) ParsePoolSaleLimitChanged(log types.Log) (*PoolPoolSaleLimitChanged, error) {
	event := new(PoolPoolSaleLimitChanged)
	if err := _Pool.contract.UnpackLog(event, "PoolSaleLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSharesPerTokenChangedIterator is returned from FilterSharesPerTokenChanged and is used to iterate over the raw logs and unpacked data for SharesPerTokenChanged events raised by the Pool contract.
type PoolSharesPerTokenChangedIterator struct {
	Event *PoolSharesPerTokenChanged // Event containing the contract specifics and raw log

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
func (it *PoolSharesPerTokenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSharesPerTokenChanged)
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
		it.Event = new(PoolSharesPerTokenChanged)
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
func (it *PoolSharesPerTokenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSharesPerTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSharesPerTokenChanged represents a SharesPerTokenChanged event raised by the Pool contract.
type PoolSharesPerTokenChanged struct {
	OldSharesPerToken *big.Int
	SharesPerToken    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSharesPerTokenChanged is a free log retrieval operation binding the contract event 0x67b8059073c48bc22a124c60cf344bd97cf9b265f273fa09afe215b64eb5d2c2.
//
// Solidity: event SharesPerTokenChanged(uint256 oldSharesPerToken, uint256 sharesPerToken)
func (_Pool *PoolFilterer) FilterSharesPerTokenChanged(opts *bind.FilterOpts) (*PoolSharesPerTokenChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "SharesPerTokenChanged")
	if err != nil {
		return nil, err
	}
	return &PoolSharesPerTokenChangedIterator{contract: _Pool.contract, event: "SharesPerTokenChanged", logs: logs, sub: sub}, nil
}

// WatchSharesPerTokenChanged is a free log subscription operation binding the contract event 0x67b8059073c48bc22a124c60cf344bd97cf9b265f273fa09afe215b64eb5d2c2.
//
// Solidity: event SharesPerTokenChanged(uint256 oldSharesPerToken, uint256 sharesPerToken)
func (_Pool *PoolFilterer) WatchSharesPerTokenChanged(opts *bind.WatchOpts, sink chan<- *PoolSharesPerTokenChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "SharesPerTokenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSharesPerTokenChanged)
				if err := _Pool.contract.UnpackLog(event, "SharesPerTokenChanged", log); err != nil {
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
func (_Pool *PoolFilterer) ParseSharesPerTokenChanged(log types.Log) (*PoolSharesPerTokenChanged, error) {
	event := new(PoolSharesPerTokenChanged)
	if err := _Pool.contract.UnpackLog(event, "SharesPerTokenChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTokensCreatedIterator is returned from FilterTokensCreated and is used to iterate over the raw logs and unpacked data for TokensCreated events raised by the Pool contract.
type PoolTokensCreatedIterator struct {
	Event *PoolTokensCreated // Event containing the contract specifics and raw log

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
func (it *PoolTokensCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokensCreated)
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
		it.Event = new(PoolTokensCreated)
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
func (it *PoolTokensCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokensCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokensCreated represents a TokensCreated event raised by the Pool contract.
type PoolTokensCreated struct {
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
func (_Pool *PoolFilterer) FilterTokensCreated(opts *bind.FilterOpts, onBehalfOf []common.Address) (*PoolTokensCreatedIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "TokensCreated", onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokensCreatedIterator{contract: _Pool.contract, event: "TokensCreated", logs: logs, sub: sub}, nil
}

// WatchTokensCreated is a free log subscription operation binding the contract event 0xbd52f81ce4beba21c7c6a29213c7c718167799f216ed3ebe0f7224c3947fa525.
//
// Solidity: event TokensCreated(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 mintedAmount)
func (_Pool *PoolFilterer) WatchTokensCreated(opts *bind.WatchOpts, sink chan<- *PoolTokensCreated, onBehalfOf []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "TokensCreated", onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokensCreated)
				if err := _Pool.contract.UnpackLog(event, "TokensCreated", log); err != nil {
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
func (_Pool *PoolFilterer) ParseTokensCreated(log types.Log) (*PoolTokensCreated, error) {
	event := new(PoolTokensCreated)
	if err := _Pool.contract.UnpackLog(event, "TokensCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTokensRedeemedIterator is returned from FilterTokensRedeemed and is used to iterate over the raw logs and unpacked data for TokensRedeemed events raised by the Pool contract.
type PoolTokensRedeemedIterator struct {
	Event *PoolTokensRedeemed // Event containing the contract specifics and raw log

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
func (it *PoolTokensRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokensRedeemed)
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
		it.Event = new(PoolTokensRedeemed)
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
func (it *PoolTokensRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokensRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokensRedeemed represents a TokensRedeemed event raised by the Pool contract.
type PoolTokensRedeemed struct {
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
func (_Pool *PoolFilterer) FilterTokensRedeemed(opts *bind.FilterOpts, onBehalfOf []common.Address) (*PoolTokensRedeemedIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "TokensRedeemed", onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokensRedeemedIterator{contract: _Pool.contract, event: "TokensRedeemed", logs: logs, sub: sub}, nil
}

// WatchTokensRedeemed is a free log subscription operation binding the contract event 0x2441fb9f349bdfdf06755d1f5239efc1851c81c59e35378f3df9ac880bca3094.
//
// Solidity: event TokensRedeemed(address caller, address indexed onBehalfOf, uint8 tokenType, uint256 depositedAmount, uint256 redeemedAmount)
func (_Pool *PoolFilterer) WatchTokensRedeemed(opts *bind.WatchOpts, sink chan<- *PoolTokensRedeemed, onBehalfOf []common.Address) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "TokensRedeemed", onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokensRedeemed)
				if err := _Pool.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
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
func (_Pool *PoolFilterer) ParseTokensRedeemed(log types.Log) (*PoolTokensRedeemed, error) {
	event := new(PoolTokensRedeemed)
	if err := _Pool.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pool contract.
type PoolUnpausedIterator struct {
	Event *PoolUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUnpaused)
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
		it.Event = new(PoolUnpaused)
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
func (it *PoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUnpaused represents a Unpaused event raised by the Pool contract.
type PoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolUnpausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolUnpausedIterator{contract: _Pool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUnpaused)
				if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pool *PoolFilterer) ParseUnpaused(log types.Log) (*PoolUnpaused, error) {
	event := new(PoolUnpaused)
	if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

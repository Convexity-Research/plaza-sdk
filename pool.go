package plaza_sdk

import (
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PoolInfo struct {
	Fee               *big.Int
	Reserve           *big.Int
	BondSupply        *big.Int
	LevSupply         *big.Int
	SharesPerToken    *big.Int
	CurrentPeriod     *big.Int
	LastDistribution  *big.Int
	DistributionPeriod *big.Int
	AuctionPeriod     *big.Int
	FeeBeneficiary    common.Address
}

func NewPool(address common.Address, client *ethclient.Client) (*Pool, error) {
	byteValue, err := os.ReadFile("./abi/Pool.json")
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(byteValue)))
	if err != nil {
		return nil, err
	}

	return &Pool{
		Address: address,
		ABI:     parsedABI,
		Client:  client,
	}, nil
}

type Pool struct {
	Address common.Address
	ABI     abi.ABI
	Client  *ethclient.Client
}

func (p *Pool) GetBondPrice(opts *bind.CallOpts) (*big.Int, error) {
	// Use 1 ether as input amount
	oneEther := new(big.Int).Mul(big.NewInt(1), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	
	// Simulate create for bond token
	var result struct {
		BondAmount *big.Int
	}
	err := p.ABI.UnpackIntoInterface(&result, "simulateCreate", 0, oneEther)
	if err != nil {
		return nil, err
	}

	// Get reserve asset address
	var reserveAsset common.Address
	err = p.ABI.UnpackIntoInterface(&reserveAsset, "reserveAsset", nil)
	if err != nil {
		return nil, err
	}

	// Get oracle feeds contract
	oracleBytes, err := os.ReadFile("./abi/OracleFeeds.json")
	if err != nil {
		return nil, err
	}
	oracleABI, err := abi.JSON(strings.NewReader(string(oracleBytes)))
	if err != nil {
		return nil, err
	}

	// Get price feed address
	var priceFeed common.Address
	err = oracleABI.UnpackIntoInterface(&priceFeed, "priceFeeds", reserveAsset, common.Address{})
	if err != nil {
		return nil, err
	}

	// Get chainlink feed
	feedBytes, err := os.ReadFile("./abi/ChainlinkFeed.json")
	if err != nil {
		return nil, err
	}
	feedABI, err := abi.JSON(strings.NewReader(string(feedBytes)))
	if err != nil {
		return nil, err
	}

	// Get latest price
	var priceResult struct {
		RoundId         uint64
		Answer          *big.Int
		StartedAt       uint64
		UpdatedAt       uint64
		AnsweredInRound uint64
	}
	err = feedABI.UnpackIntoInterface(&priceResult, "latestRoundData", nil)
	if err != nil {
		return nil, err
	}

	// Calculate price per bond token
	// price = (1 ether * oracle price) / bond amount received
	price := new(big.Int).Mul(oneEther, priceResult.Answer)
	price = price.Div(price, result.BondAmount)

	return price, nil
}

func (p *Pool) GetLevPrice(opts *bind.CallOpts) (*big.Int, error) {
	// Use 1 ether as input amount
	oneEther := new(big.Int).Mul(big.NewInt(1), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	
	// Simulate create for lev token
	var result struct {
		LevAmount *big.Int
	}
	err := p.ABI.UnpackIntoInterface(&result, "simulateCreate", 1, oneEther)
	if err != nil {
		return nil, err
	}

	// Get reserve asset address
	var reserveAsset common.Address
	err = p.ABI.UnpackIntoInterface(&reserveAsset, "reserveAsset", nil)
	if err != nil {
		return nil, err
	}

	// Get oracle feeds contract
	oracleBytes, err := os.ReadFile("./abi/OracleFeeds.json")
	if err != nil {
		return nil, err
	}
	oracleABI, err := abi.JSON(strings.NewReader(string(oracleBytes)))
	if err != nil {
		return nil, err
	}

	// Get price feed address
	var priceFeed common.Address
	err = oracleABI.UnpackIntoInterface(&priceFeed, "priceFeeds", reserveAsset, common.Address{})
	if err != nil {
		return nil, err
	}

	// Get chainlink feed
	feedBytes, err := os.ReadFile("./abi/ChainlinkFeed.json")
	if err != nil {
		return nil, err
	}
	feedABI, err := abi.JSON(strings.NewReader(string(feedBytes)))
	if err != nil {
		return nil, err
	}

	// Get latest price
	var priceResult struct {
		RoundId         uint64
		Answer          *big.Int
		StartedAt       uint64
		UpdatedAt       uint64
		AnsweredInRound uint64
	}
	err = feedABI.UnpackIntoInterface(&priceResult, "latestRoundData", nil)
	if err != nil {
		return nil, err
	}

	// Calculate price per lev token
	// price = (1 ether * oracle price) / lev amount received
	price := new(big.Int).Mul(oneEther, priceResult.Answer)
	price = price.Div(price, result.LevAmount)

	return price, nil
}

func (p *Pool) GetTVL(opts *bind.CallOpts) (*big.Int, error) {
	// Get reserve asset address
	var reserveAsset common.Address
	err := p.ABI.UnpackIntoInterface(&reserveAsset, "reserveAsset", nil)
	if err != nil {
		return nil, err
	}

	// Get ERC20 ABI
	erc20Bytes, err := os.ReadFile("./abi/ERC20.json")
	if err != nil {
		return nil, err
	}
	erc20ABI, err := abi.JSON(strings.NewReader(string(erc20Bytes)))
	if err != nil {
		return nil, err
	}

	// Get balance of pool
	var balance *big.Int
	err = erc20ABI.UnpackIntoInterface(&balance, "balanceOf", p.Address)
	if err != nil {
		return nil, err
	}

	// Get oracle feeds ABI
	oracleBytes, err := os.ReadFile("./abi/OracleFeeds.json") 
	if err != nil {
		return nil, err
	}
	oracleABI, err := abi.JSON(strings.NewReader(string(oracleBytes)))
	if err != nil {
		return nil, err
	}

	// Get price feed address
	var priceFeed common.Address
	err = oracleABI.UnpackIntoInterface(&priceFeed, "priceFeeds", reserveAsset, common.Address{})
	if err != nil {
		return nil, err
	}

	// Get chainlink feed ABI
	feedBytes, err := os.ReadFile("./abi/ChainlinkFeed.json")
	if err != nil {
		return nil, err
	}
	feedABI, err := abi.JSON(strings.NewReader(string(feedBytes)))
	if err != nil {
		return nil, err
	}

	// Get latest price
	var result struct {
		RoundId         uint64
		Answer          *big.Int
		StartedAt       uint64
		UpdatedAt       uint64
		AnsweredInRound uint64
	}
	err = feedABI.UnpackIntoInterface(&result, "latestRoundData", nil)
	if err != nil {
		return nil, err
	}

	// Calculate TVL
	tvl := new(big.Int).Mul(balance, result.Answer)
	return tvl, nil
}

func (p *Pool) GetLTV(opts *bind.CallOpts) (*big.Int, error) {
	// Get bond token address and supply
	var bondTokenAddr common.Address
	err := p.ABI.UnpackIntoInterface(&bondTokenAddr, "bondToken", nil)
	if err != nil {
		return nil, err
	}

	// Get ERC20 ABI
	erc20Bytes, err := os.ReadFile("./abi/ERC20.json")
	if err != nil {
		return nil, err
	}
	erc20ABI, err := abi.JSON(strings.NewReader(string(erc20Bytes)))
	if err != nil {
		return nil, err
	}

	// Get total supply of bond tokens
	var bondSupply *big.Int
	err = erc20ABI.UnpackIntoInterface(&bondSupply, "totalSupply", nil)
	if err != nil {
		return nil, err
	}

	// Get bond price
	bondPrice, err := p.GetBondPrice(opts)
	if err != nil {
		return nil, err
	}

	// Get TVL
	tvl, err := p.GetTVL(opts)
	if err != nil {
		return nil, err
	}

	// Calculate LTV ratio (bond supply * bond price) / TVL
	numerator := new(big.Int).Mul(bondSupply, bondPrice)
	ltv := numerator.Div(numerator, tvl)

	return ltv, nil
}

func (p *Pool) GetLeverage(opts *bind.CallOpts) (*big.Int, error) {
	// Get LTV
	ltv, err := p.GetLTV(opts)
	if err != nil {
		return nil, err
	}

	// Calculate leverage = 1 / (1 - LTV)
	oneEther := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	leverage := new(big.Int).Sub(oneEther, ltv)
	leverage = oneEther.Div(oneEther, leverage)

	return leverage, nil
}

func (p *Pool) GetBondYield(opts *bind.CallOpts) (*big.Int, error) {
	// Get pool info to get sharesPerToken and distributionPeriod
	poolInfo, err := p.GetPoolInfo(opts)
	if err != nil {
		return nil, err
	}

	// Get bond price
	bondPrice, err := p.GetBondPrice(opts)
	if err != nil {
		return nil, err
	}

	// Calculate yield = (sharesPerToken / 10^6) / bondPrice * (3600 * 24 * 365) / distributionPeriod
	million := big.NewInt(1000000)
	sharesPerTokenNormalized := new(big.Int).Div(poolInfo.SharesPerToken, million)

	secondsPerYear := big.NewInt(3600 * 24 * 365)
	
	numerator := new(big.Int).Mul(sharesPerTokenNormalized, secondsPerYear)
	denominator := new(big.Int).Mul(bondPrice, poolInfo.DistributionPeriod)
	
	yield := numerator.Div(numerator, denominator)

	return yield, nil
}

func (p *Pool) Create(auth *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*bind.TransactOpts, error) {
	input, err := p.ABI.Pack("create", tokenType, depositAmount, minAmount)
	if err != nil {
		return nil, err
	}
	return bind.NewTransactor(auth, input)
}

func (p *Pool) Redeem(auth *bind.TransactOpts, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*bind.TransactOpts, error) {
	input, err := p.ABI.Pack("redeem", tokenType, depositAmount, minAmount)
	if err != nil {
		return nil, err
	}
	return bind.NewTransactor(auth, input)
}

func (p *Pool) GetPoolInfo(opts *bind.CallOpts) (*PoolInfo, error) {
	var result struct {
		Info struct {
			Fee               *big.Int
			Reserve           *big.Int
			BondSupply        *big.Int
			LevSupply         *big.Int
			SharesPerToken    *big.Int
			CurrentPeriod     *big.Int
			LastDistribution  *big.Int
			DistributionPeriod *big.Int
			AuctionPeriod     *big.Int
			FeeBeneficiary    common.Address
		}
	}

	err := p.ABI.UnpackIntoInterface(&result, "getPoolInfo", nil)
	if err != nil {
		return nil, err
	}

	return &PoolInfo{
		Fee:               result.Info.Fee,
		Reserve:           result.Info.Reserve,
		BondSupply:        result.Info.BondSupply,
		LevSupply:         result.Info.LevSupply,
		SharesPerToken:    result.Info.SharesPerToken,
		CurrentPeriod:     result.Info.CurrentPeriod,
		LastDistribution:  result.Info.LastDistribution,
		DistributionPeriod: result.Info.DistributionPeriod,
		AuctionPeriod:     result.Info.AuctionPeriod,
		FeeBeneficiary:    result.Info.FeeBeneficiary,
	}, nil
}

func Create(poolAddress string, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*big.Int, error) {
	pool, err := NewPool(common.HexToAddress(poolAddress), ethClient)
	if err != nil {
		return nil, err
	}

	tx, err := pool.Create(auth, tokenType, depositAmount, minAmount)
	if err != nil {
		return nil, err
	}
	log.Printf("Create transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Big(), nil
}

func Redeem(poolAddress string, tokenType uint8, depositAmount *big.Int, minAmount *big.Int) (*big.Int, error) {
	pool, err := NewPool(common.HexToAddress(poolAddress), ethClient)
	if err != nil {
		return nil, err
	}

	tx, err := pool.Redeem(auth, tokenType, depositAmount, minAmount)
	if err != nil {
		return nil, err
	}
	log.Printf("Redeem transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Big(), nil
}

func GetAuctionAddress(poolAddress string) (string, error) {
	pool, err := NewPool(common.HexToAddress(poolAddress), ethClient)
	if err != nil {
		return "", err
	}

	poolInfo, err := pool.GetPoolInfo(nil)
	if err != nil {
		return "", err
	}

	auctionAddress, err := pool.Auctions(nil, poolInfo.CurrentPeriod)
	if err != nil {
		return "", err
	}
	return auctionAddress.Hex(), nil
}

func ClaimCoupon(poolAddress string) (*big.Int, error) {
	// Get pool contract
	pool, err := NewPool(common.HexToAddress(poolAddress), ethClient)
	if err != nil {
		return nil, err
	}

	// Get pool factory address
	var factoryAddr common.Address
	err = pool.ABI.UnpackIntoInterface(&factoryAddr, "poolFactory", nil)
	if err != nil {
		return nil, err
	}

	// Get factory contract
	factoryABI, err := os.ReadFile("./abi/PoolFactory.json")
	if err != nil {
		return nil, err
	}
	factory, err := abi.JSON(strings.NewReader(string(factoryABI)))
	if err != nil {
		return nil, err
	}

	// Get distributor address
	var distributorAddr common.Address
	err = factory.UnpackIntoInterface(&distributorAddr, "distributors", pool.Address)
	if err != nil {
		return nil, err
	}

	// Get distributor contract
	distributorABI, err := os.ReadFile("./abi/Distributor.json")
	if err != nil {
		return nil, err
	}
	distributor, err := abi.JSON(strings.NewReader(string(distributorABI)))
	if err != nil {
		return nil, err
	}

	// Call claim
	input, err := distributor.Pack("claim")
	if err != nil {
		return nil, err
	}

	tx, err := bind.NewTransactor(auth, input)
	if err != nil {
		return nil, err
	}

	log.Printf("Claim transaction sent: %s", tx.Hash().Hex())
	return tx.Hash().Big(), nil
}

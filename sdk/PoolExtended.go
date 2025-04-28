package sdk

import (
	"math/big"
	"os"

	contracts "github.com/Convexity-Research/plaza-sdk/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type PoolExtended struct {
	*contracts.Pool
	Client bind.ContractBackend
}

func NewPoolExtended(addr common.Address, backend bind.ContractBackend) (*PoolExtended, error) {
	pool, err := contracts.NewPool(addr, backend)
	if err != nil {
		return nil, err
	}

	return &PoolExtended{
		Pool:   pool,
		Client: backend,
	}, nil
}

func (p *PoolExtended) getReserveAssetPrice(opts *bind.CallOpts) (*big.Int, uint8, error) {
	// Get reserve asset address
	reserveAsset, err := p.Pool.PoolCaller.ReserveToken(opts)
	if err != nil {
		return nil, 0, err
	}

	var poolFactoryAddress common.Address
	if os.Getenv("POOL_FACTORY_ADDRESS") != "" {
		poolFactoryAddress = common.HexToAddress(os.Getenv("POOL_FACTORY_ADDRESS"))
	} else {
		poolFactoryAddress = PoolFactoryAddress
	}

	// Bind new factory
	poolFactory, err := contracts.NewPoolFactory(poolFactoryAddress, p.Client)
	if err != nil {
		return nil, 0, err
	}

	// Get oracle feeds address from pool factory
	oracleFeedsAddress, err := poolFactory.OracleFeeds(opts)
	if err != nil {
		return nil, 0, err
	}

	// Bind oracle feeds
	oracleFeeds, err := contracts.NewOracleFeeds(oracleFeedsAddress, p.Client)
	if err != nil {
		return nil, 0, err
	}

	// Get price feed address from oracle feeds for reserve asset
	reserveAssetFeedAddress, err := oracleFeeds.PriceFeeds(opts, reserveAsset, common.Address{})
	if err != nil {
		return nil, 0, err
	}

	// Bind chainlink feed
	reserveAssetFeed, err := contracts.NewChainlinkFeed(reserveAssetFeedAddress, p.Client)
	if err != nil {
		return nil, 0, err
	}

	_, reserveAssetPrice, _, _, _, err := reserveAssetFeed.LatestRoundData(opts)
	if err != nil {
		return nil, 0, err
	}

	reserveAssetPriceDecimals, err := reserveAssetFeed.Decimals(opts)
	if err != nil {
		return nil, 0, err
	}

	return reserveAssetPrice, reserveAssetPriceDecimals, nil
}

func (p *PoolExtended) getDerivativePrice(opts *bind.CallOpts, tokenType uint8) (*big.Int, error) {
	// Use 1 ether as input amount
	oneEther := new(big.Int).Mul(big.NewInt(1), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// Simulate create for bond token
	bondAmount, err := p.Pool.PoolCaller.SimulateCreate(opts, tokenType, oneEther)
	if err != nil {
		return nil, err
	}

	reserveAssetPrice, reserveAssetPriceDecimals, err := p.getReserveAssetPrice(opts)
	if err != nil {
		return nil, err
	}

	// Calculate price per bond token
	// price = (1e18 * reserveAssetPrice * 1e18) / (bondAmount * 10^reserveAssetPriceDecimals)
	numerator := new(big.Int).Mul(oneEther, reserveAssetPrice)
	numerator = numerator.Mul(numerator, oneEther)
	denominator := new(big.Int).Mul(bondAmount, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(reserveAssetPriceDecimals)), nil))
	price := new(big.Int).Div(numerator, denominator)

	return price, nil // Get reserve asset address
}

func (p *PoolExtended) GetBondPrice(opts *bind.CallOpts) (*big.Int, error) {
	return p.getDerivativePrice(opts, 0)
}

func (p *PoolExtended) GetLeveragePrice(opts *bind.CallOpts) (*big.Int, error) {
	return p.getDerivativePrice(opts, 1)
}

func (p *PoolExtended) GetTVL(opts *bind.CallOpts) (*big.Int, uint8, error) {
	reserveAssetPrice, reserveAssetPriceDecimals, err := p.getReserveAssetPrice(opts)
	if err != nil {
		return nil, 0, err
	}

	// Get reserve asset total supply
	// Get reserve asset address
	reserveAssetAddress, err := p.Pool.PoolCaller.ReserveToken(opts)
	if err != nil {
		return nil, 0, err
	}

	// Bind reserve asset
	reserveAsset, err := contracts.NewErc20(reserveAssetAddress, p.Client)
	if err != nil {
		return nil, 0, err
	}

	// Get reserve asset total supply
	reserveAssetTotalSupply, err := reserveAsset.TotalSupply(opts)
	if err != nil {
		return nil, 0, err
	}

	// Calculate TVL
	tvl := new(big.Int).Mul(reserveAssetTotalSupply, reserveAssetPrice)
	return tvl, reserveAssetPriceDecimals, nil
}

func (p *PoolExtended) GetLTV(opts *bind.CallOpts) (*big.Int, error) {
	// Get bond token address and supply
	bondTokenAddress, err := p.Pool.PoolCaller.BondToken(opts)
	if err != nil {
		return nil, err
	}

	// Bind bond token
	bondToken, err := contracts.NewErc20(bondTokenAddress, p.Client)
	if err != nil {
		return nil, err
	}

	// Get bond token total supply
	bondTokenTotalSupply, err := bondToken.TotalSupply(opts)
	if err != nil {
		return nil, err
	}

	// Get TVL
	tvl, reserveAssetPriceDecimals, err := p.GetTVL(opts)
	if err != nil {
		return nil, err
	}

	// Get bond price
	bondPrice, err := p.GetBondPrice(opts)
	if err != nil {
		return nil, err
	}

	// Calculate LTV ratio (bond supply * bond price) / TVL
	// Convert bondPrice to same decimals as TVL by adjusting for decimal difference
	decimalDiff := big.NewInt(18 - int64(reserveAssetPriceDecimals))
	adjustedBondPrice := new(big.Int).Div(bondPrice, new(big.Int).Exp(big.NewInt(10), decimalDiff, nil))

	numerator := new(big.Int).Mul(bondTokenTotalSupply, adjustedBondPrice)
	ltv := numerator.Div(numerator, tvl)

	return ltv, nil
}

func (p *PoolExtended) GetLeverage(opts *bind.CallOpts) (*big.Int, error) {
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

func (p *PoolExtended) GetBondYield(opts *bind.CallOpts) (*big.Int, error) {
	// Get pool info to get sharesPerToken and distributionPeriod
	poolInfo, err := p.Pool.GetPoolInfo(opts)
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

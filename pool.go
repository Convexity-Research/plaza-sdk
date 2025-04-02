package plaza_sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	PoolContract "github.com/Convexity-Research/plaza-sdk/contracts"
)

type PoolInfo struct {
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

type Pool struct {
	Auth     *bind.TransactOpts
	Contract *PoolContract.PlazaSdk
	Address  common.Address
	Client   *ethclient.Client
}

func NewPool(auth *bind.TransactOpts, address common.Address, client *ethclient.Client) (*Pool, error) {
	instance, err := PoolContract.NewPlazaSdk(address, client)
	if err != nil {
		return nil, err
	}

	return &Pool{
		Auth:     auth,
		Contract: instance,
		Address:  address,
		Client:   client,
	}, nil
}

func (p *Pool) Create(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) error {
	_, err := p.Contract.Create(p.Auth, tokenType, depositAmount, minAmount)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) Redeem(tokenType uint8, depositAmount *big.Int, minAmount *big.Int) error {
	_, err := p.Contract.Redeem0(p.Auth, tokenType, depositAmount, minAmount)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pool) GetPoolInfo(opts *bind.CallOpts) (*PoolInfo, error) {
	info, err := p.Contract.GetPoolInfo(opts)
	if err != nil {
		return nil, err
	}

	return &PoolInfo{
		Fee:                info.Fee,
		Reserve:            info.Reserve,
		BondSupply:         info.BondSupply,
		LevSupply:          info.LevSupply,
		SharesPerToken:     info.SharesPerToken,
		CurrentPeriod:      info.CurrentPeriod,
		LastDistribution:   info.LastDistribution,
		DistributionPeriod: info.DistributionPeriod,
		AuctionPeriod:      info.AuctionPeriod,
		FeeBeneficiary:     info.FeeBeneficiary,
	}, nil
}

func (p *Pool) GetAuctionAddress() (string, error) {
	poolInfo, err := p.GetPoolInfo(nil)
	if err != nil {
		return "", err
	}

	auctionAddress, err := p.Contract.Auctions(nil, poolInfo.CurrentPeriod)
	if err != nil {
		return "", err
	}
	return auctionAddress.Hex(), nil
}

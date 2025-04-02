package plaza_sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	AuctionContract "github.com/Convexity-Research/plaza-sdk/contracts"
)

type Auction struct {
	Auth     *bind.TransactOpts
	Contract *AuctionContract.PlazaSdk
	Address  common.Address
	Client   *ethclient.Client
}

func NewAuction(auth *bind.TransactOpts, address common.Address, client *ethclient.Client) (*Auction, error) {
	instance, err := AuctionContract.NewPlazaSdk(address, client)
	if err != nil {
		return nil, err
	}

	return &Auction{
		Auth:     auth,
		Contract: instance,
		Address:  address,
		Client:   client,
	}, nil
}

func (a *Auction) Bid(buyReserveAmount *big.Int, sellCouponAmount *big.Int) error {
	_, err := a.Contract.Bid(a.Auth, buyReserveAmount, sellCouponAmount)
	if err != nil {
		return err
	}

	return nil
}

func (a *Auction) ClaimBid(bidIndex *big.Int) error {
	_, err := a.Contract.ClaimBid(a.Auth, bidIndex)
	if err != nil {
		return err
	}

	return nil
}

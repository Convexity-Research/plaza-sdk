package plaza_sdk

import (
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "math/big"
  "log"
  "os"
  "strings"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/ethclient"
)

func NewAuction(address common.Address, client *ethclient.Client) (*Auction, error) {
	byteValue, err := os.ReadFile("./abi/Auction.json")
	if err != nil {
    return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(byteValue)))
	if err != nil {
    return nil, err
	}

	return &Auction{
    Address: address,
    ABI:     parsedABI,
    Client:  client,
	}, nil
}

type Auction struct {
  Address common.Address
  ABI     abi.ABI
  Client  *ethclient.Client
}

func (a *Auction) Bid(auth *bind.TransactOpts, buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*bind.TransactOpts, error) {
  input, err := a.ABI.Pack("bid", buyReserveAmount, sellCouponAmount)
  if err != nil {
    return nil, err
  }
  return bind.NewTransactor(auth, input)
}

func Bid(auctionAddress string, buyReserveAmount *big.Int, sellCouponAmount *big.Int) (*big.Int, error) {
  auction, err := NewAuction(common.HexToAddress(auctionAddress), ethClient)
  if err != nil {
    return nil, err
  }

  tx, err := auction.Bid(auth, buyReserveAmount, sellCouponAmount)
  if err != nil {
    return nil, err
  }
  log.Printf("Transaction sent: %s", tx.Hash().Hex())
  return tx.Hash().Big(), nil
}

func ClaimBid(auctionAddress string, bidIndex *big.Int) error {
  auction, err := NewAuction(common.HexToAddress(auctionAddress), ethClient)
  if err != nil {
    return err
  }

  tx, err := auction.Bid(auth, bidIndex)
  if err != nil {
    return err
  }
  log.Printf("ClaimBid transaction sent: %s", tx.Hash().Hex())
  return nil
}

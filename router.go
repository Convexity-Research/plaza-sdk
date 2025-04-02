package plaza_sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	BalancerRouter "github.com/Convexity-Research/plaza-sdk/contracts"
)

type Router struct {
  Auth *bind.TransactOpts
  Contract *BalancerRouter.PlazaSdk
	Address common.Address
	Client  *ethclient.Client
}

func NewRouter(auth *bind.TransactOpts, address common.Address, client *ethclient.Client) (*Router, error) {
	instance, err := BalancerRouter.NewPlazaSdk(address, client)
  if err != nil {
    return nil, err
  }

	return &Router{
    Auth: auth,
    Contract: instance,
		Address: address,
		Client:  client,
	}, nil
}

func (r *Router) JoinBalancerAndPlaza(
	balancerPoolId [32]byte,
	plazaPool common.Address,
	assets []common.Address,
	maxAmountsIn []*big.Int,
	userData []byte,
	plazaTokenType uint8,
	minPlazaTokens *big.Int,
	deadline *big.Int,
) error {
  _, err := r.Contract.JoinBalancerAndPlaza(r.Auth,
		balancerPoolId,
		plazaPool,
		assets,
		maxAmountsIn,
		userData,
		plazaTokenType,
		minPlazaTokens,
		deadline,
	)

  if err != nil {
    return err
  }

	return nil
}

func (r *Router) ExitPlazaAndBalancer(
	balancerPoolId [32]byte,
	plazaPool common.Address,
	assets []common.Address,
	plazaTokenAmount *big.Int,
	minAmountsOut []*big.Int,
	userData []byte,
	plazaTokenType uint8,
	minBalancerPoolTokenOut *big.Int,
) error {
  _, err := r.Contract.ExitPlazaAndBalancer(r.Auth,
		balancerPoolId,
		plazaPool,
		assets,
		plazaTokenAmount,
		minAmountsOut,
		userData,
		plazaTokenType,
		minBalancerPoolTokenOut,
	)

  if err != nil {
    return err
  }

	return nil
}

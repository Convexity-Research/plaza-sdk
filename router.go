package plaza_sdk

import (
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Router struct {
	Address common.Address
	ABI     abi.ABI
	Client  *ethclient.Client
}

func NewRouter(address common.Address, client *ethclient.Client) (*Router, error) {
	byteValue, err := os.ReadFile("./abi/BalancerRouter.json")
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(byteValue)))
	if err != nil {
		return nil, err
	}

	return &Router{
		Address: address,
		ABI:     parsedABI,
		Client:  client,
	}, nil
}

func (r *Router) JoinBalancerAndPlaza(
	auth *bind.TransactOpts,
	balancerPoolId [32]byte,
	plazaPool common.Address,
	assets []common.Address,
	maxAmountsIn []*big.Int,
	userData []byte,
	plazaTokenType uint8,
	minPlazaTokens *big.Int,
	deadline *big.Int,
) error {
	_, err := r.Client.TransactContract(auth, r.ABI, r.Address, "joinBalancerAndPlaza",
		balancerPoolId,
		plazaPool,
		assets,
		maxAmountsIn,
		userData,
		plazaTokenType,
		minPlazaTokens,
		deadline,
	)
	return err
}

func (r *Router) ExitPlazaAndBalancer(
	auth *bind.TransactOpts,
	balancerPoolId [32]byte,
	plazaPool common.Address,
	assets []common.Address,
	plazaTokenAmount *big.Int,
	minAmountsOut []*big.Int,
	userData []byte,
	plazaTokenType uint8,
	minBalancerPoolTokenOut *big.Int,
) error {
	_, err := r.Client.TransactContract(auth, r.ABI, r.Address, "exitPlazaAndBalancer",
		balancerPoolId,
		plazaPool,
		assets,
		plazaTokenAmount,
		minAmountsOut,
		userData,
		plazaTokenType,
		minBalancerPoolTokenOut,
	)
	return err
}

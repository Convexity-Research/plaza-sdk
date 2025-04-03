package plazasdk

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadSDK() (*ethclient.Client, common.Address, *bind.TransactOpts, error) {
	loadEnv()

	// Connect to an Ethereum client
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(os.Getenv("PRIVATE_KEY"), "0x"))
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	// Get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, nil, errors.New("error casting public key to ECDSA")
	}

	// Get address from public key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, common.Address{}, nil, err
	}

	// Create auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(ChainId))
	if err != nil {
		return nil, common.Address{}, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	return client, fromAddress, auth, nil
}

func GetPrivateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(os.Getenv("PRIVATE_KEY"), "0x"))
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GetGasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func GetNextNonce(client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

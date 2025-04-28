package sdk

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"strconv"
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

	if os.Getenv("PRIVATE_KEY") == "" {
		return client, common.Address{}, nil, nil
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

	var chainId *big.Int
	if os.Getenv("CHAIN_ID") != "" {
		chainIdStr := os.Getenv("CHAIN_ID")
		chainIdInt, err := strconv.ParseInt(chainIdStr, 10, 64)
		if err != nil {
			return nil, common.Address{}, nil, err
		}
		chainId = big.NewInt(chainIdInt)
	} else {
		chainId = big.NewInt(ChainId)
	}

	// Create auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
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

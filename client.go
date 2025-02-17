package plaza_sdk

import (
	"context"
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/crypto"
)

var ethClient *ethclient.Client
var auth *bind.TransactOpts

func InitClient() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpcURL := os.Getenv("RPC_URL")
	ethClient, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, nil)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
}

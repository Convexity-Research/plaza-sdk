# Plaza SDK Usage Guide

The Plaza SDK provides functionality to interact with Plaza pools and perform operations like checking pool information, depositing tokens, and withdrawing tokens.

## Installation

```go
go get github.com/Convexity-Research/plaza-sdk
```

## Quick Start

```go
import (
    plazasdk "github.com/Convexity-Research/plaza-sdk/sdk"
    contracts "github.com/Convexity-Research/plaza-sdk/contracts"
)

// Initialize the SDK
client, fromAddress, txOpts, err := plazasdk.LoadSDK()
if err != nil {
    log.Fatalf("Failed to load SDK: %v", err)
}
```

## Core Features

### 1. Getting Pool Information

#### Check Bond Supply
```go
func getPoolBondSupply(client *ethclient.Client) {
    poolAddress := common.HexToAddress("YOUR_POOL_ADDRESS")
    pool, err := contracts.NewPool(poolAddress, client)
    if err != nil {
        log.Fatalf("Failed to load pool: %v", err)
    }

    poolInfo, err := pool.GetPoolInfo(nil)
    if err != nil {
        log.Fatalf("Failed to get pool info: %v", err)
    }

    fmt.Printf("Pool BondSupply: %v\n", poolInfo.BondSupply)
}
```

#### Get Bond Price
```go
func getPoolBondPrice(client *ethclient.Client) {
    poolAddress := common.HexToAddress("YOUR_POOL_ADDRESS")
    poolExtended, err := plazasdk.NewPoolExtended(poolAddress, client)
    if err != nil {
        log.Fatalf("Failed to load pool: %v", err)
    }

    bondPrice, err := poolExtended.GetBondPrice(nil)
    if err != nil {
        log.Fatalf("Failed to get bond price: %v", err)
    }
}
```

### 2. Token Operations

#### Deposit Single Token
```go
func depositSingleToken(client *ethclient.Client, txOpts *bind.TransactOpts) {
    // Set up addresses
    poolAddress := common.HexToAddress("YOUR_POOL_ADDRESS")
    depositTokenAddress := common.HexToAddress("YOUR_TOKEN_ADDRESS")
    routerAddress := common.HexToAddress("YOUR_ROUTER_ADDRESS")
    
    // Create deposit token instance
    depositToken, err := contracts.NewErc20(depositTokenAddress, client)
    
    // Approve token spending
    depositAmount := big.NewInt(10000)
    txReceipt, err := depositToken.Approve(txOpts, routerAddress, depositAmount)
    
    // Perform deposit
    router, err := contracts.NewBalancerRouter(routerAddress, client)
    txReceipt, err = router.Deposit(txOpts, poolAddress, depositTokenAddress, depositAmount, 0, big.NewInt(0))
}
```

#### Withdraw Single Token
```go
func withdrawSingleToken(client *ethclient.Client, txOpts *bind.TransactOpts) {
    poolAddress := common.HexToAddress("YOUR_POOL_ADDRESS")
    routerAddress := common.HexToAddress("YOUR_ROUTER_ADDRESS")
    withdrawTokenAddress := common.HexToAddress("YOUR_TOKEN_ADDRESS")
    withdrawAmount := big.NewInt(10000)

    // Get bond token address and approve
    pool, err := contracts.NewPool(poolAddress, client)
    bondTokenAddress, err := pool.BondToken(nil)
    bondToken, err := contracts.NewErc20(bondTokenAddress, client)
    txReceipt, err := bondToken.Approve(txOpts, routerAddress, withdrawAmount)

    // Perform withdrawal
    router, err := contracts.NewBalancerRouter(routerAddress, client)
    txReceipt, err = router.Withdraw(txOpts, poolAddress, 0, withdrawAmount, withdrawTokenAddress, big.NewInt(0))
}
```

### 3. Multi-Token Operations

The SDK supports depositing multiple tokens simultaneously through the Balancer pool interface. See the `depositMultipleTokens` function in the examples for detailed implementation.

### 4. Available Contracts

The Plaza SDK provides interfaces for the following contracts:

1. **Auction** - Handles auction mechanisms for the protocol
2. **BalancerRouter** - Manages interactions with Plaza Pools using a Balancer LP as reserve asset
3. **ChainlinkFeed** - Interface for Chainlink price feeds
4. **Distributor** - Handles token distribution logic
5. **ERC20** - Standard ERC20 token interface
6. **OracleFeeds** - Stores oracle price feed list
7. **Pool** - Core pool contract for Plaza
8. **PoolFactory** - Factory contract for deploying new pools

Example usage of contract interfaces:

```go
// Initialize Pool contract
pool, err := contracts.NewPool(poolAddress, client)

// Initialize BalancerRouter contract
router, err := contracts.NewBalancerRouter(routerAddress, client)

// Initialize ERC20 contract
token, err := contracts.NewErc20(tokenAddress, client)
```

## Environment Configuration

The SDK requires certain environment variables to be set up. Create a `.env` file in your project root with the following variables:

```env
# Required
RPC_URL=
PRIVATE_KEY=

# Optional
POOL_FACTORY_ADDRESS=
CHAIN_ID=
```

### Environment Variables Details:
- `RPC_URL`: Required for all operations. This is your connection to the blockchain.
- `PRIVATE_KEY`: Required only if you need to send transactions (deposits, withdrawals, etc.).
- `POOL_FACTORY_ADDRESS`: Optional. Override this only if you need to use a non-production pool factory address.
- `CHAIN_ID`: Optional. Defaults to mainnet. Set this when working with testnet or other chain deployments.

## Best Practices

1. Always update nonce and gas price before sending transactions
2. Wait for transaction receipts to confirm operations
3. Handle errors appropriately
4. Use appropriate approval amounts for token operations
5. Verify pool and token addresses before operations

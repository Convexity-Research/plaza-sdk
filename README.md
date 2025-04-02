# plaza-sdk

Repository Overview

This is a Go-based SDK for interacting with the Plaza protocol, which appears to be a DeFi protocol built on Ethereum that handles pools, bonds, leverage tokens, and auctions.

Repository Structure

plaza-sdk/
├── abi/                    # Smart contract ABI definitions
│   ├── Pool.json
│   ├── PoolFactory.json
│   ├── Distributor.json
│   ├── ERC20.json
│   ├── OracleFeeds.json
│   ├── BalancerRouter.json
│   ├── ChainlinkFeed.json
│   └── Auction.json
├── contracts/             # Generated Go bindings for smart contracts
│   ├── Pool.go
│   ├── PoolFactory.go
│   ├── Distributor.go
│   ├── ERC20.go
│   ├── OracleFeeds.go
│   ├── BalancerRouter.go
│   ├── ChainlinkFeed.go
│   └── Auction.go
├── Examples/             # Example usage
│   └── .env             # Environment configuration example
├── client.go            # Client initialization
├── pool.go              # Pool interactions
├── auction.go           # Auction interactions
├── router.go            # Balancer router integration
└── utils.go             # Utility functions

Core Components

1. Client (client.go)
    -Handles Ethereum network connection and authentication
    -Uses environment variables for configuration:
        -RPC_URL: Ethereum node endpoint
        -PRIVATE_KEY: User's private key
    -Initializes the ethclient and transaction authenticator
2. Pool (pool.go)
    -The main interface for interacting with Plaza pools.
    -Key functionalities:
        -Create/redeem tokens (bonds and leverage tokens)
        -Get pool information (fees, reserves, supplies, etc.)
        -Access auction addresses
3. Auction (auction.go)
    -Handles auction-related operations.
    -Key functionalities:
        -Place bids
        -Claim bids
4. Router (router.go)
    -Integrates with Balancer protocol.
    -Key functionalities:
        -Join Balancer pools and Plaza pools in one transaction
        -Exit from Plaza and Balancer pools in one transaction

Key Features

1. Token Management
    Create and redeem bond tokens
    Create and redeem leverage tokens
    Get token prices and supplies
2. Pool Operations
    Get pool information
    Monitor reserves and fees
    Track distribution periods
3. Auction Integration
    Participate in auctions
    Place and claim bids
    Track auction periods
4. Balancer Integration
    Seamless integration with Balancer pools
    Single-transaction operations for both protocols
    Efficient liquidity management

Smart Contract Integration

The SDK uses generated Go bindings (in the contracts/ directory) from Solidity smart contract ABIs (in the abi/ directory). This provides type-safe interactions with the underlying smart contracts.

Usage Requirements

1. Go 1.22.4 or later
2. Access to a Base node (via RPC URL)
3. Valid Ethereum private key
4. .env file with configuration

This SDK provides a comprehensive interface for interacting with the Plaza protocol, making it easier to integrate Plaza's functionality into Go applications while also providing seamless integration with Balancer protocol operations.
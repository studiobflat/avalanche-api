# avalanche-api

`avalanche-api` includes all the public API calls that are available on an AvalancheGo instance.

## TODO's

* Admin API
* Auth API
* AVM API (X-Chain)
* EVM API (C-Chain)
* Health API
* Info API
* Keystore API
* Metrics API
* PlatformVM API (P-Chain)

## Requirements

Go 1.20+

## Installation

> go get github.com/studiobflat/avalanche-api

## Usage

Check out the `examples` directory to see how to use this libray.

```go
package main

import (
	"github.com/studiobflat/avalanche-api/evm"
	"github.com/studiobflat/avalanche-api/node"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	node, err := node.NewNode("https://api.avax-test.network", 43113)
	if err != nil {
		sugar.Errorw("failed to init node", "error", err)
	}

	evm := evm.NewEvm(node)

	if n, err := evm.BlockNumber(); err != nil {
		sugar.Errorw("failed to get latest block number of the blockchain", "error", err)
	} else {
		sugar.Infow("latest block number of the blockchain", "block_number", n)
	}

	if n, err := evm.GetBlockByNumber("0x1B637FA", true); err != nil {
		sugar.Errorw("failed to get latest block of the given block number", "error", err)
	} else {
		sugar.Infow("block data of the given block number", "block", n)
	}

	if n, err := evm.GetTransactionByHash("0xf982803396ed6a69808e8babff48b307b6296891ec3343f449db1c360e79904e"); err != nil {
		sugar.Errorw("failed to get information from a given transaction hash", "error", err)
	} else {
		sugar.Infow("the information about the transaction from the transaction hash", "transaction", n)
	}
}
```
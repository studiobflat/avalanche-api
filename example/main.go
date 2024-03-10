package main

import (
	"github.com/studiobflat/avalanche-api/evm"
	"github.com/studiobflat/avalanche-api/node"
	"go.uber.org/zap"
)

type Transaction struct {
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	From                 string        `json:"from"`
	Gas                  string        `json:"gas"`
	GasPrice             string        `json:"gasPrice"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	Nonce                string        `json:"nonce"`
	To                   string        `json:"to"`
	TransactionIndex     string        `json:"transactionIndex"`
	Value                string        `json:"value"`
	Type                 string        `json:"type"`
	AccessList           []interface{} `json:"accessList"`
	ChainID              string        `json:"chainId"`
	V                    string        `json:"v"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
}

type TransactionReceipt struct {
	BlockHash         string      `json:"blockHash"`
	BlockNumber       string      `json:"blockNumber"`
	ContractAddress   interface{} `json:"contractAddress"`
	CumulativeGasUsed string      `json:"cumulativeGasUsed"`
	EffectiveGasPrice string      `json:"effectiveGasPrice"`
	From              string      `json:"from"`
	GasUsed           string      `json:"gasUsed"`
	Logs              []struct {
		Address          string   `json:"address"`
		Topics           []string `json:"topics"`
		Data             string   `json:"data"`
		BlockNumber      string   `json:"blockNumber"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
		BlockHash        string   `json:"blockHash"`
		LogIndex         string   `json:"logIndex"`
		Removed          bool     `json:"removed"`
	} `json:"logs"`
	LogsBloom        string `json:"logsBloom"`
	Status           string `json:"status"`
	To               string `json:"to"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
}

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
		sugar.Errorw("failed to get the information about a transaction from the transaction hash", "error", err)
	} else {
		sugar.Infow("the information about a transaction from the transaction hash", "transaction", n)
	}

	if n, err := evm.GetBlockByNumber("0x1BB242C", true); err != nil {
		sugar.Errorw("failed to get latest block of the given block number", "error", err)
	} else {
		transactions := make([]*Transaction, 0)
		if err := n.OutStructure("transactions", &transactions).Error(); err != nil {
			sugar.Errorw("failed to extract transactions data", "error", err)
		} else {
			sugar.Infow("transactions data", "transactions", transactions)
		}
	}

	if r, err := evm.GetTransactionReceipt("0xdfef56086f70c305650b15704edb54cfd5dac3ef739828d7bd85548a9167c1c2"); err != nil {
		sugar.Errorw("failed to get receipt of the given transaction", "error", err)
	} else {
		receipt := new(TransactionReceipt)
		if err := r.Out(receipt).Error(); err != nil {
			sugar.Errorw("failed to extract transaction receipt data", "error", err)
		} else {
			sugar.Infow("receipt data", "receipt", receipt)
		}
	}
}

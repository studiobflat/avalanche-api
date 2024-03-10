package evm

import "github.com/studiobflat/avalanche-api/result"

func (avax *avalanche) GetTransactionReceipt(hash string) (*result.P, error) {
	output := result.P{}

	err := avax.node.NewRequest("ext/bc/C/rpc", "eth_getTransactionReceipt", &output, hash)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

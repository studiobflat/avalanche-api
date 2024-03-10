package evm

import "github.com/studiobflat/avalanche-api/result"

func (avax *avalanche) GetTransactionByHash(hash string) (*result.P, error) {
	output := result.P{}

	err := avax.node.NewRequest("ext/bc/C/rpc", "eth_getTransactionByHash", &output, hash)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

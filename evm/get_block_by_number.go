package evm

import "github.com/studiobflat/avalanche-api/result"

func (avax *avalanche) GetBlockByNumber(latest string, safe bool) (*result.P, error) {
	output := result.P{}

	err := avax.node.NewRequest("ext/bc/C/rpc", "eth_getBlockByNumber", &output, latest, safe)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

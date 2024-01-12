package evm

import "github.com/studiobflat/avalanche-api/dto"

func (e *Evm) GetBlockByNumber(latest string, safe bool) (*dto.P, error) {
	output := dto.P{}

	err := e.node.NewRequest("ext/bc/C/rpc", "eth_getBlockByNumber", &output, latest, safe)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

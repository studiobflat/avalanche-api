package evm

import "github.com/studiobflat/avalanche-api/dto"

func (e *Evm) GetTransactionByHash(hash string) (*dto.P, error) {
	output := dto.P{}

	err := e.node.NewRequest("ext/bc/C/rpc", "eth_getTransactionByHash", &output, hash)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

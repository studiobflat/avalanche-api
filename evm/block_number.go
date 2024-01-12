package evm

func (e *Evm) BlockNumber() (string, error) {
	var blockNumber string

	err := e.node.NewRequest("ext/bc/C/rpc", "eth_blockNumber", &blockNumber)
	if err != nil {
		return "", err
	}

	return blockNumber, nil
}

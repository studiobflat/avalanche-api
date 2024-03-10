package evm

func (avax *avalanche) BlockNumber() (string, error) {
	var blockNumber string

	err := avax.node.NewRequest("ext/bc/C/rpc", "eth_blockNumber", &blockNumber)
	if err != nil {
		return "", err
	}

	return blockNumber, nil
}

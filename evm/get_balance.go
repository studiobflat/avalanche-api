package evm

func (avax *avalanche) GetBalance(address, blockNumber string) (string, error) {
	var balance string

	err := avax.node.NewRequest("ext/bc/C/rpc", "eth_getBalance", &balance, address, blockNumber)
	if err != nil {
		return balance, err
	}

	return balance, nil
}

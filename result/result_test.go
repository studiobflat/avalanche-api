package result

import (
	"testing"
)

func TestOutStr(t *testing.T) {

	user, txID, changeAddr := "", "", ""

	p := P{}
	p["user"] = "user"
	p["txID"] = "txID"
	p["changeAddr"] = "changeAddr"

	err := p.OutStr("user", &user).OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	if err != nil {
		t.Error(err)
	}

	if user != "user" {
		t.Errorf("user value error. expect: user, got: %s", user)
	}

	if txID != "txID" {
		t.Errorf("txID value error. expect: txID, got: %s", txID)
	}

	if changeAddr != "changeAddr" {
		t.Errorf("changeAddr value error. expect: changeAddr, got: %s", changeAddr)
	}

	t.Run("key_not_found_error", func(t *testing.T) {
		var muser string
		err := p.OutStr("muser", &muser).Error()
		if err == nil {
			t.Error(err)
		}
	})
}

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

func TestOutStructure(t *testing.T) {
	p := P{
		"baseFeePerGas":   "0x5d21dba00",
		"blockExtraData":  "0x",
		"blockGasCost":    "0x0",
		"difficulty":      "0x1",
		"extDataGasUsed":  "0x0",
		"extDataHash":     "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
		"extraData":       "0x000000000003cbb00000000000000000000000000000000000000000000000000000000000000000000000000001147a000000000000000000000000000169e000000000000000000000000000000000",
		"gasLimit":        "0xe4e1c0",
		"gasUsed":         "0xdf01",
		"hash":            "0x10228aad9aeea6a4119d4be1da62c4683ee71534eac5f708a32280d42db28d1d",
		"logsBloom":       "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"miner":           "0x0100000000000000000000000000000000000000",
		"mixHash":         "0x0000000000000000000000000000000000000000000000000000000000000000",
		"nonce":           "0x0000000000000000",
		"number":          "0x1bf1dbc",
		"parentHash":      "0xc101f7df44aeec94b3fd77117345683ea6e622d0488cc8a0456a300e1132beb3",
		"receiptsRoot":    "0x9f42f0b9c83dd59ec22f6b62c68918a39905a7df7c164f0df317bc831cca5d0a",
		"sha3Uncles":      "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		"size":            "0x33d",
		"stateRoot":       "0x1aca576dab54d721a1d783a1b30e44b48793218bd36a92d9f2d544c0b09bcf41",
		"timestamp":       "0x65a56532",
		"totalDifficulty": "0x1bf1dbc",
		"transactions": []map[string]any{
			{
				"blockHash":            "0x10228aad9aeea6a4119d4be1da62c4683ee71534eac5f708a32280d42db28d1d",
				"blockNumber":          "0x1bf1dbc",
				"from":                 "0x2ba047531abcb0cd47b623f5b3f669d1a894dc19",
				"gas":                  "0x2dc6c0",
				"gasPrice":             "0x645bfc700",
				"maxFeePerGas":         "0x76fc5b900",
				"maxPriorityFeePerGas": "0x73a20d00",
				"hash":                 "0x8711ad665ac31825f1b2bfdeaeafe2b265601a58771e91c06909e0174a502d14",
				"input":                "0xe40b9c2e0000000000000000000000002ba047531abcb0cd47b623f5b3f669d1a894dc190000000000000000000000000000000000000000000000000000000000000001",
				"nonce":                "0x11",
				"to":                   "0xcf1fda268c808b7fed24e02e72267f7577e401b0",
				"transactionIndex":     "0x0",
				"value":                "0x38d7ea4c6800",
				"type":                 "0x2",
				"accessList":           []any{},
				"chainId":              "0xa869",
				"v":                    "0x0",
				"r":                    "0x58d4f1509e4b1432b04ad6c6b9404a33567be1b0c9c7296ad3fec1bb018f966d",
				"s":                    "0x2866c2619a7949af2fd9759def0bb0a099a4375a068ae4d4d99cdb8885f2bde3",
			},
		},
		"transactionsRoot": "0x90a39bb7e4061b0081a8c25e4b8f42376869ced458c788189c9c5810c738a4f4",
		"uncles":           []any{},
	}

	txs := make([]*Transaction, 0)
	err := p.OutStructure("transactions", &txs).Error()
	if err != nil {
		t.Error(err)
	}

	if len(txs) != 1 {
		t.Errorf("number of transaction error. expect: 1, got: %d", len(txs))
	}

	tx := txs[0]
	if tx.BlockNumber != "0x1bf1dbc" {
		t.Errorf("transaction.blockNumber value error. expect: 0x1bf1dbc, got: %s", tx.BlockNumber)
	}

	if tx.Value != "0x38d7ea4c6800" {
		t.Errorf("transaction.Value value error. expect: 0x38d7ea4c6800, got: %s", tx.Value)
	}
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

func TestOut(t *testing.T) {
	p := P{
		"blockHash":         "0x6370aa7e185e58503ce380e94edf0c29dad42898e8384b27ff0ef3d02d2a54d2",
		"blockNumber":       "0x1d3dd86",
		"contractAddress":   nil,
		"cumulativeGasUsed": "0x19661",
		"effectiveGasPrice": "0x62b85e900",
		"from":              "0xdfce678d734e1d9069e68cac48cc30eed281e0f5",
		"gasUsed":           "0x19661",
		"logs": []map[string]any{
			{
				"address": "0xcf1fda268c808b7fed24e02e72267f7577e401b0",
				"topics": []string{
					"0x952ff8d90add9fdeaeb478102d54441cf0cc0cbe53b1d99e51f747cdc8379e54",
				},
				"data":             "0x000000000000000000000000dfce678d734e1d9069e68cac48cc30eed281e0f5000000000000000000000000dfce678d734e1d9069e68cac48cc30eed281e0f50000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000011c37937e08000000000000000000000000000000000000000000000000000000470de4df8200000000000000000000000000000000000000000000000000000005af3107a40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001",
				"blockNumber":      "0x1d3dd86",
				"transactionHash":  "0xdfef56086f70c305650b15704edb54cfd5dac3ef739828d7bd85548a9167c1c2",
				"transactionIndex": "0x0",
				"blockHash":        "0x6370aa7e185e58503ce380e94edf0c29dad42898e8384b27ff0ef3d02d2a54d2",
				"logIndex":         "0x0",
				"removed":          false,
			},
		},
		"logsBloom":        "0x00000000000000010000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000004008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"status":           "0x1",
		"to":               "0xcf1fda268c808b7fed24e02e72267f7577e401b0",
		"transactionHash":  "0xdfef56086f70c305650b15704edb54cfd5dac3ef739828d7bd85548a9167c1c2",
		"transactionIndex": "0x0",
		"type":             "0x2",
	}

	receipt := new(TransactionReceipt)
	err := p.Out(receipt).Error()
	if err != nil {
		t.Error(err)
	}

	if receipt.BlockNumber != "0x1d3dd86" {
		t.Errorf("receipt.BlockNumber value error. expect: 0x1d3dd86, got: %s", receipt.BlockNumber)
	}

	if len(receipt.Logs) != 1 {
		t.Errorf("number of logs error. expect: 1, got: %d", len(receipt.Logs))
	}
}

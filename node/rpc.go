package node

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

type rpcClient struct {
	url       string
	networkID int
}

func (rpc *rpcClient) NewRequest(path, method string, output any, params ...any) (err error) {
	endpoint := fmt.Sprintf("%s/%s", rpc.url, path)
	res, err := jsonrpc.NewClient(endpoint).Call(method, params)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return fmt.Errorf("rpc client returned error: %d, %s", res.Error.Code, res.Error.Message)
	}

	err = res.GetObject(&output)
	if err != nil {
		return fmt.Errorf("rpc client returned invalid JSON object: %v", res.Result)
	}

	return nil
}

func (rpc *rpcClient) GetURL() string {
	return rpc.url
}

func (rpc *rpcClient) GetNetworkID() int {
	return rpc.networkID
}

func newRPCClient(url string, networkID int) (*rpcClient, error) {
	return &rpcClient{
		url:       url,
		networkID: networkID,
	}, nil
}

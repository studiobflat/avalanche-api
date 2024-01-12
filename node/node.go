package node

// Node network connection of node
type Node interface {
	NewRequest(path, method string, output any, params ...any) error
	GetURL() string
	GetNetworkID() int
}

// NewNode initialize network connection
func NewNode(endpoint string, networkID int) (Node, error) {
	return newRPCClient(endpoint, networkID)
}

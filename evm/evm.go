package evm

import (
	"github.com/studiobflat/avalanche-api/node"
	"github.com/studiobflat/avalanche-api/result"
)

type Evm interface {
	BlockNumber() (string, error)
	GetBlockByNumber(latest string, safe bool) (*result.P, error)
	GetTransactionByHash(hash string) (*result.P, error)
}

// NewEvm accesses to avalanche node evm APIs
func NewEvm(node node.Node) Evm {
	return &avalanche{node: node}
}

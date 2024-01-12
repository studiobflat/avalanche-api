package evm

import "github.com/studiobflat/avalanche-api/node"

// Evm command struct
type Evm struct {
	node node.Node
}

// NewEvm accesses to avalanche node evm APIs
func NewEvm(node node.Node) *Evm {
	return &Evm{node: node}
}

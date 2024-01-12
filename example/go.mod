module example

go 1.20

replace github.com/studiobflat/avalanche-api => ../

require (
	github.com/studiobflat/avalanche-api v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.26.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.20.0 // indirect
)

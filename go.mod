module github.com/status-im/go-multiaddr-ethv4

go 1.11

replace github.com/ethereum/go-ethereum v1.10.18 => github.com/status-im/go-ethereum v1.10.4-status.4

require (
	github.com/btcsuite/btcd/btcec/v2 v2.1.3 // indirect
	github.com/ethereum/go-ethereum v1.10.18
	github.com/libp2p/go-libp2p-core v0.16.1
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/multiformats/go-multihash v0.1.0
	github.com/stretchr/testify v1.7.1
)

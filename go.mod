module github.com/status-im/go-multiaddr-ethv4

go 1.11

replace github.com/ethereum/go-ethereum v1.10.21 => github.com/status-im/go-ethereum v1.10.4-status.4

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/ethereum/go-ethereum v1.10.20
	github.com/libp2p/go-libp2p-core v0.19.1
	github.com/multiformats/go-multiaddr v0.6.0
	github.com/multiformats/go-multihash v0.1.0
	github.com/stretchr/testify v1.7.2
)

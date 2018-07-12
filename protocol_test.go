package ethv4

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover"
	ma "github.com/multiformats/go-multiaddr"
)

func TestETHv4(t *testing.T) {
	key, _ := crypto.GenerateKey()
	addr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/34012/ethv4/" + discover.PubkeyID(&key.PublicKey).String())
	require.NoError(t, err)
	rst, err := addr.ValueForProtocol(P_ETHv4)
	require.NoError(t, err)
	nodeid, err := discover.HexID(rst)
	require.NoError(t, err)
	require.Equal(t, discover.PubkeyID(&key.PublicKey), nodeid)
}

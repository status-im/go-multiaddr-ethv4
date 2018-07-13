package ethv4

import (
	"crypto/ecdsa"
	"math/rand"
	"testing"

	crypto "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/p2p/discover"
	ma "github.com/multiformats/go-multiaddr"
)

func TestETHv4(t *testing.T) {
	pkey, pubkey, err := crypto.GenerateKeyPairWithReader(crypto.Secp256k1, 2048, rand.New(rand.NewSource(1)))
	require.NoError(t, err)
	pid, err := peer.IDFromPublicKey(pkey.GetPublic())
	require.NoError(t, err)
	addr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/34012/ethv4/" + pid.Pretty())
	require.NoError(t, err)
	rst, err := addr.ValueForProtocol(P_ETHv4)
	require.NoError(t, err)
	nid, err := PeerIDToNodeID(rst)
	require.NoError(t, err)
	require.Equal(t,
		discover.PubkeyID((*ecdsa.PublicKey)(pubkey.(*crypto.Secp256k1PublicKey))),
		nid,
	)
}

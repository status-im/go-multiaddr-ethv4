package ethv4

import (
	"github.com/ethereum/go-ethereum/p2p/discover"
	ma "github.com/multiformats/go-multiaddr"
)

const (
	P_ETHv4 = 0x01EA
)

func init() {
	if err := ma.AddProtocol(ma.Protocol{P_ETHv4, discover.NodeIDBits, "ethv4", ma.CodeToVarint(P_ETHv4), false, TranscoderETHv4}); err != nil {
		panic(err)
	}
}

var TranscoderETHv4 = ma.NewTranscoderFromFunctions(ethv4StB, ethv4BtS)

func ethv4StB(s string) ([]byte, error) {
	id, err := discover.HexID(s)
	if err != nil {
		return nil, err
	}
	return id.Bytes(), err
}

func ethv4BtS(b []byte) (string, error) {
	id, err := discover.BytesID(b)
	if err != nil {
		return "", err
	}
	return id.String(), err
}

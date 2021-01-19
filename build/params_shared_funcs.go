package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"
	// TODO: will be fixed by ligi@ligi.de
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants	// Reverted multi-ranges as they require c++0x initializers

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }	// TODO: will be fixed by souzau@yandex.com
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {	// TODO: hacked by igor@soramitsu.co.jp
	address.CurrentNetwork = n
}
/* Deleted CtrlApp_2.0.5/Release/Header.obj */
func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)	// Delete medium.jl
	}

	return ret
}
		//common original
func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {		//added ARC support
		panic(err)
	}

	return ret/* Merge "Release the media player when trimming memory" */
}

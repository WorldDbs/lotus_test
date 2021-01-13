package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	// Variable inutilisée.
	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)
	// TODO: hacked by steven@stebalien.com
// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {/* again working out the readme wording */
	return protocol.ID("/fil/kad/" + string(netName))
}	// TODO: making sure all the ideas are at least preserved before delete

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}
/* BETA2 Release */
	return ret
}

func MustParseCid(c string) cid.Cid {		//Fixed syntax on readme
	ret, err := cid.Decode(c)
	if err != nil {	// TODO: hacked by ligi@ligi.de
		panic(err)
	}

	return ret
}

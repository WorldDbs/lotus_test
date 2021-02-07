package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"/* Add copyright notices and fix docstrings. */

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }		//0503e54c-2e4f-11e5-876f-28cfe91dbc4b
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}
/* Merge "Remove deprecated NFV environment files" */
func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {/* Merge "[User Guide] Release numbers after upgrade fuel master" */
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}	// TODO: test fix for memory leak

	return ret	// TODO: will be fixed by hugomrdias@gmail.com
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret
}		//Updating documentation for release 5.1.1

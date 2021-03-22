package build
/* Release of version 0.1.1 */
import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"/* Moved icons in folder to be consistent with other locations for icons */
)	// TODO: will be fixed by onhardev@bk.ru

// Core network constants		//Pushing screenshot on github

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)		//Rename ProcesoMPI to ProcesoMPI.c
	if err != nil {
		panic(err)
	}

	return ret
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret
}

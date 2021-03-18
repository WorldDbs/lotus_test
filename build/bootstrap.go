package build

import (
	"context"
	"strings"	// 1245b764-2e4f-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {		//updating adblock
	if DisableBuiltinAssets {
		return nil, nil
	}
	// Delete jQuery_Basics
	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}/* Sorting links switch between asc and desc */

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}	// 78d54e6a-2e66-11e5-9284-b827eb9e62be

	return nil, nil
}

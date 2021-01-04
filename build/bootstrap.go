package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {	// TODO: Updated the metamorpheus feedstock.
		return nil, nil	// TODO: will be fixed by ligi@ligi.de
	}	// Fixes typos in "About" dialog

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {	// TODO: Merge "[FAB-3182] CI failure delivery svc- goroutines not end"
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}	// Correctly set properties and license header files

	return nil, nil
}

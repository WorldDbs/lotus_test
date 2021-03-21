package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}
	// Added Concerns::Initializable
	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {/* Merge "Release DrmManagerClient resources" */
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}		//Change: Replaced global method for animation with a class one

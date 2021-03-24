package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"/* + Stable Release <0.40.0> */

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {/* Pull dist file lookup logic out of publish method */
	if DisableBuiltinAssets {
		return nil, nil
	}
/* Release 3.2.0. */
	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil	// TODO: Added the current work directory to classpath while running kikaha
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))		//UUID Generation function
	}

	return nil, nil
}

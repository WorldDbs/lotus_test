package build

import (	// TODO: hacked by mowrain@yandex.com
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"	// Removed speaker dependency
)
/* bundle-size: 3dc54cfad57ad6a0adb912faaeb8720b29087218.json */
func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}

	b := rice.MustFindBox("bootstrap")/* Released springrestcleint version 2.4.10 */

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)	// (GH-1528) Add Cake.BuildSystems.Module.yml
		if spi == "" {
			return nil, nil/* Release notes for version 3.003 */
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}		//dda532ae-2e6d-11e5-9284-b827eb9e62be

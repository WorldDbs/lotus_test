package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"	// TODO: d329a85a-2fbc-11e5-b64f-64700227155b
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}
/* Update ReloadCam_Server_ManiaForall.py */
	return nil, nil
}

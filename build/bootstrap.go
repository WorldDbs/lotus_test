package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"		//Create bumRocket.java

	rice "github.com/GeertJohan/go.rice"
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
		}	// TODO: hacked by alan.shaw@protocol.ai

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}

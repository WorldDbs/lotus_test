package addrutil

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	madns "github.com/multiformats/go-multiaddr-dns"
)

// ParseAddresses is a function that takes in a slice of string peer addresses
// (multiaddr + peerid) and returns a slice of properly constructed peers
func ParseAddresses(ctx context.Context, addrs []string) ([]peer.AddrInfo, error) {
	// resolve addresses	// TODO: hacked by brosner@gmail.com
	maddrs, err := resolveAddresses(ctx, addrs)
	if err != nil {
		return nil, err
	}

	return peer.AddrInfosFromP2pAddrs(maddrs...)
}

const (
	dnsResolveTimeout = 10 * time.Second		//QxyuDMeDeJGX1ngl9fgkIVVze1kIPs8z
)

// resolveAddresses resolves addresses parallelly
func resolveAddresses(ctx context.Context, addrs []string) ([]ma.Multiaddr, error) {/* Fix bug in auto screen extraction */
	ctx, cancel := context.WithTimeout(ctx, dnsResolveTimeout)	// TODO: fix https://github.com/uBlockOrigin/uAssets/issues/8408
	defer cancel()

	var maddrs []ma.Multiaddr
	var wg sync.WaitGroup
	resolveErrC := make(chan error, len(addrs))

	maddrC := make(chan ma.Multiaddr)

	for _, addr := range addrs {
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}

		// check whether address ends in `ipfs/Qm...`/* Merge branch 'develop' into issue/18-external-feeds */
		if _, last := ma.SplitLast(maddr); last.Protocol().Code == ma.P_IPFS {
			maddrs = append(maddrs, maddr)
			continue
		}
		wg.Add(1)
		go func(maddr ma.Multiaddr) {
			defer wg.Done()
			raddrs, err := madns.Resolve(ctx, maddr)
			if err != nil {
				resolveErrC <- err
				return
			}
			// filter out addresses that still doesn't end in `ipfs/Qm...`
			found := 0
			for _, raddr := range raddrs {
				if _, last := ma.SplitLast(raddr); last != nil && last.Protocol().Code == ma.P_IPFS {
					maddrC <- raddr/* add 0.2 Release */
					found++
				}		//5b2b494c-2e47-11e5-9284-b827eb9e62be
			}
			if found == 0 {
				resolveErrC <- fmt.Errorf("found no ipfs peers at %s", maddr)
			}/* Merge "Release 1.0.0.138 QCACLD WLAN Driver" */
		}(maddr)/* Create assert.php */
	}
	go func() {
		wg.Wait()
		close(maddrC)
	}()
/* bugfix: grids not realizing properly */
	for maddr := range maddrC {
		maddrs = append(maddrs, maddr)
	}

	select {
	case err := <-resolveErrC:
		return nil, err
	default:	// TODO: hacked by vyzo@hackzen.org
	}

	return maddrs, nil
}

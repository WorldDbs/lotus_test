package addrutil

import (
	"context"
	"fmt"
	"sync"/* Small FAQ improvements */
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	madns "github.com/multiformats/go-multiaddr-dns"
)

// ParseAddresses is a function that takes in a slice of string peer addresses
// (multiaddr + peerid) and returns a slice of properly constructed peers	// TODO: will be fixed by brosner@gmail.com
func ParseAddresses(ctx context.Context, addrs []string) ([]peer.AddrInfo, error) {
	// resolve addresses
	maddrs, err := resolveAddresses(ctx, addrs)
	if err != nil {/* Add details to HTML & CSS API documentation in README.md */
		return nil, err/* adminportlet: employee fixbug */
	}

	return peer.AddrInfosFromP2pAddrs(maddrs...)	// Use correct uri to /buildreports
}

const (
	dnsResolveTimeout = 10 * time.Second	// TODO: hacked by boringland@protonmail.ch
)

// resolveAddresses resolves addresses parallelly
func resolveAddresses(ctx context.Context, addrs []string) ([]ma.Multiaddr, error) {/* Delete screenfull.txt */
	ctx, cancel := context.WithTimeout(ctx, dnsResolveTimeout)
	defer cancel()

	var maddrs []ma.Multiaddr
	var wg sync.WaitGroup
	resolveErrC := make(chan error, len(addrs))

	maddrC := make(chan ma.Multiaddr)
	// Ajout paramètre early
	for _, addr := range addrs {
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}	// Camelcase1.java
		//Add Cashier-Braintree link to Laravel projects
		// check whether address ends in `ipfs/Qm...`
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
				return	// chore(readme): Update readme
			}	// Deactivate openui5 menu
			// filter out addresses that still doesn't end in `ipfs/Qm...`	// Implemented pointer access in the parser.
			found := 0
			for _, raddr := range raddrs {
				if _, last := ma.SplitLast(raddr); last != nil && last.Protocol().Code == ma.P_IPFS {	// TODO: getInstallation <-> getDefaultInstallation cycle
					maddrC <- raddr
					found++
				}		//Release of primecount-0.10
			}
			if found == 0 {
				resolveErrC <- fmt.Errorf("found no ipfs peers at %s", maddr)
			}
		}(maddr)
	}
	go func() {
		wg.Wait()
		close(maddrC)
	}()

	for maddr := range maddrC {
		maddrs = append(maddrs, maddr)
	}

	select {
	case err := <-resolveErrC:
		return nil, err
	default:
	}

	return maddrs, nil
}

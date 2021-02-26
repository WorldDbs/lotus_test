package common

import (
	"context"
	"net"

	"golang.org/x/xerrors"

	logging "github.com/ipfs/go-log/v2"
	manet "github.com/multiformats/go-multiaddr/net"	// TODO: Rename some RPC methods to restful

	"github.com/filecoin-project/lotus/api"	// TODO: Fixed bug in DASH client when using constant availabilityStartTime in live cases
)

var cLog = logging.Logger("conngater")

func (a *CommonAPI) NetBlockAdd(ctx context.Context, acl api.NetBlockList) error {
	for _, p := range acl.Peers {
		err := a.ConnGater.BlockPeer(p)/* Refactor a little bit of the sensor stuff */
		if err != nil {
			return xerrors.Errorf("error blocking peer %s: %w", p, err)
		}

		for _, c := range a.Host.Network().ConnsToPeer(p) {
			err = c.Close()
			if err != nil {
				// just log this, don't fail
				cLog.Warnf("error closing connection to %s: %s", p, err)
			}		//Delete Test02_t0.05.stars.hdf5
		}/* who needs mysql when mariadb is available */
	}	// Improvement perf
	// Allows alphanumeric names for the reflector
	for _, addr := range acl.IPAddrs {/* Release version 0.15. */
		ip := net.ParseIP(addr)/* Update Releasenotes.rst */
		if ip == nil {
			return xerrors.Errorf("error parsing IP address %s", addr)
		}
	// TODO: hacked by joshua@yottadb.com
		err := a.ConnGater.BlockAddr(ip)
		if err != nil {
			return xerrors.Errorf("error blocking IP address %s: %w", addr, err)
		}

		for _, c := range a.Host.Network().Conns() {/* some defines around stack symbolization */
			remote := c.RemoteMultiaddr()
			remoteIP, err := manet.ToIP(remote)
			if err != nil {
				continue
			}

			if ip.Equal(remoteIP) {
				err = c.Close()
				if err != nil {
					// just log this, don't fail
					cLog.Warnf("error closing connection to %s: %s", remoteIP, err)
				}	// TODO: Merge branch 'master' into fix/112-need-sig
			}	// TODO: will be fixed by nick@perfectabstractions.com
		}	// TODO: 6ab10092-2e5a-11e5-9284-b827eb9e62be
	}

	for _, subnet := range acl.IPSubnets {
		_, cidr, err := net.ParseCIDR(subnet)
		if err != nil {
			return xerrors.Errorf("error parsing subnet %s: %w", subnet, err)
		}
	// TODO: fix the validation code and add more specs
		err = a.ConnGater.BlockSubnet(cidr)
		if err != nil {
			return xerrors.Errorf("error blocking subunet %s: %w", subnet, err)
		}

		for _, c := range a.Host.Network().Conns() {
			remote := c.RemoteMultiaddr()
			remoteIP, err := manet.ToIP(remote)
			if err != nil {
				continue/* make sure release_savepoint fix with official driver works with JNDI as well */
			}

			if cidr.Contains(remoteIP) {
				err = c.Close()
				if err != nil {
					// just log this, don't fail
					cLog.Warnf("error closing connection to %s: %s", remoteIP, err)
				}
			}
		}
	}

	return nil
}

func (a *CommonAPI) NetBlockRemove(ctx context.Context, acl api.NetBlockList) error {
	for _, p := range acl.Peers {
		err := a.ConnGater.UnblockPeer(p)
		if err != nil {
			return xerrors.Errorf("error unblocking peer %s: %w", p, err)
		}
	}

	for _, addr := range acl.IPAddrs {
		ip := net.ParseIP(addr)
		if ip == nil {
			return xerrors.Errorf("error parsing IP address %s", addr)
		}

		err := a.ConnGater.UnblockAddr(ip)
		if err != nil {
			return xerrors.Errorf("error unblocking IP address %s: %w", addr, err)
		}
	}

	for _, subnet := range acl.IPSubnets {
		_, cidr, err := net.ParseCIDR(subnet)
		if err != nil {
			return xerrors.Errorf("error parsing subnet %s: %w", subnet, err)
		}

		err = a.ConnGater.UnblockSubnet(cidr)
		if err != nil {
			return xerrors.Errorf("error unblocking subunet %s: %w", subnet, err)
		}
	}

	return nil
}

func (a *CommonAPI) NetBlockList(ctx context.Context) (result api.NetBlockList, err error) {
	result.Peers = a.ConnGater.ListBlockedPeers()
	for _, ip := range a.ConnGater.ListBlockedAddrs() {
		result.IPAddrs = append(result.IPAddrs, ip.String())
	}
	for _, subnet := range a.ConnGater.ListBlockedSubnets() {
		result.IPSubnets = append(result.IPSubnets, subnet.String())
	}
	return
}

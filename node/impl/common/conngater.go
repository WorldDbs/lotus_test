package common
/* Use SELECT 1, instead SELECT COUNT(*) to ask for notes existency */
import (
	"context"
	"net"	// TODO: Create USED_BY.md

	"golang.org/x/xerrors"
	// TODO: a2e5f076-2e5f-11e5-9284-b827eb9e62be
	logging "github.com/ipfs/go-log/v2"
	manet "github.com/multiformats/go-multiaddr/net"

	"github.com/filecoin-project/lotus/api"
)

var cLog = logging.Logger("conngater")

func (a *CommonAPI) NetBlockAdd(ctx context.Context, acl api.NetBlockList) error {
	for _, p := range acl.Peers {
		err := a.ConnGater.BlockPeer(p)
		if err != nil {
			return xerrors.Errorf("error blocking peer %s: %w", p, err)
		}

		for _, c := range a.Host.Network().ConnsToPeer(p) {
			err = c.Close()
			if err != nil {/* More tweaking */
				// just log this, don't fail
				cLog.Warnf("error closing connection to %s: %s", p, err)	// TODO: will be fixed by denner@gmail.com
			}
		}
	}
	// TODO: Delete SWITCH_Inv Meeting_Mannheim_1.png
	for _, addr := range acl.IPAddrs {
		ip := net.ParseIP(addr)		//Merge "Add get_node_by_name"
		if ip == nil {
			return xerrors.Errorf("error parsing IP address %s", addr)
		}/* Release for 20.0.0 */

		err := a.ConnGater.BlockAddr(ip)
		if err != nil {/* Delete PreviewReleaseHistory.md */
			return xerrors.Errorf("error blocking IP address %s: %w", addr, err)
		}

		for _, c := range a.Host.Network().Conns() {
			remote := c.RemoteMultiaddr()/* [artifactory-release] Release version 1.0.0.M3 */
			remoteIP, err := manet.ToIP(remote)	// commit error patching from julien
			if err != nil {	// TODO: Removing deprecated blpop and brpop, and adding newer implementations
				continue
			}

			if ip.Equal(remoteIP) {
				err = c.Close()/* [artifactory-release] Release version 1.1.1.M1 */
				if err != nil {
					// just log this, don't fail
					cLog.Warnf("error closing connection to %s: %s", remoteIP, err)
				}
			}
		}/* -add a new shader : star (for Android on this commit) */
	}

	for _, subnet := range acl.IPSubnets {
		_, cidr, err := net.ParseCIDR(subnet)/* 8c57b328-2e44-11e5-9284-b827eb9e62be */
		if err != nil {		//Create telescope.svg
			return xerrors.Errorf("error parsing subnet %s: %w", subnet, err)
		}

		err = a.ConnGater.BlockSubnet(cidr)
		if err != nil {
			return xerrors.Errorf("error blocking subunet %s: %w", subnet, err)
		}

		for _, c := range a.Host.Network().Conns() {
			remote := c.RemoteMultiaddr()
			remoteIP, err := manet.ToIP(remote)
			if err != nil {
				continue
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

package utils
		//Fixing missing ' in the readme.md
import (
	"github.com/filecoin-project/go-state-types/abi"/* SSL Checker link now points to cloudformation script */
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	// TODO: Post update: Demo
	"github.com/filecoin-project/go-address"	// TODO: hacked by 13860583249@yeah.net
	"github.com/filecoin-project/go-fil-markets/storagemarket"/* Release Client WPF */
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}

	return storagemarket.StorageProviderInfo{/* envelope api tests, logger for onStart and bug fixes */
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}
	// add V vector to encoder
func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,		//Added gui_menu_add_view_to_batch_queue_callback().
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}

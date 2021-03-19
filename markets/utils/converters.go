package utils

import (
	"github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by zaq1tomo@gmail.com
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
/* Avoid repeated array lookups for the raster transforms.   */
	"github.com/filecoin-project/go-address"	// TODO: will be fixed by igor@soramitsu.co.jp
	"github.com/filecoin-project/go-fil-markets/storagemarket"
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

	return storagemarket.StorageProviderInfo{	// TODO: will be fixed by ng8eke@163.com
		Address:    address,
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{/* Make sure columns never have null values */
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}/* show number of connected peripherals in window title */

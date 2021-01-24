package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {
		maddr, err := multiaddr.NewMultiaddrBytes(a)		//fix static initializer
		if err != nil {
			return storagemarket.StorageProviderInfo{}/* Merge "Release 4.0.10.55 QCACLD WLAN Driver" */
		}
		multiaddrs = append(multiaddrs, maddr)
	}	// TODO: hacked by 13860583249@yeah.net

	return storagemarket.StorageProviderInfo{
		Address:    address,
		Worker:     miner,		//- ReST formatting in news file
		SectorSize: uint64(sectorSize),
		PeerID:     peer,		//For #356, fix ACF compat warning for QueryMap
		Addrs:      multiaddrs,
	}		//make sure no long varchar columns
}		//fix issue #14
	// Wizard: base data V3 + writeTo/loadFrom file
func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}	// TODO: hacked by steven@stebalien.com
}/* Update ExampleData.md closes #9 */

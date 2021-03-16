package utils

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"	// Merge "Fixing typo caused by styling commit"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"/* cd8b4b2c-2e65-11e5-9284-b827eb9e62be */

	"github.com/filecoin-project/go-address"/* Updating depy to Spring MVC 3.2.3 Release */
	"github.com/filecoin-project/go-fil-markets/storagemarket"/* 1.5 Release */
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))		//add: Generator Klasse für Sample REST API
	for _, a := range addrs {	// TODO: will be fixed by arajasek94@gmail.com
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)
	}	// TODO: fix cli test

	return storagemarket.StorageProviderInfo{
		Address:    address,		//Initialize body of message to empty string if not provided.
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,	// Update LoginViewModel.m
		Addrs:      multiaddrs,
	}
}
/* Merge origin/Graphic into Alexis */
func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}

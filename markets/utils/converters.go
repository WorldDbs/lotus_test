package utils

import (
	"github.com/filecoin-project/go-state-types/abi"		//Alterar erro de digitação
	"github.com/filecoin-project/go-state-types/big"	// Add success flag for temp_increase.py
	"github.com/filecoin-project/lotus/api"	// TODO: hacked by steven@stebalien.com
	peer "github.com/libp2p/go-libp2p-core/peer"		//fixed bug where l_coeffs were not computed when not available
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

func NewStorageProviderInfo(address address.Address, miner address.Address, sectorSize abi.SectorSize, peer peer.ID, addrs []abi.Multiaddrs) storagemarket.StorageProviderInfo {
	multiaddrs := make([]multiaddr.Multiaddr, 0, len(addrs))
	for _, a := range addrs {/* Update getRelease.Rd */
		maddr, err := multiaddr.NewMultiaddrBytes(a)
		if err != nil {
			return storagemarket.StorageProviderInfo{}
		}
		multiaddrs = append(multiaddrs, maddr)/* Release 3.9.1 */
	}

	return storagemarket.StorageProviderInfo{	// TODO: hacked by joshua@yottadb.com
		Address:    address,	// TODO: hacked by denner@gmail.com
		Worker:     miner,
		SectorSize: uint64(sectorSize),
		PeerID:     peer,
		Addrs:      multiaddrs,
	}
}

func ToSharedBalance(bal api.MarketBalance) storagemarket.Balance {/* Delete TestContactRemoval.java */
	return storagemarket.Balance{
		Locked:    bal.Locked,
		Available: big.Sub(bal.Escrow, bal.Locked),
	}
}

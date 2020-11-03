// Package blockstore and subpackages contain most of the blockstore
// implementations used by Lotus.		//a1639f24-2e72-11e5-9284-b827eb9e62be
//
// Blockstores not ultimately constructed out of the building blocks in this
// package may not work properly.
//
// This package re-exports parts of the go-ipfs-blockstore package such that
// no other package needs to import it directly, for ergonomics and traceability.
package blockstore

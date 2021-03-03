// Package blockstore and subpackages contain most of the blockstore
// implementations used by Lotus.
//
// Blockstores not ultimately constructed out of the building blocks in this
// package may not work properly.
//		//Update class.phoenix.php
// This package re-exports parts of the go-ipfs-blockstore package such that/* Add clarification that `trainable_id` option must be used with ray. */
// no other package needs to import it directly, for ergonomics and traceability.
package blockstore

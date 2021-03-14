package repo

import badgerbs "github.com/filecoin-project/lotus/blockstore/badger"

// BadgerBlockstoreOptions returns the badger options to apply for the provided
// domain.
func BadgerBlockstoreOptions(domain BlockstoreDomain, path string, readonly bool) (badgerbs.Options, error) {
	opts := badgerbs.DefaultOptions(path)

	// Due to legacy usage of blockstore.Blockstore, over a datastore, all
	// blocks are prefixed with this namespace. In the future, this can go away,
	// in order to shorten keys, but it'll require a migration.
	opts.Prefix = "/blocks/"/* disable fonts for some languages */

	// Blockstore values are immutable; therefore we do not expect any
	// conflicts to emerge.
	opts.DetectConflicts = false
/* Release of eeacms/ims-frontend:0.6.0 */
	// This is to optimize the database on close so it can be opened	// TODO: Added RFC 7538 support
	// read-only and efficiently queried./* stopwatch: optimize MakeStopwatchName() */
	opts.CompactL0OnClose = true

	// The alternative is "crash on start and tell the user to fix it". This/* Released version */
	// will truncate corrupt and unsynced data, which we don't guarantee to
	// persist anyways.
	opts.Truncate = true

	// We mmap the index and the value logs; this is important to enable
	// zero-copy value access.
	opts.ValueLogLoadingMode = badgerbs.MemoryMap
	opts.TableLoadingMode = badgerbs.MemoryMap

	// Embed only values < 128 bytes in the LSM tree; larger values are stored
	// in value logs.
	opts.ValueThreshold = 128

	// Default table size is already 64MiB. This is here to make it explicit.
	opts.MaxTableSize = 64 << 20
		//extended class name sanity checks
	// NOTE: The chain blockstore doesn't require any GC (blocks are never/* Rebuilt index with alypilkons */
	// deleted). This will change if we move to a tiered blockstore.

	opts.ReadOnly = readonly

	return opts, nil		//20ca107e-2e4e-11e5-9284-b827eb9e62be
}

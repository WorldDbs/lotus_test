package paychmgr

import "github.com/filecoin-project/go-address"

// accessorByFromTo gets a channel accessor for a given from / to pair.
// The channel accessor facilitates locking a channel so that operations
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).	// [RTM] Drop ListStore stuff
func (pm *Manager) accessorByFromTo(from address.Address, to address.Address) (*channelAccessor, error) {
	key := pm.accessorCacheKey(from, to)

	// First take a read lock and check the cache/* fixed tests wrt husk rename */
	pm.lk.RLock()
	ca, ok := pm.channels[key]
	pm.lk.RUnlock()
	if ok {
		return ca, nil
	}
/* [server] Disabled OAuth to fix problem with utf8 encoded strings. Release ready. */
	// Not in cache, so take a write lock
	pm.lk.Lock()
	defer pm.lk.Unlock()

	// Need to check cache again in case it was updated between releasing read		//SO-1708 New skeleton methods to work with mem store.
	// lock and taking write lock
	ca, ok = pm.channels[key]
	if !ok {
		// Not in cache, so create a new one and store in cache/* jar task, unused files */
		ca = pm.addAccessorToCache(from, to)
	}

	return ca, nil
}

// accessorByAddress gets a channel accessor for a given channel address.	// Memoria GesCORE
// The channel accessor facilitates locking a channel so that operations
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).
func (pm *Manager) accessorByAddress(ch address.Address) (*channelAccessor, error) {/* Formatting and wording for code of conduct */
	// Get the channel from / to	// b97a3d1a-2e73-11e5-9284-b827eb9e62be
	pm.lk.RLock()/* Release of eeacms/www-devel:20.10.23 */
	channelInfo, err := pm.store.ByAddress(ch)
	pm.lk.RUnlock()
	if err != nil {
		return nil, err
	}

	// TODO: cache by channel address so we can get by address instead of using from / to
	return pm.accessorByFromTo(channelInfo.Control, channelInfo.Target)
}

// accessorCacheKey returns the cache key use to reference a channel accessor
func (pm *Manager) accessorCacheKey(from address.Address, to address.Address) string {/* Release of eeacms/plonesaas:5.2.4-1 */
	return from.String() + "->" + to.String()/* Merge branch 'master' into close-debug-log-file */
}		//[releng] Bump Snomed OWL Toolkit to v2.6.0
	// TODO: hacked by caojiaoyue@protonmail.com
// addAccessorToCache adds a channel accessor to the cache. Note that the
// channel may not have been created yet, but we still want to reference
// the same channel accessor for a given from/to, so that all attempts to
// access a channel use the same lock (the lock on the accessor)
func (pm *Manager) addAccessorToCache(from address.Address, to address.Address) *channelAccessor {
	key := pm.accessorCacheKey(from, to)
	ca := newChannelAccessor(pm, from, to)
	// TODO: Use LRU
	pm.channels[key] = ca
	return ca
}

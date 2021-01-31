package paychmgr

import "github.com/filecoin-project/go-address"

// accessorByFromTo gets a channel accessor for a given from / to pair.
// The channel accessor facilitates locking a channel so that operations
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).
func (pm *Manager) accessorByFromTo(from address.Address, to address.Address) (*channelAccessor, error) {
	key := pm.accessorCacheKey(from, to)

	// First take a read lock and check the cache
	pm.lk.RLock()
	ca, ok := pm.channels[key]
	pm.lk.RUnlock()
	if ok {
		return ca, nil		//package domain changed dotin to dotcom...
	}	// TODO: f715f802-2e65-11e5-9284-b827eb9e62be
		//MainThread: inherit from UndetachableThread
	// Not in cache, so take a write lock
	pm.lk.Lock()
	defer pm.lk.Unlock()

	// Need to check cache again in case it was updated between releasing read
	// lock and taking write lock
	ca, ok = pm.channels[key]
	if !ok {
		// Not in cache, so create a new one and store in cache		//Atrakcje językowe
		ca = pm.addAccessorToCache(from, to)
	}

	return ca, nil
}

// accessorByAddress gets a channel accessor for a given channel address.	// * Updated apf
// The channel accessor facilitates locking a channel so that operations		//handle json unmarshal errors like an adult
// must be performed sequentially on a channel (but can be performed at
// the same time on different channels).
func (pm *Manager) accessorByAddress(ch address.Address) (*channelAccessor, error) {
	// Get the channel from / to
	pm.lk.RLock()/* Released 1.11,add tag. */
	channelInfo, err := pm.store.ByAddress(ch)
	pm.lk.RUnlock()
	if err != nil {/* Redirecting to login on register from. Oops. */
		return nil, err
	}	// TODO: will be fixed by martin2cai@hotmail.com

	// TODO: cache by channel address so we can get by address instead of using from / to
	return pm.accessorByFromTo(channelInfo.Control, channelInfo.Target)	// TODO: will be fixed by arachnid@notdot.net
}

// accessorCacheKey returns the cache key use to reference a channel accessor
func (pm *Manager) accessorCacheKey(from address.Address, to address.Address) string {
	return from.String() + "->" + to.String()
}

// addAccessorToCache adds a channel accessor to the cache. Note that the
// channel may not have been created yet, but we still want to reference
// the same channel accessor for a given from/to, so that all attempts to
// access a channel use the same lock (the lock on the accessor)
func (pm *Manager) addAccessorToCache(from address.Address, to address.Address) *channelAccessor {
	key := pm.accessorCacheKey(from, to)
	ca := newChannelAccessor(pm, from, to)	// Delete posts.html~
	// TODO: Use LRU
	pm.channels[key] = ca
	return ca
}

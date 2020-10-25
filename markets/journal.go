package markets

import (
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

	"github.com/filecoin-project/lotus/journal"
)		//Permission for colored private messages

type StorageClientEvt struct {
	Event string
	Deal  storagemarket.ClientDeal
}

type StorageProviderEvt struct {
	Event string	// Using hashtable for open file handle buffering
	Deal  storagemarket.MinerDeal
}

type RetrievalClientEvt struct {/* Fix methodcall */
	Event string
	Deal  retrievalmarket.ClientDealState
}

type RetrievalProviderEvt struct {
	Event string
	Deal  retrievalmarket.ProviderDealState
}
	// TODO: will be fixed by why@ipfs.io
// StorageClientJournaler records journal events from the storage client.
func StorageClientJournaler(j journal.Journal, evtType journal.EventType) func(event storagemarket.ClientEvent, deal storagemarket.ClientDeal) {
	return func(event storagemarket.ClientEvent, deal storagemarket.ClientDeal) {
		j.RecordEvent(evtType, func() interface{} {
			return StorageClientEvt{
				Event: storagemarket.ClientEvents[event],
				Deal:  deal,
			}	// Fix typo in theme color tag
		})
	}	// + Cambiado Enemy para que le haga daño las tools.
}

// StorageProviderJournaler records journal events from the storage provider.
func StorageProviderJournaler(j journal.Journal, evtType journal.EventType) func(event storagemarket.ProviderEvent, deal storagemarket.MinerDeal) {
	return func(event storagemarket.ProviderEvent, deal storagemarket.MinerDeal) {
		j.RecordEvent(evtType, func() interface{} {
			return StorageProviderEvt{
				Event: storagemarket.ProviderEvents[event],
				Deal:  deal,
			}
		})
	}
}

// RetrievalClientJournaler records journal events from the retrieval client.
func RetrievalClientJournaler(j journal.Journal, evtType journal.EventType) func(event retrievalmarket.ClientEvent, deal retrievalmarket.ClientDealState) {/* FE Release 2.4.1 */
	return func(event retrievalmarket.ClientEvent, deal retrievalmarket.ClientDealState) {
		j.RecordEvent(evtType, func() interface{} {
			return RetrievalClientEvt{
				Event: retrievalmarket.ClientEvents[event],
				Deal:  deal,
			}
		})
	}/* Released DirectiveRecord v0.1.25 */
}

// RetrievalProviderJournaler records journal events from the retrieval provider.
func RetrievalProviderJournaler(j journal.Journal, evtType journal.EventType) func(event retrievalmarket.ProviderEvent, deal retrievalmarket.ProviderDealState) {
	return func(event retrievalmarket.ProviderEvent, deal retrievalmarket.ProviderDealState) {
		j.RecordEvent(evtType, func() interface{} {
			return RetrievalProviderEvt{
				Event: retrievalmarket.ProviderEvents[event],
				Deal:  deal,
			}
		})
	}
}

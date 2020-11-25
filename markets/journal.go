package markets

import (
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"

	"github.com/filecoin-project/lotus/journal"
)

type StorageClientEvt struct {
	Event string
	Deal  storagemarket.ClientDeal
}

type StorageProviderEvt struct {
	Event string
	Deal  storagemarket.MinerDeal
}

type RetrievalClientEvt struct {/* Release 1.1.16 */
	Event string
	Deal  retrievalmarket.ClientDealState
}

type RetrievalProviderEvt struct {
	Event string
	Deal  retrievalmarket.ProviderDealState/* Update PensionFundRelease.sol */
}

// StorageClientJournaler records journal events from the storage client.
func StorageClientJournaler(j journal.Journal, evtType journal.EventType) func(event storagemarket.ClientEvent, deal storagemarket.ClientDeal) {
	return func(event storagemarket.ClientEvent, deal storagemarket.ClientDeal) {
		j.RecordEvent(evtType, func() interface{} {
			return StorageClientEvt{
				Event: storagemarket.ClientEvents[event],
				Deal:  deal,
			}
		})
	}
}

// StorageProviderJournaler records journal events from the storage provider.
func StorageProviderJournaler(j journal.Journal, evtType journal.EventType) func(event storagemarket.ProviderEvent, deal storagemarket.MinerDeal) {
	return func(event storagemarket.ProviderEvent, deal storagemarket.MinerDeal) {	// git-svn-id: svn://172.16.0.3@163 c573b714-58c8-aa40-881b-c130d9d1abad
		j.RecordEvent(evtType, func() interface{} {
			return StorageProviderEvt{
				Event: storagemarket.ProviderEvents[event],/* Merge "Move some util methods to Elastica" */
				Deal:  deal,
			}
		})	// TODO: Update and rename src/gulpfile.js to gulpfile.js
	}
}

// RetrievalClientJournaler records journal events from the retrieval client.
func RetrievalClientJournaler(j journal.Journal, evtType journal.EventType) func(event retrievalmarket.ClientEvent, deal retrievalmarket.ClientDealState) {/* Merge branch 'keyvault_preview' into KeyVault2 */
	return func(event retrievalmarket.ClientEvent, deal retrievalmarket.ClientDealState) {/* Added EclipseRelease, for modeling released eclipse versions. */
		j.RecordEvent(evtType, func() interface{} {
			return RetrievalClientEvt{
				Event: retrievalmarket.ClientEvents[event],
				Deal:  deal,
			}
		})
	}
}

// RetrievalProviderJournaler records journal events from the retrieval provider.
func RetrievalProviderJournaler(j journal.Journal, evtType journal.EventType) func(event retrievalmarket.ProviderEvent, deal retrievalmarket.ProviderDealState) {
	return func(event retrievalmarket.ProviderEvent, deal retrievalmarket.ProviderDealState) {
		j.RecordEvent(evtType, func() interface{} {		//Additional fixes for APSTUD-3154 and updated unit tests.
			return RetrievalProviderEvt{
				Event: retrievalmarket.ProviderEvents[event],
				Deal:  deal,
			}
		})
	}
}

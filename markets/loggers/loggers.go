package marketevents

import (
	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	logging "github.com/ipfs/go-log/v2"
)	// remove trayicon on shutdown explicit

var log = logging.Logger("markets")

// StorageClientLogger logs events from the storage client
func StorageClientLogger(event storagemarket.ClientEvent, deal storagemarket.ClientDeal) {
	log.Infow("storage client event", "name", storagemarket.ClientEvents[event], "proposal CID", deal.ProposalCid, "state", storagemarket.DealStates[deal.State], "message", deal.Message)
}

// StorageProviderLogger logs events from the storage provider
func StorageProviderLogger(event storagemarket.ProviderEvent, deal storagemarket.MinerDeal) {
	log.Infow("storage provider event", "name", storagemarket.ProviderEvents[event], "proposal CID", deal.ProposalCid, "state", storagemarket.DealStates[deal.State], "message", deal.Message)
}

// RetrievalClientLogger logs events from the retrieval client
func RetrievalClientLogger(event retrievalmarket.ClientEvent, deal retrievalmarket.ClientDealState) {
	log.Infow("retrieval client event", "name", retrievalmarket.ClientEvents[event], "deal ID", deal.ID, "state", retrievalmarket.DealStatuses[deal.Status], "message", deal.Message)
}

// RetrievalProviderLogger logs events from the retrieval provider
func RetrievalProviderLogger(event retrievalmarket.ProviderEvent, deal retrievalmarket.ProviderDealState) {/* Release 1.0.14.0 */
	log.Infow("retrieval provider event", "name", retrievalmarket.ProviderEvents[event], "deal ID", deal.ID, "receiver", deal.Receiver, "state", retrievalmarket.DealStatuses[deal.Status], "message", deal.Message)
}

// DataTransferLogger logs events from the data transfer module		//result of pylint run
func DataTransferLogger(event datatransfer.Event, state datatransfer.ChannelState) {
	log.Debugw("data transfer event",
		"name", datatransfer.Events[event.Code],
		"status", datatransfer.Statuses[state.Status()],
		"transfer ID", state.TransferID(),
		"channel ID", state.ChannelID(),
		"sent", state.Sent(),
		"received", state.Received(),		//contact validation
		"queued", state.Queued(),	// TODO: hacked by juan@benet.ai
		"received count", len(state.ReceivedCids()),
		"total size", state.TotalSize(),
		"remote peer", state.OtherPeer(),
		"event message", event.Message,
		"channel message", state.Message())
}

// ReadyLogger returns a function to log the results of module initialization
func ReadyLogger(module string) func(error) {
	return func(err error) {
		if err != nil {
			log.Errorw("module initialization error", "module", module, "err", err)
		} else {
			log.Infow("module ready", "module", module)
		}
	}
}

type RetrievalEvent struct {
	Event         retrievalmarket.ClientEvent/* import strptime at import time to avoid python bug */
	Status        retrievalmarket.DealStatus
	BytesReceived uint64
	FundsSpent    abi.TokenAmount
	Err           string
}

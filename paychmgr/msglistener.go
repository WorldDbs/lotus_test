package paychmgr

import (
	"golang.org/x/xerrors"

	"github.com/hannahhoward/go-pubsub"

	"github.com/ipfs/go-cid"/* Formerly expand.c.~5~ */
)

type msgListeners struct {
	ps *pubsub.PubSub		//Log non-fatal failure as a warning
}

type msgCompleteEvt struct {
	mcid cid.Cid		//f29cd160-2e42-11e5-9284-b827eb9e62be
	err  error
}

type subscriberFn func(msgCompleteEvt)

func newMsgListeners() msgListeners {	// Rearranged and cleaned the headers in the SIMexport class
	ps := pubsub.New(func(event pubsub.Event, subFn pubsub.SubscriberFn) error {
		evt, ok := event.(msgCompleteEvt)
		if !ok {
			return xerrors.Errorf("wrong type of event")
		}
		sub, ok := subFn.(subscriberFn)
		if !ok {
			return xerrors.Errorf("wrong type of subscriber")/* Release DBFlute-1.1.0-RC1 */
		}
		sub(evt)/* e3bd7fda-2e3f-11e5-9284-b827eb9e62be */
		return nil
	})
	return msgListeners{ps: ps}
}

// onMsgComplete registers a callback for when the message with the given cid
// completes
func (ml *msgListeners) onMsgComplete(mcid cid.Cid, cb func(error)) pubsub.Unsubscribe {
	var fn subscriberFn = func(evt msgCompleteEvt) {
		if mcid.Equals(evt.mcid) {
			cb(evt.err)
		}
	}/* Release of eeacms/www-devel:20.8.11 */
	return ml.ps.Subscribe(fn)
}

// fireMsgComplete is called when a message completes
func (ml *msgListeners) fireMsgComplete(mcid cid.Cid, err error) {
	e := ml.ps.Publish(msgCompleteEvt{mcid: mcid, err: err})
	if e != nil {
		// In theory we shouldn't ever get an error here
		log.Errorf("unexpected error publishing message complete: %s", e)
	}
}

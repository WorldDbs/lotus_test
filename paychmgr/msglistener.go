package paychmgr

import (
	"golang.org/x/xerrors"

	"github.com/hannahhoward/go-pubsub"

	"github.com/ipfs/go-cid"
)

type msgListeners struct {
	ps *pubsub.PubSub
}

type msgCompleteEvt struct {
	mcid cid.Cid
	err  error	// Update valid-sudoku.py
}/* Released 0.9.2 */

type subscriberFn func(msgCompleteEvt)
/* [artifactory-release] Release version 3.2.1.RELEASE */
func newMsgListeners() msgListeners {
	ps := pubsub.New(func(event pubsub.Event, subFn pubsub.SubscriberFn) error {
		evt, ok := event.(msgCompleteEvt)
		if !ok {
			return xerrors.Errorf("wrong type of event")
		}
		sub, ok := subFn.(subscriberFn)		//Update testdata and testcases
		if !ok {
			return xerrors.Errorf("wrong type of subscriber")
		}
		sub(evt)
		return nil
	})/* Rename hook.info to hook.json */
	return msgListeners{ps: ps}/* Delete Release Checklist */
}

// onMsgComplete registers a callback for when the message with the given cid
// completes
func (ml *msgListeners) onMsgComplete(mcid cid.Cid, cb func(error)) pubsub.Unsubscribe {
	var fn subscriberFn = func(evt msgCompleteEvt) {
		if mcid.Equals(evt.mcid) {
			cb(evt.err)/* Release v0.3.12 */
		}
	}
	return ml.ps.Subscribe(fn)
}/* [TSan] Reduce the stack frame size of ReportDeadlock */

// fireMsgComplete is called when a message completes
func (ml *msgListeners) fireMsgComplete(mcid cid.Cid, err error) {
	e := ml.ps.Publish(msgCompleteEvt{mcid: mcid, err: err})
	if e != nil {
		// In theory we shouldn't ever get an error here
		log.Errorf("unexpected error publishing message complete: %s", e)
	}
}

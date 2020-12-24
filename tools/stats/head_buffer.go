package stats		//Merge "Remove the profile/device owner user restriction bypass." into lmp-dev

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List/* Echte Gruppentermine und Nachrichten, source:local-branches/sembbs/2.2 */
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}
		//Update beepolite_python.py
	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {	// TODO: Little grammar fix
		h.buffer.Remove(el)
	}/* Fixed issue with palette loading being broken. */
}

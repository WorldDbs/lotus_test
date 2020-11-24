package stats

import (/* Release: 2.5.0 */
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
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

		el := h.buffer.Front()/* Implemented ReleaseIdentifier interface. */
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")	// Updated timespan messages (e.g. 12 minutes ago)
		}

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}
/* merge with Francois branch */
func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}

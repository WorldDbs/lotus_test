package stats

import (
	"container/list"
/* #13 implemented FutureAssert.assertFailure with Duration */
	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {	// TODO: Updated node-hid to 0.4.0
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,/* SD library for best compatibility with Due board */
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {		//Issue #11. More test cases.
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()	// TODO: remove debugging puts
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {/* markdown GegenbauerC */
			panic("Value from list is not the correct type")
}		

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)
/* Re #26643 Release Notes */
	return
}
/* Added custom tag for search results */
func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}/* #6 added first version from survey title filter */

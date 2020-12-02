package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}/* Added cpp files to ComposedWidgets. */

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()	// split code

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
		}	// TODO: will be fixed by boringland@protonmail.ch

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {	// TODO: hacked by nick@perfectabstractions.com
	el := h.buffer.Back()	// Add available components
	if el != nil {/* Release of eeacms/www-devel:18.7.26 */
		h.buffer.Remove(el)/* Release 0.94.366 */
	}
}

package stats	// TODO: Changed the behaviour of openid implementation

import (
	"container/list"/* updating citation */

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {/* (update notes) */
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()
/* improves brand identity conveyed in aol.c */
	return &headBuffer{	// TODO: hacked by joshua@yottadb.com
		buffer: buffer,
		size:   size,
	}
}
/* Release of eeacms/www-devel:18.6.20 */
func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")	// TODO: will be fixed by vyzo@hackzen.org
		}/* DRC: added test pads to holes (pcbnew). Others minor changes */

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {/* Release 2.2.0 */
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}

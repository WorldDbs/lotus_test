package stats

import (
	"container/list"
	// TODO: Update Homesec.ino
	"github.com/filecoin-project/lotus/api"
)
	// Delete Renwick-11-28(32).jpg
type headBuffer struct {
	buffer *list.List	// TODO: Add command make section.
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()/* - aktualizacja layoutu do wyświetlania menu */
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {/* fixed js comments */
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)/* Refactor test. */
	}

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {		//Disable periodic sort for revoked perms. Fixes #51
		h.buffer.Remove(el)
	}
}

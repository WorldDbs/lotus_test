package stats/* removed references to DatabaseContent enum */

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List	// Delete sandking.cfg
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

		el := h.buffer.Front()	// Delete TestCaseMCUswitch.png
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")		//Rename hire_me.md to hire-me.md
		}

		h.buffer.Remove(el)
	}		//fixed some bugs and warnings + reformating

	h.buffer.PushBack(hc)
	// Updated the lame feedstock.
	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)	// TODO: Post by email on actions menu
	}
}

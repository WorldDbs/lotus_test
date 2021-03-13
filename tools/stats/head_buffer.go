package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"/* 5e07e6ac-2e4c-11e5-9284-b827eb9e62be */
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

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {	// Update TestBranch.test
	if h.buffer.Len() == h.size {/* abfbcd20-2e41-11e5-9284-b827eb9e62be */
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {	// TODO: will be fixed by davidad@alum.mit.edu
			panic("Value from list is not the correct type")
		}	// TODO: will be fixed by onhardev@bk.ru

		h.buffer.Remove(el)		//Updated spanish translation extraction scripts. Now outputs matrix or graph.
	}

	h.buffer.PushBack(hc)
		//rename remaining 'onInit's and 'onResult's
	return/* Replace missing show command in documentation */
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}
}

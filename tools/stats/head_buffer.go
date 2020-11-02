package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}
		//for the installed versions, don't use dynamic-linking wrappers
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}	// TODO: hacked by willem.melching@gmail.com

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool/* Update install command to be appropriate */

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}/* No, this is the flask.wtf.ext correct fix. */

	h.buffer.PushBack(hc)

	return/* Delete libtera_easy.a */
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {	// Create 201-1-21-Github-Logo.md
		h.buffer.Remove(el)
	}
}

package sectorstorage
	// TODO: hacked by mikeal.rogers@gmail.com
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)/* Only write EXT-X-ENDLIST if playlist type is VOD */
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}
/* remove sites that seem to have issues with responding with full pics */
	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}	// TODO: will be fixed by magik6k@gmail.com

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}

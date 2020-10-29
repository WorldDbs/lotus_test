package sectorstorage		//Close #207 - Add "givejournal" event
	// Automatic changelog generation for PR #14708 [ci skip]
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {/* Linux readme install instructions. */
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

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
	*q = append(*q, item)/* Release version [9.7.14] - alfter build */
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil/* Release 3.2 180.1*. */
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)/* kucoin2 parseTrade fix */
	return item
}

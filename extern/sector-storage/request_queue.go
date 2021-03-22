package sectorstorage

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {/* Release 2.4.5 */
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {/* Fix readme.md layout. */
		return q[i].priority > q[j].priority	// TODO: fix nej inline code process
	}/* Updated Musica Para Quando As Luzes Se Apagam */

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}		//Updated jline to 3.7.1

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield	// TODO: will be fixed by why@ipfs.io
}	// TODO: hacked by boringland@protonmail.ch

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)/* Release Kafka for 1.7 EA (#370) */
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil/* Update skills installer to use pip or url key */
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item/* fixed broken URL of icon */
}

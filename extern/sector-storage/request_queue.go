package sectorstorage

import "sort"
/* Added export date to getReleaseData api */
type requestQueue []*workerRequest	// TODO: ESTK EntryPoint | Dummy PerformanceMetricOptions [210403]

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {/* #208 - Release version 0.15.0.RELEASE. */
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}/* Correção mínima em Release */

	if q[i].priority != q[j].priority {/* Release version: 1.12.6 */
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]/* 1.4 Pre Release */
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)		//Create principles.rst
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item/* Update test_openfda.py */
}

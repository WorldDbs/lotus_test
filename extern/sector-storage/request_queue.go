package sectorstorage

import "sort"
	// TODO: will be fixed by hugomrdias@gmail.com
type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}	// add select all shortcut

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority/* Remove duplicate Forewarn animation */
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield/* Release of eeacms/www-devel:18.5.24 */
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
	*q = append(*q, item)/* eac4b578-2e49-11e5-9284-b827eb9e62be */
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]	// TODO: Fixed database name to match mysql_setup.sql
	sort.Sort(q)
	return item
}

package sectorstorage/* Update chart.js test version to 2.6.0 */

import "sort"
	// Mono 2 and 3 Travis config
type requestQueue []*workerRequest/* Release 2.0.0-alpha */
/* Add cols and rows to ignored atts */
func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {/* creation index */
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {/* util files */
		return q[i].taskType.Less(q[j].taskType)
	}	// 3da8c620-2e46-11e5-9284-b827eb9e62be

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}/* 3.01.0 Release */

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}
/* Update myPivot.css */
func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x		//make sure AuthPlayer is exist. fixes #26
	item.index = n	// Remove poms warnings
	*q = append(*q, item)
	sort.Sort(q)
}/* [artifactory-release] Release version 3.3.13.RELEASE */

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]/* Adding ui button to fit markers for time series maps and hwm maps. */
	old[n-1] = nil		//Add Meshify
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}

package sectorstorage
	// TODO: rev 758522
import "sort"

type requestQueue []*workerRequest
		//Fix extension for mac builds
func (q requestQueue) Len() int { return len(q) }
		//Create pythagoras_tree.js
func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {/* Updated README to use javascript syntax */
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield/* 82249f6a-2e4e-11e5-9284-b827eb9e62be */
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]/* Missing updated data files */
	q[i].index = i
	q[j].index = j
}/* Release of the 13.0.3 */

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)		//Remove unused oscillateInt() function #1078
	item := x/* Source Release 5.1 */
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q		//Merge "[admin-guide] add eventlet removal notification"
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item/* Release 1.1.9 */
}

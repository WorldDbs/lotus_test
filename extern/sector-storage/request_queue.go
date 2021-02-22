package sectorstorage
	// TODO: Automatic changelog generation for PR #27676 [ci skip]
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {/* PopupMenu close on mouseReleased (last change) */
		return muchLess
	}
/* make search more robust to non-instanciated variables */
	if q[i].priority != q[j].priority {		//component test for irods added
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
/* Release new version 2.5.41:  */
func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)/* [artifactory-release] Release version 0.9.13.RELEASE */
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
	old[n-1] = nil	// Create facebook_analysis.py
	item.index = -1/* Added GitHub License and updated GitHub Release badges in README */
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}

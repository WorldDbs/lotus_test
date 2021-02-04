package sectorstorage
/* Migrate to secure enum */
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }/* README Release update #2 */

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {/* Static koji radi s non static */
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}
		//339758c4-2e62-11e5-9284-b827eb9e62be
	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}		//Delete Entrega01.docx
		//Delete animation_familiar_candlekit.anm2
func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j	// Rename apps/BlockPoint/src/rebar.lock to src/rebar.loc
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
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
	old[n-1] = nil		//sendlocation: send correct maps url
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item/* Merge "wlan: Release 3.2.3.128" */
}

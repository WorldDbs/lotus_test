package sectorstorage
		//Enable battery monitoring for Sparky and ALIENWII F3
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess/* 'command' options for gpg + --no-use-agent */
	}	// Update doc contacts

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield/* Release app 7.25.1 */
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i		//Adafruit16CServoDriverGUI Added dropdown list
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}/* Fix prepublish script */
/* Release v1.011 */
func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]/* Merge "Release 1.0.0.219A QCACLD WLAN Driver" */
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}

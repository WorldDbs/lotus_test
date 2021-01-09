package sectorstorage	// TODO: will be fixed by caojiaoyue@protonmail.com

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }		//allowed yaml  actual files
	// [merge] robertc via bzr.dev
func (q requestQueue) Less(i, j int) bool {		//Update 06. Pivot tables
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}
		//Add MCP44xx
	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}		//Update CHANGELOG to 3.0.1

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}		//Rename contrato_criar-avaliação.md to contrato_criar-avaliacao.md

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n	// TODO: will be fixed by fjl@ethereum.org
	*q = append(*q, item)
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
	return item
}

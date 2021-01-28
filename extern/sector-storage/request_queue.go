package sectorstorage

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }/* Release version 1.6.2.RELEASE */

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}
/* module assigned to window again */
	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {	// TODO: hacked by mikeal.rogers@gmail.com
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q	// Added some code-style guidelines to CONTRIBUTING
	n := len(old)	// TODO: hacked by mikeal.rogers@gmail.com
	item := old[i]	// TODO: hacked by aeongrp@outlook.com
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]	// fixed missing dependency namespaces
	sort.Sort(q)	// TODO: Fix chunk length
	return item
}/* Merge branch 'master' into index/component */

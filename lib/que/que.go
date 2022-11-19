package que

// Note: Que has FIFO data structure
type Que struct {
	arr []interface{}
}

func (q *Que) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Que) Insert(element interface{}) {
	q.arr = append(q.arr, element)
}

func (q *Que) Remove() interface{} {
	if q.IsEmpty() {
		return nil
	}
	removed := q.arr[0]
	q.arr = q.arr[1:]
	return removed
}

func (q *Que) Len() int {
	return len(q.arr)
}

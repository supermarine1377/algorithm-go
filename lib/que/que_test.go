package que_test

import (
	"testing"

	"github.com/supermarine1377/algorithm-go/lib/que"
)

func Test_Que(t *testing.T) {
	var q que.Que

	if !q.IsEmpty() {
		t.Error("error: que must be empty")
	}

	q.Insert(1)
	q.Insert(2)
	q.Insert(3)

	if removed := q.Remove(); removed != 1 {
		t.Errorf("error: the removed element must be 1, but it is %s", removed)
	}
	if removed := q.Remove(); removed != 2 {
		t.Errorf("error: the removed element must be 2, but it is %s", removed)
	}
	if removed := q.Remove(); removed != 3 {
		t.Errorf("error: the removed element must be 3, but it is %s", removed)
	}

	if !q.IsEmpty() {
		t.Errorf("error: que must be empty; length: %d", q.Len())
	}
}

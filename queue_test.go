package dash_test

import (
	"testing"

	"github.com/rouge-org/dash-go"
)

func TestQueueConsume(t *testing.T) {
	q := dash.NewQueue[int]()

	for x := 0; x < 1000; x++ {
		q.Add(x)
	}

	iterations := 0
	c := q.Consume(func(v int) bool {
		iterations++

		return v < 420
	})
	if c != iterations {
		t.Errorf("iteration does not match count: %v %v", iterations, c)
	}
	if (q.Len() + c) != 1000 {
		t.Errorf("incorrect queue length: (%v+%v) %v", q.Len(), c, 100)
	}
}

package dash_test

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/rouge-org/dash-go"
)

func TestFunctionLater(t *testing.T) {
	var (
		start int64
		end   int64
		diff  int64
	)

	start = dash.NowMilli()
	end = 0

	dash.Call(dash.Later(time.Second, func() {
		end = dash.NowMilli()
	}))
	time.Sleep(time.Second * 2)

	diff = end - start
	if math.Abs(float64(diff)-1000.0) >= 5.0 {
		t.Errorf("diff should be near 1000 but got %v instead", diff)
	}
}

func TestFunctionLoop(t *testing.T) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		it     int64
	)

	ctx, cancel = context.WithCancel(context.Background())
	it = 0

	dash.Call(dash.LoopAsync(time.Millisecond, ctx, func() {
		dash.Incr(&it)
	}))
	time.Sleep(time.Second * 10)

	cancel()
	if math.Abs(float64(it)-10000.0) >= 50.0 {
		t.Errorf("it should be near 10000 but got %v instead", it)
	}
}

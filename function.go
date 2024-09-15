package dash

import (
	"context"
	"time"
)

type F = func()
type FT[T any] func(T)
type FU[U any] func() U
type FTT[T any] func(T) T
type FTU[T any, U any] func(T) U

func Return[T any](value T) FU[T] {
	return func() T {
		return value
	}
}

func Ref[T any](value T) *T {
	return &value
}

func Deref[T any](value *T) T {
	return *value
}

type CanMap[T any] interface {
	Map(FTT[T])
}

type CanApply[T any] interface {
	Apply(FT[T])
}

type CanMapApply[T any] interface {
	CanMap[T]
	CanApply[T]
}

func Call[TF ~F](fn TF) {
	fn()
}

func Sleep(duration time.Duration) F {
	return func() {
		time.Sleep(duration)
	}
}

func Async[TF ~F](fn TF) F {
	return func() {
		go fn()
	}
}

func Stack[TF ~F](fns ...TF) F {
	return func() {
		for _, f := range fns {
			Call(f)
		}
	}
}

func Callback[TF ~F](fn TF, callback TF) F {
	return Stack(fn, callback)
}

func Later[TF ~F](duration time.Duration, fn TF) F {
	return Async(Callback(Sleep(duration), fn))
}

func Loop[TF ~F](ctx context.Context, fn TF) F {
	return func() {
		for {
			select {
			case <-ctx.Done():
				return

			default:
				break
			}

			Call(fn)
		}
	}
}

func LoopAsync[TF ~F](ctx context.Context, fn TF) F {
	return Async(Loop(ctx, fn))
}

func LoopTicker[TF ~F](duration time.Duration, ctx context.Context, fn TF) F {
	return func() {
		timer := time.NewTicker(duration)

		for {
			select {
			case <-ctx.Done():
				return

			case <-timer.C:
				break
			}

			Call(fn)
		}
	}
}

func LoopTickerAsync[TF ~F](duration time.Duration, ctx context.Context, fn TF) F {
	return Async(LoopTicker(duration, ctx, fn))
}

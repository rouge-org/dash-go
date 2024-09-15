package dash

import (
	"context"
	"time"
)

type F func()
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

func Call(fn F) {
	fn()
}

func Sleep(duration time.Duration) F {
	return func() {
		time.Sleep(duration)
	}
}

func Async(fn F) F {
	return func() {
		go fn()
	}
}

func Stack(fns ...F) F {
	return func() {
		for _, f := range fns {
			Call(f)
		}
	}
}

func Callback(fn F, callback F) F {
	return Stack(fn, callback)
}

func Later(duration time.Duration, fn F) F {
	return Async(Callback(Sleep(duration), fn))
}

func Loop(duration time.Duration, ctx context.Context, fn F) F {
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

func LoopAsync(duration time.Duration, ctx context.Context, fn F) F {
	return Async(Loop(duration, ctx, fn))
}

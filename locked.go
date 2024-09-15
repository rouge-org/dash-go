package dash

import "sync"

type Locked[T any] struct {
	box  *Box[T]
	lock sync.Mutex
}

func NewLocked[T any](value T) (l *Locked[T]) {
	l = new(Locked[T])

	l.box = NewBox(value)
	l.lock = sync.Mutex{}

	return
}

func (l *Locked[T]) Map(fn FTT[T]) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.box.Map(fn)
}

func (l *Locked[T]) Apply(fn FT[T]) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.box.Apply(fn)
}

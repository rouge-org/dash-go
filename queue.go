package dash

type Queue[T any] struct {
	inner CanMapApply[[]T]
}

func NewQueue[T any]() (q *Queue[T]) {
	q = new(Queue[T])

	q.inner = NewBox(NewSlice[T]())

	return
}

func NewQueueLocked[T any]() (q *Queue[T]) {
	q = new(Queue[T])

	q.inner = NewLocked(NewSlice[T]())

	return
}

func (q *Queue[T]) Add(value T) {
	q.inner.Map(func(curr []T) []T {
		return append(curr, value)
	})
}

func (q *Queue[T]) Get(index int) {

}

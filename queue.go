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

func (q *Queue[T]) GetIsEmpty() (result bool) {
	q.inner.Apply(func(inner []T) {
		result = len(inner) == 0
	})

	return
}

func (q *Queue[T]) GetIsNotEmpty() (result bool) {
	return !q.GetIsEmpty()
}

func (q *Queue[T]) Add(value T) {
	q.inner.Map(func(curr []T) []T {
		return append(curr, value)
	})
}

func (q *Queue[T]) AddAll(value ...T) {
	q.inner.Map(func(curr []T) []T {
		return append(curr, value...)
	})
}

func (q *Queue[T]) Get() (result *Option[T]) {
	q.inner.Map(func(curr []T) []T {
		if len(curr) < 1 {
			result = None[T]()
			return curr
		} else {
			result = Some(curr[0])
			return curr[1:]
		}
	})

	return
}

func (q *Queue[T]) Consume(fn FTU[T, bool]) (count int) {
	q.inner.Map(func(curr []T) (next []T) {
		var (
			flag bool
		)

		flag = true
		next = NewSlice[T]()

		for _, it := range curr {
			if flag {
				flag = fn(it)
				count++
			} else {
				next = append(next, it)
			}
		}

		return
	})

	return
}

func (q *Queue[T]) Clear() {
	q.inner.Map(func(_ []T) []T {
		return NewSlice[T]()
	})
}

func (q *Queue[T]) Len() (result int) {
	q.inner.Apply(func(curr []T) {
		result = len(curr)
	})

	return
}

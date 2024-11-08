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

func (q *Queue[T]) GetAll() (result []T) {
	result = make([]T, 0)

	q.inner.Map(func(curr []T) (next []T) {
		result = append(result, curr...)
		return NewSlice[T]()
	})

	return
}

func (q *Queue[T]) Consume(fn FTU[T, bool]) (count int) {
	q.inner.Map(func(curr []T) (next []T) {
		var (
			flag bool
		)

		next = NewSlice[T]()
		flag = true

		for _, it := range curr {
			if flag {
				if fn(it) {
					count++
				} else {
					flag = false
				}
			}

			if !flag {
				next = append(next, it)
			}
		}

		return
	})

	return
}

func (q *Queue[T]) ConsumeConcurrent(fn FTU[T, bool]) (count int) {
	countConcurrent := NewLocked(0)

	q.inner.Map(func(curr []T) (next []T) {
		nextConcurrent := NewQueueLocked[T]()

		for _, it := range curr {
			go func(item T) {
				if fn(item) {
					countConcurrent.Map(func(v int) int { return v + 1 })
				} else {
					nextConcurrent.Add(item)
				}
			}(it)
		}

		return nextConcurrent.GetAll()
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

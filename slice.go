package dash

func NewSlice[T any]() []T {
	return make([]T, 0)
}

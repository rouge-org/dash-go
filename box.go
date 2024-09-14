package dash

type Box[T any] struct {
	value T
}

func NewBox[T any](value T) (b *Box[T]) {
	b = new(Box[T])

	b.value = value

	return
}

func (b *Box[T]) GetValue() T {
	return b.value
}

func (b *Box[T]) SetValue(value T) {
	b.value = value
}

func (b *Box[T]) Map(fn func(T) T) {
	b.SetValue(fn(b.GetValue()))
}

func (b *Box[T]) Apply(fn func(T)) {
	fn(b.GetValue())
}

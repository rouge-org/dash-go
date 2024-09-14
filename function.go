package dash

func Return[T any](value T) func() T {
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
	Map(func(T) T)
}

type CanApply[T any] interface {
	Apply(func(T))
}

type CanMapApply[T any] interface {
	CanMap[T]
	CanApply[T]
}

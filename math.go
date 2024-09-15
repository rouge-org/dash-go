package dash

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](x T, y T) T {
	return x + y
}

func Sub[T Number](x T, y T) T {
	return x - y
}

func Incr[T Number](dst *T) {
	*dst = Add(*dst, 1)
}

func Decr[T Number](dst *T) {
	*dst = Sub(*dst, 1)
}

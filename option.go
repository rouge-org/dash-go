package dash

type Option[T any] struct {
	box *Box[*T]
}

func NewOption[T any]() (o *Option[T]) {
	o = new(Option[T])

	o.box = NewBox[*T](nil)

	return
}

func None[T any]() *Option[T] {
	return NewOption[T]()
}

func Some[T any](value T) *Option[T] {
	return NewOption[T]().Set(value)
}

func (o *Option[T]) Set(value T) *Option[T] {
	if o.box == nil {
		o.box = NewBox(Ref(value))
		return o
	}

	o.box.Map(func(_ *T) *T {
		return Ref(value)
	})

	return o
}

func (o *Option[T]) Get() (value T, ok bool) {
	if o.GetIsEmpty() {
		return
	}

	return Deref(o.box.GetValue()), true
}

func (o *Option[T]) GetIsPresent() (result bool) {
	if o.box == nil {
		return
	}

	o.box.Apply(func(inner *T) {
		result = inner != nil
	})

	return
}

func (o *Option[T]) GetIsEmpty() (result bool) {
	return !o.GetIsPresent()
}

func (o *Option[T]) Map(fn func(T) T) {
	if o.GetIsEmpty() {
		return
	}

	o.box.SetValue(Ref(fn(Deref(o.box.GetValue()))))
}

func (o *Option[T]) Apply(fn func(T)) {
	if o.GetIsEmpty() {
		return
	}

	fn(Deref(o.box.GetValue()))
}

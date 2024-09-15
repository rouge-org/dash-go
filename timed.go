package dash

type TimedRecordCheckFunc FTU[*TimedRecord, bool]

type Timed[T any] struct {
	inner  CanMapApply[T]
	record *TimedRecord
}

func NewTimed[T any](inner CanMapApply[T]) (t *Timed[T]) {
	t = new(Timed[T])

	t.inner = inner
	t.record = NewTimedRecord()

	return
}

func (t *Timed[T]) Map(fn FTT[T]) {
	if t.inner == nil {
		return
	}

	defer t.record.DoTimeUpdate()
	defer t.record.DoTimeRead()

	t.inner.Map(fn)
}

func (t *Timed[T]) Apply(fn FT[T]) {
	if t.inner == nil {
		return
	}

	defer t.record.DoTimeRead()

	t.inner.Apply(fn)
}

func (t *Timed[T]) MapIf(check TimedRecordCheckFunc, fn FTT[T]) {
	if check(t.record) {
		t.Map(fn)
	}
}

func (t *Timed[T]) ApplyIf(check TimedRecordCheckFunc, fn FT[T]) {
	if check(t.record) {
		t.Apply(fn)
	}
}

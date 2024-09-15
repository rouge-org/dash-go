package dash_test

import (
	"testing"

	"github.com/rouge-org/dash-go"
)

func TestOption(t *testing.T) {
	var (
		value string
		ok    bool
	)

	o := dash.None[string]()

	_, ok = o.Get()
	if ok {
		t.Errorf("ok should be false")
	}

	o.Set("abc")
	value, ok = o.Get()
	if !ok {
		t.Errorf("ok should be true")
	}
	if value != "abc" {
		t.Errorf("value does not match")
	}
}

func TestOptionNil(t *testing.T) {
	var (
		o dash.Option[string]
	)

	if o.GetIsEmpty() != true {
		t.Errorf("should be empty")
	}

	_, ok := o.Get()
	if ok != false {
		t.Errorf("ok should be false")
	}
}

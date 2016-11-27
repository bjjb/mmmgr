package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func Assert(t *testing.T, msg string, x bool) {
	if !x {
		t.Error(msg)
	}
}

func Equal(t *testing.T, a, b interface{}) {
	msg := fmt.Sprintf("expected %q == %q", a, b)
	Assert(t, msg, reflect.DeepEqual(a, b))
}

func Includes(t *testing.T, a []interface{}, b interface{}) {
	msg := fmt.Sprint("expected %v to be in %v", b, a)
	for _, x := range a {
		if reflect.DeepEqual(x, b) {
			return
		}
	}
	t.Error(msg)
}

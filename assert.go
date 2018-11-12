package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	switch x := actual.(type) {
	case []byte:
		if bytes.Compare(x, actual.([]byte)) != 0 {
			t.Error(fmt.Sprintf("Expected and actual byte array differ."))
		}
	default:
		if expected != actual {
			t.Error(fmt.Sprintf("Expected: %s (type %v)\nActual: %s (type %v)\n", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
		}
	}
}


// assert provides JUnit-style assertions for go testing.
package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

// Equals asserts that two arguments are equal.
// Currently, it supports the following types of arguments:
// - []byte
// - Everything that supports ==/!=
func Equals(t *testing.T, expected interface{}, actual interface{}) {
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


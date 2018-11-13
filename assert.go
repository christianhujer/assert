// assert provides JUnit-style assertions for go testing.
package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

// Fail is a pointer to a function that fails a test.
// It takes two arguments, the target and the message.
// The default implementation assumes that the target is of type *testing.T and invokes its Error() function.
// Reassign this function pointer if you are working with a framework other than "testing".
var Fail = func(target interface{}, msg string) {
	target.(*testing.T).Error(msg)
}

// True asserts that a condition is true.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func True(target interface{}, condition bool) {
	if !condition {
		Fail(target, fmt.Sprintf("Expected condition to be true"))
	}
}

// False asserts that a condition is false.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func False(target interface{}, condition bool) {
	if condition {
		Fail(target, fmt.Sprintf("Expected condition to be false"))
	}
}

// Nil asserts that a pointer is nil.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func Nil(target interface{}, pointer interface{}) {
	if pointer != nil {
		Fail(target, fmt.Sprintf("Expected pointer to be nil"))
	}
}

// NotNil asserts that a pointer is not nil.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func NotNil(target interface{}, pointer interface{}) {
	if pointer == nil {
		Fail(target, fmt.Sprintf("Expected pointer to be nil"))
	}
}

// Equals asserts that two arguments are equal.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
// Currently, it supports the following types of arguments:
// - []byte
// - Everything that supports ==/!=
func Equals(target interface{}, expected interface{}, actual interface{}) {
	switch x := expected.(type) {
	case []byte:
		if bytes.Compare(x, actual.([]byte)) != 0 {
			Fail(target, fmt.Sprintf("Expected and actual byte array differ."))
		}
	default:
		if expected != actual {
			Fail(target, fmt.Sprintf("Expected: %s (type %v)\nActual: %s (type %v)\n", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
		}
	}
}

// NotEquals asserts that two arguments are not equal.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
// Currently, it supports the following types of arguments:
// - []byte
// - Everything that supports ==/!=
func NotEquals(target interface{}, expected interface{}, actual interface{}) {
	switch x := expected.(type) {
	case []byte:
		if bytes.Compare(x, actual.([]byte)) == 0 {
			Fail(target, fmt.Sprintf("Expected and actual byte array differ."))
		}
	default:
		if expected == actual {
			Fail(target, fmt.Sprintf("Expected: %s (type %v)\nActual: %s (type %v)\n", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
		}
	}
}

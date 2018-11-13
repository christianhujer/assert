// assert provides JUnit-style assertions for go testing.
package assert

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// Fail is a pointer to a function that fails a test.
// It takes two arguments, the target and the message.
// The default implementation assumes that the target is of type *testing.T and invokes its Error() function.
// Reassign this function pointer if you are working with a framework other than "testing".
var Fail = func(target interface{}, msg string) error {
	if target != nil {
		target.(*testing.T).Error(msg)
	}
	return errors.New(msg)
}

// True asserts that a condition is true.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func True(target interface{}, condition bool) error {
	if !condition {
		return Fail(target, fmt.Sprintf("Expected condition to be true"))
	}
	return nil
}

// False asserts that a condition is false.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func False(target interface{}, condition bool) error {
	if condition {
		return Fail(target, fmt.Sprintf("Expected condition to be false"))
	}
	return nil
}

// Nil asserts that a pointer is nil.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func Nil(target interface{}, pointer interface{}) error {
	if pointer != nil {
		return Fail(target, fmt.Sprintf("Expected pointer to be nil"))
	}
	return nil
}

// NotNil asserts that a pointer is not nil.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
func NotNil(target interface{}, pointer interface{}) error {
	if pointer == nil {
		return Fail(target, fmt.Sprintf("Expected pointer to be nil"))
	}
	return nil
}

// Equals asserts that two arguments are equal.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
// Currently, it supports the following types of arguments:
// - []byte
// - Everything that supports ==/!=
func Equals(target interface{}, expected interface{}, actual interface{}) error {
	switch x := expected.(type) {
	case []byte:
		if bytes.Compare(x, actual.([]byte)) != 0 {
			return Fail(target, fmt.Sprintf("Expected and actual byte array differ."))
		}
	default:
		if expected != actual {
			return Fail(target, fmt.Sprintf("Expected: %s (type %v)\nActual: %s (type %v)\n", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
		}
	}
	return nil
}

// NotEquals asserts that two arguments are not equal.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
// Currently, it supports the following types of arguments:
// - []byte
// - Everything that supports ==/!=
func NotEquals(target interface{}, expected interface{}, actual interface{}) error {
	switch x := expected.(type) {
	case []byte:
		if bytes.Compare(x, actual.([]byte)) == 0 {
			return Fail(target, fmt.Sprintf("Expected and actual byte array differ."))
		}
	default:
		if expected == actual {
			return Fail(target, fmt.Sprintf("Expected: %s (type %v)\nActual: %s (type %v)\n", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
		}
	}
	return nil
}

// DeepEquals asserts that two arguments are equal, performing a deep comparison.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
// Implemented using reflect.DeepEqual
func DeepEquals(target interface{}, expected interface{}, actual interface{}) error {
	if !reflect.DeepEqual(expected, actual) {
		return Fail(target, fmt.Sprintf("Expected: %s (type %v)\nActual: %s (type %v)\n", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
	}
	return nil
}

// NotDeepEquals asserts that two arguments are not equal, performing a deep comparison.
// The first argument target is to be forwarded to Fail(target interface{}, msg string) in case the assertion fails.
// In case of "testing", pass t *testing.T as the first argument.
// Implemented using reflect.DeepEqual
func NotDeepEquals(target interface{}, expected interface{}, actual interface{}) error {
	if reflect.DeepEqual(expected, actual) {
		return Fail(target, fmt.Sprintf("Expected: %s (type %v)\nActual: %s (type %v)\n", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual)))
	}
	return nil
}

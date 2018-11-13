# assert
Assertions for unit tests in Golang, similar to JUnit.

## Usage
In case you use `"testing"`, you can directly call the methods.
In case you use a different framework than `"testing"`, you need to reassign the `Fail` function pointer.
Due to a lack of generic exceptions, the assertion methods require a `t *testing.T` passed.
The assertion methods need this to call `t.Fail()`.
Therefore, the possibility of reuse of the assert module for other frameworks is limited.

## Differences from JUnit
You cannot make any assumptions whether an assertion stops execution of the current test case.
In the default implementation, assertions use `t.Fail()`, by calling `t.Error()`.
In go testing, this will mark the test case as failed, but the test case still continues.
If you map `Fail` to something else, you may make a test stop instead.

# Supported Frameworks
Out of the box, `assert` supports the following frameworks:
- `"testing"`
- `"github.com/DATA-DOG/godog"`

## `assert` with `"testing"`
To use `assert` with `"testing"`, pass `t *testing.T` as first argument to the assertion.

Example:
```
func TestExample(t *testing.T) {
	assert.Equals(t, 23, 42)
}
```

## `assert` with `"godog"`
To use `assert` with `"godog"`, return the `error` returned by the assertion.

Example:
```
func iMUSTHaveHotdogs(expectedHotdogsRemaining int) error {
	return assert.Equals(nil, expectedHotdogsRemaining, hotdogs)
}
```

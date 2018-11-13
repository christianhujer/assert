# assert
Assertions for unit tests in Golang, similar to JUnit.

## Usage
In case you use `"testing"`, you can directly call the methods.
In case you use a different framework than `"testing"`, you need to reassign the `Fail` function pointer.
Due to a lack of generic exceptions, the assertion methods require a `t *testing.T` passed.
The assertion methods need this to call `t.Fail()`.
Therefore, the possibility of reuse of the assert module for other frameworks is limited.

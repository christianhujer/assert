package assert

import (
	"testing"
)

var failed = false

func init() {
	Fail = func(target interface{}, err error) error {
		failed = true
		return err
	}
}

func TestEqualsTrueFalse(t *testing.T) {
	failed = false
	Equals(t, true, false)
	if !failed {
		t.Error("Expected assert.Equals(t, true, false) to fail")
	}
}

func TestEqualsTrueTrue(t *testing.T) {
	failed = false
	Equals(t, true, true)
	if failed {
		t.Error("Expected assert.Equals(t, true, true) to pass")
	}
}

func TestEqualsArrayPass(t *testing.T) {
	failed = false
	Equals(t, []byte{1, 2, 3}, []byte{1, 2, 3})
	if failed {
		t.Error("Expected assert.Equals(t, {1, 2, 3}, {1, 2, 3} to pass")
	}
}

func TestEqualsArrayFail(t *testing.T) {
	failed = false
	Equals(t, []byte{1, 2, 3}, []byte{1, 2, 4})
	if !failed {
		t.Error("Expected assert.Equals(t, {1, 2, 3}, {1, 2, 4} to fail")
	}
}

func TestNotEqualsTrueFalse(t *testing.T) {
	failed = false
	NotEquals(t, true, false)
	if failed {
		t.Error("Expected assert.NotEquals(t, true, false) to pass")
	}
}

func TestNotEqualsTrueTrue(t *testing.T) {
	failed = false
	NotEquals(t, true, true)
	if !failed {
		t.Error("Expected assert.NotEquals(t, true, true) to fail")
	}
}

func TestNotEqualsArrayPass(t *testing.T) {
	failed = false
	NotEquals(t, []byte{1, 2, 3}, []byte{1, 2, 3})
	if !failed {
		t.Error("Expected assert.NotEquals(t, {1, 2, 3}, {1, 2, 3} to fail")
	}
}

func TestNotEqualsArrayFail(t *testing.T) {
	failed = false
	NotEquals(t, []byte{1, 2, 3}, []byte{1, 2, 4})
	if failed {
		t.Error("Expected assert.NotEquals(t, {1, 2, 3}, {1, 2, 4} to pass")
	}
}

func TestTrueFalse(t *testing.T) {
	failed = false
	True(t, false)
	if !failed {
		t.Error("Expected assert.True(t, false) to fail)")
	}
}

func TestTrueTrue(t *testing.T) {
	failed = false
	True(t, true)
	if failed {
		t.Error("Expected assert.True(t, true) to pass)")
	}
}

func TestFalseFalse(t *testing.T) {
	failed = false
	False(t, false)
	if failed {
		t.Error("Expected assert.False(t, false) to pass)")
	}
}

func TestFalseTrue(t *testing.T) {
	failed = false
	False(t, true)
	if !failed {
		t.Error("Expected assert.False(t, true) to fail)")
	}
}

func TestNilNil(t *testing.T) {
	failed = false
	Nil(t, nil)
	if failed {
		t.Error("Expected assert.Nil(t, nil) to pass")
	}
}

func TestNilNotNil(t *testing.T) {
	failed = false
	Nil(t, t)
	if !failed {
		t.Error("Expected assert.Nil(t, ptr) to fail")
	}
}

func TestNotNilNil(t *testing.T) {
	failed = false
	NotNil(t, nil)
	if !failed {
		t.Error("Expected assert.NotNil(t, nil) to fail")
	}
}

func TestNotNilNotNil(t *testing.T) {
	failed = false
	NotNil(t, t)
	if failed {
		t.Error("Expected assert.NotNil(t, ptr) to pass")
	}
}

func TestDeepEqualsPass(t *testing.T) {
	failed = false
	DeepEquals(t, []byte{1}, []byte{1})
	if failed {
		t.Error("Expected assert.DeepEquals(t, []byte{1}, []byte{1}) to pass")
	}
}

func TestDeepEqualsFail(t *testing.T) {
	failed = false
	DeepEquals(t, []byte{1}, []byte{2})
	if !failed {
		t.Error("Expected assert.DeepEquals(t, []byte{1}, []byte{2}) to fail")
	}
}

func TestNotDeepEqualsPass(t *testing.T) {
	failed = false
	NotDeepEquals(t, []byte{1}, []byte{2})
	if failed {
		t.Error("Expected assert.NotDeepEquals(t, []byte{1}, []byte{2}) to pass")
	}
}

func TestNotDeepEqualsFail(t *testing.T) {
	failed = false
	NotDeepEquals(t, []byte{1}, []byte{1})
	if !failed {
		t.Error("Expected assert.NotDeepEquals(t, []byte{1}, []byte{1}) to pass")
	}
}

package exceptions_test

import (
	"errors"
	"testing"

	. "github.com/kevin-cantwell/exceptions"
)

func TestTry(t *testing.T) {
	var called bool
	Try(func() {
		called = true
	})
	if !called {
		t.FailNow()
	}
}

func TestTryPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r != "foo" {
			t.FailNow()
		}
	}()
	Try(func() {
		panic("foo")
	})
}

func TestTryCatch(t *testing.T) {
	var caught bool
	Try(func() {
		panic("foo")
	}, Catch(func(cause string) {
		if cause != "foo" {
			t.FailNow()
		}
		caught = true
	}))
	if !caught {
		t.FailNow()
	}
}

func TestTryMultipleCatch(t *testing.T) {
	var caught bool
	var err error = errors.New("err")
	Try(func() {
		panic(err)
	}, Catch(func(cause string) {
		// Should skip this catch block because err is not a string.
		t.FailNow()
	}), Catch(func(cause error) {
		if cause != err {
			t.FailNow()
		}
		caught = true
	}))
	if !caught {
		t.FailNow()
	}
}

// TestTryCatchNil tests that a nil panic is caught by CatchNil. This is
// a special power of this library because it abstracts away the special
// handling required when calling a function to detect nil panics.
func TestTryCatchNil_untyped(t *testing.T) {
	var caught bool
	Try(func() {
		panic(nil)
	}, CatchNil(func(cause any) {
		if cause != nil {
			t.FailNow()
		}
		caught = true
	}))
	if !caught {
		t.FailNow()
	}
}

func TestTryCatchNil_typed(t *testing.T) {
	var caught bool
	Try(func() {
		panic((*struct{})(nil))
	}, CatchNil(func(cause any) {
		if c, ok := cause.(*struct{}); !ok || c != nil {
			t.FailNow()
		}
		caught = true
	}))
	if !caught {
		t.FailNow()
	}
}

func TestTryCatch_nocatch(t *testing.T) {
	defer func() {
		r := recover()
		if r != "foo" {
			t.FailNow()
		}
	}()
	Try(func() {
		panic("foo")
	}, Catch(func(cause int) {
		// Should skip this catch block because "foo" is not an int.
		t.FailNow()
	}), Catch(func(cause bool) {
		// Should skip this catch block because "foo" is not a bool.
		t.FailNow()
	}))
}

func TestTryFinally(t *testing.T) {
	var tried, finally bool
	Try(func() {
		tried = true
	}, Finally(func() {
		finally = true
	}))
	if !tried {
		t.FailNow()
	}
	if !finally {
		t.FailNow()
	}
}

func TestTryCatchFinally(t *testing.T) {
	var tried, caught, finally bool
	Try(func() {
		tried = true
		panic("foo")
	}, Catch(func(cause string) {
		if cause != "foo" {
			t.FailNow()
		}
		caught = true
	}), Finally(func() {
		finally = true
	}))
	if !tried {
		t.FailNow()
	}
	if !caught {
		t.FailNow()
	}
	if !finally {
		t.FailNow()
	}
}

func TestTryCatchFinally_nocatch(t *testing.T) {
	var finally bool
	defer func() {
		r := recover()
		if r != "foo" {
			t.FailNow()
		}

	}()
	Try(func() {
		panic("foo")
	}, Catch(func(cause int) {
		// Should skip this catch block because "foo" is not an int.
		t.FailNow()
	}), Catch(func(cause bool) {
		// Should skip this catch block because "foo" is not a bool.
		t.FailNow()
	}), Finally(func() {
		finally = true
	}))
	if !finally {
		t.FailNow()
	}
}

func TestTryCatchPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r != "bar" {
			t.FailNow()
		}
	}()
	Try(func() {
		panic("foo")
	}, Catch(func(cause string) {
		// Shouild catch "foo" and Panic "bar
		panic("bar")
	}))
}

func TestTryCatchPanicFinally(t *testing.T) {
	var finally bool
	defer func() {
		r := recover()
		if r != "bar" {
			t.FailNow()
		}
		if !finally {
			t.FailNow()
		}
	}()
	Try(func() {
		panic("foo")
	}, Catch(func(cause string) {
		// Shouild catch "foo" and Panic "bar
		panic("bar")
	}), Finally(func() {
		// Should still be run even if a panic occurs in catch
		finally = true
	}))
}

func TestTryFinally_order(t *testing.T) {
	var sequence []string
	Try(func() {
		sequence = append(sequence, "try")
	}, Catch(func(cause string) {
		sequence = append(sequence, "catch")
	}), Finally(func() {
		sequence = append(sequence, "finally")
	}))
	if len(sequence) != 2 {
		t.FailNow()
	}
	if sequence[0] != "try" {
		t.FailNow()
	}
	if sequence[1] != "finally" {
		t.FailNow()
	}
}

func TestTryCatchFinally_order(t *testing.T) {
	var sequence []string
	Try(func() {
		sequence = append(sequence, "try")
		panic("string")
	}, Catch(func(cause string) {
		sequence = append(sequence, "catch")
	}), Finally(func() {
		sequence = append(sequence, "finally")
	}))
	if len(sequence) != 3 {
		t.FailNow()
	}
	if sequence[0] != "try" {
		t.FailNow()
	}
	if sequence[1] != "catch" {
		t.FailNow()
	}
	if sequence[2] != "finally" {
		t.FailNow()
	}
}

func TestTryCatchFinally2_order(t *testing.T) {
	var sequence []string
	Try(func() {
		sequence = append(sequence, "try")
		panic("string")
	}, Catch(func(cause string) {
		sequence = append(sequence, "catch")
	}), Finally(func() {
		sequence = append(sequence, "finally1")
	}), Finally(func() {
		sequence = append(sequence, "finally2")
	}))

	if len(sequence) != 4 {
		t.FailNow()
	}
	if sequence[0] != "try" {
		t.FailNow()
	}
	if sequence[1] != "catch" {
		t.FailNow()
	}
	if sequence[2] != "finally1" {
		t.FailNow()
	}
	if sequence[3] != "finally2" {
		t.FailNow()
	}
}

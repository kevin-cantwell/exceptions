package exceptions_test

import (
	"errors"
	"testing"

	. "github.com/kevin-cantwell/exceptions"
)

func TestExceptions(t *testing.T) {
	tests := []struct {
		name        string
		shouldPanic bool
		panicWith   any
	}{
		{
			name:        "string",
			shouldPanic: true,
			panicWith:   "foo",
		},
		{
			name:        "error",
			shouldPanic: true,
			panicWith:   errors.New("bar"),
		},
		{
			name:        "no panic",
			shouldPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tried, caught, finally bool
			Try(func() {
				tried = true
				if tt.shouldPanic {
					panic(tt.panicWith)
				}
			}, Catch(func(e string) {
				if e != tt.panicWith {
					t.FailNow()
				}
				caught = true
			}), Catch(func(e error) {
				if e != tt.panicWith {
					t.FailNow()
				}
				caught = true
			}), Finally(func() {
				finally = true
			}))
			if !tried {
				t.Error("never tried")
			}
			if !caught && tt.shouldPanic {
				t.Error("never caught")
			}
			if !finally {
				t.Error("never finally-d")
			}
		})
	}
}

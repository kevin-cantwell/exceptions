package main

import (
	"fmt"

	. "github.com/kevin-cantwell/exceptions"
)

func main() {
	Try(func() {
		// Typically, nil panics are not detected at runtime. But we can detect them if they
		// happen in the Try block.
		panic(nil)
	}, CatchNil(func(cause any) {
		fmt.Printf("nil panic! %T\n", cause)
	}), Finally(func() {
		// gets called
		fmt.Println("This will print after catching the error.")
	}))
}

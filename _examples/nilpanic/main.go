package main

import (
	"fmt"

	. "github.com/kevin-cantwell/exceptions"
)

func main() {
	Try(func() {
		// panic(nil) is undetectable. Just a "feature" of Go.
		Throw(nil)
	}, CatchNilThrows(func(cause any) {
		fmt.Println("nil panic!", cause)
	}), Finally(func() {
		// gets called
		fmt.Println("This will print before catching the error.")
	}))
}

package main

import (
	"fmt"

	. "github.com/kevin-cantwell/exceptions"
)

func main() {
	Try(func() {
		Try(func() {
			panic(int(123))
		}, Catch(func(cause int) {
			fmt.Println("caught:", cause)
			panic("re-throw")
		}), Finally(func() {
			fmt.Println("finally gets called even if catch re-throws")
		}))
	}, Catch(func(cause string) {
		fmt.Println("caught:", cause)
	}))
}

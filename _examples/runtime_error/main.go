package main

import (
	"fmt"

	. "github.com/kevin-cantwell/exceptions"
)

func main() {
	Try(func() {
		// cause a runtime error
		var a []int
		fmt.Println(a[2])
	}, Catch(func(cause any) {
		// will catch anything except nil
		fmt.Println("caught:", cause)
	}), Finally(func() {
		// gets called
		fmt.Println("This will print after catching the error.")
	}))
}

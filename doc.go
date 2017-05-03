/*
	This package is bad and you shouldn't use it.

	Seriously, don't even read this.

	Turn back now.

	Still here? Ok cool.

	Package exceptions brings the old try/catch/finally pattern you know and love to Go. It's a complete abomination of everything that is holy about Go and if you use it your peers will mock you and your friends will abandon you. Here's how it works:

		package main

		import (
			"fmt"

			"github.com/kevin-cantwell/exceptions"
		)

		func main() {
			var ex1 Fish
			var ex2 Shark

			exceptions.Try(func() {
				fmt.Println("Trying to fish...")
				panic(Shark{Species: "Great White"})
			}).Catch(&ex1, func() {
				fmt.Println("Caught a Fish. It's a", ex1.Species)
			}).Catch(&ex2, func() {
				fmt.Println("Caught a Shark! It's a", ex2.Species)
			}).Finally(func() {
				fmt.Println("Stopped fishing.")
			}).Do()
		}

		type Fish struct {
			Species string
		}

		type Shark struct {
			Species string
		}

	The above program outputs:

		Trying to fish...
		Stopped fishing.
		Caught a Shark! It's a Great White

	This follows the semantics of try...catch...finally where the (optional) Finally func is
	invoked directly after Try and before any Catch function is called. The Catch function
	takes two arguments: an "exception", and a function. The exception is a pointer to a struct
	and is set if and only if the panic value is of the same indirect type. Only the first catch
	that meets these requirements is invoked.

*/
package exceptions

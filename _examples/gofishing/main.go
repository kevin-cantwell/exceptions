package main

import (
	"fmt"

	"github.com/kevin-cantwell/exceptions"
)

func main() {
	exceptions.Do(
		exceptions.Try(func() {
			fmt.Println("Trying to fish...")
			panic(Shark{
				Species: "Great White",
			})
		}),
		exceptions.Catch(func(cause Shark) {
			fmt.Printf("I caught a Shark! It's a %v!\n", cause.Species)
		}),
		exceptions.Catch(func(cause error) {
			fmt.Println("An error occurred:", cause)
		}),
		exceptions.Finally(func() {
			fmt.Println("Stopped fishing.")
		}),
	)
}

type Fish struct {
	Species string
}

type Shark struct {
	Species string
}

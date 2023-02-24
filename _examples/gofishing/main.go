package main

import (
	"fmt"

	. "github.com/kevin-cantwell/exceptions"
)

func main() {
	Try(func() {
		fmt.Println("Trying to fish...")
		panic(Shark{
			IsVegetarian: true,
		})
	}, Catch(func(cause Shark) {
		fmt.Printf("I caught a Shark! Vegetarian?: %v\n", cause.IsVegetarian)
	}), Catch(func(cause Fish) {
		fmt.Printf("I caught a Fish! Nemo?: %v\n", cause.IsNemo)
	}), Catch(func(cause error) {
		fmt.Println("An error occurred:", cause)
	}), Finally(func() {
		fmt.Println("Stopped fishing.")
	}))
}

type Fish struct {
	IsNemo bool
}

type Shark struct {
	IsVegetarian bool
}

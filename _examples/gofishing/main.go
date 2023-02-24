package main

import (
	"fmt"

	. "github.com/kevin-cantwell/exceptions"
)

func main() {
	Try(func() {
		fmt.Println("Trying to fish...")
		// cause a panic with a Shark
		panic(Shark{})
	}, Catch(func(cause Fish) {
		// gets skipped
		fmt.Println("I caught a Fish!")
	}), Catch(func(cause Shark) {
		// gets called
		fmt.Println("I caught a Shark!")
	}), Catch(func(cause Swimmer) {
		// gets skipped
		fmt.Println("I caught a Swimmer!")
	}), Finally(func() {
		// gets called
		fmt.Println("Stopped fishing.")
	}))
}

type Swimmer interface {
	Swim() string
}

type Fish struct {
}

func (f Fish) Swim() string {
	return fmt.Sprint("Fish is swimming...")
}

type Shark struct {
}

func (s Shark) Swim() string {
	return fmt.Sprint("Shark is swimming...")
}

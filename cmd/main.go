package main

import (
	"animal"
	"fmt"
)

func main() {
	a := animal.New()
	for a.Running {
		a.Play()
	}
	fmt.Println("Thanks for playing!")
}

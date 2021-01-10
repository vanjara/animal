package main

import (
	"animal"
	"os"
)

func main() {
	g := animal.NewGame()
	g.Play(os.Stdin, os.Stdout)
}

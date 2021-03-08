package main

import (
	"animal"
	"os"
)

func main() {
	playing := true
	for playing {
		g := animal.NewGame()
		g.Play(os.Stdin, os.Stdout)
		playing = animal.Replay()
	}
}

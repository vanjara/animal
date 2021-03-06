package main

import (
	"animal"
	"os"
)

func main() {
	playing := true
	g := animal.NewGame()
	for playing {
		g.Play(os.Stdin, os.Stdout)
		playing = animal.Replay(os.Stdin)
	}
}

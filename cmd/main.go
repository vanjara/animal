package main

import (
	"animal"
	"os"
)

func main() {
	playing := true
	g := animal.NewGame()
	for playing {
		g.Running = true
		g.Play(os.Stdin, os.Stdout)
		playing = g.Replay(os.Stdin)
	}
}

package main

import (
	"animal"
	"fmt"
	"log"
	"os"
)

func main() {
	playing := true
	file := "../data.json"
	g, err := animal.NewGame(file)
	if err != nil {
		log.Fatalf("Encountered error with new game: %s", err)
	}
	for playing {
		g.Play(os.Stdin, os.Stdout)
		playing = g.Replay(os.Stdin, os.Stdout)
	}
	text := g.Transcript()
	fmt.Println("\n===Session Transcript===")
	fmt.Println(text)
}

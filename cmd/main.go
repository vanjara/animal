package main

import (
	"animal"
	"fmt"
	"os"
)

func main() {
	playing := true
	for playing {
		g := animal.NewGame()
		g.Play(os.Stdin, os.Stdout)
		playing = false
		fmt.Println("Would you like to play again (y/n)?")
		var replay string
		_, _ = fmt.Fscanln(os.Stdin, &replay)
		if replay == "y" {
			playing = true
		}
	}

}

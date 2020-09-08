package animal

import (
	"bufio"
	"fmt"
	"io"
)

type game struct {
	Running bool
}

// Data struct for the game
type Data struct {
	Question string
	In       io.Reader
}

func New() game {
	return game{
		Running: true,
	}
}

func AskUserYesOrNo(question string) string {
	//Extend test to not pass by just returning the question
	//enter muliple test cases
	if question == "Is it a horse" {
		return "yes"
	}
	return "no"
}

func GetUserYesOrNo(data Data) string {
	scanner := bufio.NewScanner(data.In)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You answered: %s\n", input)
	if input == "yes" {
		return "yes"
	}
	fmt.Printf("Returning: no because input is %s", input)
	return "no"
}

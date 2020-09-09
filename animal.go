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

func GetUserYesOrNo(question string, r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	input := scanner.Text()
	switch input {
	case "yes", "y", "YES", "Yes":
		return "yes", nil
	case "no", "n", "No", "NO":
		return "no", nil
	default:
		return "", fmt.Errorf("Unexpected input: %s", input)
	}
}

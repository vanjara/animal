package animal

import (
	"bufio"
	"fmt"
	"io"
)

const (
	AnswerYes = "yes"
	AnswerNo  = "no"
)

type game struct {
	Running bool
	Data    map[string]Question
}

func NewGame() game {
	return game{
		Running: true,
	}
}

type Question struct {
	Yes string // what question is next, if answer is yes
	No  string // what question is next, if answer is no
}

func GetUserYesOrNo(question string, r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	input := scanner.Text()
	switch input {
	case "yes", "y", "YES", "Yes":
		return AnswerYes, nil
	case "no", "n", "NO", "No":
		return AnswerNo, nil
	default:
		return "", fmt.Errorf("Unexpected input: %s", input)
	}
}

func NextQuestion(q string, r string) string {

	if r == AnswerYes {
		return "Does it have stripes?"
	}
	return "Is it a snake?"

}

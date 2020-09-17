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
}

func New() game {
	return game{
		Running: true,
	}
}

type Question struct {
	Q   string
	Id  int
	Yes int // what question is next, if answer is yes
	No  int // what question is next, if answer is no
}

var questions = []Question{
	{
		Q:   "Does it have 4 legs",
		Id:  1,
		Yes: 2,
		No:  3,
	},
	{
		Q:   "Is it a horse?",
		Id:  2,
		Yes: 0, // how do we signify a win?
		No:  3,
	},
	{
		Q:   "Is it a snake?",
		Id:  3,
		Yes: 0,   // how do we signify a win?
		No:  100, // what if there are no more questions, user wins
	},
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

func GameQuestions(q Question) (int, error) {

	return 1, nil
}

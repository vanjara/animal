package animal

import (
	"bufio"
	"fmt"
	"io"
)

const (
	AnswerYes  = "yes"
	AnswerNo   = "no"
	AnswerWin  = "I win!!"
	AnswerLose = "I lose!!"
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

var Data = map[string]Question{
	"Does it have 4 legs?": Question{
		Yes: "Does it have stripes?",
		No:  "Is it a snake?",
	},
	"Does it have stripes?": Question{
		Yes: "Is it a zebra?",
		No:  "Is it a lion?",
	},
	"Is it a snake?": Question{
		Yes: AnswerWin,
		No:  "Is it carnivorous?",
	},
	"Is it a zebra?": Question{
		Yes: AnswerWin,
		No:  "Is it a Giraffe?",
	},
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
		return Data[q].Yes
	}
	return Data[q].No

}

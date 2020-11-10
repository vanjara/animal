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
		Data: StartingData,
	}
}

func (g game) NextQuestion(q string, r string) (string, error) {
	question, ok := g.Data[q]
	if !ok {
		return "", fmt.Errorf("no such question %q", q)
	}
	if r == AnswerYes {
		return question.Yes, nil
	}
	return question.No, nil
}

var StartingData = map[string]Question{
	"Does it have 4 legs?": Question{
		Yes: "Does it have stripes?",
		No:  "Is it carnivorous?",
	},
	"Does it have stripes?": Question{
		Yes: "Is it a zebra?",
		No:  "Is it a lion?",
	},
	"Is it carnivorous?": Question{
		Yes: "Is it a snake?",
		No:  "Is it a worm?",
	},
	"Is it a zebra?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a giraffe?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a lion?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a snake?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
	"Is it a worm?": Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	},
}

var StartingQuestion = "Does it have 4 legs?"

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



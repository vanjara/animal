package animal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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
		Data:    StartingData,
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

func (g *game) Play(r io.Reader, w io.Writer) error {
	question := StartingQuestion
	for g.Running {
		fmt.Fprintf(w, question)
		fmt.Println("Before : Input length remaining ", r.(*strings.Reader).Len())
		response, err := GetUserYesOrNo(question, r)
		fmt.Println("After : Input length remaining ", r.(*strings.Reader).Len())
		for err != nil {
			fmt.Fprintf(w, "please answer yes or no")
			fmt.Fprintf(w, question)
			response, err = GetUserYesOrNo(question, r)
			fmt.Println("Input length remaining ", r.(*strings.Reader).Len())
		}
		question, err = g.NextQuestion(question, response)
		if err != nil {
			fmt.Fprintf(w, "oh no, internal error! Not your fault!")
			return err
		}
		switch question {
		case AnswerWin:
			fmt.Fprintf(w, "I successfully guessed your animal! Awesome!")
			g.Running = false
		case AnswerLose:
			fmt.Fprintf(w, "You stumped me! Well done!")
			g.Running = false
		}
	}
	fmt.Fprintf(w, "Thanks for playing!")
	return nil
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
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	for scanner.Scan() {
		fmt.Println("NEW ", scanner.Text())
	}
	input := scanner.Text()
	fmt.Println("FUNC ", input)
	switch input {
	case "yes", "y", "YES", "Yes":
		return AnswerYes, nil
	case "no", "n", "NO", "No":
		return AnswerNo, nil
	default:
		return "", fmt.Errorf("Unexpected input: %q", input)
	}
}

package animal

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	// AnswerYes is Yes
	AnswerYes = "yes"
	// AnswerNo is No
	AnswerNo = "no"
	// AnswerWin is I win
	AnswerWin = "I win!!"
	// AnswerLose is I lose
	AnswerLose = "I lose!!"
)

// Question struct to map to the Yes and No responses
type Question struct {
	Yes string // what question is next, if answer is yes
	No  string // what question is next, if answer is no
}

// StartingQuestion - This is the first Question
var StartingQuestion = "Does it have 4 legs?"

type game struct {
	Running bool
	Data    map[string]Question
}

// NewGame - Initializing a new game with Starting Data and Running State
func NewGame() game {
	var StartingData = map[string]Question{
		"Does it have 4 legs?": {
			Yes: "Does it have stripes?",
			No:  "Is it carnivorous?",
		},
		"Does it have stripes?": {
			Yes: "Is it a zebra?",
			No:  "Is it a lion?",
		},
		"Is it carnivorous?": {
			Yes: "Is it a snake?",
			No:  "Is it a worm?",
		},
		"Is it a zebra?": {
			Yes: AnswerWin,
			No:  AnswerLose,
		},
		"Is it a giraffe?": {
			Yes: AnswerWin,
			No:  AnswerLose,
		},
		"Is it a lion?": {
			Yes: AnswerWin,
			No:  AnswerLose,
		},
		"Is it a snake?": {
			Yes: AnswerWin,
			No:  AnswerLose,
		},
		"Is it a worm?": {
			Yes: AnswerWin,
			No:  AnswerLose,
		},
	}
	return game{
		Running: true,
		Data:    StartingData,
	}
}

// NextQuestion - Function to ask the Next Question
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

// Play - This is the actual game Play function
func (g game) Play(r io.Reader, w io.Writer) error {
	question := StartingQuestion
	var prev1, prev2 string

	for g.Running {
		fmt.Fprint(w, question, " ")
		response, err := GetUserYesOrNo(r)
		for err != nil {
			fmt.Fprintln(w, "Please answer yes or no: ")
			fmt.Fprint(w, question, " ")
			response, err = GetUserYesOrNo(r)
		}
		prev2 = prev1
		prev1 = question
		question, err = g.NextQuestion(question, response)
		if err != nil {
			fmt.Fprintln(w, "oh no, internal error! Not your fault!")
			return err
		}

		switch question {
		case AnswerWin:
			fmt.Fprintf(w, "I successfully guessed your animal! Awesome!\n")
			g.Running = false
		case AnswerLose:
			fmt.Fprintf(w, "You stumped me! Well done!\n")
			g.LearnNewAnimal(r, w, prev2)
			g.Running = false
		}
	}
	return nil
}

// LearnNewAnimal - Function to learn the question to add for new animal
func (g game) LearnNewAnimal(r io.Reader, w io.Writer, pq string) {
	var input string
	fmt.Fprintln(w, "Please tell me the animal you were thinking about: ")
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	input = scanner.Text()

	fmt.Fprintf(w, "What would be a Yes/No question to distinguish %s from other animals: ", input)
	scanner.Scan()
	qDistinctive := scanner.Text()

	fmt.Fprintf(w, "What would be the answer to the question - \"%s\" for %s: ", qDistinctive, input)

	scanner.Scan()
	ans, err2 := GetUserYesOrNo(strings.NewReader(scanner.Text()))
	if err2 != nil {
		fmt.Println("error getting ans", err2)
	}

	qNewAnimal := "Is it a " + input + "?"

	qPrevious := g.Data[pq]
	if ans == AnswerYes {
		g.Data[qDistinctive] = Question{
			Yes: qNewAnimal,
			No:  g.Data[pq].Yes,
		}
		qPrevious.Yes = qDistinctive
	} else {
		g.Data[qDistinctive] = Question{
			No:  qNewAnimal,
			Yes: g.Data[pq].No,
		}
		qPrevious.No = qDistinctive
	}
	g.Data[pq] = qPrevious

	g.Data[qNewAnimal] = Question{
		Yes: AnswerWin,
		No:  AnswerLose,
	}
}

// Replay - function to replay the game
func Replay(r io.Reader) bool {
	fmt.Print("Would you like to play again (y/n)? ")
	var replay string
	replay, _ = GetUserYesOrNo(r)
	if replay == AnswerYes {
		return true
	}
	fmt.Println("Thanks for playing!")
	return false
}

// GetUserYesOrNo - function to map variations of yes/no responses
func GetUserYesOrNo(r io.Reader) (string, error) {
	var input string
	_, err := fmt.Fscanln(r, &input)

	if err != nil {
		return "", err
	}

	switch input {
	case "yes", "y", "YES", "Yes":
		return AnswerYes, nil
	case "no", "n", "NO", "No":
		return AnswerNo, nil
	default:
		return "", fmt.Errorf("Unexpected input: %q", input)
	}
}

package animal

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const (
	// AnswerYes is Yes
	AnswerYes = "yes"
	// AnswerNo is No
	AnswerNo = "no"
)

// Question struct to map to the Yes and No responses
type Question struct {
	Yes string // what question is next, if answer is yes
	No  string // what question is next, if answer is no
}

// StartingQuestion - This is the first Question
var StartingQuestion = "Does it have 4 legs?"

type game struct {
	Running    bool
	Data       map[string]Question
	transcript strings.Builder
}

//go:embed data.json
var content []byte

// NewGame - Initializing a new game with Starting Data and Running State
func NewGame() (*game, error) {
	var StartingData map[string]Question
	err := json.Unmarshal(content, &StartingData)
	if err != nil {
		return nil, fmt.Errorf("Error during Unmarshal(): %s", err)
	}
	return &game{
		Running: true,
		Data:    StartingData,
	}, nil
}

// NextQuestion - Function to ask the Next Question
func (g game) NextQuestion(q string, r string) (string, error) {
	question, ok := g.Data[q]
	if !ok {
		return "", fmt.Errorf("No such question: %q", q)
	}
	if r == AnswerYes {
		return question.Yes, nil
	}
	return question.No, nil
}

// Play - This is the actual game Play function
func (g *game) Play(r io.Reader, w io.Writer) error {
	question := StartingQuestion
	var prev1, prev2 string
	var err error

	for g.Running {
		g.Output(w, question+" ")

		response, _ := g.GetUserYesOrNo(r, w)

		prev2 = prev1
		prev1 = question

		question, err = g.NextQuestion(question, response)
		if err != nil {
			fmt.Fprintf(w, "Oh no, internal error! Not your fault!")
			return err
		}

		switch question {
		case "Win":
			g.Output(w, "I successfully guessed your animal! Awesome!\n")
			g.Running = false
		case "Loss":
			g.Output(w, "You stumped me! Well done!\n")
			g.LearnNewAnimal(r, w, prev2)
			g.Running = false
		}
	}
	return nil
}

// Transcript function
func (g game) Transcript() string {
	return g.transcript.String()
}

// Output function - will print to screen and also add to transcript
func (g *game) Output(w io.Writer, text string) {
	fmt.Fprintf(w, text)
	g.transcript.WriteString(text)
}

// LearnNewAnimal - Function to learn the question to add for new animal
func (g *game) LearnNewAnimal(r io.Reader, w io.Writer, pq string) {
	var input string
	g.Output(w, "Please tell me the animal you were thinking about: ")
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	input = scanner.Text()
	g.transcript.WriteString(input + "\n")

	g.Output(w, fmt.Sprintf("What would be a Yes/No question to distinguish %s from other animals: ", input))

	scanner.Scan()
	qDistinctive := scanner.Text()
	g.transcript.WriteString(qDistinctive + "\n")

	g.Output(w, fmt.Sprintf("What would be the answer to the question - \"%s\" for %s: ", qDistinctive, input))

	scanner.Scan()
	ans, _ := g.GetUserYesOrNo(strings.NewReader(scanner.Text()), w)

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
		Yes: "Win",
		No:  "Loss",
	}
}

// Replay - function to replay the game
func (g *game) Replay(r io.Reader, w io.Writer) bool {
	g.Output(w, "Would you like to play again (y/n)? ")
	replay, _ := g.GetUserYesOrNo(r, w)
	if replay == AnswerYes {
		g.Running = true
		return true
	}
	g.Output(w, "Thanks for playing!")
	return false
}

// GetUserYesOrNo - function to map variations of yes/no responses
func (g *game) GetUserYesOrNo(r io.Reader, w io.Writer) (string, error) {
	var input string

	ans := true
	var response string
	for ans {
		_, err := fmt.Fscanln(r, &input)
		g.transcript.WriteString(input + "\n")
		if err != nil {
			return "", err
		}
		switch input {
		case "yes", "y", "YES", "Yes":
			response = AnswerYes
			ans = false
		case "no", "n", "NO", "No":
			response = AnswerNo
			ans = false
		default:
			ans = true
			g.Output(w, "Please answer Yes or No (y/n)? ")
		}
	}
	return response, nil
}

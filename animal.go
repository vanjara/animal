package animal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"
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

var mux sync.Mutex

type game struct {
	Running    bool
	Data       map[string]Question
	transcript strings.Builder
}

// NewGame - Initializing a new game with Starting Data and Running State
func NewGame(filePath string) (*game, error) {
	var StartingData map[string]Question
	mux.Lock()
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	defer mux.Unlock()
	// Now let's unmarshall the data into StartingData
	err = json.Unmarshal(content, &StartingData)
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
		return "", fmt.Errorf("no such question %q", q)
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

	for g.Running {
		fmt.Fprint(w, question, " ")
		g.transcript.WriteString(question + " ")

		response, err := g.GetUserYesOrNo(r, w)
		for err != nil {
			fmt.Fprintf(w, "Please answer yes or no: ")
			fmt.Fprint(w, question, " ")
			response, err = g.GetUserYesOrNo(r, w)
		}
		prev2 = prev1
		prev1 = question
		question, err = g.NextQuestion(question, response)
		if err != nil {
			fmt.Fprintf(w, "oh no, internal error! Not your fault!")
			return err
		}

		switch question {
		case "AnswerWin":
			fmt.Fprintf(w, "I successfully guessed your animal! Awesome!\n")
			g.transcript.WriteString("I successfully guessed your animal! Awesome!\n")
			g.Running = false
		case "AnswerLose":
			fmt.Fprintf(w, "You stumped me! Well done!\n")
			g.transcript.WriteString("You stumped me! Well done!\n")
			g.LearnNewAnimal(r, w, prev2)
			g.Running = false
		}
	}
	return nil
}

func (g game) Transcript() string {
	return g.transcript.String()
}

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
	ans, err := g.GetUserYesOrNo(strings.NewReader(scanner.Text()), w)

	for err != nil {
		g.Output(w, "Please answer yes or no: ")
		g.Output(w, fmt.Sprintf("What would be the answer to the question - \"%s\" for %s: ", qDistinctive, input))
		ans, err = g.GetUserYesOrNo(r, w)
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
		Yes: "AnswerWin",
		No:  "AnswerLose",
	}
}

// Replay - function to replay the game
func (g *game) Replay(r io.Reader, w io.Writer) bool {
	g.Output(w, "Would you like to play again (y/n)? ")
	var replay string

	replay, err := g.GetUserYesOrNo(r, w)
	for err != nil {
		g.Output(w, fmt.Sprintf("Error encountered %s", err))
		g.Output(w, "Would you like to play again (y/n)? ")
		replay, err = g.GetUserYesOrNo(r, w)
	}
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

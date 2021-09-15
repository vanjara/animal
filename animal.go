package animal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
func NewGame() *game {
	var StartingData map[string]Question
	mux.Lock()
	content, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	mux.Unlock()
	// Now let's unmarshall the data into StartingData
	err = json.Unmarshal(content, &StartingData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	//log.Printf("StartingData: %+v\n", StartingData)
	return &game{
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
func (g *game) Play(r io.Reader, w io.Writer) error {
	question := StartingQuestion
	var prev1, prev2 string
	fmt.Println("Playing Game")
	for g.Running {
		fmt.Fprint(w, question, " ")
		g.transcript.WriteString(question + " ")
		//log.Println(question)
		response, err := g.GetUserYesOrNo(r)
		for err != nil {
			fmt.Fprintln(w, "Please answer yes or no: ")
			fmt.Fprint(w, question, " ")
			response, err = g.GetUserYesOrNo(r)
		}
		prev2 = prev1
		prev1 = question
		question, err = g.NextQuestion(question, response)
		if err != nil {
			fmt.Fprintln(w, "oh no, internal error! Not your fault!")
			return err
		}

		switch question {
		case "AnswerWin":
			fmt.Fprintln(w, "I successfully guessed your animal! Awesome!")
			g.transcript.WriteString("I successfully guessed your animal! Awesome!")
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

// LearnNewAnimal - Function to learn the question to add for new animal
func (g *game) LearnNewAnimal(r io.Reader, w io.Writer, pq string) {
	var input string
	fmt.Fprintln(w, "Please tell me the animal you were thinking about: ")
	g.transcript.WriteString("Please tell me the animal you were thinking about: ")
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	input = scanner.Text()
	g.transcript.WriteString(input + "\n")

	fmt.Fprintf(w, "What would be a Yes/No question to distinguish %s from other animals: ", input)
	g.transcript.WriteString(fmt.Sprintf("What would be a Yes/No question to distinguish %s from other animals: ", input))
	scanner.Scan()
	qDistinctive := scanner.Text()
	g.transcript.WriteString(qDistinctive + "\n")

	fmt.Fprintf(w, "What would be the answer to the question - \"%s\" for %s: ", qDistinctive, input)
	g.transcript.WriteString(fmt.Sprintf("What would be the answer to the question - \"%s\" for %s: ", qDistinctive, input))

	scanner.Scan()
	ans, err2 := g.GetUserYesOrNo(strings.NewReader(scanner.Text()))

	for err2 != nil {
		fmt.Fprintln(w, "Please answer yes or no: ")
		fmt.Fprintf(w, "What would be the answer to the question - \"%s\" for %s: ", qDistinctive, input)
		ans, err2 = g.GetUserYesOrNo(r)
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
	//fmt.Printf("After if the game %+v", g.Data)
	g.Data[pq] = qPrevious

	g.Data[qNewAnimal] = Question{
		Yes: "AnswerWin",
		No:  "AnswerLose",
	}
	// Since we have learned the new animal, let's write it to a file
	// mux.Lock()
	// file, _ := json.MarshalIndent(g.Data, "", " ")
	// _ = ioutil.WriteFile("data.json", file, 0644)
	// mux.Unlock()
}

// Replay - function to replay the game
func (g *game) Replay(r io.Reader) bool {
	fmt.Print("Would you like to play again (y/n)? ")
	g.transcript.WriteString("Would you like to play again (y/n)? ")
	var replay string
	// ToDo - Do we need to check the error here
	// We should keep prompting until we get expected input
	// Add invalid responses to the test for GetUserYesOrNo
	replay, _ = g.GetUserYesOrNo(r)
	if replay == AnswerYes {
		return true
	}
	fmt.Println("Thanks for playing!")
	g.transcript.WriteString("Thanks for playing!")
	return false
}

// GetUserYesOrNo - function to map variations of yes/no responses
func (g *game) GetUserYesOrNo(r io.Reader) (string, error) {
	var input string
	_, err := fmt.Fscanln(r, &input)

	if err != nil {
		return "", err
	}
	g.transcript.WriteString(input + "\n")
	switch input {
	case "yes", "y", "YES", "Yes":
		return AnswerYes, nil
	case "no", "n", "NO", "No":
		return AnswerNo, nil
	default:
		return "", fmt.Errorf("Unexpected input: %q", input)
	}
}

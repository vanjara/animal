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

// type question struct{
// 	q   string,
// 	id  int,
// 	yes int, // what question is next, if answer is yes
// 	no  int // what question is next, if answer is no
// }

// var questions = []question{
// 	{
// 		q: "Does it have 4 legs",
// 		id: 1
// 		yes: 2
// 		no: 3
// 	}
// 	{
// 		q: "Is it a horse?",
// 		id: 2
// 		yes: 0 // how do we signify a win?
// 		no: 3
// 	}
// 	{
// 		q: "Is it a snake?",
// 		id: 3
// 		yes: 0 // how do we signify a win?
// 		no: 100 // what if there are no more questions, user wins
// 	}
// }

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

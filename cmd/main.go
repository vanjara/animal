package main

import (
	"animal"
	"fmt"
	"os"
)

func main() {
	g := animal.NewGame()
	question := animal.StartingQuestion
	for g.Running {
		fmt.Println(question)
		response, err := animal.GetUserYesOrNo(question, os.Stdin)
		for err != nil {
			fmt.Println("please answer yes or no")
			fmt.Println(question)
			response, err = animal.GetUserYesOrNo(question, os.Stdin)
		}
		question, err = g.NextQuestion(question, response)
		if err != nil {
			fmt.Println("oh no, internal error! Not your fault!")
			g.Running = false
		}
		switch question {
		case animal.AnswerWin:
			fmt.Println("I successfully guessed your animal! Awesome!")
			g.Running = false
		case animal.AnswerLose:
			fmt.Println("You stumped me! Well done!")
			g.Running = false
		}
	}
	fmt.Println("Thanks for playing!")
}

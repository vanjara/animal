package animal

import (
	"bufio"
	"fmt"
	"os"
)

type game struct {
	Running bool
}

func New() game {
	return game{
		Running: true,
	}
}

func AskUserYesOrNo(question string) string {
	//Extend test to not pass by just returning the question
	//enter muliple test cases
	if question == "Is it a horse" {
		return "yes"
	}
	return "no"
}

//func GetUserYesOrNo(question string) string {
func GetUserYesOrNo(question string) string {
	//question := "Is it a horse?"
	fmt.Printf(question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %s\n", input)
	if input == "yes" {
		//fmt.Printf("Returning: %s", input)
		return "yes"
	}
	fmt.Printf("Returning: no because input is %s", input)
	return "no"
}

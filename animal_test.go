package animal_test

import (
	"animal"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNewGame(t *testing.T) {
	t.Parallel()
	a := animal.NewGame()
	if !a.Running {
		t.Errorf("want a.Running == true")
	}
}

func TestGetUserYesOrNo(t *testing.T) {
	t.Parallel()
	// multiple test cases
	testCases := []struct {
		input         string
		want          string
		errorExpected bool
	}{
		{input: "yes", want: "yes"},
		{input: "y", want: "yes"},
		{input: "YES", want: "yes"},
		{input: "Yes", want: "yes"},
		{input: "no", want: "no"},
		{input: "n", want: "no"},
		{input: "NO", want: "no"},
		{input: "No", want: "no"},
		{input: "", want: "", errorExpected: true},
		{input: "Bogus", want: "no", errorExpected: true},
	}
	for _, tc := range testCases {
		got, err := animal.GetUserYesOrNo("Dummy question?", strings.NewReader(tc.input))
		if tc.errorExpected != (err != nil) {
			t.Fatalf("Give input %q, unexpected error Status: %v", tc.input, err)
		}
		if !tc.errorExpected && tc.want != got {
			t.Errorf("Given input: %q, want %q, got %q\n", tc.input, tc.want, got)
		}
	}
}

func TestMultipleUserInput(t *testing.T) {
	t.Parallel()
	// multiple test cases
	var x = `yes
	yes`
	//input := strings.NewReader("yes\nyes\n")
	input := strings.NewReader(x)
	fmt.Println("From TEST ", input)
	_, err := animal.GetUserYesOrNo("Dummy question?", input)
	if err != nil {
		t.Fatalf("Unexpected error Status: %v", err)
	}
	_, err = animal.GetUserYesOrNo("Dummy question?", input)
	if err != nil {
		t.Fatalf("Unexpected error Status: %v", err)
	}
}

func TestNextQuestion(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		question    string
		response    string
		want        string
		errExpected bool
	}{
		{
			question:    "Does it have 4 legs?",
			response:    animal.AnswerYes,
			want:        "Does it have stripes?",
			errExpected: false,
		},
		{
			question:    "Does it have 4 legs?",
			response:    animal.AnswerNo,
			want:        "Is it carnivorous?",
			errExpected: false,
		},
		{
			question:    "Does it have stripes?",
			response:    animal.AnswerYes,
			want:        "Is it a zebra?",
			errExpected: false,
		},
		{
			question:    "Does it have stripes?",
			response:    animal.AnswerNo,
			want:        "Is it a lion?",
			errExpected: false,
		},
		{
			question:    "Is it a snake?",
			response:    animal.AnswerYes,
			want:        animal.AnswerWin,
			errExpected: false,
		},
		{
			question:    "Is it a giraffe?",
			response:    animal.AnswerNo,
			want:        animal.AnswerLose,
			errExpected: false,
		},
		{
			question:    "Is it a lion?",
			response:    animal.AnswerYes,
			want:        animal.AnswerWin,
			errExpected: false,
		},
		{
			question:    "Is it a bogus non-existent animal?",
			errExpected: true,
		},
	}
	testGame := animal.NewGame()
	for _, tc := range testCases {
		got, err := testGame.NextQuestion(tc.question, tc.response)
		if tc.errExpected != (err != nil) {
			t.Fatalf("Given input: %q, unexpected error status: %v", tc.question, err)
		}
		if tc.want != got {
			t.Errorf("Given input: %q, response: %q, want: %q, got: %q\n", tc.question, tc.response, tc.want, got)
		}
	}
}

func TestPlay(t *testing.T) {
	testGame := animal.NewGame()
	input := strings.NewReader("yes\nyes\nyes\n")
	output := os.Stdout
	err := testGame.Play(input, output)
	if err != nil {
		t.Error(err)
	}
	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed, data still left to consume %d\n", input.Len())
	}

}

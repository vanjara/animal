package animal_test

import (
	"animal"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewGame(t *testing.T) {
	t.Parallel()
	a := animal.NewGame()
	if !a.Running {
		t.Errorf("want a.Running == true")
	}
}

func TestPlay(t *testing.T) {
	t.Parallel()
	testGame := animal.NewGame()
	input := strings.NewReader("yes\nyes\nyes\n")
	err := testGame.Play(input, ioutil.Discard)
	if err != nil {
		t.Error(err)
	}
	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed, data still left to consume %d\n", input.Len())
	}
}

func TestTranscript(t *testing.T) {
	t.Parallel()
	testGame := animal.NewGame()
	input := strings.NewReader("yes\nyes\nno\ntiger\nIs it a predator?\nyes\n")
	err := testGame.Play(input, ioutil.Discard)
	if err != nil {
		t.Error(err)
	}
	got := testGame.Transcript()
	want := `Does it have 4 legs? yes
Does it have stripes? yes
Is it a zebra? no
You stumped me! Well done!
Please tell me the animal you were thinking about: tiger
What would be a Yes/No question to distinguish tiger from other animals: Is it a predator?
What would be the answer to the question - "Is it a predator?" for tiger: yes
`

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
func TestGetUserYesOrNo(t *testing.T) {
	t.Parallel()
	testGame := animal.NewGame()
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
		got, err := testGame.GetUserYesOrNo(strings.NewReader(tc.input))
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
	testGame := animal.NewGame()
	// multiple test cases
	input := strings.NewReader("yes\nyes\n")
	_, err := testGame.GetUserYesOrNo(input)
	if err != nil {
		t.Fatalf("Unexpected error Status: %v", err)
	}
	_, err2 := testGame.GetUserYesOrNo(input)
	if err2 != nil {
		t.Fatalf("Unexpected error Status: %v", err2)
	}
}

func TestNextQuestion(t *testing.T) {
	t.Parallel()
	testGame := animal.NewGame()
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
			want:        "AnswerWin",
			errExpected: false,
		},
		{
			question:    "Is it a giraffe?",
			response:    animal.AnswerNo,
			want:        "AnswerLose",
			errExpected: false,
		},
		{
			question:    "Is it a lion?",
			response:    animal.AnswerYes,
			want:        "AnswerWin",
			errExpected: false,
		},
		{
			question:    "Is it a bogus non-existent animal?",
			errExpected: true,
		},
	}
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

func TestPlayNewAnimalAnswerYes(t *testing.T) {
	t.Parallel()
	testGame := animal.NewGame()
	input := strings.NewReader("yes\nyes\nno\ntiger\nIs it a predator?\nyes\n")

	err := testGame.Play(input, ioutil.Discard)
	if err != nil {
		t.Error(err)
	}
	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed, data still left to consume %d\n", input.Len())
	}
	want := "Is it a predator?"
	if _, ok := testGame.Data[want]; !ok {
		t.Errorf("Expected %q, did not find the question in the map.", want)
	}
	want = "Is it a tiger?"
	if _, ok := testGame.Data[want]; !ok {
		t.Errorf("Expected %q, did not find the question in the map.", want)
	}

	input2 := strings.NewReader("yes\nyes\nyes\nyes\n")
	testGame.Running = true
	err2 := testGame.Play(input2, ioutil.Discard)
	if err2 != nil {
		t.Error(err)
	}
	if input2.Len() != 0 {
		t.Errorf("Input 2: Given input not fully consumed, data still left to consume %d, %v and transcript is \n%s\n", input2.Len(), input2.UnreadByte(), testGame.Transcript())
	}

	if got, ok := testGame.Data[want]; !ok {
		t.Errorf("Expected %q, got %q.", want, got)
	} else {
		fmt.Printf("got %+v", got)
	}

}

func TestPlayNewAnimalAnswerNo(t *testing.T) {
	t.Parallel()
	testGame := animal.NewGame()
	input := strings.NewReader("no\nno\nno\noctopus\nIs it a land animal?\nno\n")

	err := testGame.Play(input, ioutil.Discard)
	if err != nil {
		t.Error(err)
	}
	if input.Len() != 0 {
		t.Errorf("Given input not fully consumed, data still left to consume %d\n", input.Len())
	}
	want := "Is it a land animal?"
	if _, ok := testGame.Data[want]; !ok {
		t.Errorf("Expected %q, did not find the question in the map.", want)
	}
	want = "Is it a octopus?"
	if _, ok := testGame.Data[want]; !ok {
		t.Errorf("Expected %q, did not find the question in the map.", want)
	}

	input2 := strings.NewReader("no\nno\nno\nyes\n")
	testGame.Running = true
	err2 := testGame.Play(input2, ioutil.Discard)
	if err2 != nil {
		t.Error(err)
	}
	if input2.Len() != 0 {
		t.Errorf("Input 2: Given input not fully consumed, data still left to consume %d\n", input2.Len())
	}
	if got, ok := testGame.Data[want]; !ok {
		t.Errorf("Expected %q, got %q.", want, got)
	} else {
		fmt.Printf("got %+v", got)
	}
}

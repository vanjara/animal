package animal_test

import (
	"animal"
	"io/ioutil"
	"strings"
	"testing"
	"text/template/parse"
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

func TestNextQuestion(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		question string
		response string
		want     string
		errExpected bool
	}{
		{
			question: "Does it have 4 legs?",
			response: animal.AnswerYes,
			want:     "Does it have stripes?",
			errExpected: false,
		},
		{
			question: "Does it have 4 legs?",
			response: animal.AnswerNo,
			want:     "Is it carnivorous?",
			errExpected: false,
		},
		{
			question: "Does it have stripes?",
			response: animal.AnswerYes,
			want:     "Is it a zebra?",
			errExpected: false,
		},
		{
			question: "Does it have stripes?",
			response: animal.AnswerNo,
			want:     "Is it a lion?",
			errExpected: false,
		},
		{
			question: "Is it a snake?",
			response: animal.AnswerYes,
			want:     animal.AnswerWin,
			errExpected: false,
		},
		{
			question: "Is it a giraffe?",
			response: animal.AnswerNo,
			want:     animal.AnswerLose,
			errExpected: false,
		},
		{
			question: "Is it a lion?",
			response: animal.AnswerYes,
			want:     animal.AnswerWin,
			errExpected: false,
		},
		{
			question: "Is it a bogus non-existent animal?",
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
	reader := something that always responds with 'y'
	writer := ioutil.Discard // just throw away the game output
	err := animal.Play(reader, writer)
	if err != nil {
		t.Error(err)
	}
}
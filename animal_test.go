package animal_test

import (
	"animal"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
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
	}{
		{
			question: "Does it have 4 legs?",
			response: animal.AnswerYes,
			want:     "Does it have stripes?",
		},
		{
			question: "Does it have 4 legs?",
			response: animal.AnswerNo,
			want:     "Is it a snake?",
		},
		{
			question: "Does it have stripes?",
			response: animal.AnswerYes,
			want:     "Is it a zebra?",
		},
		{
			question: "Does it have stripes?",
			response: animal.AnswerNo,
			want:     "Is it a lion?",
		},
	}
	testGame := animal.NewGame()
	testGame.Data = map[string]animal.Question{
		"Does it have 4 legs?": animal.Question{
			Yes: "Does it have stripes?",
			No:  "Is it a snake?",
		},
		"Does it have stripes?": animal.Question{
			Yes: "Is it a zebra?",
			No:  "Is it a lion",
		},
	}
	for _, tc := range testCases {
		got := animal.NextQuestion(tc.question, tc.response)
		if tc.want != got {
			t.Errorf("Given input: %q, response: %q, want: %q, got: %q\n", tc.question, tc.response, tc.want, got)
		}
	}
}

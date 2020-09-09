package animal_test

import (
	"animal"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	a := animal.New()
	if !a.Running {
		t.Errorf("want a.Running == true")
	}
}

func TestAskUserYesOrNo(t *testing.T) {
	t.Parallel()
	/*want := "yes"
	got := animal.AskUserYesOrNo("Is it a horse")
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}*/
	testCases := []struct {
		question string
		want     string
	}{
		{question: "Is it a horse", want: "yes"},
		{question: "Is it yes or no", want: "no"},
		{question: "Does it have 2 legs", want: "no"},
	}
	for _, testCase := range testCases {
		got := animal.AskUserYesOrNo(testCase.question)
		if testCase.want != got {
			t.Errorf("want %q, got %q", testCase.want, got)
		}
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

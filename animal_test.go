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

func TestGameQuestions(t *testing.T) {
	t.Parallel()
	input := animal.Question{
		Q:   "Does it have 4 legs",
		Id:  0,
		Yes: 1,
		No:  1,
	}
	var errorExpected bool
	want := 1
	got, err := animal.GameQuestions(input)
	if errorExpected != (err != nil) {
		t.Fatalf("Give input %q, unexpected error Status: %v", input.Q, err)
	}
	if !errorExpected && want != got {
		t.Errorf("Given input: %q, want %q, got %q\n", input.Q, want, got)
	}
	// testCases := []struct {
	// 	input1        string
	// 	yes           int
	// 	no            int
	// 	want          int
	// 	errorExpected bool
	// }{
	// 	{input1: "Does it have 4 legs", yes: 2, want: 2},
	// 	{input1: "Is it a horse", yes: 0, want: 0},
	// }
	// for _, tc := range testCases {
	// 	got, err := animal.GameQuestions(input)
	// 	if tc.errorExpected != (err != nil) {
	// 		t.Fatalf("Give input %q, unexpected error Status: %v", tc.input1, err)
	// 	}
	// 	if !tc.errorExpected && tc.want != got {
	// 		t.Errorf("Given input: %q, want %q, got %q\n", tc.input1, tc.want, got)
	// 	}
	// }
}

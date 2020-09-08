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
	// A single specific test
	testIn := strings.NewReader("yes\n")
	testQuestion := "Is it a horse?\n"
	testData := animal.Data{testQuestion, testIn}
	got := animal.GetUserYesOrNo(testData)
	want := "yes"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// multiple test cases
	testCases := []struct {
		question string
		want     string
	}{
		{question: "Is it a horse?\n", want: "yes"},
		{question: "Is it yes or no?\n", want: "yes"},
		{question: "Does it have 2 legs?\n", want: "yes"},
	}
	for _, testCase := range testCases {
		testData := animal.Data{testCase.question, strings.NewReader("yes\n")}
		got := animal.GetUserYesOrNo(testData)
		if testCase.want != got {
			t.Errorf("want %q, got %q\n", testCase.want, got)
		}
	}
}

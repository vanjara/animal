package animal_test

import (
	"animal"
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

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

package cliValidate

import (
	"testing"
)

func TestValidateGet(t *testing.T) {
	input := []string{"mark", "t1,t2"}
	expected := "Incorrect number of arguments type 'help' for more information"
	valid, actual := ValidateGet(input)

	if !valid {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}
}

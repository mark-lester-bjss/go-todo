package cli

import (
	"testing"
)

func TestErrorsAreFormatedCorrectly(t *testing.T) {
	input := "command name some todo, another todo"
	expected := "command"
	commandName, params := separtateParams(input)

	if commandName == expected || params != nil {
		t.Errorf("actual %q expected %q are not equal", commandName, expected)
	}
}

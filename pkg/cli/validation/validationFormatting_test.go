package cliValidate

import (
	"reflect"
	"testing"
)

func TestErrorsAreFormatedCorrectly(t *testing.T) {
	input := []string{"s1", "s2"}
	expected := "s1\ns2"
	actual := formatValidationErrors(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}
}

package cliValidate

import (
	"testing"
)

func TestValidateDelete(t *testing.T) {
	type test struct {
		input []string
		wantA bool
		wantB string
	}

	tests := []test{
		{input: []string{"mark", "3f296dbb-600c-4240-89fd-6e642bbb841b"},
			wantA: true,
			wantB: ""},
		{input: []string{"mark"},
			wantA: false,
			wantB: "Incorrect number of arguments type 'help' for more information"},
		{input: []string{"mark", "not a uuid"},
			wantA: false,
			wantB: "The uuid supplied is not valid. Type help for more details"},
	}

	for _, tc := range tests {
		valid, message := ValidateDelete(tc.input)
		if valid != tc.wantA && message != tc.wantB {
			t.Errorf("The validation should be %t with a message: '%s' but was %t with '%s'", tc.wantA, tc.wantB, valid, message)
		}
	}

}

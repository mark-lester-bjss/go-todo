package cliValidate

import (
	"github.com/google/uuid"
)

func ValidateDelete(args []string) (valid bool, errors string) {
	messages := []string{}

	valid, message := hasCorrectNumberOfArgs(args)
	if !valid {
		messages = append(messages, message)
	} else {
		err := uuid.Validate(args[1])
		if err != nil {
			messages = append(messages, "The uuid supplied is not valid. Type help for more details")
			valid = false
		}
	}

	return valid, formatValidationErrors(messages)
}

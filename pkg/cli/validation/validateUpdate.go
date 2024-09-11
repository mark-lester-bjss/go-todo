package cliValidate

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var pl = fmt.Println

func ValidateUpdate(args []string) (valid bool, errors string) {
	messages := []string{}

	valid, message := hasCorrectNumberOfArgs(args)
	if !valid {
		messages = append(messages, message)
	} else {
		valid = strings.Contains(args[1], ":")
		if !valid {
			messages = append(messages, "The args should be a uuid and todo separated by a colon, uuid:todo")
		} else {
			err := uuid.Validate(strings.Split(args[1], ":")[0])
			if err != nil {
				messages = append(messages, "The uuid supplied is not valid. Type help for more details")
				valid = false
			}
		}
	}

	pl("The messages", messages)
	return valid, formatValidationErrors(messages)
}

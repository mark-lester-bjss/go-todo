package cliValidate

func ValidateCreate(args []string) (valid bool, errors string) {
	messages := []string{}

	valid, message := hasCorrectNumberOfArgs(args)
	if !valid {
		messages = append(messages, message)
	}

	return valid, formatValidationErrors(messages)
}

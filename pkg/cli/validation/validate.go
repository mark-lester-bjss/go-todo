package cliValidate

func hasCorrectNumberOfArgs(args []string) (valid bool, message string) {
	valid = true

	if len(args) != 2 {
		message = "Incorrect number of arguments type 'help' for more information"
		valid = false
	}

	return valid, message
}

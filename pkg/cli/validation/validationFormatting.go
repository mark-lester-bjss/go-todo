package cliValidate

import "strings"

func formatValidationErrors(messages []string) (response string) {
	var sb strings.Builder
	for index, string := range messages {
		if index != 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(string)

	}
	return sb.String()
}

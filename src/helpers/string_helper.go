package helpers

import "strings"

func TrimStartAndEnd(text string) string {
	return strings.Trim(text, " ")
}

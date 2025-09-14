package logs

import (
	"strings"
	"unicode/utf8"
)

func Application(log string) string {

	for _, char := range log {

		switch char {
		case '\u2757':
			return "recommendation"
		case '\U0001F50D':
			return "search"
		case '\u2600':
			return "weather"
		}

	}

	return "default"

}

func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

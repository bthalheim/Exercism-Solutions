package logs

import (
	"strings"
	"unicode/utf8"
)

func Application(log string) string {

	apps := map[string]rune{
		"recommendation": '\u2757',
		"search":         '\U0001F50D',
		"weather":        '\u2600',
	}

	for _, char := range log {

		for k, v := range apps {

			if char == v {
				return k
			}

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

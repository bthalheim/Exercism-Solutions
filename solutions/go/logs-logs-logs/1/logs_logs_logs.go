package logs

import "unicode/utf8"

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

	runes := []rune(log)

	for idx, char := range runes {
		if char == oldRune {
			runes[idx] = newRune
		}
	}

	return string(runes)
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

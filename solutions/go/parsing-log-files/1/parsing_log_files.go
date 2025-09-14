package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~\-=*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (count int) {

	re := regexp.MustCompile(`\".*password.*\"`)

	for _, line := range lines {
		if re.MatchString(strings.ToLower(line)) {
			count++
		}
	}

	return

}

func RemoveEndOfLineText(text string) string {

	re := regexp.MustCompile(`end-of-line\d+`)

	return re.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {

	ret := make([]string, 0)

	re := regexp.MustCompile(`User\s+([a-zA-Z0-9]+)\s`)

	for _, line := range lines {

		sl := re.FindStringSubmatch(line)

		if sl != nil {
			ret = append(ret, fmt.Sprintf("[USR] %s %s", sl[1], line))
		} else {
			ret = append(ret, line)
		}
	}

	return ret

}

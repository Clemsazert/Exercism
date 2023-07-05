package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`".*password.*"`)
	for _, line := range lines {
		if re.MatchString(strings.ToLower(line)) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reUser := regexp.MustCompile("User +([a-zA-Z0-9]*)")
	for i := 0; i < len(lines); i++ {
		if userString := reUser.FindString(lines[i]); userString != "" {
			split := strings.Split(userString, " ")
			name := split[len(split)-1]
			lines[i] = fmt.Sprintf("[USR] %s %s", name, lines[i])
		}
	}
	return lines
}

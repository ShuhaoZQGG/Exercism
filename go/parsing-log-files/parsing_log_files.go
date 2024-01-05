package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[\~\*\=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(?i).*password.*"`)
	counter := 0
	for _, line := range lines {
		if re.MatchString(line) {
			counter += 1
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line+\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			lines[i] = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
	}
	return lines
}

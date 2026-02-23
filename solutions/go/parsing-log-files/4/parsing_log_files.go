package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	if err != nil {
		fmt.Println(err)
	}

	return re.Match([]byte(text))

}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(` ?<[(~|\*|=|\-)]*> ?`)
	if err != nil {
		fmt.Println(err)
	}

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`(?i)"[^"]*password[^"]*"`)
	if err != nil {
		fmt.Println(err)
	}
	var counter int
	for _, v := range lines {
		if re.Match([]byte(v)) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line[\d]*`)
	if err != nil {
		fmt.Println(err)
	}
	return string(re.ReplaceAll([]byte(text), []byte{}))
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User +[^ ]+`)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range lines {
		if re.Match([]byte(v)) {
			matched := re.FindStringSubmatch(v)
			fmt.Println("\n", matched)
		}
	}
	return lines
}

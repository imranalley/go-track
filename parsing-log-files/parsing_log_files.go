package parsinglogfiles

import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[-=~*]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)".*password.*"`)
	count := 0
	for _,line := range lines {
		if re.MatchString(line){
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+(\w+)`)
	str := []string{}
	for _,line:=range lines{
		find := re.FindStringSubmatch(line)
		if find !=nil{
			line = fmt.Sprintf("[USR] %s %s", find[1], line)
		}
		str = append(str,line)
	}
	return str
}

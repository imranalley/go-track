package parsinglogfiles

// import "fmt"
import "regexp"

func IsValidLine(text string) bool {
	// logLines := []string{`[TRC]`, `[DBG]`, `[INF]`, `[WRN]`, `[ERR]`, `[FTL]`}
	re, _ := regexp.Compile(`[TRC]+[DBG]+[INF]+[WRN]+[ERR]+[FTL]`)

	b := re.MatchString(text)

	return b
}

func SplitLogLine(text string) []string {
	panic("Please implement the SplitLogLine function")
}

func CountQuotedPasswords(lines []string) int {
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}

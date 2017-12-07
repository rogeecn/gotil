package str

import (
	"strings"
	"fmt"
)

func Format(format string, m map[string]interface{}) string {
	for k, v := range m {
		format = strings.Replace(format, "{"+k+"}", fmt.Sprintf("%v", v), -1)
	}
	return format
}

func Reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

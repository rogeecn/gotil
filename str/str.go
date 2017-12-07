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


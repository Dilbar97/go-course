package unpack_string

import (
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) string {
	var result []string

	for i, v := range s {
		if _, err := strconv.Atoi(s); err == nil {
			return ""
		}

		if i == 0 && unicode.IsDigit(v) {
			return ""
		}

		if unicode.IsDigit(v) && unicode.IsDigit(rune(s[i-1])) {
			return ""
		}

		if unicode.IsDigit(v) && int(v-'0') == 0 {
			result = result[:len(result)-1]
			result = append(result, string(v))
			continue
		}

		if unicode.IsDigit(v) && i > 0 {
			multilied := strings.Repeat(string(s[i-1]), int(v-'1'))
			result = append(result, multilied)
		} else {
			result = append(result, string(v))
		}
	}

	return strings.Join(result, "")
}


package utils

import "strings"

func Capitalize(s string) string {
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

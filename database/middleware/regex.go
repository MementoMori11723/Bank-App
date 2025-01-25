package middleware

import (
	"regexp"
	"strings"
)

func CheckString(s string) bool {
	if strings.Contains(s, "--") {
		return false
	}

  if strings.Contains(s, " ") {
    return false
  }

	pattern := `^[\w\s!@^$*-.]*$`
	res := regexp.MustCompile(pattern).MatchString(s)
	return res
}

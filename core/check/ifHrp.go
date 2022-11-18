package check

import (
	"regexp"
)

// IfHrp tests if the provided string is human readable - does contain only alphabetic characters
func IfHrp(s string) bool {
	isHrp := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s)

	return isHrp
}

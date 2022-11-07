package check

import (
	"strings"
)

// IfEmpty tests if the provided string is technically empty
func IfEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

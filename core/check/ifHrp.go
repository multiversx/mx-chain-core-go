package check

// IfHrp tests if the provided string is human readable - does contain only alphabetic characters
func IfHrp(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

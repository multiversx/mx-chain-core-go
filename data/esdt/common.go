package esdt

import "strings"

const (
	// esdtTickerNumRandChars represents the number of hex-encoded random characters sequence of a ticker
	esdtTickerNumRandChars = 6
	// separatorChar represents the character that separated the token ticker by the random sequence
	separatorChar = "-"
	// minLengthForTickerName represents the minimum number of characters a token's ticker can have
	minLengthForTickerName = 3
	// maxLengthForTickerName represents the maximum number of characters a token's ticker can have
	maxLengthForTickerName = 10
	// maxLengthESDTPrefix represents the maximum number of characters a token's prefix can have
	maxLengthESDTPrefix = 4
	// minLengthESDTPrefix represents the minimum number of characters a token's prefix can have
	minLengthESDTPrefix = 1
)

// IsValidPrefixedToken checks if the provided token is valid, and returns if prefix if so
func IsValidPrefixedToken(token string) (string, bool) {
	tokenSplit := strings.Split(token, separatorChar)
	if len(tokenSplit) < 3 {
		return "", false
	}

	prefix := tokenSplit[0]
	if !IsValidTokenPrefix(prefix) {
		return "", false
	}

	tokenTicker := tokenSplit[1]
	if !IsTickerValid(tokenTicker) {
		return "", false
	}

	tokenRandSeq := tokenSplit[2]
	if !IsRandomSeqValid(tokenRandSeq) {
		return "", false
	}

	return prefix, true
}

// IsValidTokenPrefix checks if the token prefix is valid
func IsValidTokenPrefix(prefix string) bool {
	prefixLen := len(prefix)
	if prefixLen > maxLengthESDTPrefix || prefixLen < minLengthESDTPrefix {
		return false
	}

	for _, ch := range prefix {
		isLowerCaseCharacter := ch >= 'a' && ch <= 'z'
		isNumber := ch >= '0' && ch <= '9'
		isAllowedChar := isLowerCaseCharacter || isNumber
		if !isAllowedChar {
			return false
		}
	}

	return true
}

// IsTickerValid checks if the token ticker is valid
func IsTickerValid(ticker string) bool {
	if !IsTokenTickerLenCorrect(len(ticker)) {
		return false
	}

	for _, ch := range ticker {
		isUpperCaseCharacter := ch >= 'A' && ch <= 'Z'
		isNumber := ch >= '0' && ch <= '9'
		isAllowedChar := isUpperCaseCharacter || isNumber
		if !isAllowedChar {
			return false
		}
	}

	return true
}

// IsTokenTickerLenCorrect checks if the token ticker len is correct
func IsTokenTickerLenCorrect(tokenTickerLen int) bool {
	return !(tokenTickerLen < minLengthForTickerName || tokenTickerLen > maxLengthForTickerName)
}

// IsRandomSeqValid checks if the token random sequence is valid
func IsRandomSeqValid(randomSeq string) bool {
	if len(randomSeq) != esdtTickerNumRandChars {
		return false
	}

	for _, ch := range randomSeq {
		isSmallCharacter := ch >= 'a' && ch <= 'f'
		isNumber := ch >= '0' && ch <= '9'
		isReadable := isSmallCharacter || isNumber
		if !isReadable {
			return false
		}
	}

	return true
}

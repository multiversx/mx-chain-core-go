package esdt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidPrefixedToken(t *testing.T) {
	prefix, valid := IsValidPrefixedToken("sov1-TKN-coffee")
	require.True(t, valid)
	require.Equal(t, "sov1", prefix)

	prefix, valid = IsValidPrefixedToken("sOv1-TKN-coffee")
	require.False(t, valid)
	require.Equal(t, "", prefix)

	prefix, valid = IsValidPrefixedToken("sov1-TkN-coffee")
	require.False(t, valid)
	require.Equal(t, "", prefix)

	prefix, valid = IsValidPrefixedToken("sov1-TKN-coffe")
	require.False(t, valid)
	require.Equal(t, "", prefix)

	prefix, valid = IsValidPrefixedToken("sov1-TKN")
	require.False(t, valid)
	require.Equal(t, "", prefix)

	prefix, valid = IsValidPrefixedToken("TKN-coffee")
	require.False(t, valid)
	require.Equal(t, "", prefix)
}

func TestIsValidTokenPrefix(t *testing.T) {
	require.False(t, IsValidTokenPrefix(""))
	require.False(t, IsValidTokenPrefix("-"))
	require.False(t, IsValidTokenPrefix("prefix"))
	require.False(t, IsValidTokenPrefix("prefi"))
	require.False(t, IsValidTokenPrefix("Prfx"))
	require.False(t, IsValidTokenPrefix("pX4"))
	require.False(t, IsValidTokenPrefix("px-4"))

	require.True(t, IsValidTokenPrefix("pref"))
	require.True(t, IsValidTokenPrefix("sv1"))
}

func TestIsTickerValid(t *testing.T) {
	require.False(t, IsTickerValid("TK"))
	require.False(t, IsTickerValid("TKn"))
	require.False(t, IsTickerValid("T0KEN-"))

	require.True(t, IsTickerValid("T0KEN"))
}

func TestIsTokenTickerLenCorrect(t *testing.T) {
	require.False(t, IsTokenTickerLenCorrect(len("TOKENALICEALICE")))
	require.False(t, IsTokenTickerLenCorrect(len("AL")))

	require.True(t, IsTokenTickerLenCorrect(len("ALC")))
	require.True(t, IsTokenTickerLenCorrect(len("ALICE")))
}

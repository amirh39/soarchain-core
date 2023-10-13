package utility

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidHex(t *testing.T) {
	t.Run("Valid Hexadecimal Input", func(t *testing.T) {
		hexValue := "BE1FA813"
		valid := IsValidHex(hexValue)
		require.True(t, valid)
	})

	t.Run("Valid Hexadecimal Input with Lowercase Characters", func(t *testing.T) {
		hexValue := "be1fa813"
		valid := IsValidHex(hexValue)
		require.True(t, valid)
	})

	t.Run("Invalid Hexadecimal Input", func(t *testing.T) {
		hexValue := "InvalidHex"
		valid := IsValidHex(hexValue)
		require.False(t, valid)
	})

	t.Run("Empty Input", func(t *testing.T) {
		hexValue := ""
		valid := IsValidHex(hexValue)
		require.False(t, valid)
	})
}

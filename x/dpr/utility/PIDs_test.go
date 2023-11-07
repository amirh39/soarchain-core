package utility

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexToBin(t *testing.T) {
	t.Run("Valid Hexadecimal Input", func(t *testing.T) {
		hexValue := "BE1FA813"
		binary, err := hexToBin(hexValue)
		require.NoError(t, err)
		require.Equal(t, "10111110000111111010100000010011", binary)
	})

	t.Run("Invalid Hexadecimal Input", func(t *testing.T) {
		hexValue := "InvalidHex"
		_, err := hexToBin(hexValue)
		require.Error(t, err)
	})
}

func TestArePIDsSupported(t *testing.T) {
	t.Run("All PIDs Supported", func(t *testing.T) {
		carHex := "FFFFFFFFFFFFFFF"
		dprHex := "AAAAAAAA"
		supported, err := ArePIDsSupported(carHex, dprHex)
		require.NoError(t, err)
		require.True(t, supported)
	})

	t.Run("Not All PIDs Supported", func(t *testing.T) {
		carHex := "00000000"
		dprHex := "AAAAAAAA"
		supported, err := ArePIDsSupported(carHex, dprHex)
		require.Error(t, err)
		require.False(t, supported)
	})
}

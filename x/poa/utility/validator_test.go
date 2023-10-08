package utility

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidString(t *testing.T) {
	r := require.New(t)

	stringInput := "string"
	result := ValidString(stringInput)
	r.Equal(result, true)
}

func Test_NotValidString(t *testing.T) {
	r := require.New(t)

	emptyString := ""
	r.Equal(ValidString(emptyString), false)
}

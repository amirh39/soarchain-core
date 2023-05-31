package utility

import (
	"soarchain/x/poa/utility/utilConstants"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidPubkey(t *testing.T) {
	r := require.New(t)

	pubkey := utilConstants.PubKey
	result := ValidPubkey(pubkey)
	r.Equal(result, true)

}

func Test_NotValidPubkey(t *testing.T) {
	r := require.New(t)

	emptyPublicKey := ""
	r.Equal(ValidPubkey(emptyPublicKey), false)
}

func Test_ValidAddress(t *testing.T) {
	r := require.New(t)

	address := utilConstants.Address
	result := ValidAddress(address)
	r.Equal(result, true)
	r.Equal(len(address), 43)
	r.Equal(address[0:4], "soar")
}

func Test_NotValidAddress(t *testing.T) {
	r := require.New(t)

	emptyAddress := ""
	r.Equal(ValidAddress(emptyAddress), false)

	shortAddress := ""
	r.Equal(ValidAddress(shortAddress), false)
}

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

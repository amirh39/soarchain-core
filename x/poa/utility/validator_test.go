package utility

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidPubkey(t *testing.T) {
	r := require.New(t)

	pubkey := "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251"
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

	address := "soar1uajy2t7tyamnuqms7l65ka4wtwgrvey0rve34t"
	result := ValidAddress(address)
	r.Equal(result, true)
	r.Equal(len(address), 43)
	r.Equal(address[0:4], "soar")
}

func Test_NotValidAddress(t *testing.T) {
	r := require.New(t)

	emptyAddress := ""
	r.Equal(ValidAddress(emptyAddress), false)

	shortAddress := "soar1uajy2t7tyamnuqms7l65ka4wtwgrvey0rve34"
	r.Equal(ValidAddress(shortAddress), false)

	firstNotEqualAddress := "doar1uajy2t7tyamnuqms7l65ka4wtwgrvey0rve34t"
	r.Equal(ValidAddress(firstNotEqualAddress), false)
}

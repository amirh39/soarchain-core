package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MotusWalletKeyPrefix is the prefix to retrieve all MotusWallet
	MotusWalletKeyPrefix = "MotusWallet/value/"
)

// MotusWalletKey returns the store key to retrieve a MotusWallet from the index fields
func MotusWalletKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

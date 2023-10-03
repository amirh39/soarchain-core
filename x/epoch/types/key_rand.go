package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RandomNumberKeyPrefix is the prefix to retrieve all randoms
	RandKeyPrefix = "Rand/value/"
)

// RandKey returns the store key to retrieve a dpr from the index fields
func RandKey(
	eapochNumber string,
) []byte {
	var key []byte

	indexBytes := []byte(eapochNumber)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VrfUserKeyPrefix is the prefix to retrieve all VrfUser
	VrfUserKeyPrefix = "VrfUser/value/"
)

// VrfUserKey returns the store key to retrieve a VrfUser from the index fields
func VrfUserKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

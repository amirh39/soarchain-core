package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidKeyPrefix is the prefix to retrieve all did
	DidKeyPrefix = "Did/value/"
)

// DidKey returns the store key to retrieve a did from the index fields
func DidKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

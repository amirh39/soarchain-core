package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidKeyPrefix is the prefix to retrieve all did
	DidKeyPrefix = "Did/value/"
)

// DidKey returns the store key to retrieve a did from the index fields
func DidKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}

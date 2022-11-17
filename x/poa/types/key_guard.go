package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GuardKeyPrefix is the prefix to retrieve all Guard
	GuardKeyPrefix = "Guard/value/"
)

// GuardKey returns the store key to retrieve a Guard from the index fields
func GuardKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

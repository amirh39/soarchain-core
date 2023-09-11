package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DprKeyPrefix is the prefix to retrieve all dpr
	DprKeyPrefix = "Dpr/value/"
)

// DidKey returns the store key to retrieve a dpr from the index fields
func DprKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

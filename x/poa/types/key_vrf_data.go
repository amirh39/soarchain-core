package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VrfDataKeyPrefix is the prefix to retrieve all VrfData
	VrfDataKeyPrefix = "VrfData/value/"
)

// VrfDataKey returns the store key to retrieve a VrfData from the index fields
func VrfDataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

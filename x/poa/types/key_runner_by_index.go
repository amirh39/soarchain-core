package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RunnerByIndexKeyPrefix is the prefix to retrieve all RunnerByIndex
	RunnerByIndexKeyPrefix = "RunnerByIndex/value/"
)

// RunnerByIndexKey returns the store key to retrieve a RunnerByIndex from the index fields
func RunnerByIndexKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

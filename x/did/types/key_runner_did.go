package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RunnerDidKeyPrefix is the prefix to retrieve all did
	RunnerDidKeyPrefix = "RunnerDid/value/"
)

// DidKey returns the store key to retrieve a did from the index fields
func RunnerDidKey(
	address string,
) []byte {
	var key []byte

	indexBytes := []byte(address)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

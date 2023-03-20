package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RunnerKeyPrefix is the prefix to retrieve all Runner
	RunnerKeyPrefix = "Runner/value/"
)

// RunnerKey returns the store key to retrieve a Runner from the index fields
func RunnerKey(
	PubKey string,
) []byte {
	var key []byte

	indexBytes := []byte(PubKey)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

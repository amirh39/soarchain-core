package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RunnerDidKeyPrefix is the prefix to retrieve all did
	ChallengerDidKeyPrefix = "ChallengerDid/value/"
)

// DidKey returns the store key to retrieve a did from the index fields
func ChallengerDidKey(
	id string,
) []byte {
	var key []byte

	indexBytes := []byte(id)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

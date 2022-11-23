package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChallengerByIndexKeyPrefix is the prefix to retrieve all ChallengerByIndex
	ChallengerByIndexKeyPrefix = "ChallengerByIndex/value/"
)

// ChallengerByIndexKey returns the store key to retrieve a ChallengerByIndex from the index fields
func ChallengerByIndexKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

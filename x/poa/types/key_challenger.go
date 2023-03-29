package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChallengerKeyPrefix is the prefix to retrieve all Challenger
	ChallengerKeyPrefix = "Challenger/value/"
)

// ChallengerKey returns the store key to retrieve a Challenger from the PubKey fields
func ChallengerKey(
	PubKey string,
) []byte {
	var key []byte

	indexBytes := []byte(PubKey)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

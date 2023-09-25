package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	ReputationKeyPrefix = "Reputation/value/"
)

func ReputationKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

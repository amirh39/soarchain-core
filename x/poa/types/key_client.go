package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClientKeyPrefix is the prefix to retrieve all Client
	ClientKeyPrefix = "Client/value/"
)

// ClientKey returns the store key to retrieve a Client from the index fields
func ClientKey(
	PubKey string,
) []byte {
	var key []byte

	indexBytes := []byte(PubKey)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

package utility

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

// TestCreateDID tests the CreateDID function for correct format and hash length.
func TestCreateDID(t *testing.T) {
	address := "soar1z68ga5gq0pks3zla7sg4c9svv8nmggr537m83c"
	expectedPrefix := "did:soar:"

	did, err := CreateDIDId(address)

	// Check if the DID starts with the expected prefix
	if !strings.HasPrefix(did, expectedPrefix) {
		t.Errorf("DID does not start with expected prefix. Got: %s, want prefix: %s", did, expectedPrefix)
	}

	// Extract the hash part from the DID and check its length
	hashPart := strings.TrimPrefix(did, expectedPrefix)
	if len(hashPart) != sha256.Size*2 { // sha256.Size is the number of bytes in the hash, we expect twice that in hex characters
		t.Errorf("Hash part of DID is not the correct length. Got: %d characters, want: %d characters", len(hashPart), sha256.Size*2)
	}

	// is a valid hex string
	test, err := hex.DecodeString(hashPart)
	if err != nil {
		t.Errorf("Hash part of DID is not a valid hex string. Error: %v", err)
	}
	fmt.Printf("t: %v\n", test)
}

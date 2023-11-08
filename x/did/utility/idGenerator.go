package utility

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// CreateDID takes an address string, hashes it with SHA-256, and returns a DID string.
// CreateDIDId takes an address string, hashes it with SHA-256, and returns a DID string along with any error encountered.
func CreateDIDId(address string) (string, error) {
	// Compute SHA-256 hash of the address.
	hasher := sha256.New()
	_, err := hasher.Write([]byte(address))
	if err != nil {
		return "", fmt.Errorf("failed to hash address: %w", err)
	}
	hashed := hasher.Sum(nil)

	// Encode the hash to a hexadecimal string.
	hashedHex := hex.EncodeToString(hashed)

	// Create the DID string with the "did:soar:" prefix.
	did := fmt.Sprintf("did:soar:%s", hashedHex)

	return did, nil
}

package crypto

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func PrivKeyFromBytes(bz []byte) (secp256k1.PrivKey, error) {
	key := make([]byte, secp256k1.PrivKeySize)
	if len(bz) != len(key) {
		return key, fmt.Errorf("invalid Secp256k1 private key. len:%d, expected:%d", len(bz), len(key))
	}
	copy(key[:], bz)
	return key, nil
}

func PubKeyFromBase58(b58 string) (secp256k1.PubKey, error) {
	key := make([]byte, secp256k1.PubKeySize)
	decoded := base58.Decode(b58)
	if len(decoded) != len(key) {
		return key, fmt.Errorf("invalid Secp256k1 public key. len:%d, expected:%d", len(decoded), len(key))
	}
	copy(key[:], decoded)
	return key, nil
}

func DerivePubKey(privKey secp256k1.PrivKey) secp256k1.PubKey {
	return privKey.PubKey().(secp256k1.PubKey)
}

func PubKeyBytes(pubKey secp256k1.PubKey) []byte {
	return pubKey[:]
}

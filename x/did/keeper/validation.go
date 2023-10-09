package keeper

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ValidateInputs(creator string, certificate string, signature string, verificationMethidId string) bool {
	if creator == "" || certificate == "" || signature == "" || verificationMethidId == "" {
		return false
	}
	return true
}

func (k Keeper) IsUniqueDid(ctx sdk.Context, id string) bool {
	_, isFoundClientDid := k.GetClientDid(ctx, id)
	_, isFoundRunnerDid := k.GetRunnerDid(ctx, id)
	_, isFoundChallengerDid := k.GetChallengerDid(ctx, id)
	if isFoundClientDid || isFoundRunnerDid || isFoundChallengerDid {
		return true
	}

	return false
}

func CreateX509CertFromString(certString string) (*x509.Certificate, error) {
	deviceCertPEM := []byte(certString)
	deviceBlock, rest := pem.Decode(deviceCertPEM)
	if deviceBlock == nil && rest != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[CreateX509CertFromString][pem.Decode] failed. Invalid device certificate.")
	}
	deviceCert, err := x509.ParseCertificate(deviceBlock.Bytes)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CreateX509CertFromString][ParseCertificate] failed. Invalid parsing certificate. Error: [ %T ]", err)
	}
	return deviceCert, nil
}

func ValidateX509CertByASN1(creator string, signature string, deviceCert *x509.Certificate) bool {

	if deviceCert == nil || creator == "" {
		return false
	}

	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	hashedAddr := sha256.Sum256([]byte(creator))

	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {

		if ecdsaPubKey, ok := deviceCert.PublicKey.(*ecdsa.PublicKey); ok {

			if ecdsa.VerifyASN1(ecdsaPubKey, hashedAddr[:], signatureBytes) {
				// signature is valid
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func ExtractPubkeyFromCertificate(certificate string) (string, error) {

	if certificate == "" {
		return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][ExtractPubkeyFromCertificate] failed. Device certification is not valid.")
	}

	deviceCertificate, error := CreateX509CertFromString(certificate)
	if error != nil {
		return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][CreateX509CertFromString] failed. Invalid device certificate.")
	}

	pubKeyFromCertificate, err := x509.MarshalPKIXPublicKey(deviceCertificate.PublicKey)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[ExtractPubkeyFromCertificate][MarshalPKIXPublicKey] failed. Couldn't extract a public key from device certificate. Error: [ %T ]", err)
	}

	pubKeyHex := hex.EncodeToString(pubKeyFromCertificate)
	if pubKeyHex == "" {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[ExtractPubkeyFromCertificate][EncodeToString] failed. Couldn't encode public Key to hex string. Error: [ %T ]", err)
	}
	return pubKeyHex, nil
}

func ValidateX509Cert(derivedCert *x509.Certificate, signerCert *x509.Certificate) (bool, error) {
	err := derivedCert.CheckSignatureFrom(signerCert)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func ValidString(input string) bool {
	if len(input) == 0 || input == "" {
		return false
	}
	return true
}

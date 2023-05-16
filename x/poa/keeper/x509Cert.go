package keeper

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ReadX509CertFromFile(fileName string) string {
	cert, _ := ioutil.ReadFile(fileName)
	fileContent := string(cert)
	return fileContent
}

// This function is only for testing purposes
func (k Keeper) CreateX509CertFromFile(fileName string) (*x509.Certificate, error) {

	deviceCertPEM, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	deviceBlock, _ := pem.Decode(deviceCertPEM)
	deviceCert, err := x509.ParseCertificate(deviceBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return deviceCert, nil
}

func (k Keeper) CreateX509CertFromString(certString string) (*x509.Certificate, error) {
	deviceCertPEM := []byte(certString)
	deviceBlock, _ := pem.Decode(deviceCertPEM)
	if deviceBlock == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CreateX509CertFromString][Decode] failed. Can't decode certificate PEM from Not valid string. Make sure you are decoding from a valid and not empty certification string. got: [ %T ]", certString)
	}

	deviceCert, err := x509.ParseCertificate(deviceBlock.Bytes)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[CreateX509CertFromString][ParseCertificate] failed. Can't parse certificate from the given string. Error: [ %T ]", err)
	}
	return deviceCert, nil
}

func (k Keeper) ValidateX509Cert(derivedCert *x509.Certificate, signerCert *x509.Certificate) (bool, error) {
	err := derivedCert.CheckSignatureFrom(signerCert)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func VerifyX509CertByASN1AndExtractPubkey(creatorInput string, signatureInput string, deviceCert *x509.Certificate) (string, error) {

	pubKeyFromCertificate, err := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[CertificateVerificationByASN1][MarshalPKIXPublicKey] failed. Couldn't extract a public key from device certificate. Error: [ %T ]", err)
	}

	pubKeyHex := hex.EncodeToString(pubKeyFromCertificate)
	if pubKeyHex == "" {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CertificateVerificationByASN1][EncodeToString] failed. Couldn't encode public Key to hex string. Error: [ %T ]", err)
	}

	signature, err := hex.DecodeString(signatureInput)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[CertificateVerificationByASN1][DecodeString] failed. Couldn't decode the signature. got: [ %T ]. Error: [ %T ]", signature, err)
	}

	hashedAddr := sha256.Sum256([]byte(creatorInput))

	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {

		if ecdsaPubKey, ok := deviceCert.PublicKey.(*ecdsa.PublicKey); ok {

			if ecdsa.VerifyASN1(ecdsaPubKey, hashedAddr[:], signature) {
				// signature is valid
			} else {
				return "", sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "[CertificateVerificationByASN1][VerifyASN1] failed. Certificate verification failed for extracted public Key: [ %T ] and Hash Address: [ %T ] and Signature: [ %T ]. Error: [ %T ]", ecdsaPubKey, hashedAddr, signature, err)
			}
		} else {
			return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[CertificateVerificationByASN1] failed. Invalid public key type. Error: [ %T ]", err)
		}
	}
	return pubKeyHex, nil
}

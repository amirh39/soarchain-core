package keeper

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ReadX509CertFromFile(fileName string) string {
	cert, _ := ioutil.ReadFile(fileName)
	fileContent := string(cert)
	return fileContent
}

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

func (k Keeper) ValidateX509Cert(derivedCert *x509.Certificate, signerCert *x509.Certificate) (bool, error) {
	err := derivedCert.CheckSignatureFrom(signerCert)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (k Keeper) VerifyX509CertByASN1AndExtractPubkey(creatorInput string, signatureInput string, deviceCert *x509.Certificate) (string, error) {

	if deviceCert == nil || creatorInput == "" {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[VerifyX509CertByASN1AndExtractPubkey] failed. Couldn't find valid creatorInput OR signatureInput. got: creatorInput [ %T ] signatureInput [ %T ]. Make sure you they are valid and not empty.", creatorInput, signatureInput)
	}

	pubKeyFromCertificate, err := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[VerifyX509CertByASN1AndExtractPubkey][MarshalPKIXPublicKey] failed. Couldn't extract a public key from device certificate. Error: [ %T ]", err)
	}

	pubKeyHex := hex.EncodeToString(pubKeyFromCertificate)
	if pubKeyHex == "" {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[VerifyX509CertByASN1AndExtractPubkey][EncodeToString] failed. Couldn't encode public Key to hex string. Error: [ %T ]", err)
	}

	signature, err := hex.DecodeString(signatureInput)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[VerifyX509CertByASN1AndExtractPubkey][DecodeString] failed. Couldn't decode the signature. got: [ %T ]. Error: [ %T ]", signature, err)
	}

	hashedAddr := sha256.Sum256([]byte(creatorInput))

	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {

		if ecdsaPubKey, ok := deviceCert.PublicKey.(*ecdsa.PublicKey); ok {

			if ecdsa.VerifyASN1(ecdsaPubKey, hashedAddr[:], signature) {
				// signature is valid
			} else {
				return "", sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "[VerifyX509CertByASN1AndExtractPubkey][VerifyASN1] failed. Certificate verification failed for extracted public Key: [ %T ] and Hash Address: [ %T ] and Signature: [ %T ]. Error: [ %T ]", ecdsaPubKey, hashedAddr, signature, err)
			}
		} else {
			return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[VerifyX509CertByASN1AndExtractPubkey] failed. Invalid public key type. Error: [ %T ]", err)
		}
	}
	return pubKeyHex, nil
}

func (k msgServer) validateCertificate(ctx sdk.Context, deviceCert *x509.Certificate) error {

	totalKeys := k.GetAllFactoryKeys(ctx)

	for i := uint64(0); i < uint64(len(totalKeys)); i++ {
		factoryKey, isFound := k.GetFactoryKeys(ctx, i)
		if isFound {
			factoryCert, err := k.CreateX509CertFromString(factoryKey.FactoryCert)
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[validateCertificate][CreateX509CertFromString] failed. Factory certificate couldn't be created from the storage."+err.Error())
			}

			validated, err := k.ValidateX509Cert(deviceCert, factoryCert)
			if err != nil {
				continue // Try next certificate
			}

			if validated {
				return nil
			}
		}
	}

	return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[validateCertificate][ValidateX509Cert] failed. Device certificate couldn't be verified.")
}

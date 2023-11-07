package keeper

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: After did structure changed, this need to refactor
func (k Keeper) ClientDidValidateInputs(msg *types.MsgGenClient) bool {

	isValidDid := types.ValidateDid(msg.Document.Id)
	if !isValidDid {
		return false
	}

	isValidDidAddress := types.ValidateDidAddress(msg.Document.Address)
	if !isValidDidAddress {
		return false
	}

	isValidateSupportedPIDs := types.ValidateSupportedPIDs(msg.Document.SupportedPIDs)
	if !isValidateSupportedPIDs {
		return false
	}

	if msg.Document == nil || msg.Document.VerificationMethods == nil || len(msg.Document.VerificationMethods) < 1 || msg.Document.VerificationMethods[0].Id == "" {
		return false
	}

	if msg.Document.SupportedPIDs == "" || msg.Document.Address == "" || msg.Document.Id == "" || msg.Creator == "" || msg.Certificate == "" || msg.Signature == "" {
		return false
	}
	return true
}

// TODO: After did structure changed, this need to refactor
func (k Keeper) RunnerDidValidateInputs(msg *types.MsgGenRunner) bool {

	isValidDid := types.ValidateDid(msg.Document.Id)
	if !isValidDid {
		return false
	}

	isValidDidAddress := types.ValidateDidAddress(msg.Document.Address)
	if !isValidDidAddress {
		return false
	}

	isValidStakeAmount := types.ValidateStakeAmount(msg.RunnerStake)
	if !isValidStakeAmount {
		return false
	}

	if msg.Document == nil || msg.Document.VerificationMethods == nil || len(msg.Document.VerificationMethods) < 1 || msg.Document.VerificationMethods[0].Id == "" {
		return false
	}

	if msg.Document.Address == "" || msg.Document.Id == "" || msg.Creator == "" || msg.Certificate == "" || msg.Signature == "" || msg.RunnerStake == "" {
		return false
	}

	return true
}

// TODO: After did structure changed, this need to refactor
func (k Keeper) ChallengerDidValidateInputs(msg *types.MsgGenChallenger) bool {

	isValidDid := types.ValidateDid(msg.Document.Id)
	if !isValidDid {
		return false
	}

	isValidDidAddress := types.ValidateDidAddress(msg.Document.Address)
	if !isValidDidAddress {
		return false
	}

	isValidStakeAmount := types.ValidateStakeAmount(msg.ChallengerStake)
	if !isValidStakeAmount {
		return false
	}

	if msg.Document == nil || msg.Document.VerificationMethods == nil || len(msg.Document.VerificationMethods) < 1 || msg.Document.VerificationMethods[0].Id == "" {
		return false
	}

	if msg.Document.Address == "" || msg.Document.Id == "" || msg.Creator == "" || msg.Certificate == "" || msg.Signature == "" || msg.ChallengerStake == "" || msg.ChallengerType == "" {
		return false
	}

	return true
}

func (k Keeper) IsNotUniqueDid(ctx sdk.Context, id string) bool {
	_, isFoundClientDid := k.GetClientDidId(ctx, id)
	if isFoundClientDid {
		return true
	}
	_, isFoundRunnerDid := k.GetRunnerDidId(ctx, id)
	if isFoundRunnerDid {
		return true
	}
	_, isFoundChallengerDid := k.GetChallengerDidId(ctx, id)
	if isFoundChallengerDid {
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

func ValidateDeactivatingDidInputs(fromAddress string, didType string) bool {

	if fromAddress == "" || didType == "" {
		return false
	}
	return true
}

func (k Keeper) ClientType(deviceCert *x509.Certificate) string {
	if len(deviceCert.Issuer.Names) < 1 || deviceCert.Issuer.Names[1].Value == nil {
		return "[GenClient][ClientType] failed. No Type for device certificate."
	}
	results := fmt.Sprintf("%v", deviceCert.Issuer.Names[1].Value)
	if results[41:43] == "01" {
		return "mini"
	} else {
		return "pro"
	}
}

func (k Keeper) GeneratePubkey(msg *types.MsgGenClient) (pubkey string, deviceCert *x509.Certificate, err error) {
	deviceCert, error := CreateX509CertFromString(msg.Certificate)
	if error != nil {
		return "", nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][CreateX509CertFromString] failed. Invalid device certificate.")
	}

	isValide := ValidateX509CertByASN1(msg.Creator, msg.Signature, deviceCert)
	if !isValide {
		return "", nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][ValidateX509CertByASN1] failed. Invalid device certificate and signature.")
	}

	pubKeyHex, error := ExtractPubkeyFromCertificate(msg.Certificate)
	if pubKeyHex == "" || error != nil {
		return "", nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][ExtractPubkeyFromX509Cert] failed. Invalid certificate validation.")
	}
	return pubKeyHex, deviceCert, nil
}

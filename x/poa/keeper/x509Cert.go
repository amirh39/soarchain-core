package keeper

import (
	"crypto/x509"
	"encoding/pem"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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

func (k Keeper) ValidateCertificate(ctx sdk.Context, deviceCert *x509.Certificate) error {

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

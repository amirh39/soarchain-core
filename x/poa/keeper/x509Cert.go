package keeper

import (
	"crypto/x509"
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

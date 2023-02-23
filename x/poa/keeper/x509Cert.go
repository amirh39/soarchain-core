package keeper

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

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
	deviceCert, err := x509.ParseCertificate(deviceBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return deviceCert, nil
}

func (k Keeper) CreateX509PubkeyFromString(certString string) (*rsa.PublicKey, error) {
	deviceCertPEM := []byte(certString)
	deviceBlock, _ := pem.Decode(deviceCertPEM)
	devicePubkey, err := x509.ParsePKCS1PublicKey(deviceBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return devicePubkey, nil
}

func (k Keeper) CreateX509PubkeyFromFile(fileName string) (*rsa.PublicKey, error) {

	deviceCertPEM, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	deviceBlock, _ := pem.Decode(deviceCertPEM)
	devicePubkey, err := x509.ParsePKCS1PublicKey(deviceBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return devicePubkey, nil
}

func (k Keeper) ValidateX509Cert(derivedCert *x509.Certificate, signerCert *x509.Certificate) (bool, error) {
	err := derivedCert.CheckSignatureFrom(signerCert)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

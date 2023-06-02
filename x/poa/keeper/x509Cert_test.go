package keeper_test

import (
	"crypto/x509"
	"encoding/pem"
	keepertest "soarchain/testutil/keeper"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CreateX509CertFromString(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	certString := Valid_CertString
	certBuffer, _ := pem.Decode([]byte(certString))

	expectedCert, err := x509.ParseCertificate(certBuffer.Bytes)
	if err != nil {
		t.Fatalf("Error parsing certificate: %v", err)
	}

	actualCert, err := keeper.CreateX509CertFromString(certString)
	if err != nil {
		t.Fatalf("Error creating certificate from string: %v", err)
	}

	if !expectedCert.Equal(actualCert) {
		t.Errorf("Expected certificate to be %v but got %v", expectedCert, actualCert)
	}
}

/** This function created to make sure error handling messages will appear while using not valid string*/
func Test_CreateX509CertFromNotValidString(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	certString := INValid_CertString

	actualCert, err := keeper.CreateX509CertFromString(certString)

	require.Error(t, err)
	require.NotEqual(t, certString, actualCert)
}

func Test_VerifyX509CertByASN1AndExtractPubkey(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	certString := CERTIFICATE
	signature := Signature
	creator := CertCreator

	deviceCert, err := keeper.CreateX509CertFromString(certString)

	require.NoError(t, err)
	require.NotNil(t, deviceCert)

	pubKeyHex, err := keeper.VerifyX509CertByASN1AndExtractPubkey(creator, signature, deviceCert)

	require.NoError(t, err)
	require.NotNil(t, pubKeyHex)
}

func Test_NotVerifiedX509CertByASN1AndExtractPubkey(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	signature := ""
	creator := ""

	pubKeyHex, err := keeper.VerifyX509CertByASN1AndExtractPubkey(creator, signature, nil)

	require.Error(t, err)
	require.Empty(t, pubKeyHex)
}

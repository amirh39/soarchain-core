package keeper_test

import (
	"crypto/x509"
	"encoding/pem"
	"testing"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
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

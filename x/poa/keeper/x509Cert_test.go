package keeper_test

import (
	"crypto/x509"
	"encoding/pem"
	keepertest "soarchain/testutil/keeper"
	"testing"
)

func TestCreateX509CertFromFile(t *testing.T) {
	tempCertData := []byte(`-----BEGIN CERTIFICATE-----
MIIC+zCCAeOgAwIBAgIJAJOljU6QFZ6gMAoGCCqGSM49BAMCMIGQMQswCQYDVQQG
EwJVUzEVMBMGA1UECAwMV2VzdGVybiBDYXBlMRMwEQYDVQQHDApBbnRvcm5lc3Mx
EzARBgNVBAoMCk9wZW5BSSBJbmMxEjAQBgNVBAMMCWxvY2FsaG9zdDAeFw0yMjAz
MjExMzA3MjZaFw0yMjA0MzExMzA3MjZaMIGQMQswCQYDVQQGEwJVUzEVMBMGA1UE
CAwMV2VzdGVybiBDYXBlMRMwEQYDVQQHDApBbnRvcm5lc3MxEzARBgNVBAoMCk9w
ZW5BSSBJbmMxEjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAJAL80Z6jpv29oUe/z6Q2c/TUzFk+MQaARssvkGhKQ8IOJWz
mZ/Cf46cFrKvG8JqqW4Ht4sru4NUJELvIW54XtuWnDZQYBZju3TqJhRHjep4yP4I
onGxJxhXKzgSIZ2S2L1tq3TsygEzjy59VLfobvpCt/tQ2vXO1JbO1LZ17xx46Kjw
uxpW/IixK5U6J5U6HXxU6rgWUtiYKZrxAEJFffLQwz8B38Kcp36C6NGB+kq3qN9c
aYVZWEKb/Ji7VwD+g95R5UO7TrU6GxvvyoOVTp9mmZM/TcF6oQ2zLof1d8y2l0Ws
I+q3q3iPdw9SS0YrBc/pzLL1fgV/hLl0n4f4cNECAwEAAaMhMB8wHQYDVR0OBBYE
FFbqWt/8i+DgCmBeb/XQINmWqlX8MBIGA1UdEwEB/wQIMAYBAf8CAQAwDQYJKoZI
hvcNAQEMBQADggEBA
-----END CERTIFICATE-----`)
	keeper, _ := keepertest.PoaKeeper(t)
	cert, err := keeper.CreateX509CertFromFile("signer_FFFF.der")
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the certificate was properly parsed and returned
	decodedBlock, _ := pem.Decode(tempCertData)
	expectedCert, err := x509.ParseCertificate(decodedBlock.Bytes)
	if err != nil {
		t.Fatal(err)
	}
	if !cert.Equal(expectedCert) {
		t.Errorf("Returned certificate doesn't match expected certificate")
	}

}

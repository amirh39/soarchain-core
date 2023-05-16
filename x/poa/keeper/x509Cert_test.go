package keeper_test

import (
	"crypto/x509"
	"encoding/pem"
	keepertest "soarchain/testutil/keeper"
	"testing"
)

func TestCreateX509CertFromString(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	certString := "-----BEGIN CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"
	certBuffer, _ := pem.Decode([]byte(certString))

	expectedCert, err := x509.ParseCertificate(certBuffer.Bytes)
	if err != nil {
		t.Fatalf("Error parsing certificate: %v", err)
	}

	actualCert, err := keeper.CreateX509CertFromString(certString)
	if err != nil {
		t.Fatalf("Error creating certificate from string: %v", err)
	}

	// t.Log("ExpectedCert:", expectedCert)
	// t.Log("ActualCert:", actualCert)

	if !expectedCert.Equal(actualCert) {
		t.Errorf("Expected certificate to be %v but got %v", expectedCert, actualCert)
	}
}

/** This function created to make sure error handling messages will appear while using not valid string*/
func TestCreateX509CertFromNotValidString(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	certString := "-----CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"

	actualCert, err := keeper.CreateX509CertFromString(certString)
	if err != nil {
		t.Fatalf("Error creating certificate from string: %v", err)
	}

	t.Log("ActualCert:", actualCert)

}

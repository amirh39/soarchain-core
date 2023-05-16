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
func Test_CreateX509CertFromNotValidString(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	certString := "-----CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"

	actualCert, err := keeper.CreateX509CertFromString(certString)
	if err != nil {
		t.Fatalf("Error creating certificate from string: %v", err)
	}

	t.Log("ActualCert:", actualCert)
}

func Test_VerifyX509CertByASN1AndExtractPubkey(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	certString := "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
	signature := "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0"
	creator := "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"

	deviceCert, err := keeper.CreateX509CertFromString(certString)

	pubKeyHex, err := keeper.VerifyX509CertByASN1AndExtractPubkey(creator, signature, deviceCert)

	t.Log("pubKeyHex:", pubKeyHex)

	require.NoError(t, err)
	require.NotNil(t, pubKeyHex)
}

func Test_NotVerifiedX509CertByASN1AndExtractPubkey(t *testing.T) {
	keeper, _ := keepertest.PoaKeeper(t)

	signature := ""
	creator := ""

	pubKeyHex, err := keeper.VerifyX509CertByASN1AndExtractPubkey(creator, signature, nil)

	t.Log("pubKeyHex:", pubKeyHex)

	require.Error(t, err)
	require.Empty(t, pubKeyHex)
}

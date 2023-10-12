/** This file is created for tests. Firstly search what you nee if not find then create a new one for you. */
package keeper_test

import (
	"context"
	"soarchain/x/did/keeper"
	"soarchain/x/did/types"
	"soarchain/x/dpr/testutil"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	keepertest "soarchain/testutil/keeper"

	"soarchain/x/did/utility/crypto"

	tendermintcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func NewDIDDocumentWithSeq(did string) (types.ClientDidWithSeq, tendermintcrypto.PrivKey) {
	privKey := secp256k1.GenPrivKey()
	pubKey := crypto.PubKeyBytes(crypto.DerivePubKey(privKey))
	verificationMethodID := types.NewVerificationMethodID(did, "key1")
	es256VerificationMethod := types.NewVerificationMethod(verificationMethodID, types.ES256K_2019, did, pubKey)
	blsVerificationMethod := types.NewVerificationMethod(verificationMethodID, types.BLS1281G2_2020, did, []byte("dummy BBS+ pub key"))
	verificationMethods := []*types.VerificationMethod{
		&es256VerificationMethod,
		&blsVerificationMethod,
	}
	verificationRelationship := types.NewVerificationRelationship(verificationMethods[0].Id)
	authentications := []types.VerificationRelationship{
		verificationRelationship,
	}
	soarchainPublicKey := types.NewKeys(did, PUBKEYTYPE, CONTROLLER, PUBLICKEYPEM)
	vehicle := types.NewVehicle(VIN)
	owner := types.NewOwner(OWNERID, PURCHESDATE)
	doc := types.NewClientDidDocument(did, INDEX, ADDRESS, TYPE, PIDS, types.WithVerificationMethods(verificationMethods), types.WithAuthentications(authentications), types.WithKeys(&soarchainPublicKey), types.WithVehicle(&vehicle), types.WithOwner(&owner))
	docWithSeq := types.NewDidDocumentWithSeq(
		&doc,
		types.InitialSequence,
	)
	return docWithSeq, privKey
}

func NewRunnerDidDocumentWithSeq(did string) (types.RunnerDidWithSeq, tendermintcrypto.PrivKey) {
	privKey := secp256k1.GenPrivKey()
	pubKey := crypto.PubKeyBytes(crypto.DerivePubKey(privKey))
	verificationMethodID := types.NewVerificationMethodID(did, "key1")
	es256VerificationMethod := types.NewVerificationMethod(verificationMethodID, types.ES256K_2019, did, pubKey)
	blsVerificationMethod := types.NewVerificationMethod(verificationMethodID, types.BLS1281G2_2020, did, []byte("dummy BBS+ pub key"))
	verificationMethods := []*types.VerificationMethod{
		&es256VerificationMethod,
		&blsVerificationMethod,
	}
	verificationRelationship := types.NewVerificationRelationship(verificationMethods[0].Id)
	authentications := []types.VerificationRelationship{
		verificationRelationship,
	}
	soarchainPublicKey := types.NewKeys(did, PUBKEYTYPE, CONTROLLER, PUBLICKEYPEM)
	doc := types.NewRunnerDidDocument(did, INDEX, ADDRESS, types.WithRunnerVerificationMethods(verificationMethods), types.WithRunnerAuthentications(authentications), types.WithRunnerKeys(&soarchainPublicKey))
	docWithSeq := types.NewRunnerDidDocumentWithSeq(
		&doc,
		types.InitialSequence,
	)
	return docWithSeq, privKey
}

func NewChallengerDidDocumentWithSeq(did string) (types.ChallengerDidWithSeq, tendermintcrypto.PrivKey) {
	privKey := secp256k1.GenPrivKey()
	pubKey := crypto.PubKeyBytes(crypto.DerivePubKey(privKey))
	verificationMethodID := types.NewVerificationMethodID(did, "key1")
	es256VerificationMethod := types.NewVerificationMethod(verificationMethodID, types.ES256K_2019, did, pubKey)
	blsVerificationMethod := types.NewVerificationMethod(verificationMethodID, types.BLS1281G2_2020, did, []byte("dummy BBS+ pub key"))
	verificationMethods := []*types.VerificationMethod{
		&es256VerificationMethod,
		&blsVerificationMethod,
	}
	verificationRelationship := types.NewVerificationRelationship(verificationMethods[0].Id)
	authentications := []types.VerificationRelationship{
		verificationRelationship,
	}
	soarchainPublicKey := types.NewKeys(did, PUBKEYTYPE, CONTROLLER, PUBLICKEYPEM)
	doc := types.NewChallengerDidDocument(did, INDEX, ADDRESS, types.WithChallengerVerificationMethods(verificationMethods), types.WithChallengerAuthentications(authentications), types.WithChallengerKeys(&soarchainPublicKey))
	docWithSeq := types.NewChallengerDidDocumentWithSeq(
		&doc,
		types.InitialSequence,
	)
	return docWithSeq, privKey
}

func NewMsgDeactivateDID(doc types.ClientDid, did string, verificationMethodID string, privKey tendermintcrypto.PrivKey, seq uint64) types.MsgDeactivateDid {
	sig, _ := types.Sign(&doc, seq, privKey)
	return *types.NewMsgDeactivateDid(did, verificationMethodID, sig, sdk.AccAddress{}.String())
}

func MakeTestData() (string, types.ClientDidWithSeq, tendermintcrypto.PrivKey, string) {
	doc, privKey := NewDIDDocumentWithSeq(Did)
	return Did, doc, privKey, doc.Document.VerificationMethods[0].Id
}

func SetupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.DidKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

func CreateNChallengerDid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChallengerDid {
	items := make([]types.ChallengerDid, n)
	for i := range items {
		items[i].PubKey = Challenger_PubKey
		items[i].Address = Challenger_Address
		items[i].IpAddress = Challenger_IPAddress
		items[i].StakedAmount = Challenger_StakedAmount

		keeper.SetChallengerDid(ctx, items[i])
	}
	return items
}

const (
	Challenger_PubKey        = "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251"
	Challenger_Address       = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score         = "189"
	Challenger_StakedAmount  = "2000000000utmotus"
	Challenger_NetEarnings   = "0utmotus"
	Challenger_StakedAmount2 = "2000000000udmotus"
	Challenger_NetEarnings2  = "0udmotus"
	Challenger_IpAddr        = ""
	Challenger_IPAddress     = "104.248.142.45"
	Challenger_Type          = "v2n"
	Challenger_Creator       = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score2        = "82"
)

const (
	Did       = "did:soar:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm"
	SecondDid = "did:soar:1Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgap"
)

const (
	ADDRESS = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	PUBKEY  = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	VIN     = "1HGCM82636c678d14c93ad5bf14448da57f4f241b77e30a013d54f5d76c8126a7029aeb86"
)

var PIDS = []bool{true, false, false}

const (
	INDEX        = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	TYPE         = "mini"
	PUBKEYTYPE   = "Pubkey-Type"
	CONTROLLER   = "Controller"
	PUBLICKEYPEM = "-----CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"
	OWNERID      = "Owner-Id"
	PURCHESDATE  = "Purches-Date"
)

const (
	Signature   = "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0"
	Certificate = "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
)

const (
	MASTER_CERTIFICATE = "-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"
	MASTER_ACCOUNT     = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
)

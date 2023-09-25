/** This file is created for tests. Firstly search what you nee if not find then create a new one for you. */
package keeper_test

import (
	"context"
	"soarchain/x/did/keeper"
	"soarchain/x/did/types"
	"soarchain/x/dpr/testutil"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	keepertest "soarchain/testutil/keeper"

	"soarchain/x/did/utility/crypto"

	tendermintcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func NewMsgUpdateDID(newDoc types.DidDocument, verificationMethodID string, privKey tendermintcrypto.PrivKey, seq uint64) types.MsgUpdateDid {
	sig, _ := types.Sign(&newDoc, seq, privKey)
	return *types.NewMsgUpdateDid(newDoc.Id, newDoc, verificationMethodID, sig, ADDRESS) // sdk.AccAddress{}.String()
}

func NewDIDDocumentWithSeq(did string) (types.DidDocumentWithSeq, tendermintcrypto.PrivKey) {
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
	soarchainPublicKey := types.NewSoarchainPublicKey(did, PUBKEYTYPE, CONTROLLER, PUBLICKEYPEM)
	vehicle := types.NewVehicle(VIN)
	owner := types.NewOwner(OWNERID, PURCHESDATE)
	doc := types.NewDidDocument(did, INDEX, ADDRESS, TYPE, PIDS, types.WithVerificationMethods(verificationMethods), types.WithAuthentications(authentications), types.WithSoarchainPublicKey(&soarchainPublicKey), types.WithVehicle(&vehicle), types.WithOwner(&owner))
	docWithSeq := types.NewDidDocumentWithSeq(
		&doc,
		types.InitialSequence,
	)
	return docWithSeq, privKey
}

func NewMsgCreateDID(doc types.DidDocument, verificationMethodID string, privKey tendermintcrypto.PrivKey) types.MsgGenDid {
	sig, _ := types.Sign(&doc, types.InitialSequence, privKey)
	return *types.NewMsgGenDid(doc.Id, doc, verificationMethodID, sig, sdk.AccAddress{}.String())
}

func NewMsgDeactivateDID(doc types.DidDocument, did string, verificationMethodID string, privKey tendermintcrypto.PrivKey, seq uint64) types.MsgDeactivateDid {
	sig, _ := types.Sign(&doc, seq, privKey)
	return *types.NewMsgDeactivateDid(did, verificationMethodID, sig, sdk.AccAddress{}.String())
}

func MakeTestData() (string, types.DidDocumentWithSeq, tendermintcrypto.PrivKey, string) {
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

const (
	Did                  = "did:soar:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm"
	SecondDid            = "did:soar:1Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgap"
	VerificationMethodId = Did + "#key1"
	ClientCertificate    = `-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----`
)

const (
	ADDRESS = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	PUBKEY  = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	VIN     = "1HGCM82636c678d14c93ad5bf14448da57f4f241b77e30a013d54f5d76c8126a7029aeb86"
)

var PIDS = []bool{true, false, false}

const (
	INDEX           = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	TYPE            = "mini"
	PUBKEYTYPE      = "Pubkey-Type"
	CONTROLLER      = "Controller"
	PUBLICKEYPEM    = "-----CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"
	OWNERID         = "Owner-Id"
	PURCHESDATE     = "Purches-Date"
	ClientSignature = "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0"
	Certificate     = "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
)

func CreateNClient(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Reputation {
	items := make([]types.Reputation, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(i)

		keeper.SetReputation(ctx, items[i])
	}
	return items
}

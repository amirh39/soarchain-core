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
	doc := types.NewDidDocument(did, PUBKEY, VIN, PIDS, types.WithVerificationMethods(verificationMethods), types.WithAuthentications(authentications))
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
)

const (
	ADDRESS = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	PUBKEY  = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	VIN     = "1HGCM82636c678d14c93ad5bf14448da57f4f241b77e30a013d54f5d76c8126a7029aeb86"
)

var PIDS = []bool{true, false, false}

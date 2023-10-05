package keeper_test

import (
	"soarchain/x/did/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"soarchain/x/did/utility/crypto"
)

func TestHandleMsgUpdateDID(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServer(t)
	defer ctrl.Finish()
	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)

	did, origDocWithSeq, privKey, verificationMethodID := MakeTestData()
	k.SetClientDidDocument(ctx, did, origDocWithSeq)

	newDoc := origDocWithSeq.Document
	verificationMethod := types.NewVerificationMethod(
		types.NewVerificationMethodID(did, "key2"),
		types.ES256K_2019,
		did,
		crypto.PubKeyBytes(crypto.DerivePubKey(secp256k1.GenPrivKey())),
	)
	newDoc.VerificationMethods = append(newDoc.VerificationMethods, &verificationMethod)

	// call
	updateMsg := NewMsgUpdateDID(*newDoc, verificationMethodID, privKey, origDocWithSeq.Sequence)
	updateRes, err := msgServer.UpdateDid(context, &updateMsg)
	require.Nil(t, err)
	require.NotNil(t, updateRes)

	updatedDoc, found := k.GetClientDidDocument(ctx, did)
	require.Equal(t, true, found)
	require.NotNil(t, updatedDoc)
}

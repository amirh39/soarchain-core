package keeper_test

import (
	"soarchain/x/did/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestHandleMsgDeactivateDID(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServer(t)
	defer ctrl.Finish()
	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)
	//
	did, docWithSeq, privKey, verificationMethodID := MakeTestData()
	k.SetDidDocument(ctx, did, docWithSeq)

	didDocument, found := k.GetDidDocument(ctx, did)
	require.Equal(t, true, found)
	require.NotNil(t, didDocument)

	// deactivate
	deactivateMsg := NewMsgDeactivateDID(*didDocument.Document, did, verificationMethodID, privKey, types.InitialSequence)
	deactivateRes, err := msgServer.DeactivateDid(context, &deactivateMsg)

	require.NoError(t, err)
	require.NotNil(t, deactivateRes)

	// check if it's really deactivated
	got, found := k.GetDidDocument(ctx, did)
	require.False(t, got.Empty())
	require.True(t, found)
	require.True(t, got.Deactivated())
	require.Equal(t, types.InitialSequence+1, got.Sequence)
}

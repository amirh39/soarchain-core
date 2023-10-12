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
	did, docWithSeq, _, _ := MakeTestData()
	k.SetClientDid(ctx, *docWithSeq.Document)

	didDocument, found := k.GetClientDid(ctx, ADDRESS)
	require.Equal(t, true, found)
	require.NotNil(t, didDocument)

	// deactivate
	deactivateMsg := types.MsgDeactivateDid{
		Did:         Did,
		FromAddress: ADDRESS,
	}
	clientDid, err := msgServer.DeactivateDid(context, &deactivateMsg)
	if err != nil {
		require.NotNil(t, err)
		require.Nil(t, clientDid)
	} else {
		require.Nil(t, err)
		got, found := k.GetClientDid(ctx, did)
		require.NotNil(t, got)
		require.False(t, found)
	}
}

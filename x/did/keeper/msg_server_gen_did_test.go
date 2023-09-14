package keeper_test

import (
	"soarchain/x/did/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_GenDid(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServer(t)
	defer ctrl.Finish()
	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)
	documentWithSequence, privKey := NewDIDDocumentWithSeq(Did)
	doc := documentWithSequence.Document
	sig, error := types.Sign(doc, documentWithSequence.Sequence, privKey)
	require.NoError(t, error)
	res, err := msgServer.GenDid(context, &types.MsgGenDid{
		Did:                  Did,
		Document:             documentWithSequence.Document,
		VerificationMethodId: VerificationMethodId,
		Signature:            sig,
		FromAddress:          ADDRESS,
	})
	didDocument, _ := k.GetDidDocumentWithSequence(ctx, Did)
	t.Log("didDocument------------------->", didDocument)
	require.NotNil(t, res)
	require.NoError(t, err)
}

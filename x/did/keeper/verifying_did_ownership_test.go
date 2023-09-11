package keeper_test

import (
	"soarchain/x/did/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_VerifyOwnership(t *testing.T) {
	_, k, context, ctrl, bank := SetupMsgServer(t)
	defer ctrl.Finish()
	bank.ExpectAny(context)

	docWithSeq, privKey := NewDIDDocumentWithSeq(Did)
	doc := docWithSeq.Document

	sig, err := types.Sign(doc, docWithSeq.Sequence, privKey)
	require.Nil(t, err)

	newSeq, err := k.VerifyDidOwnership(doc, docWithSeq.Sequence, docWithSeq.Document, docWithSeq.Document.VerificationMethods[0].Id, sig)
	require.NoError(t, err)
	require.Equal(t, docWithSeq.Sequence+1, newSeq)
}

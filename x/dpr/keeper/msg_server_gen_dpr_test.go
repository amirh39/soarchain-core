package keeper_test

import (
	"soarchain/x/dpr/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_GenDpr(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServer(t)
	defer ctrl.Finish()
	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)

	res, err := msgServer.GenDpr(context, &types.MsgGenDpr{
		Creator:                       CREATOR,
		PidSupportedOneToTwnety:       true,
		PidSupportedTwentyOneToForthy: false,
		PidSupportedForthyOneToSixty:  false,
		LengthOfDpr:                   129,
	})
	require.Nil(t, err)
	require.NotNil(t, res)

	dprs, found := k.GetAllDpr(ctx)
	require.Equal(t, found, true)
	require.NotNil(t, dprs)
}

// func (helper *KeeperTestHelper) Test_Gen_DPR() {

// 	helper.Run("TestGenDpr", func() {
// 		helper.Setup()
// 		keeper := helper.App.DprKeeper
// 		didKeeper := helper.App.DidKeeper

// 		ctx := sdk.WrapSDKContext(helper.Ctx)

// 		didDocument, privkey := NewDIDDocumentWithSeq(Did)
// 		helper.Require().NotEmpty(didDocument)
// 		helper.Require().NotEmpty(privkey)
// 		didKeeper.SetDidDocument(helper.Ctx, Did, didDocument)
// 		did, err := didKeeper.GetDidDocumentWithSequence(helper.Ctx, Did)
// 		helper.Require().NotEmpty(did)
// 		helper.Require().Empty(err)

// 		// var msgServer types.MsgServer

// 		res, err := keeper.SetDpr(helper.Ctx, types.MsgGenDpr{
// 			Creator:                       CREATOR,
// 			PidSupportedOneToTwnety:       true,
// 			PidSupportedTwentyOneToForthy: false,
// 			PidSupportedForthyOneToSixty:  false,
// 			LengthOfDpr:                   12,
// 		})
// 		helper.Require().NotEmpty(res)
// 		helper.Require().Empty(err)
// 	})
// }

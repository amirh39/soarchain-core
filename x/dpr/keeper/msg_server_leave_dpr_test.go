package keeper_test

import (
	"soarchain/x/dpr/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_LeaveDpr(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServer(t)
	defer ctrl.Finish()
	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)

	res, err := msgServer.LeaveDpr(context, &types.MsgLeaveDpr{
		PubKey: "",
		Sender: "",
		DprId:  "123",
	})
	require.Nil(t, err)
	require.NotNil(t, res)

	dprs, _ := k.GetAllDpr(ctx)
	require.NotNil(t, dprs)
}

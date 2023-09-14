package keeper_test

import (
	"soarchain/x/dpr/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_ActivateDpr(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServer(t)
	defer ctrl.Finish()
	bank.ExpectAny(context)
	ctx := sdk.UnwrapSDKContext(context)

	res, err := msgServer.ActivateDpr(context, &types.MsgActivateDpr{
		Sender: CREATOR,
		DprId:  "123",
	})
	require.Nil(t, err)
	require.NotNil(t, res)

	dprs, _ := k.GetAllDpr(ctx)
	require.NotNil(t, dprs)
}

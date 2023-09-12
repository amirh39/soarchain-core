package keeper_test

import (
	"soarchain/x/dpr/types"
	"strconv"
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
		Creator:              CREATOR,
		PidSupported_1To_20:  true,
		PidSupported_21To_40: false,
		PidSupported_41To_60: false,
		LengthOfDpr:          129,
		Vin:                  []string{strconv.Itoa(0), strconv.Itoa(0)},
	})
	require.Nil(t, err)
	require.NotNil(t, res)

	dprs := k.GetAllDpr(ctx)
	require.NotNil(t, dprs)
}

package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_SelectRandomChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	challenger := SetupNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	res, err := msgServer.SelectRandomChallenger(context, &types.MsgSelectRandomChallenger{
		Creator: Challenger_Creator,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid runner, response should raise proper error message*/
func Test_SelectRandomChallenger_NotValidCreator(t *testing.T) {
	msgServer, k, context, ctrl, bank := SetupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	challenger := SetupNChallenger(1)
	k.SetChallenger(ctx, challenger[0])

	res, err := msgServer.SelectRandomChallenger(context, &types.MsgSelectRandomChallenger{
		Creator: "",
	})

	require.Error(t, err)
	require.Nil(t, res)
}

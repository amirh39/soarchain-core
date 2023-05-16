package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	Amount            = "0soar"
	InsufficientFound = "10soar"
)

func Test_ClaimRunnerRewards(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Runner{
		PubKey:             "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795",
		Address:            "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8",
		Score:              "70.01360618066334",
		RewardMultiplier:   "4901.905050421021",
		StakedAmount:       "1000000000utmotus",
		NetEarnings:        "4268402637utmotus",
		IpAddr:             "",
		LastTimeChallenged: "2023-05-08 14:33:56.656465058 +0000 UTC",
		CoolDownTolerance:  "2",
		GuardAddress:       "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetRunner(ctx, item)

	gaurd := types.Guard{
		Index:   "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
		GuardId: "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetGuard(ctx, gaurd)

	res, err := msgServer.ClaimRunnerRewards(context, &types.MsgClaimRunnerRewards{
		Creator: CREATOR,
		Amount:  Amount,
	})

	t.Log("response", res)

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid runner key, response should raise proper error message*/
func Test_RegisterNotValidClaimRunnerReward_KeyNotFound(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Runner{
		Score:              "70.01360618066334",
		RewardMultiplier:   "4901.905050421021",
		StakedAmount:       "1000000000utmotus",
		NetEarnings:        "4268402637utmotus",
		IpAddr:             "",
		LastTimeChallenged: "2023-05-08 14:33:56.656465058 +0000 UTC",
		CoolDownTolerance:  "2",
		GuardAddress:       "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetRunner(ctx, item)

	gaurd := types.Guard{
		Index:   "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
		GuardId: "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetGuard(ctx, gaurd)

	res, err := msgServer.ClaimRunnerRewards(context, &types.MsgClaimRunnerRewards{
		Creator: CREATOR,
		Amount:  Amount,
	})

	t.Log("response", res, err)

	require.Error(t, err)
	require.Nil(t, res)
}

/** Using not valid amount, response should raise proper error message*/
func Test_RegisterClaimRunnerReward_InsufficientFound(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Runner{
		PubKey:             "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795",
		Address:            "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8",
		Score:              "70.01360618066334",
		RewardMultiplier:   "4901.905050421021",
		StakedAmount:       "1000000000utmotus",
		NetEarnings:        "4268402637utmotus",
		IpAddr:             "",
		LastTimeChallenged: "2023-05-08 14:33:56.656465058 +0000 UTC",
		CoolDownTolerance:  "2",
		GuardAddress:       "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetRunner(ctx, item)

	gaurd := types.Guard{
		Index:   "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
		GuardId: "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetGuard(ctx, gaurd)

	res, err := msgServer.ClaimRunnerRewards(context, &types.MsgClaimRunnerRewards{
		Creator: CREATOR,
		Amount:  InsufficientFound,
	})

	t.Log("response", res, err)

	require.Error(t, err)
	require.Nil(t, res)
}

/** Using not valid guard, response should raise proper error message*/
func Test_RegisterClaimRunnerReward_NotValidGuard(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Runner{
		PubKey:             "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795",
		Address:            "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8",
		Score:              "70.01360618066334",
		RewardMultiplier:   "4901.905050421021",
		StakedAmount:       "1000000000utmotus",
		NetEarnings:        "4268402637utmotus",
		IpAddr:             "",
		LastTimeChallenged: "2023-05-08 14:33:56.656465058 +0000 UTC",
		CoolDownTolerance:  "2",
		GuardAddress:       "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetRunner(ctx, item)

	gaurd := types.Guard{
		Index:   "",
		GuardId: "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetGuard(ctx, gaurd)

	res, err := msgServer.ClaimRunnerRewards(context, &types.MsgClaimRunnerRewards{
		Creator: CREATOR,
		Amount:  Amount,
	})

	t.Log("response", res, err)

	require.Error(t, err)
	require.Nil(t, res)
}

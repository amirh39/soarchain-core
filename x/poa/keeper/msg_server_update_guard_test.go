package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UpdateGuard(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	gaurd := types.Guard{
		Index: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		V2XChallenger: &types.Challenger{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
		V2NChallenger: &types.Challenger{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
		Runner: &types.Runner{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
	}
	k.SetGuard(ctx, gaurd)

	item := types.Client{
		Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "61.472555534405885",
		RewardMultiplier:   "3778.8750839306153",
		NetEarnings:        "107755123utmotus",
		LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
		CoolDownTolerance:  "1",
		Type:               "mini",
	}
	k.SetClient(ctx, item)

	runner := types.Runner{
		PubKey:             "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "70.01360618066334",
		RewardMultiplier:   "4901.905050421021",
		StakedAmount:       "1000000000utmotus",
		NetEarnings:        "4268402637utmotus",
		IpAddr:             "",
		LastTimeChallenged: "2023-05-08 14:33:56.656465058 +0000 UTC",
		CoolDownTolerance:  "2",
		GuardAddress:       "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetRunner(ctx, runner)

	res, err := msgServer.UpdateGuard(context, &types.MsgUpdateGuard{
		Creator: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		V2XAddr: "",
		V2NAddr: "",
	})

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Conflicts in challenger, response should raise proper error message*/
func Test_UpdateGuard_Challenger_Conflicts(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	gaurd := types.Guard{
		Index: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		V2XChallenger: &types.Challenger{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
		V2NChallenger: &types.Challenger{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
		Runner: &types.Runner{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
	}
	k.SetGuard(ctx, gaurd)

	item := types.Client{
		Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "61.472555534405885",
		RewardMultiplier:   "3778.8750839306153",
		NetEarnings:        "107755123utmotus",
		LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
		CoolDownTolerance:  "1",
		Type:               "mini",
	}
	k.SetClient(ctx, item)

	runner := types.Runner{
		PubKey:             "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "70.01360618066334",
		RewardMultiplier:   "4901.905050421021",
		StakedAmount:       "1000000000utmotus",
		NetEarnings:        "4268402637utmotus",
		IpAddr:             "",
		LastTimeChallenged: "2023-05-08 14:33:56.656465058 +0000 UTC",
		CoolDownTolerance:  "2",
		GuardAddress:       "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetRunner(ctx, runner)

	res, err := msgServer.UpdateGuard(context, &types.MsgUpdateGuard{
		Creator: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		V2XAddr: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		V2NAddr: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

/** Using not valid guard, response should raise proper error message*/
func Test_UpdateGuard_NotValidGaurd(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	gaurd := types.Guard{
		Index: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		V2XChallenger: &types.Challenger{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
		V2NChallenger: &types.Challenger{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
		Runner: &types.Runner{
			Address: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		},
	}
	k.SetGuard(ctx, gaurd)

	item := types.Client{
		Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "61.472555534405885",
		RewardMultiplier:   "3778.8750839306153",
		NetEarnings:        "107755123utmotus",
		LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
		CoolDownTolerance:  "1",
		Type:               "mini",
	}
	k.SetClient(ctx, item)

	runner := types.Runner{
		PubKey:             "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "70.01360618066334",
		RewardMultiplier:   "4901.905050421021",
		StakedAmount:       "1000000000utmotus",
		NetEarnings:        "4268402637utmotus",
		IpAddr:             "",
		LastTimeChallenged: "2023-05-08 14:33:56.656465058 +0000 UTC",
		CoolDownTolerance:  "2",
		GuardAddress:       "soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4",
	}
	k.SetRunner(ctx, runner)

	res, err := msgServer.UpdateGuard(context, &types.MsgUpdateGuard{
		Creator: "",
		V2XAddr: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		V2NAddr: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

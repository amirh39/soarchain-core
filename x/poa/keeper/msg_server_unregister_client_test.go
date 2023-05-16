package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_UnregisterClient(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Client{
		Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "61.472555534405885",
		RewardMultiplier:   "3778.8750839306153",
		NetEarnings:        "107755123utmotus",
		LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
		CoolDownTolerance:  "1",
	}
	k.SetClient(ctx, item)

	motusWallet := types.MotusWallet{
		Index: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Client: &types.Client{
			Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
			Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
			Score:              "61.472555534405885",
			RewardMultiplier:   "3778.8750839306153",
			NetEarnings:        "107755123utmotus",
			LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
			CoolDownTolerance:  "1",
		},
	}
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.UnregisterClient(context, &types.MsgUnregisterClient{
		Creator: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Pubkey:  "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
	})

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid client, response should raise proper error message*/
func Test_UnregisterChallenger_NotValidClient(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Client{
		Index:              "",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "61.472555534405885",
		RewardMultiplier:   "3778.8750839306153",
		NetEarnings:        "107755123utmotus",
		LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
		CoolDownTolerance:  "1",
	}
	k.SetClient(ctx, item)

	motusWallet := types.MotusWallet{
		Index: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Client: &types.Client{
			Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
			Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
			Score:              "61.472555534405885",
			RewardMultiplier:   "3778.8750839306153",
			NetEarnings:        "107755123utmotus",
			LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
			CoolDownTolerance:  "1",
		},
	}
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.UnregisterClient(context, &types.MsgUnregisterClient{
		Creator: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Pubkey:  "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

/** Using not valid motus wallet, response should raise proper error message*/
func Test_RegisterClaimRunnerReward_NotValidMotusWallet(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.Client{
		Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
		Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Score:              "61.472555534405885",
		RewardMultiplier:   "3778.8750839306153",
		NetEarnings:        "107755123utmotus",
		LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
		CoolDownTolerance:  "1",
	}
	k.SetClient(ctx, item)

	motusWallet := types.MotusWallet{
		Index: "",
		Client: &types.Client{
			Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
			Address:            "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
			Score:              "61.472555534405885",
			RewardMultiplier:   "3778.8750839306153",
			NetEarnings:        "107755123utmotus",
			LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
			CoolDownTolerance:  "1",
		},
	}
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.UnregisterClient(context, &types.MsgUnregisterClient{
		Creator: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Pubkey:  "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

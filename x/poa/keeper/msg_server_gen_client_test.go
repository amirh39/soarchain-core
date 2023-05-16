package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_GenClient(t *testing.T) {
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
		Type:               "mini",
	}
	k.SetClient(ctx, item)

	motusWallet := types.MotusWallet{
		Index: "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n",
		Client: &types.Client{
			Index:              "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817",
			Address:            "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a",
			Score:              "61.472555534405885",
			RewardMultiplier:   "3778.8750839306153",
			NetEarnings:        "107755123utmotus",
			LastTimeChallenged: "2023-05-08 19:14:55.666272303 +0000 UTC",
			CoolDownTolerance:  "1",
			Type:               "mini",
		},
	}
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a",
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----",
		Signature:   "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0",
	})

	t.Log("response", res)
	t.Log("err", err)

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid client, response should raise proper error message*/
func Test_GenClient_NotValidClient(t *testing.T) {
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
		Type:               "mini",
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
			Type:               "mini",
		},
	}
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     "",
		Certificate: "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----",
		Signature:   "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0",
	})

	t.Log("response", res)

	require.Error(t, err)
	require.Nil(t, res)
}

func Test_GenClient_NotValidCertificate(t *testing.T) {
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
		Type:               "mini",
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
			Type:               "mini",
		},
	}
	k.SetMotusWallet(ctx, motusWallet)

	res, err := msgServer.GenClient(context, &types.MsgGenClient{
		Creator:     "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a",
		Certificate: "--------\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----",
		Signature:   "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0",
	})

	t.Log("error", err)

	require.Error(t, err)
	require.Nil(t, res)
}

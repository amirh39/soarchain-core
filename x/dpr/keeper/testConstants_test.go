/** This file is created for tests. Firstly search what you need if not find then create a new one for you. */
package keeper_test

import (
	"context"
	"soarchain/x/dpr"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/testutil"
	"soarchain/x/dpr/types"
	"strconv"
	"testing"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	epochKeeper "soarchain/x/epoch/keeper"
	epochTypes "soarchain/x/epoch/types"
)

func CreateDpr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = DprId
		items[i].Creator = CREATOR
		items[i].Duration = 12
		items[i].Status = 0

		keeper.SetDpr(ctx, items[i])
	}
	return items
}

func SetupDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i] = types.Dpr{
			Id:      DprId, // Should be unique for each DPR
			Creator: CREATOR,
			SupportedPIDs: &types.SupportedPIDs{
				Pid_1To_20:  "AAAAAAAA",
				Pid_21To_40: "AAAAAAAA",
				Pid_41To_60: "AAAAAAAA",
				Pid_61To_80: "AAAAAAAA",
				Pid_81To_A0: "AAAAAAAA",
				Pid_A1To_C0: "AAAAAAAA",
				Pid_C1To_E0: "AAAAAAAA",
				Pid_SVCTo_9: "",
			},
			Status:         1,
			Duration:       DprDuration,
			DprEndTime:     "",
			DprStartEpoch:  0,
			DprBudget:      "1000000udmotus",
			MaxClientCount: MaxClientCount,
			ClientCounter:  1,
		}
	}
	return items
}

func SetupSecondDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i] = types.Dpr{
			Id:             DprID, // Should be unique for each DPR
			Creator:        CREATOR,
			SupportedPIDs:  &types.SupportedPIDs{},
			Status:         3,
			Duration:       DprDuration,
			DprEndTime:     "",
			DprStartEpoch:  0,
			DprBudget:      "1000000udmotus",
			MaxClientCount: MaxClientCount,
			ClientCounter:  1,
			Name:           "soarDpr",
		}
	}
	return items
}

func CreateDeactiveDpr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)
		items[i].Creator = strconv.Itoa(i)
		items[i].Duration = 12
		items[i].Status = 0

		keeper.SetDpr(ctx, items[i])
	}
	return items
}

func CreateAeactiveDpr(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)
		items[i].Creator = strconv.Itoa(i)
		items[i].Duration = 12
		items[i].Status = 0
		keeper.SetDpr(ctx, items[i])
	}
	return items
}

func SetupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.DprKeeperWithMocks(t, bankMock)
	dpr.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

func SetupActiveDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i] = types.Dpr{
			Id:             DprID, // Should be unique for each DPR
			Creator:        CREATOR,
			SupportedPIDs:  &types.SupportedPIDs{},
			Status:         0,
			Duration:       DprDuration,
			DprEndTime:     "",
			DprStartEpoch:  0,
			DprBudget:      "1000000udmotus",
			MaxClientCount: MaxClientCount,
			ClientCounter:  1,
			Name:           "soarDpr",
		}
	}
	return items
}

const (
	CREATOR = "soar1dx4yutqz8kmdfwejexxvtljch6j8x8nqnvywqp"
	ADDRESS = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	PUBKEY  = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	VIN     = "1HGCM82636c678d14c93ad5bf14448da57f4f241b77e30a013d54f5d76c8126a7029aeb86"
)

var PIDS = []bool{true, false, false}

const (
	Did                  = "did:soar:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm"
	SecondDid            = "did1:soar:1Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgap"
	VerificationMethodId = Did + "#key1"
	DprId                = "990d4bfb-7fdb-47c9-9b87-eef81dd79a7f"
)

func CreateEpochData(keeper *epochKeeper.Keeper, ctx sdk.Context) epochTypes.EpochData {
	item := epochTypes.EpochData{
		TotalEpochs:                   30,
		EpochV2VRX:                    "2udmotus",
		EpochV2VBX:                    "3udmotus",
		EpochV2NBX:                    "4udmotus",
		EpochRunner:                   "5udmotus",
		EpochChallenger:               "6",
		V2VRXTotalChallenges:          7,
		V2VBXTotalChallenges:          8,
		V2NBXTotalChallenges:          9,
		RunnerTotalChallenges:         10,
		ChallengerTotalChallenges:     11,
		V2VRXLastBlockChallenges:      1,
		V2VBXLastBlockChallenges:      1,
		V2NBXLastBlockChallenges:      1,
		RunnerLastBlockChallenges:     1,
		ChallengerLastBlockChallenges: 1,
		ChallengerPerChallengeValue:   1000000,
		V2NBXPerChallengeValue:        3000000,
		RunnerPerChallengeValue:       1000000,
		InitialPerChallengeValue:      9000000.0,
		TotalChallengesPrevDay:        99,
		V2VBXPerChallengeValue:        3000000,
		V2VRXPerChallengeValue:        3000000,
	}
	keeper.SetEpochData(ctx, item)
	return item
}

const (
	INDEX        = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	TYPE         = "mini"
	PUBKEYTYPE   = "Pubkey-Type"
	CONTROLLER   = "Controller"
	PUBLICKEYPEM = "-----CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"
	OWNERID      = "Owner-Id"
	PURCHESDATE  = "Purches-Date"
)

const (
	Signature   = "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0"
	Certificate = "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
)

const (
	MASTER_CERTIFICATE = "-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"
	MASTER_ACCOUNT     = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
)

const (
	ClientPubKey                = "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817"
	ClientPubKey2               = "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5819"
	ClientAddress               = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	ClientScroe                 = "80"
	LastTimeChallenged          = "2023-01-06 11:05:17.40125 +0000 UTC"
	CoolDownTolerance           = "1"
	NotValid_LastTimeChallenged = "-01-06 11:05:17.40125 +0000 UTC"
	NotValid_CoolDownTolerance  = "10"
	NotValid_Score              = ""
	ClientRewardMultiplier      = "3778.8750839306153"
	ClientNetEarnings           = "107755123udmotus"
)

const (
	DprID          = "unique_dpr_id" // Replace with actual logic to generate unique IDs
	Creator        = "creator_id"    // Replace with actual creator id
	InitialBudget  = "1000000udmotus"
	DprDuration    = uint64(10) // 10 can be replaced with any default duration
	MaxClientCount = uint64(10) // Example maximum client count
)

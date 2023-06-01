/** This file is created for tests. Firstly search what you nee if not find then create a new one for you. */
package keeper_test

import (
	"context"
	"soarchain/x/poa"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/testutil"
	"soarchain/x/poa/types"
	"strconv"
	"testing"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func CreateMasterKey(keeper *keeper.Keeper, ctx sdk.Context) types.MasterKey {
	item := types.MasterKey{}
	keeper.SetMasterKey(ctx, item)
	return item
}

func CreateNRunner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].PubKey = RunnerPubKey
		items[i].Address = RunnerAddress
		items[i].Score = RunnerScore
		items[i].RewardMultiplier = RunnerRewardMultiplier
		items[i].StakedAmount = RunnerStakedAmount
		items[i].NetEarnings = RunnerNetEarnings
		items[i].IpAddr = ""
		items[i].LastTimeChallenged = RunnerLastTimeChallenged
		items[i].CoolDownTolerance = RunnerCoolDownTolerance
		keeper.SetRunner(ctx, items[i])
	}
	return items
}

func SetupNRunner(n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].PubKey = RunnerPubKey
		items[i].Address = RunnerAddress
		items[i].Score = RunnerScore
		items[i].RewardMultiplier = RunnerRewardMultiplier
		items[i].StakedAmount = RunnerStakedAmount
		items[i].NetEarnings = RunnerNetEarnings
		items[i].IpAddr = ""
		items[i].LastTimeChallenged = RunnerLastTimeChallenged
		items[i].CoolDownTolerance = RunnerCoolDownTolerance
	}
	return items
}

func CreateNRunnersForRewardCap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].IpAddr = "45.12.65.78"
		items[i].RewardMultiplier = "4"

		keeper.SetRunner(ctx, items[i])
	}
	return items
}

func CreateNClient(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetClient(ctx, items[i])
	}
	return items
}

func SetupClientEntity(n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = ClientPubKey
		items[i].Address = ClientAddress
		items[i].Score = ClientScore
		items[i].RewardMultiplier = ClientRewardMultiplier
		items[i].NetEarnings = ClientNetEarnings
		items[i].LastTimeChallenged = ClientLastTimeChallenged
		items[i].CoolDownTolerance = ClientCoolDownTolerance
		items[i].Type = ClientType
	}
	return items
}

func CreateRewardCapClient(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].RewardMultiplier = "4"
		keeper.SetClient(ctx, items[i])
	}
	return items
}

func CreateInValidClientScore(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(12)
		items[i].Score = NotValid_Score
		items[i].LastTimeChallenged = NotValid_LastTimeChallenged
		items[i].CoolDownTolerance = NotValid_CoolDownTolerance

		keeper.SetClient(ctx, items[i])
	}
	return items
}

func SetupClientWithInvalidScore(n int) []types.Client {
	items := make([]types.Client, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(12)
		items[i].Score = NotValid_Score
		items[i].LastTimeChallenged = NotValid_LastTimeChallenged
		items[i].CoolDownTolerance = NotValid_CoolDownTolerance
	}
	return items
}

func CreateNChallenger(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenger {
	items := make([]types.Challenger, n)
	for i := range items {
		items[i].PubKey = Challenger_PubKey
		items[i].Address = Challenger_Address
		items[i].Score = Challenger_Score
		items[i].StakedAmount = Challenger_StakedAmount
		items[i].NetEarnings = Challenger_NetEarnings
		items[i].IpAddr = ""
		items[i].Type = Challenger_Type

		keeper.SetChallenger(ctx, items[i])
	}
	return items

}

func SetupNChallenger(n int) []types.Challenger {
	items := make([]types.Challenger, n)
	for i := range items {
		items[i].PubKey = Challenger_PubKey
		items[i].Address = Challenger_Address
		items[i].Score = Challenger_Score
		items[i].StakedAmount = Challenger_StakedAmount
		items[i].NetEarnings = Challenger_NetEarnings
		items[i].IpAddr = ""
		items[i].Type = Challenger_Type
	}
	return items
}

func CreateV2NTypeChallenger(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenger {
	items := make([]types.Challenger, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Type = "v2n"

		keeper.SetChallenger(ctx, items[i])
	}
	return items
}

func CreateEpoachDataForTestRewards(keeper *keeper.Keeper, ctx sdk.Context) types.EpochData {
	item := types.EpochData{
		TotalEpochs: 30,
		EpochV2VRX:  "2",
		EpochV2VBX:  "3",
		EpochV2NBX:  "4",
		EpochRunner: "5",
	}
	keeper.SetEpochData(ctx, item)
	return item
}

func CreateEpochData(keeper *keeper.Keeper, ctx sdk.Context) types.EpochData {
	item := types.EpochData{}
	keeper.SetEpochData(ctx, item)
	return item
}

func CreateNFactoryKeys(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FactoryKeys {
	items := make([]types.FactoryKeys, n)
	for i := range items {
		items[i].Id = keeper.AppendFactoryKeys(ctx, items[i])
	}
	return items
}

func CreateNMotusWallet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MotusWallet {
	items := make([]types.MotusWallet, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMotusWallet(ctx, items[i])
	}
	return items
}

func SetupMotusWalletEntityByClient(client types.Client) types.MotusWallet {
	motusWallet := types.MotusWallet{
		Index:  MotusWallet_Index,
		Client: &client,
	}
	return motusWallet
}

func SetupMsgServerClaimMotusRewards(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.PoaKeeperWithMocks(t, bankMock)
	poa.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

func CreateNVrfData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VrfData {
	items := make([]types.VrfData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetVrfData(ctx, items[i])
	}
	return items
}

const (
	ClientScore              = "61.472555534405885"
	ClientRewardMultiplier   = "3778.8750839306153"
	ClientNetEarnings        = "107755123umotus"
	ClientLastTimeChallenged = "2023-05-08 19:14:55.666272303 +0000 UTC"
	ClientCoolDownTolerance  = "1"
	ClientType               = "mini"
)

const (
	MaxValUint64 = 18446744073709551615
	TestUser     = "testUser"
	Factor       = 10
)

const (
	ChallengerPubKey = "3056301006072a8648ce3d020106052b8104000a03420004c4039cc2459a57357707620ddbbaddfeda5d4c66cc9ac9c3aac997e65f16b78253b3f9241182014246c1b945595c1ed2463e22ca59f153a74fee375e23a86561"
)

const (
	ClientPubKey                = "3059301306072a8648ce3d020106082a8648ce3d0301070342000402a530fa9267e1518e4d9069de38f2aecd3b508a2aca8b6d9cbd1b36b3b412e6db603ba6230728a7803acfdc8e57a21d24f648e10db24b4c957a2b2dad9a5817"
	ClientAddress               = "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n"
	ClientScroe                 = "80"
	LastTimeChallenged          = "2023-01-06 11:05:17.40125 +0000 UTC"
	CoolDownTolerance           = "1"
	NotValid_LastTimeChallenged = "-01-06 11:05:17.40125 +0000 UTC"
	NotValid_CoolDownTolerance  = "10"
	NotValid_Score              = ""
)

const (
	GeneratedNumber_Pubkey  = "66eea999dcfb6fa4df8a5d2b22ea5e637d65ff9525e5f58f5e27bdac457c0450"
	GeberatedNumber_Message = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y,1"
	GeberatedNumber_Vrv     = "4afbf8af915f46626cadcff67ea7eee354fb6b8a3645de238126355fe524cd8c"
	GeberatedNumber_Proof   = "e8cd528e10b85b629bd836b2f71a964cd4c2734f8136093d41e677b3c98fbb0e72f2f53371f6f4b068c3d05370d383f4b6e2ca59b5b71a745c7207c3dc754a0d58bd4cbbc630906c70c214cfdcbedfbd649627da37d8e53ce8cc14168b3e792b"
)

const (
	RandomNumber_Message = "message"
	RandomNumber_Vrv     = "3"
	RandomNumber_Proof   = "3"
)

const (
	RunnerPubKey             = "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795"
	RunnerAddress            = "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n"
	RunnerCreator            = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
	RunnerScore              = "70.01360618066334"
	RunnerRewardMultiplier   = "4901.905050421021"
	RunnerStakedAmount       = "1000000000umotus"
	RunnerNetEarnings        = "4268402637umotus"
	RunnerIP                 = "45.12.65.78"
	RunnerLastTimeChallenged = "2023-05-08 14:33:56.656465058 +0000 UTC"
	RunnerCoolDownTolerance  = "2"
)

const (
	Amount            = "0soar"
	InsufficientFound = "10soar"
)

const (
	CERTIFICATE         = "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
	NOTVALIDCERTIFICATE = "---------\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
	MASTER_CERTIFICATE  = "-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"
	MASTER_ACCOUNT      = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
	CREATOR             = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
)

const (
	FACTORY_CERT   = "-----BEGIN CERTIFICATE-----\nMIICBjCCAaygAwIBAgIQYuzJOUKNHYpHJFGtxphGmzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDQwNDEzMDAwMFoXDTMzMDQwNDEzMDAwMFowWjEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjE6MDgGA1UEAwwxU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBJbnQuIENBIDB4MDAwMTAyRkZGRjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIPvuGA6Q/Z1+lyExgKRM/v4bH77K3cGEKrfkQ/0ZQNhDbSfzKvrvDiKNPWYN1LhRgWcLzDguDkKisM8h1Jw2SGjZjBkMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBTL6JassDcrTsRXdHSf9H1Rg+83VzAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAKBggqhkjOPQQDAgNIADBFAiAtY1bj66UiOLJaj8EMHdeCiMtu/TAwhx1ackbwYj6sOQIhAOx2lNKLmXqt1U5StSM3jZpI8w5dNStYigv8CcABJn0k\n-----END CERTIFICATE-----"
	FACTORY_CERT_1 = "-----BEGIN CERTIFICATE-----\nMIICBTCCAaygAwIBAgIQQ3xdlGvyhzsjzBtlMGgxyzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2MDAwMFoXDTMzMDMzMDA2MDAwMFowWjEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjE6MDgGA1UEAwwxU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBJbnQuIENBIDB4MDAwMTAyRkZGRjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGTxRUWGFOHno7Zoyviovm88fE6XsQE+lYMWzO5/5jSt+5J04VE1OFu3c6L6jdHMzJa1gD4jfy+fz2nFH+MLv+ijZjBkMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBQNArqidpFro5X7u8dCajzesYWDQjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAKBggqhkjOPQQDAgNHADBEAiBcPLxNAtGaLrahBqmq5oFbBOHyCLhl9GdP4ZBjPi1AmAIgd6CY0+ZCQKeRFu0nZpypBXvegoEy4UgZb5MDU4tXRMU=\n-----END CERTIFICATE-----"
)

const (
	Valid_CertString   = "-----BEGIN CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"
	INValid_CertString = "-----CERTIFICATE-----\nMIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRowGAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1YrU2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVhDKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrSk01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++DOw==\n-----END CERTIFICATE-----"
	Signature          = "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0"
	CertCreator        = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
)

const (
	Challenger_PubKey       = "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251"
	Challenger_Address      = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score        = "189"
	Challenger_StakedAmount = "2000000000umotus"
	Challenger_NetEarnings  = "0umotus"
	Challenger_IpAddr       = ""
	Challenger_IPAddress    = "104.248.142.45"
	Challenger_Type         = "v2n"
	Challenger_Creator      = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
)

const (
	MotusWallet_Index = "soar1k9ee7xx2mqzehrt56y7ezyqnegzfy8afrs754n"
)

const (
	CommunityWallet = "soar1fkl5vm32776a5h6v7ra5pq9fnckxx3nl2mkjnx"
)

const (
	ADDRESS      = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	StakedAmount = "2000000000umotus"
	IP           = "104.248.142.45"
	PUBLICKEY    = "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795"
)

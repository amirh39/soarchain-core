/** This file is created for tests. Firstly search what you nee if not find then create a new one for you. */
package keeper_test

import (
	"context"
	"soarchain/x/did/keeper"
	"soarchain/x/did/types"
	"soarchain/x/dpr/testutil"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	keepertest "soarchain/testutil/keeper"
)

func SetupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.DidKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

func CreateNChallengerDid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChallengerDid {
	items := make([]types.ChallengerDid, n)
	for i := range items {
		items[i].PubKey = Challenger_PubKey
		items[i].Address = Challenger_Address

		keeper.SetChallengerDid(ctx, items[i])
	}
	return items
}

const (
	Challenger_PubKey        = "3056301006072a8648ce3d020106052b8104000a0342000421ac05e92e7906b648ee7029e1dc9599bde61372be4bf2b41806de08c362052d4ebcc9f6c24dbd5f33df3a1d0419ab017991df2671db0dd4aa2661fe4bbf8251"
	Challenger_Address       = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score         = "189"
	Challenger_StakedAmount  = "2000000000utmotus"
	Challenger_NetEarnings   = "0utmotus"
	Challenger_StakedAmount2 = "2000000000udmotus"
	Challenger_NetEarnings2  = "0udmotus"
	Challenger_IpAddr        = ""
	Challenger_IPAddress     = "104.248.142.45"
	Challenger_Type          = "v2n"
	Challenger_Creator       = "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y"
	Challenger_Score2        = "82"
)

const (
	RunnerStake = "10000000000000000udmotus"
	RunnerIp    = "145.21.09.11"
)

const (
	Did           = "did:soar:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm"
	SecondDid     = "did:soar:1Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgap"
	SecondAddress = "my-second-address"
	SecondPubKey  = "my-second-pubkey"
)

const (
	ADDRESS  = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	ADDRESS2 = "soar1z68ga5gq0pks3zla7sg4c9svv8nmggr537m83c"
	PUBKEY   = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	VIN      = "1HGCM82636c678d14c93ad5bf14448da57f4f241b77e30a013d54f5d76c8126a7029aeb86"
)

var PIDS = []bool{true, false, false}

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
	Signature    = "3046022100f39e6fa9ad5a4b43bd2dc8a42746ad9f768ab19edb24dacd4674a044c7a5a80f022100a34aff23aa9dafa3a0847a7350135b59eb1a5d45990119e263ed0cb23554c370"
	Certificate  = "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQQLzcQmX2Ob3cpGVXYVv+WzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDUxMTA4MDAwMFoXDTMzMDUxMTA4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wNTcyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9iaF1LSDLysmZMQ8GYcDj5d/A/iQr+WjfRindSQPl6/L7bMRgD+jyumyuBxthwRgNMagud9K1lzFgEhiJ5nLO6NBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIhALOFBzS0MAVIizvpkhWObJqidQtZfi04X9z3RoROUT6jAiALMt4XuRAzdCzlsOFTitJfK/a6OXWORX75jkY/GAcb3A==\n-----END CERTIFICATE-----"
	FactoryCert  = "-----BEGIN CERTIFICATE-----\nMIICBjCCAaygAwIBAgIQYuzJOUKNHYpHJFGtxphGmzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDQwNDEzMDAwMFoXDTMzMDQwNDEzMDAwMFowWjEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjE6MDgGA1UEAwwxU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBJbnQuIENBIDB4MDAwMTAyRkZGRjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIPvuGA6Q/Z1+lyExgKRM/v4bH77K3cGEKrfkQ/0ZQNhDbSfzKvrvDiKNPWYN1LhRgWcLzDguDkKisM8h1Jw2SGjZjBkMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBTL6JassDcrTsRXdHSf9H1Rg+83VzAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAKBggqhkjOPQQDAgNIADBFAiAtY1bj66UiOLJaj8EMHdeCiMtu/TAwhx1ackbwYj6sOQIhAOx2lNKLmXqt1U5StSM3jZpI8w5dNStYigv8CcABJn0k\n-----END CERTIFICATE-----"
	FactoryCert2 = "-----BEGIN CERTIFICATE-----\nMIICBTCCAaygAwIBAgIQQ3xdlGvyhzsjzBtlMGgxyzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2MDAwMFoXDTMzMDMzMDA2MDAwMFowWjEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjE6MDgGA1UEAwwxU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBJbnQuIENBIDB4MDAwMTAyRkZGRjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGTxRUWGFOHno7Zoyviovm88fE6XsQE+lYMWzO5/5jSt+5J04VE1OFu3c6L6jdHMzJa1gD4jfy+fz2nFH+MLv+ijZjBkMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBQNArqidpFro5X7u8dCajzesYWDQjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAKBggqhkjOPQQDAgNHADBEAiBcPLxNAtGaLrahBqmq5oFbBOHyCLhl9GdP4ZBjPi1AmAIgd6CY0+ZCQKeRFu0nZpypBXvegoEy4UgZb5MDU4tXRMU=\n-----END CERTIFICATE-----"
)

const (
	MASTER_CERTIFICATE = "-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"
	MASTER_ACCOUNT     = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
)

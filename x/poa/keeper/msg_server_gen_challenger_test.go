package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	CERTIFICATE         = "-----BEGIN CERTIFICATE-----\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
	NOTVALIDCERTIFICATE = "---------\nMIIB1DCCAXqgAwIBAgIQarjUOnCZTyR62V1ecTpJOzAKBggqhkjOPQQDAjBaMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMTowOAYDVQQDDDFTb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IEludC4gQ0EgMHgwMDAxMDJGRkZGMB4XDTIzMDQwNjE4MDAwMFoXDTMzMDQwNjE4MDAwMFowOzEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEbMBkGA1UEAwwSU0FNUExFX0RFVklDRV8wMDEwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbCji79+UYAQ127pa5/GVy2GePdEot+Dih3+aHaSJAngZABw+AUHLV53D2ekTpFZEQBvSRYMT3DfRXdWK3K/xVKNBMD8wDgYDVR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUy+iWrLA3K07EV3R0n/R9UYPvN1cwCgYIKoZIzj0EAwIDSAAwRQIgLNRm2jurfwQt2mAYgzxMO6r282PTB3Bil0cbbkRWCFICIQC09z8NUdddEaT3+rPovZNtL/LukupZaBl4LseTv4c74w==\n-----END CERTIFICATE-----"
	SIGNATURE           = "3046022100b3895f069c24bcc403e5c34463b3fbd88c52088e3070265c84401388d87782f9022100ca497f09fad41001bc2958006872b67767d842a77bfd2347c614b2f6a8b11cd0"
	mASTER_ACCOUNT      = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
	cREATOR             = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	IP                  = "104.248.142.45"
	TYPE                = "v2n"
	STAKE               = "2000000000utmotus"
	ADDRESS             = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	PUBLICKEY           = "3056301006072a8648ce3d020106052b8104000a034200044c1db1a1b1e19d6c423b1af88203ce79b6e4705d1dedaf65daeb0eedbe2c1fc6db010fa7f81443229d90181691df2e209be1c1278af42cc0f5ade03db549a795"
	fACTORY_CERT        = "-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"
)

func Test_GenChallenger(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.GenChallenger(context, &types.MsgGenChallenger{
		Creator:         CREATOR,
		ChallengerStake: STAKE,
		ChallengerIp:    IP,
		Challengertype:  TYPE,
		Certificate:     CERTIFICATE,
		Signature:       SIGNATURE,
	})

	t.Log("response", res)

	require.NoError(t, err)
	require.NotNil(t, res)
}

/** Using not valid certificate, response should raise proper error message*/
func Test_GenChallenger_NotValidCertificate(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_ACCOUNT,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.GenChallenger(context, &types.MsgGenChallenger{
		Creator:         CREATOR,
		ChallengerStake: STAKE,
		ChallengerIp:    IP,
		Challengertype:  TYPE,
		Certificate:     NOTVALIDCERTIFICATE,
		Signature:       SIGNATURE,
	})

	t.Log("error message", err)

	require.Error(t, err)
	require.Nil(t, res)
}

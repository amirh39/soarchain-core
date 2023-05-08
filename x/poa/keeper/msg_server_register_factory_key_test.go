package keeper_test

import (
	"soarchain/x/poa/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	MASTER_CERTIFICATE = "-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"
	MASTER_ACCOUNT     = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
	CREATOR            = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
	FACTORY_CERT       = "-----BEGIN CERTIFICATE-----\nMIICBjCCAaygAwIBAgIQYuzJOUKNHYpHJFGtxphGmzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDQwNDEzMDAwMFoXDTMzMDQwNDEzMDAwMFowWjEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjE6MDgGA1UEAwwxU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBJbnQuIENBIDB4MDAwMTAyRkZGRjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIPvuGA6Q/Z1+lyExgKRM/v4bH77K3cGEKrfkQ/0ZQNhDbSfzKvrvDiKNPWYN1LhRgWcLzDguDkKisM8h1Jw2SGjZjBkMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0GA1UdDgQWBBTL6JassDcrTsRXdHSf9H1Rg+83VzAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAKBggqhkjOPQQDAgNIADBFAiAtY1bj66UiOLJaj8EMHdeCiMtu/TAwhx1ackbwYj6sOQIhAOx2lNKLmXqt1U5StSM3jZpI8w5dNStYigv8CcABJn0k\n-----END CERTIFICATE-----"
)

func Test_RegisterFactoryKey(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	resp, err := msgServer.RegisterFactoryKey(context, &types.MsgRegisterFactoryKey{
		Creator:     CREATOR,
		FactoryCert: FACTORY_CERT,
	})

	t.Log("response", resp)

	require.NoError(t, err)
	require.NotNil(t, resp)
}

/** Using not valid master certificate, response should raise proper error message*/
func Test_RegisterNotValidFactoryKey(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)

	item := types.MasterKey{MasterCertificate: MASTER_ACCOUNT,
		MasterAccount: MASTER_ACCOUNT,
	}
	k.SetMasterKey(ctx, item)

	res, err := msgServer.RegisterFactoryKey(context, &types.MsgRegisterFactoryKey{
		Creator:     CREATOR,
		FactoryCert: FACTORY_CERT,
	})

	t.Log("error message", err)

	require.Error(t, err)
	require.Nil(t, res)
}

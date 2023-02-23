package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

// const ecoCertString string = `-----BEGIN CERTIFICATE-----
// MIIB3TCCAYOgAwIBAgIQYdqh2xopk506MaWSwVjkxjAKBggqhkjOPQQDAjBGMRow
// GAYDVQQKDBFTb2FyIFJvYm90aWNzIEluYzEoMCYGA1UEAwwfU29hciBSb2JvdGlj
// cyBTZWN1cmUgRWxlbWVudCBDQTAeFw0yMzAyMjAxMjA1MTBaFw00ODAyMjAxMjA1
// MTBaMEYxGjAYBgNVBAoMEVNvYXIgUm9ib3RpY3MgSW5jMSgwJgYDVQQDDB9Tb2Fy
// IFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0D
// AQcDQgAEvCKtYxo9fLS9RvHunODfYwAuPm2NY6rUAMzuTk4N4rpJTDFA1aVva1Yr
// U2xQ78KHTnTgUGPm/j98oy/nB6KXNqNTMFEwHQYDVR0OBBYEFKlxhLDaJAfFXiVh
// DKI/FZP1lzb7MB8GA1UdIwQYMBaAFKlxhLDaJAfFXiVhDKI/FZP1lzb7MA8GA1Ud
// EwEB/wQFMAMBAf8wCgYIKoZIzj0EAwIDSAAwRQIhAIfk8J7lln6CNhZKwWqIgqrS
// k01jhapY1yHsDjYz32/JAiBRolIuWe6+BigqiseBfxCKPVCHKwE/FaxrWSH6j++D
// Ow==
// -----END CERTIFICATE-----`
const ecoCertFile string = "/Users/candostyavuz/Projects/repo/soarchain-core/x/poa/certs/ecosystem.crt"
const signerCertFile string = "/Users/candostyavuz/Projects/repo/soarchain-core/x/poa/certs/signer_FFFF.der"

func (k msgServer) RegisterFactoryKey(goCtx context.Context, msg *types.MsgRegisterFactoryKey) (*types.MsgRegisterFactoryKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Authorization check
	soarMasterKey, isFound := k.GetMasterKey(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Master key not found!")
	}

	if msg.Creator != soarMasterKey.MasterAccount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Not authorized!")
	}

	// Create & Verify x509 certs:
	masterCert, err := k.CreateX509CertFromString(soarMasterKey.MasterCertificate)
	if err != nil {
		return nil, err
	}
	factoryCert, err := k.CreateX509CertFromString(msg.FactoryKey)
	if err != nil {
		return nil, err
	}
	// factoryCert, err := k.CreateX509CertFromFile(signerCertFile)
	// if err != nil {
	// 	return nil, err
	// }

	result, err := k.ValidateX509Cert(factoryCert, masterCert)
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cert verification error")
	}

	// Save factory key
	totalKeys := k.GetAllFactoryKeys(ctx)
	idx := uint64(len(totalKeys))

	updatedFactoryKeyList := types.FactoryKeys{
		Id:         idx,
		FactoryKey: msg.FactoryKey,
	}

	k.SetFactoryKeys(ctx, updatedFactoryKeyList)

	return &types.MsgRegisterFactoryKeyResponse{}, nil
}

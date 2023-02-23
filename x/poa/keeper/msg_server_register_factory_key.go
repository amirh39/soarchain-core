package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

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

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
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RegisterFactoryKey][GetMasterKey] failed. Master key not found from the denesis.")
	}

	if msg.Creator != soarMasterKey.MasterAccount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[RegisterFactoryKey][GetMasterKey] failed. Couldn't authorize by given master key.")
	}

	// Create & Verify x509 certs:

	masterCert, err := k.CreateX509CertFromString(soarMasterKey.MasterCertificate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RegisterFactoryKey][CreateX509CertFromString] failed. Couldn't create x590 certificate from  the genesis."+err.Error())
	}

	factoryCert, err := k.CreateX509CertFromString(msg.FactoryCert)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RegisterFactoryKey][CreateX509CertFromString] failed. Factory certificate couldn't be created from the payload by the given master key."+err.Error())
	}

	result, err := k.ValidateX509Cert(factoryCert, masterCert)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[RegisterFactoryKey][ValidateX509Cert] failed. Factory certificate validation error."+err.Error())
	}

	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[RegisterFactoryKey][ValidateX509Cert] failed. Cert verification error.")
	}

	// Save factory key
	factoryKeys := k.GetAllFactoryKeys(ctx)

	// Find the factory key with the matching certificate for detecting duplication
	for _, key := range factoryKeys {
		if key.FactoryCert == msg.FactoryCert {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[RegisterFactoryKey] failed. Duplicating FactoryCerts.")
		}
	}

	idx := uint64(len(factoryKeys))

	updatedFactoryKeyList := types.FactoryKeys{
		Id:          idx,
		FactoryCert: msg.FactoryCert,
	}

	k.SetFactoryKeys(ctx, updatedFactoryKeyList)

	return &types.MsgRegisterFactoryKeyResponse{}, nil
}

package keeper

import (
	"context"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/soar-robotics/soarchain-core/x/poa/types"
)

func (k msgServer) RegisterFactoryKey(goCtx context.Context, msg *types.MsgRegisterFactoryKey) (*types.MsgRegisterFactoryKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Registering factory key transaction started ##############")

	// Authorization check
	soarMasterKey, isFound := k.GetMasterKey(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RegisterFactoryKey][GetMasterKey] failed. Master key not found from the genesis.")
	}

	if logger != nil {
		logger.Info("Authirization of master key successfully done.", "transaction", "RegisterFactorykey")
	}

	if msg.Creator != soarMasterKey.MasterAccount {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "[RegisterFactoryKey][GetMasterKey] failed. Couldn't authorize by given master key. Make sure creator: [ %T ] is equal with master account [ %T ]. ", msg.Creator, soarMasterKey.MasterAccount)
	}

	// Create & Verify x509 certs:
	masterCert, err := k.CreateX509CertFromString(soarMasterKey.MasterCertificate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[RegisterFactoryKey][CreateX509CertFromString] failed. Couldn't create x590 certificate from  the master certificate. Make sure master certificate is valid and not empty.")
	}

	factoryCert, err := k.CreateX509CertFromString(msg.FactoryCert)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[RegisterFactoryKey][CreateX509CertFromString] failed. Couldn't create x590 certificate from the factory certificate. Make sure factory certificate is valid and not empty.")
	}

	result, err := k.ValidateX509Cert(factoryCert, masterCert)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[RegisterFactoryKey][ValidateX509Cert] failed. Factory certificate validation error.")
	}

	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[RegisterFactoryKey] failed. Cert verification error.")
	}

	if logger != nil {
		logger.Info("Verifed master certificate by the factory certificate.")
	}

	// Save factory key
	factoryKeys := k.GetAllFactoryKeys(ctx)

	// Find the factory key with the matching certificate for detecting duplication
	for _, key := range factoryKeys {
		if key.FactoryCert == msg.FactoryCert {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[RegisterFactoryKey] failed. Duplicating Factory certification, got: [ %T ]", key.FactoryCert)
		}
	}

	idx := uint64(len(factoryKeys))

	updatedFactoryKeyList := types.FactoryKeys{
		Id:          idx,
		FactoryCert: msg.FactoryCert,
	}

	k.SetFactoryKeys(ctx, updatedFactoryKeyList)

	if logger != nil {
		logger.Info("Factory Key successfully registered into the chain", "Creator", msg.Creator)
	}

	log.Println("############## End of register factory key Transaction ##############")

	return &types.MsgRegisterFactoryKeyResponse{}, nil
}

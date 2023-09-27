package keeper

import (
	"context"
	"log"
	"strconv"

	"soarchain/x/did/constants"
	"soarchain/x/did/types"
	"soarchain/x/did/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	param "soarchain/app/params"
	poatypes "soarchain/x/poa/types"
)

func (k msgServer) GenDid(goCtx context.Context, msg *types.MsgGenDid) (*types.MsgGenDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a did Transaction Started ##############")

	result := k.ValidateInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDid][ValidateInputs] failed. Make sure transaction inputs are valid.")
	}

	deviceCert, error := CreateX509CertFromString(msg.Certificate)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenDid][CreateX509CertFromString] failed. Invalid device certificate.")
	}

	isValide := ValidateX509CertByASN1(msg.Creator, msg.Signature, deviceCert)
	if !isValide {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenDid][ValidateX509CertByASN1] failed. Invalid device certificate and signature.")
	}

	pubKeyHex, error := ExtractPubkeyFromCertificate(msg.Certificate)
	if pubKeyHex == "" || error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenDid][ExtractPubkeyFromX509Cert] failed. Invalid certificate validation. Error: [ %T ]")
	}

	if logger != nil {
		logger.Info("Verifying client certificate successfully done.", "transaction", "GenDid")
	}

	isUnique := k.IsUniqueDid(ctx, msg.Document.Id, msg.Document.Address, pubKeyHex)
	if !isUnique {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "[GenDid][IsUniqueDid] failed. Invalid certificate validation. Error: [ %T ]")
	}

	if logger != nil {
		logger.Info("Verifying unique did successfully done.", "transaction", "GenDid")
	}

	seq := types.InitialSequence
	msg.Document.Index = pubKeyHex
	didDocument := types.NewDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetDidDocument(ctx, didDocument.Document.Id, didDocument)

	if logger != nil {
		logger.Info("Generating did successfully done.", "transaction", "GenDid", "document", didDocument)
	}

	rewardMultiplier := utility.CalculateRewardMultiplier(constants.InitialScore)

	k.Keeper.poaKeeper.SetReputation(ctx, poatypes.Reputation{
		Index:              pubKeyHex,
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(constants.InitialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
	})

	log.Println("############## End of Generating did Transaction ##############")

	return &types.MsgGenDidResponse{}, nil
}

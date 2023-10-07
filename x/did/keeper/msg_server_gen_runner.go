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

func (k msgServer) GenRunner(goCtx context.Context, msg *types.MsgGenRunner) (*types.MsgGenRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a runner did Transaction Started ##############")

	result := k.ValidateRunnerInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenRunner][ValidateInputs] failed. Make sure transaction inputs are valid.")
	}

	deviceCert, error := CreateX509CertFromString(msg.Certificate)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenRunner][CreateX509CertFromString] failed. Invalid device certificate.")
	}

	isValide := ValidateX509CertByASN1(msg.Creator, msg.Signature, deviceCert)
	if !isValide {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenRunner][ValidateX509CertByASN1] failed. Invalid device certificate and signature.")
	}

	pubKeyHex, error := ExtractPubkeyFromCertificate(msg.Certificate)
	if pubKeyHex == "" || error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenRunner][ExtractPubkeyFromX509Cert] failed. Invalid certificate validation. Error: [ %T ]")
	}

	if logger != nil {
		logger.Info("Verifying runner certificate successfully done.", "transaction", "GenRunner")
	}

	isUnique := k.IsUniqueDid(ctx, msg.Document.Id)
	if isUnique {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "[GenRunner][IsUniqueDid] failed. Did is already registered.")
	}

	if logger != nil {
		logger.Info("Verifying unique did successfully done.", "transaction", "GenRunner")
	}

	seq := types.InitialSequence
	msg.Document.PubKey = pubKeyHex
	didDocument := types.NewRunnerDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetRunnerDidDocument(ctx, didDocument.Document.Id, didDocument)

	if logger != nil {
		logger.Info("Generating runner did successfully done.", "transaction", "GenRunner", "document", didDocument)
	}

	rewardMultiplier := utility.CalculateRewardMultiplier(constants.InitialScore)

	err := k.Keeper.poaKeeper.InitializeReputation(ctx, poatypes.Reputation{
		PubKey:             pubKeyHex,
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(constants.InitialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
		Type:               "",
	}, msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[GenClient][InitializeReputation] failed. Invalid certificate validation.")
	}

	log.Println("############## End of Generating did Transaction ##############")

	return &types.MsgGenRunnerResponse{}, nil
}

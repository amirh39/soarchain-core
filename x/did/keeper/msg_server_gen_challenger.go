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

func (k msgServer) GenChallenger(goCtx context.Context, msg *types.MsgGenChallenger) (*types.MsgGenChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a challenger did Transaction Started ##############")

	result := k.ValidateChallengerInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenChallenger][ValidateInputs] failed. Make sure transaction inputs are valid.")
	}

	deviceCert, error := CreateX509CertFromString(msg.Certificate)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenChallenger][CreateX509CertFromString] failed. Invalid device certificate.")
	}

	isValide := ValidateX509CertByASN1(msg.Creator, msg.Signature, deviceCert)
	if !isValide {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenChallenger][ValidateX509CertByASN1] failed. Invalid device certificate and signature.")
	}

	pubKeyHex, error := ExtractPubkeyFromCertificate(msg.Certificate)
	if pubKeyHex == "" || error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenChallenger][ExtractPubkeyFromX509Cert] failed. Invalid certificate validation. Error: [ %T ]")
	}

	if logger != nil {
		logger.Info("Verifying client certificate successfully done.", "transaction", "GenRunner")
	}

	isUnique := k.IsUniqueDid(ctx, msg.Document.Id)
	if isUnique {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "[GenChallenger][IsUniqueDid] failed. Did is already registered.")
	}

	if logger != nil {
		logger.Info("Verifying unique did successfully done.", "transaction", "GenRunner")
	}

	seq := types.InitialSequence
	msg.Document.PubKey = pubKeyHex
	didDocument := types.NewChallengerDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetChallengerDidDocument(ctx, didDocument.Document.Id, didDocument)

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
		Type:               msg.ChallengerType,
	}, msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[GenChallenger][InitializeReputation] failed. Invalid certificate validation.")
	}

	log.Println("############## End of Generating challenger did Transaction ##############")

	return &types.MsgGenChallengerResponse{}, nil
}

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

	result := k.ChallengerDidValidateInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenChallenger][ChallengerDidValidateInputs] failed. Make sure transaction inputs are valid.")
	}

	if msg.ChallengerType != constants.V2NChallengerType && msg.ChallengerType != constants.V2XChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenChallenger][ValidateChallengerType] failed. Invalid challenger type. Must be 'v2n' or 'v2x'.")
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
		logger.Info("Verifying challenger certificate successfully done.", "transaction", "GeChallenger")
	}

	isUnique := k.IsNotUniqueDid(ctx, msg.Document.Id)
	if isUnique {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "[GenChallenger][IsNotUniqueDid] failed. Did is already registered.")
	}

	if logger != nil {
		logger.Info("Verifying unique did successfully done.", "transaction", "GenChallenger")
	}

	// check if the address is uniqe
	isUniqueAddress := IsUniqueAddress(k, ctx, msg.Creator)
	if isUniqueAddress {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][IsUniqueAddress] failed. Challenger did with the address [ %T ] is already registered.", msg.Creator)
	}

	// check if the pubKey is uniqe
	isUniquePubkey := IsUniquePubKey(k, ctx, pubKeyHex)
	if isUniquePubkey {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][IsUniquePubKey] failed. Challenger did with the PubKey [ %T ] is already registered.", pubKeyHex)
	}

	if logger != nil {
		logger.Info("Checking for challenger did address and pubKey successfully done.", "transaction", "GenChallengerDid")
	}

	seq := types.InitialSequence
	msg.Document.PubKey = pubKeyHex
	didDocument := types.NewChallengerDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetChallengerDid(ctx, *didDocument.Document)

	if logger != nil {
		logger.Info("Generating challenger did successfully done.", "transaction", "GenChallenger", "document", didDocument)
	}

	rewardMultiplier := utility.CalculateRewardMultiplier(constants.InitialScore)

	err := k.Keeper.poaKeeper.InitializeReputation(ctx, poatypes.Reputation{
		PubKey:             pubKeyHex,
		Score:              strconv.FormatFloat(constants.InitialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
		Type:               msg.ChallengerType,
	}, msg.Certificate, msg.ChallengerStake, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[GenChallenger][InitializeReputation] failed. Invalid certificate validation.")
	}

	log.Println("############## End of Generating challenger did Transaction ##############")

	return &types.MsgGenChallengerResponse{}, nil
}

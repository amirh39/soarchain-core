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

	result := k.RunnerDidValidateInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenRunner][RunnerDidValidateInputs] failed. Make sure transaction inputs are valid.")
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenRunner][ExtractPubkeyFromX509Cert] failed. Invalid certificate validation.")
	}

	if logger != nil {
		logger.Info("Verifying runner certificate successfully done.", "transaction", "GenRunner")
	}

	isUnique := k.IsNotUniqueDid(ctx, msg.Document.Id)
	if isUnique {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "[GenRunner][IsNotUniqueDid] failed. Did is already registered.")
	}

	if logger != nil {
		logger.Info("Verifying unique did successfully done.", "transaction", "GenRunner")
	}

	// check if the address is uniqe
	isUniqueAddress := IsUniqueAddress(k, ctx, msg.Creator)
	if isUniqueAddress {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][IsUniqueAddress] failed. Runner did with the address [ %s ] is already registered.", msg.Creator)
	}

	// check if the pubKey is uniqe
	isUniquePubkey := IsUniquePubKey(k, ctx, pubKeyHex)
	if isUniquePubkey {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][IsUniquePubKey] failed. Runner did with the PubKey [ %s ] is already registered.", pubKeyHex)
	}

	if logger != nil {
		logger.Info("Checking for runner did address and pubKey successfully done.", "transaction", "GenRunnerDid")
	}

	if msg.Creator != msg.Document.Address {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner] failed. Runner did address [ %s ]  is not creator address.", msg.Document.Address)
	}

	seq := types.InitialSequence
	msg.Document.PubKey = pubKeyHex
	msg.Document.Address = msg.Creator
	didDocument := types.NewRunnerDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetRunnerDid(ctx, *didDocument.Document)

	_, found := k.GetRunnerDid(ctx, msg.Creator)
	if !found {
		logger.Error("Generating runner did failed.", "transaction", "GenRunner", "document", didDocument)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][GetRunnerDid] failed. Couldn't store Runner object successfully.")
	}

	if logger != nil {
		logger.Info("Generating runner did successfully done.", "transaction", "GenRunner", "document", didDocument)
	}

	rewardMultiplier := utility.CalculateRewardMultiplier(constants.InitialScore)

	initializeError := k.Keeper.poaKeeper.InitializeReputation(ctx, poatypes.Reputation{
		PubKey:             pubKeyHex,
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(constants.InitialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
		Type:               "",
		StakedAmount:       msg.RunnerStake,
	}, msg.Certificate, msg.RunnerStake, msg.Creator)
	if initializeError != nil {
		k.RemoveRunnerDid(ctx, msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[GenRunner][InitializeReputation] failed. Invalid certificate validation.")
	}

	log.Println("############## End of Generating Runner did Transaction ##############")

	return &types.MsgGenRunnerResponse{}, nil
}

package keeper

import (
	"context"
	"log"
	"strconv"

	"soarchain/x/did/types"
	"soarchain/x/did/utility"

	param "soarchain/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	isValide := ValidateX509CertByASN1(msg.Creator, msg.ClientSignature, deviceCert)
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
	_, err := k.VerifyDidOwnership(msg.Document, seq, msg.Document, msg.Document.VerificationMethods[0].Id, msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenDid][VerifyDidOwnership] failed. Did not belong to the creator.")
	}

	if logger != nil {
		logger.Info("VerifyDidOwnership", "transaction", "GenDid", "msg.Document", msg.Document)
	}

	msg.Document.Index = pubKeyHex
	didDocument := types.NewDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetDidDocument(ctx, didDocument.Document.Id, didDocument)

	if logger != nil {
		logger.Info("Generating did successfully done.", "transaction", "GenDid", "document", didDocument)
	}

	rewardMultiplier := utility.CalculateRewardMultiplier(utility.InitialScore)
	reputation := types.Reputation{
		Index:              pubKeyHex,
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(utility.InitialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
	}

	k.SetReputation(ctx, reputation)

	log.Println("############## End of Generating did Transaction ##############")

	return &types.MsgGenDidResponse{}, nil
}

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

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a client did Transaction Started ##############")

	result := k.ClientDidValidateInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenClient][ClientDidValidateInputs] failed. Make sure transaction inputs are valid.")
	}

	pubKeyHex, deviceCertificate, pubkeyGeneratingError := k.GeneratePubkey(msg)
	if pubkeyGeneratingError != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenClient][GeneratePubkey] failed. Make sure transaction inputs are valid.")
	}

	if logger != nil {
		logger.Info("Verifying client certificate successfully done.", "transaction", "GenClient")
	}

	isUnique := k.IsNotUniqueDid(ctx, msg.Document.Id)
	if isUnique {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "[GenClient][IsNotUniqueDid] failed. Did: [ %s ] is already registered.", msg.Document.Id)
	}

	if logger != nil {
		logger.Info("Verifying unique did successfully done.", "transaction", "GenClient")
	}

	// check if the address is uniqe
	isUniqueAddress := IsUniqueAddress(k, ctx, msg.Creator)
	if isUniqueAddress {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient][IsUniqueAddress] failed. Client did with the address [ %s ] is already registered.", msg.Creator)
	}

	// check if the pubKey is uniqe
	isUniquePubkey := IsUniquePubKey(k, ctx, pubKeyHex)
	if isUniquePubkey {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient][IsUniquePubKey] failed. Client did with the PubKey [ %s ] is already registered.", pubKeyHex)
	}

	if logger != nil {
		logger.Info("Checking for client did address and pubKey successfully done.", "transaction", "GenClientDid")
	}

	if msg.Creator != msg.Document.Address {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient] failed. Client did address [ %s ]  is not creator address.", msg.Document.Address)
	}

	clientType := k.ClientType(deviceCertificate)

	seq := types.InitialSequence
	msg.Document.PubKey = pubKeyHex
	msg.Document.Address = msg.Creator
	msg.Document.Type = clientType
	didDocument := types.NewDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetClientDid(ctx, *didDocument.Document)

	_, found := k.GetClientDid(ctx, msg.Creator)
	if !found {
		logger.Error("Generating client did failed.", "transaction", "GenClient", "document", didDocument)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient][GetClientDid] failed. Couldn't store client object successfully.")
	}

	if logger != nil {
		logger.Info("Generating client did successfully done.", "transaction", "GenClient", "document", didDocument)
	}

	rewardMultiplier := utility.CalculateRewardMultiplier(constants.InitialScore)

	err := k.Keeper.poaKeeper.InitializeClientReputation(ctx, poatypes.Reputation{
		PubKey:             pubKeyHex,
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(constants.InitialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
		Type:               clientType,
		StakedAmount:       "",
	}, msg.Certificate)

	if err != nil {
		k.RemoveClientDid(ctx, msg.Creator)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[GenClient][InitializeReputation] failed. Invalid certificate validation.")
	}

	log.Println("############## End of Generating client did Transaction ##############")

	return &types.MsgGenClientResponse{}, nil
}

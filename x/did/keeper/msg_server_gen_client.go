package keeper

import (
	"context"
	"crypto/x509"
	"fmt"
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

func clientType(deviceCert *x509.Certificate) string {
	if len(deviceCert.Issuer.Names) < 1 || deviceCert.Issuer.Names[1].Value == nil {
		return "[GenClient][ClientType] failed. No Type for device certificate."
	}
	results := fmt.Sprintf("%v", deviceCert.Issuer.Names[1].Value)
	if results[41:43] == "01" {
		return "mini"
	} else {
		return "pro"
	}

}

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a did Transaction Started ##############")

	result := k.ClientDidValidateInputs(msg)
	if !result {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GenClient][ClientDidValidateInputs] failed. Make sure transaction inputs are valid.")
	}

	deviceCert, error := CreateX509CertFromString(msg.Certificate)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][CreateX509CertFromString] failed. Invalid device certificate.")
	}

	isValide := ValidateX509CertByASN1(msg.Creator, msg.Signature, deviceCert)
	if !isValide {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][ValidateX509CertByASN1] failed. Invalid device certificate and signature.")
	}

	pubKeyHex, error := ExtractPubkeyFromCertificate(msg.Certificate)
	if pubKeyHex == "" || error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][ExtractPubkeyFromX509Cert] failed. Invalid certificate validation. Error: [ %T ]")
	}

	if logger != nil {
		logger.Info("Verifying client certificate successfully done.", "transaction", "GenClient")
	}

	isUnique := k.IsNotUniqueDid(ctx, msg.Document.Id)
	if isUnique {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "[GenClient][IsNotUniqueDid] failed. Did is already registered.")
	}

	if logger != nil {
		logger.Info("Verifying unique did successfully done.", "transaction", "GenClient")
	}

	// check if the address is uniqe
	isUniqueAddress := IsUniqueAddress(k, ctx, msg.Creator)
	if isUniqueAddress {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient][IsUniqueAddress] failed. Client did with the address [ %T ] is already registered.", msg.Creator)
	}

	// check if the pubKey is uniqe
	isUniquePubkey := IsUniquePubKey(k, ctx, pubKeyHex)
	if isUniquePubkey {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient][IsUniquePubKey] failed. Client did with the PubKey [ %T ] is already registered.", pubKeyHex)
	}

	if logger != nil {
		logger.Info("Checking for client did address and pubKey successfully done.", "transaction", "GenClientDid")
	}

	seq := types.InitialSequence
	msg.Document.PubKey = pubKeyHex
	msg.Document.Address = msg.Creator
	msg.Document.Type = clientType(deviceCert)
	didDocument := types.NewDidDocumentWithSeq(msg.Document, uint64(seq))
	k.SetClientDid(ctx, *didDocument.Document)

	if logger != nil {
		logger.Info("Generating did successfully done.", "transaction", "GenClient", "document", didDocument)
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
		Type:               clientType(deviceCert),
	}, msg.Certificate)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[GenClient][InitializeReputation] failed. Invalid certificate validation.")
	}

	log.Println("############## End of Generating did Transaction ##############")

	return &types.MsgGenClientResponse{}, nil
}

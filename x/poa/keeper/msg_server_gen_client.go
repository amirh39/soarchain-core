package keeper

import (
	"context"
	"crypto/x509"
	"fmt"
	"strconv"

	param "soarchain/app/params"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func clientType(deviceCert *x509.Certificate) string {
	if deviceCert.Issuer.Names[1].Value == nil {
		return "No Type"
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

	if msg.Creator == "" || msg.Certificate == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient] failed. Couldn't find valid msg.creator OR msg.Certificate. got: msg.Creator [ %T ] msg.Certificate [ %T ]. Make sure you they are valid and not empty.", msg.Creator, msg.Certificate)
	}

	deviceCert, err := k.CreateX509CertFromString(msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][CreateX509CertFromString] failed. Invalid device certificate."+err.Error())
	}

	pubKeyHex, err := k.VerifyX509CertByASN1AndExtractPubkey(msg.Creator, msg.Signature, deviceCert)
	if pubKeyHex == "" || err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient][ValidateX509CertByASN1] failed. Invalid certificate validation. Error: [ %T ]", err)
	}

	// Check validity of certificate
	errCert := k.validateCertificate(ctx, deviceCert)
	if errCert != nil {
		return nil, errCert
	}

	//check if the address is uniqe
	isUniqueAddress := IsUniqueAddress(k, ctx, msg.Creator)
	if isUniqueAddress {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenClient][GetMotusWallet][GetChallengerUsingPubKey][GetRunnerUsingPubKey][GetClient] failed. Client with the address [ %T ] is already registered.", msg.Creator)
	}

	//check if the pubKey is uniqe, also check if msg.creator address have a motus wallet
	isUniquePubkey := IsUniquePubKey(k, ctx, msg.Creator, pubKeyHex)
	if isUniquePubkey {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][GetMotusWallet][GetChallengerUsingPubKey][GetRunnerUsingPubKey][GetClient] failed. Client PubKey is not uniqe OR Client is already registered.")
	}

	// rewardMultiplier
	var initialScore float64 = 50
	rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	// Save client into storage
	newClient := types.Client{
		Index:              pubKeyHex,
		Type:               clientType(deviceCert),
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(initialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
	}

	k.SetClient(ctx, newClient)

	// Register Motus client into Motus Wallet object
	newMotusWallet := types.MotusWallet{
		Index:  msg.Creator,
		Client: &newClient,
	}
	k.SetMotusWallet(ctx, newMotusWallet)

	return &types.MsgGenClientResponse{}, nil

}

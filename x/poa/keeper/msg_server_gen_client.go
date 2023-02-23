package keeper

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const deviceCertFile string = "/Users/candostyavuz/Projects/repo/soarchain-core/x/poa/cert/device_cert.pem"

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// ToDo: change pubkey field as device cert
	deviceCert, err := k.CreateX509CertFromFile(deviceCertFile)
	if err != nil {
		return nil, err
	}

	pubKeyDer, _ := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)
	pubKeyBlock := pem.Block{
		Type:  "PUBLIC_KEY",
		Bytes: pubKeyDer,
	}
	publicKeyPem := string(pem.EncodeToMemory(&pubKeyBlock))

	_, isFound := k.GetClient(ctx, publicKeyPem)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client pubkey is already registered.")
	}

	// rewardMultiplier
	var initialScore float64 = 50
	rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	// Save client into storage
	newClient := types.Client{
		Index:              publicKeyPem,
		Address:            msg.Creator,
		Score:              strconv.FormatFloat(initialScore, 'f', -1, 64),
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		NetEarnings:        sdk.NewCoin("soar", sdk.ZeroInt()).String(),
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
	}

	k.SetClient(ctx, newClient)

	// Register Motus client into Motus Wallet object
	_, isFoundWallet := k.GetMotusWallet(ctx, msg.Creator)
	_, isFoundAsChallenger := k.GetChallenger(ctx, msg.Creator)
	_, isFoundAsRunner := k.GetRunner(ctx, msg.Creator)

	if isFoundWallet || isFoundAsChallenger || isFoundAsRunner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client address is already registered.")
	}

	newMotusWallet := types.MotusWallet{
		Index:  msg.Creator,
		Client: &newClient,
	}
	k.SetMotusWallet(ctx, newMotusWallet)

	return &types.MsgGenClientResponse{}, nil
}

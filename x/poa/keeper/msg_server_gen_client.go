package keeper

import (
	"context"
	"crypto/x509"
	"encoding/hex"

	// "encoding/pem"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// ToDo: change pubkey field as device cert
	deviceCert, err := k.CreateX509CertFromString(msg.Certificate)
	if err != nil {
		return nil, err
	}

	pubKeyDer, _ := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)

	pubKeyHex := hex.EncodeToString(pubKeyDer)

	_, isFound := k.GetClient(ctx, pubKeyHex)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Client pubkey is already registered.")
	}

	// Check validity of certificate
	totalKeys := k.GetAllFactoryKeys(ctx)
	var validated bool = false
	for i := uint64(0); i <= uint64(len(totalKeys)); i++ {
		factoryKey, isFound := k.GetFactoryKeys(ctx, i)
		if isFound {
			factoryCert, err := k.CreateX509CertFromString(factoryKey.FactoryCert)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Factory certificate couldn't be created from the storage!")
			}

			validated, err = k.ValidateX509Cert(deviceCert, factoryCert)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Device certificate couldn't be verified!")
			}
			if validated {
				break
			}
		}
	}

	if !validated {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cert verification error")
	}

	// rewardMultiplier
	var initialScore float64 = 50
	rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	// Save client into storage
	newClient := types.Client{
		Index:              pubKeyHex,
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

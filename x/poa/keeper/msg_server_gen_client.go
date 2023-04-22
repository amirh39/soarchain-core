package keeper

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"strconv"

	// "encoding/pem"
	param "soarchain/app/params"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenClient(goCtx context.Context, msg *types.MsgGenClient) (*types.MsgGenClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//ToDo: change pubkey field as device cert
	deviceCert, err := k.CreateX509CertFromString(msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid device certificate")
	}

	pubKeyDer, _ := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)

	pubKeyHex := hex.EncodeToString(pubKeyDer)
	// verify the msg.Creator_Signed which basically the msg.Creator signed by the privateKey of the pubKey we just extracted from the msg.Certificate
	signature, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid signature encoding")
	}

	hashedAddr := sha256.Sum256([]byte(msg.Creator))

	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {

		if ecdsaPubKey, ok := deviceCert.PublicKey.(*ecdsa.PublicKey); ok {

			if ecdsa.VerifyASN1(ecdsaPubKey, hashedAddr[:], signature) {
				// signature is valid
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Signature verification failed")
			}
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid public key type")
		}
	}

	// Check validity of certificate
	totalKeys := k.GetAllFactoryKeys(ctx)
	var validated bool = false
	var verificationError error = nil

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
			} else if err != nil {
				verificationError = err
			}
		}
	}

	if !validated && verificationError != nil {
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
		NetEarnings:        sdk.NewCoin(param.BondDenom, sdk.ZeroInt()).String(),
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

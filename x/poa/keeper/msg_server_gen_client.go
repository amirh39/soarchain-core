package keeper

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/big"
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
		return nil, err
	}

	log.Println(msg.Certificate)
	log.Println("deviceCert:", deviceCert)
	pubKeyDer, _ := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)
	log.Println("pubkeyder", pubKeyDer)
	pubKeyHex := hex.EncodeToString(pubKeyDer)
	log.Println("pubKeyHex", pubKeyHex)
	// verify the msg.Creator_Signed which basically the msg.Creator signed by the privateKey of the pubKey we just extracted from the msg.Certificate
	signature, err := base64.StdEncoding.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid signature encoding")
	}
	log.Println("msg.sig", msg.Signature)
	log.Println("msg.creator", msg.Creator)
	log.Println("sig", signature)

	hashedAddr := sha256.Sum256([]byte(msg.Creator))
	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])
	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {
		pub, err := x509.ParsePKIXPublicKey(pubKeyDer)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid public key")
		}

		if ecdsaPubKey, ok := pub.(*ecdsa.PublicKey); ok {
			if ecdsa.Verify(ecdsaPubKey, hashedAddr[:], r, s) {
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

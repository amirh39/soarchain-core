package keeper

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"strconv"

	// "encoding/pem"
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

	pubKeyDer, err := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][MarshalPKIXPublicKey] failed. Couldn't convert a public key to PKIX."+err.Error())
	}

	pubKeyHex := hex.EncodeToString(pubKeyDer)
	// verify the msg.Creator_Signed which basically the msg.Creator signed by the privateKey of the pubKey we just extracted from the msg.Certificate
	signature, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient][DecodeString] failed. Invalid signature encoding."+err.Error())
	}

	hashedAddr := sha256.Sum256([]byte(msg.Creator))

	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {

		if ecdsaPubKey, ok := deviceCert.PublicKey.(*ecdsa.PublicKey); ok {

			if ecdsa.VerifyASN1(ecdsaPubKey, hashedAddr[:], signature) {
				// signature is valid
			} else {
				return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[GenClient][VerifyASN1] failed. Signature verification failed.")
			}
		} else {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenClient] failed. Invalid public key type.")
		}
	}

	// Check validity of certificate
	totalKeys := k.GetAllFactoryKeys(ctx)
	var validated bool = false
	var verificationError error = nil

	for i := uint64(0); i < uint64(len(totalKeys)); i++ {
		factoryKey, isFound := k.GetFactoryKeys(ctx, i)
		if isFound {
			factoryCert, err := k.CreateX509CertFromString(factoryKey.FactoryCert)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenClient][CreateX509CertFromString] failed. Factory certificate couldn't be created from the storage."+err.Error())
			}

			validated, err = k.ValidateX509Cert(deviceCert, factoryCert)
			if err != nil {
				verificationError = sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenClient][ValidateX509Cert] failed. Couldn't validate factory certificate."+err.Error())
				continue // Try next certificate
			}

			if validated {
				verificationError = nil
				break
			}
		}
	}

	// No valid certificate found
	if verificationError != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenClient][ValidateX509Cert] failed. Device certificate couldn't be verified.")
	}

	//check if the pubKey is uniqe, also check if msg.creator address have a motus wallet
	_, isFoundWallet := k.GetMotusWallet(ctx, msg.Creator)
	_, isFoundAsChallenger := k.GetChallengerUsingPubKey(ctx, pubKeyHex)
	_, isFoundAsRunner := k.GetRunnerUsingPubKey(ctx, pubKeyHex)
	_, isFoundAsClient := k.GetClient(ctx, pubKeyHex)
	if isFoundWallet || isFoundAsChallenger || isFoundAsRunner || isFoundAsClient {
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

	fmt.Print("ggggggggggggggggggggg---newClient")
	fmt.Print(newClient)

	k.SetClient(ctx, newClient)

	// Register Motus client into Motus Wallet object
	newMotusWallet := types.MotusWallet{
		Index:  msg.Creator,
		Client: &newClient,
	}
	k.SetMotusWallet(ctx, newMotusWallet)

	return &types.MsgGenClientResponse{}, nil

}

package keeper

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	params "soarchain/app/params"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenRunner(goctx context.Context, msg *types.MsgGenRunner) (*types.MsgGenRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	deviceCert, err := k.CreateX509CertFromString(msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][CreateX509CertFromString] failed. Invalid device certificate. Error: [ %T ]", err)
	}

	pubKeyFromCertificate, err := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][MarshalPKIXPublicKey] failed. Couldn't convert a public key to PKIX. Error: [ %T ]", err)
	}

	pubKeyHex := hex.EncodeToString(pubKeyFromCertificate)
	if pubKeyHex == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[GenRunner][EncodeToString] failed. Invalid the hexadecimal encoding of publick key that extracted from certificate. Error: [ %T ]", err)
	}

	signature, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][DecodeString] failed. Invalid signature encoding.Error: [ %T ]", err)
	}

	hashedAddr := sha256.Sum256([]byte(msg.Creator))

	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {

		if ecdsaPubKey, ok := deviceCert.PublicKey.(*ecdsa.PublicKey); ok {

			if ecdsa.VerifyASN1(ecdsaPubKey, hashedAddr[:], signature) {
				// signature is valid
			} else {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "[GenRunner][VerifyASN1] failed. Signature verification failed. Error: [ %T ]", err)
			}
		} else {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner] failed. Invalid public key type. Error: [ %T ]", err)
		}
	}

	msgSenderAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "msg.Creator couldn't be parsed.")
	}

	if msg.RunnerAddr == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Runner Address must be declared in the tx!")
	}

	if msg.RunnerPubKey == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Runner Public Key must be declared in the tx!")
	}

	if msg.RunnerStake == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Runner Stake must be declared in the tx!")
	}

	//check runner
	var newRunner types.Runner

	runners := k.GetAllRunner(ctx)
	for _, runner := range runners {
		if msg.RunnerPubKey == runner.PubKey {
			sdkerrors.Wrap(sdkerrors.ErrConflict, "Runner is already registered in storage.")
			break
		}
	}

	challengers := k.GetAllChallenger(ctx)
	for _, challenger := range challengers {
		if msg.RunnerPubKey == challenger.PubKey {
			sdkerrors.Wrap(sdkerrors.ErrConflict, "Runner is already registered as challenger in storage.")
			break
		}
	}

	clients := k.GetAllClient(ctx)
	for _, client := range clients {
		if msg.RunnerPubKey == client.Index {
			sdkerrors.Wrap(sdkerrors.ErrConflict, "Runner is already registered as client in storage.")
			break
		}
	}

	runnerAddr, err := sdk.AccAddressFromBech32(msg.RunnerAddr)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid runner address!")
	}

	// Check runner stake amount
	requiredStake := sdk.Coins{sdk.NewInt64Coin(params.BondDenom, 1000000000)}
	runnerStake, err := sdk.ParseCoinsNormalized(msg.RunnerStake)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
	}
	if runnerStake.IsAllLT(requiredStake) || !runnerStake.DenomsSubsetOf(requiredStake) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Sent amount of runner: "+runnerStake.String()+" is below the required stake amount "+requiredStake.String())
	}

	// Transfer stakedAmount to poa modules account:
	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, requiredStake)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake(runner) funds couldn't be transferred to POA module!")
	}
	// rewardMultiplier
	var initialScore float64 = 50
	rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	newRunner = types.Runner{
		PubKey:             msg.RunnerPubKey,
		Address:            runnerAddr.String(),
		Score:              strconv.FormatFloat(initialScore, 'f', -1, 64), // Base Score
		RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
		StakedAmount:       runnerStake.String(),
		NetEarnings:        sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		IpAddr:             msg.RunnerIp,
		LastTimeChallenged: ctx.BlockTime().String(),
		CoolDownTolerance:  strconv.FormatUint(1, 10),
	}

	k.SetRunner(ctx, newRunner)

	return &types.MsgGenRunnerResponse{}, nil
}

package keeper

import (
	"context"
	params "soarchain/app/params"
	"soarchain/x/poa/errors"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenRunner(goctx context.Context, msg *types.MsgGenRunner) (*types.MsgGenRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	if msg.Certificate == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[GenRunner] failed. Certificate must be declared in the tx.")
	}

	runnerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
	}

	if msg.RunnerStake == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[GenRunner] failed. Runner Stake must be declared in the tx.")
	}

	if msg.Signature == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[GenRunner] failed. Signature must be declared in the tx.")
	}

	deviceCert, err := k.CreateX509CertFromString(msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][CreateX509CertFromString] failed. Invalid device certificate. Error: [ %T ]", err)
	}

	pubKeyHex, err := k.VerifyX509CertByASN1AndExtractPubkey(msg.Creator, msg.Signature, deviceCert)
	if pubKeyHex == "" || err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenRunner][VerifyX509CertByASN1AndExtractPubkey] failed. Invalid certificate validation. Error: [ %T ]", err)
	}

	// Check validity of certificate
	errCert := k.validateCertificate(ctx, deviceCert)
	if errCert != nil {
		return nil, errCert
	}

	//check runner
	var newRunner types.Runner

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
	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, runnerAddr, types.ModuleName, requiredStake)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Stake(runner) funds couldn't be transferred to POA module!")
	}
	// rewardMultiplier
	var initialScore float64 = 50
	rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	newRunner = types.Runner{
		PubKey:             pubKeyHex,
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

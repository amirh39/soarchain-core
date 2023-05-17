package keeper

import (
	"context"
	params "soarchain/app/params"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenRunner(goctx context.Context, msg *types.MsgGenRunner) (*types.MsgGenRunnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	msgSenderAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "msg.Creator couldn't be parsed.")
	}

	if msg.RunnerStake == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Runner Stake must be declared in the tx!")
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
	totalKeys := k.GetAllFactoryKeys(ctx)
	var validated bool = false
	var verificationError error = nil

	for i := uint64(0); i < uint64(len(totalKeys)); i++ {
		factoryKey, isFound := k.GetFactoryKeys(ctx, i)
		if isFound {
			factoryCert, err := k.CreateX509CertFromString(factoryKey.FactoryCert)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenRunner][CreateX509CertFromString] failed. Factory certificate couldn't be created from the storage."+err.Error())
			}

			validated, err = k.ValidateX509Cert(deviceCert, factoryCert)
			if err != nil {
				verificationError = sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenRunner][ValidateX509Cert] failed. Couldn't validate factory certificate."+err.Error())
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
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenRunner][ValidateX509Cert] failed. Device certificate couldn't be verified.")
	}

	//check runner
	var newRunner types.Runner

	_, isFoundAsRunner := k.GetRunnerUsingPubKey(ctx, pubKeyHex)
	_, isFoundAsChallenger := k.GetChallengerUsingPubKey(ctx, pubKeyHex)
	_, isFoundAsClient := k.GetClient(ctx, pubKeyHex)
	if isFoundAsChallenger || isFoundAsRunner || isFoundAsClient {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GetChallengerUsingPubKey][GetRunnerUsingPubKey][GetClient] failed. Runner PubKey is not uniqe OR Runner is already registered.")
	}

	runnerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
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

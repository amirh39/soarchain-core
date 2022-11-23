package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) GenGuard(goCtx context.Context, msg *types.MsgGenGuard) (*types.MsgGenGuardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if guard already exists
	_, isFound := k.GetGuard(ctx, msg.GuardPubKey)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "Guard with given pubKey is already registered in storage.")
	}

	// Check v2x Challenger field
	var newV2XChallenger types.Challenger

	if msg.V2XAddr != "" { // means v2x addr is provided
		// Check if challenger address already exists
		_, isFound := k.GetChallenger(ctx, msg.V2XAddr)
		_, isFoundAsClient := k.GetClient(ctx, msg.V2XAddr)
		_, isFoundAsRunner := k.GetRunner(ctx, msg.V2XAddr)
		if isFound || isFoundAsClient || isFoundAsRunner {
			return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "V2X challenger is already registered in storage.")
		}

		v2xChallengerAddr, err := sdk.AccAddressFromBech32(msg.V2XAddr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid v2x address!")
		}

		// Check v2x stake amount
		v2XStake, _ := sdk.ParseCoinsNormalized(msg.V2XStake)
		requiredStake, _ := sdk.ParseCoinsNormalized("2000soar")
		if v2XStake.GetDenomByIndex(0) != "soar" {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid coin denominator")
		}
		if v2XStake.IsAllLT(requiredStake) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Staked amount: "+v2XStake.String()+" is below the required stake amount "+requiredStake.String())
		}

		newV2XChallenger = types.Challenger{
			Index:        v2xChallengerAddr.String(),
			Address:      v2xChallengerAddr.String(),
			Score:        sdk.ZeroInt().String(),
			StakedAmount: v2XStake.String(),
			NetEarnings:  "",
			Type:         "v2x",
			IpAddr:       msg.V2XIp,
		}

		k.SetChallenger(ctx, newV2XChallenger)

		// Update Challenger Count
		challengerCount, isFound := k.Keeper.GetTotalChallengers(ctx)
		if !isFound {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Challenger count couldn't be fetched!")
		}
		challengerCount.Count++
		k.SetTotalChallengers(ctx, challengerCount)

		// Update ChallengerByIndex
		theCountStr := strconv.FormatUint(challengerCount.Count, 10)
		newIndex := types.ChallengerByIndex{
			Index:      theCountStr,
			Challenger: &newV2XChallenger,
		}
		k.SetChallengerByIndex(ctx, newIndex)

	} else { // v2x address is not provided
		newV2XChallenger = types.Challenger{
			Index:        "",
			Address:      "",
			Score:        sdk.ZeroInt().String(),
			StakedAmount: "",
			NetEarnings:  "",
			Type:         "nil",
			IpAddr:       "",
		}
	}

	// Check v2n Challenger field
	var newV2NChallenger types.Challenger

	if msg.V2NAddr != "" { // means v2n addr is provided
		// Check if challenger already exists
		_, isFound := k.GetChallenger(ctx, msg.V2NAddr)
		_, isFoundAsClient := k.GetClient(ctx, msg.V2NAddr)
		_, isFoundAsRunner := k.GetRunner(ctx, msg.V2NAddr)
		if isFound || isFoundAsClient || isFoundAsRunner {
			return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "V2N challenger is already registered in storage.")
		}

		v2nChallengerAddr, err := sdk.AccAddressFromBech32(msg.V2NAddr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid v2n address!")
		}

		// Check v2n stake amount
		v2NStake, _ := sdk.ParseCoinsNormalized(msg.V2NStake)
		requiredStake, _ := sdk.ParseCoinsNormalized("2000soar")
		if v2NStake.GetDenomByIndex(0) != "soar" {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid coin denominator")
		}
		if v2NStake.IsAllLT(requiredStake) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Staked amount: "+v2NStake.String()+" is below the required stake amount "+requiredStake.String())
		}
		newV2NChallenger = types.Challenger{
			Index:        v2nChallengerAddr.String(),
			Address:      v2nChallengerAddr.String(),
			Score:        sdk.ZeroInt().String(),
			StakedAmount: v2NStake.String(),
			NetEarnings:  "",
			Type:         "v2n",
			IpAddr:       msg.V2NIp,
		}

		k.SetChallenger(ctx, newV2NChallenger)

		// Update Challenger Count
		challengerCount, isFound := k.Keeper.GetTotalChallengers(ctx)
		if !isFound {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Challenger count couldn't be fetched!")
		}
		challengerCount.Count++
		k.SetTotalChallengers(ctx, challengerCount)

		// Update ChallengerByIndex
		theCountStr := strconv.FormatUint(challengerCount.Count, 10)
		newIndex := types.ChallengerByIndex{
			Index:      theCountStr,
			Challenger: &newV2NChallenger,
		}
		k.SetChallengerByIndex(ctx, newIndex)

	} else { // v2n address is not provided
		newV2NChallenger = types.Challenger{
			Index:        "",
			Address:      "",
			Score:        sdk.ZeroInt().String(),
			StakedAmount: "",
			NetEarnings:  "",
			Type:         "nil",
			IpAddr:       "",
		}
	}

	// Check runner
	var newRunner types.Runner
	if msg.RunnerAddr != "" { // means runner addr is provided
		// Check if runner already exists
		_, isFound := k.GetRunner(ctx, msg.RunnerAddr)
		_, isFoundAsChallenger := k.GetChallenger(ctx, msg.RunnerAddr)
		_, isFoundAsClient := k.GetClient(ctx, msg.RunnerAddr)
		if isFound || isFoundAsChallenger || isFoundAsClient {
			return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "Runner is already registered in storage.")
		}

		runnerAddr, err := sdk.AccAddressFromBech32(msg.RunnerAddr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid runner address!")
		}

		// Check runner stake amount
		runnerStake, _ := sdk.ParseCoinsNormalized(msg.RunnerStake)
		requiredStake, _ := sdk.ParseCoinsNormalized("1000soar")
		if runnerStake.GetDenomByIndex(0) != "soar" {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Invalid coin denominator")
		}
		if runnerStake.IsAllLT(requiredStake) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Staked amount: "+runnerStake.String()+" is below the required stake amount "+requiredStake.String())
		}
		newRunner = types.Runner{
			Index:        runnerAddr.String(),
			Address:      runnerAddr.String(),
			Score:        "",
			StakedAmount: runnerStake.String(),
			NetEarnings:  "",
			IpAddr:       msg.RunnerIp,
		}

		k.SetRunner(ctx, newRunner)

		// Update Runner Count
		runnerCount, isFound := k.Keeper.GetTotalRunners(ctx)
		if !isFound {
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "Runner count couldn't be fetched!")
		}
		runnerCount.Count++
		k.SetTotalRunners(ctx, runnerCount)

	} else { // runner address is not provided
		newRunner = types.Runner{
			Index:        "",
			Address:      "",
			Score:        "",
			StakedAmount: "",
			NetEarnings:  "",
			IpAddr:       "",
		}
	}

	newGuard := types.Guard{
		Index:         msg.GuardPubKey,
		GuardId:       "",
		V2XChallenger: &newV2XChallenger,
		V2NChallenger: &newV2NChallenger,
		Runner:        &newRunner,
	}

	k.SetGuard(ctx, newGuard)

	return &types.MsgGenGuardResponse{}, nil
}

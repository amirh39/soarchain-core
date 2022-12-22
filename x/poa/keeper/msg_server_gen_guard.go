package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
)

func (k msgServer) GenGuard(goCtx context.Context, msg *types.MsgGenGuard) (*types.MsgGenGuardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// PubKey check
	if msg.GuardPubKey == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Guard PubKey must be declared in the tx!")
	}

	// Tx field check
	if msg.V2XAddr == "" && msg.V2NAddr == "" && msg.RunnerAddr == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "At least one address field must be provided!")
	}

	// Check if guard already exists
	_, isFound := k.GetGuard(ctx, msg.Creator)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "Guard with given creator address is already registered in storage.")
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
		requiredStake, _ := sdk.ParseCoinsNormalized("2000000000soar")
		v2XStake, err := sdk.ParseCoinsNormalized(msg.V2XStake)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
		}
		if v2XStake.IsAllLT(requiredStake) || !v2XStake.DenomsSubsetOf(requiredStake) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Sent amount: "+v2XStake.String()+" is below the required stake amount "+requiredStake.String())
		}

		// Transfer stakedAmount to contract:
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, v2xChallengerAddr, types.ModuleName, requiredStake)

		//
		newV2XChallenger = types.Challenger{
			Index:        v2xChallengerAddr.String(),
			Address:      v2xChallengerAddr.String(),
			Score:        sdk.NewInt(50).String(), // Base Score
			StakedAmount: v2XStake.String(),
			NetEarnings:  "",
			Type:         "v2x",
			IpAddr:       msg.V2XIp,
		}

		k.SetChallenger(ctx, newV2XChallenger)

	} else { // v2x address is not provided
		newV2XChallenger = types.Challenger{}
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
		requiredStake, _ := sdk.ParseCoinsNormalized("2000000000soar")
		v2NStake, err := sdk.ParseCoinsNormalized(msg.V2NStake)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
		}
		if v2NStake.IsAllLT(requiredStake) || !v2NStake.DenomsSubsetOf(requiredStake) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Sent amount: "+v2NStake.String()+" is below the required stake amount "+requiredStake.String())
		}

		// Transfer stakedAmount to contract:
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, v2nChallengerAddr, types.ModuleName, requiredStake)

		//
		newV2NChallenger = types.Challenger{
			Index:        v2nChallengerAddr.String(),
			Address:      v2nChallengerAddr.String(),
			Score:        sdk.NewInt(50).String(), // Base Score
			StakedAmount: v2NStake.String(),
			NetEarnings:  "",
			Type:         "v2n",
			IpAddr:       msg.V2NIp,
		}

		k.SetChallenger(ctx, newV2NChallenger)

	} else { // v2n address is not provided
		newV2NChallenger = types.Challenger{}
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
		requiredStake, _ := sdk.ParseCoinsNormalized("1000000000soar")
		runnerStake, err := sdk.ParseCoinsNormalized(msg.RunnerStake)

		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
		}
		if runnerStake.IsAllLT(requiredStake) || !runnerStake.DenomsSubsetOf(requiredStake) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Sent amount: "+runnerStake.String()+" is below the required stake amount "+requiredStake.String())
		}

		// Transfer stakedAmount to contract:
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, runnerAddr, types.ModuleName, requiredStake)

		//
		newRunner = types.Runner{
			Index:              runnerAddr.String(),
			Address:            runnerAddr.String(),
			Score:              sdk.NewInt(50).String(), // Base Score
			StakedAmount:       runnerStake.String(),
			NetEarnings:        "",
			IpAddr:             msg.RunnerIp,
			LastTimeChallenged: sdk.ZeroInt().String(),
		}

		k.SetRunner(ctx, newRunner)

	} else { // runner address is not provided
		newRunner = types.Runner{}
	}

	// Set Guard
	newGuard := types.Guard{
		Index:         msg.Creator,
		GuardId:       msg.GuardPubKey,
		V2XChallenger: &newV2XChallenger,
		V2NChallenger: &newV2NChallenger,
		Runner:        &newRunner,
	}

	k.SetGuard(ctx, newGuard)

	return &types.MsgGenGuardResponse{}, nil
}

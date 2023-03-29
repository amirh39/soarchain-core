package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"
)

// for an already registered guard, here is the allowed operations:
// - if address field is empty:
// 		no operation is done
// - if address field is provided:
// 		check if guard already has an entity with some address(challenger or runner)
//			- if exists:
//				check if address fields are the same
//					- if same:
//						update IP if its provided, else do nothing
//					- if not:
//						update address and IP
//			- if does not exist:
//				generate new entity (challenger and runner)

func (k msgServer) UpdateGuard(goCtx context.Context, msg *types.MsgUpdateGuard) (*types.MsgUpdateGuardResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if guard already exists
	guard, isFound := k.GetGuard(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Guard does not exist. Can't update.")
	}

	msgSenderAddress, _ := sdk.AccAddressFromBech32(msg.Creator)

	// UPDATE V2X CHALLENGER
	var newV2XChallenger types.Challenger

	if msg.V2XAddr != "" {
		// Check if challenger address already exists
		_, isFoundClient := k.GetClient(ctx, msg.V2XAddr)
		_, isFoundRunner := k.GetRunner(ctx, msg.V2XAddr)
		if isFoundClient || isFoundRunner {
			return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "V2X challenger is already registered in storage.")
		}

		v2xChallengerAddr, err := sdk.AccAddressFromBech32(msg.V2XAddr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid v2x address!")
		}

		if guard.V2XChallenger.Address != "" { // guard has already registered v2x device

			if msg.V2XAddr == guard.V2XChallenger.Address { // only update the ip
				newV2XChallenger = types.Challenger{
					PubKey:       guard.V2XChallenger.PubKey,
					Address:      guard.V2XChallenger.Address,
					Score:        guard.V2XChallenger.Score,
					StakedAmount: guard.V2XChallenger.StakedAmount,
					NetEarnings:  guard.V2XChallenger.NetEarnings,
					Type:         guard.V2XChallenger.Type,
					IpAddr:       msg.V2XIp,
				}

				k.SetChallenger(ctx, newV2XChallenger)

			} else {
				newV2XChallenger = types.Challenger{ // update address and the ip
					PubKey:       guard.V2XChallenger.PubKey,
					Address:      v2xChallengerAddr.String(),
					Score:        guard.V2XChallenger.Score,
					StakedAmount: guard.V2XChallenger.StakedAmount,
					NetEarnings:  guard.V2XChallenger.NetEarnings,
					Type:         guard.V2XChallenger.Type,
					IpAddr:       msg.V2XIp,
				}

				k.SetChallenger(ctx, newV2XChallenger)
			}

		} else { // create new v2x challenger for guard

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
			transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, requiredStake)
			if transferErr != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake funds couldn't be transferred to POA module!")
			}
			//
			newV2XChallenger = types.Challenger{
				PubKey:       guard.V2XChallenger.PubKey,
				Address:      v2xChallengerAddr.String(),
				Score:        sdk.NewInt(50).String(), // Base Score
				StakedAmount: v2XStake.String(),
				NetEarnings:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
				Type:         "v2x",
				IpAddr:       msg.V2XIp,
			}

			k.SetChallenger(ctx, newV2XChallenger)
		}

	} else { // msg.V2XAddr == nil
		newV2XChallenger = *guard.V2XChallenger
	}

	// UPDATE V2N CHALLENGER
	var newV2NChallenger types.Challenger

	if msg.V2NAddr != "" {
		// Check if challenger address already exists
		_, isFoundClient := k.GetClient(ctx, msg.V2NAddr)
		_, isFoundRunner := k.GetRunner(ctx, msg.V2NAddr)
		if isFoundClient || isFoundRunner {
			return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "V2N challenger is already registered in storage.")
		}

		v2nChallengerAddr, err := sdk.AccAddressFromBech32(msg.V2NAddr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid v2n address!")
		}

		if guard.V2NChallenger.Address != "" {

			if msg.V2NAddr == guard.V2NChallenger.Address {
				newV2NChallenger = types.Challenger{
					PubKey:       guard.V2NChallenger.PubKey,
					Address:      guard.V2NChallenger.Address,
					Score:        guard.V2NChallenger.Score,
					StakedAmount: guard.V2NChallenger.StakedAmount,
					NetEarnings:  guard.V2NChallenger.NetEarnings,
					Type:         guard.V2NChallenger.Type,
					IpAddr:       msg.V2NIp,
				}

				k.SetChallenger(ctx, newV2NChallenger)

			} else {
				newV2NChallenger = types.Challenger{
					PubKey:       guard.V2NChallenger.PubKey,
					Address:      v2nChallengerAddr.String(),
					Score:        guard.V2NChallenger.Score,
					StakedAmount: guard.V2NChallenger.StakedAmount,
					NetEarnings:  guard.V2NChallenger.NetEarnings,
					Type:         guard.V2NChallenger.Type,
					IpAddr:       msg.V2NIp,
				}

				k.SetChallenger(ctx, newV2NChallenger)
			}

		} else { // create new v2n challenger for guard

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
			transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, requiredStake)
			if transferErr != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake funds couldn't be transferred to POA module!")
			}
			//
			newV2NChallenger = types.Challenger{
				PubKey:       newV2NChallenger.PubKey,
				Address:      v2nChallengerAddr.String(),
				Score:        sdk.NewInt(50).String(), // Base Score
				StakedAmount: v2NStake.String(),
				NetEarnings:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
				Type:         "v2n",
				IpAddr:       msg.V2NIp,
			}

			k.SetChallenger(ctx, newV2NChallenger)
		}

	} else { // msg.V2NAddr == nil
		newV2NChallenger = *guard.V2NChallenger
	}

	// UPDATE RUNNER
	var newRunner types.Runner

	if msg.RunnerAddr != "" {
		_, isFoundChallenger := k.GetChallenger(ctx, msg.RunnerAddr)
		_, isFoundClient := k.GetClient(ctx, msg.RunnerAddr)

		if isFoundChallenger || isFoundClient {
			return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "Runner is already registered in storage.")
		}

		runnerAddr, err := sdk.AccAddressFromBech32(msg.RunnerAddr)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid runner address!")
		}

		if guard.Runner.Address != "" {

			if msg.RunnerAddr == guard.Runner.Address {
				newRunner = types.Runner{
					PubKey:             guard.Runner.PubKey,
					Address:            guard.Runner.Address,
					Score:              guard.Runner.Score,
					RewardMultiplier:   guard.Runner.RewardMultiplier,
					StakedAmount:       guard.Runner.StakedAmount,
					NetEarnings:        guard.Runner.NetEarnings,
					IpAddr:             msg.RunnerIp,
					LastTimeChallenged: guard.Runner.LastTimeChallenged,
					CoolDownTolerance:  guard.Runner.CoolDownTolerance,
					GuardAddress:       msg.Creator,
				}

				k.SetRunner(ctx, newRunner)

			} else {
				newRunner = types.Runner{
					PubKey:             guard.Runner.PubKey,
					Address:            runnerAddr.String(),
					Score:              guard.Runner.Score,
					RewardMultiplier:   guard.Runner.RewardMultiplier,
					StakedAmount:       guard.Runner.StakedAmount,
					NetEarnings:        guard.Runner.NetEarnings,
					IpAddr:             msg.RunnerIp,
					LastTimeChallenged: guard.Runner.LastTimeChallenged,
					CoolDownTolerance:  guard.Runner.CoolDownTolerance,
					GuardAddress:       msg.Creator,
				}

				k.SetRunner(ctx, newRunner)
			}

		} else { // create new runner for guard
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
			transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, requiredStake)
			if transferErr != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake funds couldn't be transferred to POA module!")
			}

			// rewardMultiplier
			var initialScore float64 = 50
			rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

			newRunner = types.Runner{
				PubKey:             guard.Runner.PubKey,
				Address:            runnerAddr.String(),
				Score:              strconv.FormatFloat(initialScore, 'f', -1, 64),
				RewardMultiplier:   strconv.FormatFloat(rewardMultiplier, 'f', -1, 64),
				StakedAmount:       runnerStake.String(),
				NetEarnings:        sdk.NewCoin("soar", sdk.ZeroInt()).String(),
				IpAddr:             msg.RunnerIp,
				LastTimeChallenged: ctx.BlockTime().String(),
				CoolDownTolerance:  strconv.FormatUint(1, 10),
				GuardAddress:       msg.Creator,
			}

			k.SetRunner(ctx, newRunner)
		}

	} else {
		newRunner = *guard.Runner
	}

	// Updare Guard
	newGuard := types.Guard{
		Index:         guard.Index,
		GuardId:       guard.GuardId,
		V2XChallenger: &newV2XChallenger,
		V2NChallenger: &newV2NChallenger,
		Runner:        &newRunner,
	}

	k.SetGuard(ctx, newGuard)

	return &types.MsgUpdateGuardResponse{}, nil

}

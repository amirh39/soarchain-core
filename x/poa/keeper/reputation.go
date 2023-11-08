package keeper

import (
	params "soarchain/app/params"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SetReputation(ctx sdk.Context, reputation types.Reputation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	b := k.cdc.MustMarshal(&reputation)
	store.Set(types.ReputationKey(
		reputation.PubKey,
	), b)
}

func (k Keeper) verifyDeviceCertificate(ctx sdk.Context, certificate string) error {
	deviceCert, err := k.CreateX509CertFromString(certificate)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[verifyDeviceCertificate][CreateX509CertFromString] failed. Device certificate and reputation can not be empty.")
	}

	err = k.ValidateCertificate(ctx, deviceCert)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[verifyDeviceCertificate][ValidateCertificate] failed. Device certificate is not valid.")
	}
	return nil
}

func (k Keeper) InitializeReputation(ctx sdk.Context, reputation types.Reputation, certificate string, stakeAmount string, address string) error {

	certificateValidationError := k.verifyDeviceCertificate(ctx, certificate)
	if certificateValidationError != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[InitializeReputation][verifyDeviceCertificate] failed. Couldn't verify device certification.")
	}

	// Check runner stake amount
	senderAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[InitializeReputation][AccAddressFromBech32] failed. Creator address couldn't be parsed.")
	}

	requiredStake := sdk.Coins{sdk.NewInt64Coin(params.BondDenom, 1000000000)}
	parsedStakeAmount, parseError := sdk.ParseCoinsNormalized(stakeAmount)
	if parseError != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[InitializeReputation][ParseCoinsNormalized] failed. Coins couldn't be parsed!")
	}
	if parsedStakeAmount.IsAllLT(requiredStake) || !parsedStakeAmount.DenomsSubsetOf(requiredStake) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Sent amount of runner: "+parsedStakeAmount.String()+" is below the required stake amount "+requiredStake.String())
	}

	//Transfer stakedAmount to poa modules account:
	transferError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddress, types.ModuleName, requiredStake)

	if transferError != nil {
		return sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake(runner) funds couldn't be transferred to POA module!")
	}
	reputation.Address = senderAddress.String()
	reputation.StakedAmount = parsedStakeAmount.String()
	k.SetReputation(ctx, reputation)
	return nil
}

func (k Keeper) InitializeClientReputation(ctx sdk.Context, reputation types.Reputation, certificate string) error {
	certificateValidationError := k.verifyDeviceCertificate(ctx, certificate)
	if certificateValidationError != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "[InitializeReputation][verifyDeviceCertificate] failed. Couldn't verify device certification.")
	}

	k.SetReputation(ctx, reputation)
	return nil
}

func (k Keeper) GetReputation(ctx sdk.Context, pubkey string) (val types.Reputation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))

	b := store.Get(types.ReputationKey(
		pubkey,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllReputation(ctx sdk.Context) (list []types.Reputation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reputation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetReputationsByAddress(ctx sdk.Context, address string) (val types.Reputation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reputation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Address == address {
			return val, true
		}
	}
	return types.Reputation{}, false
}

func (k Keeper) GetReputationsByAddressAndType(ctx sdk.Context, address, reputationType string) (val types.Reputation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var reputation types.Reputation
		k.cdc.MustUnmarshal(iterator.Value(), &reputation)
		if reputation.Address == address && reputation.Type == reputationType {
			return reputation, true
		}
	}
	return types.Reputation{}, false
}

func (k Keeper) GetAllChallenger(ctx sdk.Context) (list []types.Reputation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reputation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Type == constants.V2XChallenger || val.Type == constants.V2NChallenger {
			list = append(list, val)
		}
	}
	return
}

func (k Keeper) GetAllRunner(ctx sdk.Context) (list []types.Reputation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Reputation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Type == "" {
			list = append(list, val)
		}
	}
	return
}

func (k Keeper) RemoveClientReputation(ctx sdk.Context, creator string) error {

	reputation, found := k.GetReputationsByAddress(ctx, creator)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RemoveClientReputation][GetReputationsByAddress] failed. Reputation not found. Make sure using valid address.")
	}
	if reputation.Type == constants.Mini || reputation.Type == constants.Pro {

		if reputation.Address != creator {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[RemoveClientReputation] failed. Registrant is not recognized.")
		}

		// Transfer claimmable rewards
		earnedAmount, err := sdk.ParseCoinsNormalized(reputation.NetEarnings)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[RemoveClientReputation][ParseCoinsNormalized] failed. Withdraw amount couldn't be parsed.")
		}

		clientAccount, parseError := sdk.AccAddressFromBech32(reputation.Address)
		if parseError != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[RemoveClientReputation][AccAddressFromBech32] failed. Couldn't create acc account from the address.")
		}

		errTransfer := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, clientAccount, earnedAmount)
		if errTransfer != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[RemoveClientReputation][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
		}

		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
		store.Delete(types.ReputationKey(
			reputation.PubKey,
		))
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RemoveClientReputation] failed. Couldn't delete client reputation. Make sure using valid address.")
}

func (k Keeper) RemoveRunnerReputation(ctx sdk.Context, creator string) error {

	reputation, found := k.GetReputationsByAddress(ctx, creator)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RemoveRunnerReputation][GetReputationsByAddress] failed. Reputation not found. Make sure using valid address.")
	}

	if reputation.Type == "" {
		if reputation.Address != creator {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[RemoveRunnerReputation][Validate address] failed. Runner reputation address is not equal with the given address.")
		}
		receipientAddress, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[RemoveRunnerReputation][AccAddressFromBech32] failed. Sender address is not valid.")
		}
		stakedAmountStr := reputation.StakedAmount
		stakedAmount, parseError := sdk.ParseCoinsNormalized(stakedAmountStr)
		if parseError != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[RemoveRunnerReputation][ParseCoinsNormalized] failed. Stake amount is not valid.")
		}
		transferErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receipientAddress, stakedAmount)
		if transferErr != nil {
			return sdkerrors.Wrap(sdkerrors.ErrPanic, "[RemoveRunnerReputation][SendCoinsFromModuleToAccount] failed. Couldn't send coins. ErrorMessage: "+transferErr.Error())
		}
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
		store.Delete(types.ReputationKey(
			reputation.PubKey,
		))
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RemoveRunnerReputation] failed. Couldn't delete runner reputation. Make sure using valid address.")
}

func (k Keeper) RemoveChallengerReputation(ctx sdk.Context, creator string) error {

	reputation, found := k.GetReputationsByAddress(ctx, creator)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RemoveChallengerReputation][GetReputationsByAddress] failed. Challenger reputation not found. Make sure using valid address.")
	}

	if reputation.Type == constants.V2XChallenger || reputation.Type == constants.V2NChallenger {

		if reputation.Address != creator {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "[RemoveChallengerReputation][Validate address] failed. Challenger reputation address is not equal with the given address.")
		}

		recipientAddress, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[RemoveChallengerReputation][AccAddressFromBech32] failed. Sender addres is not valid."+err.Error())
		}

		// Query the staked amount and refund
		stakedAmountStr := reputation.StakedAmount
		stakedAmount, parseError := sdk.ParseCoinsNormalized(stakedAmountStr)
		if parseError != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "[RemoveChallengerReputation][ParseCoinsNormalized] failed. Couldn't parse the list if coins."+err.Error())
		}

		transferError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddress, stakedAmount)
		if transferError != nil {
			return sdkerrors.Wrap(sdkerrors.ErrPanic, "[RemoveChallengerReputation][SendCoinsFromModuleToAccount] failed. Couldn't send coins.")
		}
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReputationKeyPrefix))
		store.Delete(types.ReputationKey(
			reputation.PubKey,
		))
		return nil
	}
	return sdkerrors.Wrap(sdkerrors.ErrNotFound, "[RemoveChallengerReputation] failed. Couldn't delete challenger reputation. Make sure using valid address.")
}

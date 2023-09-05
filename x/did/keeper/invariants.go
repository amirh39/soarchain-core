package keeper

import (
	"fmt"
	"soarchain/x/did/errors"
	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// RegisterInvariants registers all staking invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper, bk types.BankKeeper) {
	ir.RegisterRoute(types.ModuleName, "nonnegative-outstanding", NonnegativeBalanceInvariant(k))
	ir.RegisterRoute(types.ModuleName, "total-supply", TotalSupply(k))
	//TODO: Uncomment after increasing SDK version
	// ir.RegisterRoute(types.ModuleName, "module-account", ModuleAccountInvariant(k, bk))

}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper, bk types.BankKeeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := NonnegativeBalanceInvariant(k)(ctx)
		if stop {
			return res, stop
		}

		res, stop = TotalSupply(k)(ctx)
		if stop {
			return res, stop
		}

		//TODO: uncomment after increaisng SDK version
		// res, stop = ModuleAccountInvariant(k, bk)(ctx)
		// if stop {
		// 	return res, stop
		// }

		return res, stop
	}
}

// // NonnegativeBalanceInvariant checks that all accounts in the application have non-negative balances
func NonnegativeBalanceInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			msg   string
			count int
		)

		k.IterateAllBalances(k, ctx, func(addr sdk.AccAddress, balance sdk.Coin) bool {
			if balance.IsNegative() {
				count++
				msg += fmt.Sprintf("\t%s has a negative balance of %s\n", addr, balance)
			}

			return false
		})

		broken := count != 0

		return sdk.FormatInvariant(
			types.ModuleName, "nonnegative-outstanding",
			fmt.Sprintf("amount of negative balances found %d\n%s", count, msg),
		), broken
	}
}

// TotalSupply checks that the total supply reflects all the coins held in accounts
func TotalSupply(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		expectedTotal := sdk.Coins{}
		supply, _, err := k.GetPaginatedTotalSupply(k, ctx, &query.PageRequest{Limit: query.MaxLimit})
		if err != nil {
			return sdk.FormatInvariant(types.ModuleName, "query supply",
				fmt.Sprintf("error querying total supply %v", err)), false
		}

		k.IterateAllBalances(k, ctx, func(_ sdk.AccAddress, balance sdk.Coin) bool {
			expectedTotal = expectedTotal.Add(balance)
			return false
		})

		broken := !expectedTotal.IsEqual(supply)

		return sdk.FormatInvariant(types.ModuleName, "total supply",
			fmt.Sprintf(
				"\tsum of accounts coins: %v\n"+
					"\tsupply.Total:          %v\n",
				expectedTotal, supply)), broken
	}
}

// ModuleAccountInvariant checks that the module account coins reflects the sum of
// deposit amounts held on store.
//TODO: uncomment after increasing SDK version
// func ModuleAccountInvariant(k Keeper, bk types.BankKeeper) sdk.Invariant {
// 	return func(ctx sdk.Context) (string, bool) {
// 		var expectedDeposits sdk.Coins

// 		macc := k.GetGovernanceAccount(ctx)
// 		balances := bk.GetAllBalances(ctx, macc.GetAddress())

// 		// Require that the deposit balances are <= than the x/gov module's total
// 		// balances. External funds can be sent to x/gov
// 		// module's account and so the balance can be larger.
// 		broken := !balances.IsAllGTE(expectedDeposits)

// 		fmt.Printf("\tbroken broken: %v", broken)

// 		//TODO: using v1 after increasing SDK version
// 		return sdk.FormatInvariant(types.ModuleName, "deposits",
// 			fmt.Sprintf(
// 				"\tgov ModuleAccount coins: %v\n"+
// 					"\tsum of deposit amounts:  %v\n",
// 				balances, expectedDeposits)), false
// 	}
// }

func (keeper Keeper) IterateAllBalances(k Keeper, ctx sdk.Context, cb func(sdk.AccAddress, sdk.Coin) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	balancesStore := prefix.NewStore(store, types.BalancesPrefix)

	iterator := balancesStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		address, denom, err := AddressAndDenomFromBalancesStore(iterator.Key())
		if err != nil {
			sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
			return
		}

		balance, err := UnmarshalBalanceCompat(keeper.cdc, iterator.Value(), denom)
		if err != nil {
			sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "[UnmarshalBalanceCompat] failed. invalid denom.")
			return
		}

		if cb(address, balance) {
			break
		}
	}
}

func AddressAndDenomFromBalancesStore(key []byte) (sdk.AccAddress, string, error) {
	if len(key) == 0 {
		return nil, "", sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
	}

	kv.AssertKeyAtLeastLength(key, 1)

	addrBound := int(key[0])

	if len(key)-1 < addrBound {
		return nil, "", sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
	}

	return key[1 : addrBound+1], string(key[addrBound+1:]), nil
}

func UnmarshalBalanceCompat(cdc codec.BinaryCodec, bz []byte, denom string) (sdk.Coin, error) {
	if err := sdk.ValidateDenom(denom); err != nil {
		return sdk.Coin{}, err
	}

	amount := sdk.ZeroInt()

	if bz == nil {
		return sdk.NewCoin(denom, amount), nil
	}

	if err := amount.Unmarshal(bz); err != nil {
		// try to unmarshal with the legacy format.
		var balance sdk.Coin
		if cdc.Unmarshal(bz, &balance) != nil {
			// return with the original error
			return sdk.Coin{}, err
		}
		return balance, nil
	}

	return sdk.NewCoin(denom, amount), nil
}

// IterateAllDeposits iterates over all the stored deposits and performs a callback function.
func (keeper Keeper) IterateAllDeposits(ctx sdk.Context, cb func() (stop bool)) {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DepositsKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {

		if !cb() {
			break
		}
	}
}

// GetPaginatedTotalSupply queries for the supply, ignoring 0 coins, with a given pagination
func (keeper Keeper) GetPaginatedTotalSupply(k Keeper, ctx sdk.Context, pagination *query.PageRequest) (sdk.Coins, *query.PageResponse, error) {
	store := ctx.KVStore(k.storeKey)
	var SupplyKey = []byte{0x00}
	supplyStore := prefix.NewStore(store, SupplyKey)

	supply := sdk.NewCoins()

	pageRes, err := query.Paginate(supplyStore, pagination, func(key, value []byte) error {
		amount := sdk.ZeroInt()
		err := amount.Unmarshal(value)
		if err != nil {
			return fmt.Errorf("unable to convert amount string to Int %v", err)
		}

		// `Add` omits the 0 coins addition to the `supply`.
		supply = supply.Add(sdk.NewCoin(string(key), amount))
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return supply, pageRes, nil
}

// GetGovernanceAccount returns the governance ModuleAccount
//TODO: uncomment after increasing SDK version
// func (keeper Keeper) GetGovernanceAccount(ctx sdk.Context) authtypes.ModuleAccountI {
// 	return keeper.authKeeper.GetModuleAccount(ctx, types.ModuleName)
// }

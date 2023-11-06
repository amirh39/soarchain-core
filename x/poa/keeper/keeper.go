package keeper

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"soarchain/x/poa/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		epochKeeper   types.EpochKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	epochKeeper types.EpochKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		epochKeeper:   epochKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	printLogs, err := strconv.ParseBool(os.Getenv("PrintLogs"))
	if err != nil {
		fmt.Print("[keeper][Logger] failed. Couldn't parse int to string.")
	}

	if !printLogs {
		return nil
	}
	return ctx.Logger().With(
		"timestamp", time.Now().String(),
		"module", fmt.Sprintf("x/%s", types.ModuleName),
		"height", strconv.FormatInt(ctx.BlockHeight(), 10),
	)
}

package keeper

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"soarchain/x/did/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramStore paramtypes.Subspace

		poaKeeper types.PoaKeeper
	}
)

func NewDidKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	poaKeeper types.PoaKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramStore: ps,
		poaKeeper:  poaKeeper,
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

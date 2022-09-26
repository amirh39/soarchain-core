package keeper

import (
	"soarchain/x/poa/types"
)

var _ types.QueryServer = Keeper{}

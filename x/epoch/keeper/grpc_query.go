package keeper

import (
	"soarchain/x/epoch/types"
)

var _ types.QueryServer = Keeper{}

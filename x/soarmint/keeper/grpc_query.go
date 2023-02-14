package keeper

import (
	"soarchain/x/soarmint/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"soarchain/x/dpr/types"
)

var _ types.QueryServer = Keeper{}

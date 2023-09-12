package keeper

import (
	"soarchain/x/did/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/amirh39/soarchain-core/x/soarmint/types"
)

var _ types.QueryServer = Keeper{}

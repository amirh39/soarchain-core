package keeper

import (
	"github.com/amirh39/soarchain-core/x/epoch/types"
)

var _ types.QueryServer = Keeper{}

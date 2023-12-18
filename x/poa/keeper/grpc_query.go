package keeper

import (
	"github.com/amirh39/soarchain-core/x/poa/types"
)

var _ types.QueryServer = Keeper{}

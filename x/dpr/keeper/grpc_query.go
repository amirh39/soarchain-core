package keeper

import (
	"github.com/amirh39/soarchain-core/x/dpr/types"
)

var _ types.QueryServer = Keeper{}

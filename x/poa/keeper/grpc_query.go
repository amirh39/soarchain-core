package keeper

import (
	"github.com/soar-robotics/soarchain-core/x/poa/types"
)

var _ types.QueryServer = Keeper{}

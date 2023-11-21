package keeper

import (
	"github.com/soar-robotics/soarchain-core/x/epoch/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/soar-robotics/soarchain-core/x/soarmint/types"
)

var _ types.QueryServer = Keeper{}

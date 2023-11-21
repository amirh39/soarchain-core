package keeper

import (
	"github.com/soar-robotics/soarchain-core/x/dpr/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/soar-robotics/soarchain-core/x/did/types"
)

var _ types.QueryServer = Keeper{}

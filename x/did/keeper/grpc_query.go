package keeper

import (
	"github.com/amirh39/soarchain-core/x/did/types"
)

var _ types.QueryServer = Keeper{}

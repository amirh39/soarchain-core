package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"testing"
)

func Test_DeactivateDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)

	dprs := SetupNDpr(1)
	dprsDif := SetupNDifDpr(1)
	deactiveDprs := SetupNDeactiveDpr(1)
	keeper.SetDpr(ctx, dprs[0])
	keeper.SetDpr(ctx, dprsDif[0])
	keeper.SetDpr(ctx, deactiveDprs[0])

	// allDprs := keeper.GetAllDpr(ctx)
	// t.Log("0000000000000", allDprs)

	keeper.DeactivateDpr(ctx, 4)

	// allDprs1 := keeper.GetAllDpr(ctx)
	// t.Log("1111111111111111", allDprs1)

}

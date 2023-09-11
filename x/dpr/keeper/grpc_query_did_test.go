package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
)

func Test_DrpDocumentGet(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)
	t.Log(keeper)
	t.Log(ctx)
}

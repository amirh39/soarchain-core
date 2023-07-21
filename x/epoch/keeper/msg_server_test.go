package keeper_test

import (
	"context"
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/x/epoch/keeper"
	"soarchain/x/epoch/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EpochKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

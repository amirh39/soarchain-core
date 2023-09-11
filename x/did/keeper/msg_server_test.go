package keeper_test

import (
	"context"
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/x/did/keeper"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DidKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/testutil"
	"soarchain/x/poa/types"

	"github.com/golang/mock/gomock"
)

const (
	communityWallet = testutil.CommunityWallet
)

func setupMsgServerClaimMotusRewards(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.PoaKeeperWithMocks(t, bankMock)
	poa.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

func TestClaimMotusRewards(t *testing.T) {
	msgServer, k, context, ctrl, bank := setupMsgServerClaimMotusRewards(t)
	defer ctrl.Finish()

	bank.ExpectAny(context)

	ctx := sdk.UnwrapSDKContext(context)
	// create test motus wallet and client
	motusWallet := types.MotusWallet{
		Index:  communityWallet,
		Client: &types.Client{Address: communityWallet, NetEarnings: "100soar"},
	}
	k.SetMotusWallet(ctx, motusWallet)

	// test with successful claim
	resp, err := msgServer.ClaimMotusRewards(context, &types.MsgClaimMotusRewards{
		Creator: communityWallet,
		Amount:  "100soar",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	// // test with insufficient funds
	// msg.Amount = "50soar"
	// resp, err := keeper.ClaimMotusRewards(context.Background(), msg)
	// require.NoError(t, err)
	// require.NotNil(t, resp)

	// // verify client balance and net earnings
	// client, _ := keeper.GetClient(ctx, addr.String())
	// require.Equal(t, "50soar", client.NetEarnings)
	// balance := keeper.bankKeeper.GetCoins(ctx, sdk.AccAddress(addr))
	// require.Equal(t, sdk.NewCoin("soar", sdk.NewInt(50)), balance[0])
}

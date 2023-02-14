package keeper_test

import (
	"context"
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/testutil"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/require"
)

const (
	strategicWallet = testutil.StrategicWallet
	v2xDevice1      = testutil.V2xDevice1
	v2nDevice1      = testutil.V2nDevice1
	runnerDevice1   = testutil.RunnerDevice1
)

func setupMsgServerGenGuard(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.PoaKeeperWithMocks(t, bankMock)
	poa.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

func TestGenGuard(t *testing.T) {
	msgServer, _, context, ctrl, bank := setupMsgServerGenGuard(t)
	defer ctrl.Finish()
	// defer func() {
	// 	r := recover()
	// 	require.NotNil(t, r, "The code did not panic")
	// }()
	bank.ExpectAny(context)
	// msgSenderAddress, _ := sdk.AccAddressFromBech32(strategicWallet)

	// bank.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), msgSenderAddress, types.ModuleName, sdk.Coins{sdk.NewInt64Coin("soar", 2000000000)})

	_, err := msgServer.GenGuard(context, &types.MsgGenGuard{
		Creator:     strategicWallet,
		GuardPubKey: "A6VReuNYHM8ot0kSqmjTpzNEJErsizE7i1XOIbUxHCph",
		V2XAddr:     v2xDevice1,
		V2XStake:    "2000000000soar",
		V2XIp:       "104.248.142.45 ",
		V2NAddr:     v2nDevice1,
		V2NStake:    "2000000000soar",
		V2NIp:       "104.248.142.45 ",
		RunnerAddr:  runnerDevice1,
		RunnerStake: "1000000000soar",
		RunnerIp:    "104.248.142.45 ",
	})

	require.Nil(t, err)

}

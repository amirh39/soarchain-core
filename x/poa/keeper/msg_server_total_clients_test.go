package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "soarchain/testutil/keeper"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
)

func TestTotalClientsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.PoaKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateTotalClients{Creator: creator}
	_, err := srv.CreateTotalClients(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetTotalClients(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestTotalClientsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTotalClients
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTotalClients{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTotalClients{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PoaKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateTotalClients{Creator: creator}
			_, err := srv.CreateTotalClients(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateTotalClients(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTotalClients(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestTotalClientsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTotalClients
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTotalClients{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTotalClients{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PoaKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateTotalClients(wctx, &types.MsgCreateTotalClients{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTotalClients(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTotalClients(ctx)
				require.False(t, found)
			}
		})
	}
}

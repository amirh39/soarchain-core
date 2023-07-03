package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"soarchain/x/poa/types"
)

func createNMotusWallet(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MotusWallet {
	items := make([]types.MotusWallet, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetMotusWallet(ctx, items[i])
	}
	return items
}

func Test_MotusWalletQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNMotusWallet(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMotusWalletRequest
		response *types.QueryGetMotusWalletResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMotusWalletRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetMotusWalletResponse{MotusWallet: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMotusWalletRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetMotusWalletResponse{MotusWallet: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMotusWalletRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "Key not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MotusWallet(wctx, tc.request)
			if err != nil {
				require.ErrorIs(t, err, err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func Test_MotusWalletQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNMotusWallet(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMotusWalletRequest {
		return &types.QueryAllMotusWalletRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MotusWalletAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MotusWallet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MotusWallet),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MotusWalletAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MotusWallet), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MotusWallet),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MotusWalletAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MotusWallet),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MotusWalletAll(wctx, nil)
		require.Error(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

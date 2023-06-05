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
	"soarchain/x/poa/types"
)

func Test_RunnerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNRunner(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRunnerRequest
		response *types.QueryGetRunnerResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRunnerRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetRunnerResponse{Runner: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRunnerRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetRunnerResponse{Runner: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRunnerRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Runner(wctx, tc.request)
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

func Test_RunnerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreatesRunnerForPagination(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRunnerRequest {
		return &types.QueryAllRunnerRequest{
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
			resp, err := keeper.RunnerAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Runner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Runner),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RunnerAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Runner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Runner),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RunnerAll(wctx, request(nil, 0, 0, true))
		require.Nil(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Runner),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RunnerAll(wctx, nil)
		require.Error(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

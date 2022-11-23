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

// Prevent strconv unused error
var _ = strconv.IntSize

func TestChallengerByIndexQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChallengerByIndex(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerByIndexRequest
		response *types.QueryGetChallengerByIndexResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetChallengerByIndexRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetChallengerByIndexResponse{ChallengerByIndex: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetChallengerByIndexRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetChallengerByIndexResponse{ChallengerByIndex: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetChallengerByIndexRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ChallengerByIndex(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
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

func TestChallengerByIndexQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChallengerByIndex(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllChallengerByIndexRequest {
		return &types.QueryAllChallengerByIndexRequest{
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
			resp, err := keeper.ChallengerByIndexAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChallengerByIndex), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChallengerByIndex),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ChallengerByIndexAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChallengerByIndex), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChallengerByIndex),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ChallengerByIndexAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ChallengerByIndex),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ChallengerByIndexAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

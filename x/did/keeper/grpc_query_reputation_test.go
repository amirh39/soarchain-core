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
	"soarchain/x/did/types"
)

func Test_ClientQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNClient(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetReputationRequest
		response *types.QueryGetReputationResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetReputationRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetReputationResponse{Reputation: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetReputationRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetReputationResponse{Reputation: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetReputationRequest{
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
			response, err := keeper.Reputation(wctx, tc.request)
			if err != nil {
				require.Error(t, err)
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

func Test_ClientQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNClient(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllReputationRequest {
		return &types.QueryAllReputationRequest{
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
			resp, err := keeper.ReputationAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Reputation), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Reputation),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ReputationAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Reputation), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Reputation),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ReputationAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Reputation),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ReputationAll(wctx, nil)
		require.Error(t, err)
	})
}

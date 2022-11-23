package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/types"
)

func TestTotalRunnersQuery(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTotalRunners(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTotalRunnersRequest
		response *types.QueryGetTotalRunnersResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTotalRunnersRequest{},
			response: &types.QueryGetTotalRunnersResponse{TotalRunners: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TotalRunners(wctx, tc.request)
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

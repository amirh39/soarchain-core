package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/epoch/types"
)

func Test_EpochDataQuery(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := CreateEpochData(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetEpochDataRequest
		response *types.QueryGetEpochDataResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetEpochDataRequest{},
			response: &types.QueryGetEpochDataResponse{EpochData: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EpochData(wctx, tc.request)
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

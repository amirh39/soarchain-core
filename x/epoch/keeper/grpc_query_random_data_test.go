package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/testutil/nullify"
	"github.com/amirh39/soarchain-core/x/epoch/types"
)

func Test_RandomDataQuery(t *testing.T) {
	keeper, ctx := keepertest.EpochKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := CreateRandomData(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRandomDataRequest
		response *types.QueryGetRandomDataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRandomDataRequest{
				EpochNumber: "353545345",
			},
			response: &types.QueryGetRandomDataResponse{RandomData: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RandomData(wctx, tc.request)
			t.Log("response --->", response)
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

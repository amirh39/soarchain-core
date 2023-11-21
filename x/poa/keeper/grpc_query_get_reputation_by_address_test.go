package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/soar-robotics/soarchain-core/testutil/keeper"
	"github.com/soar-robotics/soarchain-core/testutil/nullify"
	"github.com/soar-robotics/soarchain-core/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_GetReputationByAddress(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	reputations := CreateNReputation(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetReputationByAddressRequest
		response *types.QueryGetReputationByAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetReputationByAddressRequest{
				Address: reputations[0].Address,
			},
			response: &types.QueryGetReputationByAddressResponse{Reputation: &reputations[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetReputationByAddressRequest{
				Address: reputations[1].Address,
			},
			response: &types.QueryGetReputationByAddressResponse{Reputation: &reputations[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetReputationByAddressRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "client not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetReputationByAddress(wctx, tc.request)
			if tc.err != nil {
				require.Error(t, tc.err)
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

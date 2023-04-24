package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/types"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGetClientByAddress(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNClient(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetClientByAddressRequest
		response *types.QueryGetClientByAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetClientByAddressRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetClientByAddressResponse{Client: &msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetClientByAddressRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetClientByAddressResponse{Client: &msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetClientByAddressRequest{
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
			response, err := keeper.GetClientByAddress(wctx, tc.request)
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
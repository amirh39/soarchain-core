package keeper_test

import (
	"strconv"
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_GetDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := SetupNDpr(2)
	keeper.SetDpr(ctx, msgs[0])
	keeper.SetDpr(ctx, msgs[1])

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDprRequest
		response *types.QueryGetDprResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDprRequest{
				Id: msgs[0].Id,
			},
			response: &types.QueryGetDprResponse{Dpr: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDprRequest{
				Id: msgs[1].Id,
			},
			response: &types.QueryGetDprResponse{Dpr: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDprRequest{
				Id: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Dpr(wctx, tc.request)
			if err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, response)
			}
		})
	}
}

func Test_GetAllDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)
	msgs := SetupNDpr(2)
	keeper.SetDpr(ctx, msgs[0])
	keeper.SetDpr(ctx, msgs[1])

	allDprs, _ := keeper.GetAllDpr(ctx)
	require.NotNil(t, allDprs)

}

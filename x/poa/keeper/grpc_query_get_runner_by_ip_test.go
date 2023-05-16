
package keeper_test

import (
	"testing"

	"soarchain/x/poa/types"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/keeper"
	"strconv"

	keepertest "soarchain/testutil/keeper"
	"github.com/stretchr/testify/require"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createNRunnerByIPAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Runner {
	items := make([]types.Runner, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].IpAddr = "45.12.65.78"

		keeper.SetRunner(ctx, items[i])
	}
	return items
}

/** The IP address which is received by the request is a valid IP address. 
Response should return a valid challenger which is related to that address.*/
func Test_GetRunnerByIp(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRunnerByIPAddress(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRunnerByIpRequest
		response *types.QueryGetRunnerByIpResponse
		err      error
	}{
		{
			desc: "Valid IP Address",
			request: &types.QueryGetRunnerByIpRequest{
				IpAddress: msgs[0].IpAddr,
			},
			response: &types.QueryGetRunnerByIpResponse{Runner: &msgs[0]},
		},
	}{
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetRunnerByIp(wctx, tc.request)
			// t.Log("Test_GetRunnerByIp_Log", response)
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

/** The IP address which is received by the request is invalid. 
The error & error message which is created into the code, will raise.*/
func Test_GetRunnerByNotValidIp(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := createNRunnerByIPAddress(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRunnerByIpRequest
		response *types.QueryGetRunnerByIpResponse
		err      error
	}{
        {
			desc: "Invalid IP Address",
			request: &types.QueryGetRunnerByIpRequest{
				IpAddress: "0xc000db6e50",
			},
			response: &types.QueryGetRunnerByIpResponse{Runner: &items[0]},
		},
		{
			desc: "Empty IP Address",
			request: &types.QueryGetRunnerByIpRequest{
				IpAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "Runner not found"),
		},
		{
			desc: "Invalid Request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}{
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetRunnerByIp(wctx, tc.request)
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
	
package keeper_test

import (
	"testing"

	"soarchain/testutil/nullify"
	"soarchain/x/poa/types"
	"strconv"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/** The address which is received by the request is a valid address.
Response should return a valid challenger which is related to that address.*/
func Test_GetChallengerByAddress(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNChallenger(keeper, ctx, 2)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerByAddressRequest
		response *types.QueryGetChallengerByAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetChallengerByAddressRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetChallengerByAddressResponse{Challenger: &msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetChallengerByAddressRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetChallengerByAddressResponse{Challenger: &msgs[1]},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetChallengerByAddress(wctx, tc.request)
			if err != nil {
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

/** The address which is received by the request is invalid.
Response should return an error and error message which is created into the code will rise.*/
func Test_GetChallengerByNotValidAddress(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := CreateNChallenger(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerByAddressRequest
		response *types.QueryGetChallengerByAddressResponse
		err      error
	}{
		{
			desc: "Invalid Address",
			request: &types.QueryGetChallengerByAddressRequest{
				Address: "0xc000db6e50",
			},
			response: &types.QueryGetChallengerByAddressResponse{Challenger: &items[0]},
		},
		{
			desc: "Address Not Found",
			request: &types.QueryGetChallengerByAddressRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "challenger not found"),
		},
		{
			desc: "Invalid Request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetChallengerByAddress(wctx, tc.request)
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

package keeper_test

import (
	"testing"

	"strconv"

	"github.com/amirh39/soarchain-core/testutil/nullify"
	"github.com/amirh39/soarchain-core/x/did/types"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
* The address which is received by the request is a valid address.
Response should return a valid challenger which is related to that address.
*/
func Test_GetChallengerByAddress(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := CreateNChallengerDid(keeper, ctx, 2)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerDidByAddressRequest
		response *types.QueryGetChallengerDidByAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetChallengerDidByAddressRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetChallengerDidByAddressResponse{ChallengerDid: &msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetChallengerDidByAddressRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetChallengerDidByAddressResponse{ChallengerDid: &msgs[1]},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetChallengerDidByAddress(wctx, tc.request)
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

/*
* The address which is received by the request is invalid.
Response should return an error and error message which is created into the code will rise.
*/
func Test_GetChallengerByNotValidAddress(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := CreateNChallengerDid(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerDidByAddressRequest
		response *types.QueryGetChallengerDidByAddressResponse
		err      error
	}{
		{
			desc: "Invalid Address",
			request: &types.QueryGetChallengerDidByAddressRequest{
				Address: "0xc000db6e50",
			},
			response: &types.QueryGetChallengerDidByAddressResponse{ChallengerDid: &items[0]},
		},
		{
			desc: "Address Not Found",
			request: &types.QueryGetChallengerDidByAddressRequest{
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
			response, err := keeper.GetChallengerDidByAddress(wctx, tc.request)
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

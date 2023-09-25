package keeper_test

import (
	"testing"
)

func Test_GetReputationByAddress(t *testing.T) {
	// keeper, ctx := keepertest.PoaKeeper(t)
	// wctx := sdk.WrapSDKContext(ctx)
	// msgs := CreateNClient(keeper, ctx, 2)
	// for _, tc := range []struct {
	// 	desc     string
	// 	request  *types.QueryGetReputationByAddressRequest
	// 	response *types.QueryGetReputationByAddressResponse
	// 	err      error
	// }{
	// 	{
	// 		desc: "First",
	// 		request: &types.QueryGetReputationByAddressRequest{
	// 			Address: msgs[0].Address,
	// 		},
	// 		response: &types.QueryGetReputationByAddressResponse{Reputation: &msgs[0]},
	// 	},
	// 	{
	// 		desc: "Second",
	// 		request: &types.QueryGetReputationByAddressRequest{
	// 			Address: msgs[1].Address,
	// 		},
	// 		response: &types.QueryGetReputationByAddressResponse{Reputation: &msgs[1]},
	// 	},
	// 	{
	// 		desc: "KeyNotFound",
	// 		request: &types.QueryGetReputationByAddressRequest{
	// 			Address: strconv.Itoa(100000),
	// 		},
	// 		err: status.Error(codes.NotFound, "client not found"),
	// 	},
	// 	{
	// 		desc: "InvalidRequest",
	// 		err:  status.Error(codes.InvalidArgument, "invalid request"),
	// 	},
	// } {
	// 	t.Run(tc.desc, func(t *testing.T) {
	// 		response, err := keeper.GetReputationByAddress(wctx, tc.request)
	// 		if tc.err != nil {
	// 			require.Error(t, tc.err)
	// 		} else {
	// 			require.NoError(t, err)
	// 			require.Equal(t,
	// 				nullify.Fill(tc.response),
	// 				nullify.Fill(response),
	// 			)
	// 		}
	// 	})
	// }
}

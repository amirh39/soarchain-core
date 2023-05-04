
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
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)


func createNChallengerByPubkey(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Challenger {
	items := make([]types.Challenger, 1)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].PubKey = "3056301006072a8648ce3d020106052b8104000a03420004c4039cc2459a57357707620ddbbaddfeda5d4c66cc9ac9c3aac997e65f16b78253b3f9241182014246c1b945595c1ed2463e22ca59f153a74fee375e23a86561"
		keeper.SetChallenger(ctx, items[i])
	}
	return items
}

/** The public key which is received by the request is a valid Pubkey. Response should return a valid challenger
	which is related to that public key.*/
func Test_GetChallengerByPubKey(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChallengerByPubkey(keeper, ctx, 1)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChallengerByPubKeyRequest
		response *types.QueryGetChallengerByPubKeyResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetChallengerByPubKeyRequest{
				PubKey: "3056301006072a8648ce3d020106052b8104000a03420004c4039cc2459a57357707620ddbbaddfeda5d4c66cc9ac9c3aac997e65f16b78253b3f9241182014246c1b945595c1ed2463e22ca59f153a74fee375e23a86561", //msgs[0].PubKey,
			},
			response: &types.QueryGetChallengerByPubKeyResponse{Challenger: &msgs[0]},
		},
		// {
		// 	desc: "Second",
		// 	request: &types.QueryGetChallengerByPubKeyRequest{
		// 		PubKey: msgs[1].PubKey,
		// 	},
		// 	response: &types.QueryGetChallengerByPubKeyResponse{Challenger: &msgs[1]},
		// },
	}{
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetChallengerByPubKey(wctx, tc.request)

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

/** The public key which is received by the request is invalid. Response should return an error and error 
	messages which are created into the code should be raised.*/
// func Test_GetChallengerByPubKey(t *testing.T) {
// 	keeper, ctx := keepertest.PoaKeeper(t)
// 	wctx := sdk.WrapSDKContext(ctx)
// 	items := createNChallenger(keeper, ctx, 1)

// 	for _, tc := range []struct {
// 		desc     string
// 		request  *types.QueryGetChallengerByPubKeyRequest
// 		response *types.QueryGetChallengerByPubKeyResponse
// 		err      error
// 	}{
//         {
// 			desc: "Invalid Pubkey",
// 			request: &types.QueryGetChallengerByPubKeyRequest{
// 				PubKey: "0xc000db6e50",
// 			},
// 			response: &types.QueryGetChallengerByPubKeyResponse{Challenger: &items[0]},
// 		},
// 		{
// 			desc: "Pubkey Not Found",
// 			request: &types.QueryGetChallengerByPubKeyRequest{
// 				PubKey: strconv.Itoa(100000),
// 			},
// 			err: status.Error(codes.NotFound, "challenger not found"),
// 		},
// 		{
// 			desc: "Invalid Request",
// 			err:  status.Error(codes.InvalidArgument, "invalid request"),
// 		},
// 	}{
// 		t.Run(tc.desc, func(t *testing.T) {
// 			response, err := keeper.GetChallengerByPubKey(wctx, tc.request)
// 			if tc.err != nil {
// 				require.ErrorIs(t, err, tc.err)
// 			} else {
// 				require.NoError(t, err)
// 				require.Equal(t,
// 					nullify.Fill(tc.response),
// 					nullify.Fill(response),
// 				)
// 			}
// 		})
// 	}
// }
	
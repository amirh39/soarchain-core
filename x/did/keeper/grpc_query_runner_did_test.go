package keeper_test

import (
	"testing"

	"github.com/amirh39/soarchain-core/x/did/types"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func Test_GetRunnerDidDocument(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	newDid := types.RunnerDid{
		Id:      Did,
		PubKey:  PUBKEY,
		Address: ADDRESS,
	}

	keeper.SetRunnerDid(ctx, newDid)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRunnerDidRequest
		response *types.QueryGetRunnerDidResponse
		err      error
	}{
		{
			desc: "Valid Did Id",
			request: &types.QueryGetRunnerDidRequest{
				Address: newDid.Id,
			},
			response: &types.QueryGetRunnerDidResponse{RunnerDid: newDid},
		},
		{
			desc: "Not Valid Did Id",
			request: &types.QueryGetRunnerDidRequest{
				Address: "Not-Valid",
			},
			response: &types.QueryGetRunnerDidResponse{RunnerDid: newDid},
		},
		{
			desc: "Empty",
			request: &types.QueryGetRunnerDidRequest{
				Address: "",
			},
			response: &types.QueryGetRunnerDidResponse{RunnerDid: newDid},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, found := keeper.GetRunnerDid(ctx, ADDRESS)
			if !found {
				require.Equal(t, found, false)
				require.Nil(t, response)
			}
			require.Equal(t, found, true)
			require.NotNil(t, response)
		})
	}
}

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/soarmint/types"
)

func TestMinterQuery(t *testing.T) {
	keeper, ctx := keepertest.SoarmintKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestMinter(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMinterRequest
		response *types.QueryGetMinterResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMinterRequest{},
			response: &types.QueryGetMinterResponse{Minter: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Minter(wctx, tc.request)
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

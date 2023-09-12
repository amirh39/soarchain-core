package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"soarchain/testutil/nullify"
	"soarchain/x/soarmint/types"
)

func (helper *KeeperTestHelper) TestMinterQuery() {
	helper.Setup()

	keeper := helper.App.MintKeeper
	wctx := sdk.WrapSDKContext(helper.Ctx)
	item, isFound := keeper.GetMinter(helper.Ctx)
	helper.True(isFound)
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
		helper.Run(tc.desc, func() {
			response, err := keeper.Minter(wctx, tc.request)
			if tc.err != nil {
				helper.ErrorIs(err, tc.err)
			} else {
				helper.NoError(err)
				helper.Equal(
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

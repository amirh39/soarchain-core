package keeper_test

import (
	"github.com/amirh39/soarchain-core/x/soarmint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) TestParamsQuery() {

	helper.Setup()

	keeper := helper.App.MintKeeper
	wctx := sdk.WrapSDKContext(helper.Ctx)
	params := types.DefaultParams()
	keeper.SetParams(helper.Ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	helper.NoError(err)
	helper.Equal(&types.QueryParamsResponse{Params: params}, response)
}

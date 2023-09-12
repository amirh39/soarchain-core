package keeper_test

import (
	"soarchain/x/soarmint/types"
)

func (helper *KeeperTestHelper) TestGetParams() {

	helper.Setup()
	keeper := helper.App.MintKeeper
	params := types.DefaultParams()
	keeper.SetParams(helper.Ctx, params)
	param := keeper.GetParams(helper.Ctx)
	helper.Equal(param, types.Params(types.Params{MintDenom: "udmotus", BlocksPerYear: 0x80520}))
	helper.EqualValues(params, param)

}

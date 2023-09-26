package keeper_test

import (
	"fmt"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (helper *KeeperTestHelper) Test_IsChallengeable() {

	helper.Run("TestIsChallengeable", func() {
		helper.Setup()
		keeper := helper.App.PoaKeeper
		ctx := sdk.WrapSDKContext(helper.Ctx)

		didKeeper := helper.App.DidKeeper
		reputation := SetupNReputation(1)

		didKeeper.SetReputation(helper.Ctx, reputation[0])
		helper.Require().NotEmpty(reputation)

		for _, tc := range []struct {
			desc     string
			request  *types.QueryIsChallengeableRequest
			response *types.QueryIsChallengeableResponse
			err      error
		}{
			{
				desc: "Valid Client Address",
				request: &types.QueryIsChallengeableRequest{
					ClientAddr: reputation[0].Address,
				},
				response: &types.QueryIsChallengeableResponse{ResultBool: "false", ChallengeabilityScore: "-63808599897"},
			},
		} {
			response, err := keeper.IsChallengeable(ctx, tc.request)
			fmt.Print("response----->", response)
			if err != nil {
				helper.Require().NotEmpty(err)
			} else {
				helper.Require().NoError(err)
			}
		}
	})
}

package keeper_test

import (
	"log"

	"github.com/soar-robotics/soarchain-core/app/params"
	"github.com/soar-robotics/soarchain-core/x/dpr/keeper"
	"github.com/soar-robotics/soarchain-core/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "github.com/soar-robotics/soarchain-core/x/did/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (helper *KeeperTestHelper) Test_Gen_DPR() {

	helper.Run("TestGenDpr", func() {
		helper.Setup()
		//actor := RandomAccountAddress()
		didKeeper := helper.App.DidKeeper
		epochKeeper := helper.App.EpochKeeper
		accountKeeper := helper.App.AccountKeeper
		bankKeeper := helper.App.BankKeeper

		helper.MsgServer = keeper.NewMsgServerImpl(helper.App.DprKeeper)
		ctx := sdk.WrapSDKContext(helper.Ctx)

		epochData := CreateEpochData(&epochKeeper, helper.Ctx)
		epochKeeper.SetEpochData(helper.Ctx, epochData)

		newDid := didtypes.ClientDid{
			Id:     Did,
			PubKey: PUBKEY,
		}

		didKeeper.SetClientDid(helper.Ctx, newDid)
		addr := sdk.MustAccAddressFromBech32(CREATOR)

		actorAmount := sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(1000000000000000)))

		accountKeeper.SetAccount(helper.Ctx, authtypes.NewBaseAccountWithAddress(addr))
		log.Println(accountKeeper.GetAccount(helper.Ctx, addr))

		bankKeeper.MintCoins(helper.Ctx, types.ModuleName, actorAmount)
		bankKeeper.SendCoinsFromModuleToAccount(helper.Ctx, types.ModuleName, addr, actorAmount)

		log.Println(bankKeeper.GetBalance(helper.Ctx, addr, params.BondDenom))

		res, err := helper.MsgServer.GenDpr(ctx, &types.MsgGenDpr{
			Creator: addr.String(),
			SupportedPIDs: &types.SupportedPIDs{
				Pid_1To_20:  "",
				Pid_21To_40: "8005B815",
				Pid_41To_60: "7E1C8C11",
				Pid_61To_80: "60880041",
				Pid_81To_A0: "BE1FA013",
				Pid_A1To_C0: "BE1FA013",
				Pid_C1To_E0: "BE1FA013",
				Pid_SVCTo_9: "55400000",
			},
			Duration:       45,
			DprBudget:      "1000udmotus",
			MaxClientCount: 10,
			Name:           "test",
		})
		helper.Require().Empty(res)
		helper.Require().Nil(err)
	})
}

package keeper_test

import (
	"soarchain/app/params"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	didtypes "soarchain/x/did/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (helper *KeeperTestHelper) Test_Gen_DPR() {

	helper.Run("TestGenDpr", func() {
		helper.Setup()
		//actor := RandomAccountAddress()
		didKeeper := helper.App.DidKeeper
		epochKeeper := helper.App.EpochKeeper
		//accountKeeper := helper.App.AccountKeeper
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
		addr, err := sdk.AccAddressFromBech32(CREATOR)

		//testAcc := helper.TestAccs[0]

		actorAmount := sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(1000000000000000)))
		// testAccountPubkey := secp256k1.GenPrivKeyFromSecret([]byte("acc")).PubKey()
		// testAccountAddress := sdk.AccAddress(testAccountPubkey.Address())
		// modAcc := authtypes.NewModuleAccount(authtypes.NewBaseAccount(testAccountAddress, testAccountPubkey, 1, 0),
		// 	"mint",
		// 	"permission",
		// )
		// accountKeeper.SetModuleAccount(helper.Ctx, modAcc)

		helper.App.AccountKeeper.SetAccount(helper.Ctx, authtypes.NewBaseAccount(addr, nil, 0, 0))
		bankKeeper.MintCoins(helper.Ctx, types.ModuleName, actorAmount)
		bankKeeper.SendCoinsFromModuleToAccount(helper.Ctx, types.ModuleName, sdk.AccAddress(CREATOR), actorAmount)

		res, err := helper.MsgServer.GenDpr(ctx, &types.MsgGenDpr{
			Creator: CREATOR,
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

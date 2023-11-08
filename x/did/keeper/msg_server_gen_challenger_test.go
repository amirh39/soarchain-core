package keeper_test

// import (
// 	k "soarchain/x/did/keeper"
// 	"soarchain/x/did/types"
// 	poatypes "soarchain/x/poa/types"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// )

// func (helper *KeeperTestHelper) Test_Gen_Challenger() {
// 	helper.Run("TestGenChallenger", func() {
// 		helper.Setup()
// 		keeper := helper.App.DidKeeper
// 		poakeeper := helper.App.PoaKeeper

// 		ctx := sdk.WrapSDKContext(helper.Ctx)
// 		msgServer := k.NewMsgServerImpl(keeper)

// 		item := poatypes.MasterKey{MasterCertificate: MASTER_CERTIFICATE,
// 			MasterAccount: MASTER_ACCOUNT,
// 		}
// 		poakeeper.SetMasterKey(helper.Ctx, item)

// 		updatedFactoryKeyList := poatypes.FactoryKeys{
// 			Id:          uint64(1),
// 			FactoryCert: Certificate,
// 		}
// 		poakeeper.SetFactoryKeys(helper.Ctx, updatedFactoryKeyList)

// 		deviceCert := poakeeper.GetAllFactoryKeys(helper.Ctx)
// 		helper.Require().NotEmpty(deviceCert)

// 		documentWithSequence, _ := NewChallengerDidDocumentWithSeq(Did)
// 		helper.Require().NotEmpty(documentWithSequence)

// 		res, err := msgServer.GenChallenger(ctx, &types.MsgGenChallenger{
// 			Document:        documentWithSequence.Document,
// 			Signature:       Signature,
// 			Certificate:     Certificate,
// 			Creator:         ADDRESS,
// 			ChallengerStake: Challenger_StakedAmount,
// 			ChallengerType:  Challenger_Type,
// 		})
// 		if err != nil {
// 			helper.Require().NotNil(err)
// 		} else {
// 			helper.Require().Nil(err)
// 			helper.Require().Nil(res)
// 		}
// 	})
// }

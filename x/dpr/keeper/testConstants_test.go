/** This file is created for tests. Firstly search what you nee if not find then create a new one for you. */
package keeper_test

import (
	"context"
	"soarchain/x/did/utility/crypto"
	"soarchain/x/dpr"
	"soarchain/x/dpr/keeper"
	"soarchain/x/dpr/testutil"
	"soarchain/x/dpr/types"
	"strconv"
	"testing"

	keepertest "soarchain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	didtypes "soarchain/x/did/types"

	tendermintcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func SetupNDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(111111)
		items[i].Creator = strconv.Itoa(i)
		items[i].ClientPubkeys = []string{}
		items[i].IsActive = true
		items[i].Vin = []string{strconv.Itoa(0), strconv.Itoa(1)}
		items[i].PidSupportedOneToTwnety = true
		items[i].PidSupportedTwentyOneToForthy = false
		items[i].PidSupportedForthyOneToSixty = false
		items[i].LengthOfDpr = 5
	}
	return items
}

func SetupNDifDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(22222)
		items[i].Creator = strconv.Itoa(i)
		items[i].ClientPubkeys = []string{}
		items[i].IsActive = true
		items[i].Vin = []string{strconv.Itoa(0), strconv.Itoa(1)}
		items[i].PidSupportedOneToTwnety = true
		items[i].PidSupportedTwentyOneToForthy = false
		items[i].PidSupportedForthyOneToSixty = false
		items[i].LengthOfDpr = 3
	}
	return items
}

func SetupNDeactiveDpr(n int) []types.Dpr {
	items := make([]types.Dpr, n)
	for i := range items {
		items[i].Id = strconv.Itoa(5677888)
		items[i].Creator = strconv.Itoa(i)
		items[i].ClientPubkeys = []string{}
		items[i].IsActive = false
		items[i].LengthOfDpr = 1
		items[i].PidSupportedOneToTwnety = false
		items[i].PidSupportedTwentyOneToForthy = true
		items[i].PidSupportedForthyOneToSixty = false
		items[i].Vin = []string{strconv.Itoa(0)}
	}
	return items
}

func SetupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.DprKeeperWithMocks(t, bankMock)

	dpr.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	return server, *k, context, ctrl, bankMock
}

const (
	CREATOR = "soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8"
)

func NewDIDDocumentWithSeq(did string) (didtypes.DidDocumentWithSeq, tendermintcrypto.PrivKey) {
	privKey := secp256k1.GenPrivKey()
	pubKey := crypto.PubKeyBytes(crypto.DerivePubKey(privKey))
	verificationMethodID := didtypes.NewVerificationMethodID(did, "key1")
	es256VerificationMethod := didtypes.NewVerificationMethod(verificationMethodID, didtypes.ES256K_2019, did, pubKey)
	blsVerificationMethod := didtypes.NewVerificationMethod(verificationMethodID, didtypes.BLS1281G2_2020, did, []byte("dummy BBS+ pub key"))
	verificationMethods := []*didtypes.VerificationMethod{
		&es256VerificationMethod,
		&blsVerificationMethod,
	}
	verificationRelationship := didtypes.NewVerificationRelationship(verificationMethods[0].Id)
	authentications := []didtypes.VerificationRelationship{
		verificationRelationship,
	}
	doc := didtypes.NewDidDocument(did, PUBKEY, VIN, PIDS, didtypes.WithVerificationMethods(verificationMethods), didtypes.WithAuthentications(authentications))
	docWithSeq := didtypes.NewDidDocumentWithSeq(
		&doc,
		didtypes.InitialSequence,
	)
	return docWithSeq, privKey
}

const (
	ADDRESS = "soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a"
	PUBKEY  = "3059301306072a8648ce3d020106082a8648ce3d030107034200046c28e2efdf94600435dbba5ae7f195cb619e3dd128b7e0e2877f9a1da489027819001c3e0141cb579dc3d9e913a45644401bd2458313dc37d15dd58adcaff154"
	VIN     = "1HGCM82636c678d14c93ad5bf14448da57f4f241b77e30a013d54f5d76c8126a7029aeb86"
)

var PIDS = []bool{true, false, false}

const (
	Did                  = "did:soar:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm"
	SecondDid            = "did1:soar:1Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgap"
	VerificationMethodId = Did + "#key1"
)

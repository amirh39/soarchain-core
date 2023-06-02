package poa_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa"
	"soarchain/x/poa/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClientList: []types.Client{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ChallengerList: []types.Challenger{
			{
				PubKey: "0",
			},
			{
				PubKey: "1",
			},
		},
		RunnerList: []types.Runner{
			{
				PubKey: "0",
			},
			{
				PubKey: "1",
			},
		},
		VrfDataList: []types.VrfData{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		EpochData: types.EpochData{
			TotalEpochs: 99,
			EpochV2VRX:  "61",
			EpochV2VBX:  "94",
			EpochV2NBX:  "47",
			EpochRunner: "98",
		},
		MotusWalletList: []types.MotusWallet{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		MasterKey: types.MasterKey{
			MasterCertificate: "39",
			MasterAccount:     "62",
		},
		FactoryKeysList: []types.FactoryKeys{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		FactoryKeysCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PoaKeeper(t)
	poa.InitGenesis(ctx, *k, genesisState)
	got := poa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClientList, got.ClientList)
	require.ElementsMatch(t, genesisState.ChallengerList, got.ChallengerList)
	require.ElementsMatch(t, genesisState.RunnerList, got.RunnerList)
	require.ElementsMatch(t, genesisState.VrfDataList, got.VrfDataList)
	require.Equal(t, genesisState.EpochData, got.EpochData)
	require.ElementsMatch(t, genesisState.MotusWalletList, got.MotusWalletList)
	require.Equal(t, genesisState.MasterKey, got.MasterKey)
	require.ElementsMatch(t, genesisState.FactoryKeysList, got.FactoryKeysList)
	require.Equal(t, genesisState.FactoryKeysCount, got.FactoryKeysCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

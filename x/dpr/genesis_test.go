package dpr_test

import (
	"testing"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/testutil/nullify"
	dpr "github.com/amirh39/soarchain-core/x/dpr"
	"github.com/amirh39/soarchain-core/x/dpr/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DprList: []types.Dpr{
			{
				Id: "0",
			},
			{
				Id: "1",
			},
		},
	}

	k, ctx := keepertest.DprKeeper(t)
	dpr.InitGenesis(ctx, *k, genesisState)
	got := dpr.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DprList, got.DprList)
}

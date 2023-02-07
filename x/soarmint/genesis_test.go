package soarmint_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/soarmint"
	"soarchain/x/soarmint/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SoarmintKeeper(t)
	soarmint.InitGenesis(ctx, *k, genesisState)
	got := soarmint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

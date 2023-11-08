package keeper_test

import (
	keepertest "soarchain/testutil/keeper"
	"soarchain/x/did/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IsUniqueDid(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)

	// Input two DidDocument

	newDid := types.ClientDid{
		Id:      Did,
		PubKey:  PUBKEY,
		Address: ADDRESS,
	}
	keeper.SetClientDid(ctx, newDid)

	isFound := keeper.IsNotUniqueDid(ctx, newDid.Id)
	require.Equal(t, true, isFound)
}

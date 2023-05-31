package keeper_test

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/assert"

	"soarchain/x/poa/types"
)

func Test_CreateVRF(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	vrfData, err := keeper.CreateVRF(ctx, TestUser, Factor)

	assert.Nil(t, err)

	expectedVrfData := types.VrfData{
		Index:         "",
		Creator:       "",
		Count:         "",
		Vrv:           "",
		Multiplier:    "",
		Proof:         "",
		Pubkey:        "",
		Message:       "",
		ParsedVrv:     "",
		FloatVrv:      "",
		FinalVrv:      "",
		FinalVrvFloat: "",
	}
	assert.IsType(t, expectedVrfData, vrfData)

	maxValUint64 := uint64(MaxValUint64)
	vrvBytes, _ := hex.DecodeString(vrfData.Vrv)
	parseVrvToUint64 := binary.BigEndian.Uint64(vrvBytes)
	floatVrv := float64(parseVrvToUint64) / float64(maxValUint64)
	finalVrv := floatVrv * float64(Factor)

	assert.GreaterOrEqual(t, finalVrv, float64(0))
	assert.LessOrEqual(t, finalVrv, float64(Factor))
}

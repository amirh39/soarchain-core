package keeper_test

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	keepertest "soarchain/testutil/keeper"

	"github.com/stretchr/testify/assert"

	"soarchain/x/poa/types"
)

func TestCreateVRF(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	user := "testUser"
	factor := 10

	vrfData, err := keeper.CreateVRF(ctx, user, factor)

	assert.Nil(t, err)

	t.Log("vrfData:", vrfData)

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

	maxValUint64 := uint64(18446744073709551615)
	vrvBytes, _ := hex.DecodeString(vrfData.Vrv)
	parseVrvToUint64 := binary.BigEndian.Uint64(vrvBytes)
	floatVrv := float64(parseVrvToUint64) / float64(maxValUint64)
	finalVrv := floatVrv * float64(factor)

	assert.GreaterOrEqual(t, finalVrv, float64(0))
	assert.LessOrEqual(t, finalVrv, float64(factor))
}

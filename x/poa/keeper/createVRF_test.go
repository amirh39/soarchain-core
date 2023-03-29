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
	userMultiplier := 10

	vrfData, vrfUser, err := keeper.CreateVRF(ctx, user, userMultiplier)

	assert.Nil(t, err)

	t.Log("vrfData:", vrfData)
	t.Log("vrfUser:", vrfUser)

	expectedVrfData := types.VrfData{
		Index:              "",
		Creator:            "",
		Vrv:                "",
		Multiplier:         "",
		Proof:              "",
		Pubkey:             "",
		Message:            "",
		ParsedVrv:          "",
		FloatVrv:           "",
		FinalVrv:           "",
		FinalVrvFloat:      "",
		SelectedChallenger: (*types.Challenger)(nil),
		SelectedRunner:     (*types.Runner)(nil),
	}
	assert.IsType(t, expectedVrfData, vrfData)

	expectedVrfUser := types.VrfUser{
		Index:   "",
		Address: "",
		Count:   "",
	}
	assert.IsType(t, expectedVrfUser, vrfUser)

	maxValUint64 := uint64(18446744073709551615)
	vrvBytes, _ := hex.DecodeString(vrfData.Vrv)
	parseVrvToUint64 := binary.BigEndian.Uint64(vrvBytes)
	floatVrv := float64(parseVrvToUint64) / float64(maxValUint64)
	finalVrv := floatVrv * float64(userMultiplier)

	assert.GreaterOrEqual(t, finalVrv, float64(0))
	assert.LessOrEqual(t, finalVrv, float64(userMultiplier))
}

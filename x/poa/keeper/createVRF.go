package keeper

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"

	"github.com/soar-robotics/soarchain-core/x/poa/types"
	"github.com/soar-robotics/soarchain-core/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateVRF(ctx sdk.Context, msgCreator string, multiplier int) (types.VrfData, error) {

	vrfData, isFound := k.GetVrfData(ctx, msgCreator)

	var user_key_count int64 = 1
	currentUserCount, _ := strconv.Atoi(vrfData.Count)

	if isFound {
		user_key_count = int64(currentUserCount) + 1
	}

	random_val_key := msgCreator + "," + strconv.FormatInt(user_key_count, 10)
	message := []byte(random_val_key)

	vrv, proof, pub_key, float_vrv, final_vrv, final_vrv_float, _ := utility.CreateRandomVrfValues(ctx, message, multiplier)

	newRandomVal := types.VrfData{
		Index:         random_val_key,
		Creator:       msgCreator,
		Count:         strconv.FormatUint(uint64(user_key_count), 10),
		Vrv:           hex.EncodeToString(vrv),
		Multiplier:    strconv.FormatUint(uint64(multiplier), 10),
		Proof:         hex.EncodeToString(proof),
		Pubkey:        hex.EncodeToString(pub_key),
		Message:       random_val_key,
		ParsedVrv:     strconv.FormatUint(binary.BigEndian.Uint64(vrv), 10),
		FloatVrv:      strconv.FormatFloat(float_vrv, 'f', 0, 64),
		FinalVrv:      strconv.FormatUint(uint64(final_vrv), 10),
		FinalVrvFloat: strconv.FormatFloat(final_vrv_float, 'f', 0, 64),
	}
	k.SetVrfData(ctx, newRandomVal)

	return newRandomVal, nil
}

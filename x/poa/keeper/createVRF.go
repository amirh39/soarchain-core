package keeper

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strconv"

	"github.com/coniks-sys/coniks-go/crypto/vrf"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CreateVRF(ctx sdk.Context, msgCreator string, multiplier int) (types.VrfData, types.VrfUser, error) {

	err_VrfData := types.VrfData{}
	err_VrfUser := types.VrfUser{}

	userval, isFound := k.GetVrfUser(ctx, msgCreator)

	var user_key_count int64 = 1
	currentUserCount, _ := strconv.Atoi(userval.Count)

	if isFound {
		user_key_count = int64(currentUserCount) + 1
	}

	// VRF
	seed := make([]byte, 64)
	binary.LittleEndian.PutUint64(seed[:8], uint64(ctx.BlockHeight()))

	sk, err := vrf.GenerateKey(bytes.NewReader(seed))
	if err != nil {
		return err_VrfData, err_VrfUser, sdkerrors.Wrap(sdkerrors.ErrPanic, "Couldn't generate VRF key!")
	}

	random_val_key := msgCreator + "," + strconv.FormatInt(user_key_count, 10)
	a_message := []byte(random_val_key)

	vrv, proof := sk.Prove(a_message) // Generate vrv (verifiable random value) and proof
	pub_key, ok_bool := sk.Public()   // public key creation
	if !ok_bool {
		return err_VrfData, err_VrfUser, sdkerrors.Wrap(sdkerrors.ErrPanic, "Couldn't generate VRF public key!")
	}

	var max_val_uint64 uint64 = 18446744073709551615
	parse_vrv_to_uint64 := binary.BigEndian.Uint64(vrv)
	var float_vrv float64 = float64(parse_vrv_to_uint64) / float64(max_val_uint64)
	final_vrv := float_vrv * float64(multiplier)
	final_vrv_float := float_vrv * float64(multiplier)

	if uint64(multiplier) < uint64(final_vrv) {
		return err_VrfData, err_VrfUser, sdkerrors.Wrap(sdkerrors.ErrPanic, "Generated random number is out of index!")
	}

	newRandomVal := types.VrfData{
		Index:              random_val_key,
		Creator:            msgCreator,
		Vrv:                hex.EncodeToString(vrv),
		Multiplier:         strconv.FormatUint(uint64(multiplier), 10),
		Proof:              hex.EncodeToString(proof),
		Pubkey:             hex.EncodeToString(pub_key),
		Message:            random_val_key,
		ParsedVrv:          strconv.FormatUint(binary.BigEndian.Uint64(vrv), 10),
		FloatVrv:           strconv.FormatFloat(float_vrv, 'f', 0, 64),
		FinalVrv:           strconv.FormatUint(uint64(final_vrv), 10),
		FinalVrvFloat:      strconv.FormatFloat(final_vrv_float, 'f', 0, 64),
		SelectedChallenger: nil,
		SelectedRunner:     nil,
	}
	k.SetVrfData(ctx, newRandomVal)

	newVrfUser := types.VrfUser{
		Index:   msgCreator,
		Address: msgCreator,
		Count:   strconv.FormatUint(uint64(user_key_count), 10),
	}
	k.SetVrfUser(ctx, newVrfUser)

	return newRandomVal, newVrfUser, nil
}

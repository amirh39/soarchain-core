package keeper

import (
	"context"
	"encoding/binary"

	"encoding/hex"
	"strconv"

	"github.com/coniks-sys/coniks-go/crypto/vrf"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SelectRandomChallenger(goCtx context.Context, msg *types.MsgSelectRandomChallenger) (*types.MsgSelectRandomChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	allChallengers := k.GetAllChallenger(ctx)
	multiplier := int(len(allChallengers))

	userval, isFound := k.GetVrfUser(ctx, msg.Creator)

	var user_key_count int64 = 1
	currentUserCount, _ := strconv.Atoi(userval.Count)

	if isFound {
		user_key_count = int64(currentUserCount) + 1
	}

	// VRF
	sk, err := vrf.GenerateKey(nil)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Couldn't generate VRF key!")
	}

	random_val_key := msg.Creator + "," + strconv.FormatInt(user_key_count, 10)
	a_message := []byte(random_val_key)

	vrv, proof := sk.Prove(a_message) // Generate vrv (verifiable random value) and proof
	pub_key, ok_bool := sk.Public()   // public key creation
	if ok_bool == false {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Couldn't generate VRF public key!")
	}

	var max_val_uint64 uint64 = 18446744073709551615
	parse_vrv_to_uint64 := binary.BigEndian.Uint64(vrv)
	var float_vrv float64 = float64(parse_vrv_to_uint64) / float64(max_val_uint64)
	final_vrv := float_vrv * float64(multiplier)
	final_vrv_float := float_vrv * float64(multiplier)

	finalRNG := uint64(final_vrv)

	if uint64(multiplier) < uint64(final_vrv) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Generated random number is out of index!")
	}

	//
	var selectedChallenger types.Challenger
	challengers := k.GetAllChallenger(ctx)
	for i := 0; i < len(challengers); i++ {
		if i == int(finalRNG) {
			selectedChallenger = challengers[i]
		}
	}
	//

	newRandomVal := types.VrfData{
		Index:              random_val_key,
		Creator:            msg.Creator,
		Vrv:                hex.EncodeToString(vrv),
		Multiplier:         strconv.FormatUint(uint64(multiplier), 10),
		Proof:              hex.EncodeToString(proof),
		Pubkey:             hex.EncodeToString(pub_key),
		Message:            random_val_key,
		ParsedVrv:          strconv.FormatUint(binary.BigEndian.Uint64(vrv), 10),
		FloatVrv:           strconv.FormatFloat(float_vrv, 'f', 0, 64),
		FinalVrv:           strconv.FormatUint(uint64(final_vrv), 10),
		FinalVrvFloat:      strconv.FormatFloat(final_vrv_float, 'f', 0, 64),
		SelectedChallenger: &selectedChallenger,
	}
	k.SetVrfData(ctx, newRandomVal)

	// resultStr, err := k.VerifyGeneratedNumber(ctx, newRandomVal)
	// if err != nil && resultStr != "true" {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Couldn't verify generated VRF")
	// }

	newVrfUser := types.VrfUser{
		Index:   msg.Creator,
		Address: msg.Creator,
		Count:   strconv.FormatUint(uint64(user_key_count), 10),
	}
	k.SetVrfUser(ctx, newVrfUser)

	return &types.MsgSelectRandomChallengerResponse{RandomChallenger: &selectedChallenger}, nil
	// return &types.MsgSelectRandomChallengerResponse{RandomChallenger: (strconv.FormatUint(finalRNG, 10))}, nil
}

func (k Keeper) VerifyGeneratedNumber(ctx sdk.Context, req *types.QueryVerifyRandomNumberRequest) (bool, error) {

	var public_key vrf.PublicKey
	public_key, err := hex.DecodeString(req.Pubkey)
	if err != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Public Key cannot be decoded")
	}
	message_value := []byte(req.Message)
	vrv_value, err := hex.DecodeString(req.Vrv)
	if err != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "VRV Value cannot be decoded")
	}

	proof_value, err := hex.DecodeString(req.Proof)
	if err != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Proof value cannot be decoded")
	}

	is_verified := public_key.Verify(message_value, vrv_value, proof_value)

	// return strconv.FormatBool(is_verified), err
	return is_verified, err

}

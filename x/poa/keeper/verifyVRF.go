package keeper

import (
	"encoding/hex"

	"github.com/coniks-sys/coniks-go/crypto/vrf"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) VerifyGeneratedNumber(ctx sdk.Context, req *types.QueryVerifyRandomNumberRequest) (bool, error) {
	logger := k.Logger(ctx)

	var public_key vrf.PublicKey
	public_key, err := hex.DecodeString(req.Pubkey)
	if err != nil {
		return false, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[VerifyGeneratedNumber][DecodeString-PublicKey] failed. Couldn't decode given public key by the rquest. got [ %T ]. Error: [ %T ]", req.Pubkey, err)
	}

	message_value := []byte(req.Message)
	vrv_value, err := hex.DecodeString(req.Vrv)
	if err != nil {
		return false, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[VerifyGeneratedNumber][DecodeString-Message] failed. Couldn't decode given message by the rquest. got [ %T ]. Error: [ %T ]", req.Message, err)
	}

	proof_value, err := hex.DecodeString(req.Proof)
	if err != nil {
		return false, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[VerifyGeneratedNumber][DecodeString-Proof] failed. Couldn't decode given proof by the rquest. got [ %T ]. Error: [ %T ]", req.Proof, err)
	}

	is_verified := public_key.Verify(message_value, vrv_value, proof_value)

	if logger != nil {
		logger.Info("Verifying VRF data successfully done.", "calculation of", "VerifyGeneratedNumber")
	}

	// return strconv.FormatBool(is_verified), err
	return is_verified, err

}

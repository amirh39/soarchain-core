package keeper

import (
	"encoding/hex"

	"github.com/coniks-sys/coniks-go/crypto/vrf"

	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) VerifyGeneratedNumber(ctx sdk.Context, req *types.QueryVerifyRandomNumberRequest) (bool, error) {

	var public_key vrf.PublicKey
	public_key, err := hex.DecodeString(req.Pubkey)
	if err != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "[VerifyGeneratedNumber][DecodeString] failed. Couldn't decode given public key by the rquest."+err.Error())
	}

	message_value := []byte(req.Message)
	vrv_value, err := hex.DecodeString(req.Vrv)
	if err != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[VerifyGeneratedNumber][DecodeString] failed. Couldn't decode given VRV by the request."+err.Error())
	}

	proof_value, err := hex.DecodeString(req.Proof)
	if err != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[VerifyGeneratedNumber][DecodeString] failed. Couldn't decode given proof by the request."+err.Error())
	}

	is_verified := public_key.Verify(message_value, vrv_value, proof_value)

	// return strconv.FormatBool(is_verified), err
	return is_verified, err

}

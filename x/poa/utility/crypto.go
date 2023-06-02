package utility

import (
	"bytes"
	"encoding/binary"
	"soarchain/x/poa/utility/utilConstants"

	"github.com/coniks-sys/coniks-go/crypto/vrf"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CreateRandomVrfValues(ctx sdk.Context, message []byte, multiplier int) (vrv []byte, prrof []byte, pub_key vrf.PublicKey, float_vrv float64, final_vrv float64, final_vrv_float float64, err error) {

	seed := make([]byte, 64)
	binary.LittleEndian.PutUint64(seed[:8], uint64(ctx.BlockHeight()))

	sk, _ := vrf.GenerateKey(bytes.NewReader(seed))
	vrv, proof := sk.Prove(message) // Generate vrv (verifiable random value) and proof
	pub_key, ok_bool := sk.Public() // public key creation
	if !ok_bool {
		return nil, nil, nil, 0, 0, 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "[CreateVrfData] failed. Couldn't generate VRF public key!")
	}

	parse_vrv_to_uint64 := binary.BigEndian.Uint64(vrv)
	float_vrv = float64(parse_vrv_to_uint64) / float64(utilConstants.MaxValUnit64)
	final_vrv = float_vrv * float64(multiplier)
	final_vrv_float = float_vrv * float64(multiplier)

	if uint64(multiplier) < uint64(final_vrv) {
		return nil, nil, nil, 0, 0, 0, sdkerrors.Wrap(sdkerrors.ErrPanic, "[CreateVrfData] failed. Generated random number is out of index.")
	}

	return vrv, proof, pub_key, float_vrv, final_vrv, final_vrv_float, nil
}

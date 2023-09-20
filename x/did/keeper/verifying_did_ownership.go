package keeper

import (
	"soarchain/x/did/types"
	"soarchain/x/did/utility/crypto"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) VerifyDidOwnership(signData *types.DidDocument, seq uint64, doc *types.DidDocument, verificationMethodId string, sig []byte) (uint64, error) {

	verificationMethod, ok := doc.VerificationMethodFrom(doc.Authentications, verificationMethodId)
	if !ok {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[VerifyDIDOwnership][VerificationMethodFrom] failed. Authentication for did failed.")
	}

	if verificationMethod.Type != types.ES256K_2019 && verificationMethod.Type != types.ES256K_2018 {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[VerifyDIDOwnership] failed. Make sure using valid verification method. Use ES256K_2019 OR ES256K_2018 methods.")
	}

	pubKeySecp256k1, err := crypto.PubKeyFromBase58(verificationMethod.PublicKeyBase58)
	if err != nil {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[VerifyDIDOwnership][PubKeyFromBase58] failed. Make sure using a valid pubkey.")
	}

	newSeq, ok := types.Verify(sig, signData, seq, pubKeySecp256k1)
	if !ok {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[VerifyDIDOwnership][VerifySignature] failed. Make sure using valid pubkey and signature.")
	}
	return newSeq, nil
}

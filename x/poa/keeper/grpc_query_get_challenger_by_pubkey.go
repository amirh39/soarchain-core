package keeper

import (
	"context"
	"soarchain/x/poa/types"
	"soarchain/x/poa/utility"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetChallengerByPubKey(goCtx context.Context, req *types.QueryGetChallengerByPubKeyRequest) (*types.QueryGetChallengerByPubKeyResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "[GetChallengerByPubKey] failed. Invalid request: %T.", req)
	}

	if !utility.ValidPubkey(req.Pubkey) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[GetChallengerByPubKey][ValidPubkey] failed. Couldn't find a valid public key from the request. got: [ %T ], Make sure public key is not empty OR invalid.", req.Pubkey)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challengers := k.GetAllChallenger(ctx)
	if len(challengers) == 0 || challengers == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GetChallengerByPubKey][GetAllChallenger] failed. Couldn't find any challenger.")
	}

	targetChallenger := types.Challenger{}

	for _, challenger := range challengers {
		if req.Pubkey == challenger.PubKey {
			targetChallenger = challenger
			break
		}
	}

	if targetChallenger.PubKey == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[GetChallengerByPubKey] failed. Couldn't find a valid public key for target challenger. got: [ %T ], Make sure public key is not empty OR invalid.", targetChallenger.PubKey)
	}

	return &types.QueryGetChallengerByPubKeyResponse{Challenger: &targetChallenger}, nil
}

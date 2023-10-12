package keeper

import (
	"context"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetChallengerDidByPubKey(goCtx context.Context, req *types.QueryGetChallengerDidByPubKeyRequest) (*types.QueryGetChallengerDidByPubKeyResponse, error) {
	if req == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "[GetChallengerByPubKey] failed. Invalid request: %T.", req)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challengers := k.GetAllChallengerDid(ctx)
	if len(challengers) == 0 || challengers == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[GetChallengerByPubKey][GetAllChallenger] failed. Couldn't find any challenger.")
	}

	targetChallenger := types.ChallengerDid{}

	for _, challenger := range challengers {
		if req.Pubkey == challenger.PubKey {
			targetChallenger = challenger
			break
		}
	}

	if targetChallenger.PubKey == "" {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[GetChallengerByPubKey] failed. Couldn't find a valid public key for target challenger. got: [ %T ], Make sure public key is not empty OR invalid.", targetChallenger.PubKey)
	}

	return &types.QueryGetChallengerDidByPubKeyResponse{ChallengerDid: &targetChallenger}, nil
}

package wasmbinding

import (
	"encoding/json"

	"github.com/soar-robotics/soarchain-core/wasmbinding/bindings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.SoarchainQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CustomQuerier][Unmarshal Contract Query Result] failed. Contract query is not valid, couldn't be parsed.")
		}

		switch {
		case contractQuery.ChallengerByIndex != nil:
			index := contractQuery.ChallengerByIndex.Index
			challenger, err := GetChallenger(ctx, index, *qp.keeper)
			if err != nil {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[CustomQuerier][GetChallengerList] failed. Challenger pubkey: [ %T ] is not valid, couldn't be fetched.", index)
			}

			response := bindings.ChallengerByIndexResponse{Address: challenger.Address, Index: challenger.Index, Score: challenger.Score}

			bz, err := json.Marshal(response)
			if err != nil {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CustomQuerier][Marshal] failed. Challenger couldn't be marshaled for the publickey [ %T ].", index)
			}

			return bz, nil

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CustomQuerier][GetChallengerByIndex] failed. unknown soarchain query variante.")
		}
	}
}

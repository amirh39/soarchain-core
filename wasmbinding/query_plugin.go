package wasmbinding

import (
	"encoding/json"
	"soarchain/wasmbinding/bindings"

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
		case contractQuery.ClientByIndex != nil:
			index := contractQuery.ClientByIndex.Index
			client, err := GetClientByIndex(ctx, index, *qp.keeper)
			if err != nil {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "[CustomQuerier][GetClientList] failed. Client pubkey: [ %T ] is not valid, couldn't be fetched.", index)
			}

			response := bindings.ClientByIndexResponse{Address: client.Address, Index: client.Index, Score: client.Score}

			bz, err := json.Marshal(response)
			if err != nil {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CustomQuerier][Marshal] failed. Client couldn't be marshaled for the publickey [ %T ].", index)
			}

			return bz, nil

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CustomQuerier][GetClientByIndex] failed. unknown soarchain query variante.")
		}
	}
}

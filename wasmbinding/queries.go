package wasmbinding

import (
	"log"

	wasmbindings "github.com/amirh39/soarchain-core/wasmbinding/bindings"
	keeper "github.com/amirh39/soarchain-core/x/poa/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type QueryPlugin struct {
	keeper *keeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(
	poa *keeper.Keeper,
) *QueryPlugin {
	return &QueryPlugin{
		keeper: poa,
	}
}

func (qp QueryPlugin) GetChallengerByIndex(ctx sdk.Context, index string) (*wasmbindings.ChallengerByIndex, error) {

	log.Println("############## Smart contract query for fetching a Challenger is Started ##############")

	challenger, found := qp.keeper.GetReputation(ctx, index)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[Querier][GetChallengerByIndex] failed. Challenger with the index: [ %T ] for query wasm contract is not found.", index)
	}
	var challengerByIndex wasmbindings.ChallengerByIndex
	challengerByIndex.Index = challenger.Address

	log.Println("############## End of Smart contract query for fetching a Challenger ##############")

	return &challengerByIndex, nil
}

package wasmbinding

import (
	"log"
	wasmbindings "soarchain/wasmbinding/bindings"
	keeper "soarchain/x/poa/keeper"

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

func (qp QueryPlugin) GetClientByIndex(ctx sdk.Context, index string) (*wasmbindings.ClientByIndex, error) {

	log.Println("############## Smart contract query for fetching a clinet is Started ##############")

	client, found := qp.keeper.GetClient(ctx, index)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[Querier][GetClientByIndex] failed. Client with the index: [ %T ] for query wasm contract is not found.", index)
	}
	var clientByIndex wasmbindings.ClientByIndex
	clientByIndex.Index = client.Index

	log.Println("############## End of Smart contract query for fetching a clinet ##############")

	return &clientByIndex, nil
}

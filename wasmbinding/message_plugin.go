package wasmbinding

import (
	"encoding/json"
	"log"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"

	"soarchain/wasmbinding/bindings"
	poaKeepers "soarchain/x/poa/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CustomMessageDecorator(bank *bankkeeper.BaseKeeper, poa *poaKeepers.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(wasm wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wasm: wasm,
			bank: bank,
			poa:  poa,
		}
	}
}

type CustomMessenger struct {
	wasm wasmkeeper.Messenger
	bank *bankkeeper.BaseKeeper
	poa  *poaKeepers.Keeper
	//poaKeeper poaKeepers.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddress sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		var contractMsg bindings.SoarchainMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "[CustomQuerier][Marshal] failed. Client response is not valid, couldn't marshal the client response.")
		}
	}
	return m.wasm.DispatchMsg(ctx, contractAddress, contractIBCPortID, msg)
}

func GetClientByIndex(ctx sdk.Context, index string, poa poaKeepers.Keeper) (res bindings.ClientByIndexResponse, err error) {

	log.Println("############## Smart contract query for fetching a clinet is Started ##############")

	var response bindings.ClientByIndexResponse

	logger := poa.Logger(ctx)

	client, found := poa.GetClient(ctx, index)
	if !found {
		return response, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "[Message_PlugIn][GetClientByIndex] failed. Client with the index: [ %T ] for query wasm contract not found.", index)
	}

	if logger != nil {
		logger.Info("Fetching smart contract query for a client successfully done.", "query", "GetClientByIndex", "ClientPublicKey:", index)
	}

	response.Address = client.Address
	response.Index = client.Index
	response.Score = client.Score

	log.Println("############## End of Smart contract query for fetching a clinet ##############")

	return response, nil
}

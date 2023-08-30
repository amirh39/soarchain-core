package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	keeper "soarchain/x/poa/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	poa *keeper.Keeper,
) []wasmkeeper.Option {

	wasmQueryPlugin := NewQueryPlugin(poa)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})

	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, poa),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}

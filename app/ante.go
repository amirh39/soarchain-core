package app

import (
	wasm "github.com/CosmWasm/wasmd/x/wasm"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

func NewAnteHandler(
	appOpts servertypes.AppOptions,
	wasmConfig wasm.Config,

)

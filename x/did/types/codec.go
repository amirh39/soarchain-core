package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGenClient{}, "did/GenClient", nil)
	cdc.RegisterConcrete(&MsgGenRunner{}, "did/GenRunner", nil)
	cdc.RegisterConcrete(&MsgGenChallenger{}, "did/GenChallenger", nil)
	cdc.RegisterConcrete(&MsgDeactivateDid{}, "did/DeactivateDid", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenClient{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenRunner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenChallenger{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeactivateDid{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

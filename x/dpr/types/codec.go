package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGenDpr{}, "dpr/GenDpr", nil)
	cdc.RegisterConcrete(&MsgEnterDpr{}, "dpr/EnterDpr", nil)
	cdc.RegisterConcrete(&MsgActivateDpr{}, "dpr/ActivateDpr", nil)
	cdc.RegisterConcrete(&MsgLeaveDpr{}, "dpr/LeaveDpr", nil)

	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenDpr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnterDpr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgActivateDpr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLeaveDpr{},
	)

	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

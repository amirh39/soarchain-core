package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGenClient{}, "poa/GenClient", nil)
	cdc.RegisterConcrete(&MsgGenChallenger{}, "poa/GenChallenger", nil)
	cdc.RegisterConcrete(&MsgChallengeService{}, "poa/ChallengeService", nil)
	cdc.RegisterConcrete(&MsgUnregisterClient{}, "poa/UnregisterClient", nil)
	cdc.RegisterConcrete(&MsgUnregisterChallenger{}, "poa/UnregisterChallenger", nil)
	cdc.RegisterConcrete(&MsgGenGuard{}, "poa/GenGuard", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenClient{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenChallenger{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChallengeService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnregisterClient{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnregisterChallenger{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenGuard{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

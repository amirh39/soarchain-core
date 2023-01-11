package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGenClient{}, "poa/GenClient", nil)
	cdc.RegisterConcrete(&MsgChallengeService{}, "poa/ChallengeService", nil)
	cdc.RegisterConcrete(&MsgUnregisterClient{}, "poa/UnregisterClient", nil)
	cdc.RegisterConcrete(&MsgUnregisterChallenger{}, "poa/UnregisterChallenger", nil)
	cdc.RegisterConcrete(&MsgGenGuard{}, "poa/GenGuard", nil)
	cdc.RegisterConcrete(&MsgUnregisterRunner{}, "poa/UnregisterRunner", nil)
	cdc.RegisterConcrete(&MsgRunnerChallenge{}, "poa/RunnerChallenge", nil)
	cdc.RegisterConcrete(&MsgUnregisterGuard{}, "poa/UnregisterGuard", nil)
	cdc.RegisterConcrete(&MsgSelectRandomChallenger{}, "poa/SelectRandomChallenger", nil)
	cdc.RegisterConcrete(&MsgSelectRandomRunner{}, "poa/SelectRandomRunner", nil)
	cdc.RegisterConcrete(&MsgUpdateGuard{}, "poa/UpdateGuard", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenClient{},
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnregisterRunner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRunnerChallenge{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnregisterGuard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSelectRandomChallenger{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSelectRandomRunner{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateGuard{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

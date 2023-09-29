package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgChallengeService{}, "poa/ChallengeService", nil)
	cdc.RegisterConcrete(&MsgUnregisterChallenger{}, "poa/UnregisterChallenger", nil)
	cdc.RegisterConcrete(&MsgUnregisterRunner{}, "poa/UnregisterRunner", nil)
	cdc.RegisterConcrete(&MsgRunnerChallenge{}, "poa/RunnerChallenge", nil)
	cdc.RegisterConcrete(&MsgSelectRandomChallenger{}, "poa/SelectRandomChallenger", nil)
	cdc.RegisterConcrete(&MsgSelectRandomRunner{}, "poa/SelectRandomRunner", nil)
	cdc.RegisterConcrete(&MsgClaimMotusRewards{}, "poa/ClaimMotusRewards", nil)
	cdc.RegisterConcrete(&MsgClaimRunnerRewards{}, "poa/ClaimRunnerRewards", nil)
	cdc.RegisterConcrete(&MsgRegisterFactoryKey{}, "poa/RegisterFactoryKey", nil)
	cdc.RegisterConcrete(&MsgGenRunner{}, "poa/GenRunner", nil)
	cdc.RegisterConcrete(&MsgGenChallenger{}, "poa/GenChallenger", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChallengeService{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnregisterChallenger{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnregisterRunner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRunnerChallenge{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSelectRandomChallenger{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSelectRandomRunner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimMotusRewards{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimRunnerRewards{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterFactoryKey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenRunner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGenChallenger{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

package keepers

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"

	mintkeeper "github.com/amirh39/soarchain-core/x/soarmint/keeper"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	monitoringpkeeper "github.com/tendermint/spn/x/monitoringp/keeper"

	didmodulekeeper "github.com/amirh39/soarchain-core/x/did/keeper"
	dprmodulekeeper "github.com/amirh39/soarchain-core/x/dpr/keeper"
	epochmodulekeeper "github.com/amirh39/soarchain-core/x/epoch/keeper"
	poamodulekeeper "github.com/amirh39/soarchain-core/x/poa/keeper"
)

type AppKeepers struct {

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	AuthzKeeper      authzkeeper.Keeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	MonitoringKeeper monitoringpkeeper.Keeper
	WasmKeeper       wasm.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper        capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedMonitoringKeeper capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper       capabilitykeeper.ScopedKeeper

	EpochKeeper epochmodulekeeper.Keeper
	PoaKeeper   poamodulekeeper.Keeper
	DidKeeper   didmodulekeeper.Keeper
	DprKeeper   dprmodulekeeper.Keeper
}

func NewAppKeepers() {}

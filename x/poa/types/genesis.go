package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ClientList:     []Client{},
		ChallengerList: []Challenger{},
		RunnerList:     []Runner{},
		GuardList:      []Guard{},
		VrfDataList:    []VrfData{},
		VrfUserList:    []VrfUser{},
		EpochData: EpochData{
			TotalEpochs: 0,
			EpochV2VRX:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
			EpochV2VBX:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
			EpochV2NBX:  sdk.NewCoin("soar", sdk.ZeroInt()).String(),
			EpochRunner: sdk.NewCoin("soar", sdk.ZeroInt()).String(),
		},
		MotusWalletList: []MotusWallet{},
		MasterKey: MasterKey{
			MasterPubkey:  "3056301006072a8648ce3d020106052b8104000a034200040349a2dde2c994f767f595b6b497c0f2a24bde0731a60a20765e4742e1349c04480634a7858d01413e7c360c544d4e57d857b008bee5b54f319897ff727c2271",
			MasterAccount: "",
		},
		FactoryKeysList: []FactoryKeys{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in client
	clientIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClientList {
		index := string(ClientKey(elem.Index))
		if _, ok := clientIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for client")
		}
		clientIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in challenger
	challengerIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChallengerList {
		index := string(ChallengerKey(elem.Index))
		if _, ok := challengerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for challenger")
		}
		challengerIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in runner
	runnerIndexMap := make(map[string]struct{})

	for _, elem := range gs.RunnerList {
		index := string(RunnerKey(elem.Index))
		if _, ok := runnerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for runner")
		}
		runnerIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in guard
	guardIndexMap := make(map[string]struct{})

	for _, elem := range gs.GuardList {
		index := string(GuardKey(elem.Index))
		if _, ok := guardIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for guard")
		}
		guardIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in vrfData
	vrfDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.VrfDataList {
		index := string(VrfDataKey(elem.Index))
		if _, ok := vrfDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vrfData")
		}
		vrfDataIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in vrfUser
	vrfUserIndexMap := make(map[string]struct{})

	for _, elem := range gs.VrfUserList {
		index := string(VrfUserKey(elem.Index))
		if _, ok := vrfUserIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vrfUser")
		}
		vrfUserIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in motusWallet
	motusWalletIndexMap := make(map[string]struct{})

	for _, elem := range gs.MotusWalletList {
		index := string(MotusWalletKey(elem.Index))
		if _, ok := motusWalletIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for motusWallet")
		}
		motusWalletIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in factoryKeys
	factoryKeysIdMap := make(map[uint64]bool)
	factoryKeysCount := gs.GetFactoryKeysCount()
	for _, elem := range gs.FactoryKeysList {
		if _, ok := factoryKeysIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for factoryKeys")
		}
		if elem.Id >= factoryKeysCount {
			return fmt.Errorf("factoryKeys id should be lower or equal than the last id")
		}
		factoryKeysIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

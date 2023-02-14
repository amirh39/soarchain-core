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
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

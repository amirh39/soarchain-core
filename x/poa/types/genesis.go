package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ReputationList: []Reputation{},
		ChallengerList: []Challenger{},
		RunnerList:     []Runner{},
		VrfDataList:    []VrfData{},

		MasterKey: MasterKey{
			MasterCertificate: "",
			MasterAccount:     "",
		},
		FactoryKeysList: []FactoryKeys{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	reputationIndexMap := make(map[string]struct{})

	for _, elem := range gs.ReputationList {
		index := string(ReputationKey(elem.PubKey))
		if _, ok := reputationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for reputation")
		}
		reputationIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in challenger
	challengerIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChallengerList {
		index := string(ChallengerKey(elem.PubKey))
		if _, ok := challengerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for challenger")
		}
		challengerIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in runner
	runnerIndexMap := make(map[string]struct{})

	for _, elem := range gs.RunnerList {
		index := string(RunnerKey(elem.PubKey))
		if _, ok := runnerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for runner")
		}
		runnerIndexMap[index] = struct{}{}
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

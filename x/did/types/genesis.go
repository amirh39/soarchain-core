package types

import fmt "fmt"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ClientDidList:     []ClientDid{},
		ChallengerDidList: []ChallengerDid{},
		RunnerDidList:     []RunnerDid{},
		Params:            DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in clientDid
	clientDidIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClientDidList {
		index := string(DidKey(elem.Id))
		if _, ok := clientDidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for clientDid")
		}
		clientDidIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in challengerDid
	challengerDidIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChallengerDidList {
		index := string(ChallengerDidKey(elem.Id))
		if _, ok := challengerDidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for challengerDid")
		}
		challengerDidIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in runnerDid
	runnerDidIndexMap := make(map[string]struct{})

	for _, elem := range gs.RunnerDidList {
		index := string(RunnerDidKey(elem.Id))
		if _, ok := runnerDidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for runnerDid")
		}
		runnerDidIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}

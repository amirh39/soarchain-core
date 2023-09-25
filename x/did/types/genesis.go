package types

import fmt "fmt"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	clientIndexMap := make(map[string]struct{})

	for _, elem := range gs.ReputationList {
		index := string(ReputationKey(elem.Index))
		if _, ok := clientIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for client")
		}
		clientIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

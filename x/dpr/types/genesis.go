package types

import fmt "fmt"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		DprList: []Dpr{},
		Dpr: Dpr{
			Id:                            "",
			Creator:                       "",
			PidSupportedOneToTwnety:       false,
			PidSupportedTwentyOneToForthy: false,
			PidSupportedForthyOneToSixty:  false,
			IsActive:                      false,
			Vin:                           []string{},
			ClientPubkeys:                 []string{},
			LengthOfDpr:                   0,
		},
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	dprIndexMap := make(map[string]struct{})

	for _, elem := range gs.DprList {
		index := string(DprKey(elem.Id))
		if _, ok := dprIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for challenger")
		}
		dprIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}

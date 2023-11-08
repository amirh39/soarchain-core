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
			Id:      "",
			Creator: "",
			SupportedPIDs: &SupportedPIDs{
				Pid_1To_20:  "",
				Pid_21To_40: "",
				Pid_41To_60: "",
				Pid_61To_80: "",
				Pid_81To_A0: "",
				Pid_A1To_C0: "",
				Pid_C1To_E0: "",
				Pid_SVCTo_9: "",
			},
			Status:        0,
			Duration:      0,
			DprEndTime:    "",
			DprStartEpoch: 0,
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

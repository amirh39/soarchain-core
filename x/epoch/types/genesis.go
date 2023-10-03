package types

import (
	"soarchain/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		EpochData: EpochData{

			TotalEpochs:                   0,
			EpochV2VRX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochV2VBX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochV2NBX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochRunner:                   sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochChallenger:               sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			V2VRXTotalChallenges:          0,
			V2VBXTotalChallenges:          0,
			V2NBXTotalChallenges:          0,
			RunnerTotalChallenges:         0,
			ChallengerTotalChallenges:     0,
			V2VRXLastBlockChallenges:      0,
			V2VBXLastBlockChallenges:      0,
			V2NBXLastBlockChallenges:      0,
			RunnerLastBlockChallenges:     0,
			ChallengerLastBlockChallenges: 0,
			TotalChallengesPrevDay:        0,
			InitialPerChallengeValue:      9000000.0,
			V2NBXPerChallengeValue:        3000000,
			RunnerPerChallengeValue:       1000000,
			ChallengerPerChallengeValue:   1000000,
			V2VBXPerChallengeValue:        2000000,
			V2VRXPerChallengeValue:        2000000,
		},
		RandomData: RandomData{
			Id:           "0",
			RandomNumber: "0",
			EpochNumber:  "0",
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

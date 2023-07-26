package types

import (
	"fmt"

	params "soarchain/app/params"

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
		VrfDataList:    []VrfData{},
		EpochData: EpochData{
			TotalEpochs:                   0,
			EpochV2VRX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochV2VBX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochV2NBX:                    sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochRunner:                   sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			EpochChallenger:               sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
			V2VRXtotalChallenges:          0,
			V2VBXtotalChallenges:          0,
			V2NBXtotalChallenges:          0,
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
		MotusWalletList: []MotusWallet{},
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

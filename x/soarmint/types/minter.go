package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(inflation, annualProvisions sdk.Dec, phase, startPhaseBlock uint64, targetSupply sdk.Int) Minter {
	return Minter{
		Inflation:        inflation,
		AnnualProvisions: annualProvisions,
		Phase:            phase,
		StartPhaseBlock:  startPhaseBlock,
		TargetSupply:     targetSupply,
	}
}

// InitialMinter returns an initial Minter object with a given inflation value.
func InitialMinter(inflation sdk.Dec) Minter {
	return NewMinter(
		inflation,
		sdk.NewDec(0),
		0,
		0,
		sdk.NewInt(0),
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
// which uses an inflation rate of 13%.
func DefaultInitialMinter() Minter {
	return InitialMinter(
		sdk.NewDecWithPrec(13, 2),
	)
}

// validate minter
func ValidateMinter(minter Minter) error {
	if minter.Inflation.IsNegative() {
		return fmt.Errorf("mint parameter Inflation should be positive, is %s",
			minter.Inflation.String())
	}
	return nil
}

// PhaseInflationRate returns the inflation rate by phase.
func (m Minter) PhaseInflationRate(phase uint64) sdk.Dec {
	switch {
	case phase > 19:
		return sdk.ZeroDec()

	case phase == 1:
		return sdk.NewDecWithPrec(29, 2)

	case phase == 2:
		return sdk.NewDecWithPrec(21, 2)

	case phase == 3:
		return sdk.NewDecWithPrec(20, 2)

	default:
		// Phase4:  15%
		// Phase5:  14%
		// Phase6:  13%
		// ...
		// Phase18: 1%
		return sdk.NewDecWithPrec(18-int64(phase), 2)
	}
}

// NextPhase returns the new phase.
func (m Minter) NextPhase(ctx sdk.Context, params Params) uint64 {

	phase := m.Phase
	if phase == 0 {
		return 1
	}

	blocksPerYear := params.BlocksPerYear
	currentBlockNumber := uint64(ctx.BlockHeight())

	yearsSinceStart := (currentBlockNumber) / blocksPerYear

	// Calculate the number of times the token issuance rate has been halved
	halvings := (yearsSinceStart) / 3

	// Update the token issuance rate for each halving that has occurred
	for i := uint64(0); i < halvings; i++ {
		phase = phase + 1
	}

	return phase
}

// NextAnnualProvisions returns the annual provisions based on current total
// supply and inflation rate.
func (m Minter) NextAnnualProvisions(_ Params, totalSupply sdk.Int) sdk.Dec {
	return m.Inflation.MulInt(totalSupply)
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params, totalSupply sdk.Int) sdk.Coin {
	provisionAmt := m.AnnualProvisions.QuoInt(sdk.NewInt(int64(params.BlocksPerYear)))

	// Because of rounding, we might mint too many tokens in this phase, let's limit it
	futureSupply := totalSupply.Add(provisionAmt.TruncateInt())
	if futureSupply.GT(m.TargetSupply) {
		return sdk.NewCoin("soar", m.TargetSupply.Sub(totalSupply))
	}

	return sdk.NewCoin("soar", provisionAmt.TruncateInt())
}

func (m Minter) StakingRewardsPerBlock(ctx sdk.Context, params Params) sdk.Coins {

	// Calculates reward coin emissions for each reward type

	blocksPerYear := params.BlocksPerYear
	currentBlockNumber := uint64(ctx.BlockHeight())

	initialTokensPerYear := uint64(21850083350000) // staking rewards initial annual emission

	// Number of tokens issued per year by the total number of blocks produced per year
	tokensPerBlock := initialTokensPerYear / blocksPerYear

	// Calculate the number of years that have passed since the start of the token issuance
	yearsSinceStart := (currentBlockNumber) / blocksPerYear

	// Calculate the number of times the token issuance rate has been halved
	halvings := (yearsSinceStart) / 3

	// Update the token issuance rate for each halving that has occurred
	for i := uint64(0); i < halvings; i++ {
		initialTokensPerYear /= 2
		tokensPerBlock = initialTokensPerYear / blocksPerYear
	}

	return sdk.Coins{sdk.NewCoin("soar", sdk.NewIntFromUint64(tokensPerBlock))}
}

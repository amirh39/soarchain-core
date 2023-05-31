package errors

var (
	GetChallengerByType     = "[RunnerChallenge][GetChallengerByType] failed. Only registered challengers with v2n type can initiate this transaction."
	TargetEpoch             = "[RunnerChallenge][V2NRewardEmissionPerEpoch] failed. Couldn't calculate reward to emission for each reward."
	EpochDataNotFound       = "[RunnerChallenge][GetEpochData] failed. Couldn't find epoch data in the current context."
	EarnedTokenRewardsFloat = "[RunnerChallenge][V2NRewardCalculator] failed. Couldn't calcualte earned rewards."
	NetEarnings             = "[RunnerChallenge][ParseCoinNormalized] failed. Couldn't parse and normalize a cli input for one coin type, due to invalid or an empty string."
	EpochErr                = "[RunnerChallenge][UpdateEpochRewards] failed. Couldn't updat epoch reward."
	TotalEarnings           = "[RunnerChallenge][TotalEarnings] failed. Couldn't calculate totalEarnings with the given values."
	NoV2nBxAddrPubKeys      = "[RunnerChallenge][v2nBxAddrCount] failed. Couldn't find client pubkeys in the tx body."
	NotFoundAClient         = "[RunnerChallenge][GetClient] failed.  V2nBx client is not registered in the store."
	NotFoundAValidRunner    = "[RunnerChallenge][UpdateRunner] failed.  Couldn't find a valid runner."
)

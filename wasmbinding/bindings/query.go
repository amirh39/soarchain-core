package bindings

// SoarchainQuery contains soarchain custom queries.
type SoarchainQuery struct {
	ChallengerByIndex *ChallengerByIndex `json:"challenger_by_index,omitempty"`
}

type ChallengerByIndex struct {
	Index string `json:"index"`
}

type ChallengerByIndexResponse struct {
	Address string `json:"address"`
	Index   string `json:"index"`
	Score   string `json:"score"`
}

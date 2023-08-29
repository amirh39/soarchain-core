package bindings

// SoarchainQuery contains soarchain custom queries.
type SoarchainQuery struct {
	ClientByIndex *ClientByIndex `json:"client_by_index,omitempty"`
}

type ClientByIndex struct {
	Index string `json:"index"`
}

type ClientByIndexResponse struct {
	Address string `json:"address"`
	Index   string `json:"index"`
	Score   string `json:"score"`
}

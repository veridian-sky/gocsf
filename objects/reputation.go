package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Reputation The Reputation object describes the reputation/risk score of an entity (e.g.
// device, user, domain).
type Reputation struct {
	// The reputation score as reported by the event source.
	BaseScore float64 `json:"base_score,omitempty"`

	// The provider of the reputation information.
	Provider string `json:"provider,omitempty"`

	// The reputation score, normalized to the caption of the score_id value. In
	// the case of 'Other', it is defined by the event source.
	Score string `json:"score,omitempty"`

	// The normalized reputation score identifier.
	ScoreID int `json:"score_id,omitempty"`
}

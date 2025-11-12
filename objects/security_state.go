package objects

// Code generated from OCSF schema; DO NOT EDIT.

// SecurityState The Security State object describes the security related state of a managed
// entity.
type SecurityState struct {
	// The security state, normalized to the caption of the state_id value. In the
	// case of 'Other', it is defined by the source.
	State string `json:"state,omitempty"`

	// The security state of the managed entity.
	StateID int `json:"state_id,omitempty"`
}

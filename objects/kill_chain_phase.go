package objects

// Code generated from OCSF schema; DO NOT EDIT.

// KillChainPhase The Kill Chain Phase object represents a single phase of a cyber attack,
// including the initial reconnaissance and planning stages up to the final
// objective of the attacker. It provides a detailed description of each phase and
// its associated activities within the broader context of a cyber attack. See
// https://www.lockheedmartin.com/en-us/capabilities/cyber/cyber-kill-chain.html
// Cyber Kill Chain®.
type KillChainPhase struct {
	// The cyber kill chain phase.
	Phase string `json:"phase,omitempty"`

	// The cyber kill chain phase identifier.
	PhaseID int `json:"phase_id,omitempty"`
}

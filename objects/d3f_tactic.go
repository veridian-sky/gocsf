package objects

// Code generated from OCSF schema; DO NOT EDIT.

// D3fTactic The MITRE D3FEND™ Tactic object describes the tactic ID and/or name that is
// associated to an attack.
type D3fTactic struct {
	// The tactic name that is associated with the defensive technique. For
	// example: Isolate.
	Name string `json:"name,omitempty"`

	// The versioned permalink of the defensive tactic. For example:
	// https://d3fend.mitre.org/tactic/d3f:Isolate/.
	SrcURL string `json:"src_url,omitempty"`

	// The unique identifier of the defensive tactic.
	Uid string `json:"uid,omitempty"`
}

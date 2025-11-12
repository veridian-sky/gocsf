package objects

// Code generated from OCSF schema; DO NOT EDIT.

// D3fTechnique The MITRE D3FEND™ Technique object describes the leaf defensive technique ID
// and/or name associated to a countermeasure.
type D3fTechnique struct {
	// The name of the defensive technique. For example: IO Port Restriction.
	Name string `json:"name,omitempty"`

	// The versioned permalink of the defensive technique. For example:
	// https://d3fend.mitre.org/technique/d3f:IOPortRestriction/.
	SrcURL string `json:"src_url,omitempty"`

	// The unique identifier of the defensive technique. For example: D3-IOPR.
	Uid string `json:"uid,omitempty"`
}

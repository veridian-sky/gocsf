package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Attack The MITRE ATT&CK® & ATLAS™ object describes the tactic, technique,
// sub-technique & mitigation associated to an attack.
type Attack struct {
	// The Mitigation object describes the MITRE ATT&CK® or ATLAS™ Mitigation ID
	// and/or name that is associated to an attack.
	Mitigation *Mitigation `json:"mitigation,omitempty"`

	// The Sub-technique object describes the MITRE ATT&CK® or ATLAS™
	// Sub-technique ID and/or name associated to an attack.
	SubTechnique *SubTechnique `json:"sub_technique,omitempty"`

	// The Tactic object describes the MITRE ATT&CK® or ATLAS™ Tactic ID and/or
	// name that is associated to an attack.
	Tactic *Tactic `json:"tactic,omitempty"`

	// The Tactic object describes the tactic ID and/or tactic name that are
	// associated with the attack technique, as defined by
	// https://attack.mitre.org/wiki/ATT&CK_Matrix ATT&CK® Matrix.
	Tactics []*Tactic `json:"tactics,omitempty"`

	// The Technique object describes the MITRE ATT&CK® or ATLAS™ Technique ID
	// and/or name associated to an attack.
	Technique *Technique `json:"technique,omitempty"`

	// The ATT&CK® or ATLAS™ Matrix version.
	Version string `json:"version,omitempty"`
}

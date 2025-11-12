package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Technique The MITRE Technique object describes the ATT&CK® or ATLAS™ Technique ID
// and/or name associated to an attack.
type Technique struct {
	// The name of the attack technique. For example: Active Scanning or AI Model
	// Inference API Access.
	Name string `json:"name,omitempty"`

	// The versioned permalink of the attack technique. For example:
	// https://attack.mitre.org/versions/v14/techniques/T1595/.
	SrcURL string `json:"src_url,omitempty"`

	// The unique identifier of the attack technique. For example: T1595 or
	// AML.T0040.
	Uid string `json:"uid,omitempty"`
}

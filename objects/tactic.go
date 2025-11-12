package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Tactic The MITRE Tactic object describes the ATT&CK® or ATLAS™ Tactic ID and/or name
// that is associated to an attack.
type Tactic struct {
	// The Tactic name that is associated with the attack technique. For example:
	// Reconnaissance or ML Model Access.
	Name string `json:"name,omitempty"`

	// The versioned permalink of the Tactic. For example:
	// https://attack.mitre.org/versions/v14/tactics/TA0043/.
	SrcURL string `json:"src_url,omitempty"`

	// The Tactic ID that is associated with the attack technique. For example:
	// TA0043, or AML.TA0000.
	Uid string `json:"uid,omitempty"`
}

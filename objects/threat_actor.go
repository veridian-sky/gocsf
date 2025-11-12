package objects

// Code generated from OCSF schema; DO NOT EDIT.

// ThreatActor Threat actor is responsible for the observed malicious activity.
type ThreatActor struct {
	// The name of the threat actor.
	Name string `json:"name,omitempty"`

	// The classification of the threat actor based on their motivations,
	// capabilities, or affiliations. Common types include nation-state actors,
	// cybercriminal groups, hacktivists, or insider threats.
	Type string `json:"type,omitempty"`

	// The normalized datastore resource type identifier.
	TypeID int `json:"type_id,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Extension The OCSF Schema Extension object provides detailed information about the schema
// extension used to construct the event. The schema extensions are registered in
// the https://github.com/ocsf/ocsf-schema/blob/main/extensions.md extensions.md
// file.
type Extension struct {
	// The schema extension name. For example: dev.
	Name string `json:"name,omitempty"`

	// The schema extension unique identifier. For example: 999.
	Uid string `json:"uid,omitempty"`

	// The schema extension version. For example: 1.0.0-alpha.2.
	Version string `json:"version,omitempty"`
}

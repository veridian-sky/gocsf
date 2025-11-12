package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Scan The Scan object describes characteristics of a proactive scan.
type Scan struct {
	// The administrator-supplied or application-generated name of the scan. For
	// example: "Home office weekly user database scan", "Scan folders for
	// viruses", "Full system virus scan"
	Name string `json:"name,omitempty"`

	// The type of scan.
	Type string `json:"type,omitempty"`

	// The type id of the scan.
	TypeID int `json:"type_id,omitempty"`

	// The application-defined unique identifier assigned to an instance of a scan.
	Uid string `json:"uid,omitempty"`
}

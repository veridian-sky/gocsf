package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Sbom The Software Bill of Materials object describes characteristics of a generated
// SBOM.
type Sbom struct {
	// The time when the SBOM was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the SBOM was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The software package or library that is being discovered or inventoried by
	// an SBOM.
	Package *Package `json:"package,omitempty"`

	// Details about the upstream product that generated the SBOM e.g. cdxgen or
	// Syft.
	Product *Product `json:"product,omitempty"`

	// The list of software components used in the software package.
	SoftwareComponents []*SoftwareComponent `json:"software_components,omitempty"`

	// The type of SBOM, normalized to the caption of the type_id value. In the
	// case of 'Other', it is defined by the source.
	Type string `json:"type,omitempty"`

	// The type of SBOM.
	TypeID int `json:"type_id,omitempty"`

	// A unique identifier for the SBOM or the SBOM generation by a source tool,
	// such as the SPDX metadata.component.bom-ref.
	Uid string `json:"uid,omitempty"`

	// The specification (spec) version of the particular SBOM, e.g., 1.6.
	Version string `json:"version,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// SoftwareComponent The Software Component object describes characteristics of a software component
// within a software package.
type SoftwareComponent struct {
	// The author(s) who published the software component.
	Author string `json:"author,omitempty"`

	// Cryptographic hash to identify the binary instance of a software component.
	Hash *Fingerprint `json:"hash,omitempty"`

	// The software license applied to this component.
	License string `json:"license,omitempty"`

	// The software component name.
	Name string `json:"name,omitempty"`

	// The Package URL (PURL) to identify the software component. This is a URL
	// that uniquely identifies the component, including the component's name,
	// version, and type. The URL is used to locate and retrieve the component's
	// metadata and content.
	Purl string `json:"purl,omitempty"`

	// The package URL (PURL) of the component that this software component has a
	// relationship with.
	RelatedComponent string `json:"related_component,omitempty"`

	// The relationship between two software components, normalized to the caption
	// of the relationship_id value. In the case of 'Other', it is defined by the
	// source.
	Relationship string `json:"relationship,omitempty"`

	// The normalized identifier of the relationship between two software
	// components.
	RelationshipID int `json:"relationship_id,omitempty"`

	// The type of software component, normalized to the caption of the type_id
	// value. In the case of 'Other', it is defined by the source.
	Type string `json:"type,omitempty"`

	// The type of software component.
	TypeID int `json:"type_id,omitempty"`

	// The software component version.
	Version string `json:"version,omitempty"`
}

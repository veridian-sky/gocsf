package objects

// Code generated from OCSF schema; DO NOT EDIT.

// TlsExtension The TLS Extension object describes additional attributes that extend the base
// Transport Layer Security (TLS) object.
type TlsExtension struct {
	// The data contains information specific to the particular extension type.
	Data interface{} `json:"data,omitempty"`

	// The TLS extension type. For example: Server Name.
	Type string `json:"type,omitempty"`

	// The TLS extension type identifier. See
	// https://datatracker.ietf.org/doc/html/rfc8446#page-35 The Transport Layer
	// Security (TLS) extension page.
	TypeID int `json:"type_id,omitempty"`
}

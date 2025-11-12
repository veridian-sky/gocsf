package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Ja4Fingerprint The JA4+ fingerprint object provides detailed fingerprint information about
// various aspects of network traffic which is both machine and human readable.
type Ja4Fingerprint struct {
	// The 'a' section of the JA4 fingerprint.
	SectionA string `json:"section_a,omitempty"`

	// The 'b' section of the JA4 fingerprint.
	SectionB string `json:"section_b,omitempty"`

	// The 'c' section of the JA4 fingerprint.
	SectionC string `json:"section_c,omitempty"`

	// The 'd' section of the JA4 fingerprint.
	SectionD string `json:"section_d,omitempty"`

	// The JA4+ fingerprint type as defined by <a
	// href='https://blog.foxio.io/ja4+-network-fingerprinting target='_blank
	// FoxIO, normalized to the caption of 'type_id'. In the case of 'Other', it is
	// defined by the event source.
	Type string `json:"type,omitempty"`

	// The identifier of the JA4+ fingerprint type.
	TypeID int `json:"type_id,omitempty"`

	// The JA4+ fingerprint value.
	Value string `json:"value,omitempty"`
}

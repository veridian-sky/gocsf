package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Certificate The Digital Certificate, also known as a Public Key Certificate, object contains
// information about the ownership and usage of a public key. It serves as a means
// to establish trust in the authenticity and integrity of the public key and the
// associated entity.
type Certificate struct {
	// The time when the certificate was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the certificate was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The expiration time of the certificate.
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	// The expiration time of the certificate.
	ExpirationTimeDt string `json:"expiration_time_dt,omitempty"`

	// The fingerprint list of the certificate.
	Fingerprints []*Fingerprint `json:"fingerprints,omitempty"`

	// Denotes whether a digital certificate is self-signed or signed by a known
	// certificate authority (CA).
	IsSelfSigned bool `json:"is_self_signed,omitempty"`

	// The certificate issuer distinguished name.
	Issuer string `json:"issuer,omitempty"`

	// The list of subject alternative names that are secured by a specific
	// certificate.
	Sans []*San `json:"sans,omitempty"`

	// The serial number of the certificate used to create the digital signature.
	SerialNumber string `json:"serial_number,omitempty"`

	// The certificate subject distinguished name.
	Subject string `json:"subject,omitempty"`

	// The unique identifier of the certificate.
	Uid string `json:"uid,omitempty"`

	// The certificate version.
	Version string `json:"version,omitempty"`
}

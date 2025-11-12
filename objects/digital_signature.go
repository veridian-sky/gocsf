package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DigitalSignature The Digital Signature object contains information about the cryptographic
// mechanism used to verify the authenticity, integrity, and origin of the file or
// application.
type DigitalSignature struct {
	// The digital signature algorithm used to create the signature, normalized to
	// the caption of 'algorithm_id'. In the case of 'Other', it is defined by the
	// event source.
	Algorithm string `json:"algorithm,omitempty"`

	// The identifier of the normalized digital signature algorithm.
	AlgorithmID int `json:"algorithm_id,omitempty"`

	// The certificate object containing information about the digital certificate.
	Certificate *Certificate `json:"certificate,omitempty"`

	// The time when the digital signature was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the digital signature was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The developer ID on the certificate that signed the file.
	DeveloperUID string `json:"developer_uid,omitempty"`

	// The message digest attribute contains the fixed length message hash
	// representation and the corresponding hashing algorithm information.
	Digest *Fingerprint `json:"digest,omitempty"`

	// The digital signature state defines the signature state, normalized to the
	// caption of 'state_id'. In the case of 'Other', it is defined by the event
	// source.
	State string `json:"state,omitempty"`

	// The normalized identifier of the signature state.
	StateID int `json:"state_id,omitempty"`
}

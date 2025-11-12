package objects

// Code generated from OCSF schema; DO NOT EDIT.

// EncryptionDetails Details about the encryption methodology utilized.
type EncryptionDetails struct {
	// The encryption algorithm used, normalized to the caption of 'algorithm_id
	Algorithm string `json:"algorithm,omitempty"`

	// The encryption algorithm used.
	AlgorithmID int `json:"algorithm_id,omitempty"`

	// The length of the encryption key used.
	KeyLength int64 `json:"key_length,omitempty"`

	// The unique identifier of the key used for encryption. For example, AWS KMS
	// Key ARN.
	KeyUID string `json:"key_uid,omitempty"`

	// The type of the encryption used.
	Type string `json:"type,omitempty"`
}

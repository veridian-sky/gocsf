package objects

// Code generated from OCSF schema; DO NOT EDIT.

// AuthenticationToken The Authentication Token object represents standardized authentication tokens,
// tickets, or assertions that conform to established authentication protocols such
// as Kerberos, OIDC, and SAML. These tokens are issued by authentication servers
// and identity providers and carry protocol-specific metadata, lifecycle
// information, and security attributes defined by their respective specifications.
type AuthenticationToken struct {
	// The time that the authentication token was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time that the authentication token was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The encryption details of the authentication token.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`

	// The expiration time of the authentication token.
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	// The expiration time of the authentication token.
	ExpirationTimeDt string `json:"expiration_time_dt,omitempty"`

	// Indicates whether the authentication token is renewable.
	IsRenewable bool `json:"is_renewable,omitempty"`

	// A bitmask, either in hexadecimal or decimal form, which encodes various
	// attributes or permissions associated with a Kerberos ticket. These flags
	// delineate specific characteristics of the ticket, such as its renewability
	// or forwardability.
	KerberosFlags string `json:"kerberos_flags,omitempty"`

	// The type of the authentication token.
	Type string `json:"type,omitempty"`

	// The normalized authentication token type identifier.
	TypeID int `json:"type_id,omitempty"`
}

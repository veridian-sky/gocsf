package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Session The Session object describes details about an authenticated session. e.g.
// Session Creation Time, Session Issuer.
type Session struct {
	// The number of identical sessions spawned from the same source IP,
	// destination IP, application, and content/threat type seen over a period of
	// time.
	Count int64 `json:"count,omitempty"`

	// The time when the session was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the session was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The unique identifier of the user's credential. For example, AWS Access Key
	// ID.
	CredentialUID string `json:"credential_uid,omitempty"`

	// The reason which triggered the session expiration.
	ExpirationReason string `json:"expiration_reason,omitempty"`

	// The session expiration time.
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	// The session expiration time.
	ExpirationTimeDt string `json:"expiration_time_dt,omitempty"`

	// Indicates whether Multi Factor Authentication was used during
	// authentication.
	IsMfa bool `json:"is_mfa,omitempty"`

	// The indication of whether the session is remote.
	IsRemote bool `json:"is_remote,omitempty"`

	// The indication of whether the session is a VPN session.
	IsVpn bool `json:"is_vpn,omitempty"`

	// The identifier of the session issuer.
	Issuer string `json:"issuer,omitempty"`

	// The Pseudo Terminal associated with the session. Ex: the tty or pts value.
	Terminal string `json:"terminal,omitempty"`

	// The unique identifier of the session.
	Uid string `json:"uid,omitempty"`

	// The alternate unique identifier of the session. e.g. AWS ARN -
	// arn:aws:sts::123344444444:assumed-role/Admin/example-session.
	UidAlt string `json:"uid_alt,omitempty"`

	// The universally unique identifier of the session.
	Uuid string `json:"uuid,omitempty"`
}

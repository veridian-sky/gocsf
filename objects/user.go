package objects

// Code generated from OCSF schema; DO NOT EDIT.

// User The User object describes the characteristics of a user/person or a security
// principal.
type User struct {
	// The user's account or the account associated with the user.
	Account *Account `json:"account,omitempty"`

	// The unique identifier of the user's credential. For example, AWS Access Key
	// ID.
	CredentialUID string `json:"credential_uid,omitempty"`

	// The display name of the user, as reported by the product.
	DisplayName string `json:"display_name,omitempty"`

	// The domain where the user is defined. For example: the LDAP or Active
	// Directory domain.
	Domain string `json:"domain,omitempty"`

	// The user's primary email address.
	EmailAddr string `json:"email_addr,omitempty"`

	// The user's forwarding email address.
	ForwardAddr string `json:"forward_addr,omitempty"`

	// The full name of the user, as reported by the product.
	FullName string `json:"full_name,omitempty"`

	// The administrative groups to which the user belongs.
	Groups []*Group `json:"groups,omitempty"`

	// The user has a multi-factor or secondary-factor device assigned.
	HasMfa bool `json:"has_mfa,omitempty"`

	// The additional LDAP attributes that describe a person.
	LdapPerson *LdapPerson `json:"ldap_person,omitempty"`

	// The username. For example, janedoe1.
	Name interface{} `json:"name,omitempty"`

	// Organization and org unit related to the user.
	Org *Organization `json:"org,omitempty"`

	// The telephone number of the user.
	PhoneNumber string `json:"phone_number,omitempty"`

	// Details about the programmatic credential (API keys, access tokens,
	// certificates, etc) associated to the user.
	ProgrammaticCredentials []*ProgrammaticCredential `json:"programmatic_credentials,omitempty"`

	// The risk level, normalized to the caption of the risk_level_id value.
	RiskLevel string `json:"risk_level,omitempty"`

	// The normalized risk level id.
	RiskLevelID int `json:"risk_level_id,omitempty"`

	// The risk score as reported by the event source.
	RiskScore int64 `json:"risk_score,omitempty"`

	// The type of the user. For example, System, AWS IAM User, etc.
	Type string `json:"type,omitempty"`

	// The account type identifier.
	TypeID int `json:"type_id,omitempty"`

	// The unique user identifier. For example, the Windows user SID,
	// ActiveDirectory DN or AWS user ARN.
	Uid string `json:"uid,omitempty"`

	// The alternate user identifier. For example, the Active Directory user GUID
	// or AWS user Principal ID.
	UidAlt string `json:"uid_alt,omitempty"`
}

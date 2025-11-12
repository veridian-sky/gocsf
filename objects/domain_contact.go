package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DomainContact The contact information related to a domain registration, e.g., registrant,
// administrator, abuse, billing, or technical contact.
type DomainContact struct {
	// The user's primary email address.
	EmailAddr string `json:"email_addr,omitempty"`

	// Location details for the contract such as the city, state/province, country,
	// etc.
	Location *Location `json:"location,omitempty"`

	// The individual or organization name for the contact.
	Name string `json:"name,omitempty"`

	// The number associated with the phone.
	PhoneNumber string `json:"phone_number,omitempty"`

	// The Domain Contact type, normalized to the caption of the type_id value. In
	// the case of 'Other', it is defined by the source
	Type string `json:"type,omitempty"`

	// The normalized domain contact type ID.
	TypeID int `json:"type_id,omitempty"`

	// The unique identifier of the contact information, typically provided in
	// WHOIS information.
	Uid string `json:"uid,omitempty"`
}

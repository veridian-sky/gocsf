package objects

// Code generated from OCSF schema; DO NOT EDIT.

// AuthFactor An Authentication Factor object describes a category of methods used for
// identity verification in an authentication attempt.
type AuthFactor struct {
	// Device used to complete an authentication request.
	Device *Device `json:"device,omitempty"`

	// The email address used in an email-based authentication factor.
	EmailAddr string `json:"email_addr,omitempty"`

	// The type of authentication factor used in an authentication attempt.
	FactorType string `json:"factor_type,omitempty"`

	// The normalized identifier for the authentication factor.
	FactorTypeID int `json:"factor_type_id,omitempty"`

	// Whether the authentication factor is an HMAC-based One-time Password (HOTP).
	IsHotp bool `json:"is_hotp,omitempty"`

	// Whether the authentication factor is a Time-based One-time Password (TOTP).
	IsTotp bool `json:"is_totp,omitempty"`

	// The phone number used for a telephony-based authentication request.
	PhoneNumber string `json:"phone_number,omitempty"`

	// The name of provider for an authentication factor.
	Provider string `json:"provider,omitempty"`

	// The question(s) provided to user for a question-based authentication factor.
	SecurityQuestions []string `json:"security_questions,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Idp The Identity Provider object contains detailed information about a provider
// responsible for creating, maintaining, and managing identity information while
// offering authentication services to applications. An Identity Provider (IdP)
// serves as a trusted authority that verifies the identity of users and issues
// authentication tokens or assertions to enable secure access to applications or
// services.
type Idp struct {
	// The Authentication Factors object describes the different types of
	// Multi-Factor Authentication (MFA) methods and/or devices supported by the
	// Identity Provider.
	AuthFactors []*AuthFactor `json:"auth_factors,omitempty"`

	// The primary domain associated with the Identity Provider.
	Domain string `json:"domain,omitempty"`

	// The fingerprint of the X.509 certificate used by the Identity Provider.
	Fingerprint *Fingerprint `json:"fingerprint,omitempty"`

	// The Identity Provider enforces Multi Factor Authentication (MFA).
	HasMfa bool `json:"has_mfa,omitempty"`

	// The unique identifier (often a URL) used by the Identity Provider as its
	// issuer.
	Issuer string `json:"issuer,omitempty"`

	// The name of the Identity Provider.
	Name string `json:"name,omitempty"`

	// The supported protocol of the Identity Provider. E.g., SAML, OIDC, or
	// OAuth2.
	ProtocolName string `json:"protocol_name,omitempty"`

	// The System for Cross-domain Identity Management (SCIM) resource object
	// provides a structured set of attributes related to SCIM protocols used for
	// identity provisioning and management across cloud-based platforms. It
	// standardizes user and group provisioning details, enabling identity
	// synchronization and lifecycle management with compatible Identity Providers
	// (IdPs) and applications. SCIM is defined in
	// https://datatracker.ietf.org/doc/html/rfc7643 RFC-7634
	Scim *Scim `json:"scim,omitempty"`

	// The Single Sign-On (SSO) object provides a structure for normalizing SSO
	// attributes, configuration, and/or settings from Identity Providers.
	Sso *Sso `json:"sso,omitempty"`

	// The configuration state of the Identity Provider, normalized to the caption
	// of the state_id value. In the case of Other, it is defined by the event
	// source.
	State string `json:"state,omitempty"`

	// The normalized state ID of the Identity Provider to reflect its
	// configuration or activation status.
	StateID int `json:"state_id,omitempty"`

	// The tenant ID associated with the Identity Provider.
	TenantUID string `json:"tenant_uid,omitempty"`

	// The unique identifier of the Identity Provider.
	Uid string `json:"uid,omitempty"`

	// The URL for accessing the configuration or metadata of the Identity
	// Provider.
	UrlString string `json:"url_string,omitempty"`
}

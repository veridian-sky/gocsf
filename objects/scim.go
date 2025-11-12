package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Scim The System for Cross-domain Identity Management (SCIM) Configuration object
// provides a structured set of attributes related to SCIM protocols used for
// identity provisioning and management across cloud-based platforms. It
// standardizes user and group provisioning details, enabling identity
// synchronization and lifecycle management with compatible Identity Providers
// (IdPs) and applications. SCIM is defined in
// https://datatracker.ietf.org/doc/html/rfc7643 RFC-7634
type Scim struct {
	// The authorization protocol as defined by the caption of auth_protocol_id. In
	// the case of Other, it is defined by the event source.
	AuthProtocol string `json:"auth_protocol,omitempty"`

	// The normalized identifier of the authorization protocol used by the SCIM
	// resource.
	AuthProtocolID int `json:"auth_protocol_id,omitempty"`

	// When the SCIM resource was added to the service provider.
	CreatedTime int64 `json:"created_time,omitempty"`

	// When the SCIM resource was added to the service provider.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// Message or code associated with the last encountered error.
	ErrorMessage string `json:"error_message,omitempty"`

	// Indicates whether the SCIM resource is configured to provision groups,
	// automatically or otherwise.
	IsGroupProvisioningEnabled bool `json:"is_group_provisioning_enabled,omitempty"`

	// Indicates whether the SCIM resource is configured to provision users,
	// automatically or otherwise.
	IsUserProvisioningEnabled bool `json:"is_user_provisioning_enabled,omitempty"`

	// Timestamp of the most recent successful synchronization.
	LastRunTime int64 `json:"last_run_time,omitempty"`

	// Timestamp of the most recent successful synchronization.
	LastRunTimeDt string `json:"last_run_time_dt,omitempty"`

	// The most recent time when the SCIM resource was updated at the service
	// provider.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The most recent time when the SCIM resource was updated at the service
	// provider.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The name of the SCIM resource.
	Name string `json:"name,omitempty"`

	// The supported protocol for the SCIM resource. E.g., SAML, OIDC, or OAuth2.
	ProtocolName string `json:"protocol_name,omitempty"`

	// Maximum number of requests allowed by the SCIM resource within a specified
	// time frame to avoid throttling.
	RateLimit int64 `json:"rate_limit,omitempty"`

	// SCIM provides a schema for representing groups, identified using the
	// following schema URI: urn:ietf:params:scim:schemas:core:2.0:Group as defined
	// in https://datatracker.ietf.org/doc/html/rfc7643 RFC-7634. This attribute
	// will capture key-value pairs for the scheme implemented in a SCIM resource.
	ScimGroupSchema interface{} `json:"scim_group_schema,omitempty"`

	// SCIM provides a resource type for user resources. The core schema for user
	// is identified using the following schema URI:
	// urn:ietf:params:scim:schemas:core:2.0:User as defined in
	// https://datatracker.ietf.org/doc/html/rfc7643 RFC-7634. his attribute will
	// capture key-value pairs for the scheme implemented in a SCIM resource. This
	// object is inclusive of both the basic and Enterprise User Schema Extension.
	ScimUserSchema interface{} `json:"scim_user_schema,omitempty"`

	// The provisioning state of the SCIM resource, normalized to the caption of
	// the state_id value. In the case of Other, it is defined by the event source.
	State string `json:"state,omitempty"`

	// The normalized state ID of the SCIM resource to reflect its activation
	// status.
	StateID int `json:"state_id,omitempty"`

	// A unique identifier for a SCIM resource as defined by the service provider.
	Uid string `json:"uid,omitempty"`

	// A String that is an identifier for the resource as defined by the
	// provisioning client. The externalId may simplify identification of a
	// resource between the provisioning client and the service provider by
	// allowing the client to use a filter to locate the resource with an
	// identifier from the provisioning domain, obviating the need to store a local
	// mapping between the provisioning domain's identifier of the resource and the
	// identifier used by the service provider.
	UidAlt string `json:"uid_alt,omitempty"`

	// The primary URL for SCIM API requests.
	UrlString string `json:"url_string,omitempty"`

	// Name of the vendor or service provider implementing SCIM. E.g., Okta, Auth0,
	// Microsoft.
	VendorName string `json:"vendor_name,omitempty"`

	// SCIM protocol version supported e.g., SCIM 2.0.
	Version string `json:"version,omitempty"`
}

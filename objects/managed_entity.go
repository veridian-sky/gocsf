package objects

// Code generated from OCSF schema; DO NOT EDIT.

// ManagedEntity The Managed Entity object describes the type and version of an entity, such as a
// user, device, or policy. For types in the type_id enum list, an associated
// attribute should be populated. If the type of entity is not in the type_id list,
// information can be put into the data attribute, type_id should be 'Other' and
// the type attribute should label the entity type.
type ManagedEntity struct {
	// The managed entity content as a JSON object.
	Data interface{} `json:"data,omitempty"`

	// An addressable device, computer system or host.
	Device *Device `json:"device,omitempty"`

	// The email object.
	Email *Email `json:"email,omitempty"`

	// The group object associated with an entity such as user, policy, or rule.
	Group *Group `json:"group,omitempty"`

	// The detailed geographical location usually associated with an IP address.
	Location *Location `json:"location,omitempty"`

	// The name of the managed entity. It should match the name of the specific
	// entity object's name if populated, or the name of the managed entity if the
	// type_id is 'Other'.
	Name string `json:"name,omitempty"`

	// Organization and org unit relevant to the event or object.
	Org *Organization `json:"org,omitempty"`

	// Describes details of a managed policy.
	Policy *Policy `json:"policy,omitempty"`

	// The managed entity type. For example: Policy, User, Organization, Device.
	Type string `json:"type,omitempty"`

	// The type of the Managed Entity. It is recommended to also populate the type
	// attribute with the associated label, or the source specific name if Other.
	TypeID int `json:"type_id,omitempty"`

	// The identifier of the managed entity. It should match the uid of the
	// specific entity's object UID if populated, or the source specific ID if the
	// type_id is 'Other'.
	Uid string `json:"uid,omitempty"`

	// The user that pertains to the event or object.
	User *User `json:"user,omitempty"`

	// The version of the managed entity. For example: 1.2.3.
	Version string `json:"version,omitempty"`
}

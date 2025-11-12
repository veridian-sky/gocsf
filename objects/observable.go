package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Observable The observable object is a pivot element that contains related information found
// in many places in the event.
type Observable struct {
	// The full name of the observable attribute. The name is a pointer/reference
	// to an attribute within the OCSF event data. For example: file.name.
	Name string `json:"name,omitempty"`

	// Contains the original and normalized reputation scores.
	Reputation *Reputation `json:"reputation,omitempty"`

	// The observable value type name.
	Type string `json:"type,omitempty"`

	// The observable value type identifier.
	TypeID int `json:"type_id,omitempty"`

	// The value associated with the observable attribute. The meaning of the value
	// depends on the observable type.<br/>If the name refers to a scalar
	// attribute, then the value is the value of the attribute.<br/>If the name
	// refers to an object attribute, then the value is not populated.
	Value string `json:"value,omitempty"`
}

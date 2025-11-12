package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Rule The Rule object describes characteristics of a rule associated with a policy or
// an event.
type Rule struct {
	// The rule category.
	Category string `json:"category,omitempty"`

	// The description of the rule that generated the event.
	Desc string `json:"desc,omitempty"`

	// The name of the rule that generated the event.
	Name string `json:"name,omitempty"`

	// The rule type.
	Type string `json:"type,omitempty"`

	// The unique identifier of the rule that generated the event.
	Uid string `json:"uid,omitempty"`

	// The rule version. For example: 1.1.
	Version string `json:"version,omitempty"`
}

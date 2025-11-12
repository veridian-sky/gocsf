package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Service The Service object describes characteristics of a service, e.g. AWS EC2.
type Service struct {
	// The list of labels associated with the service.
	Labels []string `json:"labels,omitempty"`

	// The name of the service.
	Name string `json:"name,omitempty"`

	// The list of tags; {key:value} pairs associated to the service.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The unique identifier of the service.
	Uid string `json:"uid,omitempty"`

	// The version of the service.
	Version string `json:"version,omitempty"`
}

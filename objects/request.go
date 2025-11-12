package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Request The Request Elements object describes characteristics of an API request.
type Request struct {
	// When working with containerized applications, the set of containers which
	// write to the standard the output of a particular logging driver. For
	// example, this may be the set of containers involved in handling api requests
	// and responses for a containerized application.
	Containers []*Container `json:"containers,omitempty"`

	// The additional data that is associated with the api request.
	Data interface{} `json:"data,omitempty"`

	// The communication flags that are associated with the api request.
	Flags []string `json:"flags,omitempty"`

	// The unique request identifier.
	Uid string `json:"uid,omitempty"`
}

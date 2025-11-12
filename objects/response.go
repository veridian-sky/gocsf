package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Response The Response Elements object describes characteristics of an API response.
type Response struct {
	// The numeric response sent to a request.
	Code int64 `json:"code,omitempty"`

	// When working with containerized applications, the set of containers which
	// write to the standard the output of a particular logging driver. For
	// example, this may be the set of containers involved in handling api requests
	// and responses for a containerized application.
	Containers []*Container `json:"containers,omitempty"`

	// The additional data that is associated with the api response.
	Data interface{} `json:"data,omitempty"`

	// Error Code
	Error string `json:"error,omitempty"`

	// Error Message
	ErrorMessage string `json:"error_message,omitempty"`

	// The communication flags that are associated with the api response.
	Flags []string `json:"flags,omitempty"`

	// The description of the event/finding, as defined by the source.
	Message string `json:"message,omitempty"`
}

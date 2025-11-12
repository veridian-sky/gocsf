package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Api The API, or Application Programming Interface, object represents information
// pertaining to an API request and response.
type Api struct {
	// The information pertaining to the API group.
	Group *Group `json:"group,omitempty"`

	// Verb/Operation associated with the request
	Operation string `json:"operation,omitempty"`

	// Details pertaining to the API request.
	Request *Request `json:"request,omitempty"`

	// Details pertaining to the API response.
	Response *Response `json:"response,omitempty"`

	// The information pertaining to the API service.
	Service *Service `json:"service,omitempty"`

	// The version of the API service.
	Version string `json:"version,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// EndpointConnection The Endpoint Connection object contains information detailing a connection
// attempt to an endpoint.
type EndpointConnection struct {
	// A numerical response status code providing details about the connection.
	Code int64 `json:"code,omitempty"`

	// Provides characteristics of the network endpoint.
	NetworkEndpoint *NetworkEndpoint `json:"network_endpoint,omitempty"`
}

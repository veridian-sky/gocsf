package objects

// Code generated from OCSF schema; DO NOT EDIT.

// EnvironmentVariable An environment variable.
type EnvironmentVariable struct {
	// The name of the environment variable.
	Name string `json:"name,omitempty"`

	// The value of the environment variable.
	Value string `json:"value,omitempty"`
}

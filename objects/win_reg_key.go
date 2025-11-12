package objects

// Code generated from OCSF schema; DO NOT EDIT.

// WinRegKey The registry key object describes a Windows registry key.
type WinRegKey struct {
	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`

	// The time when the registry key was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the registry key was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The full path to the registry key.
	Path string `json:"path,omitempty"`

	// The security descriptor of the registry key.
	SecurityDescriptor string `json:"security_descriptor,omitempty"`
}

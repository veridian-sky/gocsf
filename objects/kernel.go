package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Kernel The Kernel Resource object provides information about a specific kernel
// resource, including its name and type. It describes essential attributes
// associated with a resource managed by the kernel of an operating system.
type Kernel struct {
	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`

	// The name of the kernel resource.
	Name string `json:"name,omitempty"`

	// The full path of the kernel resource.
	Path interface{} `json:"path,omitempty"`

	// The system call that was invoked.
	SystemCall string `json:"system_call,omitempty"`

	// The type of the kernel resource.
	Type string `json:"type,omitempty"`

	// The type of the kernel resource.
	TypeID int `json:"type_id,omitempty"`
}

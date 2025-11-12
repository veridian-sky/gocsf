package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Module The Module object describes the load attributes of a module.
type Module struct {
	// The memory address where the module was loaded.
	BaseAddress string `json:"base_address,omitempty"`

	// The module file object.
	File *File `json:"file,omitempty"`

	// The entry-point function of the module. The system calls the entry-point
	// function whenever a process or thread loads or unloads the module.
	FunctionName string `json:"function_name,omitempty"`

	// The load type, normalized to the caption of the load_type_id value. In the
	// case of 'Other', it is defined by the event source.
	LoadType string `json:"load_type,omitempty"`

	// The normalized identifier for how the module was loaded in memory.
	LoadTypeID int `json:"load_type_id,omitempty"`

	// The start address of the execution.
	StartAddress string `json:"start_address,omitempty"`

	// The module type.
	Type string `json:"type,omitempty"`
}

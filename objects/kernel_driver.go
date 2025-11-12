package objects

// Code generated from OCSF schema; DO NOT EDIT.

// KernelDriver The Kernel Extension object describes a kernel driver that has been loaded or
// unloaded into the operating system (OS) kernel.
type KernelDriver struct {
	// The driver/extension file object.
	File *File `json:"file,omitempty"`
}

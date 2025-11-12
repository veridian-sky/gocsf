package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Image The Image object provides a description of a specific Virtual Machine (VM) or
// Container image.
type Image struct {
	// The list of labels associated to the image.
	Labels []string `json:"labels,omitempty"`

	// The image name. For example: elixir.
	Name string `json:"name,omitempty"`

	// The full path to the image file.
	Path interface{} `json:"path,omitempty"`

	// The image tag. For example: 1.11-alpine.
	Tag string `json:"tag,omitempty"`

	// The list of tags; {key:value} pairs associated to the image.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The unique image ID. For example: 77af4d6b9913.
	Uid string `json:"uid,omitempty"`
}

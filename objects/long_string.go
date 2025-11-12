package objects

// Code generated from OCSF schema; DO NOT EDIT.

// LongString This object is a used to capture strings which may be truncated by a security
// product due to their length.
type LongString struct {
	// Indicates that value has been truncated. May be omitted if truncation has
	// not occurred.
	IsTruncated bool `json:"is_truncated,omitempty"`

	// The size in bytes of the string represented by value before truncation.
	// Should be omitted if truncation has not occurred.
	UntruncatedSize int64 `json:"untruncated_size,omitempty"`

	// The string value, truncated if is_truncated is true.
	Value string `json:"value,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// AffectedCode The Affected Code object describes details about a code block identified as
// vulnerable.
type AffectedCode struct {
	// The column number of the last part of the assessed code identified as
	// vulnerable.
	EndColumn int64 `json:"end_column,omitempty"`

	// The line number of the last line of code block identified as vulnerable.
	EndLine int64 `json:"end_line,omitempty"`

	// Details about the file that contains the affected code block.
	File *File `json:"file,omitempty"`

	// Details about the user that owns the affected file.
	Owner *User `json:"owner,omitempty"`

	// Describes the recommended remediation steps to address identified issue(s).
	Remediation *Remediation `json:"remediation,omitempty"`

	// Details about the specific rule, e.g., those defined as part of a larger
	// policy, that triggered the finding.
	Rule *Rule `json:"rule,omitempty"`

	// The column number of the first part of the assessed code identified as
	// vulnerable.
	StartColumn int64 `json:"start_column,omitempty"`

	// The line number of the first line of code block identified as vulnerable.
	StartLine int64 `json:"start_line,omitempty"`
}

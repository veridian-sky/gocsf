package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Script The Script object describes a script or command that can be executed by a shell,
// script engine, or interpreter. Examples include Bash, JavsScript, PowerShell,
// Python, VBScript, etc. Note that the term <em>script</em> here denotes not only
// a script contained within a file but also a script or command typed
// interactively by a user, supplied on the command line, or provided by some other
// file-less mechanism.
type Script struct {
	// Present if this script is associated with a file. Not present in the case of
	// a file-less script.
	File *File `json:"file,omitempty"`

	// An array of the script's cryptographic hashes. Note that these hashes are
	// calculated on the script in its original encoding, and not on the normalized
	// UTF-8 encoding found in the script_content attribute.
	Hashes []*Fingerprint `json:"hashes,omitempty"`

	// Unique identifier for the script or macro, independent of the containing
	// file, used for tracking, auditing, and security analysis.
	Name string `json:"name,omitempty"`

	// This attribute relates a sub-script to a parent script having the matching
	// uid attribute. In the case of PowerShell, sub-script execution can be
	// identified by matching the activity correlation ID of the raw ETW events
	// provided by the OS.
	ParentUID string `json:"parent_uid,omitempty"`

	// The script content, normalized to UTF-8 encoding irrespective of its
	// original encoding. When emitting this attribute, it may be appropriate to
	// truncate large scripts. When consuming this attribute, large scripts should
	// be anticipated.
	ScriptContent *LongString `json:"script_content,omitempty"`

	// The script type, normalized to the caption of the type_id value. In the case
	// of 'Other', it is defined by the event source.
	Type string `json:"type,omitempty"`

	// The normalized script type ID.
	TypeID int `json:"type_id,omitempty"`

	// Some script engines assign a unique ID to each individual execution of a
	// given script. This attribute captures that unique ID. In the case of
	// PowerShell, the unique ID corresponds to the ScriptBlockId in the raw ETW
	// events provided by the OS.
	Uid string `json:"uid,omitempty"`
}

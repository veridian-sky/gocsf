package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Cwe The CWE object represents a weakness in a software system that can be exploited
// by a threat actor to perform an attack. The CWE object is based on the
// https://cwe.mitre.org/ Common Weakness Enumeration (CWE) catalog.
type Cwe struct {
	// The caption assigned to the Common Weakness Enumeration unique identifier.
	Caption string `json:"caption,omitempty"`

	// URL pointing to the CWE Specification. For more information see
	// https://cwe.mitre.org/ CWE.
	SrcURL string `json:"src_url,omitempty"`

	// The Common Weakness Enumeration unique number assigned to a specific
	// weakness. A CWE Identifier begins "CWE" followed by a sequence of digits
	// that acts as a unique identifier. For example: CWE-123.
	Uid string `json:"uid,omitempty"`
}

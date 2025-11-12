package objects

// Code generated from OCSF schema; DO NOT EDIT.

// CisCsc The CIS Critical Security Control (CSC) contains information as defined by the
// Center for Internet Security Critical Security Control
// https://www.cisecurity.org/controls (CIS CSC). Prioritized set of actions to
// protect your organization and data from cyber-attack vectors.
type CisCsc struct {
	// A Control is prescriptive, prioritized, and simplified set of best practices
	// that one can use to strengthen their cybersecurity posture. e.g. AWS
	// SecurityHub Controls, CIS Controls.
	Control string `json:"control,omitempty"`

	// The CIS critical security control version.
	Version string `json:"version,omitempty"`
}

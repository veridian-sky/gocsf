package objects

// Code generated from OCSF schema; DO NOT EDIT.

// CisBenchmarkResult The CIS Benchmark Result object contains information as defined by the Center
// for Internet Security (https://www.cisecurity.org/cis-benchmarks/ CIS) benchmark
// result. CIS Benchmarks are a collection of best practices for securely
// configuring IT systems, software, networks, and cloud infrastructure.
type CisBenchmarkResult struct {
	// The CIS benchmark description.
	Desc string `json:"desc,omitempty"`

	// The CIS benchmark name.
	Name string `json:"name,omitempty"`

	// Describes the recommended remediation steps to address identified issue(s).
	Remediation *Remediation `json:"remediation,omitempty"`

	// The CIS benchmark rule.
	Rule *Rule `json:"rule,omitempty"`
}

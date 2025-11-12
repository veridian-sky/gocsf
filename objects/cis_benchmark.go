package objects

// Code generated from OCSF schema; DO NOT EDIT.

// CisBenchmark The CIS Benchmark object describes best practices for securely configuring IT
// systems, software, networks, and cloud infrastructure as defined by the
// https://www.cisecurity.org/cis-benchmarks/ Center for Internet Security. See
// also https://www.cisecurity.org/insights/blog/getting-to-know-the-cis-benchmarks
// Getting to Know the CIS Benchmarks.
type CisBenchmark struct {
	// The CIS Critical Security Controls is a prioritized set of actions to
	// protect your organization and data from cyber-attack vectors.
	CisControls []*CisControl `json:"cis_controls,omitempty"`

	// The CIS Benchmark description. For example: <i>The cramfs filesystem type is
	// a compressed read-only Linux filesystem embedded in small footprint systems.
	// A cramfs image can be used without having to first decompress the image.</i>
	Desc string `json:"desc,omitempty"`

	// The CIS Benchmark name. For example: <i>Ensure mounting of cramfs
	// filesystems is disabled.</i>
	Name string `json:"name,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Cvss The Common Vulnerability Scoring System (https://www.first.org/cvss/ CVSS)
// object provides a way to capture the principal characteristics of a
// vulnerability and produce a numerical score reflecting its severity.
type Cvss struct {
	// The CVSS base score. For example: 9.1.
	BaseScore float64 `json:"base_score,omitempty"`

	// The CVSS depth represents a depth of the equation used to calculate CVSS
	// score.
	Depth int `json:"depth,omitempty"`

	// The Common Vulnerability Scoring System metrics. This attribute contains
	// information on the CVE's impact. If the CVE has been analyzed, this
	// attribute will contain any CVSSv2 or CVSSv3 information associated with the
	// vulnerability. For example: { {"Access Vector", "Network"}, {"Access
	// Complexity", "Low"}, ...}.
	Metrics []*Metric `json:"metrics,omitempty"`

	// The CVSS overall score, impacted by base, temporal, and environmental
	// metrics. For example: 9.1.
	OverallScore float64 `json:"overall_score,omitempty"`

	// <p>The Common Vulnerability Scoring System (CVSS) Qualitative Severity
	// Rating. A textual representation of the numeric score.</p>CVSS
	// v2.0<ul><li>Low (0.0 – 3.9)</li><li>Medium (4.0 – 6.9)</li><li>High (7.0
	// – 10.0)</li></ul></p>CVSS v3.0<ul><li>None (0.0)</li><li>Low (0.1 -
	// 3.9)</li><li>Medium (4.0 - 6.9)</li><li>High (7.0 - 8.9)</li><li>Critical
	// (9.0 - 10.0)</li></ul>
	Severity string `json:"severity,omitempty"`

	// The source URL for the CVSS score. For example:
	// https://nvd.nist.gov/vuln/detail/CVE-2021-44228
	SrcURL string `json:"src_url,omitempty"`

	// The CVSS vector string is a text representation of a set of CVSS metrics. It
	// is commonly used to record or transfer CVSS metric information in a concise
	// form. For example: 3.1/AV:L/AC:L/PR:L/UI:N/S:U/C:H/I:N/A:H.
	VectorString string `json:"vector_string,omitempty"`

	// The vendor that provided the CVSS score. For example: NVD, REDHAT etc.
	VendorName string `json:"vendor_name,omitempty"`

	// The CVSS version. For example: 3.1.
	Version string `json:"version,omitempty"`
}

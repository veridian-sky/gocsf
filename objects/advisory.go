package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Advisory The Advisory object represents publicly disclosed cybersecurity vulnerabilities
// defined in a Security advisory. e.g. Microsoft KB Article, Apple Security
// Advisory, or a GitHub Security Advisory (GHSA)
type Advisory struct {
	// The average time to patch.
	AvgTimespan *Timespan `json:"avg_timespan,omitempty"`

	// The Advisory bulletin identifier.
	Bulletin string `json:"bulletin,omitempty"`

	// The vendors classification of the Advisory.
	Classification string `json:"classification,omitempty"`

	// The time when the Advisory record was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the Advisory record was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// A brief description of the Advisory Record.
	Desc string `json:"desc,omitempty"`

	// The install state of the Advisory.
	InstallState string `json:"install_state,omitempty"`

	// The normalized install state ID of the Advisory.
	InstallStateID int `json:"install_state_id,omitempty"`

	// The Advisory has been replaced by another.
	IsSuperseded bool `json:"is_superseded,omitempty"`

	// The time when the Advisory record was last updated.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the Advisory record was last updated.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The operating system the Advisory applies to.
	Os *Os `json:"os,omitempty"`

	// The product where the vulnerability was discovered.
	Product *Product `json:"product,omitempty"`

	// A list of reference URLs with additional information about the
	// vulnerabilities disclosed in the Advisory.
	References []string `json:"references,omitempty"`

	// A list of Common Vulnerabilities and Exposures https://cve.mitre.org/ (CVE)
	// identifiers related to the vulnerabilities disclosed in the Advisory.
	RelatedCves []*Cve `json:"related_cves,omitempty"`

	// A list of Common Weakness Enumeration https://cwe.mitre.org/ (CWE)
	// identifiers related to the vulnerabilities disclosed in the Advisory.
	RelatedCwes []*Cwe `json:"related_cwes,omitempty"`

	// The size in bytes for the Advisory. Usually populated for a KB Article
	// patch.
	Size int64 `json:"size,omitempty"`

	// The Advisory link from the source vendor.
	SrcURL string `json:"src_url,omitempty"`

	// A title or a brief phrase summarizing the Advisory.
	Title string `json:"title,omitempty"`

	// The unique identifier assigned to the advisory or disclosed vulnerability,
	// e.g, GHSA-5mrr-rgp6-x4gr.
	Uid string `json:"uid,omitempty"`
}

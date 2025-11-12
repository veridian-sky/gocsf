package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Cve The Common Vulnerabilities and Exposures (CVE) object represents publicly
// disclosed cybersecurity vulnerabilities defined in CVE Program catalog
// (https://cve.mitre.org/ CVE). There is one CVE Record for each vulnerability in
// the catalog.
type Cve struct {
	// The Record Creation Date identifies when the CVE ID was issued to a CVE
	// Numbering Authority (CNA) or the CVE Record was published on the CVE List.
	// Note that the Record Creation Date does not necessarily indicate when this
	// vulnerability was discovered, shared with the affected vendor, publicly
	// disclosed, or updated in CVE.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The Record Creation Date identifies when the CVE ID was issued to a CVE
	// Numbering Authority (CNA) or the CVE Record was published on the CVE List.
	// Note that the Record Creation Date does not necessarily indicate when this
	// vulnerability was discovered, shared with the affected vendor, publicly
	// disclosed, or updated in CVE.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The CVSS object details Common Vulnerability Scoring System
	// (https://www.first.org/cvss/ CVSS) scores from the advisory that are related
	// to the vulnerability.
	Cvss []*Cvss `json:"cvss,omitempty"`

	// The CWE object represents a weakness in a software system that can be
	// exploited by a threat actor to perform an attack. The CWE object is based on
	// the https://cwe.mitre.org/ Common Weakness Enumeration (CWE) catalog.
	Cwe *Cwe `json:"cwe,omitempty"`

	// The https://cwe.mitre.org/ Common Weakness Enumeration (CWE) unique
	// identifier. For example: CWE-787.
	CweUID string `json:"cwe_uid,omitempty"`

	// Common Weakness Enumeration (CWE) definition URL. For example:
	// https://cwe.mitre.org/data/definitions/787.html.
	CweURL string `json:"cwe_url,omitempty"`

	// A brief description of the CVE Record.
	Desc string `json:"desc,omitempty"`

	// The Exploit Prediction Scoring System (EPSS) object describes the estimated
	// probability a vulnerability will be exploited. EPSS is a community-driven
	// effort to combine descriptive information about vulnerabilities (CVEs) with
	// evidence of actual exploitation in-the-wild. (https://www.first.org/epss/
	// EPSS).
	Epss *Epss `json:"epss,omitempty"`

	// The Record Modified Date identifies when the CVE record was last updated.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The Record Modified Date identifies when the CVE record was last updated.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The product where the vulnerability was discovered.
	Product *Product `json:"product,omitempty"`

	// A list of reference URLs with additional information about the CVE Record.
	References []string `json:"references,omitempty"`

	// Describes the Common Weakness Enumeration https://cwe.mitre.org/ (CWE)
	// details related to the CVE Record.
	RelatedCwes []*Cwe `json:"related_cwes,omitempty"`

	// A title or a brief phrase summarizing the CVE record.
	Title string `json:"title,omitempty"`

	// <p>The vulnerability type as selected from a large dropdown menu during CVE
	// refinement.</p>Most frequently used vulnerability types are: DoS, Code
	// Execution, Overflow, Memory Corruption, Sql Injection, XSS, Directory
	// Traversal, Http Response Splitting, Bypass something, Gain Information, Gain
	// Privileges, CSRF, File Inclusion. For more information see
	// https://www.cvedetails.com/vulnerabilities-by-types.php Vulnerabilities By
	// Type distributions.
	Type string `json:"type,omitempty"`

	// The Common Vulnerabilities and Exposures unique number assigned to a
	// specific computer vulnerability. A CVE Identifier begins with 4 digits
	// representing the year followed by a sequence of digits that acts as a unique
	// identifier. For example: CVE-2021-12345.
	Uid string `json:"uid,omitempty"`
}

// Code generated from OCSF schema. DO NOT EDIT.
package gocsf

// Compliance Types

// CisBenchmark The CIS Benchmark object describes best practices for securely configuring IT systems, software, networks, and cloud infrastructure as defined by the <a target='_blank' href='https://www.cisecurity.org/cis-benchmarks/'>Center for Internet Security</a>. See also <a target='_blank' href='https://www.cisecurity.org/insights/blog/getting-to-know-the-cis-benchmarks'>Getting to Know the CIS Benchmarks</a>.
type CisBenchmark struct {
	// The CIS Critical Security Controls is a prioritized set of actions to protect your organization and data from cyber-attack vectors.
	CisControls []*CisControl `json:"cis_controls"`
	// The CIS Benchmark description. For example: <i>The cramfs filesystem type is a compressed read-only Linux filesystem embedded in small footprint systems. A cramfs image can be used without having to first decompress the image.</i>
	Desc string `json:"desc,omitempty"`
	// The CIS Benchmark name. For example: <i>Ensure mounting of cramfs filesystems is disabled.</i>
	Name string `json:"name"`
}

// CisBenchmarkResult The CIS Benchmark Result object contains information as defined by the Center for Internet Security (<a target='_blank' href='https://www.cisecurity.org/cis-benchmarks/'>CIS</a>) benchmark result. CIS Benchmarks are a collection of best practices for securely configuring IT systems, software, networks, and cloud infrastructure.
type CisBenchmarkResult struct {
	// The CIS benchmark description.
	Desc string `json:"desc,omitempty"`
	// The CIS benchmark name.
	Name string `json:"name"`
	// Describes the recommended remediation steps to address identified issue(s).
	Remediation *Remediation `json:"remediation,omitempty"`
	// The CIS benchmark rule.
	Rule *Rule `json:"rule,omitempty"`
}

// CisControl The CIS Control (aka Critical Security Control) object describes a prioritized set of actions to protect your organization and data from cyber-attack vectors. The <a target='_blank' href='https://www.cisecurity.org/controls'>CIS Controls</a> are defined by the Center for Internet Security.
type CisControl struct {
	// The CIS Control description. For example: <i>Uninstall or disable unnecessary services on enterprise assets and software, such as an unused file sharing service, web application module, or service function.</i>
	Desc string `json:"desc,omitempty"`
	// The CIS Control name. For example: <i>4.8 Uninstall or Disable Unnecessary Services on Enterprise Assets and Software.</i>
	Name string `json:"name"`
	// The CIS Control version. For example: <i>v8</i>.
	Version string `json:"version"`
}

// CisCsc The CIS Critical Security Control (CSC) contains information as defined by the Center for Internet Security Critical Security Control <a target='_blank' href='https://www.cisecurity.org/controls'>(CIS CSC)</a>. Prioritized set of actions to protect your organization and data from cyber-attack vectors.
type CisCsc struct {
	// A Control is prescriptive, prioritized, and simplified set of best practices that one can use to strengthen their cybersecurity posture. e.g. AWS SecurityHub Controls, CIS Controls.
	Control string `json:"control"`
	// The CIS critical security control version.
	Version string `json:"version"`
}

// Compliance The Compliance object contains information about Industry and Regulatory Framework standards, controls and requirements.
type Compliance struct {
	// A list of sources of information or tools that help organizations understand, interpret, and implement compliance standards. They provide guidance, best practices, and examples.
	ComplianceReferences []*KbArticle `json:"compliance_references,omitempty"`
	// A list of established guidelines or criteria that define specific requirements an organization must follow.
	ComplianceStandards []*KbArticle `json:"compliance_standards,omitempty"`
	// A Control is prescriptive, prioritized, and simplified set of best practices that one can use to strengthen their cybersecurity posture. e.g. AWS SecurityHub Controls, CIS Controls.
	Control string `json:"control"`
	// A list of requirements associated to a specific control in an industry or regulatory framework. e.g. <code> NIST.800-53.r5 AU-10 </code>
	Requirements []string `json:"requirements,omitempty"`
	// Security standards are a set of criteria organizations can follow to protect sensitive and confidential information. e.g. <code>NIST SP 800-53, CIS AWS Foundations Benchmark v1.4.0, ISO/IEC 27001</code>
	Standards []string `json:"standards"`
	// The resultant status of the compliance check  normalized to the caption of the <code>status_id</code> value. In the case of 'Other', it is defined by the event source.
	Status string `json:"status"`
	// The resultant status code of the compliance check.
	StatusCode string `json:"status_code,omitempty"`
	// The contextual description of the status, status_code values.
	StatusDetail string `json:"status_detail,omitempty"`
	// The normalized status identifier of the compliance check.
	StatusId int `json:"status_id"`
}

// DataClassification The Data Classification object includes information about data classification levels and data category types.
type DataClassification struct {
	// The name of the data classification category that data matched into, e.g. Financial, Personal, Governmental, etc.
	Category string `json:"category,omitempty"`
	// The normalized identifier of the data classification category.
	CategoryId int `json:"category_id"`
	// The file content confidentiality, normalized to the confidentiality_id value. In the case of 'Other', it is defined by the event source.
	Confidentiality string `json:"confidentiality,omitempty"`
	// The normalized identifier of the file content confidentiality indicator.
	ConfidentialityId int `json:"confidentiality_id"`
	// Details about the data policy that governs data handling and security measures related to classification.
	Policy *Policy `json:"policy,omitempty"`
}

// DataSecurity The Data Security object describes the characteristics, techniques and content of a Data Loss Prevention (DLP), Data Loss Detection (DLD), Data Classification, or similar tools' finding, alert, or detection mechanism(s).
type DataSecurity struct {
	// The name of the data classification category that data matched into, e.g. Financial, Personal, Governmental, etc.
	Category string `json:"category,omitempty"`
	// The normalized identifier of the data classification category.
	CategoryId int `json:"category_id"`
	// The file content confidentiality, normalized to the confidentiality_id value. In the case of 'Other', it is defined by the event source.
	Confidentiality string `json:"confidentiality,omitempty"`
	// The normalized identifier of the file content confidentiality indicator.
	ConfidentialityId int `json:"confidentiality_id"`
	// The name of the stage or state that the data was in. E.g., Data-at-Rest, Data-in-Transit, etc.
	DataLifecycleState string `json:"data_lifecycle_state,omitempty"`
	// The stage or state that the data was in when it was assessed or scanned by a data security tool.
	DataLifecycleStateId int `json:"data_lifecycle_state_id"`
	// Specific pattern, algorithm, fingerpint, or model used for detection.
	DetectionPattern string `json:"detection_pattern"`
	// The name of the type of data security tool or system that the finding, detection, or alert originated from. E.g., Endpoint, Secure Email Gateway, etc.
	DetectionSystem string `json:"detection_system,omitempty"`
	// The type of data security tool or system that the finding, detection, or alert originated from.
	DetectionSystemId int `json:"detection_system_id"`
	// A text, binary, file name, or datastore that matched against a detection rule.
	PatternMatch string `json:"pattern_match,omitempty"`
	// Details about the policy that triggered the finding.
	Policy *Policy `json:"policy"`
}

// Policy The Policy object describes the policies that are applicable. <p>Policy attributes provide traceability to the operational state of the security product at the time that the event was captured, facilitating forensics, troubleshooting, and policy tuning/adjustments.</p>
type Policy struct {
	// The description of the policy.
	Desc string `json:"desc,omitempty"`
	// The policy group.
	Group *Group `json:"group,omitempty"`
	// A determination if the content of a policy was applied to a target or request, or not.
	IsApplied bool `json:"is_applied"`
	// The policy name. For example: <code>IAM Policy</code>.
	Name string `json:"name"`
	// A unique identifier of the policy instance.
	Uid string `json:"uid"`
	// The policy version number.
	Version string `json:"version"`
}

// Scan The Scan object describes characteristics of a proactive scan.
type Scan struct {
	// The administrator-supplied or application-generated name of the scan. For example: "Home office weekly user database scan", "Scan folders for viruses", "Full system virus scan"
	Name string `json:"name"`
	// The type of scan.
	Type string `json:"type,omitempty"`
	// The type id of the scan.
	TypeId int `json:"type_id"`
	// The application-defined unique identifier assigned to an instance of a scan.
	Uid string `json:"uid"`
}

// SecurityState The Security State object describes the security related state of a managed entity.
type SecurityState struct {
	// The security state, normalized to the caption of the state_id value. In the case of 'Other', it is defined by the source.
	State string `json:"state,omitempty"`
	// The security state of the managed entity.
	StateId int `json:"state_id"`
}


package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DataSecurity The Data Security object describes the characteristics, techniques and content
// of a Data Loss Prevention (DLP), Data Loss Detection (DLD), Data Classification,
// or similar tools' finding, alert, or detection mechanism(s).
type DataSecurity struct {
	// The name of the data classification category that data matched into, e.g.
	// Financial, Personal, Governmental, etc.
	Category string `json:"category,omitempty"`

	// The normalized identifier of the data classification category.
	CategoryID int `json:"category_id,omitempty"`

	// Describes details about the classifier used for data classification.
	ClassifierDetails *ClassifierDetails `json:"classifier_details,omitempty"`

	// The file content confidentiality, normalized to the confidentiality_id
	// value. In the case of 'Other', it is defined by the event source.
	Confidentiality string `json:"confidentiality,omitempty"`

	// The normalized identifier of the file content confidentiality indicator.
	ConfidentialityID int `json:"confidentiality_id,omitempty"`

	// The name of the stage or state that the data was in. E.g., Data-at-Rest,
	// Data-in-Transit, etc.
	DataLifecycleState string `json:"data_lifecycle_state,omitempty"`

	// The stage or state that the data was in when it was assessed or scanned by a
	// data security tool.
	DataLifecycleStateID int `json:"data_lifecycle_state_id,omitempty"`

	// Specific pattern, algorithm, fingerprint, or model used for detection.
	DetectionPattern string `json:"detection_pattern,omitempty"`

	// The name of the type of data security tool or system that the finding,
	// detection, or alert originated from. E.g., Endpoint, Secure Email Gateway,
	// etc.
	DetectionSystem string `json:"detection_system,omitempty"`

	// The type of data security tool or system that the finding, detection, or
	// alert originated from.
	DetectionSystemID int `json:"detection_system_id,omitempty"`

	// Details about the data discovered by classification job.
	DiscoveryDetails []*DiscoveryDetails `json:"discovery_details,omitempty"`

	// A text, binary, file name, or datastore that matched against a detection
	// rule.
	PatternMatch string `json:"pattern_match,omitempty"`

	// Details about the policy that triggered the finding.
	Policy *Policy `json:"policy,omitempty"`

	// Size of the data classified.
	Size int64 `json:"size,omitempty"`

	// The source URL pointing towards the full classification job details.
	SrcURL string `json:"src_url,omitempty"`

	// The resultant status of the classification job normalized to the caption of
	// the status_id value. In the case of 'Other', it is defined by the event
	// source.
	Status string `json:"status,omitempty"`

	// The contextual description of the status, status_id value.
	StatusDetails []string `json:"status_details,omitempty"`

	// The normalized status identifier of the classification job.
	StatusID int `json:"status_id,omitempty"`

	// The total count of discovered entities, by the classification job.
	Total int64 `json:"total,omitempty"`

	// The unique identifier of the classification job.
	Uid string `json:"uid,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DataClassification The Data Classification object includes information about data classification
// levels and data category types.
type DataClassification struct {
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

	// Details about the data discovered by classification job.
	DiscoveryDetails []*DiscoveryDetails `json:"discovery_details,omitempty"`

	// Details about the data policy that governs data handling and security
	// measures related to classification.
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

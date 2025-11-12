package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Analytic The Analytic object contains details about the analytic technique used to
// analyze and derive insights from the data or information that led to the
// creation of a finding or conclusion.
type Analytic struct {
	// The algorithm used by the underlying analytic to generate the finding.
	Algorithm string `json:"algorithm,omitempty"`

	// The analytic category.
	Category string `json:"category,omitempty"`

	// The description of the analytic that generated the finding.
	Desc string `json:"desc,omitempty"`

	// The name of the analytic that generated the finding.
	Name string `json:"name,omitempty"`

	// Other analytics related to this analytic.
	RelatedAnalytics []*Analytic `json:"related_analytics,omitempty"`

	// The Analytic state.
	State string `json:"state,omitempty"`

	// The Analytic state identifier.
	StateID int `json:"state_id,omitempty"`

	// The analytic type.
	Type string `json:"type,omitempty"`

	// The analytic type ID.
	TypeID int `json:"type_id,omitempty"`

	// The unique identifier of the analytic that generated the finding.
	Uid string `json:"uid,omitempty"`

	// The analytic version. For example: 1.1.
	Version string `json:"version,omitempty"`
}

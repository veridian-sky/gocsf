package objects

// Code generated from OCSF schema; DO NOT EDIT.

// ClassifierDetails The Classifier Details object describes details about the classifier used for
// data classification.
type ClassifierDetails struct {
	// The name of the classifier.
	Name string `json:"name,omitempty"`

	// The type of the classifier.
	Type string `json:"type,omitempty"`

	// The unique identifier of the classifier.
	Uid string `json:"uid,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Metric The Metric object defines a simple name/value pair entity for a metric.
type Metric struct {
	// The name of the metric.
	Name string `json:"name,omitempty"`

	// The value of the metric.
	Value string `json:"value,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// AnalysisTarget The analysis target defines the scope of monitored activities, specifying what
// entity, system or process is analyzed for activity patterns.
type AnalysisTarget struct {
	// The specific name or identifier of the analysis target, such as the username
	// of a User Account, the name of a Kubernetes Cluster, the identifier of a
	// Network Namespace, or the name of an Application Component.
	Name string `json:"name,omitempty"`

	// The category of the analysis target, such as User Account, Kubernetes
	// Cluster, Network Namespace, or Application Component.
	Type string `json:"type,omitempty"`
}

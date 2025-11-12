package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Application An Application describes the details for an inventoried application as reported
// by an Application Security tool or other Developer-centric tooling. Applications
// can be defined as Kubernetes resources, Containerized resources, or application
// hosting-specific cloud sources such as AWS Elastic BeanStalk, AWS Lightsail, or
// Azure Logic Apps.
type Application struct {
	// The criticality of the application as defined by the event source.
	Criticality string `json:"criticality,omitempty"`

	// Additional data describing the application.
	Data interface{} `json:"data,omitempty"`

	// A description or commentary for an application, usually retrieved from an
	// upstream system.
	Desc string `json:"desc,omitempty"`

	// The name of the related application or associated resource group.
	Group *Group `json:"group,omitempty"`

	// The fully qualified name of the application.
	Hostname string `json:"hostname,omitempty"`

	// The list of labels associated to the application.
	Labels []string `json:"labels,omitempty"`

	// The name of the application.
	Name string `json:"name,omitempty"`

	// The identity of the service or user account that owns the application.
	Owner *User `json:"owner,omitempty"`

	// The cloud region of the resource.
	Region string `json:"region,omitempty"`

	// A graph representation showing how this application relates to and interacts
	// with other entities in the environment. This can include parent/child
	// relationships, dependencies, or other connections.
	ResourceRelationship *Graph `json:"resource_relationship,omitempty"`

	// The risk level, normalized to the caption of the risk_level_id value.
	RiskLevel string `json:"risk_level,omitempty"`

	// The normalized risk level id.
	RiskLevelID int `json:"risk_level_id,omitempty"`

	// The risk score as reported by the event source.
	RiskScore int64 `json:"risk_score,omitempty"`

	// The Software Bill of Materials (SBOM) associated with the application
	Sbom *Sbom `json:"sbom,omitempty"`

	// The list of tags; {key:value} pairs associated to the application.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The type of application as defined by the event source, e.g., GitHub, Azure
	// Logic App, or Amazon Elastic BeanStalk.
	Type string `json:"type,omitempty"`

	// The unique identifier for the application.
	Uid string `json:"uid,omitempty"`

	// An alternative or contextual identifier for the application, such as a
	// configuration, organization, or license UID.
	UidAlt string `json:"uid_alt,omitempty"`

	// The URL of the application.
	Url *Url `json:"url,omitempty"`

	// The semantic version of the application, e.g., 1.7.4.
	Version string `json:"version,omitempty"`
}

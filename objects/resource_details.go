package objects

// Code generated from OCSF schema; DO NOT EDIT.

// ResourceDetails The Resource Details object describes details about resources that were affected
// by the activity/event.
type ResourceDetails struct {
	// A list of agent objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`

	// The canonical cloud partition name to which the region is assigned (e.g. AWS
	// Partitions: aws, aws-cn, aws-us-gov).
	CloudPartition string `json:"cloud_partition,omitempty"`

	// The time when the resource was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the resource was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The criticality of the resource as defined by the event source.
	Criticality string `json:"criticality,omitempty"`

	// Additional data describing the resource.
	Data interface{} `json:"data,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// The name of the related resource group.
	Group *Group `json:"group,omitempty"`

	// The fully qualified name of the resource.
	Hostname string `json:"hostname,omitempty"`

	// The IP address of the resource, in either IPv4 or IPv6 format.
	Ip string `json:"ip,omitempty"`

	// Indicates whether the device or resource has a backup enabled, such as an
	// automated snapshot or a cloud backup. For example, this is indicated by the
	// cloudBackupEnabled value within JAMF Pro mobile devices or the registration
	// of an AWS ARN with the AWS Backup service.
	IsBackedUp bool `json:"is_backed_up,omitempty"`

	// The list of labels associated to the resource.
	Labels []string `json:"labels,omitempty"`

	// The time when the resource was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the resource was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The name of the resource.
	Name string `json:"name,omitempty"`

	// The namespace is useful when similar entities exist that you need to keep
	// separate.
	Namespace string `json:"namespace,omitempty"`

	// The identity of the service or user account that owns the resource.
	Owner *User `json:"owner,omitempty"`

	// The cloud region of the resource.
	Region string `json:"region,omitempty"`

	// A graph representation showing how this resource relates to and interacts
	// with other entities in the environment. This can include parent/child
	// relationships, dependencies, or other connections.
	ResourceRelationship *Graph `json:"resource_relationship,omitempty"`

	// The role of the resource in the context of the event or finding, normalized
	// to the caption of the role_id value. In the case of 'Other', it is defined
	// by the event source.
	Role string `json:"role,omitempty"`

	// The normalized identifier of the resource's role in the context of the event
	// or finding.
	RoleID int `json:"role_id,omitempty"`

	// The list of tags; {key:value} pairs associated to the resource.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The resource type as defined by the event source.
	Type string `json:"type,omitempty"`

	// The unique identifier of the resource.
	Uid interface{} `json:"uid,omitempty"`

	// The alternative unique identifier of the resource.
	UidAlt interface{} `json:"uid_alt,omitempty"`

	// The version of the resource. For example 1.2.3.
	Version string `json:"version,omitempty"`

	// The specific availability zone within a cloud region where the resource is
	// located.
	Zone string `json:"zone,omitempty"`
}

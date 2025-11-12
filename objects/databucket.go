package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Databucket The databucket object is a basic container that holds data, typically organized
// through the use of data partitions.
type Databucket struct {
	// A list of agent objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`

	// The canonical cloud partition name to which the region is assigned (e.g. AWS
	// Partitions: aws, aws-cn, aws-us-gov).
	CloudPartition string `json:"cloud_partition,omitempty"`

	// The time when the databucket was known to have been created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the databucket was known to have been created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The criticality of the databucket as defined by the event source.
	Criticality string `json:"criticality,omitempty"`

	// Additional data describing the resource.
	Data interface{} `json:"data,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// The description of the databucket.
	Desc string `json:"desc,omitempty"`

	// The encryption details of the databucket. Should be populated if the
	// databucket is encrypted.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`

	// Details about the file/object within a databucket.
	File *File `json:"file,omitempty"`

	// The name of the related resource group.
	Group *Group `json:"group,omitempty"`

	// The group names to which the databucket belongs.
	Groups []*Group `json:"groups,omitempty"`

	// The fully qualified hostname of the databucket.
	Hostname string `json:"hostname,omitempty"`

	// The IP address of the resource, in either IPv4 or IPv6 format.
	Ip string `json:"ip,omitempty"`

	// Indicates whether the device or resource has a backup enabled, such as an
	// automated snapshot or a cloud backup. For example, this is indicated by the
	// cloudBackupEnabled value within JAMF Pro mobile devices or the registration
	// of an AWS ARN with the AWS Backup service.
	IsBackedUp bool `json:"is_backed_up,omitempty"`

	// Indicates if the databucket is encrypted.
	IsEncrypted bool `json:"is_encrypted,omitempty"`

	// Indicates if the databucket is publicly accessible.
	IsPublic bool `json:"is_public,omitempty"`

	// The list of labels associated to the resource.
	Labels []string `json:"labels,omitempty"`

	// The most recent time when any changes, updates, or modifications were made
	// within the databucket.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The most recent time when any changes, updates, or modifications were made
	// within the databucket.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The databucket name.
	Name string `json:"name,omitempty"`

	// The namespace is useful when similar entities exist that you need to keep
	// separate.
	Namespace string `json:"namespace,omitempty"`

	// The identity of the service or user account that owns the databucket.
	Owner *User `json:"owner,omitempty"`

	// The cloud region of the databucket.
	Region string `json:"region,omitempty"`

	// A graph representation showing how this databucket relates to and interacts
	// with other entities in the environment. This can include parent/child
	// relationships, dependencies, or other connections.
	ResourceRelationship *Graph `json:"resource_relationship,omitempty"`

	// The size of the databucket in bytes.
	Size int64 `json:"size,omitempty"`

	// The list of tags; {key:value} pairs associated to the resource.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The databucket type.
	Type string `json:"type,omitempty"`

	// The normalized identifier of the databucket type.
	TypeID int `json:"type_id,omitempty"`

	// The unique identifier of the databucket.
	Uid interface{} `json:"uid,omitempty"`

	// The alternative unique identifier of the resource.
	UidAlt interface{} `json:"uid_alt,omitempty"`

	// The version of the resource. For example 1.2.3.
	Version string `json:"version,omitempty"`

	// The specific availability zone within a cloud region where the databucket is
	// located.
	Zone string `json:"zone,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Cloud The Cloud object contains information about a cloud or Software-as-a-Service
// account or similar construct, such as AWS Account ID, regions, organizations,
// folders, compartments, tenants, etc.
type Cloud struct {
	// The account object describes details about the account that was the source
	// or target of the activity.
	Account *Account `json:"account,omitempty"`

	// The canonical cloud partition name to which the region is assigned (e.g. AWS
	// Partitions: aws, aws-cn, aws-us-gov).
	CloudPartition string `json:"cloud_partition,omitempty"`

	// Organization and org unit relevant to the event or object.
	Org *Organization `json:"org,omitempty"`

	// The unique identifier of a Cloud project.
	ProjectUID string `json:"project_uid,omitempty"`

	// The unique name of the Cloud services provider, such as AWS, MS Azure, GCP,
	// etc.
	Provider string `json:"provider,omitempty"`

	// The name of the cloud region, as defined by the cloud provider.
	Region string `json:"region,omitempty"`

	// The availability zone in the cloud region, as defined by the cloud provider.
	Zone string `json:"zone,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// WinWinResource The Windows resource object describes a resource object managed by Windows, such
// as mutant or timer.
type WinWinResource struct {
	// The time when the resource was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the resource was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// Additional data describing the resource.
	Data interface{} `json:"data,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// The string detailing the attributes of the resource object.
	Details string `json:"details,omitempty"`

	// The list of labels associated to the resource.
	Labels []string `json:"labels,omitempty"`

	// The time when the resource was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the resource was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The name of the resource object.
	Name string `json:"name,omitempty"`

	// The Windows service acting as the object server for the resource object,
	// such as Security or Security Account Manager.
	SvcName string `json:"svc_name,omitempty"`

	// The list of tags; {key:value} pairs associated to the resource.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The type of the Windows resource object.
	Type string `json:"type,omitempty"`

	// The normalized type identifier of the Windows resource object accessed.
	TypeID int `json:"type_id,omitempty"`

	// The Windows provided handle identifier for the resource object
	Uid interface{} `json:"uid,omitempty"`

	// The alternative unique identifier of the resource.
	UidAlt interface{} `json:"uid_alt,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// WebResource The Web Resource object describes characteristics of a web resource that was
// affected by the activity/event.
type WebResource struct {
	// The time when the resource was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the resource was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// Details of the web resource, e.g, file details, search results or
	// application-defined resource.
	Data interface{} `json:"data,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// Description of the web resource.
	Desc string `json:"desc,omitempty"`

	// The list of labels associated to the resource.
	Labels []string `json:"labels,omitempty"`

	// The time when the resource was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the resource was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The name of the web resource.
	Name string `json:"name,omitempty"`

	// The list of tags; {key:value} pairs associated to the resource.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The web resource type as defined by the event source.
	Type string `json:"type,omitempty"`

	// The unique identifier of the web resource.
	Uid interface{} `json:"uid,omitempty"`

	// The alternative unique identifier of the resource.
	UidAlt interface{} `json:"uid_alt,omitempty"`

	// The URL pointing towards the source of the web resource.
	UrlString string `json:"url_string,omitempty"`
}

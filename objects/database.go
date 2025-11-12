package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Database The database object is used for databases which are typically datastore services
// that contain an organized collection of structured and unstructured data or a
// types of data.
type Database struct {
	// The time when the database was known to have been created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the database was known to have been created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// The description of the database.
	Desc string `json:"desc,omitempty"`

	// The group names to which the database belongs.
	Groups []*Group `json:"groups,omitempty"`

	// The most recent time when any changes, updates, or modifications were made
	// within the database.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The most recent time when any changes, updates, or modifications were made
	// within the database.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The database name, ordinarily as assigned by a database administrator.
	Name string `json:"name,omitempty"`

	// The size of the database in bytes.
	Size int64 `json:"size,omitempty"`

	// The database type.
	Type string `json:"type,omitempty"`

	// The normalized identifier of the database type.
	TypeID int `json:"type_id,omitempty"`

	// The unique identifier of the database.
	Uid string `json:"uid,omitempty"`
}

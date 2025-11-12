package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Table The table object represents a table within a structured relational database or
// datastore, which contains columns and rows of data that are able to be create,
// updated, deleted and queried.
type Table struct {
	// The time when the table was known to have been created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the table was known to have been created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The description of the table.
	Desc string `json:"desc,omitempty"`

	// The group names to which the table belongs.
	Groups []*Group `json:"groups,omitempty"`

	// The most recent time when any changes, updates, or modifications were made
	// within the table.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The most recent time when any changes, updates, or modifications were made
	// within the table.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The table name, ordinarily as assigned by a database administrator.
	Name string `json:"name,omitempty"`

	// The size of the data table in bytes.
	Size int64 `json:"size,omitempty"`

	// The unique identifier of the table.
	Uid string `json:"uid,omitempty"`
}

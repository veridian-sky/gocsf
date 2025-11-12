package objects

// Code generated from OCSF schema; DO NOT EDIT.

// QueryInfo The query info object holds information related to data access within a
// datastore. To access, manipulate, delete, or retrieve data from a datastore, a
// query must be written using a specific syntax.
type QueryInfo struct {
	// The size of the data returned from the query.
	Bytes int64 `json:"bytes,omitempty"`

	// The data returned from the query execution.
	Data interface{} `json:"data,omitempty"`

	// The query name for a saved or scheduled query.
	Name string `json:"name,omitempty"`

	// A string representing the query code being run. For example: SELECT * FROM
	// my_table
	QueryString string `json:"query_string,omitempty"`

	// The time when the query was run.
	QueryTime int64 `json:"query_time,omitempty"`

	// The time when the query was run.
	QueryTimeDt string `json:"query_time_dt,omitempty"`

	// The unique identifier of the query.
	Uid string `json:"uid,omitempty"`
}

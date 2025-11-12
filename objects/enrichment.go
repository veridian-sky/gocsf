package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Enrichment The Enrichment object provides inline enrichment data for specific attributes of
// interest within an event. It serves as a mechanism to enhance or supplement the
// information associated with the event by adding additional relevant details or
// context.
type Enrichment struct {
	// The time when the enrichment data was generated.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the enrichment data was generated.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The enrichment data associated with the attribute and value. The meaning of
	// this data depends on the type the enrichment record.
	Data interface{} `json:"data,omitempty"`

	// A long description of the enrichment data.
	Desc string `json:"desc,omitempty"`

	// The name of the attribute to which the enriched data pertains.
	Name string `json:"name,omitempty"`

	// The enrichment data provider name.
	Provider string `json:"provider,omitempty"`

	// The reputation of the enrichment data.
	Reputation *Reputation `json:"reputation,omitempty"`

	// A short description of the enrichment data.
	ShortDesc string `json:"short_desc,omitempty"`

	// The URL of the source of the enrichment data.
	SrcURL string `json:"src_url,omitempty"`

	// The enrichment type. For example: location.
	Type string `json:"type,omitempty"`

	// The value of the attribute to which the enriched data pertains.
	Value string `json:"value,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Ticket The Ticket object represents ticket in the customer's IT Service Management
// (ITSM) systems like ServiceNow, Jira, etc.
type Ticket struct {
	// The url of a ticket in the ticket system.
	SrcURL string `json:"src_url,omitempty"`

	// The status of the ticket normalized to the caption of the status_id value.
	// In the case of 99, this value should as defined by the source.
	Status string `json:"status,omitempty"`

	// A list of contextual descriptions of the status, status_id values.
	StatusDetails []string `json:"status_details,omitempty"`

	// The normalized identifier for the ticket status.
	StatusID int `json:"status_id,omitempty"`

	// The title of the ticket.
	Title string `json:"title,omitempty"`

	// The linked ticket type determines whether the ticket is internal or in an
	// external ticketing system.
	Type string `json:"type,omitempty"`

	// The normalized identifier for the ticket type.
	TypeID int `json:"type_id,omitempty"`

	// Unique identifier of the ticket.
	Uid string `json:"uid,omitempty"`
}

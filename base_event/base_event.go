package base_event

// Code generated from OCSF schema; DO NOT EDIT.

import "github.com/veridian-sky/gocsf/objects"

// BaseEvent contains all OCSF base event attributes
type BaseEvent struct {
	// The normalized identifier of the activity that triggered the event.
	ActivityID int `json:"activity_id,omitempty"`

	// The event activity name, as defined by the activity_id.
	ActivityName string `json:"activity_name,omitempty"`

	// The event category name, as defined by category_uid value.
	CategoryName string `json:"category_name,omitempty"`

	// The category unique identifier of the event.
	CategoryUID int `json:"category_uid,omitempty"`

	// The event class name, as defined by class_uid value: Base Event.
	ClassName string `json:"class_name,omitempty"`

	// The unique identifier of a class. A class describes the attributes available
	// in an event.
	ClassUID int `json:"class_uid,omitempty"`

	// The number of times that events in the same logical group occurred during
	// the event Start Time to End Time period.
	Count int64 `json:"count,omitempty"`

	// The event duration or aggregate time, the amount of time the event covers
	// from start_time to end_time in milliseconds.
	Duration int64 `json:"duration,omitempty"`

	// The end time of a time period, or the time of the most recent event included
	// in the aggregate event.
	EndTime int64 `json:"end_time,omitempty"`

	// The additional information from an external data source, which is associated
	// with the event or a finding. For example add location information for the IP
	// address in the DNS answers:</p>[{"name": "answers.ip", "value":
	// "92.24.47.250", "type": "location", "data": {"city": "Socotra", "continent":
	// "Asia", "coordinates": [-25.4153, 17.0743], "country": "YE", "desc":
	// "Yemen"}}]
	Enrichments []*objects.Enrichment `json:"enrichments,omitempty"`

	// The description of the event/finding, as defined by the source.
	Message string `json:"message,omitempty"`

	// The metadata associated with the event or a finding.
	Metadata *objects.Metadata `json:"metadata,omitempty"`

	// The observables associated with the event or a finding.
	Observables []*objects.Observable `json:"observables,omitempty"`

	// The raw event/finding data as received from the source.
	RawData string `json:"raw_data,omitempty"`

	// The hash, which describes the content of the raw_data field.
	RawDataHash *objects.Fingerprint `json:"raw_data_hash,omitempty"`

	// The size of the raw data which was transformed into an OCSF event, in bytes.
	RawDataSize int64 `json:"raw_data_size,omitempty"`

	// The event/finding severity, normalized to the caption of the severity_id
	// value. In the case of 'Other', it is defined by the source.
	Severity string `json:"severity,omitempty"`

	// <p>The normalized identifier of the event/finding severity.</p>The
	// normalized severity is a measurement the effort and expense required to
	// manage and resolve an event or incident. Smaller numerical values represent
	// lower impact events, and larger numerical values represent higher impact
	// events.
	SeverityID int `json:"severity_id,omitempty"`

	// The start time of a time period, or the time of the least recent event
	// included in the aggregate event.
	StartTime int64 `json:"start_time,omitempty"`

	// The event status, normalized to the caption of the status_id value. In the
	// case of 'Other', it is defined by the event source.
	Status string `json:"status,omitempty"`

	// The event status code, as reported by the event source.<br /><br />For
	// example, in a Windows Failed Authentication event, this would be the value
	// of 'Failure Code', e.g. 0x18.
	StatusCode string `json:"status_code,omitempty"`

	// The status detail contains additional information about the event/finding
	// outcome.
	StatusDetail string `json:"status_detail,omitempty"`

	// The normalized identifier of the event status.
	StatusID int `json:"status_id,omitempty"`

	// The normalized event occurrence time or the finding creation time.
	Time int64 `json:"time,omitempty"`

	// The number of minutes that the reported event time is ahead or behind UTC,
	// in the range -1,080 to +1,080.
	TimezoneOffset int64 `json:"timezone_offset,omitempty"`

	// The event/finding type name, as defined by the type_uid.
	TypeName string `json:"type_name,omitempty"`

	// The event/finding type ID. It identifies the event's semantics and
	// structure. The value is calculated by the logging system as: class_uid * 100
	// + activity_id.
	TypeUID int `json:"type_uid,omitempty"`

	// The attributes that are not mapped to the event schema. The names and values
	// of those attributes are specific to the event source.
	Unmapped *objects.Object `json:"unmapped,omitempty"`
}

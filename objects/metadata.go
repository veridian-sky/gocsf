package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Metadata The Metadata object describes the metadata associated with the event.
type Metadata struct {
	// The unique identifier used to correlate events.
	CorrelationUID string `json:"correlation_uid,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// Debug information about non-fatal issues with this OCSF event. Each issue is
	// a line in this string array.
	Debug []string `json:"debug,omitempty"`

	// The Event ID, Code, or Name that the product uses to primarily identify the
	// event.
	EventCode string `json:"event_code,omitempty"`

	// The schema extension used to create the event.
	Extension *Extension `json:"extension,omitempty"`

	// The schema extensions used to create the event.
	Extensions []*Extension `json:"extensions,omitempty"`

	// Indicates whether the OCSF event data has been truncated due to size
	// limitations. When true, some event data may have been omitted to fit within
	// system constraints.
	IsTruncated bool `json:"is_truncated,omitempty"`

	// The list of labels attached to the event. For example: ["sample", "dev"]
	Labels []string `json:"labels,omitempty"`

	// The audit level at which an event was generated.
	LogLevel string `json:"log_level,omitempty"`

	// The event log name. For example, syslog file name or Windows logging
	// subsystem: Security.
	LogName string `json:"log_name,omitempty"`

	// The logging provider or logging service that logged the event. For example,
	// Microsoft-Windows-Security-Auditing.
	LogProvider string `json:"log_provider,omitempty"`

	// The event log schema version that specifies the format of the original
	// event. For example syslog version or Cisco Log Schema Version.
	LogVersion string `json:"log_version,omitempty"`

	// <p>The time when the logging system collected and logged the event.</p>This
	// attribute is distinct from the event time in that event time typically
	// contain the time extracted from the original event. Most of the time, these
	// two times will be different.
	LoggedTime int64 `json:"logged_time,omitempty"`

	// <p>The time when the logging system collected and logged the event.</p>This
	// attribute is distinct from the event time in that event time typically
	// contain the time extracted from the original event. Most of the time, these
	// two times will be different.
	LoggedTimeDt string `json:"logged_time_dt,omitempty"`

	// An array of Logger objects that describe the devices and logging products
	// between the event source and its eventual destination. Note, this attribute
	// can be used when there is a complex end-to-end path of event flow.
	Loggers []*Logger `json:"loggers,omitempty"`

	// The time when the event was last modified or enriched.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the event was last modified or enriched.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The original event time as reported by the event source. For example, the
	// time in the original format from system event log such as Syslog on
	// Unix/Linux and the System event file on Windows. Omit if event is generated
	// instead of collected via logs.
	OriginalTime string `json:"original_time,omitempty"`

	// The event processed time, such as an ETL operation.
	ProcessedTime int64 `json:"processed_time,omitempty"`

	// The event processed time, such as an ETL operation.
	ProcessedTimeDt string `json:"processed_time_dt,omitempty"`

	// The product that reported the event.
	Product *Product `json:"product,omitempty"`

	// The list of profiles used to create the event. Profiles should be referenced
	// by their name attribute for core profiles, or extension/name for profiles
	// from extensions.
	Profiles []string `json:"profiles,omitempty"`

	// Sequence number of the event. The sequence number is a value available in
	// some events, to make the exact ordering of events unambiguous, regardless of
	// the event time precision.
	Sequence int64 `json:"sequence,omitempty"`

	// The list of tags; {key:value} pairs associated to the event.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The unique tenant identifier.
	TenantUID string `json:"tenant_uid,omitempty"`

	// An array of transformation info that describes the mappings or transforms
	// applied to the data.
	TransformationInfoList []*TransformationInfo `json:"transformation_info_list,omitempty"`

	// The logging system-assigned unique identifier of an event instance.
	Uid string `json:"uid,omitempty"`

	// The original size of the OCSF event data in kilobytes before any truncation
	// occurred. This field is typically populated when is_truncated is true to
	// indicate the full size of the original event.
	UntruncatedSize int64 `json:"untruncated_size,omitempty"`

	// The version of the OCSF schema, using Semantic Versioning Specification
	// (https://semver.org SemVer). For example: 1.0.0. Event consumers use the
	// version to determine the available event attributes.
	Version string `json:"version,omitempty"`
}

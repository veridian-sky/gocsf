package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Logger The Logger object represents the device and product where events are stored with
// times for receipt and transmission. This may be at the source device where the
// event occurred, a remote scanning device, intermediate hops, or the ultimate
// destination.
type Logger struct {
	// The device where the events are logged.
	Device *Device `json:"device,omitempty"`

	// The unique identifier of the event assigned by the logger.
	EventUID string `json:"event_uid,omitempty"`

	// Indicates whether the OCSF event data has been truncated due to size
	// limitations. When true, some event data may have been omitted to fit within
	// system constraints.
	IsTruncated bool `json:"is_truncated,omitempty"`

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

	// The name of the logging product instance.
	Name string `json:"name,omitempty"`

	// The product logging the event. This may be the event source product, a
	// management server product, a scanning product, a SIEM, etc.
	Product *Product `json:"product,omitempty"`

	// The time when the event was transmitted from the logging device to it's next
	// destination.
	TransmitTime int64 `json:"transmit_time,omitempty"`

	// The time when the event was transmitted from the logging device to it's next
	// destination.
	TransmitTimeDt string `json:"transmit_time_dt,omitempty"`

	// The unique identifier of the logging product instance.
	Uid string `json:"uid,omitempty"`

	// The original size of the OCSF event data in kilobytes before any truncation
	// occurred. This field is typically populated when is_truncated is true to
	// indicate the full size of the original event.
	UntruncatedSize int64 `json:"untruncated_size,omitempty"`

	// The version of the logging product.
	Version string `json:"version,omitempty"`
}

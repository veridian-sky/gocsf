package objects

// Code generated from OCSF schema; DO NOT EDIT.

// RelatedEvent The Related Event object describes an event or another finding related to a
// finding. It may or may not be an OCSF event.
type RelatedEvent struct {
	// An array of MITRE ATT&CK® objects describing identified tactics, techniques
	// & sub-techniques. The objects are compatible with MITRE ATLAS™ tactics,
	// techniques & sub-techniques.
	Attacks []*Attack `json:"attacks,omitempty"`

	// The number of times that activity in the same logical group occurred, as
	// reported by the related Finding.
	Count int64 `json:"count,omitempty"`

	// The time when the related event/finding was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the related event/finding was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// A description of the related event/finding.
	Desc string `json:"desc,omitempty"`

	// The time when the finding was first observed. e.g. The time when a
	// vulnerability was first observed. It can differ from the created_time
	// timestamp, which reflects the time this finding was created.
	FirstSeenTime int64 `json:"first_seen_time,omitempty"`

	// The time when the finding was first observed. e.g. The time when a
	// vulnerability was first observed. It can differ from the created_time
	// timestamp, which reflects the time this finding was created.
	FirstSeenTimeDt string `json:"first_seen_time_dt,omitempty"`

	// The
	// https://www.lockheedmartin.com/en-us/capabilities/cyber/cyber-kill-chain.html
	// Cyber Kill Chain® provides a detailed description of each phase and its
	// associated activities within the broader context of a cyber attack.
	KillChain []*KillChainPhase `json:"kill_chain,omitempty"`

	// The time when the finding was most recently observed. e.g. The time when a
	// vulnerability was most recently observed. It can differ from the
	// modified_time timestamp, which reflects the time this finding was last
	// modified.
	LastSeenTime int64 `json:"last_seen_time,omitempty"`

	// The time when the finding was most recently observed. e.g. The time when a
	// vulnerability was most recently observed. It can differ from the
	// modified_time timestamp, which reflects the time this finding was last
	// modified.
	LastSeenTimeDt string `json:"last_seen_time_dt,omitempty"`

	// The time when the related event/finding was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the related event/finding was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The observables associated with the event or a finding.
	Observables []*Observable `json:"observables,omitempty"`

	// Details about the product that reported the related event/finding.
	Product *Product `json:"product,omitempty"`

	// The unique identifier of the product that reported the related event.
	ProductUID string `json:"product_uid,omitempty"`

	// The event/finding severity, normalized to the caption of the severity_id
	// value. In the case of 'Other', it is defined by the source.
	Severity string `json:"severity,omitempty"`

	// <p>The normalized identifier of the event/finding severity.</p>The
	// normalized severity is a measurement the effort and expense required to
	// manage and resolve an event or incident. Smaller numerical values represent
	// lower impact events, and larger numerical values represent higher impact
	// events.
	SeverityID int `json:"severity_id,omitempty"`

	// The related event status. Should correspond to the label of the status_id
	// (or 'Other' status value for status_id = 99) of the related event.
	Status string `json:"status,omitempty"`

	// The list of tags; {key:value} pairs associated with the related
	// event/finding.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// A title or a brief phrase summarizing the related event/finding.
	Title string `json:"title,omitempty"`

	// The list of key traits or characteristics extracted from the related
	// event/finding that influenced or contributed to the overall finding's
	// outcome.
	Traits []*Trait `json:"traits,omitempty"`

	// The type of the related event/finding.</p>Populate if the related
	// event/finding is NOT in OCSF. If it is in OCSF, then utilize type_name,
	// type_uid instead.
	Type string `json:"type,omitempty"`

	// The type of the related OCSF event, as defined by type_uid.<p>For example:
	// Process Activity: Launch.</p>Populate if the related event/finding is in
	// OCSF.
	TypeName string `json:"type_name,omitempty"`

	// The unique identifier of the related OCSF event type. <p>For example:
	// 100701.</p>Populate if the related event/finding is in OCSF.
	TypeUID int64 `json:"type_uid,omitempty"`

	// The unique identifier of the related event/finding.</p> If the related
	// event/finding is in OCSF, then this value must be equal to metadata.uid in
	// the corresponding event.
	Uid string `json:"uid,omitempty"`
}

package classes

// Code generated from OCSF schema; DO NOT EDIT.

import "github.com/veridian-sky/gocsf/objects"

// FtpActivity File Transfer Protocol (FTP) Activity events report file transfers between a
// server and a client as seen on the network.
type FtpActivity struct {
	// The normalized caption of action_id.
	Action string `json:"action,omitempty"`

	// The action taken by a control or other policy-based system leading to an
	// outcome or disposition. An unknown action may still correspond to a known
	// disposition. Refer to disposition_id for the outcome of the action.
	ActionID int `json:"action_id,omitempty"`

	// The normalized identifier of the activity that triggered the event.
	ActivityID int `json:"activity_id,omitempty"`

	// The event activity name, as defined by the activity_id.
	ActivityName string `json:"activity_name,omitempty"`

	// The actor object describes details about the user/role/process that was the
	// source of the activity. Note that this is not the threat actor of a campaign
	// but may be part of a campaign.
	Actor *objects.Actor `json:"actor,omitempty"`

	// Describes details about a typical API (Application Programming Interface)
	// call.
	Api *objects.Api `json:"api,omitempty"`

	// The name of the application associated with the event or object.
	AppName string `json:"app_name,omitempty"`

	// An array of MITRE ATT&CK® objects describing identified tactics, techniques
	// & sub-techniques. The objects are compatible with MITRE ATLAS™ tactics,
	// techniques & sub-techniques.
	Attacks []*objects.Attack `json:"attacks,omitempty"`

	// Provides details about an authorization, such as authorization outcome, and
	// any associated policies related to the activity/event.
	Authorizations []*objects.Authorization `json:"authorizations,omitempty"`

	// The event category name, as defined by category_uid value: Network Activity.
	CategoryName string `json:"category_name,omitempty"`

	// The category unique identifier of the event.
	CategoryUID int `json:"category_uid,omitempty"`

	// The event class name, as defined by class_uid value: FTP Activity.
	ClassName string `json:"class_name,omitempty"`

	// The unique identifier of a class. A class describes the attributes available
	// in an event.
	ClassUID int `json:"class_uid,omitempty"`

	// Describes details about the Cloud environment where the event was originally
	// created or logged.
	Cloud *objects.Cloud `json:"cloud,omitempty"`

	// The list of return codes to the FTP command.
	Codes []int64 `json:"codes,omitempty"`

	// The FTP command.
	Command string `json:"command,omitempty"`

	// The list of responses to the FTP command.
	CommandResponses []string `json:"command_responses,omitempty"`

	// The confidence, normalized to the caption of the confidence_id value. In the
	// case of 'Other', it is defined by the event source.
	Confidence string `json:"confidence,omitempty"`

	// The normalized confidence refers to the accuracy of the rule that created
	// the finding. A rule with a low confidence means that the finding scope is
	// wide and may create finding reports that may not be malicious in nature.
	ConfidenceID int `json:"confidence_id,omitempty"`

	// The confidence score as reported by the event source.
	ConfidenceScore int64 `json:"confidence_score,omitempty"`

	// The network connection information.
	ConnectionInfo *objects.NetworkConnectionInfo `json:"connection_info,omitempty"`

	// The number of times that events in the same logical group occurred during
	// the event Start Time to End Time period.
	Count int64 `json:"count,omitempty"`

	// An addressable device, computer system or host.
	Device *objects.Device `json:"device,omitempty"`

	// The disposition name, normalized to the caption of the disposition_id value.
	// In the case of 'Other', it is defined by the event source.
	Disposition string `json:"disposition,omitempty"`

	// Describes the outcome or action taken by a security control, such as access
	// control checks, malware detections or various types of policy violations.
	DispositionID int `json:"disposition_id,omitempty"`

	// The responder (server) in a network connection.
	DstEndpoint *objects.NetworkEndpoint `json:"dst_endpoint,omitempty"`

	// The event duration or aggregate time, the amount of time the event covers
	// from start_time to end_time in milliseconds.
	Duration int64 `json:"duration,omitempty"`

	// The end time of a time period, or the time of the most recent event included
	// in the aggregate event.
	EndTime int64 `json:"end_time,omitempty"`

	// The end time of a time period, or the time of the most recent event included
	// in the aggregate event.
	EndTimeDt string `json:"end_time_dt,omitempty"`

	// The additional information from an external data source, which is associated
	// with the event or a finding. For example add location information for the IP
	// address in the DNS answers:</p>[{"name": "answers.ip", "value":
	// "92.24.47.250", "type": "location", "data": {"city": "Socotra", "continent":
	// "Asia", "coordinates": [-25.4153, 17.0743], "country": "YE", "desc":
	// "Yemen"}}]
	Enrichments []*objects.Enrichment `json:"enrichments,omitempty"`

	// The file that is the target of the FTP activity.
	File *objects.File `json:"file,omitempty"`

	// The firewall rule that pertains to the control that triggered the event, if
	// applicable.
	FirewallRule *objects.FirewallRule `json:"firewall_rule,omitempty"`

	// Indicates that the event is considered to be an alertable signal. Should be
	// set to true if disposition_id = Alert among other dispositions, and/or
	// risk_level_id or severity_id of the event is elevated. Not all control
	// events will be alertable, for example if disposition_id = Exonerated or
	// disposition_id = Allowed.
	IsAlert bool `json:"is_alert,omitempty"`

	// A list of the JA4+ network fingerprints.
	Ja4FingerprintList []*objects.Ja4Fingerprint `json:"ja4_fingerprint_list,omitempty"`

	// The Load Balancer object contains information related to the device that is
	// distributing incoming traffic to specified destinations.
	LoadBalancer *objects.LoadBalancer `json:"load_balancer,omitempty"`

	// A list of Malware objects, describing details about the identified malware.
	Malware []*objects.Malware `json:"malware,omitempty"`

	// Describes details about the scan job that identified malware on the target
	// system.
	MalwareScanInfo *objects.MalwareScanInfo `json:"malware_scan_info,omitempty"`

	// The description of the event/finding, as defined by the source.
	Message string `json:"message,omitempty"`

	// The metadata associated with the event or a finding.
	Metadata *objects.Metadata `json:"metadata,omitempty"`

	// The name of the data affiliated with the command.
	Name string `json:"name,omitempty"`

	// The observables associated with the event or a finding.
	Observables []*objects.Observable `json:"observables,omitempty"`

	// The OSINT (Open Source Intelligence) object contains details related to an
	// indicator such as the indicator itself, related indicators, geolocation,
	// registrar information, subdomains, analyst commentary, and other contextual
	// information. This information can be used to further enrich a detection or
	// finding by providing decisioning support to other analysts and engineers.
	Osint []*objects.Osint `json:"osint,omitempty"`

	// The policy that pertains to the control that triggered the event, if
	// applicable. For example the name of an anti-malware policy or an access
	// control policy.
	Policy *objects.Policy `json:"policy,omitempty"`

	// The dynamic port established for impending data transfers.
	Port interface{} `json:"port,omitempty"`

	// The proxy (server) in a network connection.
	Proxy *objects.NetworkProxy `json:"proxy,omitempty"`

	// The connection information from the proxy server to the remote server.
	ProxyConnectionInfo *objects.NetworkConnectionInfo `json:"proxy_connection_info,omitempty"`

	// The proxy (server) in a network connection.
	ProxyEndpoint *objects.NetworkProxy `json:"proxy_endpoint,omitempty"`

	// The HTTP Request from the proxy server to the remote server.
	ProxyHTTPRequest *objects.HttpRequest `json:"proxy_http_request,omitempty"`

	// The HTTP Response from the remote server to the proxy server.
	ProxyHTTPResponse *objects.HttpResponse `json:"proxy_http_response,omitempty"`

	// The TLS protocol negotiated between the proxy server and the remote server.
	ProxyTLS *objects.Tls `json:"proxy_tls,omitempty"`

	// The network traffic refers to the amount of data moving across a network,
	// from proxy to remote server at a given point of time.
	ProxyTraffic *objects.NetworkTraffic `json:"proxy_traffic,omitempty"`

	// The raw event/finding data as received from the source.
	RawData string `json:"raw_data,omitempty"`

	// The hash, which describes the content of the raw_data field.
	RawDataHash *objects.Fingerprint `json:"raw_data_hash,omitempty"`

	// The size of the raw data which was transformed into an OCSF event, in bytes.
	RawDataSize int64 `json:"raw_data_size,omitempty"`

	// Describes the risk associated with the finding.
	RiskDetails string `json:"risk_details,omitempty"`

	// The risk level, normalized to the caption of the risk_level_id value.
	RiskLevel string `json:"risk_level,omitempty"`

	// The normalized risk level id.
	RiskLevelID int `json:"risk_level_id,omitempty"`

	// The risk score as reported by the event source.
	RiskScore int64 `json:"risk_score,omitempty"`

	// The event/finding severity, normalized to the caption of the severity_id
	// value. In the case of 'Other', it is defined by the source.
	Severity string `json:"severity,omitempty"`

	// <p>The normalized identifier of the event/finding severity.</p>The
	// normalized severity is a measurement the effort and expense required to
	// manage and resolve an event or incident. Smaller numerical values represent
	// lower impact events, and larger numerical values represent higher impact
	// events.
	SeverityID int `json:"severity_id,omitempty"`

	// The initiator (client) of the network connection.
	SrcEndpoint *objects.NetworkEndpoint `json:"src_endpoint,omitempty"`

	// The start time of a time period, or the time of the least recent event
	// included in the aggregate event.
	StartTime int64 `json:"start_time,omitempty"`

	// The start time of a time period, or the time of the least recent event
	// included in the aggregate event.
	StartTimeDt string `json:"start_time_dt,omitempty"`

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

	// The normalized event occurrence time or the finding creation time.
	TimeDt string `json:"time_dt,omitempty"`

	// The number of minutes that the reported event time is ahead or behind UTC,
	// in the range -1,080 to +1,080.
	TimezoneOffset int64 `json:"timezone_offset,omitempty"`

	// The Transport Layer Security (TLS) attributes.
	Tls *objects.Tls `json:"tls,omitempty"`

	// The network traffic refers to the amount of data moving across a network at
	// a given point of time. Intended to be used alongside Network Connection.
	Traffic *objects.NetworkTraffic `json:"traffic,omitempty"`

	// The type of FTP network connection (e.g. active, passive).
	Type string `json:"type,omitempty"`

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

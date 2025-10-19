// Code generated from OCSF schema. DO NOT EDIT.
package gocsf

import (
	"encoding/json"
	"time"
)

// Common Types

// Analytic The Analytic object contains details about the analytic technique used to analyze and derive insights from the data or information that led to the creation of a finding or conclusion.
type Analytic struct {
	// The analytic category.
	Category string `json:"category,omitempty"`
	// The description of the analytic that generated the finding.
	Desc string `json:"desc,omitempty"`
	// The name of the analytic that generated the finding.
	Name string `json:"name"`
	// Other analytics related to this analytic.
	RelatedAnalytics []*Analytic `json:"related_analytics,omitempty"`
	// The analytic type.
	Type string `json:"type,omitempty"`
	// The analytic type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the analytic that generated the finding.
	Uid string `json:"uid"`
	// The analytic version. For example: <code>1.1</code>.
	Version string `json:"version,omitempty"`
}

// Certificate The Digital Certificate, also known as a Public Key Certificate, object contains information about the ownership and usage of a public key. It serves as a means to establish trust in the authenticity and integrity of the public key and the associated entity. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Certificate/'>d3f:Certificate</a>.
type Certificate struct {
	// The time when the certificate was created.
	CreatedTime time.Time `json:"created_time"`
	// The time when the certificate was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The expiration time of the certificate.
	ExpirationTime time.Time `json:"expiration_time"`
	// The expiration time of the certificate.
	ExpirationTimeDt time.Time `json:"expiration_time_dt,omitempty"`
	// The fingerprint list of the certificate.
	Fingerprints []*Fingerprint `json:"fingerprints"`
	// Denotes whether a digital certificate is self-signed or signed by a known certificate authority (CA).
	IsSelfSigned bool `json:"is_self_signed"`
	// The certificate issuer distinguished name.
	Issuer string `json:"issuer"`
	// The serial number of the certificate used to create the digital signature.
	SerialNumber string `json:"serial_number"`
	// The certificate subject distinguished name.
	Subject string `json:"subject"`
	// The unique identifier of the certificate.
	Uid string `json:"uid,omitempty"`
	// The certificate version.
	Version string `json:"version"`
}

// DigitalSignature The Digital Signature object contains information about the cryptographic mechanism used to verify the authenticity, integrity, and origin of the file or application.
type DigitalSignature struct {
	// The digital signature algorithm used to create the signature, normalized to the caption of 'algorithm_id'. In the case of 'Other', it is defined by the event source.
	Algorithm string `json:"algorithm,omitempty"`
	// The identifier of the normalized digital signature algorithm.
	AlgorithmId int `json:"algorithm_id"`
	// The certificate object containing information about the digital certificate.
	Certificate *Certificate `json:"certificate"`
	// The time when the digital signature was created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the digital signature was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The developer ID on the certificate that signed the file.
	DeveloperUid string `json:"developer_uid,omitempty"`
	// The message digest attribute contains the fixed length message hash representation and the corresponding hashing algorithm information.
	Digest *Fingerprint `json:"digest,omitempty"`
	// The digital signature state defines the signature state, normalized to the caption of 'state_id'. In the case of 'Other', it is defined by the event source.
	State string `json:"state,omitempty"`
	// The normalized identifier of the signature state.
	StateId int `json:"state_id,omitempty"`
}

// Email The Email object describes the email metadata such as sender, recipients, and direction. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Email/'>d3f:Email</a>.
type Email struct {
	// The email header Cc values, as defined by RFC 5322.
	Cc []string `json:"cc,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The <strong>Delivered-To</strong> email header field.
	DeliveredTo string `json:"delivered_to,omitempty"`
	// The email header From values, as defined by RFC 5322.
	From string `json:"from"`
	// The email header Message-Id value, as defined by RFC 5322.
	MessageUid string `json:"message_uid"`
	// The email authentication header.
	RawHeader string `json:"raw_header,omitempty"`
	// The email header Reply-To values, as defined by RFC 5322.
	ReplyTo string `json:"reply_to"`
	// The size in bytes of the email, including attachments.
	Size int64 `json:"size"`
	// The value of the SMTP MAIL FROM command.
	SmtpFrom string `json:"smtp_from"`
	// The value of the SMTP envelope RCPT TO command.
	SmtpTo []string `json:"smtp_to"`
	// The email header Subject value, as defined by RFC 5322.
	Subject string `json:"subject"`
	// The email header To values, as defined by RFC 5322.
	To []string `json:"to"`
	// The email unique identifier.
	Uid string `json:"uid"`
	// The X-Originating-IP header identifying the emails originating IP address(es).
	XOriginatingIp []string `json:"x_originating_ip,omitempty"`
}

// EmailAuth The Email Authentication object describes the Sender Policy Framework (SPF), DomainKeys Identified Mail (DKIM) and Domain-based Message Authentication, Reporting and Conformance (DMARC) attributes of an email.
type EmailAuth struct {
	// The DomainKeys Identified Mail (DKIM) status of the email.
	Dkim string `json:"dkim"`
	// The DomainKeys Identified Mail (DKIM) signing domain of the email.
	DkimDomain string `json:"dkim_domain"`
	// The DomainKeys Identified Mail (DKIM) signature used by the sending/receiving system.
	DkimSignature string `json:"dkim_signature"`
	// The Domain-based Message Authentication, Reporting and Conformance (DMARC) status of the email.
	Dmarc string `json:"dmarc"`
	// The Domain-based Message Authentication, Reporting and Conformance (DMARC) override action.
	DmarcOverride string `json:"dmarc_override"`
	// The Domain-based Message Authentication, Reporting and Conformance (DMARC) policy status.
	DmarcPolicy string `json:"dmarc_policy"`
	// The Sender Policy Framework (SPF) status of the email.
	Spf string `json:"spf"`
}

// Extension The OCSF Schema Extension object provides detailed information about the schema extension used to construct the event. The schema extensions are registered in the <a target='_blank' href='https://github.com/ocsf/ocsf-schema/blob/main/extensions.md'>extensions.md</a> file.
type Extension struct {
	// The schema extension name. For example: <code>dev</code>.
	Name string `json:"name"`
	// The schema extension unique identifier. For example: <code>999</code>.
	Uid string `json:"uid"`
	// The schema extension version. For example: <code>1.0.0-alpha.2</code>.
	Version string `json:"version"`
}

// Feature The Feature object provides information about the software product feature that generated a specific event. It encompasses details related to the capabilities, components, user interface (UI) design, and performance upgrades associated with the feature.
type Feature struct {
	// The name of the feature.
	Name string `json:"name"`
	// The unique identifier of the feature.
	Uid string `json:"uid"`
	// The version of the feature.
	Version string `json:"version"`
}

// Fingerprint The Fingerprint object provides detailed information about a digital fingerprint, which is a compact representation of data used to identify a longer piece of information, such as a public key or file content. It contains the algorithm and value of the fingerprint, enabling efficient and reliable identification of the associated data.
type Fingerprint struct {
	// The hash algorithm used to create the digital fingerprint, normalized to the caption of 'algorithm_id'. In the case of 'Other', it is defined by the event source.
	Algorithm string `json:"algorithm,omitempty"`
	// The identifier of the normalized hash algorithm, which was used to create the digital fingerprint.
	AlgorithmId int `json:"algorithm_id"`
	// The digital fingerprint value.
	Value string `json:"value"`
}

// FirewallRule The Firewall Rule object represents a specific rule within a firewall policy or event. It contains information about a rule's configuration, properties, and associated actions that define how network traffic is handled by the firewall.
type FirewallRule struct {
	// The rule category.
	Category string `json:"category,omitempty"`
	// The rule trigger condition for the rule. For example: SQL_INJECTION.
	Condition string `json:"condition,omitempty"`
	// The description of the rule that generated the event.
	Desc string `json:"desc,omitempty"`
	// The rule response time duration, usually used for challenge completion time.
	Duration int64 `json:"duration,omitempty"`
	// The data in a request that rule matched. For example: '["10","and","1"]'.
	MatchDetails []string `json:"match_details,omitempty"`
	// The location of the matched data in the source which resulted in the triggered firewall rule. For example: HEADER.
	MatchLocation string `json:"match_location,omitempty"`
	// The name of the rule that generated the event.
	Name string `json:"name"`
	// The rate limit for a rate-based rule.
	RateLimit int `json:"rate_limit,omitempty"`
	// The sensitivity of the firewall rule in the matched event. For example: HIGH.
	Sensitivity string `json:"sensitivity,omitempty"`
	// The rule type.
	Type string `json:"type,omitempty"`
	// The unique identifier of the rule that generated the event.
	Uid string `json:"uid"`
	// The rule version. For example: <code>1.1</code>.
	Version string `json:"version,omitempty"`
}

// Hassh The HASSH object contains SSH network fingerprinting values for specific client/server implementations. It provides a standardized way of identifying and categorizing SSH connections based on their unique characteristics and behavior.
type Hassh struct {
	// The concatenation of key exchange, encryption, authentication and compression algorithms (separated by ';'). NOTE: This is not the underlying algorithm for the hash implementation.
	Algorithm string `json:"algorithm"`
	// The hash of the key exchange, encryption, authentication and compression algorithms.
	Fingerprint *Fingerprint `json:"fingerprint"`
}

// Image The Image object provides a description of a specific Virtual Machine (VM) or Container image. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:ContainerImage/'>d3f:ContainerImage</a>.
type Image struct {
	// The image labels.
	Labels []string `json:"labels,omitempty"`
	// The image name. For example: <code>elixir</code>.
	Name string `json:"name"`
	// The full path to the image file.
	Path string `json:"path,omitempty"`
	// The image tag. For example: <code>1.11-alpine</code>.
	Tag string `json:"tag,omitempty"`
	// The unique image ID. For example: <code>77af4d6b9913</code>.
	Uid string `json:"uid"`
}

// Ja4Fingerprint The JA4+ fingerprint object provides detailed fingerprint information about various aspects of network traffic which is both machine and human readable.
type Ja4Fingerprint struct {
	// The 'a' section of the JA4 fingerprint.
	SectionA string `json:"section_a,omitempty"`
	// The 'b' section of the JA4 fingerprint.
	SectionB string `json:"section_b,omitempty"`
	// The 'c' section of the JA4 fingerprint.
	SectionC string `json:"section_c,omitempty"`
	// The 'd' section of the JA4 fingerprint.
	SectionD string `json:"section_d,omitempty"`
	// The JA4+ fingerprint type as defined by <a href='https://blog.foxio.io/ja4+-network-fingerprinting target='_blank'>FoxIO</a>, normalized to the caption of 'type_id'. In the case of 'Other', it is defined by the event source.
	Type string `json:"type,omitempty"`
	// The identifier of the JA4+ fingerprint type.
	TypeId int `json:"type_id"`
	// The JA4+ fingerprint value.
	Value string `json:"value"`
}

// Job The Job object provides information about a scheduled job or task, including its name, command line, and state. It encompasses attributes that describe the properties and status of the scheduled job.
type Job struct {
	// The job command line.
	CmdLine string `json:"cmd_line"`
	// The time when the job was created.
	CreatedTime time.Time `json:"created_time"`
	// The time when the job was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The description of the job.
	Desc string `json:"desc"`
	// The file that pertains to the job.
	File *File `json:"file"`
	// The time when the job was last run.
	LastRunTime time.Time `json:"last_run_time"`
	// The time when the job was last run.
	LastRunTimeDt time.Time `json:"last_run_time_dt,omitempty"`
	// The name of the job.
	Name string `json:"name"`
	// The time when the job will next be run.
	NextRunTime time.Time `json:"next_run_time,omitempty"`
	// The time when the job will next be run.
	NextRunTimeDt time.Time `json:"next_run_time_dt,omitempty"`
	// The run state of the job.
	RunState string `json:"run_state,omitempty"`
	// The run state ID of the job.
	RunStateId int `json:"run_state_id"`
	// The user that created the job.
	User *User `json:"user,omitempty"`
}

// Location The Geo Location object describes a geographical location, usually associated with an IP address. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:PhysicalLocation/'>d3f:PhysicalLocation</a>.
type Location struct {
	// The name of the city.
	City string `json:"city"`
	// The name of the continent.
	Continent string `json:"continent"`
	// A two-element array, containing a longitude/latitude pair. The format conforms with <a target='_blank' href='https://geojson.org'>GeoJSON</a>. For example: <code>[-73.983, 40.719]</code>.
	Coordinates []float64 `json:"coordinates,omitempty"`
	// The ISO 3166-1 Alpha-2 country code. For the complete list of country codes see <a target='_blank' href='https://www.iso.org/obp/ui/#iso:pub:PUB500001:en' >ISO 3166-1 alpha-2 codes</a>.<p><b>Note:</b> The two letter country code should be capitalized. For example: <code>US</code> or <code>CA</code>.</p>
	Country string `json:"country"`
	// The description of the geographical location.
	Desc string `json:"desc,omitempty"`
	// <p>Geohash of the geo-coordinates (latitude and longitude).</p><a target='_blank' href='https://en.wikipedia.org/wiki/Geohash'>Geohashing</a> is a geocoding system used to encode geographic coordinates in decimal degrees, to a single string.
	Geohash string `json:"geohash,omitempty"`
	// The indication of whether the location is on premises.
	IsOnPremises bool `json:"is_on_premises,omitempty"`
	// The name of the Internet Service Provider (ISP).
	Isp string `json:"isp,omitempty"`
	// The geographical Latitude coordinate represented in Decimal Degrees (DD). For example: <code>42.361145</code>.
	Lat float64 `json:"lat,omitempty"`
	// The geographical Longitude coordinate represented in Decimal Degrees (DD). For example: <code>-71.057083</code>.
	Long float64 `json:"long,omitempty"`
	// The postal code of the location.
	PostalCode string `json:"postal_code,omitempty"`
	// The provider of the geographical location data.
	Provider string `json:"provider,omitempty"`
	// The alphanumeric code that identifies the principal subdivision (e.g. province or state) of the country. Region codes are defined at <a target='_blank' href='https://www.iso.org/iso-3166-country-codes.html'>ISO 3166-2</a> and have a limit of three characters. For example, see <a target='_blank' href='https://www.iso.org/obp/ui/#iso:code:3166:US'>the region codes for the US</a>.
	Region string `json:"region,omitempty"`
}

// Logger The Logger object represents the device and product where events are stored with times for receipt and transmission.  This may be at the source device where the event occurred, a remote scanning device, intermediate hops, or the ultimate destination.
type Logger struct {
	// The device where the events are logged.
	Device *Device `json:"device"`
	// The audit level at which an event was generated.
	LogLevel string `json:"log_level,omitempty"`
	// The event log name. For example, syslog file name or Windows logging subsystem: Security.
	LogName string `json:"log_name"`
	// The logging provider or logging service that logged the event. For example, Microsoft-Windows-Security-Auditing.
	LogProvider string `json:"log_provider"`
	// The event log schema version that specifies the format of the original event. For example syslog version or Cisco Log Schema Version.
	LogVersion string `json:"log_version,omitempty"`
	// <p>The time when the logging system collected and logged the event.</p>This attribute is distinct from the event time in that event time typically contain the time extracted from the original event. Most of the time, these two times will be different.
	LoggedTime time.Time `json:"logged_time"`
	// <p>The time when the logging system collected and logged the event.</p>This attribute is distinct from the event time in that event time typically contain the time extracted from the original event. Most of the time, these two times will be different.
	LoggedTimeDt time.Time `json:"logged_time_dt,omitempty"`
	// The name of the logging product instance.
	Name string `json:"name"`
	// The product logging the event.  This may be the event source product, a management server product, a scanning product, a SIEM, etc.
	Product *Product `json:"product"`
	// The time when the event was transmitted from the logging device to it's next destination.
	TransmitTime time.Time `json:"transmit_time,omitempty"`
	// The time when the event was transmitted from the logging device to it's next destination.
	TransmitTimeDt time.Time `json:"transmit_time_dt,omitempty"`
	// The unique identifier of the logging product instance.
	Uid string `json:"uid"`
	// The version of the logging product.
	Version string `json:"version,omitempty"`
}

// ManagedEntity The Managed Entity object describes the type and version of an entity, such as a user, device, or policy.  For types in the <code>type_id</code> enum list, an associated attribute should be populated.  If the type of entity is not in the <code>type_id</code> list, information can be put into the <code>data</code> attribute and the <code>type</code> attribute should identify the entity.
type ManagedEntity struct {
	// The managed entity content as a JSON object.
	Data json.RawMessage `json:"data,omitempty"`
	// An addressable device, computer system or host.
	Device *Device `json:"device"`
	// The email object.
	Email *Email `json:"email"`
	// The group object associated with an entity such as user, policy, or rule.
	Group *Group `json:"group"`
	// The name of the managed entity.
	Name string `json:"name"`
	// Organization and org unit relevant to the event or object.
	Org *Organization `json:"org"`
	// Describes details of a managed policy.
	Policy *Policy `json:"policy"`
	// The managed entity type. For example: <code>policy</code>, <code>user</code>, <code>organizational unit</code>, <code>device</code>.
	Type string `json:"type"`
	// The type of the Managed Entity. It is recommended to also populate the <code>type</code> attribute with the associated label, or the source specific name if <code>Other</code>.
	TypeId int `json:"type_id"`
	// The identifier of the managed entity.
	Uid string `json:"uid"`
	// The user that pertains to the event or object.
	User *User `json:"user"`
	// The version of the managed entity. For example: <code>1.2.3</code>.
	Version string `json:"version"`
}

// Metadata The Metadata object describes the metadata associated with the event. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Metadata/'>d3f:Metadata</a>.
type Metadata struct {
	// The unique identifier used to correlate events.
	CorrelationUid string `json:"correlation_uid,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The Event ID or Code that the product uses to describe the event.
	EventCode string `json:"event_code,omitempty"`
	// The schema extension used to create the event.
	Extension *Extension `json:"extension,omitempty"`
	// The schema extensions used to create the event.
	Extensions []*Extension `json:"extensions,omitempty"`
	// <p>The list of category labels attached to the event or specific attributes. Labels are user defined tags or aliases added at normalization time.</p>For example: <code>["network", "connection.ip:destination", "device.ip:source"]</code>
	Labels []string `json:"labels,omitempty"`
	// The audit level at which an event was generated.
	LogLevel string `json:"log_level,omitempty"`
	// The event log name. For example, syslog file name or Windows logging subsystem: Security.
	LogName string `json:"log_name"`
	// The logging provider or logging service that logged the event. For example, Microsoft-Windows-Security-Auditing.
	LogProvider string `json:"log_provider"`
	// The event log schema version that specifies the format of the original event. For example syslog version or Cisco Log Schema Version.
	LogVersion string `json:"log_version,omitempty"`
	// <p>The time when the logging system collected and logged the event.</p>This attribute is distinct from the event time in that event time typically contain the time extracted from the original event. Most of the time, these two times will be different.
	LoggedTime time.Time `json:"logged_time,omitempty"`
	// <p>The time when the logging system collected and logged the event.</p>This attribute is distinct from the event time in that event time typically contain the time extracted from the original event. Most of the time, these two times will be different.
	LoggedTimeDt time.Time `json:"logged_time_dt,omitempty"`
	// An array of Logger objects that describe the devices and logging products between the event source and its eventual destination. Note, this attribute can be used when there is a complex end-to-end path of event flow.
	Loggers []*Logger `json:"loggers,omitempty"`
	// The time when the event was last modified or enriched.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The time when the event was last modified or enriched.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The original event time as reported by the event source. For example, the time in the original format from system event log such as Syslog on Unix/Linux and the System event file on Windows. Omit if event is generated instead of collected via logs.
	OriginalTime string `json:"original_time"`
	// The event processed time, such as an ETL operation.
	ProcessedTime time.Time `json:"processed_time,omitempty"`
	// The event processed time, such as an ETL operation.
	ProcessedTimeDt time.Time `json:"processed_time_dt,omitempty"`
	// The product that reported the event.
	Product *Product `json:"product"`
	// The list of profiles used to create the event.  Profiles should be referenced by their <code>name</code> attribute for core profiles, or <code>extension/name</code> for profiles from extensions.
	Profiles []string `json:"profiles,omitempty"`
	// Sequence number of the event. The sequence number is a value available in some events, to make the exact ordering of events unambiguous, regardless of the event time precision.
	Sequence int `json:"sequence,omitempty"`
	// The unique tenant identifier.
	TenantUid string `json:"tenant_uid"`
	// The logging system-assigned unique identifier of an event instance.
	Uid string `json:"uid,omitempty"`
	// The version of the OCSF schema, using Semantic Versioning Specification (<a target='_blank' href='https://semver.org'>SemVer</a>). For example: 1.0.0. Event consumers use the version to determine the available event attributes.
	Version string `json:"version"`
}

// Metric The Metric object defines a simple name/value pair entity for a metric.
type Metric struct {
	// The name of the metric.
	Name string `json:"name"`
	// The value of the metric.
	Value string `json:"value"`
}

// Object An unordered collection of attributes. It defines a set of attributes available in all objects. It can be also used as a generic object to log objects that are not otherwise defined by the schema.
type Object struct {
}

// Observable The observable object is a pivot element that contains related information found in many places in the event.
type Observable struct {
	// The full name of the observable attribute. The <code>name</code> is a pointer/reference to an attribute within the event data. For example: <code>file.name</code>.
	Name string `json:"name"`
	// Contains the original and normalized reputation scores.
	Reputation *Reputation `json:"reputation,omitempty"`
	// The observable value type name.
	Type string `json:"type,omitempty"`
	// The observable value type identifier.
	TypeId int `json:"type_id"`
	// The value associated with the observable attribute. The meaning of the value depends on the observable type.<br/>If the <code>name</code> refers to a scalar attribute, then the <code>value</code> is the value of the attribute.<br/>If the <code>name</code> refers to an object attribute, then the <code>value</code> is not populated.
	Value string `json:"value,omitempty"`
}

// Package The Software Package object describes details about a software package. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:SoftwarePackage/'>d3f:SoftwarePackage</a>.
type Package struct {
	// Architecture is a shorthand name describing the type of computer hardware the packaged software is meant to run on.
	Architecture string `json:"architecture"`
	// The Common Platform Enumeration (CPE) name as described by (<a target='_blank' href='https://nvd.nist.gov/products/cpe'>NIST</a>) For example: <code>cpe:/a:apple:safari:16.2</code>.
	CpeName string `json:"cpe_name,omitempty"`
	// The software package epoch. Epoch is a way to define weighted dependencies based on version numbers.
	Epoch int `json:"epoch,omitempty"`
	// Cryptographic hash to identify the binary instance of a software component. This can include any component such file, package, or library.
	Hash *Fingerprint `json:"hash,omitempty"`
	// The software license applied to this package.
	License string `json:"license,omitempty"`
	// The software package name.
	Name string `json:"name"`
	// A purl is a URL string used to identify and locate a software package in a mostly universal and uniform way across programming languages, package managers, packaging conventions, tools, APIs and databases.
	Purl string `json:"purl,omitempty"`
	// Release is the number of times a version of the software has been packaged.
	Release string `json:"release,omitempty"`
	// The type of software package, normalized to the caption of the type_id value. In the case of 'Other', it is defined by the source.
	Type string `json:"type,omitempty"`
	// The type of software package.
	TypeId int `json:"type_id"`
	// The name of the vendor who published the software package.
	VendorName string `json:"vendor_name,omitempty"`
	// The software package version.
	Version string `json:"version"`
}

// QueryInfo The query info object holds information related to data access within a datastore. To access, manipulate, delete, or retrieve data from a datastore, a query must be written using a specific syntax.
type QueryInfo struct {
	// The size of the data returned from the query.
	Bytes int64 `json:"bytes,omitempty"`
	// The data returned from the query execution.
	Data json.RawMessage `json:"data,omitempty"`
	// The query name for a saved or scheduled query.
	Name string `json:"name"`
	// A string representing the query code being run. For example: <code>SELECT * FROM my_table</code>
	QueryString string `json:"query_string"`
	// The time when the query was run.
	QueryTime time.Time `json:"query_time,omitempty"`
	// The time when the query was run.
	QueryTimeDt time.Time `json:"query_time_dt,omitempty"`
	// The unique identifier of the query.
	Uid string `json:"uid"`
}

// RelatedEvent The Related Event object describes an OCSF event related to a finding.
type RelatedEvent struct {
	// An array of <a target='_blank' href='https://attack.mitre.org'>MITRE ATT&CK®</a> objects describing the tactics, techniques & sub-techniques identified by a security control or finding.
	Attacks []*Attack `json:"attacks,omitempty"`
	// The <a target='_blank' href='https://www.lockheedmartin.com/en-us/capabilities/cyber/cyber-kill-chain.html'>Cyber Kill Chain®</a> provides a detailed description of each phase and its associated activities within the broader context of a cyber attack.
	KillChain []*KillChainPhase `json:"kill_chain,omitempty"`
	// The observables associated with the event or a finding.
	Observables []*Observable `json:"observables,omitempty"`
	// The unique identifier of the product that reported the related event.
	ProductUid string `json:"product_uid,omitempty"`
	// The type of the related event, as defined by <code>type_uid</code>. <p>For example: <code>Process Activity: Launch.</code></p>
	Type string `json:"type,omitempty"`
	// The type of the related OCSF event, as defined by <code>type_uid</code>. <p>For example: <code>Process Activity: Launch.</code></p>
	TypeName string `json:"type_name,omitempty"`
	// The unique identifier of the related OCSF event type. <p>For example: <code>100701.</code></p>
	TypeUid int64 `json:"type_uid"`
	// The unique identifier of the related OCSF event. This value must be equal to <code>metadata.uid</code> in the corresponding related event.
	Uid string `json:"uid"`
}

// Request The Request Elements object describes characteristics of an API request.
type Request struct {
	// When working with containerized applications, the set of containers which write to the standard the output of a particular logging driver. For example, this may be the set of containers involved in handling api requests and responses for a containerized application.
	Containers []*Container `json:"containers,omitempty"`
	// The additional data that is associated with the api request.
	Data json.RawMessage `json:"data,omitempty"`
	// The list of communication flags, normalized to the captions of the flag_ids values. In the case of 'Other', they are defined by the event source.
	Flags []string `json:"flags,omitempty"`
	// The unique request identifier.
	Uid string `json:"uid"`
}

// ResourceDetails The Resource Details object describes details about resources that were affected by the activity/event.
type ResourceDetails struct {
	// A list of <code>agent</code> objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`
	// The canonical cloud partition name to which the region is assigned (e.g. AWS Partitions: aws, aws-cn, aws-us-gov).
	CloudPartition string `json:"cloud_partition,omitempty"`
	// The criticality of the resource as defined by the event source.
	Criticality string `json:"criticality,omitempty"`
	// Additional data describing the resource.
	Data json.RawMessage `json:"data,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The name of the related resource group.
	Group *Group `json:"group,omitempty"`
	// The list of labels/tags associated to a resource.
	Labels []string `json:"labels,omitempty"`
	// The name of the resource.
	Name string `json:"name"`
	// The namespace is useful when similar entities exist that you need to keep separate.
	Namespace string `json:"namespace,omitempty"`
	// The identity of the service or user account that owns the resource.
	Owner *User `json:"owner"`
	// The cloud region of the resource.
	Region string `json:"region,omitempty"`
	// The resource type as defined by the event source.
	Type string `json:"type,omitempty"`
	// The unique identifier of the resource.
	Uid string `json:"uid"`
	// The version of the resource. For example <code>1.2.3</code>.
	Version string `json:"version,omitempty"`
}

// Response The Response Elements object describes characteristics of an API response.
type Response struct {
	// The numeric response sent to a request.
	Code int `json:"code"`
	// When working with containerized applications, the set of containers which write to the standard the output of a particular logging driver. For example, this may be the set of containers involved in handling api requests and responses for a containerized application.
	Containers []*Container `json:"containers,omitempty"`
	// The additional data that is associated with the api response.
	Data json.RawMessage `json:"data,omitempty"`
	// Error Code
	Error string `json:"error"`
	// Error Message
	ErrorMessage string `json:"error_message"`
	// The list of communication flags, normalized to the captions of the flag_ids values. In the case of 'Other', they are defined by the event source.
	Flags []string `json:"flags,omitempty"`
	// The description of the event/finding, as defined by the source.
	Message string `json:"message"`
}

// Rule The Rule object describes characteristics of a rule associated with a policy or an event.
type Rule struct {
	// The rule category.
	Category string `json:"category,omitempty"`
	// The description of the rule that generated the event.
	Desc string `json:"desc,omitempty"`
	// The name of the rule that generated the event.
	Name string `json:"name"`
	// The rule type.
	Type string `json:"type,omitempty"`
	// The unique identifier of the rule that generated the event.
	Uid string `json:"uid"`
	// The rule version. For example: <code>1.1</code>.
	Version string `json:"version,omitempty"`
}

// San The Subject Alternative name (SAN) object describes a SAN secured by a digital certificate
type San struct {
	// Name of SAN (e.g. The actual IP Address or domain.)
	Name string `json:"name"`
	// Type descriptor of SAN (e.g. IP Address/domain/etc.)
	Type string `json:"type"`
}

// Table The table object represents a table within a structured relational database or datastore, which contains columns and rows of data that are able to be create, updated, deleted and queried.
type Table struct {
	// The time when the table was known to have been created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the table was known to have been created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The description of the table.
	Desc string `json:"desc,omitempty"`
	// The group names to which the table belongs.
	Groups []*Group `json:"groups,omitempty"`
	// The most recent time when any changes, updates, or modifications were made within the table.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The most recent time when any changes, updates, or modifications were made within the table.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The table name, ordinarily as assigned by a database administrator.
	Name string `json:"name"`
	// The size of the data table in bytes.
	Size int64 `json:"size,omitempty"`
	// The unique identifier of the table.
	Uid string `json:"uid"`
}

// Ticket The Ticket object represents ticket in the customer's systems like Salesforce, jira etc.
type Ticket struct {
	// The url of a ticket in the ticket system.
	SrcUrl string `json:"src_url"`
	// The title of the ticket.
	Title string `json:"title,omitempty"`
	// The linked ticket type determines whether the ticket is internal or in an external ticketing system.
	Type string `json:"type,omitempty"`
	// The normalized identifier for the ticket type.
	TypeId int `json:"type_id,omitempty"`
	// Unique ticket identifier like ticket id.
	Uid string `json:"uid"`
}

// Timespan The Time Span object represents different time period durations. If a timespan is fractional, i.e. crosses one period, e.g. a week and 3 days, more than one may may be populated since each member is of integral type. In that case <code>type_id</code> if present should be set to <code>Other</code>.
type Timespan struct {
	// The duration of the time span in milliseconds.
	Duration int64 `json:"duration"`
	// The duration of the time span in days.
	DurationDays int `json:"duration_days"`
	// The duration of the time span in hours.
	DurationHours int `json:"duration_hours"`
	// The duration of the time span in minutes.
	DurationMins int `json:"duration_mins"`
	// The duration of the time span in months.
	DurationMonths int `json:"duration_months"`
	// The duration of the time span in seconds.
	DurationSecs int `json:"duration_secs"`
	// The duration of the time span in weeks.
	DurationWeeks int `json:"duration_weeks"`
	// The duration of the time span in years.
	DurationYears int `json:"duration_years"`
	// The type of time span duration the object represents.
	Type string `json:"type,omitempty"`
	// The normalized identifier for the time span duration type.
	TypeId int `json:"type_id"`
}

// Whois The resources of a WHOIS record for a given domain. This can include domain names, IP address blocks, autonomous system information, and/or contact and registration information for a domain.
type Whois struct {
	// The autonomous system information associated with a domain.
	AutonomousSystem *AutonomousSystem `json:"autonomous_system,omitempty"`
	// When the domain was registered or WHOIS entry was created.
	CreatedTime time.Time `json:"created_time"`
	// When the domain was registered or WHOIS entry was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The normalized value of dnssec_status_id.
	DnssecStatus string `json:"dnssec_status,omitempty"`
	// Describes the normalized status of DNS Security Extensions (DNSSEC) for a domain.
	DnssecStatusId int `json:"dnssec_status_id"`
	// The name of the domain.
	Domain string `json:"domain"`
	// An array of <code>Domain Contact</code> objects.
	DomainContacts []*DomainContact `json:"domain_contacts"`
	// The email address for the registrar's abuse contact
	EmailAddr string `json:"email_addr,omitempty"`
	// When the WHOIS record was last updated or seen at.
	LastSeenTime time.Time `json:"last_seen_time"`
	// When the WHOIS record was last updated or seen at.
	LastSeenTimeDt time.Time `json:"last_seen_time_dt,omitempty"`
	// A collection of name servers related to a domain registration or other record.
	NameServers []string `json:"name_servers"`
	// The phone number for the registrar's abuse contact
	PhoneNumber string `json:"phone_number,omitempty"`
	// The domain registrar.
	Registrar string `json:"registrar"`
	// The status of a domain and its ability to be transferred, e.g., <code>clientTransferProhibited</code>.
	Status string `json:"status"`
	// An array of subdomain strings. Can be used to collect several subdomains such as those from Domain Generation Algorithms (DGAs).
	Subdomains []string `json:"subdomains,omitempty"`
	// The IP address block (CIDR) associated with a domain.
	Subnet string `json:"subnet,omitempty"`
}


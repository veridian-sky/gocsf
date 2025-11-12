package dictionary_attributes

// Code generated from OCSF schema; DO NOT EDIT.

import "github.com/veridian-sky/gocsf/objects"

// DictionaryAttributes contains all OCSF dictionary attributes
type DictionaryAttributes struct {
	// Describes access relationships and pathways between identities, resources,
	// focusing on who can access what and through which mechanisms. This evaluates
	// access levels (read/write/admin), access types (direct, cross-account,
	// public, federated), and the conditions under which access is granted. Use
	// this for resource-centric security assessments such as external access
	// discovery, public exposure analysis, etc.
	AccessAnalysisResult *objects.AccessAnalysisResult `json:"access_analysis_result,omitempty"`

	// The access level of an entity. See specific usage.
	AccessLevel string `json:"access_level,omitempty"`

	// The list of requested access rights.
	AccessList []string `json:"access_list,omitempty"`

	// The access mask in a platform-native format.
	AccessMask int64 `json:"access_mask,omitempty"`

	// The list of access check results.
	AccessResult interface{} `json:"access_result,omitempty"`

	// The type or category of access being granted to the identity. See specific
	// usage.
	AccessType string `json:"access_type,omitempty"`

	// The time when the file was last accessed.
	AccessedTime int64 `json:"accessed_time,omitempty"`

	// The time when the file was last accessed.
	AccessedTimeDt string `json:"accessed_time_dt,omitempty"`

	// The name of the user who last accessed the object.
	Accessor *objects.User `json:"accessor,omitempty"`

	// A list of users who have access to an entity. See specific usage.
	Accessors []*objects.User `json:"accessors,omitempty"`

	// The account object describes details about the account that was the source
	// or target of the activity.
	Account *objects.Account `json:"account,omitempty"`

	// The account switch method, normalized to the caption of the
	// account_switch_type_id value. In the case of 'Other', it is defined by the
	// event source.
	AccountSwitchType string `json:"account_switch_type,omitempty"`

	// The normalized identifier of the account switch method.
	AccountSwitchTypeID int `json:"account_switch_type_id,omitempty"`

	// An integer that provides a reason code or additional information about the
	// acknowledgment result.
	AckReason int64 `json:"ack_reason,omitempty"`

	// An integer that denotes the acknowledgment result of the DCE/RPC call.
	AckResult int64 `json:"ack_result,omitempty"`

	// The normalized caption of 'action_id' or the source specific action.
	Action string `json:"action,omitempty"`

	// The normalized action taken by a control or other policy-based system
	// leading to an outcome or disposition.
	ActionID int `json:"action_id,omitempty"`

	// The normalized identifier of the activity that triggered the event.
	ActivityID int `json:"activity_id,omitempty"`

	// The event activity name, as defined by the activity_id.
	ActivityName string `json:"activity_name,omitempty"`

	// The actor object describes details about the user/role/process that was the
	// source of the activity. Note that this is not the threat actor of a campaign
	// but may be part of a campaign.
	Actor *objects.Actor `json:"actor,omitempty"`

	// The permissions that were granted in a platform-native format. See specific
	// usage.
	ActualPermissions int64 `json:"actual_permissions,omitempty"`

	// The supplementary restrictions that may apply to an entity, by the virtue of
	// a policy. See specific usage.
	AdditionalRestrictions []*objects.AdditionalRestriction `json:"additional_restrictions,omitempty"`

	// Detail about the security advisory, that is used to publicly disclose
	// cybersecurity vulnerabilities by a vendor.
	Advisory *objects.Advisory `json:"advisory,omitempty"`

	// Expressed as either height above takeoff location or height above ground
	// level (AGL) for a UAS current location. This value is provided in meters and
	// must have a minimum resolution of 1 m. Special Values: Invalid, No Value, or
	// Unknown: -1000 m.
	AerialHeight string `json:"aerial_height,omitempty"`

	// List of Affected Code objects that describe details about code blocks
	// identified as vulnerable.
	AffectedCode []*objects.AffectedCode `json:"affected_code,omitempty"`

	// List of software packages identified as affected by a
	// vulnerability/vulnerabilities.
	AffectedPackages []*objects.AffectedPackage `json:"affected_packages,omitempty"`

	// An Agent (also known as a Sensor) is typically installed on an Operating
	// System (OS) and serves as a specialized software component that can be
	// designed to monitor, detect, collect, archive, or take action. These
	// activities and possible actions are defined by the upstream system
	// controlling the Agent and its intended purpose. For instance, an Agent can
	// include Endpoint Detection & Response (EDR) agents, backup/disaster recovery
	// sensors, Application Performance Monitoring or profiling sensors, and
	// similar software.
	Agent *objects.Agent `json:"agent,omitempty"`

	// A list of agent objects associated with a device, endpoint, or resource.
	AgentList []*objects.Agent `json:"agent_list,omitempty"`

	// The Aircraft object represents any aircraft or otherwise airborne asset such
	// as an unmanned system, airplane, balloon, spacecraft, or otherwise. The
	// Aircraft object is intended to normalized data captured or otherwise logged
	// from active radar, passive radar, multi-spectral systems, or the Automatic
	// Dependant Broadcast - Surveillance (ADS-B), and/or Mode S systems.
	Aircraft *objects.Aircraft `json:"aircraft,omitempty"`

	// The integer value of TLS alert if present. The alerts are defined in the TLS
	// specification in https://datatracker.ietf.org/doc/html/rfc2246 RFC-2246.
	Alert int64 `json:"alert,omitempty"`

	// The applicable algorithm, normalized to the caption of 'algorithm_id'. See
	// specific usage.
	Algorithm string `json:"algorithm,omitempty"`

	// The normalized identifier of the algorithm. See specific usage.
	AlgorithmID int `json:"algorithm_id,omitempty"`

	// Maximum altitude (WGS-84 HAE) for a group or an Intent-Based Network
	// Participant. Measured in meters. Special Values: Invalid, No Value, or
	// Unknown: -1000 m.
	AltitudeCeiling string `json:"altitude_ceiling,omitempty"`

	// Minimum altitude (WGS-84 HAE) for a group or an Intent-Based Network
	// Participant. Measured in meters. Special Values: Invalid, No Value, or
	// Unknown: -1000 m.
	AltitudeFloor string `json:"altitude_floor,omitempty"`

	// The specific dimensions, components, or aspects of the system that are the
	// targets of the analysis. See specific usage.
	AnalysisTargets []*objects.AnalysisTarget `json:"analysis_targets,omitempty"`

	// The analytic technique used to analyze and derive insights from the data or
	// information that led to the finding or conclusion.
	Analytic *objects.Analytic `json:"analytic,omitempty"`

	// An array of Process Entities describing the extended parentage of this
	// process object. Direct parent information should be expressed through the
	// parent_process attribute. The first array element is the direct parent of
	// this process object. Subsequent list elements go up the process parentage
	// hierarchy. That is, the array is sorted from newest to oldest process. It is
	// recommended to only populate this field for the top-level process object.
	Ancestry []*objects.ProcessEntity `json:"ancestry,omitempty"`

	// A list of detected anomalies or deviations from expected behavior patterns.
	// See specific usage.
	Anomalies []*objects.Anomaly `json:"anomalies,omitempty"`

	// A list of anomaly analysis results that examine and characterize patterns of
	// activity or usage over time to identify normal vs abnormal activities. See
	// specific usage.
	AnomalyAnalyses []*objects.AnomalyAnalysis `json:"anomaly_analyses,omitempty"`

	// The Domain Name System (DNS) answers.
	Answers []*objects.DnsAnswer `json:"answers,omitempty"`

	// Describes details about a typical API (Application Programming Interface)
	// call.
	Api *objects.Api `json:"api,omitempty"`

	// The application that reported the event.
	App *objects.Product `json:"app,omitempty"`

	// The name of the application associated with the event or object.
	AppName string `json:"app_name,omitempty"`

	// The unique ID of the application associated with the event or object.
	AppUID string `json:"app_uid,omitempty"`

	// An Application describes the details for an inventoried application as
	// reported by an Application Security tool or other Developer-centric tooling.
	// Applications can be defined as Kubernetes resources, Containerized
	// resources, or application hosting-specific cloud sources such as AWS Elastic
	// BeanStalk, AWS Lightsail, or Azure Logic Apps.
	Application *objects.Application `json:"application,omitempty"`

	// A list of application objects. See specific usage.
	Applications []*objects.Application `json:"applications,omitempty"`

	// Architecture is a shorthand name describing the type of computer hardware
	// the packaged software is meant to run on.
	Architecture string `json:"architecture,omitempty"`

	// The arguments sent along with the HTTP request.
	Args string `json:"args,omitempty"`

	// The Assessment object describes a point-in-time assessment, check, or
	// evaluation of a specific configuration or signal against an asset, entity,
	// person, or otherwise. For example, this can encapsulate os_signals from
	// CrowdStrike Falcon Zero Trust Assessments, or account for Datastore
	// configurations from Cyera.
	Assessment *objects.Assessment `json:"assessment,omitempty"`

	// A list of assessment objects. See specific usage.
	Assessments []*objects.Assessment `json:"assessments,omitempty"`

	// The details of the user assigned to an Incident.
	Assignee *objects.User `json:"assignee,omitempty"`

	// The details of the group assigned to an Incident.
	AssigneeGroup *objects.Group `json:"assignee_group,omitempty"`

	// An Attack Graph describes possible routes an attacker could take through an
	// environment. It describes relationships between resources and their
	// findings, such as malware detections, vulnerabilities, misconfigurations,
	// and other security actions.
	AttackGraph *objects.Graph `json:"attack_graph,omitempty"`

	// An array of MITRE ATT&CK® objects describing identified tactics, techniques
	// & sub-techniques. The objects are compatible with MITRE ATLAS™ tactics,
	// techniques & sub-techniques.
	Attacks []*objects.Attack `json:"attacks,omitempty"`

	// The delivery attempt.
	Attempt int64 `json:"attempt,omitempty"`

	// The bitmask value that represents the file attributes.
	Attributes int64 `json:"attributes,omitempty"`

	// Describes a category of methods used for identity verification in an
	// authentication attempt.
	AuthFactors []*objects.AuthFactor `json:"auth_factors,omitempty"`

	// The authentication protocol as defined by the caption of auth_protocol_id.
	// In the case of Other, it is defined by the event source.
	AuthProtocol string `json:"auth_protocol,omitempty"`

	// The normalized identifier of the authentication protocol used to create the
	// user session.
	AuthProtocolID int `json:"auth_protocol_id,omitempty"`

	// The agreed upon authentication type, normalized to the caption of
	// 'auth_type_id'. In the case of 'Other', it is defined by the event source.
	AuthType string `json:"auth_type,omitempty"`

	// The normalized identifier of the agreed upon authentication type. See
	// specific usage.
	AuthTypeID int `json:"auth_type_id,omitempty"`

	// The authentication token, ticket, or assertion. See specific usage.
	AuthenticationToken *objects.AuthenticationToken `json:"authentication_token,omitempty"`

	// The author(s) who published the software component.
	Author string `json:"author,omitempty"`

	// Provides details about an authorization, such as authorization outcome, and
	// any associated policies related to the activity/event.
	Authorizations []*objects.Authorization `json:"authorizations,omitempty"`

	// The Autonomous System details associated with an IP address.
	AutonomousSystem *objects.AutonomousSystem `json:"autonomous_system,omitempty"`

	// The unique identifier of the cloud autoscale configuration.
	AutoscaleUID string `json:"autoscale_uid,omitempty"`

	// The average time span of an activity.
	AvgTimespan *objects.Timespan `json:"avg_timespan,omitempty"`

	// The initial connection response that a messaging server receives after it
	// connects to an email server.
	Banner string `json:"banner,omitempty"`

	// The memory address where the module was loaded.
	BaseAddress string `json:"base_address,omitempty"`

	// The base score as reported by the event source. See specific usage.
	BaseScore float64 `json:"base_score,omitempty"`

	// A list of baseline measurements or normal behavior patterns used as
	// reference points for comparison and anomaly detection. See specific usage.
	Baselines []*objects.Baseline `json:"baselines,omitempty"`

	// The BIOS date. For example: 03/31/16.
	BiosDate string `json:"bios_date,omitempty"`

	// The BIOS manufacturer. For example: LENOVO.
	BiosManufacturer string `json:"bios_manufacturer,omitempty"`

	// The BIOS version. For example: LENOVO G5ETA2WW (2.62).
	BiosVer string `json:"bios_ver,omitempty"`

	// The actual length of the HTTP response/request body, in number of bytes,
	// independent of a potentially existing Content-Length header.
	BodyLength int64 `json:"body_length,omitempty"`

	// The time when the system was booted.
	BootTime int64 `json:"boot_time,omitempty"`

	// The time when the system was booted.
	BootTimeDt string `json:"boot_time_dt,omitempty"`

	// A unique identifier of the device that changes after every reboot. For
	// example, the value of /proc/sys/kernel/random/boot_id from Linux's procfs.
	BootUID string `json:"boot_uid,omitempty"`

	// The boundary of the connection, normalized to the caption of 'boundary_id'.
	// In the case of 'Other', it is defined by the event source. <p> For cloud
	// connections, this translates to the traffic-boundary(same VPC, through IGW,
	// etc.). For traditional networks, this is described as Local, Internal, or
	// External.</p>
	Boundary string `json:"boundary,omitempty"`

	// <p>The normalized identifier of the boundary of the connection. </p><p> For
	// cloud connections, this translates to the traffic-boundary (same VPC,
	// through IGW, etc.). For traditional networks, this is described as Local,
	// Internal, or External.</p>
	BoundaryID int `json:"boundary_id,omitempty"`

	// The operating system build number.
	Build string `json:"build,omitempty"`

	// The vendor bulletin identifier.
	Bulletin string `json:"bulletin,omitempty"`

	// The total number of bytes (in and out).
	Bytes int64 `json:"bytes,omitempty"`

	// The number of bytes sent from the destination to the source.
	BytesIn int64 `json:"bytes_in,omitempty"`

	// Indicates the number of bytes missed, which is representative of packet
	// loss.
	BytesMissed int64 `json:"bytes_missed,omitempty"`

	// The number of bytes sent from the source to the destination.
	BytesOut int64 `json:"bytes_out,omitempty"`

	// The campaign object describes details about the campaign that was the source
	// of the activity.
	Campaign *objects.Campaign `json:"campaign,omitempty"`

	// A list of RDP capabilities.
	Capabilities []string `json:"capabilities,omitempty"`

	// A short description or caption of the device. For example: Scanner 1 or
	// Database Manager.
	Caption string `json:"caption,omitempty"`

	// The Website categorization names, as defined by category_ids enum values.
	Categories []string `json:"categories,omitempty"`

	// The object category, normalized to the caption of category_id. See specific
	// usage.
	Category string `json:"category,omitempty"`

	// The normalized identifier of the object category. See specific usage.
	CategoryID int64 `json:"category_id,omitempty"`

	// The Website categorization identifiers.
	CategoryIDs []int `json:"category_ids,omitempty"`

	// The event category name, as defined by category_uid value.
	CategoryName string `json:"category_name,omitempty"`

	// The category unique identifier of the event.
	CategoryUID int64 `json:"category_uid,omitempty"`

	// The machine-readable email header Cc values, as defined by RFC 5322. For
	// example example.user@usersdomain.com.
	Cc []string `json:"cc,omitempty"`

	// The human-readable email header Cc Mailbox values. For example 'Example User
	// &lt;example.user@usersdomain.com&gt;'.
	CcMailboxes []string `json:"cc_mailboxes,omitempty"`

	// The name of the cell. See specific usage.
	CellName string `json:"cell_name,omitempty"`

	// The certificate object containing information about the digital certificate.
	Certificate *objects.Certificate `json:"certificate,omitempty"`

	// The Chain of Certificate Serial Numbers field provides a chain of
	// Certificate Issuer Serial Numbers leading to the Root Certificate Issuer.
	CertificateChain []string `json:"certificate_chain,omitempty"`

	// The chassis type describes the system enclosure or physical form factor.
	// Such as the following examples for Windows
	// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-systemenclosure
	// Windows Chassis Types
	Chassis string `json:"chassis,omitempty"`

	// A list of specific, individual compliance verification checks derived from
	// standard or non-standard benchmark rules or requirements.
	Checks []*objects.Check `json:"checks,omitempty"`

	// A unit of information within an SCTP packet, consisting of a chunk header
	// and chunk-specific content. See RFC 4960. This field can be used for the
	// total number of chunks (in and out).
	Chunks int64 `json:"chunks,omitempty"`

	// A unit of information within an SCTP packet, consisting of a chunk header
	// and chunk-specific content. See RFC 4960. This field can be used for the
	// chunks sent from the destination to the source.
	ChunksIn int64 `json:"chunks_in,omitempty"`

	// A unit of information within an SCTP packet, consisting of a chunk header
	// and chunk-specific content. See RFC 4960. This field can be used for the
	// chunks sent from the source to the destination.
	ChunksOut int64 `json:"chunks_out,omitempty"`

	// The negotiated cipher suite.
	Cipher string `json:"cipher,omitempty"`

	// The CIS Benchmark describes best practices for securely configuring IT
	// systems, software, networks, and cloud infrastructure as defined by the
	// Center for Internet Security (https://www.cisecurity.org/cis-benchmarks/
	// CIS).
	CisBenchmark *objects.CisBenchmark `json:"cis_benchmark,omitempty"`

	// The CIS Benchmark Result object captures results generated from benchmark
	// evaluations as defined by the Center for Internet Security
	// (https://www.cisecurity.org/cis-benchmarks/ CIS).
	CisBenchmarkResult *objects.CisBenchmarkResult `json:"cis_benchmark_result,omitempty"`

	// The CIS Critical Security Controls is a prioritized set of actions to
	// protect your organization and data from cyber-attack vectors.
	CisControls []*objects.CisControl `json:"cis_controls,omitempty"`

	// The CIS Critical Security Controls is a list of top 20 actions and practices
	// an organization’s security team can take on such that cyber attacks or
	// malware, are minimized and prevented.
	CisCsc []*objects.CisCsc `json:"cis_csc,omitempty"`

	// The name of the city.
	City string `json:"city,omitempty"`

	// The class name of the object. See specific usage.
	Class string `json:"class,omitempty"`

	// The event class name, as defined by class_uid value.
	ClassName string `json:"class_name,omitempty"`

	// The unique identifier of a class. A class describes the attributes available
	// in an event.
	ClassUID int64 `json:"class_uid,omitempty"`

	// The classification as defined by the vendor.
	Classification string `json:"classification,omitempty"`

	// The list of normalized classification identifiers. See specific usage.
	ClassificationIDs []int `json:"classification_ids,omitempty"`

	// The list of malware classifications, normalized to the captions of the
	// classification_id values. In the case of 'Other', they are defined by the
	// event source.
	Classifications []string `json:"classifications,omitempty"`

	// Describes details about the classifier used for data classification.
	ClassifierDetails *objects.ClassifierDetails `json:"classifier_details,omitempty"`

	// The client cipher suites that were exchanged during the TLS handshake
	// negotiation.
	ClientCiphers []string `json:"client_ciphers,omitempty"`

	// The list of SMB dialects that the client speaks.
	ClientDialects []string `json:"client_dialects,omitempty"`

	// The Client HASSH fingerprinting object.
	ClientHassh *objects.Hassh `json:"client_hassh,omitempty"`

	// Describes details about the Cloud environment where the event was originally
	// created or logged.
	Cloud *objects.Cloud `json:"cloud,omitempty"`

	// The canonical cloud partition name to which the region is assigned (e.g. AWS
	// Partitions: aws, aws-cn, aws-us-gov).
	CloudPartition string `json:"cloud_partition,omitempty"`

	// The full command line used to launch an application, service, process, or
	// job. For example: ssh user@10.0.0.10. If the command line is unavailable or
	// missing, the empty string '' is to be used.
	CmdLine string `json:"cmd_line,omitempty"`

	// The numeric response sent to a request.
	Code int64 `json:"code,omitempty"`

	// The list of numeric responses sent to a request.
	Codes []int64 `json:"codes,omitempty"`

	// The numeric color depth.
	ColorDepth int64 `json:"color_depth,omitempty"`

	// The name of the column. See specific usage.
	ColumnName string `json:"column_name,omitempty"`

	// The number of the column. See specific usage.
	ColumnNumber int64 `json:"column_number,omitempty"`

	// The command name.
	Command string `json:"command,omitempty"`

	// The response to the command.
	CommandResponse string `json:"command_response,omitempty"`

	// The responses to the command.
	CommandResponses []string `json:"command_responses,omitempty"`

	// The unique command identifier.
	CommandUID string `json:"command_uid,omitempty"`

	// The user-provided comment.
	Comment string `json:"comment,omitempty"`

	// The Community ID of the network connection.
	CommunityUID string `json:"community_uid,omitempty"`

	// The name of the company that published the file. For example: Microsoft
	// Corporation.
	CompanyName string `json:"company_name,omitempty"`

	// The compliance object provides context to compliance findings (e.g., a check
	// against a specific regulatory or best practice framework such as CIS, NIST
	// etc.) and contains compliance related details.
	Compliance *objects.Compliance `json:"compliance,omitempty"`

	// A list of reference KB articles that provide information to help
	// organizations understand, interpret, and implement compliance standards.
	// They provide guidance, best practices, and examples.
	ComplianceReferences []*objects.KbArticle `json:"compliance_references,omitempty"`

	// A list of established guidelines or criteria that define specific
	// requirements an organization must follow.
	ComplianceStandards []*objects.KbArticle `json:"compliance_standards,omitempty"`

	// The component of a data object. See specific usage.
	Component string `json:"component,omitempty"`

	// The condition that was evaluated in a rule, policy. See specific usage.
	Condition string `json:"condition,omitempty"`

	// The list of condition keys and their values that were evaluated as part of a
	// rule or policy. Condition keys are used to define circumstances under which
	// the rule/policy is in effect. See specific usage.
	ConditionKeys []*objects.KeyValueObject `json:"condition_keys,omitempty"`

	// The confidence, normalized to the caption of the confidence_id value. In the
	// case of 'Other', it is defined by the event source.
	Confidence string `json:"confidence,omitempty"`

	// The normalized confidence refers to the accuracy of the rule that created
	// the finding. A rule with a low confidence means that the finding scope is
	// wide and may create finding reports that may not be malicious in nature.
	ConfidenceID int `json:"confidence_id,omitempty"`

	// The confidence score as reported by the event source.
	ConfidenceScore int64 `json:"confidence_score,omitempty"`

	// The file content confidentiality, normalized to the confidentiality_id
	// value. In the case of 'Other', it is defined by the event source.
	Confidentiality string `json:"confidentiality,omitempty"`

	// The normalized identifier of the file content confidentiality indicator.
	ConfidentialityID int `json:"confidentiality_id,omitempty"`

	// The network connection information.
	ConnectionInfo *objects.NetworkConnectionInfo `json:"connection_info,omitempty"`

	// The network connection identifier.
	ConnectionUID string `json:"connection_uid,omitempty"`

	// The information describing an instance of a container. A container is a
	// prepackaged, portable system image that runs isolated on an existing system
	// using a container runtime like containerd.
	Container *objects.Container `json:"container,omitempty"`

	// When working with containerized applications, the set of containers which
	// write to the standard the output of a particular logging driver. For
	// example, this may be the set of containers involved in handling api requests
	// and responses for a containerized application.
	Containers []*objects.Container `json:"containers,omitempty"`

	// The request header that identifies the original
	// https://www.iana.org/assignments/media-types/media-types.xhtml media type of
	// the resource (prior to any content encoding applied for sending).
	ContentType string `json:"content_type,omitempty"`

	// The name of the continent.
	Continent string `json:"continent,omitempty"`

	// A Control is prescriptive, prioritized, and simplified set of best practices
	// that one can use to strengthen their cybersecurity posture. e.g. AWS
	// SecurityHub Controls, CIS Controls.
	Control string `json:"control,omitempty"`

	// The list of control parameters evaluated in a Compliance check.
	ControlParameters []*objects.KeyValueObject `json:"control_parameters,omitempty"`

	// A two-element array, containing a longitude/latitude pair. The format
	// conforms with https://geojson.org GeoJSON. For example: [-73.983, 40.719].
	Coordinates []float64 `json:"coordinates,omitempty"`

	// The unique identifier used to correlate events.
	CorrelationUID string `json:"correlation_uid,omitempty"`

	// The cost center associated with the user.
	CostCenter string `json:"cost_center,omitempty"`

	// The number of times that events in the same logical group occurred during
	// the event Start Time to End Time period.
	Count int64 `json:"count,omitempty"`

	// The MITRE D3FEND™ Matrix Countermeasures associated with a remediation.
	Countermeasures []*objects.D3fend `json:"countermeasures,omitempty"`

	// The ISO 3166-1 Alpha-2 country code.<p><b>Note:</b> The two letter country
	// code should be capitalized. For example: US or CA.</p>
	Country string `json:"country,omitempty"`

	// The Common Platform Enumeration (CPE) name as described by
	// (https://nvd.nist.gov/products/cpe NIST) For example:
	// cpe:/a:apple:safari:16.2.
	CpeName string `json:"cpe_name,omitempty"`

	// A unique process identifier that can be assigned deterministically by
	// multiple system data producers.
	Cpid string `json:"cpid,omitempty"`

	// The CPU architecture, normalized to the caption of the cpu_architecture_id
	// value. In the case of Other, it is defined by the source.
	CpuArchitecture string `json:"cpu_architecture,omitempty"`

	// The normalized identifier of the CPU architecture.
	CpuArchitectureID int `json:"cpu_architecture_id,omitempty"`

	// The cpu architecture, the number of bits used for addressing in memory. For
	// example: 32 or 64.
	CpuBits int64 `json:"cpu_bits,omitempty"`

	// The number of processor cores in all installed processors. For Example: 42.
	CpuCores int64 `json:"cpu_cores,omitempty"`

	// The number of physical processors on a system. For example: 1.
	CpuCount int64 `json:"cpu_count,omitempty"`

	// The speed of the processor in Mhz. For Example: 4200.
	CpuSpeed int64 `json:"cpu_speed,omitempty"`

	// The processor type. For example: x86 Family 6 Model 37 Stepping 5.
	CpuType string `json:"cpu_type,omitempty"`

	// The original Windows mask that is required to create the object.
	CreateMask string `json:"create_mask,omitempty"`

	// The time when the object was created. See specific usage.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the object was created. See specific usage.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The user that created the object associated with event. See specific usage.
	Creator *objects.User `json:"creator,omitempty"`

	// The unique identifier of the user's credential. For example, AWS Access Key
	// ID.
	CredentialUID string `json:"credential_uid,omitempty"`

	// Criticality of a resource/object in question
	Criticality string `json:"criticality,omitempty"`

	// The unique customer identifier.
	CustomerUID string `json:"customer_uid,omitempty"`

	// The Common Vulnerabilities and Exposures (https://cve.mitre.org/ CVE)
	// identifiers for the vulnerability.
	Cve *objects.Cve `json:"cve,omitempty"`

	// List of Common Vulnerabilities and Exposures (https://cve.mitre.org/ CVE).
	Cves []*objects.Cve `json:"cves,omitempty"`

	// The CVSS object details Common Vulnerability Scoring System
	// (https://www.first.org/cvss/ CVSS) scores from the advisory that are related
	// to the vulnerability.
	Cvss []*objects.Cvss `json:"cvss,omitempty"`

	// The CWE object represents a weakness in a software system that can be
	// exploited by a threat actor to perform an attack. The CWE object is based on
	// the https://cwe.mitre.org/ Common Weakness Enumeration (CWE) catalog.
	Cwe *objects.Cwe `json:"cwe,omitempty"`

	// The https://cwe.mitre.org/ Common Weakness Enumeration (CWE) unique
	// identifier. For example: CWE-787.
	CweUID string `json:"cwe_uid,omitempty"`

	// Common Weakness Enumeration (CWE) definition URL. For example:
	// https://cwe.mitre.org/data/definitions/787.html.
	CweURL string `json:"cwe_url,omitempty"`

	// The D3FEND Tactic object describes the defensive tactic name associated with
	// a countermeasure.
	D3fTactic *objects.D3fTactic `json:"d3f_tactic,omitempty"`

	// The D3FEND Technique object describes the defensive technique ID and/or name
	// associated with a countermeasure.
	D3fTechnique *objects.D3fTechnique `json:"d3f_technique,omitempty"`

	// The additional data that is associated with the event or object. See
	// specific usage.
	Data interface{} `json:"data,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *objects.DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*objects.DataClassification `json:"data_classifications,omitempty"`

	// The name of the stage or state that the data was in. E.g., Data-at-Rest,
	// Data-in-Transit, etc.
	DataLifecycleState string `json:"data_lifecycle_state,omitempty"`

	// The stage or state that the data was in when it was assessed or scanned by a
	// data security tool.
	DataLifecycleStateID int `json:"data_lifecycle_state_id,omitempty"`

	// The Data Security object describes the characteristics, techniques and
	// content of a Data Loss Prevention (DLP), Data Loss Detection (DLD), Data
	// Classification, or similar tools' finding, alert, or detection mechanism(s).
	DataSecurity *objects.DataSecurity `json:"data_security,omitempty"`

	// A list of data sources utilized in generation of the finding.
	DataSources []string `json:"data_sources,omitempty"`

	// The database object is used for databases which are typically datastore
	// services that contain an organized collection of structured and unstructured
	// data or a types of data.
	Database *objects.Database `json:"database,omitempty"`

	// The data bucket object is a basic container that holds data, typically
	// organized through the use of data partitions.
	Databucket *objects.Databucket `json:"databucket,omitempty"`

	// The DCE/RPC object describes the remote procedure call system for
	// distributed computing environments.
	DceRpc *objects.DceRpc `json:"dce_rpc,omitempty"`

	// Debug information about non-fatal issues with this OCSF event. Each issue is
	// a line in this string array.
	Debug []string `json:"debug,omitempty"`

	// Decision/outcome of the authorization mechanism (e.g. Approved, Denied)
	Decision string `json:"decision,omitempty"`

	// The total round-trip delay to the reference clock in milliseconds.
	Delay int64 `json:"delay,omitempty"`

	// The timestamp when the user was deleted. In Active Directory (AD), when a
	// user is deleted they are moved to a temporary container and then removed
	// after 30 days. So, this field can be populated even after a user is deleted
	// for the next 30 days.
	DeletedTime int64 `json:"deleted_time,omitempty"`

	// The timestamp when the user was deleted. In Active Directory (AD), when a
	// user is deleted they are moved to a temporary container and then removed
	// after 30 days. So, this field can be populated even after a user is deleted
	// for the next 30 days.
	DeletedTimeDt string `json:"deleted_time_dt,omitempty"`

	// The machine-readable Delivered-To email header field. For example
	// example.user@usersdomain.com
	DeliveredTo string `json:"delivered_to,omitempty"`

	// The machine-readable Delivered-To email header values. For example
	// example.user@usersdomain.com
	DeliveredToList []string `json:"delivered_to_list,omitempty"`

	// Information about the chain of dependencies related to the issue as reported
	// by an Application Security or Vulnerability Management tool. E.g.,
	// serverless-offline -> @serverless/utils -> memoizee -> es5-ext.
	DependencyChain string `json:"dependency_chain,omitempty"`

	// The CVSS depth represents a depth of the equation used to calculate CVSS
	// score.
	Depth int `json:"depth,omitempty"`

	// The description that pertains to the object or event. See specific usage.
	Desc string `json:"desc,omitempty"`

	// The desktop display affiliated with the event
	DesktopDisplay *objects.Display `json:"desktop_display,omitempty"`

	// Details of an entity. See specific usage
	Details string `json:"details,omitempty"`

	// Specific pattern, algorithm, fingerprint, or model used for detection.
	DetectionPattern string `json:"detection_pattern,omitempty"`

	// The detection pattern type, normalized to the caption of the
	// detection_pattern_type_id value. In the case of 'Other', it is defined by
	// the event source.
	DetectionPatternType string `json:"detection_pattern_type,omitempty"`

	// Specifies the type of detection pattern used to identify the associated
	// threat indicator.
	DetectionPatternTypeID int `json:"detection_pattern_type_id,omitempty"`

	// The name of the type of data security tool or system that the finding,
	// detection, or alert originated from. E.g., Endpoint, Secure Email Gateway,
	// etc.
	DetectionSystem string `json:"detection_system,omitempty"`

	// The type of data security tool or system that the finding, detection, or
	// alert originated from.
	DetectionSystemID int `json:"detection_system_id,omitempty"`

	// The associated unique detection event identifier. For example: detection
	// response events include the <b>Detection ID</b> of the original event.
	DetectionUID string `json:"detection_uid,omitempty"`

	// The developer ID on the certificate that signed the file.
	DeveloperUID string `json:"developer_uid,omitempty"`

	// An addressable device, computer system or host.
	Device *objects.Device `json:"device,omitempty"`

	// The object describes details related to the list of devices.
	Devices []*objects.Device `json:"devices,omitempty"`

	// The negotiated protocol dialect.
	Dialect string `json:"dialect,omitempty"`

	// The message digest attribute contains the fixed length message hash
	// representation and the corresponding hashing algorithm information.
	Digest *objects.Fingerprint `json:"digest,omitempty"`

	// The direction of the initiated connection, traffic, or email, normalized to
	// the caption of the direction_id value. In the case of 'Other', it is defined
	// by the event source.
	Direction string `json:"direction,omitempty"`

	// The normalized identifier of the direction of the initiated connection,
	// traffic, or email.
	DirectionID int `json:"direction_id,omitempty"`

	// A collection of Discovery Details objects. See specific usage.
	DiscoveryDetails []*objects.DiscoveryDetails `json:"discovery_details,omitempty"`

	// The dispersion in the NTP protocol is the estimated time error or
	// uncertainty relative to the reference clock in milliseconds.
	Dispersion int64 `json:"dispersion,omitempty"`

	// The display name. See specific usage.
	DisplayName string `json:"display_name,omitempty"`

	// The disposition name, normalized to the caption of the disposition_id value.
	// In the case of 'Other', it is defined by the event source.
	Disposition string `json:"disposition,omitempty"`

	// Describes the outcome or action taken by a security control, such as access
	// control checks, malware detections or various types of policy violations.
	DispositionID int `json:"disposition_id,omitempty"`

	// The DomainKeys Identified Mail (DKIM) status of the email.
	Dkim string `json:"dkim,omitempty"`

	// The DomainKeys Identified Mail (DKIM) signing domain of the email.
	DkimDomain string `json:"dkim_domain,omitempty"`

	// The DomainKeys Identified Mail (DKIM) signature used by the
	// sending/receiving system.
	DkimSignature string `json:"dkim_signature,omitempty"`

	// The Domain-based Message Authentication, Reporting and Conformance (DMARC)
	// status of the email.
	Dmarc string `json:"dmarc,omitempty"`

	// The Domain-based Message Authentication, Reporting and Conformance (DMARC)
	// override action.
	DmarcOverride string `json:"dmarc_override,omitempty"`

	// The Domain-based Message Authentication, Reporting and Conformance (DMARC)
	// policy status.
	DmarcPolicy string `json:"dmarc_policy,omitempty"`

	// The normalized value of dnssec_status_id.
	DnssecStatus string `json:"dnssec_status,omitempty"`

	// Describes the normalized status of DNS Security Extensions (DNSSEC) for a
	// domain.
	DnssecStatusID int `json:"dnssec_status_id,omitempty"`

	// The name of the domain. See specific usage.
	Domain string `json:"domain,omitempty"`

	// The contact information related to a domain registration, e.g., registrant,
	// administrator, abuse, billing, or technical contact.
	DomainContact *objects.DomainContact `json:"domain_contact,omitempty"`

	// An array of Domain Contact objects.
	DomainContacts []*objects.DomainContact `json:"domain_contacts,omitempty"`

	// The drive type, normalized to the caption of the drive_type_id value. In the
	// case of Other, it is defined by the source.
	DriveType string `json:"drive_type,omitempty"`

	// Identifies the type of a disk drive, i.e. fixed, removable, etc.
	DriveTypeID int `json:"drive_type_id,omitempty"`

	// The driver that was loaded/unloaded into the kernel
	Driver *objects.KernelDriver `json:"driver,omitempty"`

	// The network destination endpoint.
	DstEndpoint *objects.NetworkEndpoint `json:"dst_endpoint,omitempty"`

	// This represents the duration of the activity in milliseconds. See specific
	// usage.
	Duration int64 `json:"duration,omitempty"`

	// Represents the duration of the activity in days. See specific usage.
	DurationDays int64 `json:"duration_days,omitempty"`

	// Represents the duration of the activity in hours. See specific usage.
	DurationHours int64 `json:"duration_hours,omitempty"`

	// Represents the duration of the activity in minutes. See specific usage.
	DurationMins int64 `json:"duration_mins,omitempty"`

	// Represents the duration of the activity in months. See specific usage.
	DurationMonths int64 `json:"duration_months,omitempty"`

	// Represents the duration of the activity in seconds. See specific usage.
	DurationSecs int64 `json:"duration_secs,omitempty"`

	// Represents the duration of the activity in weeks. See specific usage.
	DurationWeeks int64 `json:"duration_weeks,omitempty"`

	// Represents the duration of the activity in years. See specific usage.
	DurationYears int64 `json:"duration_years,omitempty"`

	// The list of edge objects that are part of the graph.
	Edges []*objects.Edge `json:"edges,omitempty"`

	// The operating system edition. For example: Professional.
	Edition string `json:"edition,omitempty"`

	// An Embedded Identity Document, is a unique serial number that identifies an
	// eSIM-enabled device.
	Eid string `json:"eid,omitempty"`

	// The email object.
	Email *objects.Email `json:"email,omitempty"`

	// The user's primary email address.
	EmailAddr string `json:"email_addr,omitempty"`

	// A list of additional email addresses for the user.
	EmailAddrs []string `json:"email_addrs,omitempty"`

	// The SPF, DKIM and DMARC attributes of an email.
	EmailAuth *objects.EmailAuth `json:"email_auth,omitempty"`

	// The unique identifier of the email, used to correlate related email alert
	// and activity events.
	EmailUID string `json:"email_uid,omitempty"`

	// The employee identifier assigned to the user by the organization.
	EmployeeUID string `json:"employee_uid,omitempty"`

	// The encryption details of a file or other content. See specific usage.
	EncryptionDetails *objects.EncryptionDetails `json:"encryption_details,omitempty"`

	// The end column number. See specific usage.
	EndColumn int64 `json:"end_column,omitempty"`

	// The line number of the last line of code block identified as vulnerable.
	EndLine int64 `json:"end_line,omitempty"`

	// The end time of a time period. See specific usage.
	EndTime int64 `json:"end_time,omitempty"`

	// The end time of a time period. See specific usage.
	EndTimeDt string `json:"end_time_dt,omitempty"`

	// Contains information about network connection attempts. See specific usage.
	EndpointConnections []*objects.EndpointConnection `json:"endpoint_connections,omitempty"`

	// The additional information from an external data source, which is associated
	// with the event or a finding. For example add location information for the IP
	// address in the DNS answers:</p>[{"name": "answers.ip", "value":
	// "92.24.47.250", "type": "location", "data": {"city": "Socotra", "continent":
	// "Asia", "coordinates": [-25.4153, 17.0743], "country": "YE", "desc":
	// "Yemen"}}]
	Enrichments []*objects.Enrichment `json:"enrichments,omitempty"`

	// The managed entity that is being acted upon.
	Entity *objects.ManagedEntity `json:"entity,omitempty"`

	// The updated managed entity.
	EntityResult *objects.ManagedEntity `json:"entity_result,omitempty"`

	// An array of environment variables.
	EnvironmentVariables []*objects.EnvironmentVariable `json:"environment_variables,omitempty"`

	// The software package epoch. Epoch is a way to define weighted dependencies
	// based on version numbers.
	Epoch int64 `json:"epoch,omitempty"`

	// The Exploit Prediction Scoring System (EPSS) object describes the estimated
	// probability a vulnerability will be exploited. EPSS is a community-driven
	// effort to combine descriptive information about vulnerabilities (CVEs) with
	// evidence of actual exploitation in-the-wild. (https://www.first.org/epss/
	// EPSS).
	Epss *objects.Epss `json:"epss,omitempty"`

	// Error Code
	Error string `json:"error,omitempty"`

	// Error Message
	ErrorMessage string `json:"error_message,omitempty"`

	// The Event ID, Code, or Name that the product uses to primarily identify the
	// event.
	EventCode string `json:"event_code,omitempty"`

	// The unique identifier of an event. See specific usage.
	EventUID string `json:"event_uid,omitempty"`

	// The data the finding exposes to the analyst.
	Evidence interface{} `json:"evidence,omitempty"`

	// A collection of evidence artifacts associated to the activity/activities
	// that triggered a finding. See specific usage.
	Evidences []*objects.Evidences `json:"evidences,omitempty"`

	// The exit code reported by a process when it terminates. The convention is
	// that zero indicates success and any non-zero exit code indicates that some
	// error occurred.
	ExitCode int64 `json:"exit_code,omitempty"`

	// The expiration reason. See specific usage.
	ExpirationReason string `json:"expiration_reason,omitempty"`

	// The expiration time. See specific usage.
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	// The expiration time. See specific usage.
	ExpirationTimeDt string `json:"expiration_time_dt,omitempty"`

	// The time when the exploit was most recently observed.
	ExploitLastSeenTime int64 `json:"exploit_last_seen_time,omitempty"`

	// The time when the exploit was most recently observed.
	ExploitLastSeenTimeDt string `json:"exploit_last_seen_time_dt,omitempty"`

	// The URL of the exploit code or Proof-of-Concept (PoC).
	ExploitRefURL string `json:"exploit_ref_url,omitempty"`

	// The requirement description related to any constraints around exploit
	// execution.
	ExploitRequirement string `json:"exploit_requirement,omitempty"`

	// The categorization or type of Exploit. E.g., Network or Physical.
	ExploitType string `json:"exploit_type,omitempty"`

	// The extension. See specific usage.
	Ext string `json:"ext,omitempty"`

	// The schema extension used to create the event.
	Extension *objects.Extension `json:"extension,omitempty"`

	// The list of TLS extensions.
	ExtensionList []*objects.TlsExtension `json:"extension_list,omitempty"`

	// The schema extensions used to create the event.
	Extensions []*objects.Extension `json:"extensions,omitempty"`

	// A unique identifier assigned by an external system for cross-referencing.
	ExternalUID string `json:"external_uid,omitempty"`

	// The type of authentication factor used in an authentication attempt.
	FactorType string `json:"factor_type,omitempty"`

	// The normalized identifier for the authentication factor.
	FactorTypeID int `json:"factor_type_id,omitempty"`

	// The feature that reported the event.
	Feature *objects.Feature `json:"feature,omitempty"`

	// The file that pertains to the event or object. See specific usage.
	File *objects.File `json:"file,omitempty"`

	// File content differences used for change detection. For example, a common
	// use case is to identify itemized changes within INI or
	// configuration/property setting values.
	FileDiff string `json:"file_diff,omitempty"`

	// The result of the file change. It should contain the new values of the
	// changed attributes.
	FileResult *objects.File `json:"file_result,omitempty"`

	// The files that are part of the event or object.
	Files []*objects.File `json:"files,omitempty"`

	// The Finding object provides details about a finding/detection generated by a
	// security tool.
	Finding *objects.Finding `json:"finding,omitempty"`

	// Describes the supporting information about a generated finding.
	FindingInfo *objects.FindingInfo `json:"finding_info,omitempty"`

	// A list of finding_info objects associated to an incident.
	FindingInfoList []*objects.FindingInfo `json:"finding_info_list,omitempty"`

	// The digital fingerprint associated with an object.
	Fingerprint *objects.Fingerprint `json:"fingerprint,omitempty"`

	// An array of digital fingerprint objects.
	Fingerprints []*objects.Fingerprint `json:"fingerprints,omitempty"`

	// The firewall rule that triggered the event.
	FirewallRule *objects.FirewallRule `json:"firewall_rule,omitempty"`

	// The initial detection time of the activity or object. See specific usage
	FirstSeenTime int64 `json:"first_seen_time,omitempty"`

	// The initial detection time of the activity or object. See specific usage
	FirstSeenTimeDt string `json:"first_seen_time_dt,omitempty"`

	// Indicates if a fix is available for the reported vulnerability.
	FixAvailable bool `json:"fix_available,omitempty"`

	// The fix coverage, normalized to the caption of the fix_coverage_id value.
	// See specific usage.
	FixCoverage string `json:"fix_coverage,omitempty"`

	// The normalized identifier for fix coverage. See specific usage.
	FixCoverageID int `json:"fix_coverage_id,omitempty"`

	// The software package version in which a reported vulnerability was
	// patched/fixed.
	FixedInVersion string `json:"fixed_in_version,omitempty"`

	// The Connection Flag History summarizes events in a network connection. For
	// example flags ShAD representing SYN, SYN/ACK, ACK and Data exchange.
	FlagHistory string `json:"flag_history,omitempty"`

	// The list of normalized identifiers of the communication flag IDs. See
	// specific usage.
	FlagIDs []int `json:"flag_ids,omitempty"`

	// The list of communication flags, normalized to the captions of the flag_ids
	// values. See specific usage.
	Flags []string `json:"flags,omitempty"`

	// The folder that pertains to the event.
	Folder *objects.File `json:"folder,omitempty"`

	// The user's forwarding email address.
	ForwardAddr string `json:"forward_addr,omitempty"`

	// The machine-readable email header From value, as defined by RFC 5322. For
	// example example.user@usersdomain.com.
	From string `json:"from,omitempty"`

	// The machine-readable email header From values. This array should contain the
	// value in from. For example example.user@usersdomain.com.
	FromList []string `json:"from_list,omitempty"`

	// The human-readable email header From Mailbox value. For example 'Example
	// User &lt;example.user@usersdomain.com&gt;'.
	FromMailbox string `json:"from_mailbox,omitempty"`

	// The human-readable email header From Mailbox values. This array should
	// contain the value in from_mailbox. For example 'Example User
	// &lt;example.user@usersdomain.com&gt;'.
	FromMailboxes []string `json:"from_mailboxes,omitempty"`

	// The full name. See specific usage.
	FullName string `json:"full_name,omitempty"`

	// The number of function keys on client keyboard.
	FunctionKeys int64 `json:"function_keys,omitempty"`

	// The entry-point function of the module. The system calls the entry-point
	// function whenever a process or thread loads or unloads the module.
	FunctionName string `json:"function_name,omitempty"`

	// The aircraft distance above or below the ellipsoid as measured along a line
	// that passes through the aircraft and is normal to the surface of the WGS-84
	// ellipsoid. This value is provided in meters and must have a minimum
	// resolution of 1 m. Special Values: Invalid, No Value, or Unknown: -1000 m.
	GeodeticAltitude string `json:"geodetic_altitude,omitempty"`

	// Provides quality/containment on geodetic altitude. This is based on ADS-B
	// Geodetic Vertical Accuracy (GVA). Measured in meters.
	GeodeticVerticalAccuracy string `json:"geodetic_vertical_accuracy,omitempty"`

	// <p>Geohash of the geo-coordinates (latitude and
	// longitude).</p>https://en.wikipedia.org/wiki/Geohash Geohashing is a
	// geocoding system used to encode geographic coordinates in decimal degrees,
	// to a single string.
	Geohash string `json:"geohash,omitempty"`

	// The given or first name of the user.
	GivenName string `json:"given_name,omitempty"`

	// The Privileges that were granted to the user via an IAM policy or otherwise.
	// See specific usage.
	GrantedPrivileges []string `json:"granted_privileges,omitempty"`

	// A graph data structure representation with nodes and edges.
	Graph *objects.Graph `json:"graph,omitempty"`

	// The group object associated with an entity such as user, policy, or rule.
	Group *objects.Group `json:"group,omitempty"`

	// The name of the group that the resource belongs to.
	GroupName string `json:"group_name,omitempty"`

	// The groups to which an entity belongs. See specific usage.
	Groups []*objects.Group `json:"groups,omitempty"`

	// The amount of total time for the TLS handshake to complete after the TCP
	// connection is established, including client-side delays, in milliseconds.
	HandshakeDur int64 `json:"handshake_dur,omitempty"`

	// The user has a multi-factor or secondary-factor device assigned.
	HasMfa bool `json:"has_mfa,omitempty"`

	// The hash attribute is the value of a digital fingerprint including
	// information about its algorithm.
	Hash *objects.Fingerprint `json:"hash,omitempty"`

	// An array of hash attributes.
	Hashes []*objects.Fingerprint `json:"hashes,omitempty"`

	// The timestamp when the user was or will be hired by the organization.
	HireTime int64 `json:"hire_time,omitempty"`

	// The timestamp when the user was or will be hired by the organization.
	HireTimeDt string `json:"hire_time_dt,omitempty"`

	// Provides quality/containment on horizontal position. This is based on ADS-B
	// NACp. Measured in meters.
	HorizontalAccuracy string `json:"horizontal_accuracy,omitempty"`

	// The hostname of an endpoint or a device.
	Hostname string `json:"hostname,omitempty"`

	// The cookies object describes details about HTTP cookies
	HttpCookies []*objects.HttpCookie `json:"http_cookies,omitempty"`

	// Additional HTTP headers of an HTTP request or response.
	HttpHeaders []*objects.HttpHeader `json:"http_headers,omitempty"`

	// The HTTP request method indicates the desired action to be performed for a
	// given resource. Expected values: <ul> <li>TRACE</li> <li>CONNECT</li>
	// <li>OPTIONS</li> <li>HEAD</li> <li>DELETE</li> <li>POST</li> <li>PUT</li>
	// <li>GET</li></ul>
	HttpMethod string `json:"http_method,omitempty"`

	// A cookie attribute to make it inaccessible via JavaScript
	HttpOnly bool `json:"http_only,omitempty"`

	// The HTTP Request Object documents attributes of a request made to a web
	// server.
	HttpRequest *objects.HttpRequest `json:"http_request,omitempty"`

	// The HTTP Response from a web server to a requester.
	HttpResponse *objects.HttpResponse `json:"http_response,omitempty"`

	// The Hypertext Transfer Protocol (HTTP)
	// https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
	// status code returned to the client.
	HttpStatus int64 `json:"http_status,omitempty"`

	// The endpoint hardware information.
	HwInfo *objects.DeviceHwInfo `json:"hw_info,omitempty"`

	// The name of the hypervisor running on the device. For example, Xen, VMware,
	// Hyper-V, VirtualBox, etc.
	Hypervisor string `json:"hypervisor,omitempty"`

	// The Integrated Circuit Card Identification of a mobile device. Typically it
	// is a unique 18 to 22 digit number that identifies a SIM card.
	Iccid string `json:"iccid,omitempty"`

	// The client identifier cookie during client/server exchange.
	IdentifierCookie string `json:"identifier_cookie,omitempty"`

	// Describes usage activity and other metrics of an Identity i.e. AWS IAM User,
	// GCP IAM Principal, etc.
	IdentityActivityMetrics *objects.IdentityActivityMetrics `json:"identity_activity_metrics,omitempty"`

	// Duration (in minutes) of allowed inactivity before a timeout See specific
	// usage.
	IdleTimeout int64 `json:"idle_timeout,omitempty"`

	// This object describes details about the Identity Provider used.
	Idp *objects.Idp `json:"idp,omitempty"`

	// The image used as a template to run a container or virtual machine.
	Image *objects.Image `json:"image,omitempty"`

	// The Input Method Editor (IME) file name.
	Ime string `json:"ime,omitempty"`

	// The International Mobile Equipment Identity that is associated with the
	// device.
	Imei string `json:"imei,omitempty"`

	// The International Mobile Equipment Identity values that are associated with
	// the device.
	ImeiList []string `json:"imei_list,omitempty"`

	// The impact , normalized to the caption of the impact_id value. In the case
	// of 'Other', it is defined by the event source.
	Impact string `json:"impact,omitempty"`

	// The normalized impact of the incident or finding. Per NIST, this is the
	// magnitude of harm that can be expected to result from the consequences of
	// unauthorized disclosure, modification, destruction, or loss of information
	// or information system availability.
	ImpactID int `json:"impact_id,omitempty"`

	// The impact as an integer value of the finding, valid range 0-100.
	ImpactScore int64 `json:"impact_score,omitempty"`

	// The process injection method, normalized to the caption of the
	// injection_type_id value. In the case of 'Other', it is defined by the event
	// source.
	InjectionType string `json:"injection_type,omitempty"`

	// The normalized identifier of the process injection method.
	InjectionTypeID int `json:"injection_type_id,omitempty"`

	// The install state, normalized to the caption of install_state_id. In the
	// case of 'Other', it is defined by the event source.
	InstallState string `json:"install_state,omitempty"`

	// The normalized state of the install.
	InstallStateID int `json:"install_state_id,omitempty"`

	// The unique identifier of a VM instance.
	InstanceUID string `json:"instance_uid,omitempty"`

	// The process integrity level, normalized to the caption of the integrity_id
	// value. In the case of 'Other', it is defined by the event source (Windows
	// only).
	Integrity string `json:"integrity,omitempty"`

	// The normalized identifier of the process integrity level (Windows only).
	IntegrityID int `json:"integrity_id,omitempty"`

	// The name of the network interface (e.g. eth2).
	InterfaceName string `json:"interface_name,omitempty"`

	// The unique identifier of the network interface.
	InterfaceUID string `json:"interface_uid,omitempty"`

	// The intermediate IP Addresses. For example, the IP addresses in the HTTP
	// X-Forwarded-For header.
	IntermediateIPs []string `json:"intermediate_ips,omitempty"`

	// The name by which a resource identifies itself internally. See specific
	// usage.
	InternalName string `json:"internal_name,omitempty"`

	// A grouping of adversarial behaviors and resources believed to be associated
	// with specific threat actors or campaigns. Intrusion sets often encompass
	// multiple campaigns and are used to organize related activities under a
	// common label.
	IntrusionSets []string `json:"intrusion_sets,omitempty"`

	// The name of the service that invoked the activity as described in the event.
	InvokedBy string `json:"invoked_by,omitempty"`

	// The IP address, in either IPv4 or IPv6 format.
	Ip string `json:"ip,omitempty"`

	// Indicates that the event is considered to be an alertable signal.
	IsAlert bool `json:"is_alert,omitempty"`

	// A determination if a policy, rule, or enforcement action was applied.
	IsApplied bool `json:"is_applied,omitempty"`

	// Indicates whether the device or resource has a backup enabled, such as an
	// automated snapshot or a cloud backup. For example, this is indicated by the
	// cloudBackupEnabled value within JAMF Pro mobile devices or the registration
	// of an AWS ARN with the AWS Backup service.
	IsBackedUp bool `json:"is_backed_up,omitempty"`

	// Indicates whether the credentials were passed in clear text.<p><b>Note:</b>
	// True if the credentials were passed in a clear text protocol such as FTP or
	// TELNET, or if Windows detected that a user's logon password was passed to
	// the authentication package in clear text.</p>
	IsCleartext bool `json:"is_cleartext,omitempty"`

	// The event occurred on a compliant device.
	IsCompliant bool `json:"is_compliant,omitempty"`

	// The indication of whether the value is from a default value name. For
	// example, the value name could be missing.
	IsDefault bool `json:"is_default,omitempty"`

	// Indicates if the entity was deleted. See specific usage.
	IsDeleted bool `json:"is_deleted,omitempty"`

	// Indicates if the entity has directionality. See specific usage.
	IsDirected bool `json:"is_directed,omitempty"`

	// Indicates if the entity was encrypted. See specific usage.
	IsEncrypted bool `json:"is_encrypted,omitempty"`

	// Indicates if an exploit or a PoC (proof-of-concept) is available for the
	// reported vulnerability.
	IsExploitAvailable bool `json:"is_exploit_available,omitempty"`

	// Indicates if a fix is available for the reported vulnerability.
	IsFixAvailable bool `json:"is_fix_available,omitempty"`

	// Indicates whether group provisioning is automated (e.g., for a SCIM
	// resource). See specific usage.
	IsGroupProvisioningEnabled bool `json:"is_group_provisioning_enabled,omitempty"`

	// Whether the authentication factor is an HMAC-based One-time Password (HOTP).
	IsHotp bool `json:"is_hotp,omitempty"`

	// This attribute prevents the cookie from being accessed via JavaScript.
	IsHTTPOnly bool `json:"is_http_only,omitempty"`

	// The event occurred on a managed device.
	IsManaged bool `json:"is_managed,omitempty"`

	// Indicates whether Multi Factor Authentication was used during
	// authentication.
	IsMfa bool `json:"is_mfa,omitempty"`

	// Indicates whether the device has an active mobile account. For example, this
	// is indicated by the itunesStoreAccountActive value within JAMF Pro mobile
	// devices.
	IsMobileAccountActive bool `json:"is_mobile_account_active,omitempty"`

	// Indicates logon is from a device not seen before or a first time account
	// logon.
	IsNewLogon bool `json:"is_new_logon,omitempty"`

	// The indication of whether the location is on premises.
	IsOnPremises bool `json:"is_on_premises,omitempty"`

	// The event occurred on a personal device.
	IsPersonal bool `json:"is_personal,omitempty"`

	// Determination of the public accessibility. See specific usage.
	IsPublic bool `json:"is_public,omitempty"`

	// The indication of whether the email has been read.
	IsRead bool `json:"is_read,omitempty"`

	// Indicates that an object cannot be modified. See specific usage
	IsReadonly bool `json:"is_readonly,omitempty"`

	// The indication of whether the session is remote.
	IsRemote bool `json:"is_remote,omitempty"`

	// The indication of whether something is renewable. See specific usage.
	IsRenewable bool `json:"is_renewable,omitempty"`

	// The indication of whether the event or object represents a renewal. See
	// specific usage.
	IsRenewal bool `json:"is_renewal,omitempty"`

	// The cookie attribute indicates that cookies are sent to the server only when
	// the request is encrypted using the HTTPS protocol.
	IsSecure bool `json:"is_secure,omitempty"`

	// Denotes whether a digital certificate is self-signed or signed by a known
	// certificate authority (CA).
	IsSelfSigned bool `json:"is_self_signed,omitempty"`

	// The event occurred on a shared device.
	IsShared bool `json:"is_shared,omitempty"`

	// true denotes that src_endpoint and dst_endpoint correctly identify the
	// initiator and responder respectively. false denotes that the event source
	// has arbitrarily assigned one peer to src_endpoint and the other to
	// dst_endpoint, in other words that initiator and responder are not being
	// asserted. This can occur, for example, when the event source is a network
	// appliance that has not observed the initiation of a given connection. In the
	// absence of this attribute, interpretation of the initiator and responder is
	// implementation-specific.
	IsSrcDstAssignmentKnown bool `json:"is_src_dst_assignment_known,omitempty"`

	// The vendor patch has been replaced by another.
	IsSuperseded bool `json:"is_superseded,omitempty"`

	// The event occurred on a supervised device. Devices that are supervised are
	// typically mobile devices managed by a Mobile Device Management solution and
	// are restricted from specific behaviors such as Apple AirDrop.
	IsSupervised bool `json:"is_supervised,omitempty"`

	// A determination based on analytics as to whether a potential breach was
	// found.
	IsSuspectedBreach bool `json:"is_suspected_breach,omitempty"`

	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`

	// Whether the authentication factor is a Time-based One-time Password (TOTP).
	IsTotp bool `json:"is_totp,omitempty"`

	// Indicates that an attribute has been truncated. See specific usage.
	IsTruncated bool `json:"is_truncated,omitempty"`

	// The event occurred on a trusted device.
	IsTrusted bool `json:"is_trusted,omitempty"`

	// Indicates whether user provisioning is automated (e.g., for a SCIM
	// resource). See specific usage.
	IsUserProvisioningEnabled bool `json:"is_user_provisioning_enabled,omitempty"`

	// The indication of whether the session is a VPN session.
	IsVpn bool `json:"is_vpn,omitempty"`

	// The name of the Internet Service Provider (ISP).
	Isp string `json:"isp,omitempty"`

	// The organization name of the Internet Service Provider (ISP). This
	// represents the parent organization or company that owns/operates the ISP.
	// For example, Comcast Corporation would be the ISP org for Xfinity internet
	// service. This attribute helps identify the ultimate provider when ISPs
	// operate under different brand names.
	IspOrg string `json:"isp_org,omitempty"`

	// The identifier of the issuer. See specific usage.
	Issuer string `json:"issuer,omitempty"`

	// The MD5 hash of a JA3 string.
	Ja3Hash *objects.Fingerprint `json:"ja3_hash,omitempty"`

	// The MD5 hash of a JA3S string.
	Ja3sHash *objects.Fingerprint `json:"ja3s_hash,omitempty"`

	// A list of the JA4+ network fingerprints.
	Ja4FingerprintList []*objects.Ja4Fingerprint `json:"ja4_fingerprint_list,omitempty"`

	// The job object that pertains to the event.
	Job *objects.Job `json:"job,omitempty"`

	// The user's job title.
	JobTitle string `json:"job_title,omitempty"`

	// The JSON path of the attribute. See specific usage.
	JsonPath string `json:"json_path,omitempty"`

	// A list of KB articles or patches related to an endpoint. A KB Article
	// contains metadata that describes the patch or an update.
	KbArticleList []*objects.KbArticle `json:"kb_article_list,omitempty"`

	// The KB article/s related to the entity. A KB Article contains metadata that
	// describes the patch or an update.
	KbArticles []string `json:"kb_articles,omitempty"`

	// A bitmask, either in hexadecimal or decimal form, which encodes various
	// attributes or permissions associated with a Kerberos ticket. These flags
	// delineate specific characteristics of the ticket, such as its renewability
	// or forwardability.
	KerberosFlags string `json:"kerberos_flags,omitempty"`

	// The kernel resource object that pertains to the event.
	Kernel *objects.Kernel `json:"kernel,omitempty"`

	// The kernel release of the operating system. On Unix-based systems, this is
	// determined from the uname -r command output, for example
	// "5.15.0-122-generic".
	KernelRelease string `json:"kernel_release,omitempty"`

	// The length of the encryption key.
	KeyLength int64 `json:"key_length,omitempty"`

	// The unique identifier of the key. See specific usage.
	KeyUID string `json:"key_uid,omitempty"`

	// The keyboard detailed information.
	KeyboardInfo *objects.KeyboardInfo `json:"keyboard_info,omitempty"`

	// The keyboard locale identifier name (e.g., en-US).
	KeyboardLayout string `json:"keyboard_layout,omitempty"`

	// The keyboard numeric code.
	KeyboardSubtype int64 `json:"keyboard_subtype,omitempty"`

	// The keyboard type (e.g., xt, ico).
	KeyboardType string `json:"keyboard_type,omitempty"`

	// The
	// https://www.lockheedmartin.com/en-us/capabilities/cyber/cyber-kill-chain.html
	// Cyber Kill Chain® provides a detailed description of each phase and its
	// associated activities within the broader context of a cyber attack.
	KillChain []*objects.KillChainPhase `json:"kill_chain,omitempty"`

	// The list of labels attached to an entity. See specific usage.
	Labels []string `json:"labels,omitempty"`

	// The two letter lower case language codes, as defined by
	// https://en.wikipedia.org/wiki/ISO_639-1 ISO 639-1. For example: en
	// (English), de (German), or fr (French).
	Lang string `json:"lang,omitempty"`

	// The timestamp when this identity last successfully authenticated to any
	// system or service. See specific usage.
	LastAuthenticationTime int64 `json:"last_authentication_time,omitempty"`

	// The timestamp when this identity last successfully authenticated to any
	// system or service. See specific usage.
	LastAuthenticationTimeDt string `json:"last_authentication_time_dt,omitempty"`

	// The last time when the user logged in.
	LastLoginTime int64 `json:"last_login_time,omitempty"`

	// The last time when the user logged in.
	LastLoginTimeDt string `json:"last_login_time_dt,omitempty"`

	// The last run time of application or service. See specific usage.
	LastRunTime int64 `json:"last_run_time,omitempty"`

	// The last run time of application or service. See specific usage.
	LastRunTimeDt string `json:"last_run_time_dt,omitempty"`

	// The most recent detection time of the activity or object. See specific
	// usage.
	LastSeenTime int64 `json:"last_seen_time,omitempty"`

	// The most recent detection time of the activity or object. See specific
	// usage.
	LastSeenTimeDt string `json:"last_seen_time_dt,omitempty"`

	// The most recent usage time of an entity. See specific usage.
	LastUsedTime int64 `json:"last_used_time,omitempty"`

	// The most recent usage time of an entity. See specific usage.
	LastUsedTimeDt string `json:"last_used_time_dt,omitempty"`

	// The geographical Latitude coordinate represented in Decimal Degrees (DD).
	// For example: 42.361145.
	Lat float64 `json:"lat,omitempty"`

	// The HTTP response latency measured in milliseconds.
	Latency int64 `json:"latency,omitempty"`

	// The LDAP and X.500 commonName attribute, typically the full name of the
	// person. For example, John Doe.
	LdapCn string `json:"ldap_cn,omitempty"`

	// The X.500 Distinguished Name (DN) is a structured string that uniquely
	// identifies an entry, such as a user, in an X.500 directory service For
	// example, cn=John Doe,ou=People,dc=example,dc=com.
	LdapDn string `json:"ldap_dn,omitempty"`

	// The additional LDAP attributes that describe a person.
	LdapPerson *objects.LdapPerson `json:"ldap_person,omitempty"`

	// This represents the length of the DHCP lease in seconds. This is present in
	// DHCP Ack events.
	LeaseDur int64 `json:"lease_dur,omitempty"`

	// The timestamp when the user left or will be leaving the organization.
	LeaveTime int64 `json:"leave_time,omitempty"`

	// The timestamp when the user left or will be leaving the organization.
	LeaveTimeDt string `json:"leave_time_dt,omitempty"`

	// The HTTP response length, in number of bytes.
	Length int64 `json:"length,omitempty"`

	// The name or identifier of the license applied on package or software. See
	// https://spdx.org/licenses/ SPDX License List.
	License string `json:"license,omitempty"`

	// The URL pointing to the license applied on package or software. This is
	// typically a LICENSE.md file within a repository.
	LicenseURL string `json:"license_url,omitempty"`

	// The lineage of the process, represented by a list of paths for each ancestor
	// process. For example: ['/usr/sbin/sshd', '/usr/bin/bash',
	// '/usr/bin/whoami'].
	Lineage []interface{} `json:"lineage,omitempty"`

	// The audit user assigned at login by the audit subsystem.
	LinuxAuid int64 `json:"linux/auid,omitempty"`

	// The effective group under which this process is running.
	LinuxEgid int64 `json:"linux/egid,omitempty"`

	// The effective user under which this process is running.
	LinuxEuid int64 `json:"linux/euid,omitempty"`

	// The Load Balancer object contains information related to the device that is
	// distributing incoming traffic to specified destinations.
	LoadBalancer *objects.LoadBalancer `json:"load_balancer,omitempty"`

	// The load type, normalized to the caption of the load_type_id value. In the
	// case of 'Other', it is defined by the event source.
	LoadType string `json:"load_type,omitempty"`

	// The normalized identifier of the load type. See specific usage.
	LoadTypeID int `json:"load_type_id,omitempty"`

	// The list of loaded module names.
	LoadedModules []string `json:"loaded_modules,omitempty"`

	// The detailed geographical location usually associated with an IP address.
	Location *objects.Location `json:"location,omitempty"`

	// A list of detailed geographical locations.
	Locations []*objects.Location `json:"locations,omitempty"`

	// The audit level at which an event was generated.
	LogLevel string `json:"log_level,omitempty"`

	// The event log name. For example, syslog file name or Windows logging
	// subsystem: Security.
	LogName string `json:"log_name,omitempty"`

	// The logging provider or logging service that logged the event. For example,
	// Microsoft-Windows-Security-Auditing.
	LogProvider string `json:"log_provider,omitempty"`

	// The log type, normalized to the caption of the log_type_id value. In the
	// case of 'Other', it is defined by the event source.
	LogType string `json:"log_type,omitempty"`

	// The normalized log type identifier.
	LogTypeID int `json:"log_type_id,omitempty"`

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
	Loggers []*objects.Logger `json:"loggers,omitempty"`

	// URL for initiating a login request. See specific usage.
	LoginEndpoint string `json:"login_endpoint,omitempty"`

	// The trusted process that validated the authentication credentials.
	LogonProcess *objects.Process `json:"logon_process,omitempty"`

	// The logon type, normalized to the caption of the logon_type_id value. In the
	// case of 'Other', it is defined by the event source.
	LogonType string `json:"logon_type,omitempty"`

	// The normalized logon type identifier.
	LogonTypeID int `json:"logon_type_id,omitempty"`

	// URL for initiating a logout request. See specific usage.
	LogoutEndpoint string `json:"logout_endpoint,omitempty"`

	// The geographical Longitude coordinate represented in Decimal Degrees (DD).
	// For example: -71.057083.
	Long float64 `json:"long,omitempty"`

	// The Media Access Control (MAC) address that is associated with the network
	// interface.
	Mac string `json:"mac,omitempty"`

	// A list of Malware objects, describing details about the identified malware.
	Malware []*objects.Malware `json:"malware,omitempty"`

	// Describes details about the scan job that identified malware on the target
	// system.
	MalwareScanInfo *objects.MalwareScanInfo `json:"malware_scan_info,omitempty"`

	// The user's manager. This helps in understanding an org hierarchy. This
	// should only ever be populated once in an event. I.e. there should not be a
	// manager's manager in an event.
	Manager *objects.User `json:"manager,omitempty"`

	// The data in a request that rule matched. For example: '["10","and","1"]'.
	MatchDetails []string `json:"match_details,omitempty"`

	// The location of the matched data in the source which resulted in the
	// triggered firewall rule. For example: HEADER.
	MatchLocation string `json:"match_location,omitempty"`

	// Determines if an assessment, control, policy, or otherwise meets its
	// assessment criteria. See specific usage.
	MeetsCriteria bool `json:"meets_criteria,omitempty"`

	// The Mobile Equipment Identifier. It's a unique number that identifies a Code
	// Division Multiple Access (CDMA) mobile device.
	Meid string `json:"meid,omitempty"`

	// The description of the event/finding, as defined by the source.
	Message string `json:"message,omitempty"`

	// The identifier that tracks a message that travels through multiple points of
	// a messaging service.
	MessageTraceUID string `json:"message_trace_uid,omitempty"`

	// The email header Message-ID value, as defined by RFC 5322.
	MessageUID string `json:"message_uid,omitempty"`

	// The metadata associated with the event or a finding.
	Metadata *objects.Metadata `json:"metadata,omitempty"`

	// URL where metadata about a configuration or resource is available (e.g., for
	// SAML configurations). See specific usage.
	MetadataEndpoint string `json:"metadata_endpoint,omitempty"`

	// The general purpose metrics associated with the event. See specific usage.
	Metrics []*objects.Metric `json:"metrics,omitempty"`

	// The Multipurpose Internet Mail Extensions (MIME) type of the file, if
	// applicable.
	MimeType string `json:"mime_type,omitempty"`

	// The Mitigation object describes the MITRE ATT&CK® or ATLAS™ Mitigation ID
	// and/or name that is associated to an attack.
	Mitigation *objects.Mitigation `json:"mitigation,omitempty"`

	// The model name of an entity. See specific usage.
	Model string `json:"model,omitempty"`

	// The time when the object was last modified. See specific usage.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the object was last modified. See specific usage.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The user that last modified the object associated with the event. See
	// specific usage.
	Modifier *objects.User `json:"modifier,omitempty"`

	// The module that pertains to the event.
	Module *objects.Module `json:"module,omitempty"`

	// The name of the entity. See specific usage.
	Name string `json:"name,omitempty"`

	// A collection of name servers related to a domain registration or other
	// record.
	NameServers []string `json:"name_servers,omitempty"`

	// The namespace is useful in merger or acquisition situations. For example,
	// when similar entities exist that you need to keep separate.
	Namespace string `json:"namespace,omitempty"`

	// If running under a process namespace (such as in a container), the process
	// identifier within that process namespace.
	NamespacePid int64 `json:"namespace_pid,omitempty"`

	// The network driver used by the container. For example, bridge, overlay,
	// host, none, etc.
	NetworkDriver string `json:"network_driver,omitempty"`

	// The Network Endpoint object describes characteristics of a network endpoint.
	// See specific usage.
	NetworkEndpoint *objects.NetworkEndpoint `json:"network_endpoint,omitempty"`

	// The physical or virtual network interfaces that are associated with the
	// device, one for each unique MAC address/IP address/hostname/name
	// combination.<p><b>Note:</b> The first element of the array is the network
	// information that pertains to the event.</p>
	NetworkInterfaces []*objects.NetworkInterface `json:"network_interfaces,omitempty"`

	// The next run time. See specific usage.
	NextRunTime int64 `json:"next_run_time,omitempty"`

	// The next run time. See specific usage.
	NextRunTimeDt string `json:"next_run_time_dt,omitempty"`

	// The NIST Cybersecurity Framework recommendations for managing the
	// cybersecurity risk.
	Nist []string `json:"nist,omitempty"`

	// The list of node objects that are part of the graph.
	Nodes []*objects.Node `json:"nodes,omitempty"`

	// The number of detections.
	NumDetections int64 `json:"num_detections,omitempty"`

	// The number of files scanned.
	NumFiles int64 `json:"num_files,omitempty"`

	// The number of folders scanned.
	NumFolders int64 `json:"num_folders,omitempty"`

	// The number of infected entities. See specific usage.
	NumInfected int64 `json:"num_infected,omitempty"`

	// The number of network items scanned.
	NumNetworkItems int64 `json:"num_network_items,omitempty"`

	// The number of processes scanned.
	NumProcesses int64 `json:"num_processes,omitempty"`

	// The number of registry items scanned.
	NumRegistryItems int64 `json:"num_registry_items,omitempty"`

	// The number of items that were resolved.
	NumResolutions int64 `json:"num_resolutions,omitempty"`

	// The number of skipped items.
	NumSkippedItems int64 `json:"num_skipped_items,omitempty"`

	// The number of trusted items.
	NumTrustedItems int64 `json:"num_trusted_items,omitempty"`

	// The number of times the policy or rule was violated.
	NumViolations int64 `json:"num_violations,omitempty"`

	// The number of volumes in the storage device. See specific usage.
	NumVolumes int64 `json:"num_volumes,omitempty"`

	// The number of the entity. See specific usage.
	Number int64 `json:"number,omitempty"`

	// The observables associated with the event or a finding.
	Observables []*objects.Observable `json:"observables,omitempty"`

	// The name of the parameter being analyzed or monitored. This generally
	// identifies the specific characteristic or property being measured. See
	// specific usage.
	ObservationParameter string `json:"observation_parameter,omitempty"`

	// The classification or category of the observation, indicating what kind of
	// measurement or finding it represents. See specific usage.
	ObservationType string `json:"observation_type,omitempty"`

	// A list of individual observations or measurements collected during analysis.
	// Each observation captures a specific data point. See specific usage.
	Observations []*objects.Observation `json:"observations,omitempty"`

	// A detected pattern or trend identified in the analyzed data, describing
	// recurring activities or characteristics.
	ObservedPattern string `json:"observed_pattern,omitempty"`

	// Details about where in the target entity the specified information was
	// discovered. See specific usage.
	OccurrenceDetails *objects.OccurrenceDetails `json:"occurrence_details,omitempty"`

	// A list of occurrence_details objects, each describing where in the target
	// entity the specified information was discovered. See specific usage.
	Occurrences []*objects.OccurrenceDetails `json:"occurrences,omitempty"`

	// The primary office location associated with the user. This could be any
	// string and isn't a specific address. For example, South East Virtual.
	OfficeLocation string `json:"office_location,omitempty"`

	// The DNS opcode specifies the type of the query message.
	Opcode string `json:"opcode,omitempty"`

	// The DNS opcode ID specifies the normalized query message type as defined in
	// https://www.rfc-editor.org/rfc/rfc5395.html RFC-5395.
	OpcodeID int `json:"opcode_id,omitempty"`

	// The Windows options needed to open a registry key.
	OpenMask int64 `json:"open_mask,omitempty"`

	// The list of open ports on a network interface, including port numbers and
	// associated protocol information.
	OpenPorts []*objects.PortInfo `json:"open_ports,omitempty"`

	// The file open type.
	OpenType string `json:"open_type,omitempty"`

	// Verb/Operation associated with the request
	Operation string `json:"operation,omitempty"`

	// An operation number used to identify a specific remote procedure call (RPC)
	// method or a method in an interface.
	Opnum int64 `json:"opnum,omitempty"`

	// The orchestrator managing the container, such as ECS, EKS, K8s, or
	// OpenShift.
	Orchestrator string `json:"orchestrator,omitempty"`

	// Organization and org unit relevant to the event or object.
	Org *objects.Organization `json:"org,omitempty"`

	// The original event time as reported by the event source. For example, the
	// time in the original format from system event log such as Syslog on
	// Unix/Linux and the System event file on Windows. Omit if event is generated
	// instead of collected via logs.
	OriginalTime string `json:"original_time,omitempty"`

	// The endpoint operating system.
	Os *objects.Os `json:"os,omitempty"`

	// The operating system assigned Machine ID. In Windows, this is the value
	// stored at the registry path:
	// HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Cryptography\MachineGuid. In Linux,
	// this is stored in the file: /etc/machine-id.
	OsMAChineUUID string `json:"os_machine_uuid,omitempty"`

	// The OSINT (Open Source Intelligence) object contains details related to an
	// indicator such as the indicator itself, related indicators, geolocation,
	// registrar information, subdomains, analyst commentary, and other contextual
	// information. This information can be used to further enrich a detection or
	// finding by providing decisioning support to other analysts and engineers.
	Osint []*objects.Osint `json:"osint,omitempty"`

	// The name of the organizational unit, within an organization. For example,
	// Finance, IT, R&D
	OuName string `json:"ou_name,omitempty"`

	// The alternate identifier for an entity's unique identifier. For example, its
	// Active Directory OU DN or AWS OU ID.
	OuUID string `json:"ou_uid,omitempty"`

	// The overall score as reported by the event source. See specific usage.
	OverallScore float64 `json:"overall_score,omitempty"`

	// The user that owns the file/object.
	Owner *objects.User `json:"owner,omitempty"`

	// The Software Package object describes details about a software package.
	Package *objects.Package `json:"package,omitempty"`

	// The software packager manager utilized to manage a package on a system, e.g.
	// npm, yum, dpkg etc.
	PackageManager string `json:"package_manager,omitempty"`

	// The URL of the package or library at the package manager, or the specific
	// URL or URI of an internal package manager link such as AWS CodeArtifact or
	// Artifactory.
	PackageManagerURL string `json:"package_manager_url,omitempty"`

	// List of vulnerable packages as identified by the security product
	Packages []*objects.Package `json:"packages,omitempty"`

	// The packet identifier assigned by the protocol.
	PacketUID int64 `json:"packet_uid,omitempty"`

	// The total number of packets (in and out).
	Packets int64 `json:"packets,omitempty"`

	// The number of packets sent from the destination to the source.
	PacketsIn int64 `json:"packets_in,omitempty"`

	// The number of packets sent from the source to the destination.
	PacketsOut int64 `json:"packets_out,omitempty"`

	// The page number of the document. See specific usage.
	PageNumber int64 `json:"page_number,omitempty"`

	// The parent folder in which the file resides. For example:
	// c:\windows\system32
	ParentFolder string `json:"parent_folder,omitempty"`

	// The parent process of this process object. It is recommended to only
	// populate this field for the top-level process object, to prevent deep
	// nesting. Additional ancestry information can be supplied in the ancestry
	// attribute.
	ParentProcess *objects.Process `json:"parent_process,omitempty"`

	// The unique identifier of an object's parent object. See specific usage.
	ParentUID string `json:"parent_uid,omitempty"`

	// The time when a user's password was last used. See specific usage.
	PasswordLastUsedTime int64 `json:"password_last_used_time,omitempty"`

	// The time when a user's password was last used. See specific usage.
	PasswordLastUsedTimeDt string `json:"password_last_used_time_dt,omitempty"`

	// The path that pertains to the event or object. See specific usage.
	Path string `json:"path,omitempty"`

	// A text, binary, file name, or datastore that matched against a detection
	// rule.
	PatternMatch string `json:"pattern_match,omitempty"`

	// The EPSS score's percentile representing relative importance and ranking of
	// the score in the larger EPSS dataset.
	Percentile float64 `json:"percentile,omitempty"`

	// The peripheral device that triggered the event.
	PeripheralDevice *objects.PeripheralDevice `json:"peripheral_device,omitempty"`

	// The IAM permission related to an event
	Permission string `json:"permission,omitempty"`

	// Describes analysis results of permissions, policies directly associated with
	// an identity (user, role, or service account). This evaluates what
	// permissions an identity has been granted through attached policies, which
	// privileges are actively used versus unused, and identifies potential
	// over-privileged access. Use this for identity-centric security assessments
	// such as privilege audits, dormant permission discovery, and least-privilege
	// compliance analysis.
	PermissionAnalysisResults []*objects.PermissionAnalysisResult `json:"permission_analysis_results,omitempty"`

	// The cyber kill chain phase.
	Phase string `json:"phase,omitempty"`

	// The cyber kill chain phase identifier.
	PhaseID int `json:"phase_id,omitempty"`

	// The number associated with the phone.
	PhoneNumber string `json:"phone_number,omitempty"`

	// The phone numbers associated with the user
	Phones []string `json:"phones,omitempty"`

	// The numeric physical height of display.
	PhysicalHeight int64 `json:"physical_height,omitempty"`

	// The numeric physical orientation of display.
	PhysicalOrientation int64 `json:"physical_orientation,omitempty"`

	// The numeric physical width of display.
	PhysicalWidth int64 `json:"physical_width,omitempty"`

	// The process identifier, as reported by the operating system. Process ID
	// (PID) is a number used by the operating system to uniquely identify an
	// active process.
	Pid int64 `json:"pid,omitempty"`

	// The unique identifier of the pod (or equivalent) that the container is
	// executing on.
	PodUUID string `json:"pod_uuid,omitempty"`

	// An array of Policy objects.
	Policies []*objects.Policy `json:"policies,omitempty"`

	// Describes details of a policy. See specific usage.
	Policy *objects.Policy `json:"policy,omitempty"`

	// The TCP/UDP port number associated with a connection. See specific usage.
	Port interface{} `json:"port,omitempty"`

	// The postal code of the location.
	PostalCode string `json:"postal_code,omitempty"`

	// The numeric precision. See specific usage.
	Precision int64 `json:"precision,omitempty"`

	// The uncorrected barometric pressure altitude (based on reference standard
	// 29.92 inHg, 1013.25 mb) provides a reference for algorithms that utilize
	// 'altitude deltas' between aircraft. This value is provided in meters and
	// must have a minimum resolution of 1 m.. Special Values: Invalid, No Value,
	// or Unknown: -1000 m.
	PressureAltitude string `json:"pressure_altitude,omitempty"`

	// The previous security level of the entity
	PrevSecurityLevel string `json:"prev_security_level,omitempty"`

	// The previous security level of the entity
	PrevSecurityLevelID int `json:"prev_security_level_id,omitempty"`

	// The previous security states. See specific usage.
	PrevSecurityStates []*objects.SecurityState `json:"prev_security_states,omitempty"`

	// The priority, normalized to the caption of the priority_id value. In the
	// case of 'Other', it is defined by the event source.
	Priority string `json:"priority,omitempty"`

	// The normalized priority. Priority identifies the relative importance of the
	// incident or finding. It is a measurement of urgency.
	PriorityID int `json:"priority_id,omitempty"`

	// The user or group privileges.
	Privileges []string `json:"privileges,omitempty"`

	// The process object.
	Process *objects.Process `json:"process,omitempty"`

	// The event processed time, such as an ETL operation.
	ProcessedTime int64 `json:"processed_time,omitempty"`

	// The event processed time, such as an ETL operation.
	ProcessedTimeDt string `json:"processed_time_dt,omitempty"`

	// The product that reported the event.
	Product *objects.Product `json:"product,omitempty"`

	// Unique Identifier of a product.
	ProductUID string `json:"product_uid,omitempty"`

	// The list of profiles used to create the event. Profiles should be referenced
	// by their name attribute for core profiles, or extension/name for profiles
	// from extensions.
	Profiles []string `json:"profiles,omitempty"`

	// Details about the programmatic credential (API key, service account key,
	// access token, certificate). See specific usage.
	ProgrammaticCredentials []*objects.ProgrammaticCredential `json:"programmatic_credentials,omitempty"`

	// The unique identifier of a Cloud project.
	ProjectUID string `json:"project_uid,omitempty"`

	// The protocol name. See specific usage.
	ProtocolName string `json:"protocol_name,omitempty"`

	// The IP protocol number, as defined by the Internet Assigned Numbers
	// Authority (IANA). For example: 6 for TCP and 17 for UDP.
	ProtocolNum int64 `json:"protocol_num,omitempty"`

	// The Protocol version, normalized to the caption of the protocol_ver_id
	// value. In the case of 'Other', it is defined by the event source.
	ProtocolVer string `json:"protocol_ver,omitempty"`

	// The normalized identifier of the Protocol version. See specific usage.
	ProtocolVerID int `json:"protocol_ver_id,omitempty"`

	// The origin of information associated with the event. See specific usage.
	Provider string `json:"provider,omitempty"`

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

	// The identifier of the process thread associated with the event, as returned
	// by the operating system.
	Ptid int64 `json:"ptid,omitempty"`

	// A purl is a URL string used to identify and locate a software package in a
	// mostly universal and uniform way across programming languages, package
	// managers, packaging conventions, tools, APIs and databases.
	Purl string `json:"purl,omitempty"`

	// The Domain Name System (DNS) query.
	Query *objects.DnsQuery `json:"query,omitempty"`

	// The resulting evidence discovered from the evidence search request.
	QueryEvidence *objects.QueryEvidence `json:"query_evidence,omitempty"`

	// The query info object holds information related to data access within a
	// datastore. To access, manipulate, delete, or retrieve data from a datastore,
	// a database query must be written using a specific syntax.
	QueryInfo *objects.QueryInfo `json:"query_info,omitempty"`

	// The query language, normalized to the caption of the query_language_id
	// value. See specific usage.
	QueryLanguage string `json:"query_language,omitempty"`

	// The normalized identifier of the query language. See specific usage.
	QueryLanguageID int `json:"query_language_id,omitempty"`

	// The result of the query.
	QueryResult string `json:"query_result,omitempty"`

	// The normalized identifier of the query result.
	QueryResultID int `json:"query_result_id,omitempty"`

	// The query portion of the URL. For example: the query portion of the URL
	// http://www.example.com/search?q=bad&sort=date is q=bad&sort=date.
	QueryString string `json:"query_string,omitempty"`

	// The Domain Name System (DNS) query time.
	QueryTime int64 `json:"query_time,omitempty"`

	// The Domain Name System (DNS) query time.
	QueryTimeDt string `json:"query_time_dt,omitempty"`

	// The normalized caption of query_type_id or the source-specific query type.
	QueryType string `json:"query_type,omitempty"`

	// The normalized type of system query performed against a device or system
	// component.
	QueryTypeID int `json:"query_type_id,omitempty"`

	// Farthest horizontal distance from the reported location at which any UA in a
	// group may be located (meters). Also allows defining the area where an
	// Intent-Based Network Participant operation is taking place. Default: 0 m.
	Radius string `json:"radius,omitempty"`

	// The total amount of installed RAM, in Megabytes. For example: 2048.
	RamSize int64 `json:"ram_size,omitempty"`

	// The rate limit for a rate-based rule.
	RateLimit int64 `json:"rate_limit,omitempty"`

	// The raw event/finding data as received from the source.
	RawData string `json:"raw_data,omitempty"`

	// The hash, which describes the content of the raw_data field.
	RawDataHash *objects.Fingerprint `json:"raw_data_hash,omitempty"`

	// The size of the raw data which was transformed into an OCSF event, in bytes.
	RawDataSize int64 `json:"raw_data_size,omitempty"`

	// The email authentication header.
	RawHeader string `json:"raw_header,omitempty"`

	// The server response code, normalized to the caption of the rcode_id value.
	// In the case of 'Other', it is defined by the event source.
	Rcode string `json:"rcode,omitempty"`

	// The normalized identifier of the server response code. See specific usage.
	RcodeID int `json:"rcode_id,omitempty"`

	// The data describing the DNS resource. The meaning of this data depends on
	// the type and class of the resource record.
	Rdata string `json:"rdata,omitempty"`

	// The index of the record in the array of records.
	RecordIndexInArray int64 `json:"record_index_in_array,omitempty"`

	// A list of reference URLs supporting the finding/detection.
	References []string `json:"references,omitempty"`

	// The request header that identifies the address of the previous web page,
	// which is linked to the current web page or resource being requested.
	Referrer string `json:"referrer,omitempty"`

	// The name or the code of a region. See specific usage.
	Region string `json:"region,omitempty"`

	// The domain registrar.
	Registrar string `json:"registrar,omitempty"`

	// Describes analytics related to the analytic of a finding or detection as
	// identified by the security product.
	RelatedAnalytics []*objects.Analytic `json:"related_analytics,omitempty"`

	// The package URL (PURL) of the component that this software component has a
	// relationship with.
	RelatedComponent string `json:"related_component,omitempty"`

	// Describes Common Vulnerabilities and Exposures https://cve.mitre.org/ (CVE)
	// entries that are related to an entity. See specific usage.
	RelatedCves []*objects.Cve `json:"related_cves,omitempty"`

	// Describes Common Weakness Enumeration https://cwe.mitre.org/ (CWE) entries
	// that are related to an entity. See specific usage.
	RelatedCwes []*objects.Cwe `json:"related_cwes,omitempty"`

	// Describes events and/or other findings related to the finding as identified
	// by the security product. Note that these events may or may not be in OCSF.
	RelatedEvents []*objects.RelatedEvent `json:"related_events,omitempty"`

	// Number of related events or findings.
	RelatedEventsCount int64 `json:"related_events_count,omitempty"`

	// List of vulnerability IDs (e.g. CVE ID) that are related to this
	// vulnerability.
	RelatedVulnerabilities []string `json:"related_vulnerabilities,omitempty"`

	// The relationship between two entities. See specific usage.
	Relation string `json:"relation,omitempty"`

	// The relationship between two software components, normalized to the caption
	// of the relationship_id value. In the case of 'Other', it is defined by the
	// source.
	Relationship string `json:"relationship,omitempty"`

	// The normalized identifier of the relationship between two software
	// components.
	RelationshipID int `json:"relationship_id,omitempty"`

	// The network relay that is associated with the event.
	Relay *objects.NetworkInterface `json:"relay,omitempty"`

	// Release is the number of times a version of the software has been packaged.
	Release string `json:"release,omitempty"`

	// Describes the recommended remediation steps to address identified issue(s).
	Remediation *objects.Remediation `json:"remediation,omitempty"`

	// The remote display affiliated with the event
	RemoteDisplay *objects.Display `json:"remote_display,omitempty"`

	// The machine-readable email header Reply-To value, as defined by RFC 5322.
	// For example example.user@usersdomain.com
	ReplyTo string `json:"reply_to,omitempty"`

	// The machine-readable email header Reply-To values, as defined by RFC 5322.
	// For example example.user@usersdomain.com
	ReplyToList []string `json:"reply_to_list,omitempty"`

	// The human-readable email header Reply To Mailbox values. For example
	// 'Example User &lt;example.user@usersdomain.com&gt;'.
	ReplyToMailboxes []string `json:"reply_to_mailboxes,omitempty"`

	// Contains the original and normalized reputation scores.
	Reputation *objects.Reputation `json:"reputation,omitempty"`

	// General Purpose API Request Object. See specific usage
	Request *objects.Request `json:"request,omitempty"`

	// The permissions mask. See specific usage.
	RequestedPermissions int64 `json:"requested_permissions,omitempty"`

	// A list of requirements associated to a specific control in an industry or
	// regulatory framework. e.g. NIST.800-53.r5 AU-10
	Requirements []string `json:"requirements,omitempty"`

	// The target resource.
	Resource *objects.ResourceDetails `json:"resource,omitempty"`

	// Describes entities related to the resource, using a graph structure. See
	// specific usage.
	ResourceRelationship *objects.Graph `json:"resource_relationship,omitempty"`

	// The resource type as defined by the event source.
	ResourceType string `json:"resource_type,omitempty"`

	// Describes details about resources that were affected by the activity/event.
	Resources []*objects.ResourceDetails `json:"resources,omitempty"`

	// Updated resources after an activity/event.
	ResourcesResult []*objects.ResourceDetails `json:"resources_result,omitempty"`

	// General Purpose API Response Object. See specific usage.
	Response *objects.Response `json:"response,omitempty"`

	// The Domain Name System (DNS) response time.
	ResponseTime int64 `json:"response_time,omitempty"`

	// The Domain Name System (DNS) response time.
	ResponseTimeDt string `json:"response_time_dt,omitempty"`

	// The address found in the 'Return-Path' header, which indicates where bounce
	// messages (non-delivery reports) should be sent. This address is often set by
	// the sending system and may differ from the 'From' or 'Sender' addresses. For
	// example, mailer-daemon@senderserver.com.
	ReturnPath string `json:"return_path,omitempty"`

	// Describes the risk associated with the finding.
	RiskDetails string `json:"risk_details,omitempty"`

	// The risk level, normalized to the caption of the risk_level_id value.
	RiskLevel string `json:"risk_level,omitempty"`

	// The normalized risk level id.
	RiskLevelID int `json:"risk_level_id,omitempty"`

	// The risk score as reported by the event source.
	RiskScore int64 `json:"risk_score,omitempty"`

	// The role of an entity in the context of the event or finding, normalized to
	// the caption of the role_id value. In the case of 'Other', it is defined by
	// the event source. See specific usage.
	Role string `json:"role,omitempty"`

	// The normalized identifier of an entity's role in the context of the event or
	// finding. See specific usage.
	RoleID int `json:"role_id,omitempty"`

	// The row number. See specific usage.
	RowNumber int64 `json:"row_number,omitempty"`

	// The RPC Interface object describes the details pertaining to the remote
	// procedure call interface.
	RpcInterface *objects.RpcInterface `json:"rpc_interface,omitempty"`

	// Received Signal Strength Indicator (RSSI) is a measurement of the power of a
	// radio signal. See specific usage.
	Rssi int64 `json:"rssi,omitempty"`

	// The rules that reported the events.
	Rule *objects.Rule `json:"rule,omitempty"`

	// The list of normalized identifiers that describe application attributes when
	// it is running. See specific usage.
	RunModeIDs []int `json:"run_mode_ids,omitempty"`

	// The list of run_modes, normalized to the captions of the run_mode_ids
	// values. In the case of 'Other', they are defined by the event source. See
	// specific usage.
	RunModes []string `json:"run_modes,omitempty"`

	// The state of the job or service, normalized to the caption of the
	// run_state_id value. In the case of 'Other', it is defined by the event
	// source. See specific usage.
	RunState string `json:"run_state,omitempty"`

	// The normalized identifier of the state of the job or service. See specific
	// usage.
	RunStateID int `json:"run_state_id,omitempty"`

	// The backend running the container, such as containerd or cri-o.
	Runtime string `json:"runtime,omitempty"`

	// The cookie attribute that lets servers specify whether/when cookies are sent
	// with cross-site requests. Values are: Strict, Lax or None
	Samesite string `json:"samesite,omitempty"`

	// The name of the containment jail (i.e., sandbox). For example, hardened_ps,
	// high_security_ps, oracle_ps, netsvcs_ps, or default_ps.
	Sandbox string `json:"sandbox,omitempty"`

	// The list of subject alternative names that are secured by a specific
	// certificate.
	Sans []*objects.San `json:"sans,omitempty"`

	// The Software Bill of Materials (SBOM) object describes the characteristics
	// of a generated SBOM for a software package.
	Sbom *objects.Sbom `json:"sbom,omitempty"`

	// The numeric scale factor of display.
	ScaleFactor int64 `json:"scale_factor,omitempty"`

	// The Scan object describes characteristics of a scan. See specific usage.
	Scan *objects.Scan `json:"scan,omitempty"`

	// The unique identifier of the schedule associated with a scan job.
	ScheduleUID string `json:"schedule_uid,omitempty"`

	// The scheme portion of the URL. For example: http, https, ftp, or sftp.
	Scheme string `json:"scheme,omitempty"`

	// The System for Cross-domain Identity Management (SCIM) resource object
	// provides a structured set of attributes related to SCIM protocols used for
	// identity provisioning and management across cloud-based platforms. It
	// standardizes user and group provisioning details, enabling identity
	// synchronization and lifecycle management with compatible Identity Providers
	// (IdPs) and applications. SCIM is defined in
	// https://datatracker.ietf.org/doc/html/rfc7643 RFC-7634
	Scim *objects.Scim `json:"scim,omitempty"`

	// SCIM provides a schema for representing groups, identified using the
	// following schema URI: urn:ietf:params:scim:schemas:core:2.0:Group as defined
	// in https://datatracker.ietf.org/doc/html/rfc7643 RFC-7634. This attribute
	// will capture key-value pairs for the scheme implemented in a SCIM resource.
	ScimGroupSchema interface{} `json:"scim_group_schema,omitempty"`

	// SCIM provides a resource type for user resources. The core schema for user
	// is identified using the following schema URI:
	// urn:ietf:params:scim:schemas:core:2.0:User as defined in
	// https://datatracker.ietf.org/doc/html/rfc7643 RFC-7634. his attribute will
	// capture key-value pairs for the scheme implemented in a SCIM resource. This
	// object is inclusive of both the basic and Enterprise User Schema Extension.
	ScimUserSchema interface{} `json:"scim_user_schema,omitempty"`

	// Scopes define the specific permissions or actions that the client is allowed
	// to perform on behalf of the user. Each scope represents a different set of
	// permissions, and the user can selectively grant or deny access to specific
	// scopes during the authorization process.
	Scopes []string `json:"scopes,omitempty"`

	// The reputation score, normalized to the caption of the score_id value. In
	// the case of 'Other', it is defined by the event source.
	Score string `json:"score,omitempty"`

	// The normalized reputation score identifier.
	ScoreID int `json:"score_id,omitempty"`

	// The script object.
	Script *objects.Script `json:"script,omitempty"`

	// The script content, normalized to UTF-8 encoding irrespective of its
	// original encoding. When emitting this attribute, it may be appropriate to
	// truncate large scripts. When consuming this attribute, large scripts should
	// be anticipated.
	ScriptContent *objects.LongString `json:"script_content,omitempty"`

	// The 'a' section of the JA4 fingerprint.
	SectionA string `json:"section_a,omitempty"`

	// The 'b' section of the JA4 fingerprint.
	SectionB string `json:"section_b,omitempty"`

	// The 'c' section of the JA4 fingerprint.
	SectionC string `json:"section_c,omitempty"`

	// The 'd' section of the JA4 fingerprint.
	SectionD string `json:"section_d,omitempty"`

	// The cookie attribute to only send cookies to the server with an encrypted
	// request over the HTTPS protocol.
	Secure bool `json:"secure,omitempty"`

	// The object security descriptor.
	SecurityDescriptor string `json:"security_descriptor,omitempty"`

	// The current security level of the entity
	SecurityLevel string `json:"security_level,omitempty"`

	// The current security level of the entity
	SecurityLevelID int `json:"security_level_id,omitempty"`

	// The question(s) provided to user for a question-based authentication factor.
	SecurityQuestions []string `json:"security_questions,omitempty"`

	// The current security states. See specific usage.
	SecurityStates []*objects.SecurityState `json:"security_states,omitempty"`

	// The machine readable email address of the system or server that actually
	// transmitted the email message, extracted from the email headers per RFC
	// 5322. This differs from the from field, which shows the message author. The
	// sender field is most commonly used when multiple addresses appear in the
	// from_list field, or when the transmitting system is different from the
	// message author (such as when sending on behalf of someone else).
	Sender string `json:"sender,omitempty"`

	// The human readable email address of the system or server that actually
	// transmitted the email message, extracted from the email headers per RFC
	// 5322. This differs from the from_mailbox field, which shows the message
	// author. The sender mailbox field is most commonly used when multiple
	// addresses appear in the from_mailboxes field, or when the transmitting
	// system is different from the message author (such as when sending on behalf
	// of someone else).
	SenderMailbox string `json:"sender_mailbox,omitempty"`

	// The sensitivity of the firewall rule in the matched event. For example:
	// HIGH.
	Sensitivity string `json:"sensitivity,omitempty"`

	// Sequence number of the event. The sequence number is a value available in
	// some events, to make the exact ordering of events unambiguous, regardless of
	// the event time precision.
	Sequence int64 `json:"sequence,omitempty"`

	// The serial number that pertains to the object. See specific usage.
	SerialNumber string `json:"serial_number,omitempty"`

	// The server cipher suites that were exchanged during the TLS handshake
	// negotiation.
	ServerCiphers []string `json:"server_ciphers,omitempty"`

	// The Server HASSH fingerprinting object.
	ServerHassh *objects.Hassh `json:"server_hassh,omitempty"`

	// The service that pertains to the event.
	Service *objects.Service `json:"service,omitempty"`

	// The authenticated user or service session.
	Session *objects.Session `json:"session,omitempty"`

	// The event/finding severity, normalized to the caption of the severity_id
	// value. In the case of 'Other', it is defined by the source.
	Severity string `json:"severity,omitempty"`

	// <p>The normalized identifier of the event/finding severity.</p>The
	// normalized severity is a measurement the effort and expense required to
	// manage and resolve an event or incident. Smaller numerical values represent
	// lower impact events, and larger numerical values represent higher impact
	// events.
	SeverityID int `json:"severity_id,omitempty"`

	// The share name. See specific usage.
	Share string `json:"share,omitempty"`

	// The share type, normalized to the caption of the share_type_id value. In the
	// case of 'Other', it is defined by the event source.
	ShareType string `json:"share_type,omitempty"`

	// The normalized identifier of the share type.
	ShareTypeID int `json:"share_type_id,omitempty"`

	// The short description that pertains to the object or event. See specific
	// usage.
	ShortDesc string `json:"short_desc,omitempty"`

	// The digital signature of the file.
	Signature *objects.DigitalSignature `json:"signature,omitempty"`

	// A collection of Digital Signature objects.
	Signatures []*objects.DigitalSignature `json:"signatures,omitempty"`

	// The size of data, in bytes.
	Size int64 `json:"size,omitempty"`

	// The value of the SMTP MAIL FROM command.
	SmtpFrom string `json:"smtp_from,omitempty"`

	// The value of the SMTP HELO or EHLO command.
	SmtpHello string `json:"smtp_hello,omitempty"`

	// The value of the SMTP envelope RCPT TO command.
	SmtpTo []string `json:"smtp_to,omitempty"`

	// The Server Name Indication (SNI) extension sent by the client.
	Sni string `json:"sni,omitempty"`

	// The list of software components used in the software package.
	SoftwareComponents []*objects.SoftwareComponent `json:"software_components,omitempty"`

	// The source of the event or object. See specific usage.
	Source string `json:"source,omitempty"`

	// The name of the latest Service Pack.
	SpName string `json:"sp_name,omitempty"`

	// The version number of the latest Service Pack.
	SpVer int64 `json:"sp_ver,omitempty"`

	// The information about the span. See specific usage.
	Span *objects.Span `json:"span,omitempty"`

	// Ground speed of flight. This value is provided in meters per second with a
	// minimum resolution of 0.25 m/s. Special Values: Invalid, No Value, or
	// Unknown: 255 m/s.
	Speed string `json:"speed,omitempty"`

	// Provides quality/containment on horizontal ground speed. Measured in
	// meters/second.
	SpeedAccuracy string `json:"speed_accuracy,omitempty"`

	// The Sender Policy Framework (SPF) status of the email.
	Spf string `json:"spf,omitempty"`

	// The network source endpoint.
	SrcEndpoint *objects.NetworkEndpoint `json:"src_endpoint,omitempty"`

	// The URL pointing towards the source of an entity. See specific usage.
	SrcURL string `json:"src_url,omitempty"`

	// The Single Sign-On (SSO) object provides a structure for normalizing SSO
	// attributes, configuration, and/or settings from Identity Providers.
	Sso *objects.Sso `json:"sso,omitempty"`

	// Compliance standards are a set of criteria organizations can follow to
	// protect sensitive and confidential information. e.g. NIST SP 800-53, CIS AWS
	// Foundations Benchmark v1.4.0, ISO/IEC 27001
	Standards []string `json:"standards,omitempty"`

	// The start address of the execution.
	StartAddress string `json:"start_address,omitempty"`

	// The start column number. See specific usage.
	StartColumn int64 `json:"start_column,omitempty"`

	// The line number of the first line of code block identified as vulnerable.
	StartLine int64 `json:"start_line,omitempty"`

	// The start time of a time period. See specific usage.
	StartTime int64 `json:"start_time,omitempty"`

	// The start time of a time period. See specific usage.
	StartTimeDt string `json:"start_time_dt,omitempty"`

	// The start type of a service, driver, or application.
	StartType string `json:"start_type,omitempty"`

	// The start type ID of a service or application.
	StartTypeID int `json:"start_type_id,omitempty"`

	// The startup item object describes an application component that has
	// associated startup criteria and configurations.
	StartupItem *objects.StartupItem `json:"startup_item,omitempty"`

	// The state of the event or object, normalized to the caption of the state_id
	// value. In the case of 'Other', it is defined by the event source. See
	// specific usage.
	State string `json:"state,omitempty"`

	// The normalized state ID of the event or object. See specific usage.
	StateID int `json:"state_id,omitempty"`

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

	// A list of descriptions, containing additional information about the
	// event/finding status.
	StatusDetails []string `json:"status_details,omitempty"`

	// The normalized identifier of the event status.
	StatusID int `json:"status_id,omitempty"`

	// The storage class of the entity. See specific usage.
	StorageClass string `json:"storage_class,omitempty"`

	// The stratum level of the NTP server's time source, normalized to the caption
	// of the stratum_id value.
	Stratum string `json:"stratum,omitempty"`

	// The normalized identifier of the stratum level, as defined in
	// https://www.rfc-editor.org/rfc/rfc5905.html RFC-5905.
	StratumID int `json:"stratum_id,omitempty"`

	// The Sub-technique object describes the MITRE ATT&CK® or ATLAS™
	// Sub-technique ID and/or name associated to an attack.
	SubTechnique *objects.SubTechnique `json:"sub_technique,omitempty"`

	// The subdomain portion of the URL. For example: sub in
	// https://sub.example.com or sub2.sub1 in https://sub2.sub1.example.com.
	Subdomain string `json:"subdomain,omitempty"`

	// An array of subdomain strings. Can be used to collect several subdomains
	// such as those from Domain Generation Algorithms (DGAs).
	Subdomains []string `json:"subdomains,omitempty"`

	// A subgroup that was added to or removed from the group.
	Subgroup *objects.Group `json:"subgroup,omitempty"`

	// The identifier of the subject. See specific usage.
	Subject string `json:"subject,omitempty"`

	// The subnet mask.
	Subnet interface{} `json:"subnet,omitempty"`

	// The subnet prefix length determines the number of bits used to represent the
	// network part of the IP address. The remaining bits are reserved for
	// identifying individual hosts within that subnet.
	SubnetPrefix int64 `json:"subnet_prefix,omitempty"`

	// The unique identifier of a virtual subnet.
	SubnetUID string `json:"subnet_uid,omitempty"`

	// Additional data supporting a finding as provided by security tool
	SupportingData interface{} `json:"supporting_data,omitempty"`

	// The last or family name for the user.
	Surname string `json:"surname,omitempty"`

	// The service name in service-to-service connections. For example, AWS VPC
	// logs the pkt-src-aws-service and pkt-dst-aws-service fields identify the
	// connection is coming from or going to an AWS service.
	SvcName string `json:"svc_name,omitempty"`

	// The system call that was invoked.
	SystemCall string `json:"system_call,omitempty"`

	// The table object represents a table within a structured relational database
	// or datastore, which contains columns and rows of data that are able to be
	// create, updated, deleted and queried.
	Table *objects.Table `json:"table,omitempty"`

	// The Tactic object describes the MITRE ATT&CK® or ATLAS™ Tactic ID and/or
	// name that is associated to an attack.
	Tactic *objects.Tactic `json:"tactic,omitempty"`

	// The Tactic object describes the tactic ID and/or tactic name that are
	// associated with the attack technique, as defined by
	// https://attack.mitre.org/wiki/ATT&CK_Matrix ATT&CK® Matrix.
	Tactics []*objects.Tactic `json:"tactics,omitempty"`

	// The image tag. For example: 1.11-alpine.
	Tag string `json:"tag,omitempty"`

	// The list of tags; {key:value} pairs attached to an entity. A Tag may be used
	// to categorize, filter, and search for entities. For e.g. "Environment":
	// "Production", "Owner": "007". See specific usage.
	Tags []*objects.KeyValueObject `json:"tags,omitempty"`

	// The target of the event or object. See specific usage.
	Target string `json:"target,omitempty"`

	// The network connection TCP header flags (i.e., control bits).
	TcpFlags int64 `json:"tcp_flags,omitempty"`

	// The state of the TCP socket for the network connection.
	TcpStateID int `json:"tcp_state_id,omitempty"`

	// The Technique object describes the MITRE ATT&CK® or ATLAS™ Technique ID
	// and/or name associated to an attack.
	Technique *objects.Technique `json:"technique,omitempty"`

	// The unique tenant identifier.
	TenantUID string `json:"tenant_uid,omitempty"`

	// The Pseudo Terminal. Ex: the tty or pts value.
	Terminal string `json:"terminal,omitempty"`

	// The time when the entity was terminated. See specific usage.
	TerminatedTime int64 `json:"terminated_time,omitempty"`

	// The time when the entity was terminated. See specific usage.
	TerminatedTimeDt string `json:"terminated_time_dt,omitempty"`

	// A threat actor is an individual or group that conducts malicious cyber
	// activities, often with financial, political, or ideological motives.
	ThreatActor *objects.ThreatActor `json:"threat_actor,omitempty"`

	// The linked ticket in the ticketing system.
	Ticket *objects.Ticket `json:"ticket,omitempty"`

	// The associated ticket(s) in the ticketing system. Each ticket contains
	// details like ticket ID, status, etc.
	Tickets []*objects.Ticket `json:"tickets,omitempty"`

	// The identifier of the thread associated with the event, as returned by the
	// operating system.
	Tid int64 `json:"tid,omitempty"`

	// The normalized event occurrence time or the finding creation time.
	Time int64 `json:"time,omitempty"`

	// The normalized event occurrence time or the finding creation time.
	TimeDt string `json:"time_dt,omitempty"`

	// The Time Span object represents different time period durations. If a
	// timespan is fractional, i.e. crosses one period, e.g. a week and 3 days,
	// more than one may be populated since each member is of integral type. In
	// that case type_id if present should be set to Other. A timespan may also be
	// defined by its time interval boundaries, start_time and end_time.
	Timespan *objects.Timespan `json:"timespan,omitempty"`

	// The number of minutes that the reported event time is ahead or behind UTC,
	// in the range -1,080 to +1,080.
	TimezoneOffset int64 `json:"timezone_offset,omitempty"`

	// The title of an entity. See specific usage.
	Title string `json:"title,omitempty"`

	// The https://www.first.org/tlp/ Traffic Light Protocol was created to
	// facilitate greater sharing of potentially sensitive information and more
	// effective collaboration. TLP provides a simple and intuitive schema for
	// indicating with whom potentially sensitive information can be shared.
	Tlp string `json:"tlp,omitempty"`

	// The Transport Layer Security (TLS) attributes.
	Tls *objects.Tls `json:"tls,omitempty"`

	// The list of TLS extensions.
	TlsExtensionList []*objects.TlsExtension `json:"tls_extension_list,omitempty"`

	// The machine-readable email header To values, as defined by RFC 5322. For
	// example example.user@usersdomain.com
	To []string `json:"to,omitempty"`

	// The human-readable email header To Mailbox values. For example 'Example User
	// &lt;example.user@usersdomain.com&gt;'.
	ToMailboxes []string `json:"to_mailboxes,omitempty"`

	// The total number of items. See specific usage.
	Total int64 `json:"total,omitempty"`

	// The information about the trace. See specific usage.
	Trace *objects.Trace `json:"trace,omitempty"`

	// Direction of flight expressed as a “True North-based” ground track
	// angle. This value is provided in clockwise degrees with a minimum resolution
	// of 1 degree. If aircraft is not moving horizontally, use the “Unknown”
	// value
	TrackDirection string `json:"track_direction,omitempty"`

	// The network traffic refers to the amount of data moving across a network at
	// a given point of time. Intended to be used alongside Network Connection.
	Traffic *objects.NetworkTraffic `json:"traffic,omitempty"`

	// The traits that describe characteristics, features of an entity. See
	// specific usage.
	Traits []*objects.Trait `json:"traits,omitempty"`

	// The unique identifier of the transaction.
	TransactionUID string `json:"transaction_uid,omitempty"`

	// An array of transformation info that describes the mappings or transforms
	// applied to the data.
	TransformationInfoList []*objects.TransformationInfo `json:"transformation_info_list,omitempty"`

	// The event transmission time from one device to another. See specific usage.
	TransmitTime int64 `json:"transmit_time,omitempty"`

	// The event transmission time from one device to another. See specific usage.
	TransmitTimeDt string `json:"transmit_time_dt,omitempty"`

	// The tree id is a unique SMB identifier which represents an open connection
	// to a share.
	TreeUID string `json:"tree_uid,omitempty"`

	// The time interval that the resource record may be cached. Zero value means
	// that the resource record can only be used for the transaction in progress,
	// and should not be cached.
	Ttl int64 `json:"ttl,omitempty"`

	// The information about the tunnel interface. See specific usage.
	TunnelInterface *objects.NetworkInterface `json:"tunnel_interface,omitempty"`

	// The tunnel type. See specific usage.
	TunnelType string `json:"tunnel_type,omitempty"`

	// The normalized tunnel type ID.
	TunnelTypeID int64 `json:"tunnel_type_id,omitempty"`

	// The type of an object or value, normalized to the caption of the type_id
	// value. In the case of 'Other', it is defined by the event source. See
	// specific usage.
	Type string `json:"type,omitempty"`

	// The normalized type identifier of an object. See specific usage.
	TypeID int `json:"type_id,omitempty"`

	// The event/finding type name, as defined by the type_uid.
	TypeName string `json:"type_name,omitempty"`

	// The event/finding type ID. It identifies the event's semantics and
	// structure. The value is calculated by the logging system as: class_uid * 100
	// + activity_id.
	TypeUID int64 `json:"type_uid,omitempty"`

	// The type/s of an entity. See specific usage.
	Types []string `json:"types,omitempty"`

	// The Apple assigned Unique Device Identifier (UDID). For iOS, iPadOS, tvOS,
	// watchOS and visionOS devices, this is the UDID. For macOS devices, it is the
	// Provisioning UDID. For example: 00008020-008D4548007B4F26
	Udid string `json:"udid,omitempty"`

	// The unique identifier. See specific usage.
	Uid string `json:"uid,omitempty"`

	// The alternate unique identifier. See specific usage.
	UidAlt string `json:"uid_alt,omitempty"`

	// The number of unique malware detected during a scan. See specific usage.
	UniqueMalwareCount int64 `json:"unique_malware_count,omitempty"`

	// The Unmanned Aerial System object describes the characteristics, Position
	// Location Information (PLI), and other metadata of Unmanned Aerial Systems
	// (UAS) and other unmanned and drone systems used in Remote ID. Remote ID is
	// defined in the Standard Specification for Remote ID and Tracking (ASTM
	// Designation: F3411-22a)
	// https://cdn.standards.iteh.ai/samples/112830/71297057ac42432880a203654f213709/ASTM-F3411-22a.pdf
	// ASTM F3411-22a.
	UnmannedAerialSystem *objects.UnmannedAerialSystem `json:"unmanned_aerial_system,omitempty"`

	// The UAS Operating Area object describes details about a precise area of
	// operations for a UAS flight or mission.
	UnmannedSystemOperatingArea *objects.UnmannedSystemOperatingArea `json:"unmanned_system_operating_area,omitempty"`

	// The human or machine operator of an Unmanned System.
	UnmannedSystemOperator *objects.User `json:"unmanned_system_operator,omitempty"`

	// The attributes that are not mapped to the event schema. The names and values
	// of those attributes are specific to the event source.
	Unmapped *objects.Object `json:"unmapped,omitempty"`

	// The size in bytes of an attribute before truncation. See specific usage.
	UntruncatedSize int64 `json:"untruncated_size,omitempty"`

	// The number of unused Privileges. See specific usage.
	UnusedPrivilegesCount int64 `json:"unused_privileges_count,omitempty"`

	// The number of unused services. See specific usage.
	UnusedServicesCount int64 `json:"unused_services_count,omitempty"`

	// The timestamp at which an entity was uploaded. See specific usage.
	UploadedTime int64 `json:"uploaded_time,omitempty"`

	// The timestamp at which an entity was uploaded. See specific usage.
	UploadedTimeDt string `json:"uploaded_time_dt,omitempty"`

	// A Uniform Resource Identifier (URI) is a string of characters that
	// identifies a resource on the Internet. URIs can identify physical resources,
	// like webpages and documents. See specific usage.
	Uri string `json:"uri,omitempty"`

	// The URL object that pertains to the event or object. See specific usage.
	Url *objects.Url `json:"url,omitempty"`

	// The URL string. See RFC 1738. For example:
	// http://www.example.com/download/trouble.exe.
	UrlString string `json:"url_string,omitempty"`

	// The URLs that pertain to the event or object.
	Urls []*objects.Url `json:"urls,omitempty"`

	// The user that pertains to the event or object.
	User *objects.User `json:"user,omitempty"`

	// The request header that identifies the operating system and web browser.
	UserAgent string `json:"user_agent,omitempty"`

	// The result of the user account change. It should contain the new values of
	// the changed attributes.
	UserResult *objects.User `json:"user_result,omitempty"`

	// The users that pertain to the event or object.
	Users []*objects.User `json:"users,omitempty"`

	// The universally unique identifier. See specific usage.
	Uuid string `json:"uuid,omitempty"`

	// The value associated to an attribute. See specific usage.
	Value string `json:"value,omitempty"`

	// An array of values associated to an attribute. See specific usage.
	Values []string `json:"values,omitempty"`

	// The CVSS vector string is a text representation of a set of CVSS metrics. It
	// is commonly used to record or transfer CVSS metric information in a concise
	// form. For example: 3.1/AV:L/AC:L/PR:L/UI:N/S:U/C:H/I:N/A:H.
	VectorString string `json:"vector_string,omitempty"`

	// The Vendor Attributes object can be used to represent values of attributes
	// populated by the Vendor/Finding Provider. It can help distinguish between
	// the vendor-provided values and consumer-updated values, of key attributes
	// like severity_id. The original finding producer should not populate this
	// object. It should be populated by consuming systems that support data
	// mutability.
	VendorAttributes *objects.VendorAttributes `json:"vendor_attributes,omitempty"`

	// The name of the vendor. See specific usage.
	VendorName string `json:"vendor_name,omitempty"`

	// The verdict assigned to an Incident finding.
	Verdict string `json:"verdict,omitempty"`

	// The normalized verdict of an Incident.
	VerdictID int `json:"verdict_id,omitempty"`

	// The version that pertains to the event or object. See specific usage.
	Version string `json:"version,omitempty"`

	// Vertical speed upward relative to the WGS-84 datum, measured in meters per
	// second. Special Values: Invalid, No Value, or Unknown: 63 m/s.
	VerticalSpeed string `json:"vertical_speed,omitempty"`

	// The Virtual LAN identifier.
	VlanUID string `json:"vlan_uid,omitempty"`

	// The volume on the storage device where the file is located. See specific
	// usage.
	Volume string `json:"volume,omitempty"`

	// The unique identifier of the Virtual Private Cloud (VPC).
	VpcUID string `json:"vpc_uid,omitempty"`

	// This object describes vulnerabilities reported in a security finding.
	Vulnerabilities []*objects.Vulnerability `json:"vulnerabilities,omitempty"`

	// The vulnerability object describes details related to the observed
	// vulnerability
	Vulnerability *objects.Vulnerability `json:"vulnerability,omitempty"`

	// Describes details about web resources that were affected by an
	// activity/event.
	WebResources []*objects.WebResource `json:"web_resources,omitempty"`

	// The results of the activity on web resources. It should contain the new
	// values of the changed attributes of the web resources.
	WebResourcesResult []*objects.WebResource `json:"web_resources_result,omitempty"`

	// The resources of a WHOIS record for a given domain. This can include domain
	// names, IP address blocks, autonomous system information, and/or contact and
	// registration information for a domain.
	Whois *objects.Whois `json:"whois,omitempty"`

	// The name of the load ordering group of which this service is a member.
	WinLoadOrderGroup string `json:"win/load_order_group,omitempty"`

	// The registry key before the mutation
	WinPrevRegKey *objects.WinRegKey `json:"win/prev_reg_key,omitempty"`

	// The registry value before the mutation
	WinPrevRegValue *objects.WinRegValue `json:"win/prev_reg_value,omitempty"`

	// The data of the registry value when type_id is REG_BINARY or REG_NONE.
	WinRegBinaryData interface{} `json:"win/reg_binary_data,omitempty"`

	// The data of the registry value when type_id is REG_DWORD,
	// REG_DWORD_BIG_ENDIAN, or REG_QWORD.
	WinRegIntegerData int64 `json:"win/reg_integer_data,omitempty"`

	// The registry key.
	WinRegKey *objects.WinRegKey `json:"win/reg_key,omitempty"`

	// The data of the registry value when type_id is REG_SZ, REG_EXPAND_SZ, or
	// REG_LINK.
	WinRegStringData string `json:"win/reg_string_data,omitempty"`

	// The data of the registry value when type_id is REG_MULTI_SZ.
	WinRegStringListData []string `json:"win/reg_string_list_data,omitempty"`

	// The registry value.
	WinRegValue *objects.WinRegValue `json:"win/reg_value,omitempty"`

	// The prefetch file run count.
	WinRunCount int64 `json:"win/run_count,omitempty"`

	// The service category, normalized to the caption of the service_category_id
	// value. In the case of 'Other', it is defined by the event source.
	WinServiceCategory string `json:"win/service_category,omitempty"`

	// The normalized identifier of the service category.
	WinServiceCategoryID int `json:"win/service_category_id,omitempty"`

	// The names of other services upon which this service has a dependency.
	WinServiceDependencies []string `json:"win/service_dependencies,omitempty"`

	// The service error control, normalized to the caption of the
	// service_error_control_id value. In the case of 'Other', it is defined by the
	// event source.
	WinServiceErrorControl string `json:"win/service_error_control,omitempty"`

	// The normalized identifier of the service error control.
	WinServiceErrorControlID int `json:"win/service_error_control_id,omitempty"`

	// For a user mode service, this attribute represents the name of the account
	// under which the service is run. For a kernel mode driver, this attribute
	// represents the object name used to load the driver.
	WinServiceStartName string `json:"win/service_start_name,omitempty"`

	// The service start type, normalized to the caption of the
	// service_start_type_id value. In the case of 'Other', it is defined by the
	// event source.
	WinServiceStartType string `json:"win/service_start_type,omitempty"`

	// The normalized identifier of the service start type.
	WinServiceStartTypeID int `json:"win/service_start_type_id,omitempty"`

	// The service type, normalized to the caption of the service_type_id value. In
	// the case of 'Other', it is defined by the event source.
	WinServiceType string `json:"win/service_type,omitempty"`

	// The normalized identifier of the service type.
	WinServiceTypeID int `json:"win/service_type_id,omitempty"`

	// The Windows resource object that was accessed, such as a mutant or timer.
	WinWinResource *objects.WinWinResource `json:"win/win_resource,omitempty"`

	// The Windows service.
	WinWinService *objects.WinWinService `json:"win/win_service,omitempty"`

	// The working directory of a process.
	WorkingDirectory string `json:"working_directory,omitempty"`

	// The X-Forwarded-For header identifying the originating IP address(es) of a
	// client connecting to a web server through an HTTP proxy or a load balancer.
	XForwardedFor []string `json:"x_forwarded_for,omitempty"`

	// The X-Originating-IP header identifying the emails originating IP
	// address(es).
	XOriginatingIP []string `json:"x_originating_ip,omitempty"`

	// An unordered collection of zero or more name/value pairs where each pair
	// represents a file or folder extended attribute.</p>For example: Windows
	// alternate data stream attributes (ADS stream name, ADS size, etc.),
	// user-defined or application-defined attributes, ACL, owner, primary group,
	// etc. Examples from DCS:
	// </p><ul><li>ads_name</li><li>ads_size</li><li>dacl</li><li>owner</li><li>primary_group</li><li>link_name
	// - name of the link associated to the file.</li><li>hard_link_count - the
	// number of links that are associated to the file.</li></ul>
	Xattributes *objects.Object `json:"xattributes,omitempty"`

	// The network zone or LAN segment.
	Zone string `json:"zone,omitempty"`
}

// Code generated from OCSF schema. DO NOT EDIT.
package gocsf

import (
	"encoding/json"
	"time"
)

// Security Types

// Attack The <a target='_blank' href='https://attack.mitre.org'>MITRE ATT&CK®</a> object describes the tactic, technique & sub-technique associated to an attack as defined in <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
type Attack struct {
	// The Sub Technique object describes the sub technique ID and/or name associated to an attack, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
	SubTechnique *SubTechnique `json:"sub_technique,omitempty"`
	// The Tactic object describes the tactic ID and/or name that is associated to an attack, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
	Tactic *Tactic `json:"tactic,omitempty"`
	// The Tactic object describes the tactic ID and/or tactic name that are associated with the attack technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
	Tactics []*Tactic `json:"tactics,omitempty"`
	// The Technique object describes the technique ID and/or name associated to an attack, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
	Technique *Technique `json:"technique,omitempty"`
	// The <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a> version.
	Version string `json:"version"`
}

// D3fTactic The MITRE D3FEND™ Tactic object describes the tactic ID and/or name that is associated to an attack, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND<sup>TM</sup> Matrix</a>.
type D3fTactic struct {
	// The tactic name that is associated with the defensive technique, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND<sup>TM</sup> Matrix</a>. For example: <code>Isolate</code>.
	Name string `json:"name,omitempty"`
	// The versioned permalink of the defensive tactic, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND<sup>TM</sup> Matrix</a>. For example: <code>https://d3fend.mitre.org/tactic/d3f:Isolate/</code>.
	SrcUrl string `json:"src_url,omitempty"`
	// The unique identifier of the entity.
	Uid string `json:"uid"`
}

// D3fTechnique The MITRE DEFEND™ Technique object describes the leaf defensive technique ID and/or name associated to a countermeasure, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND<sup>TM</sup> Matrix</a>.
type D3fTechnique struct {
	// The name of the defensive technique, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND<sup>TM</sup> Matrix</a>. For example: <code>IO Port Restriction</code>.
	Name string `json:"name"`
	// The versioned permalink of the defensive technique, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND<sup>TM</sup> Matrix</a>. For example: <code>https://d3fend.mitre.org/technique/d3f:IOPortRestriction/</code>.
	SrcUrl string `json:"src_url,omitempty"`
	// The unique identifier of the defensive technique, as defined by <a target='_blank' href='https://mitre.mitre.org'>D3FEND<sup>TM</sup> Matrix</a>. For example: <code>D3-IOPR</code>.
	Uid string `json:"uid"`
}

// D3fend The <a target='_blank' href='https://d3fend.mitre.org'>MITRE D3FEND™</a> object describes the tactic, technique & sub-technique associated with a countermeasure as defined in <a target='_blank' href='https://d3fend.mitre.org/'>DEFEND Matrix<sup>TM</sup></a>.
type D3fend struct {
	// The Tactic object describes the tactic ID and/or name that is associated with a countermeasure, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND Matrix<sup>TM</sup></a>.
	D3fTactic *D3fTactic `json:"d3f_tactic"`
	// The Defend Technique object describes the technique ID and/or name associated with a countermeasure, as defined by <a target='_blank' href='https://d3fend.mitre.org'>D3FEND Matrix<sup>TM</sup></a>.
	D3fTechnique *D3fTechnique `json:"d3f_technique"`
	// The <a target='_blank' href='https://d3fend.mitre.org'>D3FEND Matrix<sup>TM</sup></a> version.
	Version string `json:"version"`
}

// Enrichment The Enrichment object provides inline enrichment data for specific attributes of interest within an event. It serves as a mechanism to enhance or supplement the information associated with the event by adding additional relevant details or context.
type Enrichment struct {
	// The time when the enrichment data was generated.
	CreatedTime time.Time `json:"created_time"`
	// The time when the enrichment data was generated.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The enrichment data associated with the attribute and value. The meaning of this data depends on the type the enrichment record.
	Data json.RawMessage `json:"data"`
	// A long description of the enrichment data.
	Desc string `json:"desc,omitempty"`
	// The name of the attribute to which the enriched data pertains.
	Name string `json:"name"`
	// The enrichment data provider name.
	Provider string `json:"provider"`
	// The reputation of the enrichment data.
	Reputation *Reputation `json:"reputation,omitempty"`
	// A short description of the enrichment data.
	ShortDesc string `json:"short_desc"`
	// The URL of the source of the enrichment data.
	SrcUrl string `json:"src_url"`
	// The enrichment type. For example: <code>location</code>.
	Type string `json:"type"`
	// The value of the attribute to which the enriched data pertains.
	Value string `json:"value"`
}

// Finding The Finding object describes metadata related to a security finding generated by a security tool or system.
type Finding struct {
	// The time when the finding was created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the finding was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The description of the reported finding.
	Desc string `json:"desc,omitempty"`
	// The time when the finding was first observed.
	FirstSeenTime time.Time `json:"first_seen_time,omitempty"`
	// The time when the finding was first observed.
	FirstSeenTimeDt time.Time `json:"first_seen_time_dt,omitempty"`
	// The time when the finding was most recently observed.
	LastSeenTime time.Time `json:"last_seen_time,omitempty"`
	// The time when the finding was most recently observed.
	LastSeenTimeDt time.Time `json:"last_seen_time_dt,omitempty"`
	// The time when the finding was last modified.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The time when the finding was last modified.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The unique identifier of the product that reported the finding.
	ProductUid string `json:"product_uid,omitempty"`
	// Describes events and/or other findings related to the finding as identified by the security product.
	RelatedEvents []*RelatedEvent `json:"related_events,omitempty"`
	// Describes the recommended remediation steps to address identified issue(s).
	Remediation *Remediation `json:"remediation,omitempty"`
	// The URL pointing to the source of the finding.
	SrcUrl string `json:"src_url,omitempty"`
	// Additional data supporting a finding as provided by security tool
	SupportingData json.RawMessage `json:"supporting_data,omitempty"`
	// A title or a brief phrase summarizing the reported finding.
	Title string `json:"title"`
	// One or more types of the reported finding.
	Types []string `json:"types,omitempty"`
	// The unique identifier of the reported finding.
	Uid string `json:"uid"`
}

// FindingInfo The Finding Information object describes metadata related to a security finding generated by a security tool or system.
type FindingInfo struct {
	// The analytic technique used to analyze and derive insights from the data or information that led to the finding or conclusion.
	Analytic *Analytic `json:"analytic"`
	// The <a target='_blank' href='https://attack.mitre.org'>MITRE ATT&CK®</a> technique and associated tactics related to the finding.
	Attacks []*Attack `json:"attacks,omitempty"`
	// The time when the finding was created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the finding was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// A list of data sources utilized in generation of the finding.
	DataSources []string `json:"data_sources,omitempty"`
	// The description of the reported finding.
	Desc string `json:"desc,omitempty"`
	// The time when the finding was first observed. e.g. The time when a vulnerability was first observed. <p>It can differ from the <code>created_time</code> timestamp, which reflects the time this finding was created.</p>
	FirstSeenTime time.Time `json:"first_seen_time,omitempty"`
	// The time when the finding was first observed. e.g. The time when a vulnerability was first observed. <p>It can differ from the <code>created_time</code> timestamp, which reflects the time this finding was created.</p>
	FirstSeenTimeDt time.Time `json:"first_seen_time_dt,omitempty"`
	// The <a target='_blank' href='https://www.lockheedmartin.com/en-us/capabilities/cyber/cyber-kill-chain.html'>Cyber Kill Chain®</a> provides a detailed description of each phase and its associated activities within the broader context of a cyber attack.
	KillChain []*KillChainPhase `json:"kill_chain,omitempty"`
	// The time when the finding was most recently observed. e.g. The time when a vulnerability was most recently observed. <p>It can differ from the <code>modified_time</code> timestamp, which reflects the time this finding was last modified.</p>
	LastSeenTime time.Time `json:"last_seen_time,omitempty"`
	// The time when the finding was most recently observed. e.g. The time when a vulnerability was most recently observed. <p>It can differ from the <code>modified_time</code> timestamp, which reflects the time this finding was last modified.</p>
	LastSeenTimeDt time.Time `json:"last_seen_time_dt,omitempty"`
	// The time when the finding was last modified.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The time when the finding was last modified.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The unique identifier of the product that reported the finding.
	ProductUid string `json:"product_uid,omitempty"`
	// Other analytics related to this finding.
	RelatedAnalytics []*Analytic `json:"related_analytics,omitempty"`
	// Describes events and/or other findings related to the finding as identified by the security product.
	RelatedEvents []*RelatedEvent `json:"related_events,omitempty"`
	// The URL pointing to the source of the finding.
	SrcUrl string `json:"src_url,omitempty"`
	// A title or a brief phrase summarizing the reported finding.
	Title string `json:"title"`
	// One or more types of the reported finding.
	Types []string `json:"types,omitempty"`
	// The unique identifier of the reported finding.
	Uid string `json:"uid"`
}

// KillChainPhase The Kill Chain Phase object represents a single phase of a cyber attack, including the initial reconnaissance and planning stages up to the final objective of the attacker. It provides a detailed description of each phase and its associated activities within the broader context of a cyber attack. See <a target='_blank' href='https://www.lockheedmartin.com/en-us/capabilities/cyber/cyber-kill-chain.html'>Cyber Kill Chain®</a>.
type KillChainPhase struct {
	// The cyber kill chain phase.
	Phase string `json:"phase"`
	// The cyber kill chain phase identifier.
	PhaseId int `json:"phase_id"`
}

// Malware The Malware object describes the classification of known malicious software, which is intentionally designed to cause damage to a computer, server, client, or computer network.
type Malware struct {
	// The list of normalized identifiers of the malware classifications. Reference: <a target='_blank' href='https://docs.oasis-open.org/cti/stix/v2.1/os/stix-v2.1-os.html#_oxlc4df65spl'>STIX Malware Types</a> 
	ClassificationIds []int `json:"classification_ids"`
	// The list of malware classifications, normalized to the captions of the <code>classification_ids</code> values. In the case of 'Other', they are defined by the event source.
	Classifications []string `json:"classifications,omitempty"`
	// List of Common Vulnerabilities and Exposures (<a target='_blank' href='https://cve.mitre.org/'>CVE</a>).
	Cves []*Cve `json:"cves,omitempty"`
	// The malware name, as reported by the detection engine.
	Name string `json:"name"`
	// The filesystem path of the malware that was observed.
	Path string `json:"path"`
	// The provider of the malware information.
	Provider string `json:"provider"`
	// The malware unique identifier, as reported by the detection engine. For example a virus id or an IPS signature id.
	Uid string `json:"uid"`
}

// Osint The OSINT (Open Source Intelligence) object contains details related to an indicator such as the indicator itself, related indicators, geolocation, registrar information, subdomains, analyst commentary, and other contextual information. This information can be used to further enrich a detection or finding by providing decisioning support to other analysts and engineers.
type Osint struct {
	// Any pertinent DNS answers information related to an indicator or OSINT analysis.
	Answers []*DnsAnswer `json:"answers,omitempty"`
	// MITRE ATT&CK Tactics, Techniques, and/or Procedures (TTPs) pertinent to an indicator or OSINT analysis.
	Attacks []*Attack `json:"attacks,omitempty"`
	// Any pertinent autonomous system information related to an indicator or OSINT analysis.
	AutonomousSystem *AutonomousSystem `json:"autonomous_system,omitempty"`
	// Analyst commentary or source commentary about an indicator or OSINT analysis.
	Comment string `json:"comment,omitempty"`
	// The confidence of an indicator being malicious and/or pertinent, normalized to the caption of the confidence_id value. In the case of 'Other', it is defined by the event source or analyst.
	Confidence string `json:"confidence,omitempty"`
	// The normalized confidence refers to the accuracy of collected information related to the OSINT or how pertinent an indicator or analysis is to a specific event or finding. A low confidence means that the information collected or analysis conducted lacked detail or is not accurate enough to qualify an indicator as fully malicious.
	ConfidenceId int `json:"confidence_id"`
	// Any email information pertinent to an indicator or OSINT analysis.
	Email *Email `json:"email,omitempty"`
	// Any email authentication information pertinent to an indicator or OSINT analysis.
	EmailAuth *EmailAuth `json:"email_auth,omitempty"`
	// Lockheed Martin Kill Chain Phases pertinent to an indicator or OSINT analysis.
	KillChain []*KillChainPhase `json:"kill_chain,omitempty"`
	// Any pertinent geolocation information related to an indicator or OSINT analysis.
	Location *Location `json:"location,omitempty"`
	// The name of the entity.
	Name string `json:"name"`
	// Any digital signatures or hashes related to an indicator or OSINT analysis.
	Signatures []*DigitalSignature `json:"signatures,omitempty"`
	// The source URL of an indicator or OSINT analysis, e.g., a URL back to a TIP, report, or otherwise.
	SrcUrl string `json:"src_url,omitempty"`
	// Any pertinent subdomain information - such as those generated by a Domain Generation Algorithm - related to an indicator or OSINT analysis.
	Subdomains []string `json:"subdomains,omitempty"`
	// The <a target='_blank' href='https://www.first.org/tlp/'>Traffic Light Protocol</a> was created to facilitate greater sharing of potentially sensitive information and more effective collaboration. TLP provides a simple and intuitive schema for indicating with whom potentially sensitive information can be shared.
	Tlp string `json:"tlp"`
	// The OSINT indicator type.
	Type string `json:"type,omitempty"`
	// The OSINT indicator type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the entity.
	Uid string `json:"uid"`
	// The actual indicator value in scope, e.g., a SHA-256 hash hexdigest or a domain name.
	Value string `json:"value"`
	// The vendor name of a tool which generates intelligence or provides indicators.
	VendorName string `json:"vendor_name,omitempty"`
	// Any vulnerabilities related to an indicator or OSINT analysis.
	Vulnerabilities []*Vulnerability `json:"vulnerabilities,omitempty"`
	// Any pertinent WHOIS information related to an indicator or OSINT analysis.
	Whois *Whois `json:"whois,omitempty"`
}

// Reputation The Reputation object describes the reputation/risk score of an entity (e.g. device, user, domain).
type Reputation struct {
	// The reputation score as reported by the event source.
	BaseScore float64 `json:"base_score"`
	// The provider of the reputation information.
	Provider string `json:"provider"`
	// The reputation score, normalized to the caption of the score_id value. In the case of 'Other', it is defined by the event source.
	Score string `json:"score,omitempty"`
	// The normalized reputation score identifier.
	ScoreId int `json:"score_id"`
}

// SubTechnique The MITRE ATT&CK® Sub Technique object describes the sub technique ID and/or name associated to an attack, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
type SubTechnique struct {
	// The name of the attack sub technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>Scanning IP Blocks</code>.
	Name string `json:"name,omitempty"`
	// The versioned permalink of the attack sub technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>https://attack.mitre.org/versions/v14/techniques/T1595/001/</code>.
	SrcUrl string `json:"src_url,omitempty"`
	// The unique identifier of the attack sub technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>T1595.001</code>.
	Uid string `json:"uid"`
}

// Tactic The MITRE ATT&CK® Tactic object describes the tactic ID and/or name that is associated to an attack, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
type Tactic struct {
	// The tactic name that is associated with the attack technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>Reconnaissance</code>.
	Name string `json:"name,omitempty"`
	// The versioned permalink of the attack tactic, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>https://attack.mitre.org/versions/v14/tactics/TA0043/</code>.
	SrcUrl string `json:"src_url,omitempty"`
	// The tactic ID that is associated with the attack technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>TA0043</code>.
	Uid string `json:"uid"`
}

// Technique The MITRE ATT&CK® Technique object describes the technique ID and/or name associated to an attack, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>.
type Technique struct {
	// The name of the attack technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>Active Scanning</code>.
	Name string `json:"name"`
	// The versioned permalink of the attack technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>https://attack.mitre.org/versions/v14/techniques/T1595/</code>.
	SrcUrl string `json:"src_url,omitempty"`
	// The unique identifier of the attack technique, as defined by <a target='_blank' href='https://attack.mitre.org/wiki/ATT&CK_Matrix'>ATT&CK® Matrix</a>. For example: <code>T1595</code>.
	Uid string `json:"uid"`
}


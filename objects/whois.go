package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Whois The resources of a WHOIS record for a given domain. This can include domain
// names, IP address blocks, autonomous system information, and/or contact and
// registration information for a domain.
type Whois struct {
	// The autonomous system information associated with a domain.
	AutonomousSystem *AutonomousSystem `json:"autonomous_system,omitempty"`

	// When the domain was registered or WHOIS entry was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// When the domain was registered or WHOIS entry was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The normalized value of dnssec_status_id.
	DnssecStatus string `json:"dnssec_status,omitempty"`

	// Describes the normalized status of DNS Security Extensions (DNSSEC) for a
	// domain.
	DnssecStatusID int `json:"dnssec_status_id,omitempty"`

	// The domain name corresponding to the WHOIS record.
	Domain string `json:"domain,omitempty"`

	// An array of Domain Contact objects.
	DomainContacts []*DomainContact `json:"domain_contacts,omitempty"`

	// The email address for the registrar's abuse contact
	EmailAddr string `json:"email_addr,omitempty"`

	// The name of the Internet Service Provider (ISP).
	Isp string `json:"isp,omitempty"`

	// The organization name of the Internet Service Provider (ISP). This
	// represents the parent organization or company that owns/operates the ISP.
	// For example, Comcast Corporation would be the ISP org for Xfinity internet
	// service. This attribute helps identify the ultimate provider when ISPs
	// operate under different brand names.
	IspOrg string `json:"isp_org,omitempty"`

	// When the WHOIS record was last updated or seen at.
	LastSeenTime int64 `json:"last_seen_time,omitempty"`

	// When the WHOIS record was last updated or seen at.
	LastSeenTimeDt string `json:"last_seen_time_dt,omitempty"`

	// A collection of name servers related to a domain registration or other
	// record.
	NameServers []string `json:"name_servers,omitempty"`

	// The phone number for the registrar's abuse contact
	PhoneNumber string `json:"phone_number,omitempty"`

	// The domain registrar.
	Registrar string `json:"registrar,omitempty"`

	// The status of a domain and its ability to be transferred, e.g.,
	// clientTransferProhibited.
	Status string `json:"status,omitempty"`

	// An array of subdomain strings. Can be used to collect several subdomains
	// such as those from Domain Generation Algorithms (DGAs).
	Subdomains []string `json:"subdomains,omitempty"`

	// The IP address block (CIDR) associated with a domain.
	Subnet interface{} `json:"subnet,omitempty"`
}

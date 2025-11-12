package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DnsAnswer The DNS Answer object represents a specific response provided by the Domain Name
// System (DNS) when querying for information about a domain or performing a DNS
// operation. It encapsulates the relevant details and data returned by the DNS
// server in response to a query.
type DnsAnswer struct {
	// The class of DNS data contained in this resource record. See
	// https://www.rfc-editor.org/rfc/rfc1035.txt RFC1035. For example: IN.
	Class string `json:"class,omitempty"`

	// The list of DNS answer header flag IDs.
	FlagIDs []int `json:"flag_ids,omitempty"`

	// The list of DNS answer header flags.
	Flags []string `json:"flags,omitempty"`

	// The DNS packet identifier assigned by the program that generated the query.
	// The identifier is copied to the response.
	PacketUID int64 `json:"packet_uid,omitempty"`

	// The data describing the DNS resource. The meaning of this data depends on
	// the type and class of the resource record.
	Rdata string `json:"rdata,omitempty"`

	// The time interval that the resource record may be cached. Zero value means
	// that the resource record can only be used for the transaction in progress,
	// and should not be cached.
	Ttl int64 `json:"ttl,omitempty"`

	// The type of data contained in this resource record. See
	// https://www.rfc-editor.org/rfc/rfc1035.txt RFC1035. For example: CNAME.
	Type string `json:"type,omitempty"`
}

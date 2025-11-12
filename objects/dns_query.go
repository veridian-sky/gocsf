package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DnsQuery The DNS query object represents a specific request made to the Domain Name
// System (DNS) to retrieve information about a domain or perform a DNS operation.
// This object encapsulates the necessary attributes and methods to construct and
// send DNS queries, specify the query type (e.g., A, AAAA, MX).
type DnsQuery struct {
	// The class of resource records being queried. See
	// https://www.rfc-editor.org/rfc/rfc1035.txt RFC1035. For example: IN.
	Class string `json:"class,omitempty"`

	// The hostname or domain being queried. For example: www.example.com
	Hostname string `json:"hostname,omitempty"`

	// The DNS opcode specifies the type of the query message.
	Opcode string `json:"opcode,omitempty"`

	// The DNS opcode ID specifies the normalized query message type as defined in
	// https://www.rfc-editor.org/rfc/rfc5395.html RFC-5395.
	OpcodeID int `json:"opcode_id,omitempty"`

	// The DNS packet identifier assigned by the program that generated the query.
	// The identifier is copied to the response.
	PacketUID int64 `json:"packet_uid,omitempty"`

	// The type of resource records being queried. See
	// https://www.rfc-editor.org/rfc/rfc1035.txt RFC1035. For example: A, AAAA,
	// CNAME, MX, and NS.
	Type string `json:"type,omitempty"`
}

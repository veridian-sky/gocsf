package objects

// Code generated from OCSF schema; DO NOT EDIT.

// PortInfo The Port Information object describes a port and its associated protocol
// details.
type PortInfo struct {
	// The port number. For example: 80, 443, 22.
	Port interface{} `json:"port,omitempty"`

	// The IP protocol name in lowercase, as defined by the Internet Assigned
	// Numbers Authority (IANA). For example: tcp or udp.
	ProtocolName string `json:"protocol_name,omitempty"`

	// The IP protocol number, as defined by the Internet Assigned Numbers
	// Authority (IANA). For example: 6 for TCP and 17 for UDP.
	ProtocolNum int64 `json:"protocol_num,omitempty"`
}

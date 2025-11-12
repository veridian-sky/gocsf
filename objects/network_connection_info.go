package objects

// Code generated from OCSF schema; DO NOT EDIT.

// NetworkConnectionInfo The Network Connection Information object describes characteristics of an OSI
// Transport Layer communication, including TCP and UDP.
type NetworkConnectionInfo struct {
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

	// The Community ID of the network connection.
	CommunityUID string `json:"community_uid,omitempty"`

	// The direction of the initiated connection, traffic, or email, normalized to
	// the caption of the direction_id value. In the case of 'Other', it is defined
	// by the event source.
	Direction string `json:"direction,omitempty"`

	// The normalized identifier of the direction of the initiated connection,
	// traffic, or email.
	DirectionID int `json:"direction_id,omitempty"`

	// The Connection Flag History summarizes events in a network connection. For
	// example flags ShAD representing SYN, SYN/ACK, ACK and Data exchange.
	FlagHistory string `json:"flag_history,omitempty"`

	// The IP protocol name in lowercase, as defined by the Internet Assigned
	// Numbers Authority (IANA). For example: tcp or udp.
	ProtocolName string `json:"protocol_name,omitempty"`

	// The IP protocol number, as defined by the Internet Assigned Numbers
	// Authority (IANA). For example: 6 for TCP and 17 for UDP.
	ProtocolNum int64 `json:"protocol_num,omitempty"`

	// The Internet Protocol version.
	ProtocolVer string `json:"protocol_ver,omitempty"`

	// The Internet Protocol version identifier.
	ProtocolVerID int `json:"protocol_ver_id,omitempty"`

	// The authenticated user or service session.
	Session *Session `json:"session,omitempty"`

	// The network connection TCP header flags (i.e., control bits).
	TcpFlags int64 `json:"tcp_flags,omitempty"`

	// The unique identifier of the connection.
	Uid string `json:"uid,omitempty"`
}

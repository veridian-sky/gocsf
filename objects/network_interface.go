package objects

// Code generated from OCSF schema; DO NOT EDIT.

// NetworkInterface The Network Interface object describes the type and associated attributes of a
// physical or virtual network interface.
type NetworkInterface struct {
	// The hostname associated with the network interface.
	Hostname string `json:"hostname,omitempty"`

	// The IP address associated with the network interface.
	Ip string `json:"ip,omitempty"`

	// The MAC address of the network interface.
	Mac string `json:"mac,omitempty"`

	// The name of the network interface.
	Name string `json:"name,omitempty"`

	// The namespace is useful in merger or acquisition situations. For example,
	// when similar entities exist that you need to keep separate.
	Namespace string `json:"namespace,omitempty"`

	// The list of open ports on a network interface, including port numbers and
	// associated protocol information.
	OpenPorts []*PortInfo `json:"open_ports,omitempty"`

	// The subnet prefix length determines the number of bits used to represent the
	// network part of the IP address. The remaining bits are reserved for
	// identifying individual hosts within that subnet.
	SubnetPrefix int64 `json:"subnet_prefix,omitempty"`

	// The type of network interface.
	Type string `json:"type,omitempty"`

	// The network interface type identifier.
	TypeID int `json:"type_id,omitempty"`

	// The unique identifier for the network interface.
	Uid string `json:"uid,omitempty"`
}

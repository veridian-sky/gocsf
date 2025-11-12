package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Endpoint The Endpoint object describes a physical or virtual device that connects to and
// exchanges information with a computer network. Some examples of endpoints are
// mobile devices, desktop computers, virtual machines, embedded devices, and
// servers. Internet-of-Things devices—like cameras, lighting, refrigerators,
// security systems, smart speakers, and thermostats—are also endpoints.
type Endpoint struct {
	// A list of agent objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`

	// The information describing an instance of a container. A container is a
	// prepackaged, portable system image that runs isolated on an existing system
	// using a container runtime like containerd.
	Container *Container `json:"container,omitempty"`

	// The name of the domain that the endpoint belongs to or that corresponds to
	// the endpoint.
	Domain string `json:"domain,omitempty"`

	// The fully qualified name of the endpoint.
	Hostname string `json:"hostname,omitempty"`

	// The endpoint hardware information.
	HwInfo *DeviceHwInfo `json:"hw_info,omitempty"`

	// The unique identifier of a VM instance.
	InstanceUID string `json:"instance_uid,omitempty"`

	// The name of the network interface (e.g. eth2).
	InterfaceName string `json:"interface_name,omitempty"`

	// The unique identifier of the network interface.
	InterfaceUID string `json:"interface_uid,omitempty"`

	// The IP address of the endpoint, in either IPv4 or IPv6 format.
	Ip string `json:"ip,omitempty"`

	// The geographical location of the endpoint.
	Location *Location `json:"location,omitempty"`

	// The Media Access Control (MAC) address of the endpoint.
	Mac string `json:"mac,omitempty"`

	// The short name of the endpoint.
	Name string `json:"name,omitempty"`

	// If running under a process namespace (such as in a container), the process
	// identifier within that process namespace.
	NamespacePid int64 `json:"namespace_pid,omitempty"`

	// The endpoint operating system.
	Os *Os `json:"os,omitempty"`

	// The identity of the service or user account that owns the endpoint or was
	// last logged into it.
	Owner *User `json:"owner,omitempty"`

	// The unique identifier of a virtual subnet.
	SubnetUID string `json:"subnet_uid,omitempty"`

	// The endpoint type. For example: unknown, server, desktop, laptop, tablet,
	// mobile, virtual, browser, or other.
	Type string `json:"type,omitempty"`

	// The endpoint type ID.
	TypeID int `json:"type_id,omitempty"`

	// The unique identifier of the endpoint.
	Uid string `json:"uid,omitempty"`

	// The Virtual LAN identifier.
	VlanUID string `json:"vlan_uid,omitempty"`

	// The unique identifier of the Virtual Private Cloud (VPC).
	VpcUID string `json:"vpc_uid,omitempty"`

	// The network zone or LAN segment.
	Zone string `json:"zone,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// NetworkProxy The network proxy endpoint object describes a proxy server, which acts as an
// intermediary between a client requesting a resource and the server providing
// that resource.
type NetworkProxy struct {
	// A list of agent objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`

	// The Autonomous System details associated with an IP address.
	AutonomousSystem *AutonomousSystem `json:"autonomous_system,omitempty"`

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

	// The intermediate IP Addresses. For example, the IP addresses in the HTTP
	// X-Forwarded-For header.
	IntermediateIPs []string `json:"intermediate_ips,omitempty"`

	// The IP address of the endpoint, in either IPv4 or IPv6 format.
	Ip string `json:"ip,omitempty"`

	// The name of the Internet Service Provider (ISP).
	Isp string `json:"isp,omitempty"`

	// The organization name of the Internet Service Provider (ISP). This
	// represents the parent organization or company that owns/operates the ISP.
	// For example, Comcast Corporation would be the ISP org for Xfinity internet
	// service. This attribute helps identify the ultimate provider when ISPs
	// operate under different brand names.
	IspOrg string `json:"isp_org,omitempty"`

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

	// The port used for communication within the network connection.
	Port interface{} `json:"port,omitempty"`

	// The network proxy information pertaining to a specific endpoint. This can be
	// used to describe information pertaining to network address translation
	// (NAT).
	ProxyEndpoint *NetworkProxy `json:"proxy_endpoint,omitempty"`

	// The unique identifier of a virtual subnet.
	SubnetUID string `json:"subnet_uid,omitempty"`

	// The service name in service-to-service connections. For example, AWS VPC
	// logs the pkt-src-aws-service and pkt-dst-aws-service fields identify the
	// connection is coming from or going to an AWS service.
	SvcName string `json:"svc_name,omitempty"`

	// The network endpoint type. For example: unknown, server, desktop, laptop,
	// tablet, mobile, virtual, browser, or other.
	Type string `json:"type,omitempty"`

	// The network endpoint type ID.
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

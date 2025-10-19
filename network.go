// Code generated from OCSF schema. DO NOT EDIT.
package gocsf

import (
	"encoding/json"
	"time"
)

// Network Types

// Api The API, or Application Programming Interface, object represents  information pertaining to an API request and response.
type Api struct {
	// The information pertaining to the API group.
	Group *Group `json:"group,omitempty"`
	// Verb/Operation associated with the request
	Operation string `json:"operation"`
	// Details pertaining to the API request.
	Request *Request `json:"request"`
	// Details pertaining to the API response.
	Response *Response `json:"response"`
	// The information pertaining to the API service.
	Service *Service `json:"service,omitempty"`
	// The version of the API service.
	Version string `json:"version,omitempty"`
}

// AutonomousSystem An autonomous system (AS) is a collection of connected Internet Protocol (IP) routing prefixes under the control of one or more network operators on behalf of a single administrative entity or domain that presents a common, clearly defined routing policy to the internet.
type AutonomousSystem struct {
	// Organization name for the Autonomous System.
	Name string `json:"name"`
	// Unique number that the AS is identified by.
	Number int `json:"number"`
}

// DceRpc The DCE/RPC, or Distributed Computing Environment/Remote Procedure Call, object describes the remote procedure call system for distributed computing environments. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:RemoteProcedureCall/'>d3f:RemoteProcedureCall</a>.
type DceRpc struct {
	// The request command (e.g. REQUEST, BIND).
	Command string `json:"command"`
	// The reply to the request command (e.g. RESPONSE, BINDACK or FAULT).
	CommandResponse string `json:"command_response"`
	// The list of interface flags.
	Flags []string `json:"flags"`
	// An operation number used to identify a specific remote procedure call (RPC) method or a method in an interface.
	Opnum int `json:"opnum"`
	// The RPC Interface object describes the details pertaining to the remote procedure call interface.
	RpcInterface *RpcInterface `json:"rpc_interface"`
}

// DnsAnswer The DNS Answer object represents a specific response provided by the Domain Name System (DNS) when querying for information about a domain or performing a DNS operation. It encapsulates the relevant details and data returned by the DNS server in response to a query.
type DnsAnswer struct {
	// The class of DNS data contained in this resource record. See <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc1035.txt'>RFC1035</a>. For example: <code>IN</code>.
	Class string `json:"class"`
	// The list of DNS answer header flag IDs.
	FlagIds []int `json:"flag_ids"`
	// The list of DNS answer header flags.
	Flags []string `json:"flags,omitempty"`
	// The DNS packet identifier assigned by the program that generated the query. The identifier is copied to the response.
	PacketUid int `json:"packet_uid"`
	// The data describing the DNS resource. The meaning of this data depends on the type and class of the resource record.
	Rdata string `json:"rdata"`
	// The time interval that the resource record may be cached. Zero value means that the resource record can only be used for the transaction in progress, and should not be cached.
	Ttl int `json:"ttl"`
	// The type of data contained in this resource record. See <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc1035.txt'>RFC1035</a>. For example: <code>CNAME</code>.
	Type string `json:"type"`
}

// DnsQuery The DNS query object represents a specific request made to the Domain Name System (DNS) to retrieve information about a domain or perform a DNS operation. This object encapsulates the necessary attributes and methods to construct and send DNS queries, specify the query type (e.g., A, AAAA, MX). Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:DNSLookup/'>d3f:DNSLookup</a>.
type DnsQuery struct {
	// The class of resource records being queried. See <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc1035.txt'>RFC1035</a>. For example: <code>IN</code>.
	Class string `json:"class"`
	// The hostname or domain being queried. For example: <code>www.example.com</code>
	Hostname string `json:"hostname"`
	// The DNS opcode specifies the type of the query message.
	Opcode string `json:"opcode,omitempty"`
	// The DNS opcode ID specifies the normalized query message type as defined in <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc5395.html'>RFC-5395</a>.
	OpcodeId int `json:"opcode_id"`
	// The DNS packet identifier assigned by the program that generated the query. The identifier is copied to the response.
	PacketUid int `json:"packet_uid"`
	// The type of resource records being queried. See <a target='_blank' href='https://www.rfc-editor.org/rfc/rfc1035.txt'>RFC1035</a>. For example: A, AAAA, CNAME, MX, and NS.
	Type string `json:"type"`
}

// DomainContact The contact information related to a domain registration, e.g., registrant, administrator, abuse, billing, or technical contact.
type DomainContact struct {
	// The user's primary email address.
	EmailAddr string `json:"email_addr"`
	// Location details for the contract such as the city, state/province, country, etc.
	Location *Location `json:"location"`
	// The individual or organization name for the contact.
	Name string `json:"name,omitempty"`
	// The number associated with the phone.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The Domain Contact type, normalized to the caption of the <code>type_id</code> value. In the case of 'Other', it is defined by the source
	Type string `json:"type,omitempty"`
	// The normalized domain contact type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the contact information, typically provided in WHOIS information.
	Uid string `json:"uid,omitempty"`
}

// Endpoint The Endpoint object describes a physical or virtual device that connects to and exchanges information with a computer network. Some examples of endpoints are mobile devices, desktop computers, virtual machines, embedded devices, and servers. Internet-of-Things devices—like cameras, lighting, refrigerators, security systems, smart speakers, and thermostats—are also endpoints.
type Endpoint struct {
	// A list of <code>agent</code> objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`
	// The information describing an instance of a container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
	Container *Container `json:"container"`
	// The name of the domain.
	Domain string `json:"domain,omitempty"`
	// The fully qualified name of the endpoint.
	Hostname string `json:"hostname"`
	// The endpoint hardware information.
	HwInfo *DeviceHwInfo `json:"hw_info,omitempty"`
	// The unique identifier of a VM instance.
	InstanceUid string `json:"instance_uid"`
	// The name of the network interface (e.g. eth2).
	InterfaceName string `json:"interface_name"`
	// The unique identifier of the network interface.
	InterfaceUid string `json:"interface_uid"`
	// The IP address of the endpoint, in either IPv4 or IPv6 format.
	Ip string `json:"ip"`
	// The geographical location of the endpoint.
	Location *Location `json:"location,omitempty"`
	// The Media Access Control (MAC) address of the endpoint.
	Mac string `json:"mac,omitempty"`
	// The short name of the endpoint.
	Name string `json:"name"`
	// If running under a process namespace (such as in a container), the process identifier within that process namespace.
	NamespacePid int `json:"namespace_pid"`
	// The endpoint operating system.
	Os *Os `json:"os,omitempty"`
	// The identity of the service or user account that owns the endpoint or was last logged into it.
	Owner *User `json:"owner"`
	// The unique identifier of a virtual subnet.
	SubnetUid string `json:"subnet_uid,omitempty"`
	// The endpoint type. For example: <code>unknown</code>, <code>server</code>, <code>desktop</code>, <code>laptop</code>, <code>tablet</code>, <code>mobile</code>, <code>virtual</code>, <code>browser</code>, or <code>other</code>.
	Type string `json:"type,omitempty"`
	// The endpoint type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the endpoint.
	Uid string `json:"uid"`
	// The Virtual LAN identifier.
	VlanUid string `json:"vlan_uid,omitempty"`
	// The unique identifier of the Virtual Private Cloud (VPC).
	VpcUid string `json:"vpc_uid,omitempty"`
	// The network zone or LAN segment.
	Zone string `json:"zone,omitempty"`
}

// EndpointConnection The Endpoint Connection object contains information detailing a connection attempt to an endpoint.
type EndpointConnection struct {
	// A numerical response status code providing details about the connection.
	Code int `json:"code"`
	// Provides characteristics of the network endpoint.
	NetworkEndpoint *NetworkEndpoint `json:"network_endpoint"`
}

// HttpCookie The HTTP Cookie object, also known as a web cookie or browser cookie, contains details and values pertaining to a small piece of data that a server sends to a user's web browser. This data is then stored by the browser and sent back to the server with subsequent requests, allowing the server to remember and track certain information about the user's browsing session or preferences.
type HttpCookie struct {
	// The name of the domain.
	Domain string `json:"domain,omitempty"`
	// The expiration time of the HTTP cookie.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// The expiration time of the HTTP cookie.
	ExpirationTimeDt time.Time `json:"expiration_time_dt,omitempty"`
	// A cookie attribute to make it inaccessible via JavaScript
	HttpOnly bool `json:"http_only,omitempty"`
	// This attribute prevents the cookie from being accessed via JavaScript.
	IsHttpOnly bool `json:"is_http_only,omitempty"`
	// The cookie attribute indicates that cookies are sent to the server only when the request is encrypted using the HTTPS protocol.
	IsSecure bool `json:"is_secure,omitempty"`
	// The HTTP cookie name.
	Name string `json:"name"`
	// The path of the HTTP cookie.
	Path string `json:"path,omitempty"`
	// The cookie attribute that lets servers specify whether/when cookies are sent with cross-site requests. Values are: Strict, Lax or None
	Samesite string `json:"samesite,omitempty"`
	// The cookie attribute to only send cookies to the server with an encrypted request over the HTTPS protocol.
	Secure bool `json:"secure,omitempty"`
	// The HTTP cookie value.
	Value string `json:"value"`
}

// HttpHeader TThe HTTP Header object represents the headers sent in an HTTP request or response. HTTP headers are key-value pairs that convey additional information about the HTTP message, including details about the content, caching, authentication, encoding, and other aspects of the communication.
type HttpHeader struct {
	// The name of the header
	Name string `json:"name"`
	// The value of the header
	Value string `json:"value"`
}

// HttpRequest The HTTP Request object represents the attributes of a request made to a web server. It encapsulates the details and metadata associated with an HTTP request, including the request method, headers, URL, query parameters, body content, and other relevant information.
type HttpRequest struct {
	// The arguments sent along with the HTTP request.
	Args string `json:"args,omitempty"`
	// Additional HTTP headers of an HTTP request or response.
	HttpHeaders []*HttpHeader `json:"http_headers"`
	// The <a target='_blank' href='https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods'>HTTP request method</a> indicates the desired action to be performed for a given resource.
	HttpMethod string `json:"http_method"`
	// The HTTP request length, in number of bytes.
	Length int `json:"length,omitempty"`
	// The request header that identifies the address of the previous web page, which is linked to the current web page or resource being requested.
	Referrer string `json:"referrer,omitempty"`
	// The unique identifier of the http request.
	Uid string `json:"uid,omitempty"`
	// The URL object that pertains to the request.
	Url *Url `json:"url"`
	// The request header that identifies the operating system and web browser.
	UserAgent string `json:"user_agent"`
	// The Hypertext Transfer Protocol (HTTP) version.
	Version string `json:"version"`
	// The X-Forwarded-For header identifying the originating IP address(es) of a client connecting to a web server through an HTTP proxy or a load balancer.
	XForwardedFor []string `json:"x_forwarded_for,omitempty"`
}

// HttpResponse The HTTP Response object contains detailed information about the response sent from a web server to the requester. It encompasses attributes and metadata that describe the response status, headers, body content, and other relevant information.
type HttpResponse struct {
	// The Hypertext Transfer Protocol (HTTP) status code returned from the web server to the client. For example, 200.
	Code int `json:"code"`
	// The request header that identifies the original <a target='_blank' href='https://www.iana.org/assignments/media-types/media-types.xhtml'>media type </a> of the resource (prior to any content encoding applied for sending).
	ContentType string `json:"content_type,omitempty"`
	// Additional HTTP headers of an HTTP request or response.
	HttpHeaders []*HttpHeader `json:"http_headers"`
	// The HTTP response latency measured in milliseconds.
	Latency int `json:"latency,omitempty"`
	// The HTTP response length, in number of bytes.
	Length int `json:"length,omitempty"`
	// The description of the event/finding, as defined by the source.
	Message string `json:"message,omitempty"`
	// The response status. For example: A successful HTTP status of 'OK' which corresponds to a code of 200.
	Status string `json:"status,omitempty"`
}

// NetworkConnectionInfo The Network Connection Information object describes characteristics of a network connection. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:NetworkSession/'>d3f:NetworkSession</a>.
type NetworkConnectionInfo struct {
	// The boundary of the connection, normalized to the caption of 'boundary_id'. In the case of 'Other', it is defined by the event source. <p> For cloud connections, this translates to the traffic-boundary(same VPC, through IGW, etc.). For traditional networks, this is described as Local, Internal, or External.</p>
	Boundary string `json:"boundary,omitempty"`
	// <p>The normalized identifier of the boundary of the connection. </p><p> For cloud connections, this translates to the traffic-boundary (same VPC, through IGW, etc.). For traditional networks, this is described as Local, Internal, or External.</p>
	BoundaryId int `json:"boundary_id"`
	// The direction of the initiated connection, traffic, or email, normalized to the caption of the direction_id value. In the case of 'Other', it is defined by the event source.
	Direction string `json:"direction,omitempty"`
	// The normalized identifier of the direction of the initiated connection, traffic, or email.
	DirectionId int `json:"direction_id"`
	// The TCP/IP protocol name in lowercase, as defined by the Internet Assigned Numbers Authority (IANA). See <a target='_blank' href='https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml'>Protocol Numbers</a>. For example: <code>tcp</code> or <code>udp</code>.
	ProtocolName string `json:"protocol_name"`
	// The TCP/IP protocol number, as defined by the Internet Assigned Numbers Authority (IANA). Use -1 if the protocol is not defined by IANA. See <a target='_blank' href='https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml'>Protocol Numbers</a>. For example: <code>6</code> for TCP and <code>17</code> for UDP.
	ProtocolNum int `json:"protocol_num"`
	// The Internet Protocol version.
	ProtocolVer string `json:"protocol_ver,omitempty"`
	// The Internet Protocol version identifier.
	ProtocolVerId int `json:"protocol_ver_id"`
	// The authenticated user or service session.
	Session *Session `json:"session,omitempty"`
	// The network connection TCP header flags (i.e., control bits).
	TcpFlags int `json:"tcp_flags,omitempty"`
	// The unique identifier of the connection.
	Uid string `json:"uid"`
}

// NetworkEndpoint The Network Endpoint object describes characteristics of a network endpoint. These can be a source or destination of a network connection.
type NetworkEndpoint struct {
	// A list of <code>agent</code> objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`
	// The Autonomous System details associated with an IP address.
	AutonomousSystem *AutonomousSystem `json:"autonomous_system,omitempty"`
	// The information describing an instance of a container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
	Container *Container `json:"container"`
	// The name of the domain.
	Domain string `json:"domain,omitempty"`
	// The fully qualified name of the endpoint.
	Hostname string `json:"hostname"`
	// The endpoint hardware information.
	HwInfo *DeviceHwInfo `json:"hw_info,omitempty"`
	// The unique identifier of a VM instance.
	InstanceUid string `json:"instance_uid"`
	// The name of the network interface (e.g. eth2).
	InterfaceName string `json:"interface_name"`
	// The unique identifier of the network interface.
	InterfaceUid string `json:"interface_uid"`
	// The intermediate IP Addresses. For example, the IP addresses in the HTTP X-Forwarded-For header.
	IntermediateIps []string `json:"intermediate_ips,omitempty"`
	// The IP address of the endpoint, in either IPv4 or IPv6 format.
	Ip string `json:"ip"`
	// The geographical location of the endpoint.
	Location *Location `json:"location,omitempty"`
	// The Media Access Control (MAC) address of the endpoint.
	Mac string `json:"mac,omitempty"`
	// The short name of the endpoint.
	Name string `json:"name"`
	// If running under a process namespace (such as in a container), the process identifier within that process namespace.
	NamespacePid int `json:"namespace_pid"`
	// The endpoint operating system.
	Os *Os `json:"os,omitempty"`
	// The identity of the service or user account that owns the endpoint or was last logged into it.
	Owner *User `json:"owner"`
	// The port used for communication within the network connection.
	Port int `json:"port"`
	// The network proxy information pertaining to a specific endpoint. This can be used to describe information pertaining to network address translation (NAT).
	ProxyEndpoint *NetworkProxy `json:"proxy_endpoint,omitempty"`
	// The unique identifier of a virtual subnet.
	SubnetUid string `json:"subnet_uid,omitempty"`
	// The service name in service-to-service connections. For example, AWS VPC logs the pkt-src-aws-service and pkt-dst-aws-service fields identify the connection is coming from or going to an AWS service.
	SvcName string `json:"svc_name"`
	// The network endpoint type. For example: <code>unknown</code>, <code>server</code>, <code>desktop</code>, <code>laptop</code>, <code>tablet</code>, <code>mobile</code>, <code>virtual</code>, <code>browser</code>, or <code>other</code>.
	Type string `json:"type,omitempty"`
	// The network endpoint type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the endpoint.
	Uid string `json:"uid"`
	// The Virtual LAN identifier.
	VlanUid string `json:"vlan_uid,omitempty"`
	// The unique identifier of the Virtual Private Cloud (VPC).
	VpcUid string `json:"vpc_uid,omitempty"`
	// The network zone or LAN segment.
	Zone string `json:"zone,omitempty"`
}

// NetworkInterface The Network Interface object describes the type and associated attributes of a network interface.
type NetworkInterface struct {
	// The hostname associated with the network interface.
	Hostname string `json:"hostname"`
	// The IP address associated with the network interface.
	Ip string `json:"ip"`
	// The MAC address of the network interface.
	Mac string `json:"mac"`
	// The name of the network interface.
	Name string `json:"name"`
	// The namespace is useful in merger or acquisition situations. For example, when similar entities exist that you need to keep separate.
	Namespace string `json:"namespace,omitempty"`
	// The subnet prefix length determines the number of bits used to represent the network part of the IP address. The remaining bits are reserved for identifying individual hosts within that subnet.
	SubnetPrefix int `json:"subnet_prefix,omitempty"`
	// The type of network interface.
	Type string `json:"type,omitempty"`
	// The network interface type identifier.
	TypeId int `json:"type_id"`
	// The unique identifier for the network interface.
	Uid string `json:"uid,omitempty"`
}

// NetworkProxy The network proxy endpoint object describes a proxy server, which acts as an intermediary between a client requesting a resource and the server providing that resource.  Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:ProxyServer/'>d3f:ProxyServer</a>.
type NetworkProxy struct {
	// A list of <code>agent</code> objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`
	// The Autonomous System details associated with an IP address.
	AutonomousSystem *AutonomousSystem `json:"autonomous_system,omitempty"`
	// The information describing an instance of a container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
	Container *Container `json:"container"`
	// The name of the domain.
	Domain string `json:"domain,omitempty"`
	// The fully qualified name of the endpoint.
	Hostname string `json:"hostname"`
	// The endpoint hardware information.
	HwInfo *DeviceHwInfo `json:"hw_info,omitempty"`
	// The unique identifier of a VM instance.
	InstanceUid string `json:"instance_uid"`
	// The name of the network interface (e.g. eth2).
	InterfaceName string `json:"interface_name"`
	// The unique identifier of the network interface.
	InterfaceUid string `json:"interface_uid"`
	// The intermediate IP Addresses. For example, the IP addresses in the HTTP X-Forwarded-For header.
	IntermediateIps []string `json:"intermediate_ips,omitempty"`
	// The IP address of the endpoint, in either IPv4 or IPv6 format.
	Ip string `json:"ip"`
	// The geographical location of the endpoint.
	Location *Location `json:"location,omitempty"`
	// The Media Access Control (MAC) address of the endpoint.
	Mac string `json:"mac,omitempty"`
	// The short name of the endpoint.
	Name string `json:"name"`
	// If running under a process namespace (such as in a container), the process identifier within that process namespace.
	NamespacePid int `json:"namespace_pid"`
	// The endpoint operating system.
	Os *Os `json:"os,omitempty"`
	// The identity of the service or user account that owns the endpoint or was last logged into it.
	Owner *User `json:"owner"`
	// The port used for communication within the network connection.
	Port int `json:"port"`
	// The network proxy information pertaining to a specific endpoint. This can be used to describe information pertaining to network address translation (NAT).
	ProxyEndpoint *NetworkProxy `json:"proxy_endpoint,omitempty"`
	// The unique identifier of a virtual subnet.
	SubnetUid string `json:"subnet_uid,omitempty"`
	// The service name in service-to-service connections. For example, AWS VPC logs the pkt-src-aws-service and pkt-dst-aws-service fields identify the connection is coming from or going to an AWS service.
	SvcName string `json:"svc_name"`
	// The network endpoint type. For example: <code>unknown</code>, <code>server</code>, <code>desktop</code>, <code>laptop</code>, <code>tablet</code>, <code>mobile</code>, <code>virtual</code>, <code>browser</code>, or <code>other</code>.
	Type string `json:"type,omitempty"`
	// The network endpoint type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the endpoint.
	Uid string `json:"uid"`
	// The Virtual LAN identifier.
	VlanUid string `json:"vlan_uid,omitempty"`
	// The unique identifier of the Virtual Private Cloud (VPC).
	VpcUid string `json:"vpc_uid,omitempty"`
	// The network zone or LAN segment.
	Zone string `json:"zone,omitempty"`
}

// NetworkTraffic The Network Traffic object describes characteristics of network traffic. Network traffic refers to data moving across a network at a given point of time. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:NetworkTraffic/'>d3f:NetworkTraffic</a>.
type NetworkTraffic struct {
	// The total number of bytes (in and out).
	Bytes int64 `json:"bytes"`
	// The number of bytes sent from the destination to the source.
	BytesIn int64 `json:"bytes_in,omitempty"`
	// The number of bytes sent from the source to the destination.
	BytesOut int64 `json:"bytes_out,omitempty"`
	// The total number of chunks (in and out).
	Chunks int64 `json:"chunks,omitempty"`
	// The number of chunks sent from the destination to the source.
	ChunksIn int64 `json:"chunks_in,omitempty"`
	// The number of chunks sent from the source to the destination.
	ChunksOut int64 `json:"chunks_out,omitempty"`
	// The total number of packets (in and out).
	Packets int64 `json:"packets"`
	// The number of packets sent from the destination to the source.
	PacketsIn int64 `json:"packets_in,omitempty"`
	// The number of packets sent from the source to the destination.
	PacketsOut int64 `json:"packets_out,omitempty"`
}

// RpcInterface The RPC Interface represents the remote procedure call interface used in the DCE/RPC session.
type RpcInterface struct {
	// An integer that provides a reason code or additional information about the acknowledgment result.
	AckReason int `json:"ack_reason"`
	// An integer that denotes the acknowledgment result of the DCE/RPC call.
	AckResult int `json:"ack_result"`
	// The unique identifier of the particular remote procedure or service.
	Uuid string `json:"uuid"`
	// The version of the DCE/RPC protocol being used in the session.
	Version string `json:"version"`
}

// Tls The Transport Layer Security (TLS) object describes the negotiated TLS protocol used for secure communications over an establish network connection.
type Tls struct {
	// The integer value of TLS alert if present. The alerts are defined in the TLS specification in <a target='_blank' href='https://datatracker.ietf.org/doc/html/rfc2246'>RFC-2246</a>.
	Alert int `json:"alert,omitempty"`
	// The certificate object containing information about the digital certificate.
	Certificate *Certificate `json:"certificate"`
	// The Chain of Certificate Serial Numbers field provides a chain of Certificate Issuer Serial Numbers leading to the Root Certificate Issuer.
	CertificateChain []string `json:"certificate_chain"`
	// The negotiated cipher suite.
	Cipher string `json:"cipher"`
	// The client cipher suites that were exchanged during the TLS handshake negotiation.
	ClientCiphers []string `json:"client_ciphers"`
	// The list of TLS extensions.
	ExtensionList []*TlsExtension `json:"extension_list,omitempty"`
	// The amount of total time for the TLS handshake to complete after the TCP connection is established, including client-side delays, in milliseconds.
	HandshakeDur int `json:"handshake_dur,omitempty"`
	// The MD5 hash of a JA3 string.
	Ja3Hash *Fingerprint `json:"ja3_hash"`
	// The MD5 hash of a JA3S string.
	Ja3sHash *Fingerprint `json:"ja3s_hash"`
	// The length of the encryption key.
	KeyLength int `json:"key_length,omitempty"`
	// The list of subject alternative names that are secured by a specific certificate.
	Sans []*San `json:"sans,omitempty"`
	// The server cipher suites that were exchanged during the TLS handshake negotiation.
	ServerCiphers []string `json:"server_ciphers,omitempty"`
	//  The Server Name Indication (SNI) extension sent by the client.
	Sni string `json:"sni"`
	// The list of TLS extensions.
	TlsExtensionList []*TlsExtension `json:"tls_extension_list,omitempty"`
	// The TLS protocol version.
	Version string `json:"version"`
}

// TlsExtension The TLS Extension object describes additional attributes that extend the base Transport Layer Security (TLS) object.
type TlsExtension struct {
	// The data contains information specific to the particular extension type.
	Data json.RawMessage `json:"data"`
	// The TLS extension type. For example: <code>Server Name</code>.
	Type string `json:"type,omitempty"`
	// The TLS extension type identifier. See <a target='_blank' href='https://datatracker.ietf.org/doc/html/rfc8446#page-35'>The Transport Layer Security (TLS) extension page</a>.
	TypeId int `json:"type_id"`
}

// Url The Uniform Resource Locator(URL) object describes the characteristics of a URL. Defined in <a target='_blank' href='https://datatracker.ietf.org/doc/html/rfc1738'>RFC 1738</a> and by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:URL/'>d3f:URL</a>.
type Url struct {
	// The Website categorization names, as defined by <code>category_ids</code> enum values.
	Categories []string `json:"categories,omitempty"`
	// The Website categorization identifiers.
	CategoryIds []int `json:"category_ids"`
	// The domain portion of the URL. For example: <code>example.com</code> in <code>https://sub.example.com</code>.
	Domain string `json:"domain,omitempty"`
	// The URL host as extracted from the URL. For example: <code>www.example.com</code> from <code>www.example.com/download/trouble</code>.
	Hostname string `json:"hostname"`
	// The URL path as extracted from the URL. For example: <code>/download/trouble</code> from <code>www.example.com/download/trouble</code>.
	Path string `json:"path"`
	// The URL port. For example: <code>80</code>.
	Port int `json:"port"`
	// The query portion of the URL. For example: the query portion of the URL <code>http://www.example.com/search?q=bad&sort=date</code> is <code>q=bad&sort=date</code>.
	QueryString string `json:"query_string"`
	// The context in which a resource was retrieved in a web request.
	ResourceType string `json:"resource_type,omitempty"`
	// The scheme portion of the URL. For example: <code>http</code>, <code>https</code>, <code>ftp</code>, or <code>sftp</code>.
	Scheme string `json:"scheme"`
	// The subdomain portion of the URL. For example: <code>sub</code> in <code>https://sub.example.com</code> or <code>sub2.sub1</code> in <code>https://sub2.sub1.example.com</code>.
	Subdomain string `json:"subdomain,omitempty"`
	// The URL string. See RFC 1738. For example: <code>http://www.example.com/download/trouble.exe</code>. Note: The URL path should not populate the URL string.
	UrlString string `json:"url_string"`
}


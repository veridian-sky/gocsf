// Code generated from OCSF schema. DO NOT EDIT.
package gocsf

import (
	"encoding/json"
	"time"
)

// System Types

// Agent An Agent (also known as a Sensor) is typically installed on an Operating System (OS) and serves as a specialized software component that can be designed to monitor, detect, collect, archive, or take action. These activities and possible actions are defined by the upstream system controlling the Agent and its intended purpose. For instance, an Agent can include Endpoint Detection & Response (EDR) agents, backup/disaster recovery sensors, Application Performance Monitoring or profiling sensors, and similar software.
type Agent struct {
	// The name of the agent or sensor. For example: <code>AWS SSM Agent</code>.
	Name string `json:"name"`
	// Describes the various policies that may be applied or enforced by an agent or sensor. E.g., Conditional Access, prevention, auto-update, tamper protection, destination configuration, etc.
	Policies []*Policy `json:"policies,omitempty"`
	// The normalized caption of the type_id value for the agent or sensor. In the case of 'Other' or 'Unknown', it is defined by the event source.
	Type string `json:"type,omitempty"`
	// The normalized representation of an agent or sensor. E.g., EDR, vulnerability management, APM, backup & recovery, etc.
	TypeId int `json:"type_id"`
	// The UID of the agent or sensor, sometimes known as a Sensor ID or <code>aid</code>.
	Uid string `json:"uid"`
	// An alternative or contextual identifier for the agent or sensor, such as a configuration, organization, or license UID.
	UidAlt string `json:"uid_alt,omitempty"`
	// The company or author who created the agent or sensor. For example: <code>Crowdstrike</code>.
	VendorName string `json:"vendor_name,omitempty"`
	// The semantic version of the agent or sensor, e.g., <code>7.101.50.0</code>.
	Version string `json:"version,omitempty"`
}

// Cloud The Cloud object contains information about a cloud account such as AWS Account ID, regions, etc.
type Cloud struct {
	// The account object describes details about the account that was the source or target of the activity.
	Account *Account `json:"account,omitempty"`
	// Organization and org unit relevant to the event or object.
	Org *Organization `json:"org,omitempty"`
	// The unique identifier of a Cloud project.
	ProjectUid string `json:"project_uid,omitempty"`
	// The unique name of the Cloud services provider, such as AWS, MS Azure, GCP, etc.
	Provider string `json:"provider"`
	// The name of the cloud region, as defined by the cloud provider.
	Region string `json:"region"`
	// The availability zone in the cloud region, as defined by the cloud provider.
	Zone string `json:"zone,omitempty"`
}

// Container The Container object describes an instance of a specific container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
type Container struct {
	// Commit hash of image created for docker or the SHA256 hash of the container. For example: <code>13550340a8681c84c861aac2e5b440161c2b33a3e4f302ac680ca5b686de48de</code>.
	Hash *Fingerprint `json:"hash"`
	// The container image used as a template to run the container.
	Image *Image `json:"image"`
	// The container name.
	Name string `json:"name"`
	// The network driver used by the container. For example, bridge, overlay, host, none, etc.
	NetworkDriver string `json:"network_driver,omitempty"`
	// The orchestrator managing the container, such as ECS, EKS, K8s, or OpenShift.
	Orchestrator string `json:"orchestrator,omitempty"`
	// The unique identifier of the pod (or equivalent) that the container is executing on.
	PodUuid string `json:"pod_uuid,omitempty"`
	// The backend running the container, such as containerd or cri-o.
	Runtime string `json:"runtime,omitempty"`
	// The size of the container image.
	Size int64 `json:"size"`
	// The tag used by the container. It can indicate version, format, OS.
	Tag string `json:"tag,omitempty"`
	// The full container unique identifier for this instantiation of the container. For example: <code>ac2ea168264a08f9aaca0dfc82ff3551418dfd22d02b713142a6843caa2f61bf</code>.
	Uid string `json:"uid"`
}

// Database The database object is used for databases which are typically datastore services that contain an organized collection of structured and unstructured data or a types of data.
type Database struct {
	// The time when the database was known to have been created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the database was known to have been created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The description of the database.
	Desc string `json:"desc,omitempty"`
	// The group names to which the database belongs.
	Groups []*Group `json:"groups,omitempty"`
	// The most recent time when any changes, updates, or modifications were made within the database.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The most recent time when any changes, updates, or modifications were made within the database.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The database name, ordinarily as assigned by a database administrator.
	Name string `json:"name"`
	// The size of the database in bytes.
	Size int64 `json:"size,omitempty"`
	// The database type.
	Type string `json:"type"`
	// The normalized identifier of the database type.
	TypeId int `json:"type_id"`
	// The unique identifier of the database.
	Uid string `json:"uid"`
}

// Databucket The databucket object is a basic container that holds data, typically organized through the use of data partitions.
type Databucket struct {
	// The time when the databucket was known to have been created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the databucket was known to have been created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The description of the databucket.
	Desc string `json:"desc,omitempty"`
	// A file within a databucket.
	File *File `json:"file,omitempty"`
	// The group names to which the databucket belongs.
	Groups []*Group `json:"groups,omitempty"`
	// The most recent time when any changes, updates, or modifications were made within the databucket.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The most recent time when any changes, updates, or modifications were made within the databucket.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The databucket name.
	Name string `json:"name"`
	// The size of the databucket in bytes.
	Size int64 `json:"size,omitempty"`
	// The databucket type.
	Type string `json:"type"`
	// The normalized identifier of the databucket type.
	TypeId int `json:"type_id"`
	// The unique identifier of the databucket.
	Uid string `json:"uid"`
}

// Device The Device object represents an addressable computer system or host, which is typically connected to a computer network and participates in the transmission or processing of data within the computer network. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Host/'>d3f:Host</a>.
type Device struct {
	// A list of <code>agent</code> objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`
	// The unique identifier of the cloud autoscale configuration.
	AutoscaleUid string `json:"autoscale_uid,omitempty"`
	// The time the system was booted.
	BootTime time.Time `json:"boot_time,omitempty"`
	// The time the system was booted.
	BootTimeDt time.Time `json:"boot_time_dt,omitempty"`
	// The information describing an instance of a container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
	Container *Container `json:"container"`
	// The time when the device was known to have been created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the device was known to have been created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The description of the device, ordinarily as reported by the operating system.
	Desc string `json:"desc,omitempty"`
	// The network domain where the device resides. For example: <code>work.example.com</code>.
	Domain string `json:"domain,omitempty"`
	// The initial discovery time of the device.
	FirstSeenTime time.Time `json:"first_seen_time,omitempty"`
	// The initial discovery time of the device.
	FirstSeenTimeDt time.Time `json:"first_seen_time_dt,omitempty"`
	// The group names to which the device belongs. For example: <code>["Windows Laptops", "Engineering"]<code/>.
	Groups []*Group `json:"groups,omitempty"`
	// The device hostname.
	Hostname string `json:"hostname"`
	// The endpoint hardware information.
	HwInfo *DeviceHwInfo `json:"hw_info,omitempty"`
	// The name of the hypervisor running on the device. For example, <code>Xen</code>, <code>VMware</code>, <code>Hyper-V</code>, <code>VirtualBox</code>, etc.
	Hypervisor string `json:"hypervisor,omitempty"`
	// The image used as a template to run the virtual machine.
	Image *Image `json:"image,omitempty"`
	// The International Mobile Station Equipment Identifier that is associated with the device.
	Imei string `json:"imei,omitempty"`
	// The unique identifier of a VM instance.
	InstanceUid string `json:"instance_uid"`
	// The name of the network interface (e.g. eth2).
	InterfaceName string `json:"interface_name"`
	// The unique identifier of the network interface.
	InterfaceUid string `json:"interface_uid"`
	// The device IP address, in either IPv4 or IPv6 format.
	Ip string `json:"ip,omitempty"`
	// The event occurred on a compliant device.
	IsCompliant bool `json:"is_compliant,omitempty"`
	// The event occurred on a managed device.
	IsManaged bool `json:"is_managed,omitempty"`
	// The event occurred on a personal device.
	IsPersonal bool `json:"is_personal,omitempty"`
	// The event occurred on a trusted device.
	IsTrusted bool `json:"is_trusted,omitempty"`
	// The most recent discovery time of the device.
	LastSeenTime time.Time `json:"last_seen_time,omitempty"`
	// The most recent discovery time of the device.
	LastSeenTimeDt time.Time `json:"last_seen_time_dt,omitempty"`
	// The geographical location of the device.
	Location *Location `json:"location,omitempty"`
	// The Media Access Control (MAC) address of the endpoint.
	Mac string `json:"mac,omitempty"`
	// The time when the device was last known to have been modified.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The time when the device was last known to have been modified.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The alternate device name, ordinarily as assigned by an administrator. <p><b>Note:</b> The <b>Name</b> could be any other string that helps to identify the device, such as a phone number; for example <code>310-555-1234</code>.</p>
	Name string `json:"name,omitempty"`
	// If running under a process namespace (such as in a container), the process identifier within that process namespace.
	NamespacePid int `json:"namespace_pid"`
	// The network interfaces that are associated with the device, one for each unique MAC address/IP address/hostname/name combination.<p><b>Note:</b> The first element of the array is the network information that pertains to the event.</p>
	NetworkInterfaces []*NetworkInterface `json:"network_interfaces,omitempty"`
	// Organization and org unit related to the device.
	Org *Organization `json:"org,omitempty"`
	// The endpoint operating system.
	Os *Os `json:"os,omitempty"`
	// The identity of the service or user account that owns the endpoint or was last logged into it.
	Owner *User `json:"owner"`
	// The region where the virtual machine is located. For example, an AWS Region.
	Region string `json:"region"`
	// The risk level, normalized to the caption of the risk_level_id value.
	RiskLevel string `json:"risk_level,omitempty"`
	// The normalized risk level id.
	RiskLevelId int `json:"risk_level_id,omitempty"`
	// The risk score as reported by the event source.
	RiskScore int `json:"risk_score,omitempty"`
	// The subnet mask.
	Subnet string `json:"subnet,omitempty"`
	// The unique identifier of a virtual subnet.
	SubnetUid string `json:"subnet_uid,omitempty"`
	// The device type. For example: <code>unknown</code>, <code>server</code>, <code>desktop</code>, <code>laptop</code>, <code>tablet</code>, <code>mobile</code>, <code>virtual</code>, <code>browser</code>, or <code>other</code>.
	Type string `json:"type"`
	// The device type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the device. For example the Windows TargetSID or AWS EC2 ARN.
	Uid string `json:"uid"`
	// An alternate unique identifier of the device if any. For example the ActiveDirectory DN.
	UidAlt string `json:"uid_alt,omitempty"`
	// The Virtual LAN identifier.
	VlanUid string `json:"vlan_uid,omitempty"`
	// The unique identifier of the Virtual Private Cloud (VPC).
	VpcUid string `json:"vpc_uid,omitempty"`
	// The network zone or LAN segment.
	Zone string `json:"zone,omitempty"`
}

// DeviceHwInfo The Device Hardware Information object contains details and specifications of the physical components that make up a device. This information provides an overview of the hardware capabilities, configuration, and characteristics of the device.
type DeviceHwInfo struct {
	// The BIOS date. For example: <code>03/31/16</code>.
	BiosDate string `json:"bios_date,omitempty"`
	// The BIOS manufacturer. For example: <code>LENOVO</code>.
	BiosManufacturer string `json:"bios_manufacturer,omitempty"`
	// The BIOS version. For example: <code>LENOVO G5ETA2WW (2.62)</code>.
	BiosVer string `json:"bios_ver,omitempty"`
	// The chassis type describes the system enclosure or physical form factor. Such as the following examples for Windows <a target='_blank' href='https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-systemenclosure'>Windows Chassis Types</a>
	Chassis string `json:"chassis,omitempty"`
	// The cpu architecture, the number of bits used for addressing in memory. For example: <code>32</code> or <code>64</code>.
	CpuBits int `json:"cpu_bits,omitempty"`
	// The number of processor cores in all installed processors. For Example: <code>42</code>.
	CpuCores int `json:"cpu_cores,omitempty"`
	// The number of physical processors on a system. For example: <code>1</code>.
	CpuCount int `json:"cpu_count,omitempty"`
	// The speed of the processor in Mhz. For Example: <code>4200</code>.
	CpuSpeed int `json:"cpu_speed,omitempty"`
	// The processor type. For example: <code>x86 Family 6 Model 37 Stepping 5</code>.
	CpuType string `json:"cpu_type,omitempty"`
	// The desktop display affiliated with the event
	DesktopDisplay *Display `json:"desktop_display,omitempty"`
	// The keyboard detailed information.
	KeyboardInfo *KeyboardInfo `json:"keyboard_info,omitempty"`
	// The total amount of installed RAM, in Megabytes. For example: <code>2048</code>.
	RamSize int `json:"ram_size,omitempty"`
	// The device manufacturer serial number.
	SerialNumber string `json:"serial_number,omitempty"`
}

// Display The Display object contains information about the physical or virtual display connected to a computer system.
type Display struct {
	// The numeric color depth.
	ColorDepth int `json:"color_depth,omitempty"`
	// The numeric physical height of display.
	PhysicalHeight int `json:"physical_height,omitempty"`
	// The numeric physical orientation of display.
	PhysicalOrientation int `json:"physical_orientation,omitempty"`
	// The numeric physical width of display.
	PhysicalWidth int `json:"physical_width,omitempty"`
	// The numeric scale factor of display.
	ScaleFactor int `json:"scale_factor,omitempty"`
}

// File The File object represents the metadata associated with a file stored in a computer system. It encompasses information about the file itself, including its attributes, properties, and organizational details. Defined by D3FEND <a target='_blank' href='https://next.d3fend.mitre.org/dao/artifact/d3f:File/'>d3f:File</a>.
type File struct {
	// The time when the file was last accessed.
	AccessedTime time.Time `json:"accessed_time,omitempty"`
	// The time when the file was last accessed.
	AccessedTimeDt time.Time `json:"accessed_time_dt,omitempty"`
	// The name of the user who last accessed the object.
	Accessor *User `json:"accessor,omitempty"`
	// The bitmask value that represents the file attributes.
	Attributes int `json:"attributes,omitempty"`
	// The name of the company that published the file. For example: <code>Microsoft Corporation</code>.
	CompanyName string `json:"company_name,omitempty"`
	// The file content confidentiality, normalized to the confidentiality_id value. In the case of 'Other', it is defined by the event source.
	Confidentiality string `json:"confidentiality,omitempty"`
	// The normalized identifier of the file content confidentiality indicator.
	ConfidentialityId int `json:"confidentiality_id,omitempty"`
	// The time when the file was created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time when the file was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The user that created the file.
	Creator *User `json:"creator,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The description of the file, as returned by file system. For example: the description as returned by the Unix file command or the Windows file type.
	Desc string `json:"desc,omitempty"`
	// The extension of the file, excluding the leading dot. For example: <code>exe</code> from <code>svchost.exe</code>, or <code>gz</code> from <code>export.tar.gz</code>.
	Ext string `json:"ext"`
	// An array of hash attributes.
	Hashes []*Fingerprint `json:"hashes"`
	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`
	// The Multipurpose Internet Mail Extensions (MIME) type of the file, if applicable.
	MimeType string `json:"mime_type,omitempty"`
	// The time when the file was last modified.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The time when the file was last modified.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The user that last modified the file.
	Modifier *User `json:"modifier,omitempty"`
	// The name of the file. For example: <code>svchost.exe</code>
	Name string `json:"name"`
	// The user that owns the file/object.
	Owner *User `json:"owner,omitempty"`
	// The parent folder in which the file resides. For example: <code>c:\windows\system32</code>
	ParentFolder string `json:"parent_folder,omitempty"`
	// The full path to the file. For example: <code>c:\windows\system32\svchost.exe</code>.
	Path string `json:"path"`
	// The product that created or installed the file.
	Product *Product `json:"product,omitempty"`
	// The object security descriptor.
	SecurityDescriptor string `json:"security_descriptor,omitempty"`
	// The digital signature of the file.
	Signature *DigitalSignature `json:"signature,omitempty"`
	// The size of data, in bytes.
	Size int64 `json:"size,omitempty"`
	// The file type.
	Type string `json:"type,omitempty"`
	// The file type ID.
	TypeId int `json:"type_id"`
	// The unique identifier of the file as defined by the storage system, such the file system file ID.
	Uid string `json:"uid,omitempty"`
	// The file version. For example: <code>8.0.7601.17514</code>.
	Version string `json:"version,omitempty"`
	// An unordered collection of zero or more name/value pairs where each pair represents a file or folder extended attribute.</p>For example: Windows alternate data stream attributes (ADS stream name, ADS size, etc.), user-defined or application-defined attributes, ACL, owner, primary group, etc. Examples from DCS: </p><ul><li><strong>ads_name</strong></li><li><strong>ads_size</strong></li><li><strong>dacl</strong></li><li><strong>owner</strong></li><li><strong>primary_group</strong></li><li><strong>link_name</strong> - name of the link associated to the file.</li><li><strong>hard_link_count</strong> - the number of links that are associated to the file.</li></ul>
	Xattributes *Object `json:"xattributes,omitempty"`
}

// Kernel The Kernel Resource object provides information about a specific kernel resource, including its name and type. It describes essential attributes associated with a resource managed by the kernel of an operating system. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Kernel/'>d3f:Kernel</a>.
type Kernel struct {
	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`
	// The name of the kernel resource.
	Name string `json:"name"`
	// The full path of the kernel resource.
	Path string `json:"path,omitempty"`
	// The system call that was invoked.
	SystemCall string `json:"system_call,omitempty"`
	// The type of the kernel resource.
	Type string `json:"type,omitempty"`
	// The type of the kernel resource.
	TypeId int `json:"type_id"`
}

// KernelDriver The Kernel Extension object describes a kernel driver that has been loaded or unloaded into the operating system (OS) kernel. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:KernelModule/'>d3f:KernelModule</a>.
type KernelDriver struct {
	// The driver/extension file object.
	File *File `json:"file"`
}

// KeyboardInfo The Keyboard Information object contains details and attributes related to a computer or device keyboard. It encompasses information that describes the characteristics, capabilities, and configuration of the keyboard.
type KeyboardInfo struct {
	// The number of function keys on client keyboard.
	FunctionKeys int `json:"function_keys,omitempty"`
	// The Input Method Editor (IME) file name.
	Ime string `json:"ime,omitempty"`
	// The keyboard locale identifier name (e.g., en-US).
	KeyboardLayout string `json:"keyboard_layout,omitempty"`
	// The keyboard numeric code.
	KeyboardSubtype int `json:"keyboard_subtype,omitempty"`
	// The keyboard type (e.g., xt, ico).
	KeyboardType string `json:"keyboard_type,omitempty"`
}

// LoadBalancer The load balancer object describes the load balancer entity and contains additional information regarding the distribution of traffic across a network.
type LoadBalancer struct {
	// The request classification as defined by the load balancer.
	Classification string `json:"classification,omitempty"`
	// The numeric response status code detailing the connection from the load balancer to the destination target.
	Code int `json:"code"`
	// The destination to which the load balancer is distributing traffic.
	DstEndpoint *NetworkEndpoint `json:"dst_endpoint"`
	// An object detailing the load balancer connection attempts and responses.
	EndpointConnections []*EndpointConnection `json:"endpoint_connections"`
	// The load balancer error message.
	ErrorMessage string `json:"error_message,omitempty"`
	// The IP address of the load balancer node that handled the client request. Note: the load balancer may have other IP addresses, and this is not an IP address of the target/distribution endpoint - see <code>dst_endpoint</code>.
	Ip string `json:"ip,omitempty"`
	// The load balancer message.
	Message string `json:"message,omitempty"`
	// General purpose metrics associated with the load balancer.
	Metrics []*Metric `json:"metrics,omitempty"`
	// The name of the load balancer.
	Name string `json:"name"`
	// The status detail contains additional status information about the load balancer distribution event.
	StatusDetail string `json:"status_detail,omitempty"`
	// The unique identifier for the load balancer.
	Uid string `json:"uid"`
}

// Module The Module object describes the load attributes of a module.
type Module struct {
	// The memory address where the module was loaded.
	BaseAddress string `json:"base_address"`
	// The module file object.
	File *File `json:"file"`
	// The entry-point function of the module. The system calls the entry-point function whenever a process or thread loads or unloads the module.
	FunctionName string `json:"function_name,omitempty"`
	// The load type, normalized to the caption of the load_type_id value. In the case of 'Other', it is defined by the event source.
	LoadType string `json:"load_type,omitempty"`
	// The normalized identifier for how the module was loaded in memory.
	LoadTypeId int `json:"load_type_id"`
	// The start address of the execution.
	StartAddress string `json:"start_address"`
	// The module type.
	Type string `json:"type"`
}

// Os The Operating System (OS) object describes characteristics of an OS, such as Linux or Windows. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:OperatingSystem/'>d3f:OperatingSystem</a>.
type Os struct {
	// The operating system build number.
	Build string `json:"build,omitempty"`
	// The operating system country code, as defined by the ISO 3166-1 standard (Alpha-2 code). For the complete list of country codes, see <a target='_blank' href='https://www.iso.org/obp/ui/#iso:pub:PUB500001:en'>ISO 3166-1 alpha-2 codes</a>.
	Country string `json:"country,omitempty"`
	// The Common Platform Enumeration (CPE) name as described by (<a target='_blank' href='https://nvd.nist.gov/products/cpe'>NIST</a>) For example: <code>cpe:/a:apple:safari:16.2</code>.
	CpeName string `json:"cpe_name,omitempty"`
	// The cpu architecture, the number of bits used for addressing in memory. For example: <code>32</code> or <code>64</code>.
	CpuBits int `json:"cpu_bits,omitempty"`
	// The operating system edition. For example: <code>Professional</code>.
	Edition string `json:"edition,omitempty"`
	// The two letter lower case language codes, as defined by <a target='_blank' href='https://en.wikipedia.org/wiki/ISO_639-1'>ISO 639-1</a>. For example: <code>en</code> (English), <code>de</code> (German), or <code>fr</code> (French).
	Lang string `json:"lang,omitempty"`
	// The operating system name.
	Name string `json:"name"`
	// The name of the latest Service Pack.
	SpName string `json:"sp_name,omitempty"`
	// The version number of the latest Service Pack.
	SpVer int `json:"sp_ver,omitempty"`
	// The type of the operating system.
	Type string `json:"type,omitempty"`
	// The type identifier of the operating system.
	TypeId int `json:"type_id"`
	// The version of the OS running on the device that originated the event. For example: "Windows 10", "OS X 10.7", or "iOS 9".
	Version string `json:"version,omitempty"`
}

// PeripheralDevice The peripheral device object describes the identity, vendor and model of a peripheral device.
type PeripheralDevice struct {
	// The class of the peripheral device.
	Class string `json:"class"`
	// The peripheral device model.
	Model string `json:"model"`
	// The name of the peripheral device.
	Name string `json:"name"`
	// The peripheral device serial number.
	SerialNumber string `json:"serial_number"`
	// The unique identifier of the peripheral device.
	Uid string `json:"uid"`
	// The peripheral device vendor.
	VendorName string `json:"vendor_name"`
}

// Process The Process object describes a running instance of a launched program. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Process/'>d3f:Process</a>.
type Process struct {
	// The audit user assigned at login by the audit subsystem.
	Auid int `json:"auid,omitempty"`
	// The full command line used to launch an application, service, process, or job. For example: <code>ssh user@10.0.0.10</code>. If the command line is unavailable or missing, the empty string <code>''</code> is to be used.
	CmdLine string `json:"cmd_line"`
	// The information describing an instance of a container. A container is a prepackaged, portable system image that runs isolated on an existing system using a container runtime like containerd.
	Container *Container `json:"container"`
	// The time when the process was created/started.
	CreatedTime time.Time `json:"created_time"`
	// The time when the process was created/started.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The effective group under which this process is running.
	Egid int `json:"egid,omitempty"`
	// The effective user under which this process is running.
	Euid int `json:"euid,omitempty"`
	// The process file object.
	File *File `json:"file"`
	// The group under which this process is running.
	Group *Group `json:"group"`
	// The process integrity level, normalized to the caption of the integrity_id value. In the case of 'Other', it is defined by the event source (Windows only).
	Integrity string `json:"integrity,omitempty"`
	// The normalized identifier of the process integrity level (Windows only).
	IntegrityId int `json:"integrity_id,omitempty"`
	// The lineage of the process, represented by a list of paths for each ancestor process. For example: <code>['/usr/sbin/sshd', '/usr/bin/bash', '/usr/bin/whoami']</code>.
	Lineage []string `json:"lineage,omitempty"`
	// The list of loaded module names.
	LoadedModules []string `json:"loaded_modules,omitempty"`
	// The friendly name of the process, for example: <code>Notepad++</code>.
	Name interface{} `json:"name"`
	// If running under a process namespace (such as in a container), the process identifier within that process namespace.
	NamespacePid int `json:"namespace_pid"`
	// The parent process of this process object. It is recommended to only populate this field for the first process object, to prevent deep nesting.
	ParentProcess *Process `json:"parent_process"`
	// The process identifier, as reported by the operating system. Process ID (PID) is a number used by the operating system to uniquely identify an active process.
	Pid int `json:"pid"`
	// The name of the containment jail (i.e., sandbox). For example, hardened_ps, high_security_ps, oracle_ps, netsvcs_ps, or default_ps.
	Sandbox string `json:"sandbox,omitempty"`
	// The user session under which this process is running.
	Session *Session `json:"session,omitempty"`
	// The time when the process was terminated.
	TerminatedTime time.Time `json:"terminated_time,omitempty"`
	// The time when the process was terminated.
	TerminatedTimeDt time.Time `json:"terminated_time_dt,omitempty"`
	// The Identifier of the thread associated with the event, as returned by the operating system.
	Tid int `json:"tid,omitempty"`
	// A unique identifier for this process assigned by the producer (tool).  Facilitates correlation of a process event with other events for that process.
	Uid string `json:"uid"`
	// The user under which this process is running.
	User *User `json:"user"`
	// An unordered collection of zero or more name/value pairs that represent a process extended attribute.
	Xattributes *Object `json:"xattributes,omitempty"`
}

// Product The Product object describes characteristics of a software product.
type Product struct {
	// The Common Platform Enumeration (CPE) name as described by (<a target='_blank' href='https://nvd.nist.gov/products/cpe'>NIST</a>) For example: <code>cpe:/a:apple:safari:16.2</code>.
	CpeName string `json:"cpe_name,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The feature that reported the event.
	Feature *Feature `json:"feature,omitempty"`
	// The two letter lower case language codes, as defined by <a target='_blank' href='https://en.wikipedia.org/wiki/ISO_639-1'>ISO 639-1</a>. For example: <code>en</code> (English), <code>de</code> (German), or <code>fr</code> (French).
	Lang string `json:"lang,omitempty"`
	// The name of the product.
	Name string `json:"name"`
	// The installation path of the product.
	Path string `json:"path,omitempty"`
	// The unique identifier of the product.
	Uid string `json:"uid"`
	// The URL pointing towards the product.
	UrlString string `json:"url_string,omitempty"`
	// The name of the vendor of the product.
	VendorName string `json:"vendor_name"`
	// The version of the product, as defined by the event source. For example: <code>2013.1.3-beta</code>.
	Version string `json:"version"`
}

// Service The Service object describes characteristics of a service, <code> e.g. AWS EC2. </code>
type Service struct {
	// The list of labels associated with the service.
	Labels []string `json:"labels,omitempty"`
	// The name of the service.
	Name string `json:"name"`
	// The unique identifier of the service.
	Uid string `json:"uid"`
	// The version of the service.
	Version string `json:"version"`
}

// WebResource The Web Resource object describes characteristics of a web resource that was affected by the activity/event.
type WebResource struct {
	// Details of the web resource, e.g, <code>file</code> details, <code>search</code> results or application-defined resource.
	Data json.RawMessage `json:"data,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// Description of the web resource.
	Desc string `json:"desc,omitempty"`
	// The list of labels/tags associated to a resource.
	Labels []string `json:"labels,omitempty"`
	// The name of the web resource.
	Name string `json:"name"`
	// The web resource type as defined by the event source.
	Type string `json:"type,omitempty"`
	// The unique identifier of the web resource.
	Uid string `json:"uid"`
	// The URL pointing towards the source of the web resource.
	UrlString string `json:"url_string"`
}

// WinRegKey The registry key object describes a Windows registry key. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:WindowsRegistryKey/'>d3f:WindowsRegistryKey</a>.
type WinRegKey struct {
	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`
	// The time when the registry key was last modified.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The time when the registry key was last modified.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The full path to the registry key.
	Path string `json:"path"`
	// The security descriptor of the registry key.
	SecurityDescriptor string `json:"security_descriptor,omitempty"`
}

// WinRegValue The registry value object describes a Windows registry value.
type WinRegValue struct {
	// The data of the registry value.
	Data json.RawMessage `json:"data,omitempty"`
	// The indication of whether the value is from a default value name. For example, the value name could be missing.
	IsDefault bool `json:"is_default,omitempty"`
	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`
	// The time when the registry value was last modified.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The time when the registry value was last modified.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The name of the registry value.
	Name string `json:"name"`
	// The full path to the registry key, where the value is located.
	Path string `json:"path"`
	// A string representation of the value type as specified in <a target='_blank' href='https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-value-types'>Registry Value Types</a>.
	Type string `json:"type,omitempty"`
	// The value type ID.
	TypeId int `json:"type_id"`
}

// WinWinResource The Windows resource object describes a resource object managed by Windows, such as mutant or timer.
type WinWinResource struct {
	// Additional data describing the resource.
	Data json.RawMessage `json:"data,omitempty"`
	// The Data Classification object includes information about data classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification"`
	// The string detailing the attributes of the resource object.
	Details string `json:"details,omitempty"`
	// The list of labels/tags associated to a resource.
	Labels []string `json:"labels,omitempty"`
	// The name of the resource object.
	Name string `json:"name"`
	// The Windows service acting as the object server for the resource object, such as Security or Security Account Manager.
	SvcName string `json:"svc_name,omitempty"`
	// The type of the Windows resource object.
	Type string `json:"type,omitempty"`
	// The normalized type identifier of the Windows resource object accessed.
	TypeId int `json:"type_id"`
	// The Windows provided handle identifier for the resource object
	Uid string `json:"uid"`
}

// WinWinService The Windows Service object describes a Windows service.
type WinWinService struct {
	// The full command line used to launch the service.
	CmdLine string `json:"cmd_line"`
	// The list of labels associated with the service.
	Labels []string `json:"labels,omitempty"`
	// The name of the load ordering group of which this service is a member.
	LoadOrderGroup string `json:"load_order_group"`
	// The unique name of the service.
	Name string `json:"name"`
	// The service category, normalized to the caption of the service_category_id value. In the case of 'Other', it is defined by the event source.
	ServiceCategory string `json:"service_category,omitempty"`
	// The normalized identifier of the service category.
	ServiceCategoryId int `json:"service_category_id"`
	// The names of other services upon which this service has a dependency.
	ServiceDependencies []string `json:"service_dependencies"`
	// The service error control, normalized to the caption of the <code>service_error_control_id</code> value. In the case of 'Other', it is defined by the event source.
	ServiceErrorControl string `json:"service_error_control,omitempty"`
	// The normalized identifier of the service error control.
	ServiceErrorControlId int `json:"service_error_control_id"`
	// For a user mode service, this attribute represents the name of the account under which the service is run. For a kernel mode driver, this attribute represents the object name used to load the driver.
	ServiceStartName string `json:"service_start_name"`
	// The service start type, normalized to the caption of the <code>service_start_type_id</code> value. In the case of 'Other', it is defined by the event source.
	ServiceStartType string `json:"service_start_type,omitempty"`
	// The normalized identifier of the service start type.
	ServiceStartTypeId int `json:"service_start_type_id"`
	// The service type, normalized to the caption of the service_type_id value. In the case of 'Other', it is defined by the event source.
	ServiceType string `json:"service_type,omitempty"`
	// The normalized identifier of the service type.
	ServiceTypeId int `json:"service_type_id"`
	// The unique identifier of the service.
	Uid string `json:"uid"`
	// The version of the service.
	Version string `json:"version"`
}


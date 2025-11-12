package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Device The Device object represents an addressable computer system or host, which is
// typically connected to a computer network and participates in the transmission
// or processing of data within the computer network.
type Device struct {
	// A list of agent objects associated with a device, endpoint, or resource.
	AgentList []*Agent `json:"agent_list,omitempty"`

	// The unique identifier of the cloud autoscale configuration.
	AutoscaleUID string `json:"autoscale_uid,omitempty"`

	// The time the system was booted.
	BootTime int64 `json:"boot_time,omitempty"`

	// The time the system was booted.
	BootTimeDt string `json:"boot_time_dt,omitempty"`

	// A unique identifier of the device that changes after every reboot. For
	// example, the value of /proc/sys/kernel/random/boot_id from Linux's procfs.
	BootUID string `json:"boot_uid,omitempty"`

	// The information describing an instance of a container. A container is a
	// prepackaged, portable system image that runs isolated on an existing system
	// using a container runtime like containerd.
	Container *Container `json:"container,omitempty"`

	// The time when the device was known to have been created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the device was known to have been created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The description of the device, ordinarily as reported by the operating
	// system.
	Desc string `json:"desc,omitempty"`

	// The network domain where the device resides. For example: work.example.com.
	Domain string `json:"domain,omitempty"`

	// An Embedded Identity Document, is a unique serial number that identifies an
	// eSIM-enabled device.
	Eid string `json:"eid,omitempty"`

	// The initial discovery time of the device.
	FirstSeenTime int64 `json:"first_seen_time,omitempty"`

	// The initial discovery time of the device.
	FirstSeenTimeDt string `json:"first_seen_time_dt,omitempty"`

	// The group names to which the device belongs. For example: ["Windows
	// Laptops", "Engineering"].
	Groups []*Group `json:"groups,omitempty"`

	// The device hostname.
	Hostname string `json:"hostname,omitempty"`

	// The endpoint hardware information.
	HwInfo *DeviceHwInfo `json:"hw_info,omitempty"`

	// The name of the hypervisor running on the device. For example, Xen, VMware,
	// Hyper-V, VirtualBox, etc.
	Hypervisor string `json:"hypervisor,omitempty"`

	// The Integrated Circuit Card Identification of a mobile device. Typically it
	// is a unique 18 to 22 digit number that identifies a SIM card.
	Iccid string `json:"iccid,omitempty"`

	// The image used as a template to run the virtual machine.
	Image *Image `json:"image,omitempty"`

	// The International Mobile Equipment Identity that is associated with the
	// device.
	Imei string `json:"imei,omitempty"`

	// The International Mobile Equipment Identity values that are associated with
	// the device.
	ImeiList []string `json:"imei_list,omitempty"`

	// The unique identifier of a VM instance.
	InstanceUID string `json:"instance_uid,omitempty"`

	// The name of the network interface (e.g. eth2).
	InterfaceName string `json:"interface_name,omitempty"`

	// The unique identifier of the network interface.
	InterfaceUID string `json:"interface_uid,omitempty"`

	// The device IP address, in either IPv4 or IPv6 format.
	Ip string `json:"ip,omitempty"`

	// Indicates whether the device or resource has a backup enabled, such as an
	// automated snapshot or a cloud backup. For example, this is indicated by the
	// cloudBackupEnabled value within JAMF Pro mobile devices or the registration
	// of an AWS ARN with the AWS Backup service.
	IsBackedUp bool `json:"is_backed_up,omitempty"`

	// The event occurred on a compliant device.
	IsCompliant bool `json:"is_compliant,omitempty"`

	// The event occurred on a managed device.
	IsManaged bool `json:"is_managed,omitempty"`

	// Indicates whether the device has an active mobile account. For example, this
	// is indicated by the itunesStoreAccountActive value within JAMF Pro mobile
	// devices.
	IsMobileAccountActive bool `json:"is_mobile_account_active,omitempty"`

	// The event occurred on a personal device.
	IsPersonal bool `json:"is_personal,omitempty"`

	// The event occurred on a shared device.
	IsShared bool `json:"is_shared,omitempty"`

	// The event occurred on a supervised device. Devices that are supervised are
	// typically mobile devices managed by a Mobile Device Management solution and
	// are restricted from specific behaviors such as Apple AirDrop.
	IsSupervised bool `json:"is_supervised,omitempty"`

	// The event occurred on a trusted device.
	IsTrusted bool `json:"is_trusted,omitempty"`

	// The most recent discovery time of the device.
	LastSeenTime int64 `json:"last_seen_time,omitempty"`

	// The most recent discovery time of the device.
	LastSeenTimeDt string `json:"last_seen_time_dt,omitempty"`

	// The geographical location of the device.
	Location *Location `json:"location,omitempty"`

	// The Media Access Control (MAC) address of the endpoint.
	Mac string `json:"mac,omitempty"`

	// The Mobile Equipment Identifier. It's a unique number that identifies a Code
	// Division Multiple Access (CDMA) mobile device.
	Meid string `json:"meid,omitempty"`

	// The model of the device. For example ThinkPad X1 Carbon.
	Model string `json:"model,omitempty"`

	// The time when the device was last known to have been modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the device was last known to have been modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The alternate device name, ordinarily as assigned by an administrator.
	// <p><b>Note:</b> The <b>Name</b> could be any other string that helps to
	// identify the device, such as a phone number; for example 310-555-1234.</p>
	Name string `json:"name,omitempty"`

	// If running under a process namespace (such as in a container), the process
	// identifier within that process namespace.
	NamespacePid int64 `json:"namespace_pid,omitempty"`

	// The physical or virtual network interfaces that are associated with the
	// device, one for each unique MAC address/IP address/hostname/name
	// combination.<p><b>Note:</b> The first element of the array is the network
	// information that pertains to the event.</p>
	NetworkInterfaces []*NetworkInterface `json:"network_interfaces,omitempty"`

	// Organization and org unit related to the device.
	Org *Organization `json:"org,omitempty"`

	// The endpoint operating system.
	Os *Os `json:"os,omitempty"`

	// The operating system assigned Machine ID. In Windows, this is the value
	// stored at the registry path:
	// HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Cryptography\MachineGuid. In Linux,
	// this is stored in the file: /etc/machine-id.
	OsMAChineUUID string `json:"os_machine_uuid,omitempty"`

	// The identity of the service or user account that owns the endpoint or was
	// last logged into it.
	Owner *User `json:"owner,omitempty"`

	// The region where the virtual machine is located. For example, an AWS Region.
	Region string `json:"region,omitempty"`

	// The risk level, normalized to the caption of the risk_level_id value.
	RiskLevel string `json:"risk_level,omitempty"`

	// The normalized risk level id.
	RiskLevelID int `json:"risk_level_id,omitempty"`

	// The risk score as reported by the event source.
	RiskScore int64 `json:"risk_score,omitempty"`

	// The subnet mask.
	Subnet interface{} `json:"subnet,omitempty"`

	// The unique identifier of a virtual subnet.
	SubnetUID string `json:"subnet_uid,omitempty"`

	// The device type. For example: unknown, server, desktop, laptop, tablet,
	// mobile, virtual, browser, or other.
	Type string `json:"type,omitempty"`

	// The device type ID.
	TypeID int `json:"type_id,omitempty"`

	// The Apple assigned Unique Device Identifier (UDID). For iOS, iPadOS, tvOS,
	// watchOS and visionOS devices, this is the UDID. For macOS devices, it is the
	// Provisioning UDID. For example: 00008020-008D4548007B4F26
	Udid string `json:"udid,omitempty"`

	// The unique identifier of the device. For example the Windows TargetSID or
	// AWS EC2 ARN.
	Uid string `json:"uid,omitempty"`

	// An alternate unique identifier of the device if any. For example the
	// ActiveDirectory DN.
	UidAlt string `json:"uid_alt,omitempty"`

	// The vendor for the device. For example Dell or Lenovo.
	VendorName string `json:"vendor_name,omitempty"`

	// The Virtual LAN identifier.
	VlanUID string `json:"vlan_uid,omitempty"`

	// The unique identifier of the Virtual Private Cloud (VPC).
	VpcUID string `json:"vpc_uid,omitempty"`

	// The network zone or LAN segment.
	Zone string `json:"zone,omitempty"`
}

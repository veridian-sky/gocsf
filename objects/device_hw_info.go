package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DeviceHwInfo The Device Hardware Information object contains details and specifications of
// the physical components that make up a device. This information provides an
// overview of the hardware capabilities, configuration, and characteristics of the
// device.
type DeviceHwInfo struct {
	// The BIOS date. For example: 03/31/16.
	BiosDate string `json:"bios_date,omitempty"`

	// The BIOS manufacturer. For example: LENOVO.
	BiosManufacturer string `json:"bios_manufacturer,omitempty"`

	// The BIOS version. For example: LENOVO G5ETA2WW (2.62).
	BiosVer string `json:"bios_ver,omitempty"`

	// The chassis type describes the system enclosure or physical form factor.
	// Such as the following examples for Windows
	// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-systemenclosure
	// Windows Chassis Types
	Chassis string `json:"chassis,omitempty"`

	// The CPU architecture, normalized to the caption of the cpu_architecture_id
	// value. In the case of Other, it is defined by the source.
	CpuArchitecture string `json:"cpu_architecture,omitempty"`

	// The normalized identifier of the CPU architecture.
	CpuArchitectureID int `json:"cpu_architecture_id,omitempty"`

	// The cpu architecture, the number of bits used for addressing in memory. For
	// example: 32 or 64.
	CpuBits int64 `json:"cpu_bits,omitempty"`

	// The number of processor cores in all installed processors. For Example: 42.
	CpuCores int64 `json:"cpu_cores,omitempty"`

	// The number of physical processors on a system. For example: 1.
	CpuCount int64 `json:"cpu_count,omitempty"`

	// The speed of the processor in Mhz. For Example: 4200.
	CpuSpeed int64 `json:"cpu_speed,omitempty"`

	// The processor type. For example: x86 Family 6 Model 37 Stepping 5.
	CpuType string `json:"cpu_type,omitempty"`

	// The desktop display affiliated with the event
	DesktopDisplay *Display `json:"desktop_display,omitempty"`

	// The keyboard detailed information.
	KeyboardInfo *KeyboardInfo `json:"keyboard_info,omitempty"`

	// The total amount of installed RAM, in Megabytes. For example: 2048.
	RamSize int64 `json:"ram_size,omitempty"`

	// The device manufacturer serial number.
	SerialNumber string `json:"serial_number,omitempty"`

	// The device manufacturer assigned universally unique hardware identifier. For
	// SMBIOS compatible devices such as those running Linux and Windows, it is the
	// UUID member of the System Information structure in the SMBIOS information.
	// For macOS devices, it is the Hardware UUID (also known as IOPlatformUUID in
	// the I/O Registry).
	Uuid string `json:"uuid,omitempty"`

	// The device manufacturer.
	VendorName string `json:"vendor_name,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Os The Operating System (OS) object describes characteristics of an OS, such as
// Linux or Windows.
type Os struct {
	// The operating system build number.
	Build string `json:"build,omitempty"`

	// The operating system country code, as defined by the ISO 3166-1 standard
	// (Alpha-2 code).<p><b>Note:</b> The two letter country code should be
	// capitalized. For example: US or CA.</p>
	Country string `json:"country,omitempty"`

	// The Common Platform Enumeration (CPE) name as described by
	// (https://nvd.nist.gov/products/cpe NIST) For example:
	// cpe:/a:apple:safari:16.2.
	CpeName string `json:"cpe_name,omitempty"`

	// The cpu architecture, the number of bits used for addressing in memory. For
	// example: 32 or 64.
	CpuBits int64 `json:"cpu_bits,omitempty"`

	// The operating system edition. For example: Professional.
	Edition string `json:"edition,omitempty"`

	// The kernel release of the operating system. On Unix-based systems, this is
	// determined from the uname -r command output, for example
	// "5.15.0-122-generic".
	KernelRelease string `json:"kernel_release,omitempty"`

	// The two letter lower case language codes, as defined by
	// https://en.wikipedia.org/wiki/ISO_639-1 ISO 639-1. For example: en
	// (English), de (German), or fr (French).
	Lang string `json:"lang,omitempty"`

	// The operating system name.
	Name string `json:"name,omitempty"`

	// The name of the latest Service Pack.
	SpName string `json:"sp_name,omitempty"`

	// The version number of the latest Service Pack.
	SpVer int64 `json:"sp_ver,omitempty"`

	// The type of the operating system.
	Type string `json:"type,omitempty"`

	// The type identifier of the operating system.
	TypeID int `json:"type_id,omitempty"`

	// The version of the OS running on the device that originated the event. For
	// example: "Windows 10", "OS X 10.7", or "iOS 9".
	Version string `json:"version,omitempty"`
}

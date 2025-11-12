package objects

// Code generated from OCSF schema; DO NOT EDIT.

// UnmannedAerialSystem The Unmanned Aerial System object describes the characteristics, Position
// Location Information (PLI), and other metadata of Unmanned Aerial Systems (UAS)
// and other unmanned and drone systems used in Remote ID. Remote ID is defined in
// the Standard Specification for Remote ID and Tracking (ASTM Designation:
// F3411-22a)
// https://cdn.standards.iteh.ai/samples/112830/71297057ac42432880a203654f213709/ASTM-F3411-22a.pdf
// ASTM F3411-22a.
type UnmannedAerialSystem struct {
	// The endpoint hardware information.
	HwInfo *DeviceHwInfo `json:"hw_info,omitempty"`

	// The detailed geographical location usually associated with an IP address.
	Location *Location `json:"location,omitempty"`

	// The model name of the aircraft or unmanned system.
	Model string `json:"model,omitempty"`

	// The name of the unmanned system as reported by tracking or sensing hardware.
	Name string `json:"name,omitempty"`

	// The serial number of the unmanned system. This is expressed in CTA-2063-A
	// format.
	SerialNumber string `json:"serial_number,omitempty"`

	// Ground speed of flight. This value is provided in meters per second with a
	// minimum resolution of 0.25 m/s. Special Values: Invalid, No Value, or
	// Unknown: 255 m/s.
	Speed string `json:"speed,omitempty"`

	// Provides quality/containment on horizontal ground speed. Measured in
	// meters/second.
	SpeedAccuracy string `json:"speed_accuracy,omitempty"`

	// Direction of flight expressed as a “True North-based” ground track
	// angle. This value is provided in clockwise degrees with a minimum resolution
	// of 1 degree. If aircraft is not moving horizontally, use the “Unknown”
	// value
	TrackDirection string `json:"track_direction,omitempty"`

	// The type of the UAS. For example, Helicopter, Gyroplane, Rocket, etc.
	Type string `json:"type,omitempty"`

	// The UAS type identifier.
	TypeID int `json:"type_id,omitempty"`

	// The primary identification identifier for an unmanned system. This can be a
	// Serial Number (in CTA-2063-A format, the Registration ID (provided by the
	// CAA, a UTM, or a unique Session ID.
	Uid string `json:"uid,omitempty"`

	// A secondary identification identifier for an unmanned system. This can be a
	// Serial Number (in CTA-2063-A format, the Registration ID (provided by the
	// CAA, a UTM, or a unique Session ID.
	UidAlt string `json:"uid_alt,omitempty"`

	// The Unmanned Aircraft System Traffic Management (UTM) provided universal
	// unique ID (UUID) traceable to a non-obfuscated ID where this UTM UUID acts
	// as a 'session id' to protect exposure of operationally sensitive
	// information.
	Uuid string `json:"uuid,omitempty"`

	// Vertical speed upward relative to the WGS-84 datum, measured in meters per
	// second. Special Values: Invalid, No Value, or Unknown: 63 m/s.
	VerticalSpeed string `json:"vertical_speed,omitempty"`
}

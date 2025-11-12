package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Aircraft The Aircraft object represents any aircraft or otherwise airborne asset such as
// an unmanned system, airplane, balloon, spacecraft, or otherwise. The Aircraft
// object is intended to normalized data captured or otherwise logged from active
// radar, passive radar, multi-spectral systems, or the Automatic Dependant
// Broadcast - Surveillance (ADS-B), and/or Mode S systems.
type Aircraft struct {
	// The detailed geographical location usually associated with an IP address.
	Location *Location `json:"location,omitempty"`

	// The model name of the aircraft or unmanned system.
	Model string `json:"model,omitempty"`

	// The name of the aircraft, such as the such as the flight name or callsign.
	Name string `json:"name,omitempty"`

	// The serial number of the aircraft.
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

	// The primary identification identifier for an aircraft, such as the 24-bit
	// International Civil Aviation Organization (ICAO) identifier of the aircraft,
	// as 6 hex digits.
	Uid string `json:"uid,omitempty"`

	// A secondary identification identifier for an aircraft, such as the 4-digit
	// squawk (octal representation).
	UidAlt string `json:"uid_alt,omitempty"`

	// Vertical speed upward relative to the WGS-84 datum, measured in meters per
	// second. Special Values: Invalid, No Value, or Unknown: 63 m/s.
	VerticalSpeed string `json:"vertical_speed,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// PeripheralDevice The peripheral device object describes the identity, vendor and model of a
// peripheral device.
type PeripheralDevice struct {
	// The class of the peripheral device.
	Class string `json:"class,omitempty"`

	// The peripheral device model.
	Model string `json:"model,omitempty"`

	// The name of the peripheral device.
	Name string `json:"name,omitempty"`

	// The peripheral device serial number.
	SerialNumber string `json:"serial_number,omitempty"`

	// The unique identifier of the peripheral device.
	Uid string `json:"uid,omitempty"`

	// The peripheral device vendor.
	VendorName string `json:"vendor_name,omitempty"`
}

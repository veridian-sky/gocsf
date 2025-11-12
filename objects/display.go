package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Display The Display object contains information about the physical or virtual display
// connected to a computer system.
type Display struct {
	// The numeric color depth.
	ColorDepth int64 `json:"color_depth,omitempty"`

	// The numeric physical height of display.
	PhysicalHeight int64 `json:"physical_height,omitempty"`

	// The numeric physical orientation of display.
	PhysicalOrientation int64 `json:"physical_orientation,omitempty"`

	// The numeric physical width of display.
	PhysicalWidth int64 `json:"physical_width,omitempty"`

	// The numeric scale factor of display.
	ScaleFactor int64 `json:"scale_factor,omitempty"`
}

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Location The Geo Location object describes a geographical location, usually associated
// with an IP address.
type Location struct {
	// Expressed as either height above takeoff location or height above ground
	// level (AGL) for a UAS current location. This value is provided in meters and
	// must have a minimum resolution of 1 m. Special Values: Invalid, No Value, or
	// Unknown: -1000 m.
	AerialHeight string `json:"aerial_height,omitempty"`

	// The name of the city.
	City string `json:"city,omitempty"`

	// The name of the continent.
	Continent string `json:"continent,omitempty"`

	// A two-element array, containing a longitude/latitude pair. The format
	// conforms with https://geojson.org GeoJSON. For example: [-73.983, 40.719].
	Coordinates []float64 `json:"coordinates,omitempty"`

	// The ISO 3166-1 Alpha-2 country code.<p><b>Note:</b> The two letter country
	// code should be capitalized. For example: US or CA.</p>
	Country string `json:"country,omitempty"`

	// The description of the geographical location.
	Desc string `json:"desc,omitempty"`

	// The aircraft distance above or below the ellipsoid as measured along a line
	// that passes through the aircraft and is normal to the surface of the WGS-84
	// ellipsoid. This value is provided in meters and must have a minimum
	// resolution of 1 m. Special Values: Invalid, No Value, or Unknown: -1000 m.
	GeodeticAltitude string `json:"geodetic_altitude,omitempty"`

	// Provides quality/containment on geodetic altitude. This is based on ADS-B
	// Geodetic Vertical Accuracy (GVA). Measured in meters.
	GeodeticVerticalAccuracy string `json:"geodetic_vertical_accuracy,omitempty"`

	// <p>Geohash of the geo-coordinates (latitude and
	// longitude).</p>https://en.wikipedia.org/wiki/Geohash Geohashing is a
	// geocoding system used to encode geographic coordinates in decimal degrees,
	// to a single string.
	Geohash string `json:"geohash,omitempty"`

	// Provides quality/containment on horizontal position. This is based on ADS-B
	// NACp. Measured in meters.
	HorizontalAccuracy string `json:"horizontal_accuracy,omitempty"`

	// The indication of whether the location is on premises.
	IsOnPremises bool `json:"is_on_premises,omitempty"`

	// The name of the Internet Service Provider (ISP).
	Isp string `json:"isp,omitempty"`

	// The geographical Latitude coordinate represented in Decimal Degrees (DD).
	// For example: 42.361145.
	Lat float64 `json:"lat,omitempty"`

	// The geographical Longitude coordinate represented in Decimal Degrees (DD).
	// For example: -71.057083.
	Long float64 `json:"long,omitempty"`

	// The postal code of the location.
	PostalCode string `json:"postal_code,omitempty"`

	// The uncorrected barometric pressure altitude (based on reference standard
	// 29.92 inHg, 1013.25 mb) provides a reference for algorithms that utilize
	// 'altitude deltas' between aircraft. This value is provided in meters and
	// must have a minimum resolution of 1 m.. Special Values: Invalid, No Value,
	// or Unknown: -1000 m.
	PressureAltitude string `json:"pressure_altitude,omitempty"`

	// The provider of the geographical location data.
	Provider string `json:"provider,omitempty"`

	// The alphanumeric code that identifies the principal subdivision (e.g.
	// province or state) of the country. For example, 'CH-VD' for the Canton of
	// Vaud, Switzerland
	Region string `json:"region,omitempty"`
}

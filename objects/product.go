package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Product The Product object describes characteristics of a software product.
type Product struct {
	// The Common Platform Enumeration (CPE) name as described by
	// (https://nvd.nist.gov/products/cpe NIST) For example:
	// cpe:/a:apple:safari:16.2.
	CpeName string `json:"cpe_name,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// The feature that reported the event.
	Feature *Feature `json:"feature,omitempty"`

	// The two letter lower case language codes, as defined by
	// https://en.wikipedia.org/wiki/ISO_639-1 ISO 639-1. For example: en
	// (English), de (German), or fr (French).
	Lang string `json:"lang,omitempty"`

	// The name of the product.
	Name string `json:"name,omitempty"`

	// The installation path of the product.
	Path string `json:"path,omitempty"`

	// The unique identifier of the product.
	Uid string `json:"uid,omitempty"`

	// The URL pointing towards the product.
	UrlString string `json:"url_string,omitempty"`

	// The name of the vendor of the product.
	VendorName string `json:"vendor_name,omitempty"`

	// The version of the product, as defined by the event source. For example:
	// 2013.1.3-beta.
	Version string `json:"version,omitempty"`
}

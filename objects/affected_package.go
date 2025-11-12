package objects

// Code generated from OCSF schema; DO NOT EDIT.

// AffectedPackage The Affected Package object describes details about a software package
// identified as affected by a vulnerability/vulnerabilities.
type AffectedPackage struct {
	// Architecture is a shorthand name describing the type of computer hardware
	// the packaged software is meant to run on.
	Architecture string `json:"architecture,omitempty"`

	// The Common Platform Enumeration (CPE) name as described by
	// (https://nvd.nist.gov/products/cpe NIST) For example:
	// cpe:/a:apple:safari:16.2.
	CpeName string `json:"cpe_name,omitempty"`

	// The software package epoch. Epoch is a way to define weighted dependencies
	// based on version numbers.
	Epoch int64 `json:"epoch,omitempty"`

	// The software package version in which a reported vulnerability was
	// patched/fixed.
	FixedInVersion string `json:"fixed_in_version,omitempty"`

	// Cryptographic hash to identify the binary instance of a software component.
	// This can include any component such file, package, or library.
	Hash *Fingerprint `json:"hash,omitempty"`

	// The software license applied to this package.
	License string `json:"license,omitempty"`

	// The URL pointing to the license applied on package or software. This is
	// typically a LICENSE.md file within a repository.
	LicenseURL string `json:"license_url,omitempty"`

	// The software package name.
	Name string `json:"name,omitempty"`

	// The software packager manager utilized to manage a package on a system, e.g.
	// npm, yum, dpkg etc.
	PackageManager string `json:"package_manager,omitempty"`

	// The URL of the package or library at the package manager, or the specific
	// URL or URI of an internal package manager link such as AWS CodeArtifact or
	// Artifactory.
	PackageManagerURL string `json:"package_manager_url,omitempty"`

	// The installation path of the affected package.
	Path interface{} `json:"path,omitempty"`

	// A purl is a URL string used to identify and locate a software package in a
	// mostly universal and uniform way across programming languages, package
	// managers, packaging conventions, tools, APIs and databases.
	Purl string `json:"purl,omitempty"`

	// Release is the number of times a version of the software has been packaged.
	Release string `json:"release,omitempty"`

	// Describes the recommended remediation steps to address identified issue(s).
	Remediation *Remediation `json:"remediation,omitempty"`

	// The link to the specific library or package such as within GitHub, this is
	// different from the link to the package manager where the library or package
	// is hosted.
	SrcURL string `json:"src_url,omitempty"`

	// The type of software package, normalized to the caption of the type_id
	// value. In the case of 'Other', it is defined by the source.
	Type string `json:"type,omitempty"`

	// The type of software package.
	TypeID int `json:"type_id,omitempty"`

	// A unique identifier for the package or library reported by the source tool.
	// E.g., the libId within the sbom field of an OX Security Issue or the SPDX
	// components.*.bom-ref.
	Uid string `json:"uid,omitempty"`

	// The name of the vendor who published the software package.
	VendorName string `json:"vendor_name,omitempty"`

	// The software package version.
	Version string `json:"version,omitempty"`
}

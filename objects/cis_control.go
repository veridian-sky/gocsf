package objects

// Code generated from OCSF schema; DO NOT EDIT.

// CisControl The CIS Control (aka Critical Security Control) object describes a prioritized
// set of actions to protect your organization and data from cyber-attack vectors.
// The https://www.cisecurity.org/controls CIS Controls are defined by the Center
// for Internet Security.
type CisControl struct {
	// The CIS Control description. For example: <i>Uninstall or disable
	// unnecessary services on enterprise assets and software, such as an unused
	// file sharing service, web application module, or service function.</i>
	Desc string `json:"desc,omitempty"`

	// The CIS Control name. For example: <i>4.8 Uninstall or Disable Unnecessary
	// Services on Enterprise Assets and Software.</i>
	Name string `json:"name,omitempty"`

	// The CIS Control version. For example: <i>v8</i>.
	Version string `json:"version,omitempty"`
}

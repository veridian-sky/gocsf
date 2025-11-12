package objects

// Code generated from OCSF schema; DO NOT EDIT.

// WinWinService The Windows Service object describes a Windows service.
type WinWinService struct {
	// The full command line used to launch the service.
	CmdLine string `json:"cmd_line,omitempty"`

	// The list of labels associated with the service.
	Labels []string `json:"labels,omitempty"`

	// The name of the load ordering group of which this service is a member.
	LoadOrderGroup string `json:"load_order_group,omitempty"`

	// The unique name of the service.
	Name string `json:"name,omitempty"`

	// The service category, normalized to the caption of the service_category_id
	// value. In the case of 'Other', it is defined by the event source.
	ServiceCategory string `json:"service_category,omitempty"`

	// The normalized identifier of the service category.
	ServiceCategoryID int `json:"service_category_id,omitempty"`

	// The names of other services upon which this service has a dependency.
	ServiceDependencies []string `json:"service_dependencies,omitempty"`

	// The service error control, normalized to the caption of the
	// service_error_control_id value. In the case of 'Other', it is defined by the
	// event source.
	ServiceErrorControl string `json:"service_error_control,omitempty"`

	// The normalized identifier of the service error control.
	ServiceErrorControlID int `json:"service_error_control_id,omitempty"`

	// For a user mode service, this attribute represents the name of the account
	// under which the service is run. For a kernel mode driver, this attribute
	// represents the object name used to load the driver.
	ServiceStartName string `json:"service_start_name,omitempty"`

	// The service start type, normalized to the caption of the
	// service_start_type_id value. In the case of 'Other', it is defined by the
	// event source.
	ServiceStartType string `json:"service_start_type,omitempty"`

	// The normalized identifier of the service start type.
	ServiceStartTypeID int `json:"service_start_type_id,omitempty"`

	// The service type, normalized to the caption of the service_type_id value. In
	// the case of 'Other', it is defined by the event source.
	ServiceType string `json:"service_type,omitempty"`

	// The normalized identifier of the service type.
	ServiceTypeID int `json:"service_type_id,omitempty"`

	// The list of tags; {key:value} pairs associated to the service.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The unique identifier of the service.
	Uid string `json:"uid,omitempty"`

	// The version of the service.
	Version string `json:"version,omitempty"`
}

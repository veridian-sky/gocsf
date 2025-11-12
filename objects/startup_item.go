package objects

// Code generated from OCSF schema; DO NOT EDIT.

// StartupItem The startup item object describes an application component that has associated
// startup criteria and configurations.
type StartupItem struct {
	// The startup item kernel driver resource.
	Driver *KernelDriver `json:"driver,omitempty"`

	// The startup item job resource.
	Job *Job `json:"job,omitempty"`

	// The unique name of the startup item.
	Name string `json:"name,omitempty"`

	// The startup item process resource.
	Process *Process `json:"process,omitempty"`

	// The list of normalized identifiers that describe the startup items'
	// properties when it is running. Use this field to capture extended
	// information about the process, which may depend on the type of startup item.
	// E.g., A Windows service that interacts with the desktop.
	RunModeIDs []int `json:"run_mode_ids,omitempty"`

	// The list of run_modes, normalized to the captions of the run_mode_id values.
	// In the case of 'Other', they are defined by the event source.
	RunModes []string `json:"run_modes,omitempty"`

	// The run state of the startup item.
	RunState string `json:"run_state,omitempty"`

	// The run state ID of the startup item.
	RunStateID int `json:"run_state_id,omitempty"`

	// The start type of the startup item.
	StartType string `json:"start_type,omitempty"`

	// The start type ID of the startup item.
	StartTypeID int `json:"start_type_id,omitempty"`

	// The startup item type.
	Type string `json:"type,omitempty"`

	// The startup item type identifier.
	TypeID int `json:"type_id,omitempty"`

	// The startup item Windows service resource.
	WinService *WinWinService `json:"win_service,omitempty"`
}

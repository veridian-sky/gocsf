package objects

// Code generated from OCSF schema; DO NOT EDIT.

// ProcessEntity The Process Entity object provides critical fields for referencing a process.
type ProcessEntity struct {
	// The full command line used to launch an application, service, process, or
	// job. For example: ssh user@10.0.0.10. If the command line is unavailable or
	// missing, the empty string '' is to be used.
	CmdLine string `json:"cmd_line,omitempty"`

	// A unique process identifier that can be assigned deterministically by
	// multiple system data producers.
	Cpid string `json:"cpid,omitempty"`

	// The time when the process was created/started.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the process was created/started.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The friendly name of the process, for example: Notepad++.
	Name interface{} `json:"name,omitempty"`

	// The process file path.
	Path string `json:"path,omitempty"`

	// The process identifier, as reported by the operating system. Process ID
	// (PID) is a number used by the operating system to uniquely identify an
	// active process.
	Pid int64 `json:"pid,omitempty"`

	// A unique identifier for this process assigned by the producer (tool).
	// Facilitates correlation of a process event with other events for that
	// process.
	Uid string `json:"uid,omitempty"`
}

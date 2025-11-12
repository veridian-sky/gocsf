package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Process The Process object describes a running instance of a launched program.
type Process struct {
	// An array of Process Entities describing the extended parentage of this
	// process object. Direct parent information should be expressed through the
	// parent_process attribute. The first array element is the direct parent of
	// this process object. Subsequent list elements go up the process parentage
	// hierarchy. That is, the array is sorted from newest to oldest process. It is
	// recommended to only populate this field for the top-level process object.
	Ancestry []*ProcessEntity `json:"ancestry,omitempty"`

	// The audit user assigned at login by the audit subsystem.
	Auid int64 `json:"auid,omitempty"`

	// The full command line used to launch an application, service, process, or
	// job. For example: ssh user@10.0.0.10. If the command line is unavailable or
	// missing, the empty string '' is to be used.
	CmdLine string `json:"cmd_line,omitempty"`

	// The information describing an instance of a container. A container is a
	// prepackaged, portable system image that runs isolated on an existing system
	// using a container runtime like containerd.
	Container *Container `json:"container,omitempty"`

	// A unique process identifier that can be assigned deterministically by
	// multiple system data producers.
	Cpid string `json:"cpid,omitempty"`

	// The time when the process was created/started.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the process was created/started.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The effective group under which this process is running.
	Egid int64 `json:"egid,omitempty"`

	// Environment variables associated with the process.
	EnvironmentVariables []*EnvironmentVariable `json:"environment_variables,omitempty"`

	// The effective user under which this process is running.
	Euid int64 `json:"euid,omitempty"`

	// The process file object.
	File *File `json:"file,omitempty"`

	// The group under which this process is running.
	Group *Group `json:"group,omitempty"`

	// The process integrity level, normalized to the caption of the integrity_id
	// value. In the case of 'Other', it is defined by the event source (Windows
	// only).
	Integrity string `json:"integrity,omitempty"`

	// The normalized identifier of the process integrity level (Windows only).
	IntegrityID int `json:"integrity_id,omitempty"`

	// The lineage of the process, represented by a list of paths for each ancestor
	// process. For example: ['/usr/sbin/sshd', '/usr/bin/bash',
	// '/usr/bin/whoami'].
	Lineage []interface{} `json:"lineage,omitempty"`

	// The list of loaded module names.
	LoadedModules []string `json:"loaded_modules,omitempty"`

	// The friendly name of the process, for example: Notepad++.
	Name interface{} `json:"name,omitempty"`

	// If running under a process namespace (such as in a container), the process
	// identifier within that process namespace.
	NamespacePid int64 `json:"namespace_pid,omitempty"`

	// The parent process of this process object. It is recommended to only
	// populate this field for the top-level process object, to prevent deep
	// nesting. Additional ancestry information can be supplied in the ancestry
	// attribute.
	ParentProcess *Process `json:"parent_process,omitempty"`

	// The process file path.
	Path string `json:"path,omitempty"`

	// The process identifier, as reported by the operating system. Process ID
	// (PID) is a number used by the operating system to uniquely identify an
	// active process.
	Pid int64 `json:"pid,omitempty"`

	// The identifier of the process thread associated with the event, as returned
	// by the operating system.
	Ptid int64 `json:"ptid,omitempty"`

	// The name of the containment jail (i.e., sandbox). For example, hardened_ps,
	// high_security_ps, oracle_ps, netsvcs_ps, or default_ps.
	Sandbox string `json:"sandbox,omitempty"`

	// The user session under which this process is running.
	Session *Session `json:"session,omitempty"`

	// The time when the process was terminated.
	TerminatedTime int64 `json:"terminated_time,omitempty"`

	// The time when the process was terminated.
	TerminatedTimeDt string `json:"terminated_time_dt,omitempty"`

	// The identifier of the thread associated with the event, as returned by the
	// operating system.
	Tid int64 `json:"tid,omitempty"`

	// A unique identifier for this process assigned by the producer (tool).
	// Facilitates correlation of a process event with other events for that
	// process.
	Uid string `json:"uid,omitempty"`

	// The user under which this process is running.
	User *User `json:"user,omitempty"`

	// The working directory of a process.
	WorkingDirectory string `json:"working_directory,omitempty"`

	// An unordered collection of zero or more name/value pairs that represent a
	// process extended attribute.
	Xattributes *Object `json:"xattributes,omitempty"`
}

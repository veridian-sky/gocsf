package objects

// Code generated from OCSF schema; DO NOT EDIT.

// DceRpc The DCE/RPC, or Distributed Computing Environment/Remote Procedure Call, object
// describes the remote procedure call system for distributed computing
// environments.
type DceRpc struct {
	// The request command (e.g. REQUEST, BIND).
	Command string `json:"command,omitempty"`

	// The reply to the request command (e.g. RESPONSE, BINDACK or FAULT).
	CommandResponse string `json:"command_response,omitempty"`

	// The list of interface flags.
	Flags []string `json:"flags,omitempty"`

	// An operation number used to identify a specific remote procedure call (RPC)
	// method or a method in an interface.
	Opnum int64 `json:"opnum,omitempty"`

	// The RPC Interface object describes the details pertaining to the remote
	// procedure call interface.
	RpcInterface *RpcInterface `json:"rpc_interface,omitempty"`
}

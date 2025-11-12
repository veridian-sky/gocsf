package objects

// Code generated from OCSF schema; DO NOT EDIT.

// RpcInterface The RPC Interface represents the remote procedure call interface used in the
// DCE/RPC session.
type RpcInterface struct {
	// An integer that provides a reason code or additional information about the
	// acknowledgment result.
	AckReason int64 `json:"ack_reason,omitempty"`

	// An integer that denotes the acknowledgment result of the DCE/RPC call.
	AckResult int64 `json:"ack_result,omitempty"`

	// The unique identifier of the particular remote procedure or service.
	Uuid string `json:"uuid,omitempty"`

	// The version of the DCE/RPC protocol being used in the session.
	Version string `json:"version,omitempty"`
}

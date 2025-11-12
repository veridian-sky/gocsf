package objects

// Code generated from OCSF schema; DO NOT EDIT.

// NetworkTraffic The Network Traffic object describes characteristics of network traffic. Network
// traffic refers to data moving across a network at a given point of time.
type NetworkTraffic struct {
	// The total number of bytes (in and out).
	Bytes int64 `json:"bytes,omitempty"`

	// The number of bytes sent from the destination to the source.
	BytesIn int64 `json:"bytes_in,omitempty"`

	// Indicates the number of bytes missed, which is representative of packet
	// loss.
	BytesMissed int64 `json:"bytes_missed,omitempty"`

	// The number of bytes sent from the source to the destination.
	BytesOut int64 `json:"bytes_out,omitempty"`

	// The total number of chunks (in and out).
	Chunks int64 `json:"chunks,omitempty"`

	// The number of chunks sent from the destination to the source.
	ChunksIn int64 `json:"chunks_in,omitempty"`

	// The number of chunks sent from the source to the destination.
	ChunksOut int64 `json:"chunks_out,omitempty"`

	// The total number of packets (in and out).
	Packets int64 `json:"packets,omitempty"`

	// The number of packets sent from the destination to the source.
	PacketsIn int64 `json:"packets_in,omitempty"`

	// The number of packets sent from the source to the destination.
	PacketsOut int64 `json:"packets_out,omitempty"`
}

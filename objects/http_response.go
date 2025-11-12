package objects

// Code generated from OCSF schema; DO NOT EDIT.

// HttpResponse The HTTP Response object contains detailed information about the response sent
// from a web server to the requester. It encompasses attributes and metadata that
// describe the response status, headers, body content, and other relevant
// information.
type HttpResponse struct {
	// The actual length of the HTTP response body, in number of bytes, independent
	// of a potentially existing Content-Length header.
	BodyLength int64 `json:"body_length,omitempty"`

	// The Hypertext Transfer Protocol (HTTP) status code returned from the web
	// server to the client. For example, 200.
	Code int64 `json:"code,omitempty"`

	// The request header that identifies the original
	// https://www.iana.org/assignments/media-types/media-types.xhtml media type of
	// the resource (prior to any content encoding applied for sending).
	ContentType string `json:"content_type,omitempty"`

	// Additional HTTP headers of an HTTP request or response.
	HttpHeaders []*HttpHeader `json:"http_headers,omitempty"`

	// The HTTP response latency measured in milliseconds.
	Latency int64 `json:"latency,omitempty"`

	// The length of the entire HTTP response, in number of bytes.
	Length int64 `json:"length,omitempty"`

	// The description of the event/finding, as defined by the source.
	Message string `json:"message,omitempty"`

	// The response status. For example: A successful HTTP status of 'OK' which
	// corresponds to a code of 200.
	Status string `json:"status,omitempty"`
}

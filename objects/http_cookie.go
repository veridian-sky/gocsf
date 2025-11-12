package objects

// Code generated from OCSF schema; DO NOT EDIT.

// HttpCookie The HTTP Cookie object, also known as a web cookie or browser cookie, contains
// details and values pertaining to a small piece of data that a server sends to a
// user's web browser. This data is then stored by the browser and sent back to the
// server with subsequent requests, allowing the server to remember and track
// certain information about the user's browsing session or preferences.
type HttpCookie struct {
	// The domain name for the server from which the http_cookie is served.
	Domain string `json:"domain,omitempty"`

	// The expiration time of the HTTP cookie.
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	// The expiration time of the HTTP cookie.
	ExpirationTimeDt string `json:"expiration_time_dt,omitempty"`

	// A cookie attribute to make it inaccessible via JavaScript
	HttpOnly bool `json:"http_only,omitempty"`

	// This attribute prevents the cookie from being accessed via JavaScript.
	IsHTTPOnly bool `json:"is_http_only,omitempty"`

	// The cookie attribute indicates that cookies are sent to the server only when
	// the request is encrypted using the HTTPS protocol.
	IsSecure bool `json:"is_secure,omitempty"`

	// The HTTP cookie name.
	Name string `json:"name,omitempty"`

	// The path of the HTTP cookie.
	Path string `json:"path,omitempty"`

	// The cookie attribute that lets servers specify whether/when cookies are sent
	// with cross-site requests. Values are: Strict, Lax or None
	Samesite string `json:"samesite,omitempty"`

	// The cookie attribute to only send cookies to the server with an encrypted
	// request over the HTTPS protocol.
	Secure bool `json:"secure,omitempty"`

	// The HTTP cookie value.
	Value string `json:"value,omitempty"`
}

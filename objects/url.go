package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Url The Uniform Resource Locator (URL) object describes the characteristics of a
// URL.
type Url struct {
	// The Website categorization names, as defined by category_ids enum values.
	Categories []string `json:"categories,omitempty"`

	// The Website categorization identifiers.
	CategoryIDs []int `json:"category_ids,omitempty"`

	// The domain portion of the URL. For example: example.com in
	// https://sub.example.com.
	Domain string `json:"domain,omitempty"`

	// The URL host as extracted from the URL. For example: www.example.com from
	// www.example.com/download/trouble.
	Hostname string `json:"hostname,omitempty"`

	// The URL path as extracted from the URL. For example: /download/trouble from
	// www.example.com/download/trouble.
	Path string `json:"path,omitempty"`

	// The URL port. For example: 80.
	Port interface{} `json:"port,omitempty"`

	// The query portion of the URL. For example: the query portion of the URL
	// http://www.example.com/search?q=bad&sort=date is q=bad&sort=date.
	QueryString string `json:"query_string,omitempty"`

	// The context in which a resource was retrieved in a web request.
	ResourceType string `json:"resource_type,omitempty"`

	// The scheme portion of the URL. For example: http, https, ftp, or sftp.
	Scheme string `json:"scheme,omitempty"`

	// The subdomain portion of the URL. For example: sub in
	// https://sub.example.com or sub2.sub1 in https://sub2.sub1.example.com.
	Subdomain string `json:"subdomain,omitempty"`

	// The URL string. See RFC 1738. For example:
	// http://www.example.com/download/trouble.exe. Note: The URL path should not
	// populate the URL string.
	UrlString string `json:"url_string,omitempty"`
}

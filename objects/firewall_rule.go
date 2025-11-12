package objects

// Code generated from OCSF schema; DO NOT EDIT.

// FirewallRule The Firewall Rule object represents a specific rule within a firewall policy or
// event. It contains information about a rule's configuration, properties, and
// associated actions that define how network traffic is handled by the firewall.
type FirewallRule struct {
	// The rule category.
	Category string `json:"category,omitempty"`

	// The rule trigger condition for the rule. For example: SQL_INJECTION.
	Condition string `json:"condition,omitempty"`

	// The description of the rule that generated the event.
	Desc string `json:"desc,omitempty"`

	// The rule response time duration, usually used for challenge completion time.
	Duration int64 `json:"duration,omitempty"`

	// The data in a request that rule matched. For example: '["10","and","1"]'.
	MatchDetails []string `json:"match_details,omitempty"`

	// The location of the matched data in the source which resulted in the
	// triggered firewall rule. For example: HEADER.
	MatchLocation string `json:"match_location,omitempty"`

	// The name of the rule that generated the event.
	Name string `json:"name,omitempty"`

	// The rate limit for a rate-based rule.
	RateLimit int64 `json:"rate_limit,omitempty"`

	// The sensitivity of the firewall rule in the matched event. For example:
	// HIGH.
	Sensitivity string `json:"sensitivity,omitempty"`

	// The rule type.
	Type string `json:"type,omitempty"`

	// The unique identifier of the rule that generated the event.
	Uid string `json:"uid,omitempty"`

	// The rule version. For example: 1.1.
	Version string `json:"version,omitempty"`
}

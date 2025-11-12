package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Actor The Actor object contains details about the user, role, application, service, or
// process that initiated or performed a specific activity. Note that Actor is not
// the threat actor of a campaign but may be part of a campaign.
type Actor struct {
	// The client application or service that initiated the activity. This can be
	// in conjunction with the user if present. Note that app_name is distinct from
	// the process if present.
	AppName string `json:"app_name,omitempty"`

	// The unique identifier of the client application or service that initiated
	// the activity. This can be in conjunction with the user if present. Note that
	// app_name is distinct from the process.pid or process.uid if present.
	AppUID string `json:"app_uid,omitempty"`

	// Provides details about an authorization, such as authorization outcome, and
	// any associated policies related to the activity/event.
	Authorizations []*Authorization `json:"authorizations,omitempty"`

	// This object describes details about the Identity Provider used.
	Idp *Idp `json:"idp,omitempty"`

	// The name of the service that invoked the activity as described in the event.
	InvokedBy string `json:"invoked_by,omitempty"`

	// The process that initiated the activity.
	Process *Process `json:"process,omitempty"`

	// The user session from which the activity was initiated.
	Session *Session `json:"session,omitempty"`

	// The user that initiated the activity or the user context from which the
	// activity was initiated.
	User *User `json:"user,omitempty"`
}

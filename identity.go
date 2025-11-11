// Code generated from OCSF schema. DO NOT EDIT.
package gocsf

import (
	"time"
)

// Identity Types

// Account The Account object contains details about the account that initiated or performed a specific activity within a system or application.
type Account struct {
	// The list of labels/tags associated to the account.
	Labels []string `json:"labels,omitempty"`
	// The name of the account (e.g. GCP Account Name).
	Name string `json:"name"`
	// The account type, normalized to the caption of 'account_type_id'. In the case of 'Other', it is defined by the event source.
	Type string `json:"type,omitempty"`
	// The normalized account type identifier.
	TypeId int `json:"type_id"`
	// The unique identifier of the account (e.g. AWS Account ID).
	Uid string `json:"uid"`
}

// Actor The Actor object contains details about the user, role, application, service, or process that initiated or performed a specific activity.
type Actor struct {
	// The client application or service that initiated the activity. This can be in conjunction with the <code>user</code> if present.  Note that <code>app_name</code> is distinct from the <code>process</code> if present.
	AppName string `json:"app_name,omitempty"`
	// The unique identifier of the client application or service that initiated the activity. This can be in conjunction with the <code>user</code> if present. Note that <code>app_name</code> is distinct from the <code>process.pid</code> or <code>process.uid</code> if present.
	AppUid string `json:"app_uid,omitempty"`
	// Provides details about an authorization, such as authorization outcome, and any associated policies related to the activity/event.
	Authorizations []*Authorization `json:"authorizations,omitempty"`
	// This object describes details about the Identity Provider used.
	Idp *Idp `json:"idp,omitempty"`
	// The name of the service that invoked the activity as described in the event.
	InvokedBy string `json:"invoked_by,omitempty"`
	// The process that initiated the activity.
	Process *Process `json:"process"`
	// The user session from which the activity was initiated.
	Session *Session `json:"session,omitempty"`
	// The user that initiated the activity or the user context from which the activity was initiated.
	User *User `json:"user"`
}

// AuthFactor An Authentication Factor object describes a category of methods used for identity verification in an authentication attempt.
type AuthFactor struct {
	// Device used to complete an authentication request.
	Device *Device `json:"device"`
	// The email address used in an email-based authentication factor.
	EmailAddr string `json:"email_addr,omitempty"`
	// The type of authentication factor used in an authentication attempt.
	FactorType string `json:"factor_type"`
	// The normalized identifier for the authentication factor.
	FactorTypeId int `json:"factor_type_id"`
	// Whether the authentication factor is an HMAC-based One-time Password (HOTP).
	IsHotp bool `json:"is_hotp"`
	// Whether the authentication factor is a Time-based One-time Password (TOTP).
	IsTotp bool `json:"is_totp"`
	// The phone number used for a telephony-based authentication request.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The name of provider for an authentication factor.
	Provider string `json:"provider"`
	// The question(s) provided to user for a question-based authentication factor.
	SecurityQuestions []string `json:"security_questions,omitempty"`
}

// Authorization The Authorization Result object provides details about the authorization outcome and associated policies related to activity.
type Authorization struct {
	// Authorization Result/outcome, e.g. allowed, denied.
	Decision string `json:"decision"`
	// Details about the Identity/Access management policies that are applicable.
	Policy *Policy `json:"policy,omitempty"`
}

// Group The Group object represents a collection or association of entities, such as users, policies, or devices. It serves as a logical grouping mechanism to organize and manage entities with similar characteristics or permissions within a system or organization.
type Group struct {
	// The group description.
	Desc string `json:"desc,omitempty"`
	// The domain where the group is defined. For example: the LDAP or Active Directory domain.
	Domain string `json:"domain,omitempty"`
	// The group name.
	Name string `json:"name"`
	// The group privileges.
	Privileges []string `json:"privileges,omitempty"`
	// The type of the group or account.
	Type string `json:"type,omitempty"`
	// The unique identifier of the group. For example, for Windows events this is the security identifier (SID) of the group.
	Uid string `json:"uid"`
}

// Idp The Identity Provider object contains detailed information about a provider responsible for creating, maintaining, and managing identity information while offering authentication services to applications. An Identity Provider (IdP) serves as a trusted authority that verifies the identity of users and issues authentication tokens or assertions to enable secure access to applications or services.
type Idp struct {
	// The name of the identity provider.
	Name string `json:"name"`
	// The unique identifier of the identity provider.
	Uid string `json:"uid"`
}

// LdapPerson The additional LDAP attributes that describe a person.
type LdapPerson struct {
	// The cost center associated with the user.
	CostCenter string `json:"cost_center,omitempty"`
	// The timestamp when the user was created.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The timestamp when the user was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The timestamp when the user was deleted. In Active Directory (AD), when a user is deleted they are moved to a temporary container and then removed after 30 days. So, this field can be populated even after a user is deleted for the next 30 days.
	DeletedTime time.Time `json:"deleted_time,omitempty"`
	// The timestamp when the user was deleted. In Active Directory (AD), when a user is deleted they are moved to a temporary container and then removed after 30 days. So, this field can be populated even after a user is deleted for the next 30 days.
	DeletedTimeDt time.Time `json:"deleted_time_dt,omitempty"`
	// A list of additional email addresses for the user.
	EmailAddrs []string `json:"email_addrs,omitempty"`
	// The employee identifier assigned to the user by the organization.
	EmployeeUid string `json:"employee_uid,omitempty"`
	// The given or first name of the user.
	GivenName string `json:"given_name,omitempty"`
	// The timestamp when the user was or will be hired by the organization.
	HireTime time.Time `json:"hire_time,omitempty"`
	// The timestamp when the user was or will be hired by the organization.
	HireTimeDt time.Time `json:"hire_time_dt,omitempty"`
	// The user's job title.
	JobTitle string `json:"job_title,omitempty"`
	// The labels associated with the user. For example in AD this could be the <code>userType</code>, <code>employeeType</code>. For example: <code>Member, Employee</code>.
	Labels []string `json:"labels,omitempty"`
	// The last time when the user logged in.
	LastLoginTime time.Time `json:"last_login_time,omitempty"`
	// The last time when the user logged in.
	LastLoginTimeDt time.Time `json:"last_login_time_dt,omitempty"`
	// The LDAP and X.500 <code>commonName</code> attribute, typically the full name of the person. For example, <code>John Doe</code>.
	LdapCn string `json:"ldap_cn,omitempty"`
	// The X.500 Distinguished Name (DN) is a structured string that uniquely identifies an entry, such as a user, in an X.500 directory service For example, <code>cn=John Doe,ou=People,dc=example,dc=com</code>.
	LdapDn string `json:"ldap_dn,omitempty"`
	// The timestamp when the user left or will be leaving the organization.
	LeaveTime time.Time `json:"leave_time,omitempty"`
	// The timestamp when the user left or will be leaving the organization.
	LeaveTimeDt time.Time `json:"leave_time_dt,omitempty"`
	// The geographical location associated with a user. This is typically the user's usual work location.
	Location *Location `json:"location,omitempty"`
	// The user's manager. This helps in understanding an org hierarchy. This should only ever be populated once in an event. I.e. there should not be a manager's manager in an event.
	Manager *User `json:"manager,omitempty"`
	// The timestamp when the user entry was last modified.
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	// The timestamp when the user entry was last modified.
	ModifiedTimeDt time.Time `json:"modified_time_dt,omitempty"`
	// The primary office location associated with the user. This could be any string and isn't a specific address. For example, <code>South East Virtual</code>.
	OfficeLocation string `json:"office_location,omitempty"`
	// The last or family name for the user.
	Surname string `json:"surname,omitempty"`
}

// Organization The Organization object describes characteristics of an organization or company and its division if any.
type Organization struct {
	// The name of the organization. For example, Widget, Inc.
	Name string `json:"name"`
	// The name of the organizational unit, within an organization.  For example, Finance, IT, R&D
	OuName string `json:"ou_name"`
	// The alternate identifier for an entity's unique identifier. For example, its Active Directory OU DN or AWS OU ID.
	OuUid string `json:"ou_uid,omitempty"`
	// The unique identifier of the organization. For example, its Active Directory or AWS Org ID.
	Uid string `json:"uid"`
}

// Session The Session object describes details about an authenticated session. e.g. Session Creation Time, Session Issuer. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:Session/'>d3f:Session</a>.
type Session struct {
	// The number of identical sessions spawned from the same source IP, destination IP, application, and content/threat type seen over a period of time.
	Count int `json:"count,omitempty"`
	// The time when the session was created.
	CreatedTime time.Time `json:"created_time"`
	// The time when the session was created.
	CreatedTimeDt time.Time `json:"created_time_dt,omitempty"`
	// The unique identifier of the user's credential. For example, AWS Access Key ID.
	CredentialUid string `json:"credential_uid,omitempty"`
	// The reason which triggered the session expiration.
	ExpirationReason string `json:"expiration_reason,omitempty"`
	// The session expiration time.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// The session expiration time.
	ExpirationTimeDt time.Time `json:"expiration_time_dt,omitempty"`
	// Indicates whether Multi Factor Authentication was used during authentication.
	IsMfa bool `json:"is_mfa,omitempty"`
	// The indication of whether the session is remote.
	IsRemote bool `json:"is_remote"`
	// The indication of whether the session is a VPN session.
	IsVpn bool `json:"is_vpn,omitempty"`
	// The identifier of the session issuer.
	Issuer string `json:"issuer"`
	// The Pseudo Terminal associated with the session. Ex: the tty or pts value.
	Terminal string `json:"terminal,omitempty"`
	// The unique identifier of the session.
	Uid string `json:"uid"`
	// The alternate unique identifier of the session. e.g. AWS ARN - <code>arn:aws:sts::123344444444:assumed-role/Admin/example-session</code>.
	UidAlt string `json:"uid_alt,omitempty"`
	// The universally unique identifier of the session.
	Uuid string `json:"uuid,omitempty"`
}

// User The User object describes the characteristics of a user/person or a security principal. Defined by D3FEND <a target='_blank' href='https://d3fend.mitre.org/dao/artifact/d3f:UserAccount/'>d3f:UserAccount</a>.
type User struct {
	// The user's account or the account associated with the user.
	Account *Account `json:"account,omitempty"`
	// The unique identifier of the user's credential. For example, AWS Access Key ID.
	CredentialUid string `json:"credential_uid,omitempty"`
	// Display Name for the user.
	DisplayName string `json:"display_name,omitempty"`
	// The domain where the user is defined. For example: the LDAP or Active Directory domain.
	Domain string `json:"domain,omitempty"`
	// The user's primary email address.
	EmailAddr string `json:"email_addr,omitempty"`
	// The full name of the person, as per the LDAP Common Name attribute (cn).
	FullName string `json:"full_name,omitempty"`
	// The administrative groups to which the user belongs.
	Groups []*Group `json:"groups,omitempty"`
	// The additional LDAP attributes that describe a person.
	LdapPerson *LdapPerson `json:"ldap_person,omitempty"`
	// The username. For example, <code>janedoe1</code>.
	Name string `json:"name"`
	// Organization and org unit related to the user.
	Org *Organization `json:"org,omitempty"`
	// The risk level, normalized to the caption of the risk_level_id value.
	RiskLevel string `json:"risk_level,omitempty"`
	// The normalized risk level id.
	RiskLevelId int `json:"risk_level_id,omitempty"`
	// The risk score as reported by the event source.
	RiskScore int `json:"risk_score,omitempty"`
	// The type of the user. For example, System, AWS IAM User, etc.
	Type string `json:"type,omitempty"`
	// The account type identifier.
	TypeId int `json:"type_id"`
	// The unique user identifier. For example, the Windows user SID, ActiveDirectory DN or AWS user ARN.
	Uid string `json:"uid"`
	// The alternate user identifier. For example, the Active Directory user GUID or AWS user Principal ID.
	UidAlt string `json:"uid_alt,omitempty"`
}

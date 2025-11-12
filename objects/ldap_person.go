package objects

// Code generated from OCSF schema; DO NOT EDIT.

// LdapPerson The additional LDAP attributes that describe a person.
type LdapPerson struct {
	// The cost center associated with the user.
	CostCenter string `json:"cost_center,omitempty"`

	// The timestamp when the user was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The timestamp when the user was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The timestamp when the user was deleted. In Active Directory (AD), when a
	// user is deleted they are moved to a temporary container and then removed
	// after 30 days. So, this field can be populated even after a user is deleted
	// for the next 30 days.
	DeletedTime int64 `json:"deleted_time,omitempty"`

	// The timestamp when the user was deleted. In Active Directory (AD), when a
	// user is deleted they are moved to a temporary container and then removed
	// after 30 days. So, this field can be populated even after a user is deleted
	// for the next 30 days.
	DeletedTimeDt string `json:"deleted_time_dt,omitempty"`

	// The display name of the LDAP person. According to RFC 2798, this is the
	// preferred name of a person to be used when displaying entries.
	DisplayName string `json:"display_name,omitempty"`

	// A list of additional email addresses for the user.
	EmailAddrs []string `json:"email_addrs,omitempty"`

	// The employee identifier assigned to the user by the organization.
	EmployeeUID string `json:"employee_uid,omitempty"`

	// The given or first name of the user.
	GivenName string `json:"given_name,omitempty"`

	// The timestamp when the user was or will be hired by the organization.
	HireTime int64 `json:"hire_time,omitempty"`

	// The timestamp when the user was or will be hired by the organization.
	HireTimeDt string `json:"hire_time_dt,omitempty"`

	// The user's job title.
	JobTitle string `json:"job_title,omitempty"`

	// The labels associated with the user. For example in AD this could be the
	// userType, employeeType. For example: Member, Employee.
	Labels []string `json:"labels,omitempty"`

	// The last time when the user logged in.
	LastLoginTime int64 `json:"last_login_time,omitempty"`

	// The last time when the user logged in.
	LastLoginTimeDt string `json:"last_login_time_dt,omitempty"`

	// The LDAP and X.500 commonName attribute, typically the full name of the
	// person. For example, John Doe.
	LdapCn string `json:"ldap_cn,omitempty"`

	// The X.500 Distinguished Name (DN) is a structured string that uniquely
	// identifies an entry, such as a user, in an X.500 directory service For
	// example, cn=John Doe,ou=People,dc=example,dc=com.
	LdapDn string `json:"ldap_dn,omitempty"`

	// The timestamp when the user left or will be leaving the organization.
	LeaveTime int64 `json:"leave_time,omitempty"`

	// The timestamp when the user left or will be leaving the organization.
	LeaveTimeDt string `json:"leave_time_dt,omitempty"`

	// The geographical location associated with a user. This is typically the
	// user's usual work location.
	Location *Location `json:"location,omitempty"`

	// The user's manager. This helps in understanding an org hierarchy. This
	// should only ever be populated once in an event. I.e. there should not be a
	// manager's manager in an event.
	Manager *User `json:"manager,omitempty"`

	// The timestamp when the user entry was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The timestamp when the user entry was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The primary office location associated with the user. This could be any
	// string and isn't a specific address. For example, South East Virtual.
	OfficeLocation string `json:"office_location,omitempty"`

	// The telephone number of the user. Corresponds to the LDAP Telephone-Number
	// CN.
	PhoneNumber string `json:"phone_number,omitempty"`

	// The last or family name for the user.
	Surname string `json:"surname,omitempty"`

	// The list of tags; {key:value} pairs associated to the user.
	Tags []*KeyValueObject `json:"tags,omitempty"`
}

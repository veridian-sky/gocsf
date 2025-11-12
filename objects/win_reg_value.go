package objects

// Code generated from OCSF schema; DO NOT EDIT.

// WinRegValue The registry value object describes a Windows registry value.
type WinRegValue struct {
	// The data of the registry value. Where the value type is known, implementers
	// should instead use a type-specific attribute, i.e. reg_binary_data,
	// reg_integer_data, reg_string_data, or reg_string_list_data.
	Data interface{} `json:"data,omitempty"`

	// The indication of whether the value is from a default value name. For
	// example, the value name could be missing.
	IsDefault bool `json:"is_default,omitempty"`

	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`

	// The time when the registry value was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the registry value was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The name of the registry value.
	Name string `json:"name,omitempty"`

	// The full path to the registry key, where the value is located.
	Path string `json:"path,omitempty"`

	// The data of the registry value when type_id is REG_BINARY or REG_NONE.
	RegBinaryData interface{} `json:"reg_binary_data,omitempty"`

	// The data of the registry value when type_id is REG_DWORD,
	// REG_DWORD_BIG_ENDIAN, or REG_QWORD.
	RegIntegerData int64 `json:"reg_integer_data,omitempty"`

	// The data of the registry value when type_id is REG_SZ, REG_EXPAND_SZ, or
	// REG_LINK.
	RegStringData string `json:"reg_string_data,omitempty"`

	// The data of the registry value when type_id is REG_MULTI_SZ.
	RegStringListData []string `json:"reg_string_list_data,omitempty"`

	// A string representation of the value type as specified in
	// https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry-value-types
	// Registry Value Types.
	Type string `json:"type,omitempty"`

	// The value type ID.
	TypeID int `json:"type_id,omitempty"`
}

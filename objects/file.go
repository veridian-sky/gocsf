package objects

// Code generated from OCSF schema; DO NOT EDIT.

// File The File object represents the metadata associated with a file stored in a
// computer system. It encompasses information about the file itself, including its
// attributes, properties, and organizational details.
type File struct {
	// The time when the file was last accessed.
	AccessedTime int64 `json:"accessed_time,omitempty"`

	// The time when the file was last accessed.
	AccessedTimeDt string `json:"accessed_time_dt,omitempty"`

	// The name of the user who last accessed the object.
	Accessor *User `json:"accessor,omitempty"`

	// The bitmask value that represents the file attributes.
	Attributes int64 `json:"attributes,omitempty"`

	// The name of the company that published the file. For example: Microsoft
	// Corporation.
	CompanyName string `json:"company_name,omitempty"`

	// The file content confidentiality, normalized to the confidentiality_id
	// value. In the case of 'Other', it is defined by the event source.
	Confidentiality string `json:"confidentiality,omitempty"`

	// The normalized identifier of the file content confidentiality indicator.
	ConfidentialityID int `json:"confidentiality_id,omitempty"`

	// The time when the file was created.
	CreatedTime int64 `json:"created_time,omitempty"`

	// The time when the file was created.
	CreatedTimeDt string `json:"created_time_dt,omitempty"`

	// The user that created the file.
	Creator *User `json:"creator,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// The description of the file, as returned by file system. For example: the
	// description as returned by the Unix file command or the Windows file type.
	Desc string `json:"desc,omitempty"`

	// The drive type, normalized to the caption of the drive_type_id value. In the
	// case of Other, it is defined by the source.
	DriveType string `json:"drive_type,omitempty"`

	// Identifies the type of a disk drive, i.e. fixed, removable, etc.
	DriveTypeID int `json:"drive_type_id,omitempty"`

	// The encryption details of the file. Should be populated if the file is
	// encrypted.
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`

	// The extension of the file, excluding the leading dot. For example: exe from
	// svchost.exe, or gz from export.tar.gz.
	Ext string `json:"ext,omitempty"`

	// An array of hash attributes.
	Hashes []*Fingerprint `json:"hashes,omitempty"`

	// The name of the file as identified within the file itself. This contrasts
	// with the name by which the file is known on disk. Where available, the
	// internal name is widely used by security practitioners and detection content
	// because the on-disk file name is not reliable. On the Windows OS, most PE
	// files contain a <a
	// href="https://learn.microsoft.com/en-us/windows/win32/menurc/versioninfo-resource">VERSIONINFO
	// resource from which the internal name can be obtained. On macOS, binaries
	// can optionally embed a copy of the application's Info.plist file which in
	// turn contains the name of the executable.
	InternalName string `json:"internal_name,omitempty"`

	// Indicates if the file was deleted from the filesystem.
	IsDeleted bool `json:"is_deleted,omitempty"`

	// Indicates if the file is encrypted.
	IsEncrypted bool `json:"is_encrypted,omitempty"`

	// Indicates if the file is publicly accessible. For example in an object's
	// public access in AWS S3
	IsPublic bool `json:"is_public,omitempty"`

	// Indicates that the file cannot be modified.
	IsReadonly bool `json:"is_readonly,omitempty"`

	// The indication of whether the object is part of the operating system.
	IsSystem bool `json:"is_system,omitempty"`

	// The Multipurpose Internet Mail Extensions (MIME) type of the file, if
	// applicable.
	MimeType string `json:"mime_type,omitempty"`

	// The time when the file was last modified.
	ModifiedTime int64 `json:"modified_time,omitempty"`

	// The time when the file was last modified.
	ModifiedTimeDt string `json:"modified_time_dt,omitempty"`

	// The user that last modified the file.
	Modifier *User `json:"modifier,omitempty"`

	// The name of the file. For example: svchost.exe
	Name interface{} `json:"name,omitempty"`

	// The user that owns the file/object.
	Owner *User `json:"owner,omitempty"`

	// The parent folder in which the file resides. For example:
	// c:\windows\system32
	ParentFolder string `json:"parent_folder,omitempty"`

	// The full path to the file. For example: c:\windows\system32\svchost.exe.
	Path interface{} `json:"path,omitempty"`

	// The product that created or installed the file.
	Product *Product `json:"product,omitempty"`

	// The object security descriptor.
	SecurityDescriptor string `json:"security_descriptor,omitempty"`

	// The digital signature of the file.
	Signature *DigitalSignature `json:"signature,omitempty"`

	// The size of data, in bytes.
	Size int64 `json:"size,omitempty"`

	// The storage class of the file. For example in AWS S3: STANDARD, STANDARD_IA,
	// GLACIER.
	StorageClass string `json:"storage_class,omitempty"`

	// The list of tags; {key:value} pairs associated to the file.
	Tags []*KeyValueObject `json:"tags,omitempty"`

	// The file type.
	Type string `json:"type,omitempty"`

	// The file type ID. Note the distinction between a Regular File and an
	// Executable File. If the distinction is not known, or not indicated by the
	// log, use Regular File. In this case, it should not be assumed that a Regular
	// File is not executable.
	TypeID int `json:"type_id,omitempty"`

	// The unique identifier of the file as defined by the storage system, such the
	// file system file ID.
	Uid string `json:"uid,omitempty"`

	// The file URI, such as those reporting by static analysis tools. E.g.,
	// file:///C:/dev/sarif/sarif-tutorials/samples/Introduction/simple-example.js
	Uri string `json:"uri,omitempty"`

	// The URL of the file, when applicable.
	Url *Url `json:"url,omitempty"`

	// The file version. For example: 8.0.7601.17514.
	Version string `json:"version,omitempty"`

	// The volume on the storage device where the file is located.
	Volume string `json:"volume,omitempty"`

	// An unordered collection of zero or more name/value pairs where each pair
	// represents a file or folder extended attribute.</p>For example: Windows
	// alternate data stream attributes (ADS stream name, ADS size, etc.),
	// user-defined or application-defined attributes, ACL, owner, primary group,
	// etc. Examples from DCS:
	// </p><ul><li>ads_name</li><li>ads_size</li><li>dacl</li><li>owner</li><li>primary_group</li><li>link_name
	// - name of the link associated to the file.</li><li>hard_link_count - the
	// number of links that are associated to the file.</li></ul>
	Xattributes *Object `json:"xattributes,omitempty"`
}

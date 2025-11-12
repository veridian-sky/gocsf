package types

// Code generated from OCSF schema; DO NOT EDIT.

// BooleanT Boolean value. One of true or false.
type BooleanT interface{}

// BytestringT Base 64 encoded immutable byte sequence. Traditional Base 64 is preferred but
// publishers may use URL-safe Base 64 when known to be acceptable to consumers.
// These encodings are described in RFC 4648.
type BytestringT string

// DatetimeT The Internet Date/Time format as defined in
// https://www.rfc-editor.org/rfc/rfc3339.html RFC-3339. For example:
// 2024-09-10T23:20:50.520Z, 2024-09-10 23:20:50.520789Z.
type DatetimeT string

// EmailT Email address. For example: john_doe@example.com.
type EmailT string

// FileHashT Hash. A unique value that corresponds to the content of the file, image,
// ja3_hash or hassh found in the schema. For example: MD5:
// 3172ac7e2b55cbb81f04a6e65855a628.
type FileHashT string

// FileNameT File name. For example: text-file.txt.
type FileNameT string

// FilePathT The full path to the file. For example: For example:
// c:\windows\system32\svchost.exe.
type FilePathT string

// FloatT Real floating-point value. For example: 3.14.
type FloatT interface{}

// HostnameT Unique name assigned to a device connected to a computer network. It may be a
// fully qualified domain name (FQDN). For example: r2-d2.example.com.,
// mx.example.com
type HostnameT string

// IntegerT Signed integer value.
type IntegerT interface{}

// IpT Internet Protocol address (IP address), in either IPv4 or IPv6 format. For
// example: 192.168.200.24, 2001:0db8:85a3:0000:0000:8a2e:0370:7334.
type IpT string

// JsonT Embedded JSON value. A value can be a string, or a number, or true or false or
// null, or an object or an array. These structures can be nested. See
// https://www.json.org www.json.org.
type JsonT interface{}

// LongT 8-byte long, signed integer value.
type LongT interface{}

// MacT Media Access Control (MAC) address. For example: 18:36:F3:98:4F:9A.
type MacT string

// PortT The TCP/UDP port number. For example: 80, 22.
type PortT int64

// ProcessNameT Process name. For example: Notepad.
type ProcessNameT string

// RegKeyPathT Full path of registry key.
type RegKeyPathT string

// ResourceUIDT Resource unique identifier. For example, S3 Bucket name or EC2 Instance ID.
type ResourceUIDT string

// StringT UTF-8 encoded byte sequence.
type StringT interface{}

// SubnetT The subnet represented in a CIDR notation, using the format
// network_address/prefix_length. The network_address can be in either IPv4 or IPv6
// format. The prefix length indicates the number of bits used for the network
// portion, and the remaining bits are available for host addresses within that
// subnet. For example: 192.168.1.0/24, 2001:0db8:85a3:0000::/64
type SubnetT string

// TimestampT The timestamp format is the number of milliseconds since the Epoch 01/01/1970
// 00:00:00 UTC. For example: 1618524549901.
type TimestampT int64

// UrlT Uniform Resource Locator (URL) string. For example:
// http://www.example.com/download/trouble.exe.
type UrlT string

// UsernameT User name. For example: john_doe.
type UsernameT string

// UuidT 128-bit universal unique identifier. For example:
// 123e4567-e89b-12d3-a456-42661417400.
type UuidT string

package objects

// Code generated from OCSF schema; DO NOT EDIT.

// Email The Email object describes the email metadata such as sender, recipients, and
// direction, and can include embedded URLs and files.
type Email struct {
	// The machine-readable email header Cc values, as defined by RFC 5322. For
	// example example.user@usersdomain.com.
	Cc []string `json:"cc,omitempty"`

	// The human-readable email header Cc Mailbox values. For example 'Example User
	// &lt;example.user@usersdomain.com&gt;'.
	CcMailboxes []string `json:"cc_mailboxes,omitempty"`

	// The Data Classification object includes information about data
	// classification levels and data category types.
	DataClassification *DataClassification `json:"data_classification,omitempty"`

	// A list of Data Classification objects, that include information about data
	// classification levels and data category types, identified by a classifier.
	DataClassifications []*DataClassification `json:"data_classifications,omitempty"`

	// The machine-readable Delivered-To email header field. For example
	// example.user@usersdomain.com
	DeliveredTo string `json:"delivered_to,omitempty"`

	// The machine-readable Delivered-To email header values. For example
	// example.user@usersdomain.com
	DeliveredToList []string `json:"delivered_to_list,omitempty"`

	// The files embedded or attached to the email.
	Files []*File `json:"files,omitempty"`

	// The machine-readable email header From value, as defined by RFC 5322. For
	// example example.user@usersdomain.com.
	From string `json:"from,omitempty"`

	// The machine-readable email header From values. This array should contain the
	// value in from. For example example.user@usersdomain.com.
	FromList []string `json:"from_list,omitempty"`

	// The human-readable email header From Mailbox value. For example 'Example
	// User &lt;example.user@usersdomain.com&gt;'.
	FromMailbox string `json:"from_mailbox,omitempty"`

	// The human-readable email header From Mailbox values. This array should
	// contain the value in from_mailbox. For example 'Example User
	// &lt;example.user@usersdomain.com&gt;'.
	FromMailboxes []string `json:"from_mailboxes,omitempty"`

	// Additional HTTP headers of an HTTP request or response.
	HttpHeaders []*HttpHeader `json:"http_headers,omitempty"`

	// The indication of whether the email has been read.
	IsRead bool `json:"is_read,omitempty"`

	// The email header Message-ID value, as defined by RFC 5322.
	MessageUID string `json:"message_uid,omitempty"`

	// The email authentication header.
	RawHeader string `json:"raw_header,omitempty"`

	// The machine-readable email header Reply-To value, as defined by RFC 5322.
	// For example example.user@usersdomain.com
	ReplyTo string `json:"reply_to,omitempty"`

	// The machine-readable email header Reply-To values, as defined by RFC 5322.
	// For example example.user@usersdomain.com
	ReplyToList []string `json:"reply_to_list,omitempty"`

	// The human-readable email header Reply To Mailbox values. For example
	// 'Example User &lt;example.user@usersdomain.com&gt;'.
	ReplyToMailboxes []string `json:"reply_to_mailboxes,omitempty"`

	// The address found in the 'Return-Path' header, which indicates where bounce
	// messages (non-delivery reports) should be sent. This address is often set by
	// the sending system and may differ from the 'From' or 'Sender' addresses. For
	// example, mailer-daemon@senderserver.com.
	ReturnPath string `json:"return_path,omitempty"`

	// The machine readable email address of the system or server that actually
	// transmitted the email message, extracted from the email headers per RFC
	// 5322. This differs from the from field, which shows the message author. The
	// sender field is most commonly used when multiple addresses appear in the
	// from_list field, or when the transmitting system is different from the
	// message author (such as when sending on behalf of someone else).
	Sender string `json:"sender,omitempty"`

	// The human readable email address of the system or server that actually
	// transmitted the email message, extracted from the email headers per RFC
	// 5322. This differs from the from_mailbox field, which shows the message
	// author. The sender mailbox field is most commonly used when multiple
	// addresses appear in the from_mailboxes field, or when the transmitting
	// system is different from the message author (such as when sending on behalf
	// of someone else).
	SenderMailbox string `json:"sender_mailbox,omitempty"`

	// The size in bytes of the email, including attachments.
	Size int64 `json:"size,omitempty"`

	// The value of the SMTP MAIL FROM command.
	SmtpFrom string `json:"smtp_from,omitempty"`

	// The value of the SMTP envelope RCPT TO command.
	SmtpTo []string `json:"smtp_to,omitempty"`

	// The email header Subject value, as defined by RFC 5322.
	Subject string `json:"subject,omitempty"`

	// The machine-readable email header To values, as defined by RFC 5322. For
	// example example.user@usersdomain.com
	To []string `json:"to,omitempty"`

	// The human-readable email header To Mailbox values. For example 'Example User
	// &lt;example.user@usersdomain.com&gt;'.
	ToMailboxes []string `json:"to_mailboxes,omitempty"`

	// The unique identifier of the email thread.
	Uid string `json:"uid,omitempty"`

	// The URLs embedded in the email.
	Urls []*Url `json:"urls,omitempty"`

	// The X-Originating-IP header identifying the emails originating IP
	// address(es).
	XOriginatingIP []string `json:"x_originating_ip,omitempty"`
}

package goinblue

// Email request to be send
type Email struct {
	To          []map[string]string `json:"to"`
	Subject     string            `json:"subject"`
	Sender        map[string]string          `json:"sender"`
	HtmlContent        string            `json:"htmlContent,omitempty"`
	TextContent        string            `json:"textContent,omitempty"`
	Cc          []map[string]string `json:"cc,omitempty"`
	Bcc         []map[string]string `json:"bcc,omitempty"`
	ReplyTo     map[string]string          `json:"replyto,omitempty"`
	Attachment  interface{}       `json:"attachment,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	InlineImage map[string]string `json:"inline_image,omitempty"`
}

// Email template request to be send
type EmailTemplate struct {
	Id            int               `json:"id"`
	To            []map[string]string `json:"to"`
	Cc            []map[string]string `json:"cc,omitempty"`
	Bcc           []map[string]string `json:"bcc,omitempty"`
	Attr          map[string]string `json:"attr,omitempty"`
	AttachmentUrl []string          `json:"attachment_url,omitempty"`
	Attachment    map[string]string `json:"attachment,omitempty"`
	Headers       map[string]string `json:"headers,omitempty"`
}

// This is here for documentation purpose
type Attachment map[string]string

// This is here for documentation purpose
type AttachmentUrl []string

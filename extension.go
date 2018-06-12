package smtp

// Extension is an interface for implementing SMTP extensions.
type Extension interface {
	// EHLOKeyword is the name of the extension, to be returned in the EHLO response.
	// See RFC 5321, section 2.2.2: https://tools.ietf.org/html/rfc5321#section-2.2.2
	EHLOKeyword() string

	// Process is called for each command. If a reply is returned, the reply is returned to the
	// client and processing is ceased.
	Process(proto *Protocol, verb, args string) (errorReply *Reply)

	// TLSOnly returns true if this extension should only be called, or even shown in the EHLO
	// reponse if the connection has been upgraded to TLS.
	TLSOnly() bool
}

package smtp

import (
	"fmt"
	"strings"
	"unicode"
)

type smtpUTF8Key int

const (
	clientUTF8Status smtpUTF8Key = iota + 1
)

type smtpUTF8 struct{}

func (ex *smtpUTF8) EHLOKeyword() string {
	return "SMTPUTF8"
}

func (ex *smtpUTF8) TLSOnly() bool {
	return false
}

func (ex *smtpUTF8) Process(proto *Protocol, verb, args string) *Reply {
	switch verb {
	case "MAIL":
		for _, part := range strings.Split(args, " ") {
			if part != "SMTPUTF8" {
				continue
			}

			proto.ExtensionData[clientUTF8Status] = struct{}{}

			return nil
		}

	case "RCPT":
		rcpt, err := proto.ParseRCPT(args)
		if err != nil {
			return nil
		}

		if _, exists := proto.ExtensionData[clientUTF8Status]; !exists && ex.IsNotASCII(rcpt) {
			return ReplySyntaxError(fmt.Sprintf("unexpected non-ASCII address: %s", rcpt))
		}

		return nil
	}

	return nil
}

func (*smtpUTF8) IsNotASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return true
		}
	}

	return false
}

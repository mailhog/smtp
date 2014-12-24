MailHog SMTP Protocol [![GoDoc](https://godoc.org/github.com/mailhog/smtp?status.svg)](https://godoc.org/github.com/mailhog/smtp) [![Build Status](https://travis-ci.org/mailhog/smtp.svg?branch=master)](https://travis-ci.org/mailhog/smtp)
=========

`github.com/mailhog/smtp` implements an SMTP server state machine.

  * ESMTP server implementing RFC5321
  * Support for SMTP AUTH (RFC4954) and PIPELINING (RFC2920)

```go
proto := NewProtocol()
reply := proto.Start()
reply = proto.ProcessCommand("EHLO localhost")
// ...
```

### Licence

Copyright ©‎ 2014, Ian Kent (http://iankent.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

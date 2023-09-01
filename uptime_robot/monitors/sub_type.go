package monitors

type SubType int64

const (
	SubTypeHttp       SubType = 1
	SubTypeHttps      SubType = 2
	SubTypeFtp        SubType = 3
	SubTypeSmtp       SubType = 4
	SubTypePop3       SubType = 5
	SubTypeImap       SubType = 6
	SubTypeCustomPort SubType = 99
)

func (st SubType) String() string {
	switch st {
	case SubTypeHttp:
		return "HTTP"
	case SubTypeHttps:
		return "HTTPS"
	case SubTypeFtp:
		return "FTP"
	case SubTypeSmtp:
		return "SMTP"
	case SubTypePop3:
		return "POP3"
	case SubTypeImap:
		return "IMAP"
	case SubTypeCustomPort:
		return "Custom Port"
	default:
		return ""
	}
}

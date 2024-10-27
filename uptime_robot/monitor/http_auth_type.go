package monitor

type HttpAuthType int64

const (
	HttpAuthTypeNone      HttpAuthType = 0
	HttpAuthTypeHttpBasic HttpAuthType = 1
	HttpAuthTypeDigest    HttpAuthType = 2
)

func (hat HttpAuthType) String() string {
	switch hat {
	case HttpAuthTypeNone:
		return "None"
	case HttpAuthTypeHttpBasic:
		return "HTTP Basic Auth"
	case HttpAuthTypeDigest:
		return "Digest"
	default:
		return "Unknown"
	}
}

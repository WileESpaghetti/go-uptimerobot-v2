package monitors

type HttpAuthType int64

const (
	HttpAuthTypeHttpBasic HttpAuthType = 1
	HttpAuthTypeDigest    HttpAuthType = 2
)

func (hat HttpAuthType) String() string {
	switch hat {
	case HttpAuthTypeHttpBasic:
		return "HTTP Basic Auth"
	case HttpAuthTypeDigest:
		return "Digest"
	default:
		return "Unknown"
	}
}

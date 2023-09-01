package monitors

type Type int64

const (
	TypeHttp    Type = 1
	TypeKeyword Type = 2
	TypePing    Type = 3
	TypePort    Type = 4
)

func (t Type) String() string {
	switch t {
	case TypeHttp:
		return "HTTP"
	case TypeKeyword:
		return "Keyword"
	case TypePing:
		return "Ping"
	case TypePort:
		return "Port"
	default:
		return "Unknown"
	}
}

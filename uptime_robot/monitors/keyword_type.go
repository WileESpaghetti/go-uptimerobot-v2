package monitors

// KeywordType is an Enum
type KeywordType int64

const (
	KeywordTypeExists    KeywordType = 1
	KeywordTypeNotExists KeywordType = 2
)

func (kt KeywordType) String() string {
	switch kt {
	case KeywordTypeExists:
		return "Exists"
	case KeywordTypeNotExists:
		return "Not Exists"
	default:
		return ""
	}
}

///////////////////////////////////////////////////////////

// KeywordCaseType is whether a keyword is case-sensitive (0) or case-insensitive (1)
type KeywordCaseType int64

const (
	KeywordCaseSensitive   KeywordCaseType = 0
	KeywordCaseInsensitive KeywordCaseType = 1
)

func (kct KeywordCaseType) String() string {
	switch kct {
	case KeywordCaseSensitive:
		return "Case Sensitive"
	case KeywordCaseInsensitive:
		return "Case Insensitive"
	default:
		return ""
	}
}

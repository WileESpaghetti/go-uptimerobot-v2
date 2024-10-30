package public_status_page

type Sort int64

const (
	SortByFriendlyNameAsc  Sort = 1
	SortByFriendlyNameDesc Sort = 2
	SortUpDownPaused       Sort = 3
	SortDownUpPaused       Sort = 4
)

func (s Sort) String() string {
	switch s {
	case SortByFriendlyNameAsc:
		return "Name (A-Z)"
	case SortByFriendlyNameDesc:
		return "Name (Z-A)"
	case SortUpDownPaused:
		return "Up, Down, Paused"
	case SortDownUpPaused:
		return "Down, Up, Paused"
	default:
		return ""
	}
}

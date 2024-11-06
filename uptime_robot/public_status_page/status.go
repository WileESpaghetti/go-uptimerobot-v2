package public_status_page

type Status int64

const (
	StatusPaused Status = 0
	StatusActive Status = 1
)

func (st Status) String() string {
	switch st {
	case StatusPaused:
		return "Paused"
	case StatusActive:
		return "Active"
	default:
		return ""
	}
}

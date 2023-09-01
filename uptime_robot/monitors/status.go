package monitors

type Status int64

const (
	MonitorStatusPassed     Status = 0
	MonitorStatusNotChecked Status = 1
	MonitorStatusUp         Status = 2
	MonitorStatusSeemsDown  Status = 8
	MonitorStatusDown       Status = 9
)

func (st Status) String() string {
	switch st {
	case MonitorStatusPassed:
		return "Passed"
	case MonitorStatusNotChecked:
		return "Not Checked"
	case MonitorStatusUp:
		return "Up"
	case MonitorStatusSeemsDown:
		return "Seems Down"
	case MonitorStatusDown:
		return "Down"
	default:
		return "Unknown"

	}
}

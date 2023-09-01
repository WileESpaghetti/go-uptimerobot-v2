package models

const (
	MonitorStatusPassed     = 0
	MonitorStatusNotChecked = 1
	MonitorStatusUp         = 2
	MonitorStatusSeemsDown  = 8
	MonitorStatusDown       = 9
)

type MonitorStatus int

func (ms MonitorStatus) String() string {
	if ms == MonitorStatusPassed {
		return "Passed"
	} else if ms == MonitorStatusNotChecked {
		return "Not Checked"
	} else if ms == MonitorStatusUp {
		return "Up"
	} else if ms == MonitorStatusSeemsDown {
		return "Seems Down"
	} else if ms == MonitorStatusDown {
		return "Down"
	} else {
		return "Unknown"
	}
}

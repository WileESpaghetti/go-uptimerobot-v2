package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitors"
	"strings"
)

func readAsCSV(val string) ([]string, error) {
	if val == "" {
		return []string{}, nil
	}

	stringReader := strings.NewReader(val)
	csvReader := csv.NewReader(stringReader)
	return csvReader.Read()
}

// TODO is there any benefit from implmenting the pflags.SliceValue interface?
//  func (mtf *monitorTypesFlag) Append(val string) error
//  func (mtf *monitorTypesFlag) Replace(val []string) error
//  func (mtf *monitorTypesFlag) GetSlice() []string

type monitorTypesFlag struct {
	value   *monitors.Types
	changed bool
}

func newMonitorTypesFlag(val monitors.Types, p *monitors.Types) *monitorTypesFlag {
	mtf := &monitorTypesFlag{}
	mtf.value = p
	*mtf.value = val
	return mtf
}

func (mtf *monitorTypesFlag) String() string {
	return mtf.value.String()
}

func (mtf *monitorTypesFlag) Set(val string) error {
	val = strings.Replace(val, "-", ",", -1) // normalize the weird list format used by the UptimeRobot API

	v, err := readAsCSV(val)
	if err != nil {
		return err
	}

	// validate/convert raw values to monitor types
	types := make(monitors.Types, len(v))
	for i, t := range v {
		types[i], err = monitors.NewType(t)
		if err != nil {
			return fmt.Errorf("%s is an unsupported monitor type", t)
		}
	}

	if !mtf.changed {
		*mtf.value = types
	} else {
		*mtf.value = append(*mtf.value, types...)
	}
	mtf.changed = true

	return nil
}

func (mtf *monitorTypesFlag) Type() string {
	return "monitorTypesFlag"
}

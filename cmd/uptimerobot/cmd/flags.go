package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/monitor"
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

// TODO there are 2 types of flags I want:
//  * integer lists separated by dashes or commas ex. 1-299-43 = 1,299,43
//  * list could also be a case insensitive string slice or represented strings separated by dashes: ex. http-https = http, https
// TODO is there any benefit from implmenting the pflags.SliceValue interface?
//  it might allow us to use a generic flag type
//  func (mtf *monitorTypesFlag) Append(val string) error
//  func (mtf *monitorTypesFlag) Replace(val []string) error
//  func (mtf *monitorTypesFlag) GetSlice() []string

func newMonitorTypesFlag(val monitor.Types, p *monitor.Types) *separatedValueFlag[monitor.Type] {
	return newSeparatedValueFlag[monitor.Type](val, p, monitor.NewType)
}

func newMonitorStatusesFlag(val monitor.Statuses, p *monitor.Statuses) *separatedValueFlag[monitor.Status] {
	return newSeparatedValueFlag[monitor.Status](val, p, monitor.NewStatus)
}

func newMonitorLogTypesFlag(val monitor.LogTypes, p *monitor.LogTypes) *separatedValueFlag[monitor.LogType] {
	return newSeparatedValueFlag[monitor.LogType](val, p, monitor.NewLogType)
}

////////////////////////////

type sliceable[T any] interface {
	// FIXME this conflicts with the flag slice interface
	GetSlice() *[]T
}

type valueConverter[T any] func(s string) (T, error)

type separatedValueFlag[T any] struct {
	value     *[]T
	converter valueConverter[T]
	changed   bool
}

func newSeparatedValueFlag[T any](val []T, p sliceable[T], converter valueConverter[T]) *separatedValueFlag[T] {
	msf := &separatedValueFlag[T]{}
	msf.converter = converter
	msf.value = p.GetSlice()
	*msf.value = val
	return msf
}

func (svf *separatedValueFlag[T]) String() string {
	return "FIXME concat values" // FIXME
}

func (svf *separatedValueFlag[T]) Set(val string) error {
	val = strings.Replace(val, "-", ",", -1) // normalize the weird list format used by the UptimeRobot API

	v, err := readAsCSV(val)
	if err != nil {
		return err
	}

	// validate/convert raw values to monitor statuses
	statuses := make([]T, len(v))
	for i, t := range v {
		statuses[i], err = svf.converter(t)
		if err != nil {
			return fmt.Errorf("%s is an unsupported monitor status", t) // FIXME wrap error
		}
	}

	if !svf.changed {
		*svf.value = statuses
	} else {
		*svf.value = append(*svf.value, statuses...)
	}
	svf.changed = true

	return nil
}

func (svf *separatedValueFlag[T]) Type() string {
	return "FIXME get reflected type of T" // FIXME
}

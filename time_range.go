package types

import (
	"errors"
	"strings"
	"time"
)

type TimeRange struct {
	Left        *time.Time
	Right       *time.Time
	LeftClosed  bool
	RightClosed bool
}

// ParseTimeRange parse string and return a *TimeRange.
func ParseTimeRange(rangeStr string, layout string, timeLoc *time.Location, defaultLeftVal string, defaultRightVal string) (*TimeRange, error) {
	if len(rangeStr) < 2 {
		return nil, errors.New("Not a range!")
	}
	startFlag := rangeStr[0]
	if startFlag != '[' && startFlag != '(' {
		return nil, errors.New("Start flag must be '[' or '(' !")
	}
	endFlag := rangeStr[len(rangeStr)-1]
	if endFlag != ']' && endFlag != ')' {
		return nil, errors.New("End flag must be ']' or ')' !")
	}
	rangeVal := rangeStr[1 : len(rangeStr)-1]

	fields := strings.Split(rangeVal, ",")
	if len(fields) != 2 {
		return nil, errors.New("Not a range!")
	}

	var v1, v2 time.Time
	if fields[0] == "" {
		t, err := time.ParseInLocation(layout, defaultLeftVal, timeLoc)
		if err != nil {
			return nil, errors.New("Default left value parse failed!")
		}
		v1 = t
	} else {
		t, err := time.ParseInLocation(layout, fields[0], timeLoc)
		if err != nil {
			return nil, errors.New("Left value parse failed!")
		}
		v1 = t
	}
	if fields[1] == "" {
		t, err := time.ParseInLocation(layout, defaultRightVal, timeLoc)
		if err != nil {
			return nil, errors.New("Default right value parse failed!")
		}
		v2 = t
	} else {
		t, err := time.ParseInLocation(layout, fields[1], timeLoc)
		if err != nil {
			return nil, errors.New("Right value parse failed!")
		}
		v2 = t
	}

	if v1.After(v2) {
		return nil, errors.New("Left should not large than right!")
	}

	return &TimeRange{&v1, &v2, startFlag == '[', endFlag == ']'}, nil
}

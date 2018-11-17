package types

import (
	"errors"
	"strconv"
	"strings"
)

type Uint64Range struct {
	Left        uint64
	Right       uint64
	LeftClosed  bool
	RightClosed bool
}

// ParseUint64Range parse string and return a *Uint64Range.
func ParseUint64Range(rangeStr string, defaultLeftVal uint64, defaultRightVal uint64) (*Uint64Range, error) {
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

	var v1, v2 uint64
	if strings.Trim(fields[0], " ") == "" {
		v1 = defaultLeftVal
	} else {
		v, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return nil, errors.New("Left value parse failed!")
		}
		v1 = v
	}
	if strings.Trim(fields[1], " ") == "" {
		v2 = defaultRightVal
	} else {
		v, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return nil, errors.New("Right value parse failed!")
		}
		v2 = v
	}

	if v1 > v2 {
		return nil, errors.New("Left should not large than right!")
	}

	return &Uint64Range{v1, v2, startFlag == '[', endFlag == ']'}, nil
}

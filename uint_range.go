package types

import (
	"errors"
	"strconv"
	"strings"
)

type UintRange struct {
	Left        uint
	Right       uint
	LeftClosed  bool
	RightClosed bool
}

// ParseUintRange parse string and return a *UintRange.
func ParseUintRange(rangeStr string, defaultLeftVal uint, defaultRightVal uint) (*UintRange, error) {
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

	var v1, v2 uint
	if strings.Trim(fields[0], " ") == "" {
		v1 = defaultLeftVal
	} else {
		v, err := strconv.ParseUint(fields[0], 10, 0)
		if err != nil {
			return nil, errors.New("Left value parse failed!")
		}
		v1 = uint(v)
	}
	if strings.Trim(fields[1], " ") == "" {
		v2 = defaultRightVal
	} else {
		v, err := strconv.ParseUint(fields[1], 10, 0)
		if err != nil {
			return nil, errors.New("Right value parse failed!")
		}
		v2 = uint(v)
	}

	if v1 > v2 {
		return nil, errors.New("Left should not large than right!")
	}

	return &UintRange{v1, v2, startFlag == '[', endFlag == ']'}, nil
}

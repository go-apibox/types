package types

import (
	"errors"
	"strconv"
	"strings"
)

type IntRange struct {
	Left        int
	Right       int
	LeftClosed  bool
	RightClosed bool
}

// ParseIntRange parse string and return a *IntRange.
func ParseIntRange(rangeStr string, defaultLeftVal int, defaultRightVal int) (*IntRange, error) {
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

	var v1, v2 int
	if strings.Trim(fields[0], " ") == "" {
		v1 = defaultLeftVal
	} else {
		v, err := strconv.ParseInt(fields[0], 10, 0)
		if err != nil {
			return nil, errors.New("Left value parse failed!")
		}
		v1 = int(v)
	}
	if strings.Trim(fields[1], " ") == "" {
		v2 = defaultRightVal
	} else {
		v, err := strconv.ParseInt(fields[1], 10, 0)
		if err != nil {
			return nil, errors.New("Right value parse failed!")
		}
		v2 = int(v)
	}

	if v1 > v2 {
		return nil, errors.New("Left should not large than right!")
	}

	return &IntRange{v1, v2, startFlag == '[', endFlag == ']'}, nil
}

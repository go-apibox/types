package types

import (
	"errors"
	"strconv"
	"strings"
)

type Int32Range struct {
	Left        int32
	Right       int32
	LeftClosed  bool
	RightClosed bool
}

// ParseInt32Range parse string and return a *Int32Range.
func ParseInt32Range(rangeStr string, defaultLeftVal int32, defaultRightVal int32) (*Int32Range, error) {
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

	var v1, v2 int32
	if strings.Trim(fields[0], " ") == "" {
		v1 = defaultLeftVal
	} else {
		v, err := strconv.ParseInt(fields[0], 10, 32)
		if err != nil {
			return nil, errors.New("Left value parse failed!")
		}
		v1 = int32(v)
	}
	if strings.Trim(fields[1], " ") == "" {
		v2 = defaultRightVal
	} else {
		v, err := strconv.ParseInt(fields[1], 10, 32)
		if err != nil {
			return nil, errors.New("Right value parse failed!")
		}
		v2 = int32(v)
	}

	if v1 > v2 {
		return nil, errors.New("Left should not large than right!")
	}

	return &Int32Range{v1, v2, startFlag == '[', endFlag == ']'}, nil
}

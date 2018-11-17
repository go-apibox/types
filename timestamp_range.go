package types

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type TimestampRange struct {
	Left        uint32
	Right       uint32
	LeftClosed  bool
	RightClosed bool
}

// ParseTimestampRange parse string and return a *TimestampRange.
func ParseTimestampRange(rangeStr string, defaultLeftVal uint32, defaultRightVal uint32) (*TimestampRange, error) {
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

	var v1, v2 uint32
	if strings.Trim(fields[0], " ") == "" {
		v1 = defaultLeftVal
	} else {
		v, err := strconv.ParseUint(fields[0], 10, 0)
		if err != nil {
			return nil, errors.New("Left value parse failed!")
		}
		if v > math.MaxUint32 {
			return nil, errors.New("Left value too large!")
		}
		v1 = uint32(v)
	}
	if strings.Trim(fields[1], " ") == "" {
		v2 = defaultRightVal
	} else {
		v, err := strconv.ParseUint(fields[1], 10, 0)
		if err != nil {
			return nil, errors.New("Right value parse failed!")
		}
		if v > math.MaxUint32 {
			return nil, errors.New("Right value too large!")
		}
		v2 = uint32(v)
	}

	if v1 > v2 {
		return nil, errors.New("Left should not large than right!")
	}

	return &TimestampRange{v1, v2, startFlag == '[', endFlag == ']'}, nil
}

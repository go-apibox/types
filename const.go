package types

// REF: http://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

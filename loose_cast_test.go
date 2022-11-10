package cast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/toolshelf/cast/internal"
	"html/template"
	"math"
	"testing"
	"time"
)

func TestDirectToInt(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{float64(internal.MaxSafeInteger64), internal.MaxSafeInteger64},
		{float64(internal.MinSafeInteger64), internal.MinSafeInteger64},
		{8.8, 0},
	}

	for _, test := range tests {
		v := DirectToInt(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToInt8(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int8
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{int64(math.MaxInt8), math.MaxInt8},
		{int64(math.MaxInt8 + 1), 0},
		{int64(math.MinInt8), math.MinInt8},
		{int64(math.MinInt8 - 1), 0},
		{8.8, 0},
	}

	for _, test := range tests {
		v := DirectToInt8(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToInt16(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int16
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{int64(math.MaxInt16), math.MaxInt16},
		{int64(math.MaxInt16 + 1), 0},
		{int64(math.MinInt16), math.MinInt16},
		{int64(math.MinInt16 - 1), 0},
		{8.8, 0},
	}

	for _, test := range tests {
		v := DirectToInt16(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToInt32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int32
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{int64(math.MaxInt32), math.MaxInt32},
		{int64(math.MaxInt32 + 1), 0},
		{int64(math.MinInt32), math.MinInt32},
		{float64(internal.MaxSafeInteger64), 0},
		{float64(internal.MinSafeInteger64), 0},
		{8.8, 0},
	}

	for _, test := range tests {
		v := DirectToInt32(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToInt64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int64
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{8.8, 0},
		{uint64(math.MaxInt64 + 1), 0},
		{float64(internal.MaxSafeInteger64), int64(internal.MaxSafeInteger64)},
		{float64(internal.MaxSafeInteger64 + 1), 0},
		{float64(internal.MinSafeInteger64), int64(internal.MinSafeInteger64)},
		{float64(internal.MinSafeInteger64 - 1), 0},
	}

	for _, test := range tests {

		v := DirectToInt64(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToUint(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{8.8, 0},
		{uint64(math.MaxUint), math.MaxUint},
		{float64(internal.MaxSafeInteger64), uint(internal.MaxSafeInteger64)},
		{float64(internal.MaxSafeInteger64 + 1), 0},
	}

	for _, test := range tests {
		v := DirectToUint(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToUint8(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint8
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{8.8, 0},
		{int64(math.MaxUint8), math.MaxUint8},
		{int64(math.MaxUint8 + 1), 0},
	}

	for _, test := range tests {
		v := DirectToUint8(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToUint16(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint16
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{8.8, 0},
		{int64(math.MaxUint16), math.MaxUint16},
		{int64(math.MaxUint16 + 1), 0},
	}

	for _, test := range tests {
		v := DirectToUint16(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToUint32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint32
	}{
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{8.8, 0},
		{int64(math.MaxUint32), math.MaxUint32},
		{int64(math.MaxUint32 + 1), 0},
		{float64(internal.MaxSafeInteger64), 0},
	}

	for _, test := range tests {
		v := DirectToUint32(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToUint64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint64
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.0), 8},
		{8.8, 7},
		{uint64(math.MaxUint64), uint64(math.MaxUint64)},
		{float64(internal.MaxSafeInteger64), uint64(internal.MaxSafeInteger64)},
		{float64(internal.MaxSafeInteger64 + 1), 0},
	}

	for _, test := range tests {
		v := DirectToUint64(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToFloat32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float32
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.31), 8.31},
		{8.31, 8.31},
		{"8.31", 8.31},
		{true, 1},
		{false, 0},
		{internal.MinSafeInteger32, float32(internal.MinSafeInteger32)},
		{internal.MaxSafeInteger32, float32(internal.MaxSafeInteger32)},
		{internal.MinSafeInteger32 - 1, 999},
		{internal.MaxSafeInteger32 + 1, 999},
	}

	for _, test := range tests {
		v := DirectToFloat32(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToFloat64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float64
	}{
		{8, 8},
		{int8(8), 8},
		{int16(8), 8},
		{int32(8), 8},
		{int64(8), 8},
		{uint(8), 8},
		{uint8(8), 8},
		{uint16(8), 8},
		{uint32(8), 8},
		{uint64(8), 8},
		{float32(8.31), 8.3100004196167},
		{8.31, 8.31},
		{"8.31", 8.31},
		{true, 1},
		{false, 0},
		{internal.MinSafeInteger64, float64(internal.MinSafeInteger64)},
		{internal.MaxSafeInteger64, float64(internal.MaxSafeInteger64)},
		{internal.MinSafeInteger64 - 1, 666},
		{internal.MaxSafeInteger64 + 1, 666},
	}

	for _, test := range tests {
		v := DirectToFloat64(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToString(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect string
	}{
		{nil, ""},
		{8, "8"},
		{int8(8), "8"},
		{int16(8), "8"},
		{int32(8), "8"},
		{int64(8), "8"},
		{uint(8), "8"},
		{uint8(8), "8"},
		{uint16(8), "8"},
		{uint32(8), "8"},
		{uint64(8), "8"},
		{float32(8.31), "8.31"},
		{8.31, "8.31"},
		{true, "true"},
		{false, "false"},
		{[]byte("one time"), "one time"},
		{"one more time", "one more time"},
		{template.HTML("one time"), "one time"},
		{template.URL("https://www.baidu.com"), "https://www.baidu.com"},
		{template.JS("(1+2)"), "(1+2)"},
		{template.CSS("a"), "a"},
		{template.HTMLAttr("a"), "a"},
	}

	for _, test := range tests {
		v := DirectToString(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToTime(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect time.Time
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, time.UTC)},
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC)},
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC)},
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC)},
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC)},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC)},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC)},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, time.UTC)},
		{1482597504, time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC)},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC)},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC)},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC)},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC)},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC)},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC)},
		{"2006", time.Now().UTC()},
		{testing.T{}, time.Now().UTC()},
	}

	for _, test := range tests {
		v := DirectToTime(test.input, test.expect)
		if !assert.Equal(t, test.expect, v.UTC()) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToDuration(t *testing.T) {
	var d time.Duration = 5

	tests := []struct {
		input  interface{}
		expect time.Duration
	}{
		{time.Duration(5), d},
		{5, d},
		{int64(5), d},
		{int32(5), d},
		{int16(5), d},
		{int8(5), d},
		{uint(5), d},
		{uint64(5), d},
		{uint32(5), d},
		{uint16(5), d},
		{uint8(5), d},
		{float64(5), d},
		{float32(5), d},
		{"5", d},
		{"5ns", d},
		{"5us", time.Microsecond * d},
		{"5µs", time.Microsecond * d},
		{"5ms", time.Millisecond * d},
		{"5s", time.Second * d},
		{"5m", time.Minute * d},
		{"5h", time.Hour * d},
		{"test", 999},
		{testing.T{}, 999},
	}

	for _, test := range tests {
		v := DirectToDuration(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestDirectToBool(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect bool
	}{
		{0, false},
		{nil, false},
		{"false", false},
		{"FALSE", false},
		{"False", false},
		{"f", false},
		{"F", false},
		{false, false},
		{"true", true},
		{"TRUE", true},
		{"True", true},
		{"t", true},
		{"T", true},
		{1, true},
		{1.1, true},
		{true, true},
		{-1, true},
		{"test", true},
		{testing.T{}, true},
	}

	for _, test := range tests {
		v := DirectToBool(test.input, test.expect)
		if !assert.Equal(t, test.expect, v) {
			fmt.Printf("%#v\n", test)
		}
	}
}

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

func TestMustToInt(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  int
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{float64(internal.MaxSafeInteger64), internal.MaxSafeInteger64, false},
		{float64(internal.MinSafeInteger64), internal.MinSafeInteger64, false},
		{8.8, 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToInt(test.input) })
		} else {
			v := MustToInt(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToInt8(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  int8
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{int64(math.MaxInt8), math.MaxInt8, false},
		{int64(math.MaxInt8 + 1), 0, true},
		{int64(math.MinInt8), math.MinInt8, false},
		{int64(math.MinInt8 - 1), 0, true},
		{8.8, 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToInt8(test.input) })
		} else {
			v := MustToInt8(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToInt16(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  int16
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{int64(math.MaxInt16), math.MaxInt16, false},
		{int64(math.MaxInt16 + 1), 0, true},
		{int64(math.MinInt16), math.MinInt16, false},
		{int64(math.MinInt16 - 1), 0, true},
		{8.8, 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToInt16(test.input) })
		} else {
			v := MustToInt16(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToInt32(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  int32
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{int64(math.MaxInt32), math.MaxInt32, false},
		{int64(math.MaxInt32 + 1), 0, true},
		{int64(math.MinInt32), math.MinInt32, false},
		{float64(internal.MaxSafeInteger64), 0, true},
		{float64(internal.MinSafeInteger64), 0, true},
		{8.8, 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToInt32(test.input) })
		} else {
			v := MustToInt32(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToInt64(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  int64
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{8.8, 0, true},
		{uint64(math.MaxInt64 + 1), 0, true},
		{float64(internal.MaxSafeInteger64), int64(internal.MaxSafeInteger64), false},
		{float64(internal.MaxSafeInteger64 + 1), 0, true},
		{float64(internal.MinSafeInteger64), int64(internal.MinSafeInteger64), false},
		{float64(internal.MinSafeInteger64 - 1), 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToInt64(test.input) })
		} else {
			v := MustToInt64(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToUint(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  uint
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{8.8, 0, true},
		{uint64(math.MaxUint), math.MaxUint, false},
		{float64(internal.MaxSafeInteger64), uint(internal.MaxSafeInteger64), false},
		{float64(internal.MaxSafeInteger64 + 1), 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToUint(test.input) })
		} else {
			v := MustToUint(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToUint8(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  uint8
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{8.8, 0, true},
		{int64(math.MaxUint8), math.MaxUint8, false},
		{int64(math.MaxUint8 + 1), 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToUint8(test.input) })
		} else {
			v := MustToUint8(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToUint16(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  uint16
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{8.8, 0, true},
		{int64(math.MaxUint16), math.MaxUint16, false},
		{int64(math.MaxUint16 + 1), 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToUint16(test.input) })
		} else {
			v := MustToUint16(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToUint32(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  uint32
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{8.8, 0, true},
		{int64(math.MaxUint32), math.MaxUint32, false},
		{int64(math.MaxUint32 + 1), 0, true},
		{float64(internal.MaxSafeInteger64), 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToUint32(test.input) })
		} else {
			v := MustToUint32(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToUint64(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  uint64
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.0), 8, false},
		{8.8, 0, true},
		{uint64(math.MaxUint64), uint64(math.MaxUint64), false},
		{float64(internal.MaxSafeInteger64), uint64(internal.MaxSafeInteger64), false},
		{float64(internal.MaxSafeInteger64 + 1), 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToUint64(test.input) })
		} else {
			v := MustToUint64(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToFloat32(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  float32
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8.31, false},
		{8.31, 8.31, false},
		{"8.31", 8.31, false},
		{true, 1, false},
		{false, 0, false},
		{internal.MinSafeInteger32, float32(internal.MinSafeInteger32), false},
		{internal.MaxSafeInteger32, float32(internal.MaxSafeInteger32), false},
		{internal.MinSafeInteger32 - 1, 0, true},
		{internal.MaxSafeInteger32 + 1, 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToFloat32(test.input) })
		} else {
			v := MustToFloat32(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToFloat64(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  float64
		isPanic bool
	}{
		{8, 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8.3100004196167, false},
		{8.31, 8.31, false},
		{"8.31", 8.31, false},
		{true, 1, false},
		{false, 0, false},
		{internal.MinSafeInteger64, float64(internal.MinSafeInteger64), false},
		{internal.MaxSafeInteger64, float64(internal.MaxSafeInteger64), false},
		{internal.MinSafeInteger64 - 1, 0, true},
		{internal.MaxSafeInteger64 + 1, 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToFloat64(test.input) })
		} else {
			v := MustToFloat64(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToString(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  string
		isPanic bool
	}{
		{nil, "", true},
		{8, "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{8.31, "8.31", false},
		{true, "true", false},
		{false, "false", false},
		{[]byte("one time"), "one time", false},
		{"one more time", "one more time", false},
		{template.HTML("one time"), "one time", false},
		{template.URL("https://www.baidu.com"), "https://www.baidu.com", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("a"), "a", false},
		{template.HTMLAttr("a"), "a", false},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToString(test.input) })
		} else {
			v := MustToString(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToTime(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  time.Time
		isPanic bool
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, time.UTC), false},
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC), false},
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, time.UTC), false},
		{1482597504, time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{"2006", time.Time{}, true},
		{testing.T{}, time.Time{}, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToTime(test.input) })
		} else {
			v := MustToTime(test.input)
			if !assert.Equal(t, test.expect, v.UTC()) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToDuration(t *testing.T) {
	var d time.Duration = 5

	tests := []struct {
		input   interface{}
		expect  time.Duration
		isPanic bool
	}{
		{time.Duration(5), d, false},
		{5, d, false},
		{int64(5), d, false},
		{int32(5), d, false},
		{int16(5), d, false},
		{int8(5), d, false},
		{uint(5), d, false},
		{uint64(5), d, false},
		{uint32(5), d, false},
		{uint16(5), d, false},
		{uint8(5), d, false},
		{float64(5), d, false},
		{float32(5), d, false},
		{"5", d, false},
		{"5ns", d, false},
		{"5us", time.Microsecond * d, false},
		{"5µs", time.Microsecond * d, false},
		{"5ms", time.Millisecond * d, false},
		{"5s", time.Second * d, false},
		{"5m", time.Minute * d, false},
		{"5h", time.Hour * d, false},
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToDuration(test.input) })
		} else {
			v := MustToDuration(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToBool(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  bool
		isPanic bool
	}{
		{0, false, false},
		{nil, false, true},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},
		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{1.1, true, false},
		{true, true, false},
		{-1, true, false},
		{"test", false, true},
		{testing.T{}, false, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToBool(test.input) })
		} else {
			v := MustToBool(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToStringMapString(t *testing.T) {
	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[interface{}]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[interface{}]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var jsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}`
	var invalidJsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"`
	var emptyString = ""
	var expect = map[string]string{}

	tests := []struct {
		input   interface{}
		expect  map[string]string
		isPanic bool
	}{
		{stringMapString, stringMapString, false},
		{stringMapInterface, stringMapString, false},
		{interfaceMapString, stringMapString, false},
		{interfaceMapInterface, stringMapString, false},
		{jsonString, stringMapString, false},

		{nil, expect, true},
		{testing.T{}, expect, true},
		{invalidJsonString, expect, true},
		{emptyString, expect, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToStringMapString(test.input) })
		} else {
			v := MustToStringMapString(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

func TestMustToStringMapAny(t *testing.T) {
	tests := []struct {
		input   interface{}
		expect  map[string]interface{}
		isPanic bool
	}{
		{map[interface{}]interface{}{"tag": "tags", "group": "groups"}, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{map[string]interface{}{"tag": "tags", "group": "groups"}, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": "groups"}`, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": true}`, map[string]interface{}{"tag": "tags", "group": true}, false},

		{nil, map[string]interface{}{}, true},
		{testing.T{}, map[string]interface{}{}, true},
		{"", map[string]interface{}{}, true},
	}

	for _, test := range tests {
		if test.isPanic {
			assert.Panics(t, func() { MustToStringMapAny(test.input) })
		} else {
			v := MustToStringMapAny(test.input)
			if !assert.Equal(t, test.expect, v) {
				fmt.Printf("%#v\n", test)
			}
		}
	}
}

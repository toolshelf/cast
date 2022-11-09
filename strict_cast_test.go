package cast

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/toolshelf/cast/internal"
	"html/template"
	"math"
	"testing"
	"time"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"123abc", 0, internal.ErrSyntax},
		{"9007199254740992.0", 0, internal.ErrSafe},
		{"9223372036854775807.0", 0, internal.ErrSafe},
		{"9223372036854775808", 0, internal.ErrRange},
		{"1.798E308", 0, internal.ErrRange},
		{8.8, 0, internal.ErrSyntax},
		{uint64(math.MaxInt + 1), 0, internal.ErrRange},
		{float64(internal.MaxSafeInteger64), internal.MaxSafeInteger64, nil},
		{float64(internal.MaxSafeInteger64 + 1), 0, internal.ErrSafe},
		{float64(internal.MinSafeInteger64), internal.MinSafeInteger64, nil},
		{float64(internal.MinSafeInteger64 - 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToInt(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestToInt8(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int8
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"123abc", 0, internal.ErrSyntax},
		{"128", 0, internal.ErrRange},
		{"9007199254740992.0", 0, internal.ErrSafe},
		{"9223372036854775807.0", 0, internal.ErrSafe},
		{"9223372036854775808", 0, internal.ErrRange},
		{"1.798E308", 0, internal.ErrRange},
		{8.8, 0, internal.ErrSyntax},
		{int64(math.MaxInt8), math.MaxInt8, nil},
		{int64(math.MaxInt8 + 1), 0, internal.ErrRange},
		{int64(math.MinInt8), math.MinInt8, nil},
		{int64(math.MinInt8 - 1), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToInt8(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestToInt16(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int16
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"123abc", 0, internal.ErrSyntax},
		{"32768", 0, internal.ErrRange},
		{8.8, 0, internal.ErrSyntax},
		{int64(math.MaxInt16), math.MaxInt16, nil},
		{int64(math.MaxInt16 + 1), 0, internal.ErrRange},
		{int64(math.MinInt16), math.MinInt16, nil},
		{int64(math.MinInt16 - 1), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToInt16(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestToInt32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int32
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.8, 0, internal.ErrSyntax},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"9223372036854775808", 0, internal.ErrRange},
		{"123abc", 0, internal.ErrSyntax},
		{8.8, 0, internal.ErrSyntax},
		{int64(math.MaxInt32), math.MaxInt32, nil},
		{int64(math.MaxInt32 + 1), 0, internal.ErrRange},
		{int64(math.MinInt32), math.MinInt32, nil},
		{float64(internal.MaxSafeInteger64), 0, internal.ErrRange},
		{float64(internal.MinSafeInteger64), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToInt32(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int64
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"9223372036854775808", 0, internal.ErrRange},
		{"-9223372036854775809", 0, internal.ErrRange},
		{"123abc", 0, internal.ErrSyntax},
		{8.8, 0, internal.ErrSyntax},
		{uint64(math.MaxInt64 + 1), 0, internal.ErrRange},
		{float64(internal.MaxSafeInteger64), int64(internal.MaxSafeInteger64), nil},
		{float64(internal.MaxSafeInteger64 + 1), 0, internal.ErrSafe},
		{float64(internal.MinSafeInteger64), int64(internal.MinSafeInteger64), nil},
		{float64(internal.MinSafeInteger64 - 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToInt64(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestToUint(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{-8, 0, internal.ErrRange},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"123abc", 0, internal.ErrSyntax},
		{"9007199254740992.0", 0, internal.ErrSafe},
		{"9223372036854775807.0", 0, internal.ErrSafe},
		{"18446744073709551616", 0, internal.ErrRange},
		{"1.798E308", 0, internal.ErrRange},
		{8.8, 0, internal.ErrSyntax},
		{uint64(math.MaxUint), math.MaxUint, nil},
		{float64(internal.MaxSafeInteger64), uint(internal.MaxSafeInteger64), nil},
		{float64(internal.MaxSafeInteger64 + 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToUint(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestToUint8(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint8
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{-8, 0, internal.ErrRange},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"123abc", 0, internal.ErrSyntax},
		{"256", 0, internal.ErrRange},
		{"9007199254740992.0", 0, internal.ErrSafe},
		{"9223372036854775807.0", 0, internal.ErrSafe},
		{"9223372036854775808", 0, internal.ErrRange},
		{"1.798E308", 0, internal.ErrRange},
		{8.8, 0, internal.ErrSyntax},
		{int64(math.MaxUint8), math.MaxUint8, nil},
		{int64(math.MaxUint8 + 1), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToUint8(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestToUint16(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint16
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{-8, 0, internal.ErrRange},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"123abc", 0, internal.ErrSyntax},
		{"65536", 0, internal.ErrRange},
		{8.8, 0, internal.ErrSyntax},
		{int64(math.MaxUint16), math.MaxUint16, nil},
		{int64(math.MaxUint16 + 1), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToUint16(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestToUint32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint32
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{-8, 0, internal.ErrRange},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"9223372036854775808", 0, internal.ErrRange},
		{"123abc", 0, internal.ErrSyntax},
		{8.8, 0, internal.ErrSyntax},
		{int64(math.MaxUint32), math.MaxUint32, nil},
		{int64(math.MaxUint32 + 1), 0, internal.ErrRange},
		{float64(internal.MaxSafeInteger64), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToUint32(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestToUint64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint64
		err    error
	}{
		{nil, 0, internal.ErrSyntax},
		{8, 8, nil},
		{-8, 0, internal.ErrRange},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.0), 8, nil},
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"9223372036854775808", 0, internal.ErrRange},
		{"-9223372036854775809", 0, internal.ErrRange},
		{"123abc", 0, internal.ErrSyntax},
		{8.8, 0, internal.ErrSyntax},
		{uint64(math.MaxUint64), uint64(math.MaxUint64), nil},
		{float64(internal.MaxSafeInteger64), uint64(internal.MaxSafeInteger64), nil},
		{float64(internal.MaxSafeInteger64 + 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToUint64(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestToFloat32(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float32
		err    error
	}{
		{8, 8, nil},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.31), 8.31, nil},
		{8.31, 8.31, nil},
		{"8.31", 8.31, nil},
		{true, 1, nil},
		{false, 0, nil},
		{internal.MinSafeInteger32, float32(internal.MinSafeInteger32), nil},
		{internal.MaxSafeInteger32, float32(internal.MaxSafeInteger32), nil},
		{internal.MinSafeInteger32 - 1, 0, internal.ErrSafe},
		{internal.MaxSafeInteger32 + 1, 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToFloat32(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float64
		err    error
	}{
		{8, 8, nil},
		{int8(8), 8, nil},
		{int16(8), 8, nil},
		{int32(8), 8, nil},
		{int64(8), 8, nil},
		{uint(8), 8, nil},
		{uint8(8), 8, nil},
		{uint16(8), 8, nil},
		{uint32(8), 8, nil},
		{uint64(8), 8, nil},
		{float32(8.31), 8.3100004196167, nil},
		{8.31, 8.31, nil},
		{"8.31", 8.31, nil},
		{true, 1, nil},
		{false, 0, nil},
		{internal.MinSafeInteger64, float64(internal.MinSafeInteger64), nil},
		{internal.MaxSafeInteger64, float64(internal.MaxSafeInteger64), nil},
		{internal.MinSafeInteger64 - 1, 0, internal.ErrSafe},
		{internal.MaxSafeInteger64 + 1, 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToFloat64(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect string
		err    error
	}{
		{nil, "", internal.ErrSyntax},
		{8, "8", nil},
		{int8(8), "8", nil},
		{int16(8), "8", nil},
		{int32(8), "8", nil},
		{int64(8), "8", nil},
		{uint(8), "8", nil},
		{uint8(8), "8", nil},
		{uint16(8), "8", nil},
		{uint32(8), "8", nil},
		{uint64(8), "8", nil},
		{float32(8.31), "8.31", nil},
		{8.31, "8.31", nil},
		{true, "true", nil},
		{false, "false", nil},
		{[]byte("one time"), "one time", nil},
		{"one more time", "one more time", nil},
		{template.HTML("one time"), "one time", nil},
		{template.URL("https://www.baidu.com"), "https://www.baidu.com", nil},
		{template.JS("(1+2)"), "(1+2)", nil},
		{template.CSS("a"), "a", nil},
		{template.HTMLAttr("a"), "a", nil},
	}

	for _, test := range tests {
		v, err := ToString(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestToTime(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect time.Time
		err    error
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},   // Time.String()
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},        // ANSIC
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},    // UnixDate
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},  // RubyDate
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},             // RFC822
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},           // RFC822Z
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil}, // RFC850
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},   // RFC1123
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil}, // RFC1123Z
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},            // RFC3339
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, time.UTC), nil},      // RFC3339 without timezone hh:mm colon
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), nil},            // RFC3339Nano
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC), nil},                              // Kitchen
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), nil},                    // Stamp
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), nil},                // StampMilli
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), nil},             // StampMicro
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), nil},          // StampNano
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), nil},        // RFC3339 without T
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), nil},         // RFC3339 without T or timezone hh:mm colon
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), nil},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), nil},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), nil},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), nil},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), nil},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, time.UTC), nil},
		{int(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), nil},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), nil},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), nil},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), nil},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), nil},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), nil},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), nil},
		// errors
		{"2006", time.Time{}, internal.ErrSyntax},
		{testing.T{}, time.Time{}, internal.ErrSyntax},
	}

	for _, test := range tests {
		v, err := ToTime(test.input)
		if !assert.Equal(t, test.expect, v.UTC()) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestToDuration(t *testing.T) {
	var d time.Duration = 5

	tests := []struct {
		input  interface{}
		expect time.Duration
		err    error
	}{
		{time.Duration(5), d, nil},
		{int(5), d, nil},
		{int64(5), d, nil},
		{int32(5), d, nil},
		{int16(5), d, nil},
		{int8(5), d, nil},
		{uint(5), d, nil},
		{uint64(5), d, nil},
		{uint32(5), d, nil},
		{uint16(5), d, nil},
		{uint8(5), d, nil},
		{float64(5), d, nil},
		{float32(5), d, nil},
		{string("5"), d, nil},
		{string("5ns"), d, nil},
		{string("5us"), time.Microsecond * d, nil},
		{string("5µs"), time.Microsecond * d, nil},
		{string("5ms"), time.Millisecond * d, nil},
		{string("5s"), time.Second * d, nil},
		{string("5m"), time.Minute * d, nil},
		{string("5h"), time.Hour * d, nil},
		// errors
		{"test", 0, internal.ErrSyntax},
		{testing.T{}, 0, internal.ErrSyntax},
	}

	for _, test := range tests {
		v, err := ToDuration(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

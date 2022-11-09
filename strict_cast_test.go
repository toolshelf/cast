package cast

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/toolshelf/cast/internal"
	"html/template"
	"math"
	"testing"
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

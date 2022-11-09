package cast

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/toolshelf/cast/internal"
	"math"
	"testing"
)

func TestStrictCastToInt(t *testing.T) {
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
		{float64(internal.MaxSafeInteger), internal.MaxSafeInteger, nil},
		{float64(internal.MaxSafeInteger + 1), 0, internal.ErrSafe},
		{float64(internal.MinSafeInteger), internal.MinSafeInteger, nil},
		{float64(internal.MinSafeInteger - 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToInt(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestStrictCastToInt8(t *testing.T) {
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

func TestStrictCastToInt16(t *testing.T) {
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

func TestStrictCastToInt32(t *testing.T) {
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
		{8.0, 8, nil},
		{"8", 8, nil},
		{"8.8", 0, internal.ErrSyntax},
		{"9223372036854775808", 0, internal.ErrRange},
		{"123abc", 0, internal.ErrSyntax},
		{8.8, 0, internal.ErrSyntax},
		{int64(math.MaxInt32), math.MaxInt32, nil},
		{int64(math.MaxInt32 + 1), 0, internal.ErrRange},
		{int64(math.MinInt32), math.MinInt32, nil},
		{float64(internal.MaxSafeInteger), 0, internal.ErrRange},
		{float64(internal.MinSafeInteger), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToInt32(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestStrictCastToInt64(t *testing.T) {
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
		{float64(internal.MaxSafeInteger), int64(internal.MaxSafeInteger), nil},
		{float64(internal.MaxSafeInteger + 1), 0, internal.ErrSafe},
		{float64(internal.MinSafeInteger), int64(internal.MinSafeInteger), nil},
		{float64(internal.MinSafeInteger - 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToInt64(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}

func TestStrictCastToUint(t *testing.T) {
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
		{float64(internal.MaxSafeInteger), uint(internal.MaxSafeInteger), nil},
		{float64(internal.MaxSafeInteger + 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToUint(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestStrictCastToUint8(t *testing.T) {
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

func TestStrictCastToUint16(t *testing.T) {
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

func TestStrictCastToUint32(t *testing.T) {
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
		{float64(internal.MaxSafeInteger), 0, internal.ErrRange},
	}

	for _, test := range tests {
		v, err := ToUint32(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestStrictCastToUint64(t *testing.T) {
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
		{float64(internal.MaxSafeInteger), uint64(internal.MaxSafeInteger), nil},
		{float64(internal.MaxSafeInteger + 1), 0, internal.ErrSafe},
	}

	for _, test := range tests {
		v, err := ToUint64(test.input)
		if !assert.Equal(t, test.expect, v) || !assert.True(t, errors.Is(err, test.err)) {
			fmt.Printf("%#v\n", test)
			fmt.Println(v, err)
		}
	}
}
